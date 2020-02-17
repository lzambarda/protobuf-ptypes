// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index.proto

package notifier

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcc, 0x4b, 0x49,
	0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c,
	0x2d, 0x92, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x27, 0x95, 0xa6, 0xe9,
	0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x94, 0x49, 0xf1, 0xc1, 0x94, 0x41, 0xf9, 0x42, 0x05, 0x25,
	0x95, 0x05, 0xa9, 0xc5, 0xfa, 0x48, 0x46, 0x19, 0x4d, 0x65, 0xe4, 0xe2, 0xf0, 0x83, 0x2a, 0x13,
	0xf2, 0xe5, 0x12, 0x08, 0x4e, 0xcd, 0x4b, 0x81, 0xf0, 0x93, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0x84,
	0x14, 0xf5, 0xe0, 0xa6, 0xa0, 0xcb, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0x89, 0xe9,
	0x41, 0x5c, 0xa1, 0x07, 0x73, 0x85, 0x9e, 0x2b, 0xc8, 0x15, 0x42, 0x56, 0x5c, 0x02, 0xee, 0xa9,
	0x25, 0xc8, 0x3a, 0x8a, 0x85, 0x78, 0xf4, 0x20, 0x8e, 0xd0, 0x0b, 0x2d, 0x4e, 0x2d, 0x92, 0x12,
	0x43, 0x18, 0x8e, 0xac, 0xcc, 0x80, 0x31, 0x89, 0x0d, 0x6c, 0x96, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x39, 0xf2, 0x02, 0x92, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotifierClient is the client API for Notifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotifierClient interface {
	// SendNotification creates a notification for a user, always return empty
	// response no matter the result.
	SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// GetNotifications returns a stream of Notification for the requesting
	// user.
	GetNotifications(ctx context.Context, in *ptypes.User, opts ...grpc.CallOption) (Notifier_GetNotificationsClient, error)
}

type notifierClient struct {
	cc *grpc.ClientConn
}

func NewNotifierClient(cc *grpc.ClientConn) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notifier.Notifier/SendNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) GetNotifications(ctx context.Context, in *ptypes.User, opts ...grpc.CallOption) (Notifier_GetNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Notifier_serviceDesc.Streams[0], "/notifier.Notifier/GetNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &notifierGetNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Notifier_GetNotificationsClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type notifierGetNotificationsClient struct {
	grpc.ClientStream
}

func (x *notifierGetNotificationsClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotifierServer is the server API for Notifier service.
type NotifierServer interface {
	// SendNotification creates a notification for a user, always return empty
	// response no matter the result.
	SendNotification(context.Context, *SendNotificationRequest) (*empty.Empty, error)
	// GetNotifications returns a stream of Notification for the requesting
	// user.
	GetNotifications(*ptypes.User, Notifier_GetNotificationsServer) error
}

// UnimplementedNotifierServer can be embedded to have forward compatible implementations.
type UnimplementedNotifierServer struct {
}

func (*UnimplementedNotifierServer) SendNotification(ctx context.Context, req *SendNotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (*UnimplementedNotifierServer) GetNotifications(req *ptypes.User, srv Notifier_GetNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}

func RegisterNotifierServer(s *grpc.Server, srv NotifierServer) {
	s.RegisterService(&_Notifier_serviceDesc, srv)
}

func _Notifier_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifier.Notifier/SendNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).SendNotification(ctx, req.(*SendNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_GetNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ptypes.User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotifierServer).GetNotifications(m, &notifierGetNotificationsServer{stream})
}

type Notifier_GetNotificationsServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type notifierGetNotificationsServer struct {
	grpc.ServerStream
}

func (x *notifierGetNotificationsServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

var _Notifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notifier.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNotification",
			Handler:    _Notifier_SendNotification_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNotifications",
			Handler:       _Notifier_GetNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "index.proto",
}