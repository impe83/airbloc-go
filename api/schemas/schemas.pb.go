// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schemas.proto

package schemas // import "github.com/airbloc/airbloc-go/api/schemas"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type CreateSchemaRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Schema               string   `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSchemaRequest) Reset()         { *m = CreateSchemaRequest{} }
func (m *CreateSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSchemaRequest) ProtoMessage()    {}
func (*CreateSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_schemas_771b1f65513ec2af, []int{0}
}
func (m *CreateSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSchemaRequest.Unmarshal(m, b)
}
func (m *CreateSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSchemaRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSchemaRequest.Merge(dst, src)
}
func (m *CreateSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSchemaRequest.Size(m)
}
func (m *CreateSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSchemaRequest proto.InternalMessageInfo

func (m *CreateSchemaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSchemaRequest) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

type CreateSchemaResult struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	GasUsed              uint64   `protobuf:"varint,3,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSchemaResult) Reset()         { *m = CreateSchemaResult{} }
func (m *CreateSchemaResult) String() string { return proto.CompactTextString(m) }
func (*CreateSchemaResult) ProtoMessage()    {}
func (*CreateSchemaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_schemas_771b1f65513ec2af, []int{1}
}
func (m *CreateSchemaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSchemaResult.Unmarshal(m, b)
}
func (m *CreateSchemaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSchemaResult.Marshal(b, m, deterministic)
}
func (dst *CreateSchemaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSchemaResult.Merge(dst, src)
}
func (m *CreateSchemaResult) XXX_Size() int {
	return xxx_messageInfo_CreateSchemaResult.Size(m)
}
func (m *CreateSchemaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSchemaResult.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSchemaResult proto.InternalMessageInfo

func (m *CreateSchemaResult) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func (m *CreateSchemaResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateSchemaResult) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateSchemaRequest)(nil), "airbloc.schemas.CreateSchemaRequest")
	proto.RegisterType((*CreateSchemaResult)(nil), "airbloc.schemas.CreateSchemaResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchemaServiceClient interface {
	Create(ctx context.Context, in *CreateSchemaRequest, opts ...grpc.CallOption) (*CreateSchemaResult, error)
}

type schemaServiceClient struct {
	cc *grpc.ClientConn
}

func NewSchemaServiceClient(cc *grpc.ClientConn) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) Create(ctx context.Context, in *CreateSchemaRequest, opts ...grpc.CallOption) (*CreateSchemaResult, error) {
	out := new(CreateSchemaResult)
	err := c.cc.Invoke(ctx, "/airbloc.schemas.SchemaService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
type SchemaServiceServer interface {
	Create(context.Context, *CreateSchemaRequest) (*CreateSchemaResult, error)
}

func RegisterSchemaServiceServer(s *grpc.Server, srv SchemaServiceServer) {
	s.RegisterService(&_SchemaService_serviceDesc, srv)
}

func _SchemaService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.schemas.SchemaService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).Create(ctx, req.(*CreateSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchemaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.schemas.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SchemaService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schemas.proto",
}

func init() { proto.RegisterFile("schemas.proto", fileDescriptor_schemas_771b1f65513ec2af) }

var fileDescriptor_schemas_771b1f65513ec2af = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0xcc, 0x2c, 0x4a, 0xca,
	0xc9, 0x4f, 0xd6, 0x83, 0x0a, 0x2b, 0x39, 0x72, 0x09, 0x3b, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x06,
	0x83, 0x05, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x31, 0x2e, 0x36, 0x88, 0x2e,
	0x09, 0x26, 0xb0, 0x28, 0x94, 0xa7, 0x14, 0xc6, 0x25, 0x84, 0x6a, 0x44, 0x71, 0x69, 0x4e, 0x09,
	0x48, 0x75, 0x6a, 0x45, 0x66, 0x71, 0x49, 0x31, 0xd8, 0x0c, 0x8e, 0x20, 0x28, 0x4f, 0x88, 0x8f,
	0x8b, 0x29, 0x33, 0x05, 0x6a, 0x02, 0x53, 0x66, 0x8a, 0x90, 0x04, 0x17, 0x7b, 0x7a, 0x62, 0x71,
	0x68, 0x71, 0x6a, 0x8a, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x8c, 0x6b, 0x94, 0xc2, 0xc5,
	0x0b, 0x31, 0x31, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x98, 0x8b, 0x0d, 0x62, 0x91,
	0x90, 0x8a, 0x1e, 0x9a, 0x3f, 0xf4, 0xb0, 0x78, 0x42, 0x4a, 0x99, 0x80, 0x2a, 0x90, 0x3b, 0x9d,
	0xb4, 0xa3, 0x34, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xa1, 0x1a,
	0x60, 0xb4, 0x6e, 0x7a, 0xbe, 0x7e, 0x62, 0x41, 0xa6, 0x3e, 0x54, 0x7f, 0x12, 0x1b, 0x38, 0x14,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xbc, 0xa9, 0x97, 0x56, 0x01, 0x00, 0x00,
}
