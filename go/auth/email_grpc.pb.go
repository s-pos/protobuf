// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

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

// UserAuthServiceClient is the client API for UserAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAuthServiceClient interface {
	// gRPC for send email verification
	SendEmailVerification(ctx context.Context, in *Verification, opts ...grpc.CallOption) (*VerificationReply, error)
	// gRPC for send email reset password
	SendEmailResetPassword(ctx context.Context, in *ResetPassword, opts ...grpc.CallOption) (*ResetPasswordReply, error)
}

type userAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAuthServiceClient(cc grpc.ClientConnInterface) UserAuthServiceClient {
	return &userAuthServiceClient{cc}
}

func (c *userAuthServiceClient) SendEmailVerification(ctx context.Context, in *Verification, opts ...grpc.CallOption) (*VerificationReply, error) {
	out := new(VerificationReply)
	err := c.cc.Invoke(ctx, "/auth.UserAuthService/SendEmailVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthServiceClient) SendEmailResetPassword(ctx context.Context, in *ResetPassword, opts ...grpc.CallOption) (*ResetPasswordReply, error) {
	out := new(ResetPasswordReply)
	err := c.cc.Invoke(ctx, "/auth.UserAuthService/SendEmailResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAuthServiceServer is the server API for UserAuthService service.
// All implementations must embed UnimplementedUserAuthServiceServer
// for forward compatibility
type UserAuthServiceServer interface {
	// gRPC for send email verification
	SendEmailVerification(context.Context, *Verification) (*VerificationReply, error)
	// gRPC for send email reset password
	SendEmailResetPassword(context.Context, *ResetPassword) (*ResetPasswordReply, error)
	mustEmbedUnimplementedUserAuthServiceServer()
}

// UnimplementedUserAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserAuthServiceServer struct {
}

func (UnimplementedUserAuthServiceServer) SendEmailVerification(context.Context, *Verification) (*VerificationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailVerification not implemented")
}
func (UnimplementedUserAuthServiceServer) SendEmailResetPassword(context.Context, *ResetPassword) (*ResetPasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailResetPassword not implemented")
}
func (UnimplementedUserAuthServiceServer) mustEmbedUnimplementedUserAuthServiceServer() {}

// UnsafeUserAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAuthServiceServer will
// result in compilation errors.
type UnsafeUserAuthServiceServer interface {
	mustEmbedUnimplementedUserAuthServiceServer()
}

func RegisterUserAuthServiceServer(s grpc.ServiceRegistrar, srv UserAuthServiceServer) {
	s.RegisterService(&UserAuthService_ServiceDesc, srv)
}

func _UserAuthService_SendEmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServiceServer).SendEmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserAuthService/SendEmailVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServiceServer).SendEmailVerification(ctx, req.(*Verification))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuthService_SendEmailResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServiceServer).SendEmailResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserAuthService/SendEmailResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServiceServer).SendEmailResetPassword(ctx, req.(*ResetPassword))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAuthService_ServiceDesc is the grpc.ServiceDesc for UserAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.UserAuthService",
	HandlerType: (*UserAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmailVerification",
			Handler:    _UserAuthService_SendEmailVerification_Handler,
		},
		{
			MethodName: "SendEmailResetPassword",
			Handler:    _UserAuthService_SendEmailResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/auth/email.proto",
}
