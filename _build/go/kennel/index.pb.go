// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index.proto

package kennel

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	ptypes "github.com/lzambarda/protorepo/_build/go/ptypes"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("index.proto", fileDescriptor_f750e0f7889345b5) }

var fileDescriptor_f750e0f7889345b5 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcc, 0x4b, 0x49,
	0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x4e, 0xcd, 0xcb, 0x4b, 0xcd, 0x91,
	0xe2, 0x81, 0xd0, 0x10, 0x51, 0x29, 0xa1, 0x82, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x7d, 0x24, 0x95,
	0x46, 0x76, 0x5c, 0x6c, 0xde, 0x60, 0x35, 0x42, 0x26, 0x5c, 0x42, 0xee, 0xa9, 0x25, 0x4e, 0xf9,
	0x45, 0xd9, 0xa9, 0x45, 0xc5, 0x6e, 0xf9, 0x45, 0xfe, 0xe5, 0x79, 0xa9, 0x45, 0x42, 0x3c, 0x7a,
	0x10, 0x4d, 0x7a, 0xa1, 0xc5, 0xa9, 0x45, 0x52, 0x7c, 0x7a, 0x50, 0x03, 0x21, 0xca, 0x0c, 0x18,
	0x93, 0xd8, 0xc0, 0xc6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x05, 0xbd, 0x21, 0x7f,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KennelClient is the client API for Kennel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KennelClient interface {
	GetBorkersForOwner(ctx context.Context, in *ptypes.User, opts ...grpc.CallOption) (Kennel_GetBorkersForOwnerClient, error)
}

type kennelClient struct {
	cc *grpc.ClientConn
}

func NewKennelClient(cc *grpc.ClientConn) KennelClient {
	return &kennelClient{cc}
}

func (c *kennelClient) GetBorkersForOwner(ctx context.Context, in *ptypes.User, opts ...grpc.CallOption) (Kennel_GetBorkersForOwnerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Kennel_serviceDesc.Streams[0], "/kennel.Kennel/GetBorkersForOwner", opts...)
	if err != nil {
		return nil, err
	}
	x := &kennelGetBorkersForOwnerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Kennel_GetBorkersForOwnerClient interface {
	Recv() (*Borker, error)
	grpc.ClientStream
}

type kennelGetBorkersForOwnerClient struct {
	grpc.ClientStream
}

func (x *kennelGetBorkersForOwnerClient) Recv() (*Borker, error) {
	m := new(Borker)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KennelServer is the server API for Kennel service.
type KennelServer interface {
	GetBorkersForOwner(*ptypes.User, Kennel_GetBorkersForOwnerServer) error
}

// UnimplementedKennelServer can be embedded to have forward compatible implementations.
type UnimplementedKennelServer struct {
}

func (*UnimplementedKennelServer) GetBorkersForOwner(req *ptypes.User, srv Kennel_GetBorkersForOwnerServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBorkersForOwner not implemented")
}

func RegisterKennelServer(s *grpc.Server, srv KennelServer) {
	s.RegisterService(&_Kennel_serviceDesc, srv)
}

func _Kennel_GetBorkersForOwner_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ptypes.User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KennelServer).GetBorkersForOwner(m, &kennelGetBorkersForOwnerServer{stream})
}

type Kennel_GetBorkersForOwnerServer interface {
	Send(*Borker) error
	grpc.ServerStream
}

type kennelGetBorkersForOwnerServer struct {
	grpc.ServerStream
}

func (x *kennelGetBorkersForOwnerServer) Send(m *Borker) error {
	return x.ServerStream.SendMsg(m)
}

var _Kennel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kennel.Kennel",
	HandlerType: (*KennelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBorkersForOwner",
			Handler:       _Kennel_GetBorkersForOwner_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "index.proto",
}
