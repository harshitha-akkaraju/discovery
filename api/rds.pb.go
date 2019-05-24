// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rds.proto

package api

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ListRepositoriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRepositoriesRequest) Reset()         { *m = ListRepositoriesRequest{} }
func (m *ListRepositoriesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRepositoriesRequest) ProtoMessage()    {}
func (*ListRepositoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rds_0f41e9a66425029b, []int{0}
}
func (m *ListRepositoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRepositoriesRequest.Unmarshal(m, b)
}
func (m *ListRepositoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRepositoriesRequest.Marshal(b, m, deterministic)
}
func (dst *ListRepositoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRepositoriesRequest.Merge(dst, src)
}
func (m *ListRepositoriesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRepositoriesRequest.Size(m)
}
func (m *ListRepositoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRepositoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRepositoriesRequest proto.InternalMessageInfo

type ListRepositoriesResponse struct {
	Repositories         []string `protobuf:"bytes,1,rep,name=repositories,proto3" json:"repositories,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRepositoriesResponse) Reset()         { *m = ListRepositoriesResponse{} }
func (m *ListRepositoriesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRepositoriesResponse) ProtoMessage()    {}
func (*ListRepositoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rds_0f41e9a66425029b, []int{1}
}
func (m *ListRepositoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRepositoriesResponse.Unmarshal(m, b)
}
func (m *ListRepositoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRepositoriesResponse.Marshal(b, m, deterministic)
}
func (dst *ListRepositoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRepositoriesResponse.Merge(dst, src)
}
func (m *ListRepositoriesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRepositoriesResponse.Size(m)
}
func (m *ListRepositoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRepositoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRepositoriesResponse proto.InternalMessageInfo

func (m *ListRepositoriesResponse) GetRepositories() []string {
	if m != nil {
		return m.Repositories
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRepositoriesRequest)(nil), "cloud.deps.rds.api.ListRepositoriesRequest")
	proto.RegisterType((*ListRepositoriesResponse)(nil), "cloud.deps.rds.api.ListRepositoriesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepositoryDiscoveryServiceClient is the client API for RepositoryDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryDiscoveryServiceClient interface {
	List(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error)
}

type repositoryDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryDiscoveryServiceClient(cc *grpc.ClientConn) RepositoryDiscoveryServiceClient {
	return &repositoryDiscoveryServiceClient{cc}
}

func (c *repositoryDiscoveryServiceClient) List(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error) {
	out := new(ListRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/cloud.deps.rds.api.RepositoryDiscoveryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryDiscoveryServiceServer is the server API for RepositoryDiscoveryService service.
type RepositoryDiscoveryServiceServer interface {
	List(context.Context, *ListRepositoriesRequest) (*ListRepositoriesResponse, error)
}

func RegisterRepositoryDiscoveryServiceServer(s *grpc.Server, srv RepositoryDiscoveryServiceServer) {
	s.RegisterService(&_RepositoryDiscoveryService_serviceDesc, srv)
}

func _RepositoryDiscoveryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryDiscoveryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.deps.rds.api.RepositoryDiscoveryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryDiscoveryServiceServer).List(ctx, req.(*ListRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.deps.rds.api.RepositoryDiscoveryService",
	HandlerType: (*RepositoryDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RepositoryDiscoveryService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rds.proto",
}

func init() { proto.RegisterFile("rds.proto", fileDescriptor_rds_0f41e9a66425029b) }

var fileDescriptor_rds_0f41e9a66425029b = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x4a, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4b, 0x49, 0x2d,
	0x28, 0xd6, 0x03, 0x89, 0x26, 0x16, 0x64, 0x2a, 0x49, 0x72, 0x89, 0xfb, 0x64, 0x16, 0x97, 0x04,
	0xa5, 0x16, 0xe4, 0x17, 0x67, 0x96, 0xe4, 0x17, 0x65, 0xa6, 0x16, 0x07, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x28, 0xd9, 0x71, 0x49, 0x60, 0x4a, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x29,
	0x71, 0xf1, 0x14, 0x21, 0x89, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0xa1, 0x88, 0x19, 0xd5,
	0x73, 0x49, 0xc1, 0xf5, 0x56, 0xba, 0x64, 0x16, 0x27, 0xe7, 0x97, 0xa5, 0x16, 0x55, 0x06, 0xa7,
	0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x25, 0x72, 0xb1, 0x80, 0x4c, 0x17, 0xd2, 0xd6, 0xc3, 0x74,
	0x95, 0x1e, 0x0e, 0x27, 0x49, 0xe9, 0x10, 0xa7, 0x18, 0xe2, 0x48, 0x27, 0x55, 0x2e, 0x2c, 0x3e,
	0x76, 0xe2, 0x08, 0x72, 0x09, 0x0e, 0x00, 0x85, 0x47, 0x00, 0x63, 0x14, 0x73, 0x62, 0x41, 0x66,
	0x12, 0x1b, 0x38, 0x74, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x22, 0xd2, 0x0a, 0x60, 0x2a,
	0x01, 0x00, 0x00,
}
