// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

package exchange // import "github.com/airbloc/airbloc-go/api/exchange"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/airbloc/airbloc-go/api/common"
import data "github.com/airbloc/airbloc-go/api/data"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
	return fileDescriptor_exchange_558f21d37c5fc632, []int{0, 0}
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
	return fileDescriptor_exchange_558f21d37c5fc632, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contract.Unmarshal(m, b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
}
func (dst *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(dst, src)
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
	Data                 *data.Batch `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Contract             *Contract   `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_558f21d37c5fc632, []int{1}
}
func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (dst *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(dst, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetData() *data.Batch {
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
	return fileDescriptor_exchange_558f21d37c5fc632, []int{2}
}
func (m *OrderId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderId.Unmarshal(m, b)
}
func (m *OrderId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderId.Marshal(b, m, deterministic)
}
func (dst *OrderId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderId.Merge(dst, src)
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

type SettleResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleResult) Reset()         { *m = SettleResult{} }
func (m *SettleResult) String() string { return proto.CompactTextString(m) }
func (*SettleResult) ProtoMessage()    {}
func (*SettleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_558f21d37c5fc632, []int{3}
}
func (m *SettleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleResult.Unmarshal(m, b)
}
func (m *SettleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleResult.Marshal(b, m, deterministic)
}
func (dst *SettleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleResult.Merge(dst, src)
}
func (m *SettleResult) XXX_Size() int {
	return xxx_messageInfo_SettleResult.Size(m)
}
func (m *SettleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleResult.DiscardUnknown(m)
}

var xxx_messageInfo_SettleResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Contract)(nil), "airbloc.exchange.Contract")
	proto.RegisterType((*OrderRequest)(nil), "airbloc.exchange.OrderRequest")
	proto.RegisterType((*OrderId)(nil), "airbloc.exchange.OrderId")
	proto.RegisterType((*SettleResult)(nil), "airbloc.exchange.SettleResult")
	proto.RegisterEnum("airbloc.exchange.Contract_Type", Contract_Type_name, Contract_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeServiceClient is the client API for ExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeServiceClient interface {
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderId, error)
	Settle(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*SettleResult, error)
	Reject(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error)
}

type exchangeServiceClient struct {
	cc *grpc.ClientConn
}

func NewExchangeServiceClient(cc *grpc.ClientConn) ExchangeServiceClient {
	return &exchangeServiceClient{cc}
}

func (c *exchangeServiceClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderId, error) {
	out := new(OrderId)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.ExchangeService/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) Settle(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*SettleResult, error) {
	out := new(SettleResult)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.ExchangeService/Settle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) Reject(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*common.Result, error) {
	out := new(common.Result)
	err := c.cc.Invoke(ctx, "/airbloc.exchange.ExchangeService/Reject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServiceServer is the server API for ExchangeService service.
type ExchangeServiceServer interface {
	Order(context.Context, *OrderRequest) (*OrderId, error)
	Settle(context.Context, *OrderId) (*SettleResult, error)
	Reject(context.Context, *OrderId) (*common.Result, error)
}

func RegisterExchangeServiceServer(s *grpc.Server, srv ExchangeServiceServer) {
	s.RegisterService(&_ExchangeService_serviceDesc, srv)
}

func _ExchangeService_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.ExchangeService/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_Settle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).Settle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.ExchangeService/Settle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).Settle(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.exchange.ExchangeService/Reject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).Reject(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExchangeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.exchange.ExchangeService",
	HandlerType: (*ExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Order",
			Handler:    _ExchangeService_Order_Handler,
		},
		{
			MethodName: "Settle",
			Handler:    _ExchangeService_Settle_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _ExchangeService_Reject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exchange.proto",
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor_exchange_558f21d37c5fc632) }

var fileDescriptor_exchange_558f21d37c5fc632 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x7d, 0x7d, 0x8f, 0xcf, 0x0b, 0xe9, 0x23, 0xa3, 0xd1, 0xda, 0x85, 0x62, 0x4d, 0x94, 0x18,
	0x2d, 0x49, 0x49, 0xdc, 0xb8, 0x2a, 0x68, 0x42, 0x63, 0xd4, 0x64, 0x60, 0xe5, 0x6e, 0x98, 0x4e,
	0x68, 0x0d, 0x74, 0xea, 0x74, 0x30, 0xf2, 0x4b, 0xf9, 0x3b, 0xa6, 0xd3, 0x96, 0x10, 0x04, 0x57,
	0xed, 0xbd, 0xe7, 0xdc, 0x73, 0xce, 0xbd, 0x2d, 0xe8, 0xec, 0x8b, 0x06, 0x24, 0x9a, 0x32, 0x3b,
	0x16, 0x5c, 0x72, 0xd4, 0x22, 0xa1, 0x98, 0xcc, 0x38, 0xb5, 0x8b, 0xbe, 0x09, 0x3e, 0x91, 0x24,
	0x43, 0xcd, 0x86, 0x5c, 0xc6, 0x2c, 0xc9, 0x0a, 0x6b, 0xa5, 0x41, 0x6d, 0xc0, 0x23, 0x29, 0x08,
	0x95, 0xa8, 0x07, 0xa5, 0x14, 0x33, 0xb4, 0xb6, 0xd6, 0xd1, 0x9d, 0x33, 0x7b, 0x5b, 0xc6, 0x2e,
	0x98, 0xf6, 0x78, 0x19, 0x33, 0xac, 0xc8, 0xe8, 0x09, 0x0e, 0x93, 0x39, 0x11, 0xb2, 0xc0, 0x5c,
	0xdf, 0x17, 0x2c, 0x49, 0x8c, 0xbf, 0x6d, 0xad, 0xd3, 0x70, 0x8e, 0xd7, 0x22, 0x94, 0xcf, 0xe7,
	0x3c, 0xb2, 0x73, 0x18, 0xef, 0x1c, 0x42, 0x97, 0xa0, 0x8b, 0x90, 0x06, 0x44, 0xf8, 0x21, 0x89,
	0x86, 0x24, 0x09, 0x8c, 0x7f, 0x6d, 0xad, 0xd3, 0xc4, 0x5b, 0x5d, 0xeb, 0x1c, 0x4a, 0x69, 0x04,
	0xa4, 0x03, 0x60, 0x6f, 0x30, 0x74, 0xf1, 0x83, 0xe7, 0xbe, 0xb4, 0xfe, 0xa0, 0x3a, 0x94, 0x47,
	0xcf, 0x2e, 0x1e, 0xb7, 0x34, 0x8b, 0x43, 0xf3, 0x55, 0xf8, 0x4c, 0x60, 0xf6, 0xb1, 0x60, 0x89,
	0x44, 0x57, 0x50, 0x4a, 0x8f, 0xa0, 0x96, 0x6b, 0x38, 0x07, 0xeb, 0x5c, 0xea, 0x32, 0x7d, 0x22,
	0x69, 0x80, 0x15, 0x01, 0xdd, 0x41, 0x8d, 0xe6, 0xb1, 0xf2, 0x25, 0xcc, 0xfd, 0x97, 0xc0, 0x6b,
	0xae, 0x75, 0x01, 0x55, 0x65, 0xe8, 0xf9, 0xc8, 0x80, 0x2a, 0xcf, 0x5e, 0x95, 0x5d, 0x1d, 0x17,
	0xa5, 0xa5, 0x43, 0x73, 0xc4, 0xa4, 0x9c, 0x31, 0xcc, 0x92, 0xc5, 0x4c, 0x3a, 0x2b, 0x0d, 0xfe,
	0x3f, 0xe6, 0xa2, 0x23, 0x26, 0x3e, 0x43, 0xca, 0x50, 0x1f, 0xca, 0x4a, 0x08, 0x9d, 0xfe, 0xf4,
	0xdd, 0x5c, 0xc9, 0x3c, 0xd9, 0x83, 0x7b, 0x3e, 0x1a, 0x40, 0x25, 0xf3, 0x41, 0xfb, 0x49, 0xe6,
	0x0e, 0xfd, 0xcd, 0x70, 0xe8, 0x1e, 0x2a, 0x98, 0xbd, 0x33, 0x2a, 0x7f, 0x13, 0x39, 0xda, 0xfe,
	0xc2, 0xd9, 0x70, 0xff, 0xe6, 0xed, 0x7a, 0x1a, 0xca, 0x60, 0x31, 0x49, 0xfb, 0xdd, 0x9c, 0x53,
	0x3c, 0x6f, 0xa7, 0xbc, 0x4b, 0xe2, 0xb0, 0x5b, 0xa8, 0x4d, 0x2a, 0xea, 0x77, 0xec, 0x7d, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xca, 0xe8, 0x7c, 0xdc, 0xcb, 0x02, 0x00, 0x00,
}
