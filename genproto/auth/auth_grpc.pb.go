// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: proto/auth/auth.proto

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

const (
	AuthService_Verify_FullMethodName           = "/auth.AuthService/Verify"
	AuthService_UpdateInfo_FullMethodName       = "/auth.AuthService/UpdateInfo"
	AuthService_GetNewEmailToken_FullMethodName = "/auth.AuthService/GetNewEmailToken"
	AuthService_VerifyEmailToken_FullMethodName = "/auth.AuthService/VerifyEmailToken"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	UpdateInfo(ctx context.Context, in *InfoUpdateRequest, opts ...grpc.CallOption) (*InfoUpdateResponse, error)
	GetNewEmailToken(ctx context.Context, in *EmailTokenRequest, opts ...grpc.CallOption) (*EmailTokenResponse, error)
	VerifyEmailToken(ctx context.Context, in *VerifyEmailTokenRequest, opts ...grpc.CallOption) (*VerifyEmailTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, AuthService_Verify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateInfo(ctx context.Context, in *InfoUpdateRequest, opts ...grpc.CallOption) (*InfoUpdateResponse, error) {
	out := new(InfoUpdateResponse)
	err := c.cc.Invoke(ctx, AuthService_UpdateInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetNewEmailToken(ctx context.Context, in *EmailTokenRequest, opts ...grpc.CallOption) (*EmailTokenResponse, error) {
	out := new(EmailTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_GetNewEmailToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyEmailToken(ctx context.Context, in *VerifyEmailTokenRequest, opts ...grpc.CallOption) (*VerifyEmailTokenResponse, error) {
	out := new(VerifyEmailTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_VerifyEmailToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	UpdateInfo(context.Context, *InfoUpdateRequest) (*InfoUpdateResponse, error)
	GetNewEmailToken(context.Context, *EmailTokenRequest) (*EmailTokenResponse, error)
	VerifyEmailToken(context.Context, *VerifyEmailTokenRequest) (*VerifyEmailTokenResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAuthServiceServer) UpdateInfo(context.Context, *InfoUpdateRequest) (*InfoUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInfo not implemented")
}
func (UnimplementedAuthServiceServer) GetNewEmailToken(context.Context, *EmailTokenRequest) (*EmailTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewEmailToken not implemented")
}
func (UnimplementedAuthServiceServer) VerifyEmailToken(context.Context, *VerifyEmailTokenRequest) (*VerifyEmailTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmailToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateInfo(ctx, req.(*InfoUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetNewEmailToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetNewEmailToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetNewEmailToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetNewEmailToken(ctx, req.(*EmailTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyEmailToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyEmailToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyEmailToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyEmailToken(ctx, req.(*VerifyEmailTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _AuthService_Verify_Handler,
		},
		{
			MethodName: "UpdateInfo",
			Handler:    _AuthService_UpdateInfo_Handler,
		},
		{
			MethodName: "GetNewEmailToken",
			Handler:    _AuthService_GetNewEmailToken_Handler,
		},
		{
			MethodName: "VerifyEmailToken",
			Handler:    _AuthService_VerifyEmailToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth/auth.proto",
}
