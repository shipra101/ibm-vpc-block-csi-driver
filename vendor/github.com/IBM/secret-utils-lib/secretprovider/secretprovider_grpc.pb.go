/**
 * Copyright 2022 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package secretprovider

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion6

// SecretProviderClient is the client API for SecretProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretProviderClient interface {
	GetIAMToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*IAMToken, error)
	GetDefaultIAMToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*IAMToken, error)
}

type secretProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretProviderClient(cc grpc.ClientConnInterface) SecretProviderClient {
	return &secretProviderClient{cc}
}

func (c *secretProviderClient) GetIAMToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*IAMToken, error) {
	out := new(IAMToken)
	err := c.cc.Invoke(ctx, "/secretprovider.SecretProvider/GetIAMToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretProviderClient) GetDefaultIAMToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*IAMToken, error) {
	out := new(IAMToken)
	err := c.cc.Invoke(ctx, "/secretprovider.SecretProvider/GetDefaultIAMToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretProviderServer is the server API for SecretProvider service.
// All implementations must embed UnimplementedSecretProviderServer
// for forward compatibility
type SecretProviderServer interface {
	GetIAMToken(context.Context, *Request) (*IAMToken, error)
	GetDefaultIAMToken(context.Context, *Request) (*IAMToken, error)
	mustEmbedUnimplementedSecretProviderServer()
}

// UnimplementedSecretProviderServer must be embedded to have forward compatible implementations.
type UnimplementedSecretProviderServer struct {
}

func (UnimplementedSecretProviderServer) GetIAMToken(context.Context, *Request) (*IAMToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIAMToken not implemented")
}
func (UnimplementedSecretProviderServer) GetDefaultIAMToken(context.Context, *Request) (*IAMToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultIAMToken not implemented")
}
func (UnimplementedSecretProviderServer) mustEmbedUnimplementedSecretProviderServer() {}

// UnsafeSecretProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretProviderServer will
// result in compilation errors.
type UnsafeSecretProviderServer interface {
	mustEmbedUnimplementedSecretProviderServer()
}

func RegisterSecretProviderServer(s *grpc.Server, srv SecretProviderServer) {
	s.RegisterService(&SecretProvider_ServiceDesc, srv)
}

func _SecretProvider_GetIAMToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretProviderServer).GetIAMToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secretprovider.SecretProvider/GetIAMToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretProviderServer).GetIAMToken(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretProvider_GetDefaultIAMToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretProviderServer).GetDefaultIAMToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/secretprovider.SecretProvider/GetDefaultIAMToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretProviderServer).GetDefaultIAMToken(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretProvider_ServiceDesc is the grpc.ServiceDesc for SecretProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "secretprovider.SecretProvider",
	HandlerType: (*SecretProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIAMToken",
			Handler:    _SecretProvider_GetIAMToken_Handler,
		},
		{
			MethodName: "GetDefaultIAMToken",
			Handler:    _SecretProvider_GetDefaultIAMToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secretprovider.proto",
}