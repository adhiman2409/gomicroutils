// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: proto/auth/type/update_info.proto

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

type InfoUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ImageURL    string `protobuf:"bytes,2,opt,name=imageURL,proto3" json:"imageURL,omitempty"`
	Role        string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Designation string `protobuf:"bytes,4,opt,name=designation,proto3" json:"designation,omitempty"`
	Department  string `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *InfoUpdateRequest) Reset() {
	*x = InfoUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_type_update_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoUpdateRequest) ProtoMessage() {}

func (x *InfoUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_type_update_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoUpdateRequest.ProtoReflect.Descriptor instead.
func (*InfoUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_type_update_info_proto_rawDescGZIP(), []int{0}
}

func (x *InfoUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InfoUpdateRequest) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *InfoUpdateRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *InfoUpdateRequest) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

func (x *InfoUpdateRequest) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

type InfoUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsError bool   `protobuf:"varint,1,opt,name=isError,proto3" json:"isError,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InfoUpdateResponse) Reset() {
	*x = InfoUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_type_update_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoUpdateResponse) ProtoMessage() {}

func (x *InfoUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_type_update_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoUpdateResponse.ProtoReflect.Descriptor instead.
func (*InfoUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_type_update_info_proto_rawDescGZIP(), []int{1}
}

func (x *InfoUpdateResponse) GetIsError() bool {
	if x != nil {
		return x.IsError
	}
	return false
}

func (x *InfoUpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_auth_type_update_info_proto protoreflect.FileDescriptor

var file_proto_auth_type_update_info_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x49, 0x6e,
	0x66, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x48, 0x0a, 0x12, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_type_update_info_proto_rawDescOnce sync.Once
	file_proto_auth_type_update_info_proto_rawDescData = file_proto_auth_type_update_info_proto_rawDesc
)

func file_proto_auth_type_update_info_proto_rawDescGZIP() []byte {
	file_proto_auth_type_update_info_proto_rawDescOnce.Do(func() {
		file_proto_auth_type_update_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_type_update_info_proto_rawDescData)
	})
	return file_proto_auth_type_update_info_proto_rawDescData
}

var file_proto_auth_type_update_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_auth_type_update_info_proto_goTypes = []interface{}{
	(*InfoUpdateRequest)(nil),  // 0: auth.InfoUpdateRequest
	(*InfoUpdateResponse)(nil), // 1: auth.InfoUpdateResponse
}
var file_proto_auth_type_update_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_type_update_info_proto_init() }
func file_proto_auth_type_update_info_proto_init() {
	if File_proto_auth_type_update_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_type_update_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoUpdateRequest); i {
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
		file_proto_auth_type_update_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoUpdateResponse); i {
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
			RawDescriptor: file_proto_auth_type_update_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_auth_type_update_info_proto_goTypes,
		DependencyIndexes: file_proto_auth_type_update_info_proto_depIdxs,
		MessageInfos:      file_proto_auth_type_update_info_proto_msgTypes,
	}.Build()
	File_proto_auth_type_update_info_proto = out.File
	file_proto_auth_type_update_info_proto_rawDesc = nil
	file_proto_auth_type_update_info_proto_goTypes = nil
	file_proto_auth_type_update_info_proto_depIdxs = nil
}
