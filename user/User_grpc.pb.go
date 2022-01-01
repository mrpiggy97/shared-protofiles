// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	RegisterUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_RegisterUsersClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userEntity.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_RegisterUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/userEntity.UserService/RegisterUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceRegisterUsersClient{stream}
	return x, nil
}

type UserService_RegisterUsersClient interface {
	Send(*RegisterUserRequest) error
	Recv() (*RegisterUserResponse, error)
	grpc.ClientStream
}

type userServiceRegisterUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceRegisterUsersClient) Send(m *RegisterUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceRegisterUsersClient) Recv() (*RegisterUserResponse, error) {
	m := new(RegisterUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUser(context.Context, *UserRequest) (*User, error)
	RegisterUsers(UserService_RegisterUsersServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) RegisterUsers(UserService_RegisterUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterUsers not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userEntity.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).RegisterUsers(&userServiceRegisterUsersServer{stream})
}

type UserService_RegisterUsersServer interface {
	Send(*RegisterUserResponse) error
	Recv() (*RegisterUserRequest, error)
	grpc.ServerStream
}

type userServiceRegisterUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceRegisterUsersServer) Send(m *RegisterUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceRegisterUsersServer) Recv() (*RegisterUserRequest, error) {
	m := new(RegisterUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userEntity.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterUsers",
			Handler:       _UserService_RegisterUsers_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protofiles/User.proto",
}
