// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: freon_service.proto

package freonApi

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_freon_service_proto protoreflect.FileDescriptor

var file_freon_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x1a, 0x12, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe8, 0x02, 0x0a, 0x0c,
	0x46, 0x72, 0x65, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x66, 0x72,
	0x65, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x66, 0x72, 0x65,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x66, 0x72,
	0x65, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12,
	0x65, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x66,
	0x72, 0x65, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x23, 0x2e, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x62, 0x3b, 0x66, 0x72, 0x65,
	0x6f, 0x6e, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_freon_service_proto_goTypes = []interface{}{
	(*GetListLocalizationsReq)(nil),      // 0: freon.GetListLocalizationsReq
	(*GetListTranslationsReq)(nil),       // 1: freon.GetListTranslationsReq
	(*GetTranslationReq)(nil),            // 2: freon.GetTranslationReq
	(*GetLatestTranslationFilesReq)(nil), // 3: freon.GetLatestTranslationFilesReq
	(*GetListLocalizationsRes)(nil),      // 4: freon.GetListLocalizationsRes
	(*GetListTranslationsRes)(nil),       // 5: freon.GetListTranslationsRes
	(*GetTranslationRes)(nil),            // 6: freon.GetTranslationRes
	(*GetLatestTranslationFilesRes)(nil), // 7: freon.GetLatestTranslationFilesRes
}
var file_freon_service_proto_depIdxs = []int32{
	0, // 0: freon.FreonService.GetListLocalizations:input_type -> freon.GetListLocalizationsReq
	1, // 1: freon.FreonService.GetListTranslations:input_type -> freon.GetListTranslationsReq
	2, // 2: freon.FreonService.GetTranslation:input_type -> freon.GetTranslationReq
	3, // 3: freon.FreonService.GetLatestTranslationFiles:input_type -> freon.GetLatestTranslationFilesReq
	4, // 4: freon.FreonService.GetListLocalizations:output_type -> freon.GetListLocalizationsRes
	5, // 5: freon.FreonService.GetListTranslations:output_type -> freon.GetListTranslationsRes
	6, // 6: freon.FreonService.GetTranslation:output_type -> freon.GetTranslationRes
	7, // 7: freon.FreonService.GetLatestTranslationFiles:output_type -> freon.GetLatestTranslationFilesRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_freon_service_proto_init() }
func file_freon_service_proto_init() {
	if File_freon_service_proto != nil {
		return
	}
	file_localization_proto_init()
	file_translation_proto_init()
	file_translation_file_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_freon_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_freon_service_proto_goTypes,
		DependencyIndexes: file_freon_service_proto_depIdxs,
	}.Build()
	File_freon_service_proto = out.File
	file_freon_service_proto_rawDesc = nil
	file_freon_service_proto_goTypes = nil
	file_freon_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FreonServiceClient is the client API for FreonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
type FreonServiceServer interface {
	GetListLocalizations(context.Context, *GetListLocalizationsReq) (*GetListLocalizationsRes, error)
	GetListTranslations(context.Context, *GetListTranslationsReq) (*GetListTranslationsRes, error)
	GetTranslation(context.Context, *GetTranslationReq) (*GetTranslationRes, error)
	GetLatestTranslationFiles(context.Context, *GetLatestTranslationFilesReq) (*GetLatestTranslationFilesRes, error)
}

// UnimplementedFreonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFreonServiceServer struct {
}

func (*UnimplementedFreonServiceServer) GetListLocalizations(context.Context, *GetListLocalizationsReq) (*GetListLocalizationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListLocalizations not implemented")
}
func (*UnimplementedFreonServiceServer) GetListTranslations(context.Context, *GetListTranslationsReq) (*GetListTranslationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTranslations not implemented")
}
func (*UnimplementedFreonServiceServer) GetTranslation(context.Context, *GetTranslationReq) (*GetTranslationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranslation not implemented")
}
func (*UnimplementedFreonServiceServer) GetLatestTranslationFiles(context.Context, *GetLatestTranslationFilesReq) (*GetLatestTranslationFilesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestTranslationFiles not implemented")
}

func RegisterFreonServiceServer(s *grpc.Server, srv FreonServiceServer) {
	s.RegisterService(&_FreonService_serviceDesc, srv)
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

var _FreonService_serviceDesc = grpc.ServiceDesc{
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
