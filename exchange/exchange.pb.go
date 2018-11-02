// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/exchange.proto

package exchange

import (
	fmt "fmt"
	data "github.com/airbloc/airbloc-go/api/data"
	common "github.com/airbloc/airbloc-go/common"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Contract_Type int32

const (
	Contract_RICHARDIAN Contract_Type = 0
	Contract_SMART      Contract_Type = 1
)

var Contract_Type_name = map[int32]string{
	0: "RICHARDIAN",
	1: "SMART",
}

var Contract_Type_value = map[string]int32{
	"RICHARDIAN": 0,
	"SMART":      1,
}

func (x Contract_Type) String() string {
	return proto.EnumName(Contract_Type_name, int32(x))
}

func (Contract_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{0, 0}
}

type Contract struct {
	Type                 Contract_Type   `protobuf:"varint,1,opt,name=type,proto3,enum=airbloc.exchange.Contract_Type" json:"type,omitempty"`
	SmartContractAddress *common.Address `protobuf:"bytes,2,opt,name=smartContractAddress,proto3" json:"smartContractAddress,omitempty"`
	RichardianHash       []byte          `protobuf:"bytes,3,opt,name=richardianHash,proto3" json:"richardianHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{0}
}

func (m *Contract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contract.Unmarshal(m, b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return xxx_messageInfo_Contract.Size(m)
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetType() Contract_Type {
	if m != nil {
		return m.Type
	}
	return Contract_RICHARDIAN
}

func (m *Contract) GetSmartContractAddress() *common.Address {
	if m != nil {
		return m.SmartContractAddress
	}
	return nil
}

func (m *Contract) GetRichardianHash() []byte {
	if m != nil {
		return m.RichardianHash
	}
	return nil
}

type OrderRequest struct {
	Data                 *data.BatchRequest `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Contract             *Contract          `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Offeror              string             `protobuf:"bytes,3,opt,name=offeror,proto3" json:"offeror,omitempty"`
	Offeree              string             `protobuf:"bytes,4,opt,name=offeree,proto3" json:"offeree,omitempty"`
	Option               []string           `protobuf:"bytes,5,rep,name=option,proto3" json:"option,omitempty"`
	Amount               float64            `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{1}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetData() *data.BatchRequest {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *OrderRequest) GetContract() *Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *OrderRequest) GetOfferor() string {
	if m != nil {
		return m.Offeror
	}
	return ""
}

func (m *OrderRequest) GetOfferee() string {
	if m != nil {
		return m.Offeree
	}
	return ""
}

func (m *OrderRequest) GetOption() []string {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *OrderRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type OrderId struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderId) Reset()         { *m = OrderId{} }
func (m *OrderId) String() string { return proto.CompactTextString(m) }
func (*OrderId) ProtoMessage()    {}
func (*OrderId) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{2}
}

func (m *OrderId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderId.Unmarshal(m, b)
}
func (m *OrderId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderId.Marshal(b, m, deterministic)
}
func (m *OrderId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderId.Merge(m, src)
}
func (m *OrderId) XXX_Size() int {
	return xxx_messageInfo_OrderId.Size(m)
}
func (m *OrderId) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderId.DiscardUnknown(m)
}

var xxx_messageInfo_OrderId proto.InternalMessageInfo

func (m *OrderId) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type SettleMessage struct {
	OrderId              *OrderId `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleMessage) Reset()         { *m = SettleMessage{} }
func (m *SettleMessage) String() string { return proto.CompactTextString(m) }
func (*SettleMessage) ProtoMessage()    {}
func (*SettleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{3}
}

func (m *SettleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleMessage.Unmarshal(m, b)
}
func (m *SettleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleMessage.Marshal(b, m, deterministic)
}
func (m *SettleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleMessage.Merge(m, src)
}
func (m *SettleMessage) XXX_Size() int {
	return xxx_messageInfo_SettleMessage.Size(m)
}
func (m *SettleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SettleMessage proto.InternalMessageInfo

func (m *SettleMessage) GetOrderId() *OrderId {
	if m != nil {
		return m.OrderId
	}
	return nil
}

type ContractId struct {
	ContractId           string   `protobuf:"bytes,1,opt,name=ContractId,proto3" json:"ContractId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractId) Reset()         { *m = ContractId{} }
func (m *ContractId) String() string { return proto.CompactTextString(m) }
func (*ContractId) ProtoMessage()    {}
func (*ContractId) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{4}
}

func (m *ContractId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractId.Unmarshal(m, b)
}
func (m *ContractId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractId.Marshal(b, m, deterministic)
}
func (m *ContractId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractId.Merge(m, src)
}
func (m *ContractId) XXX_Size() int {
	return xxx_messageInfo_ContractId.Size(m)
}
func (m *ContractId) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractId.DiscardUnknown(m)
}

var xxx_messageInfo_ContractId proto.InternalMessageInfo

func (m *ContractId) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

type SettleResult struct {
	ContractId           *ContractId `protobuf:"bytes,1,opt,name=contractId,proto3" json:"contractId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SettleResult) Reset()         { *m = SettleResult{} }
func (m *SettleResult) String() string { return proto.CompactTextString(m) }
func (*SettleResult) ProtoMessage()    {}
func (*SettleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{5}
}

func (m *SettleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleResult.Unmarshal(m, b)
}
func (m *SettleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleResult.Marshal(b, m, deterministic)
}
func (m *SettleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleResult.Merge(m, src)
}
func (m *SettleResult) XXX_Size() int {
	return xxx_messageInfo_SettleResult.Size(m)
}
func (m *SettleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleResult.DiscardUnknown(m)
}

var xxx_messageInfo_SettleResult proto.InternalMessageInfo

func (m *SettleResult) GetContractId() *ContractId {
	if m != nil {
		return m.ContractId
	}
	return nil
}

func init() {
	proto.RegisterEnum("airbloc.exchange.Contract_Type", Contract_Type_name, Contract_Type_value)
	proto.RegisterType((*Contract)(nil), "airbloc.exchange.Contract")
	proto.RegisterType((*OrderRequest)(nil), "airbloc.exchange.OrderRequest")
	proto.RegisterType((*OrderId)(nil), "airbloc.exchange.OrderId")
	proto.RegisterType((*SettleMessage)(nil), "airbloc.exchange.SettleMessage")
	proto.RegisterType((*ContractId)(nil), "airbloc.exchange.ContractId")
	proto.RegisterType((*SettleResult)(nil), "airbloc.exchange.SettleResult")
}

func init() { proto.RegisterFile("proto/exchange.proto", fileDescriptor_95588e2623cfabeb) }

var fileDescriptor_95588e2623cfabeb = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x9b, 0xb8, 0xcd, 0x24, 0x44, 0x61, 0x55, 0x15, 0x63, 0xa1, 0x60, 0x8c, 0x54,
	0xf9, 0x00, 0xae, 0xe4, 0x48, 0x1c, 0x80, 0x4b, 0x92, 0x22, 0xd5, 0x82, 0x82, 0xb4, 0xed, 0x89,
	0xdb, 0xc6, 0x9e, 0xc6, 0x46, 0x89, 0x37, 0xec, 0x6e, 0x24, 0xfa, 0x1a, 0xbc, 0x1c, 0x12, 0x4f,
	0x83, 0xbc, 0x5e, 0xa7, 0x49, 0xa8, 0x2f, 0x9c, 0x92, 0xf9, 0xe7, 0xdb, 0xd9, 0xff, 0x9f, 0x38,
	0x86, 0x93, 0x95, 0xe0, 0x8a, 0x9f, 0xe3, 0xcf, 0x24, 0x63, 0xc5, 0x1c, 0x43, 0x5d, 0x92, 0x01,
	0xcb, 0xc5, 0x6c, 0xc1, 0x93, 0xb0, 0xd6, 0xdd, 0x41, 0xc5, 0xa5, 0x4c, 0xb1, 0x8a, 0x71, 0x9f,
	0x54, 0x8a, 0xba, 0x5b, 0xa1, 0xac, 0x24, 0xff, 0xb7, 0x05, 0xc7, 0x53, 0x5e, 0x28, 0xc1, 0x12,
	0x45, 0x46, 0xd0, 0x2a, 0x7b, 0x8e, 0xe5, 0x59, 0x41, 0x3f, 0x7a, 0x11, 0xee, 0x8f, 0x0c, 0x6b,
	0x32, 0xbc, 0xb9, 0x5b, 0x21, 0xd5, 0x30, 0xf9, 0x04, 0x27, 0x72, 0xc9, 0x84, 0xaa, 0x7b, 0xe3,
	0x34, 0x15, 0x28, 0xa5, 0x73, 0xe0, 0x59, 0x41, 0x37, 0x7a, 0xba, 0x19, 0x92, 0xf0, 0xe5, 0x92,
	0x17, 0xa1, 0x69, 0xd3, 0x07, 0x0f, 0x91, 0x33, 0xe8, 0x8b, 0x3c, 0xc9, 0x98, 0x48, 0x73, 0x56,
	0x5c, 0x32, 0x99, 0x39, 0x87, 0x9e, 0x15, 0xf4, 0xe8, 0x9e, 0xea, 0xbf, 0x84, 0x56, 0x69, 0x81,
	0xf4, 0x01, 0x68, 0x3c, 0xbd, 0x1c, 0xd3, 0x8b, 0x78, 0xfc, 0x65, 0xf0, 0x88, 0x74, 0xa0, 0x7d,
	0x7d, 0x35, 0xa6, 0x37, 0x03, 0xcb, 0xff, 0x63, 0x41, 0xef, 0xab, 0x48, 0x51, 0x50, 0xfc, 0xb1,
	0x46, 0xa9, 0x48, 0x08, 0xad, 0x72, 0x17, 0x3a, 0x5d, 0x37, 0x72, 0x37, 0xc6, 0xf4, 0x82, 0x26,
	0x4c, 0x25, 0x99, 0x21, 0xa9, 0xe6, 0xc8, 0x5b, 0x38, 0x4e, 0x8c, 0x3d, 0x13, 0xc6, 0x6d, 0xde,
	0x08, 0xdd, 0xb0, 0xc4, 0x81, 0x23, 0x7e, 0x7b, 0x8b, 0x82, 0x0b, 0x6d, 0xbe, 0x43, 0xeb, 0x72,
	0xd3, 0x41, 0x74, 0x5a, 0x5b, 0x1d, 0x44, 0x72, 0x0a, 0x36, 0x5f, 0xa9, 0x9c, 0x17, 0x4e, 0xdb,
	0x3b, 0x0c, 0x3a, 0xd4, 0x54, 0xa5, 0xce, 0x96, 0x7c, 0x5d, 0x28, 0xc7, 0xf6, 0xac, 0xc0, 0xa2,
	0xa6, 0xf2, 0x5f, 0xc1, 0x91, 0xce, 0x16, 0xa7, 0x7a, 0x68, 0xf5, 0x55, 0x27, 0x2b, 0x87, 0x56,
	0xa5, 0x7f, 0x01, 0x8f, 0xaf, 0x51, 0xa9, 0x05, 0x5e, 0xa1, 0x94, 0x6c, 0x8e, 0x64, 0xb4, 0x8b,
	0x76, 0xa3, 0x67, 0xff, 0x06, 0x32, 0x63, 0xef, 0xa7, 0xbc, 0x06, 0xa8, 0x43, 0xc6, 0x29, 0x19,
	0x6e, 0x57, 0xe6, 0xc2, 0x2d, 0xc5, 0xff, 0x0c, 0xbd, 0xea, 0x4e, 0x8a, 0x72, 0xbd, 0x50, 0xe4,
	0x03, 0x40, 0xb2, 0xcb, 0x77, 0xa3, 0xe7, 0xcd, 0x6b, 0x8c, 0x53, 0xba, 0xc5, 0x47, 0xbf, 0x0e,
	0xe0, 0xf8, 0xa3, 0x61, 0xc8, 0x04, 0xda, 0xda, 0x1c, 0x19, 0x36, 0xb8, 0x36, 0x3f, 0x9f, 0xdb,
	0x9c, 0x8a, 0xc4, 0x60, 0x57, 0xf6, 0xc8, 0x03, 0x4f, 0xf7, 0xce, 0xb2, 0xdc, 0x61, 0x13, 0x60,
	0x92, 0xbd, 0x07, 0x9b, 0xe2, 0x77, 0x4c, 0x14, 0x69, 0xbe, 0xcf, 0x3d, 0xdd, 0x7f, 0xfc, 0xcd,
	0xe1, 0x77, 0xd0, 0x9e, 0x2e, 0xb8, 0xc4, 0xff, 0x38, 0x3b, 0x09, 0xbe, 0x9d, 0xcd, 0x73, 0x95,
	0xad, 0x67, 0xa5, 0x7e, 0x6e, 0x98, 0xfa, 0xf3, 0xcd, 0xfc, 0xfe, 0xcd, 0x30, 0xb3, 0xf5, 0x7f,
	0x7c, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x61, 0x9f, 0x5b, 0x32, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeClient interface {
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderId, error)
	Settle(ctx context.Context, in *SettleMessage, opts ...grpc.CallOption) (*SettleResult, error)
	Reject(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error)
	// after Open()
	Close(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderId, error) {
	out := new(OrderId)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Settle(ctx context.Context, in *SettleMessage, opts ...grpc.CallOption) (*SettleResult, error) {
	out := new(SettleResult)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Settle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Reject(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error) {
	out := new(common.Result)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Reject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Close(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error) {
	out := new(common.Result)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.Exchange/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServer is the server API for Exchange service.
type ExchangeServer interface {
	Order(context.Context, *OrderRequest) (*OrderId, error)
	Settle(context.Context, *SettleMessage) (*SettleResult, error)
	Reject(context.Context, *OrderId) (*common.Result, error)
	// after Open()
	Close(context.Context, *OrderId) (*common.Result, error)
}

func RegisterExchangeServer(s *grpc.Server, srv ExchangeServer) {
	s.RegisterService(&_Exchange_serviceDesc, srv)
}

func _Exchange_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Settle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Settle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Settle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Settle(ctx, req.(*SettleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Reject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Reject(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.Exchange/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Close(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.exchange.Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Order",
			Handler:    _Exchange_Order_Handler,
		},
		{
			MethodName: "Settle",
			Handler:    _Exchange_Settle_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _Exchange_Reject_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Exchange_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/exchange.proto",
}
