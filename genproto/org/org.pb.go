// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.26.1
// source: proto/org/org.proto

package org

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

var File_proto_org_org_proto protoreflect.FileDescriptor

var file_proto_org_org_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x72, 0x67, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8b, 0x01, 0x0a,
	0x0a, 0x4f, 0x72, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x49,
	0x6e, 0x69, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x13, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x4f,
	0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b,
	0x4f, 0x72, 0x67, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x14, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_proto_org_org_proto_goTypes = []interface{}{
	(*InitOrgRequest)(nil),   // 0: org.InitOrgRequest
	(*EmployeeRequest)(nil),  // 1: org.EmployeeRequest
	(*InitOrgResponse)(nil),  // 2: org.InitOrgResponse
	(*EmployeeResponse)(nil), // 3: org.EmployeeResponse
}
var file_proto_org_org_proto_depIdxs = []int32{
	0, // 0: org.OrgService.InitOrganization:input_type -> org.InitOrgRequest
	1, // 1: org.OrgService.OrgEmployee:input_type -> org.EmployeeRequest
	2, // 2: org.OrgService.InitOrganization:output_type -> org.InitOrgResponse
	3, // 3: org.OrgService.OrgEmployee:output_type -> org.EmployeeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_org_org_proto_init() }
func file_proto_org_org_proto_init() {
	if File_proto_org_org_proto != nil {
		return
	}
	file_proto_org_type_employee_proto_init()
	file_proto_org_type_organization_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_org_org_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_org_org_proto_goTypes,
		DependencyIndexes: file_proto_org_org_proto_depIdxs,
	}.Build()
	File_proto_org_org_proto = out.File
	file_proto_org_org_proto_rawDesc = nil
	file_proto_org_org_proto_goTypes = nil
	file_proto_org_org_proto_depIdxs = nil
}
