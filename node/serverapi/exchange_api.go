package serverapi

import (
	"github.com/airbloc/airbloc-go/exchange"
	"github.com/airbloc/airbloc-go/node"
	commonpb "github.com/airbloc/airbloc-go/proto/rpc/v1"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"golang.org/x/net/context"
)

type ExchangeAPI struct {
	exchange *exchange.Exchange
}

func NewExchangeAPI(backend node.Backend) (node.API, error) {
	ex := exchange.New(backend.Client())
	return &ExchangeAPI{ex}, nil
}

func (api *ExchangeAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterExchangeServer(service.GrpcServer, api)
}

// TODO
func (api *ExchangeAPI) Order(ctx context.Context, req *pb.OrderRequest) (*commonpb.Hash, error) {
	return nil, nil
}

func (api *ExchangeAPI) Settle(ctx context.Context, id *commonpb.Hash) (*commonpb.Result, error) {
	return nil, nil
}

func (api *ExchangeAPI) Reject(ctx context.Context, id *pb.OrderIdRequest) (*commonpb.Result, error) {
	return nil, nil
}

func (api *ExchangeAPI) CloseOrder(ctx context.Context, id *pb.OrderIdRequest) (*commonpb.Result, error) {
	return nil, nil
}
