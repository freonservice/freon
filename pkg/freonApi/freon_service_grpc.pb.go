// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: freon_service.proto

package freonApi

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

// FreonServiceClient is the client API for FreonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreonServiceClient interface {
	GetListLocalizations(ctx context.Context, in *GetListLocalizationsReq, opts ...grpc.CallOption) (*GetListLocalizationsRes, error)
	GetListTranslations(ctx context.Context, in *GetListTranslationsReq, opts ...grpc.CallOption) (*GetListTranslationsRes, error)
	GetTranslation(ctx context.Context, in *GetTranslationReq, opts ...grpc.CallOption) (*GetTranslationRes, error)
	GetLatestTranslationFiles(ctx context.Context, in *GetLatestTranslationFilesReq, opts ...grpc.CallOption) (*GetLatestTranslationFilesRes, error)
}

type freonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreonServiceClient(cc grpc.ClientConnInterface) FreonServiceClient {
	return &freonServiceClient{cc}
}

func (c *freonServiceClient) GetListLocalizations(ctx context.Context, in *GetListLocalizationsReq, opts ...grpc.CallOption) (*GetListLocalizationsRes, error) {
	out := new(GetListLocalizationsRes)
	err := c.cc.Invoke(ctx, "/freon.FreonService/GetListLocalizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freonServiceClient) GetListTranslations(ctx context.Context, in *GetListTranslationsReq, opts ...grpc.CallOption) (*GetListTranslationsRes, error) {
	out := new(GetListTranslationsRes)
	err := c.cc.Invoke(ctx, "/freon.FreonService/GetListTranslations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freonServiceClient) GetTranslation(ctx context.Context, in *GetTranslationReq, opts ...grpc.CallOption) (*GetTranslationRes, error) {
	out := new(GetTranslationRes)
	err := c.cc.Invoke(ctx, "/freon.FreonService/GetTranslation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freonServiceClient) GetLatestTranslationFiles(ctx context.Context, in *GetLatestTranslationFilesReq, opts ...grpc.CallOption) (*GetLatestTranslationFilesRes, error) {
	out := new(GetLatestTranslationFilesRes)
	err := c.cc.Invoke(ctx, "/freon.FreonService/GetLatestTranslationFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreonServiceServer is the server API for FreonService service.
// All implementations should embed UnimplementedFreonServiceServer
// for forward compatibility
type FreonServiceServer interface {
	GetListLocalizations(context.Context, *GetListLocalizationsReq) (*GetListLocalizationsRes, error)
	GetListTranslations(context.Context, *GetListTranslationsReq) (*GetListTranslationsRes, error)
	GetTranslation(context.Context, *GetTranslationReq) (*GetTranslationRes, error)
	GetLatestTranslationFiles(context.Context, *GetLatestTranslationFilesReq) (*GetLatestTranslationFilesRes, error)
}

// UnimplementedFreonServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFreonServiceServer struct {
}

func (UnimplementedFreonServiceServer) GetListLocalizations(context.Context, *GetListLocalizationsReq) (*GetListLocalizationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListLocalizations not implemented")
}
func (UnimplementedFreonServiceServer) GetListTranslations(context.Context, *GetListTranslationsReq) (*GetListTranslationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTranslations not implemented")
}
func (UnimplementedFreonServiceServer) GetTranslation(context.Context, *GetTranslationReq) (*GetTranslationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranslation not implemented")
}
func (UnimplementedFreonServiceServer) GetLatestTranslationFiles(context.Context, *GetLatestTranslationFilesReq) (*GetLatestTranslationFilesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestTranslationFiles not implemented")
}

// UnsafeFreonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreonServiceServer will
// result in compilation errors.
type UnsafeFreonServiceServer interface {
	mustEmbedUnimplementedFreonServiceServer()
}

func RegisterFreonServiceServer(s grpc.ServiceRegistrar, srv FreonServiceServer) {
	s.RegisterService(&FreonService_ServiceDesc, srv)
}

func _FreonService_GetListLocalizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListLocalizationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreonServiceServer).GetListLocalizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freon.FreonService/GetListLocalizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreonServiceServer).GetListLocalizations(ctx, req.(*GetListLocalizationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreonService_GetListTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTranslationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreonServiceServer).GetListTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freon.FreonService/GetListTranslations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreonServiceServer).GetListTranslations(ctx, req.(*GetListTranslationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreonService_GetTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTranslationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreonServiceServer).GetTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freon.FreonService/GetTranslation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreonServiceServer).GetTranslation(ctx, req.(*GetTranslationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreonService_GetLatestTranslationFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestTranslationFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreonServiceServer).GetLatestTranslationFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freon.FreonService/GetLatestTranslationFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreonServiceServer).GetLatestTranslationFiles(ctx, req.(*GetLatestTranslationFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FreonService_ServiceDesc is the grpc.ServiceDesc for FreonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "freon.FreonService",
	HandlerType: (*FreonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListLocalizations",
			Handler:    _FreonService_GetListLocalizations_Handler,
		},
		{
			MethodName: "GetListTranslations",
			Handler:    _FreonService_GetListTranslations_Handler,
		},
		{
			MethodName: "GetTranslation",
			Handler:    _FreonService_GetTranslation_Handler,
		},
		{
			MethodName: "GetLatestTranslationFiles",
			Handler:    _FreonService_GetLatestTranslationFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "freon_service.proto",
}
