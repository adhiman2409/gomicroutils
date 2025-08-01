// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: proto/org/type/organization.proto

package org

import (
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

type InitOrgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Eid               string `protobuf:"bytes,2,opt,name=eid,proto3" json:"eid,omitempty"`
	Domain            string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	OrgName           string `protobuf:"bytes,4,opt,name=orgName,proto3" json:"orgName,omitempty"`
	AdminEmailId      string `protobuf:"bytes,5,opt,name=adminEmailId,proto3" json:"adminEmailId,omitempty"`
	AdminPhoneNumber  string `protobuf:"bytes,6,opt,name=adminPhoneNumber,proto3" json:"adminPhoneNumber,omitempty"`
	AdminName         string `protobuf:"bytes,7,opt,name=adminName,proto3" json:"adminName,omitempty"`
	Department        string `protobuf:"bytes,8,opt,name=department,proto3" json:"department,omitempty"`
	Designation       string `protobuf:"bytes,9,opt,name=designation,proto3" json:"designation,omitempty"`
	EmailVerified     bool   `protobuf:"varint,10,opt,name=emailVerified,proto3" json:"emailVerified,omitempty"`
	Password          string `protobuf:"bytes,11,opt,name=password,proto3" json:"password,omitempty"`
	Role              string `protobuf:"bytes,12,opt,name=role,proto3" json:"role,omitempty"`
	FirstLoginPending bool   `protobuf:"varint,13,opt,name=firstLoginPending,proto3" json:"firstLoginPending,omitempty"`
	Status            string `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	UseGoogleOAuth    bool   `protobuf:"varint,15,opt,name=useGoogleOAuth,proto3" json:"useGoogleOAuth,omitempty"`
	Country           string `protobuf:"bytes,16,opt,name=country,proto3" json:"country,omitempty"`
	TimeZone          string `protobuf:"bytes,17,opt,name=timeZone,proto3" json:"timeZone,omitempty"`
}

func (x *InitOrgRequest) Reset() {
	*x = InitOrgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_org_type_organization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitOrgRequest) ProtoMessage() {}

func (x *InitOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_org_type_organization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitOrgRequest.ProtoReflect.Descriptor instead.
func (*InitOrgRequest) Descriptor() ([]byte, []int) {
	return file_proto_org_type_organization_proto_rawDescGZIP(), []int{0}
}

func (x *InitOrgRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InitOrgRequest) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

func (x *InitOrgRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *InitOrgRequest) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *InitOrgRequest) GetAdminEmailId() string {
	if x != nil {
		return x.AdminEmailId
	}
	return ""
}

func (x *InitOrgRequest) GetAdminPhoneNumber() string {
	if x != nil {
		return x.AdminPhoneNumber
	}
	return ""
}

func (x *InitOrgRequest) GetAdminName() string {
	if x != nil {
		return x.AdminName
	}
	return ""
}

func (x *InitOrgRequest) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *InitOrgRequest) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

func (x *InitOrgRequest) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *InitOrgRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *InitOrgRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *InitOrgRequest) GetFirstLoginPending() bool {
	if x != nil {
		return x.FirstLoginPending
	}
	return false
}

func (x *InitOrgRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InitOrgRequest) GetUseGoogleOAuth() bool {
	if x != nil {
		return x.UseGoogleOAuth
	}
	return false
}

func (x *InitOrgRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *InitOrgRequest) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

type InitOrgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsError bool   `protobuf:"varint,1,opt,name=isError,proto3" json:"isError,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InitOrgResponse) Reset() {
	*x = InitOrgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_org_type_organization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitOrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitOrgResponse) ProtoMessage() {}

func (x *InitOrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_org_type_organization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitOrgResponse.ProtoReflect.Descriptor instead.
func (*InitOrgResponse) Descriptor() ([]byte, []int) {
	return file_proto_org_type_organization_proto_rawDescGZIP(), []int{1}
}

func (x *InitOrgResponse) GetIsError() bool {
	if x != nil {
		return x.IsError
	}
	return false
}

func (x *InitOrgResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_org_type_organization_proto protoreflect.FileDescriptor

var file_proto_org_type_organization_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x72, 0x67, 0x22, 0x8e, 0x04, 0x0a, 0x0e, 0x49, 0x6e, 0x69,
	0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x75,
	0x73, 0x65, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x45, 0x0a, 0x0f, 0x49, 0x6e, 0x69,
	0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_org_type_organization_proto_rawDescOnce sync.Once
	file_proto_org_type_organization_proto_rawDescData = file_proto_org_type_organization_proto_rawDesc
)

func file_proto_org_type_organization_proto_rawDescGZIP() []byte {
	file_proto_org_type_organization_proto_rawDescOnce.Do(func() {
		file_proto_org_type_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_org_type_organization_proto_rawDescData)
	})
	return file_proto_org_type_organization_proto_rawDescData
}

var file_proto_org_type_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_org_type_organization_proto_goTypes = []interface{}{
	(*InitOrgRequest)(nil),  // 0: org.InitOrgRequest
	(*InitOrgResponse)(nil), // 1: org.InitOrgResponse
}
var file_proto_org_type_organization_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_org_type_organization_proto_init() }
func file_proto_org_type_organization_proto_init() {
	if File_proto_org_type_organization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_org_type_organization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitOrgRequest); i {
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
		file_proto_org_type_organization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitOrgResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_org_type_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_org_type_organization_proto_goTypes,
		DependencyIndexes: file_proto_org_type_organization_proto_depIdxs,
		MessageInfos:      file_proto_org_type_organization_proto_msgTypes,
	}.Build()
	File_proto_org_type_organization_proto = out.File
	file_proto_org_type_organization_proto_rawDesc = nil
	file_proto_org_type_organization_proto_goTypes = nil
	file_proto_org_type_organization_proto_depIdxs = nil
}
