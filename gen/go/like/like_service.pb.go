// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: like_service.proto

package like

import (
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

var File_like_service_proto protoreflect.FileDescriptor

var file_like_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x1a, 0x12, 0x6c, 0x69, 0x6b, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x83,
	0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33,
	0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x6c, 0x69, 0x6b,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x6b, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x11, 0x2e,
	0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69,
	0x6b, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_like_service_proto_goTypes = []any{
	(*ExistsRequest)(nil),  // 0: like.ExistsRequest
	(*AddRequest)(nil),     // 1: like.AddRequest
	(*DeleteRequest)(nil),  // 2: like.DeleteRequest
	(*CountRequest)(nil),   // 3: like.CountRequest
	(*FindRequest)(nil),    // 4: like.FindRequest
	(*ExistsResponse)(nil), // 5: like.ExistsResponse
	(*AddResponse)(nil),    // 6: like.AddResponse
	(*CountResponse)(nil),  // 7: like.CountResponse
	(*FindResponse)(nil),   // 8: like.FindResponse
}
var file_like_service_proto_depIdxs = []int32{
	0, // 0: like.LikeService.Exists:input_type -> like.ExistsRequest
	1, // 1: like.LikeService.Add:input_type -> like.AddRequest
	2, // 2: like.LikeService.Delete:input_type -> like.DeleteRequest
	3, // 3: like.LikeService.Count:input_type -> like.CountRequest
	4, // 4: like.LikeService.Find:input_type -> like.FindRequest
	5, // 5: like.LikeService.Exists:output_type -> like.ExistsResponse
	6, // 6: like.LikeService.Add:output_type -> like.AddResponse
	2, // 7: like.LikeService.Delete:output_type -> like.DeleteRequest
	7, // 8: like.LikeService.Count:output_type -> like.CountResponse
	8, // 9: like.LikeService.Find:output_type -> like.FindResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_like_service_proto_init() }
func file_like_service_proto_init() {
	if File_like_service_proto != nil {
		return
	}
	file_like_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_like_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_like_service_proto_goTypes,
		DependencyIndexes: file_like_service_proto_depIdxs,
	}.Build()
	File_like_service_proto = out.File
	file_like_service_proto_rawDesc = nil
	file_like_service_proto_goTypes = nil
	file_like_service_proto_depIdxs = nil
}
