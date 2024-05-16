// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: proto/auth/type/verify.proto

package auth

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

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RouteName string `protobuf:"bytes,2,opt,name=routeName,proto3" json:"routeName,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_type_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_type_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_type_verify_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *VerifyRequest) GetRouteName() string {
	if x != nil {
		return x.RouteName
	}
	return ""
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAuthorised bool   `protobuf:"varint,1,opt,name=isAuthorised,proto3" json:"isAuthorised,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	EmailId      string `protobuf:"bytes,4,opt,name=emailId,proto3" json:"emailId,omitempty"`
	PhoneNumber  string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Tenant       string `protobuf:"bytes,6,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Domain       string `protobuf:"bytes,7,opt,name=domain,proto3" json:"domain,omitempty"`
	Department   string `protobuf:"bytes,8,opt,name=department,proto3" json:"department,omitempty"`
	Role         string `protobuf:"bytes,9,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_type_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_type_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_type_verify_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyResponse) GetIsAuthorised() bool {
	if x != nil {
		return x.IsAuthorised
	}
	return false
}

func (x *VerifyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VerifyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VerifyResponse) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

func (x *VerifyResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *VerifyResponse) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *VerifyResponse) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *VerifyResponse) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *VerifyResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_proto_auth_type_verify_proto protoreflect.FileDescriptor

var file_proto_auth_type_verify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x22, 0x43, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x82, 0x02, 0x0a, 0x0e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x73, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x0f,
	0x5a, 0x0d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_type_verify_proto_rawDescOnce sync.Once
	file_proto_auth_type_verify_proto_rawDescData = file_proto_auth_type_verify_proto_rawDesc
)

func file_proto_auth_type_verify_proto_rawDescGZIP() []byte {
	file_proto_auth_type_verify_proto_rawDescOnce.Do(func() {
		file_proto_auth_type_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_type_verify_proto_rawDescData)
	})
	return file_proto_auth_type_verify_proto_rawDescData
}

var file_proto_auth_type_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_auth_type_verify_proto_goTypes = []interface{}{
	(*VerifyRequest)(nil),  // 0: auth.VerifyRequest
	(*VerifyResponse)(nil), // 1: auth.VerifyResponse
}
var file_proto_auth_type_verify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_type_verify_proto_init() }
func file_proto_auth_type_verify_proto_init() {
	if File_proto_auth_type_verify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_type_verify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_proto_auth_type_verify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
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
			RawDescriptor: file_proto_auth_type_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_auth_type_verify_proto_goTypes,
		DependencyIndexes: file_proto_auth_type_verify_proto_depIdxs,
		MessageInfos:      file_proto_auth_type_verify_proto_msgTypes,
	}.Build()
	File_proto_auth_type_verify_proto = out.File
	file_proto_auth_type_verify_proto_rawDesc = nil
	file_proto_auth_type_verify_proto_goTypes = nil
	file_proto_auth_type_verify_proto_depIdxs = nil
}
