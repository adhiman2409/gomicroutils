// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto/org/type/org_attendance.proto

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

type OrgAttendanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgName string `protobuf:"bytes,1,opt,name=orgName,proto3" json:"orgName,omitempty"`
	Domain  string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *OrgAttendanceRequest) Reset() {
	*x = OrgAttendanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_org_type_org_attendance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgAttendanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgAttendanceRequest) ProtoMessage() {}

func (x *OrgAttendanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_org_type_org_attendance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgAttendanceRequest.ProtoReflect.Descriptor instead.
func (*OrgAttendanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_org_type_org_attendance_proto_rawDescGZIP(), []int{0}
}

func (x *OrgAttendanceRequest) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *OrgAttendanceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type OrgAttendanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId                   string   `protobuf:"bytes,1,opt,name=orgId,proto3" json:"orgId,omitempty"`
	OrgName                 string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	CaptureCheckInLocation  bool     `protobuf:"varint,3,opt,name=captureCheckInLocation,proto3" json:"captureCheckInLocation,omitempty"`
	EnforceLocationCheckIn  bool     `protobuf:"varint,4,opt,name=enforceLocationCheckIn,proto3" json:"enforceLocationCheckIn,omitempty"`
	EnforceLocationCheckOut bool     `protobuf:"varint,5,opt,name=enforceLocationCheckOut,proto3" json:"enforceLocationCheckOut,omitempty"`
	RequiredLat             float64  `protobuf:"fixed64,6,opt,name=requiredLat,proto3" json:"requiredLat,omitempty"`
	RequiredLng             float64  `protobuf:"fixed64,7,opt,name=requiredLng,proto3" json:"requiredLng,omitempty"`
	FlexiHoursEnabled       bool     `protobuf:"varint,8,opt,name=flexiHoursEnabled,proto3" json:"flexiHoursEnabled,omitempty"`
	FlexiHourWindowInMin    int32    `protobuf:"varint,9,opt,name=flexiHourWindowInMin,proto3" json:"flexiHourWindowInMin,omitempty"`
	OrgCheckInTime          string   `protobuf:"bytes,10,opt,name=orgCheckInTime,proto3" json:"orgCheckInTime,omitempty"`
	OrgCheckOutTime         string   `protobuf:"bytes,11,opt,name=orgCheckOutTime,proto3" json:"orgCheckOutTime,omitempty"`
	CheckInMarginInMin      int32    `protobuf:"varint,12,opt,name=checkInMarginInMin,proto3" json:"checkInMarginInMin,omitempty"`
	DailyWorkingHours       float32  `protobuf:"fixed32,13,opt,name=dailyWorkingHours,proto3" json:"dailyWorkingHours,omitempty"`
	WorkingDaysPerWeek      float32  `protobuf:"fixed32,14,opt,name=workingDaysPerWeek,proto3" json:"workingDaysPerWeek,omitempty"`
	WeeklyOffDays           []string `protobuf:"bytes,15,rep,name=weeklyOffDays,proto3" json:"weeklyOffDays,omitempty"`
}

func (x *OrgAttendanceResponse) Reset() {
	*x = OrgAttendanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_org_type_org_attendance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgAttendanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgAttendanceResponse) ProtoMessage() {}

func (x *OrgAttendanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_org_type_org_attendance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgAttendanceResponse.ProtoReflect.Descriptor instead.
func (*OrgAttendanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_org_type_org_attendance_proto_rawDescGZIP(), []int{1}
}

func (x *OrgAttendanceResponse) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *OrgAttendanceResponse) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *OrgAttendanceResponse) GetCaptureCheckInLocation() bool {
	if x != nil {
		return x.CaptureCheckInLocation
	}
	return false
}

func (x *OrgAttendanceResponse) GetEnforceLocationCheckIn() bool {
	if x != nil {
		return x.EnforceLocationCheckIn
	}
	return false
}

func (x *OrgAttendanceResponse) GetEnforceLocationCheckOut() bool {
	if x != nil {
		return x.EnforceLocationCheckOut
	}
	return false
}

func (x *OrgAttendanceResponse) GetRequiredLat() float64 {
	if x != nil {
		return x.RequiredLat
	}
	return 0
}

func (x *OrgAttendanceResponse) GetRequiredLng() float64 {
	if x != nil {
		return x.RequiredLng
	}
	return 0
}

func (x *OrgAttendanceResponse) GetFlexiHoursEnabled() bool {
	if x != nil {
		return x.FlexiHoursEnabled
	}
	return false
}

func (x *OrgAttendanceResponse) GetFlexiHourWindowInMin() int32 {
	if x != nil {
		return x.FlexiHourWindowInMin
	}
	return 0
}

func (x *OrgAttendanceResponse) GetOrgCheckInTime() string {
	if x != nil {
		return x.OrgCheckInTime
	}
	return ""
}

func (x *OrgAttendanceResponse) GetOrgCheckOutTime() string {
	if x != nil {
		return x.OrgCheckOutTime
	}
	return ""
}

func (x *OrgAttendanceResponse) GetCheckInMarginInMin() int32 {
	if x != nil {
		return x.CheckInMarginInMin
	}
	return 0
}

func (x *OrgAttendanceResponse) GetDailyWorkingHours() float32 {
	if x != nil {
		return x.DailyWorkingHours
	}
	return 0
}

func (x *OrgAttendanceResponse) GetWorkingDaysPerWeek() float32 {
	if x != nil {
		return x.WorkingDaysPerWeek
	}
	return 0
}

func (x *OrgAttendanceResponse) GetWeeklyOffDays() []string {
	if x != nil {
		return x.WeeklyOffDays
	}
	return nil
}

var File_proto_org_type_org_attendance_proto protoreflect.FileDescriptor

var file_proto_org_type_org_attendance_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x72, 0x67, 0x22, 0x48, 0x0a, 0x14, 0x4f, 0x72,
	0x67, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x9d, 0x05, 0x0a, 0x15, 0x4f, 0x72, 0x67, 0x41, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36,
	0x0a, 0x16, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x16, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x38,
	0x0a, 0x17, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x17, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x11,
	0x66, 0x6c, 0x65, 0x78, 0x69, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x66, 0x6c, 0x65, 0x78, 0x69, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x66, 0x6c,
	0x65, 0x78, 0x69, 0x48, 0x6f, 0x75, 0x72, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x4d,
	0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x66, 0x6c, 0x65, 0x78, 0x69, 0x48,
	0x6f, 0x75, 0x72, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x26,
	0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6f, 0x72, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x4d, 0x61, 0x72, 0x67, 0x69,
	0x6e, 0x49, 0x6e, 0x4d, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x4d, 0x69, 0x6e,
	0x12, 0x2c, 0x0a, 0x11, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x2e,
	0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x73, 0x50, 0x65, 0x72,
	0x57, 0x65, 0x65, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x73, 0x50, 0x65, 0x72, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x24,
	0x0a, 0x0d, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x4f, 0x66, 0x66, 0x44, 0x61, 0x79, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x4f, 0x66, 0x66,
	0x44, 0x61, 0x79, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_org_type_org_attendance_proto_rawDescOnce sync.Once
	file_proto_org_type_org_attendance_proto_rawDescData = file_proto_org_type_org_attendance_proto_rawDesc
)

func file_proto_org_type_org_attendance_proto_rawDescGZIP() []byte {
	file_proto_org_type_org_attendance_proto_rawDescOnce.Do(func() {
		file_proto_org_type_org_attendance_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_org_type_org_attendance_proto_rawDescData)
	})
	return file_proto_org_type_org_attendance_proto_rawDescData
}

var file_proto_org_type_org_attendance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_org_type_org_attendance_proto_goTypes = []interface{}{
	(*OrgAttendanceRequest)(nil),  // 0: org.OrgAttendanceRequest
	(*OrgAttendanceResponse)(nil), // 1: org.OrgAttendanceResponse
}
var file_proto_org_type_org_attendance_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_org_type_org_attendance_proto_init() }
func file_proto_org_type_org_attendance_proto_init() {
	if File_proto_org_type_org_attendance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_org_type_org_attendance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgAttendanceRequest); i {
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
		file_proto_org_type_org_attendance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgAttendanceResponse); i {
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
			RawDescriptor: file_proto_org_type_org_attendance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_org_type_org_attendance_proto_goTypes,
		DependencyIndexes: file_proto_org_type_org_attendance_proto_depIdxs,
		MessageInfos:      file_proto_org_type_org_attendance_proto_msgTypes,
	}.Build()
	File_proto_org_type_org_attendance_proto = out.File
	file_proto_org_type_org_attendance_proto_rawDesc = nil
	file_proto_org_type_org_attendance_proto_goTypes = nil
	file_proto_org_type_org_attendance_proto_depIdxs = nil
}
