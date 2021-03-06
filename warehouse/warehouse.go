package warehouse

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/azer/logger"

	"github.com/airbloc/airbloc-go/adapter"
	"github.com/airbloc/airbloc-go/blockchain"
	"github.com/airbloc/airbloc-go/data"
	"github.com/airbloc/airbloc-go/database/localdb"
	"github.com/airbloc/airbloc-go/database/metadb"
	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/key"
	"github.com/airbloc/airbloc-go/warehouse/protocol"
	"github.com/airbloc/airbloc-go/warehouse/storage"
	"github.com/pkg/errors"
)

type DataWarehouse struct {
	kms            key.Manager
	protocols      map[string]protocol.Protocol
	localCache     *localdb.Model
	metaDatabase   *metadb.Model
	ethclient      blockchain.TxClient
	dataRegistry   *adapter.DataRegistry
	DefaultStorage storage.Storage
	log            *logger.Logger
}

func New(
	kms key.Manager,
	localDatabase localdb.Database,
	metaDatabase metadb.Database,
	ethclient blockchain.TxClient,
	defaultStorage storage.Storage,
	supportedProtocols []protocol.Protocol,
) *DataWarehouse {
	protocols := map[string]protocol.Protocol{}
	for _, protoc := range supportedProtocols {
		protocols[protoc.Name()] = protoc
	}
	contract := ethclient.GetContract(&adapter.DataRegistry{})
	return &DataWarehouse{
		kms:            kms,
		protocols:      protocols,
		localCache:     localdb.NewModel(localDatabase, "bundle"),
		metaDatabase:   metadb.NewModel(metaDatabase, "bundles"),
		ethclient:      ethclient,
		dataRegistry:   contract.(*adapter.DataRegistry),
		DefaultStorage: defaultStorage,
		log:            logger.New("warehouse"),
	}
}

func (warehouse *DataWarehouse) CreateBundle(collection common.ID) *BundleStream {
	return newBundleStream(warehouse, collection)
}

func (warehouse *DataWarehouse) validate(collection common.ID, data *common.Data) error {
	return nil
}

func (warehouse *DataWarehouse) encrypt(d *common.Data) (*common.EncryptedData, error) {
	encryptedPayload, err := warehouse.kms.Encrypt(d.Payload)
	if err != nil {
		return nil, err
	}
	return &common.EncryptedData{
		OwnerAnid: d.OwnerAnid,
		Payload:   encryptedPayload,
		Capsule:   nil,
	}, nil
}

func generateBundleNameOf(bundle *data.Bundle) string {
	tokenBytes := make([]byte, 4)
	rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	currentTime := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s-%s.bundle", currentTime, bundle.Collection.String(), token)
}

func (warehouse *DataWarehouse) Store(stream *BundleStream) (*data.Bundle, error) {
	if stream == nil {
		return nil, errors.New("No data in the stream.")
	}
	ingestedAt := time.Now()

	createdBundle := &data.Bundle{
		Provider:   common.ID{}, /* TODO: implement appID */
		Collection: stream.collection,
		DataCount:  stream.DataCount,
		IngestedAt: ingestedAt,
		Data:       stream.data,
	}
	bundleName := generateBundleNameOf(createdBundle)
	uri, err := warehouse.DefaultStorage.Save(bundleName, createdBundle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save bundle to the storage")
	}
	createdBundle.Uri = uri.String()

	// register to on-chain
	bundleIndex, err := warehouse.registerBundleOnChain(createdBundle)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register bundle to blockchain")
	}
	createdBundle.Id = fmt.Sprintf("%s/%d", createdBundle.Collection.String(), bundleIndex)

	// save metadata to make the bundle searchable
	bundleInfo := map[string]interface{}{
		"bundleId":   createdBundle.Id,
		"uri":        createdBundle.Uri,
		"provider":   createdBundle.Provider.String(),
		"collection": createdBundle.Collection.String(),
		"dataCount":  createdBundle.DataCount,
		"ingestedAt": ingestedAt,
	}
	_, err = warehouse.metaDatabase.Create(bundleInfo, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save metadata")
	}
	warehouse.log.Info("Bundle %s registered on", bundleName, logger.Attrs{
		"index": bundleIndex,
		"count": bundleInfo["dataCount"],
	})
	return createdBundle, nil
}

func (warehouse *DataWarehouse) registerBundleOnChain(bundle *data.Bundle) (int, error) {
	bundleDataHash, err := bundle.Hash()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get hash of the bundle data")
	}

	userMerkleRoot, err := bundle.SetupUserProof()
	if err != nil {
		return 0, errors.Wrap(err, "failed to setup SMT")
	}

	warehouse.log.Info("Bundle data", logger.Attrs{
		"hash":       bundleDataHash.Hex(),
		"merkleRoot": userMerkleRoot.Hex(),
	})

	tx, err := warehouse.dataRegistry.RegisterBundle(
		warehouse.ethclient.Account(),
		bundle.Collection,
		userMerkleRoot,
		bundleDataHash,
		bundle.Uri)
	if err != nil {
		return 0, errors.Wrap(err, "failed to register a bundle to DataRegistry")
	}

	receipt, err := warehouse.ethclient.WaitMined(context.Background(), tx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to wait for tx to be mined")
	}

	registerResult, err := warehouse.dataRegistry.ParseBundleRegisteredFromReceipt(receipt)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse a event from the receipt")
	}
	return int(registerResult.Index), nil
}

func (warehouse *DataWarehouse) Get(bundleId string) (*data.Bundle, error) {
	// try to fetch URI from cache. TODO: TTL of the bundle cache
	uri, err := warehouse.localCache.Get(bundleId)
	if err != nil {
		warehouse.log.Error("Warning: failed to access local DB: %s", err.Error(), logger.Attrs{"uri": uri})
	}

	if uri != nil {
		if uri, err := url.Parse(string(uri)); err == nil {
			return warehouse.Fetch(uri)
		}
	}

	// search URI from metadatabase
	query := bson.NewDocument(bson.EC.String("data.data.bundleId", bundleId))
	metadata, err := warehouse.metaDatabase.RetrieveAsset(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query on metadatabase")
	}

	parsedUri, err := url.Parse(metadata.Lookup("data", "uri").StringValue())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse URI")
	}
	return warehouse.Fetch(parsedUri)
}

func (warehouse *DataWarehouse) Fetch(uri *url.URL) (*data.Bundle, error) {
	protoc, exists := warehouse.protocols[uri.Scheme]
	if !exists {
		return nil, errors.Errorf("the protocol %s is not supported", uri.Scheme)
	}
	return protoc.Read(uri)
}
