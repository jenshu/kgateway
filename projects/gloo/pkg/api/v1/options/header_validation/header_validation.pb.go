// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/header_validation/header_validation.proto

package header_validation

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HeaderValidationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How header methods will be validated. By default, Envoy will
	// validate HTTP methods for HTTP/1 connections against the default set of allowed methods.
	// The default allowed methods can be found here:
	// https://github.com/envoyproxy/envoy/blob/2970ddbd4ade787dd51dfbe605ae2e8c5d8ffcf7/source/common/http/http1/balsa_parser.cc#L54
	// or here, if Universal Header Validation is enabled:
	// https://github.com/envoyproxy/envoy/blob/0b9f67e7f71bcba3ff49575dc61676478cb68614/source/extensions/http/header_validators/envoy_default/header_validator.cc#L53
	// Invalid methods on HTTP/1 requests will be rejected with a HTTP 400
	// response.
	// For HTTP/2, Envoy does not validate header methods by default.
	//
	// Types that are assignable to HeaderMethodValidation:
	//
	//	*HeaderValidationSettings_DisableHttp1MethodValidation
	HeaderMethodValidation isHeaderValidationSettings_HeaderMethodValidation `protobuf_oneof:"header_method_validation"`
}

func (x *HeaderValidationSettings) Reset() {
	*x = HeaderValidationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValidationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValidationSettings) ProtoMessage() {}

func (x *HeaderValidationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValidationSettings.ProtoReflect.Descriptor instead.
func (*HeaderValidationSettings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDescGZIP(), []int{0}
}

func (m *HeaderValidationSettings) GetHeaderMethodValidation() isHeaderValidationSettings_HeaderMethodValidation {
	if m != nil {
		return m.HeaderMethodValidation
	}
	return nil
}

func (x *HeaderValidationSettings) GetDisableHttp1MethodValidation() *empty.Empty {
	if x, ok := x.GetHeaderMethodValidation().(*HeaderValidationSettings_DisableHttp1MethodValidation); ok {
		return x.DisableHttp1MethodValidation
	}
	return nil
}

type isHeaderValidationSettings_HeaderMethodValidation interface {
	isHeaderValidationSettings_HeaderMethodValidation()
}

type HeaderValidationSettings_DisableHttp1MethodValidation struct {
	// Disable method validation. Envoy will not perform any validation on
	// the method provided in the HTTP header.
	DisableHttp1MethodValidation *empty.Empty `protobuf:"bytes,1,opt,name=disable_http1_method_validation,json=disableHttp1MethodValidation,proto3,oneof"`
}

func (*HeaderValidationSettings_DisableHttp1MethodValidation) isHeaderValidationSettings_HeaderMethodValidation() {
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x26, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x18, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x5f, 0x0a, 0x1f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x31, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x74,
	0x74, 0x70, 0x31, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x1a, 0x0a, 0x18, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_goTypes = []interface{}{
	(*HeaderValidationSettings)(nil), // 0: header_validation.options.gloo.solo.io.HeaderValidationSettings
	(*empty.Empty)(nil),              // 1: google.protobuf.Empty
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_depIdxs = []int32{
	1, // 0: header_validation.options.gloo.solo.io.HeaderValidationSettings.disable_http1_method_validation:type_name -> google.protobuf.Empty
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValidationSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HeaderValidationSettings_DisableHttp1MethodValidation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_header_validation_header_validation_proto_depIdxs = nil
}
