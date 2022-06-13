// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.21.1
// source: match/v1/common.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status int32

const (
	// HTTP 200
	Status_STATUS_OK Status = 0
	// HTTP 400
	Status_STATUS_BAD_REQUEST Status = 1
	// HTTP 401
	Status_STATUS_UNAUTHENTICATED Status = 5
	// HTTP 403
	Status_STATUS_FORBIDDEN Status = 8
	// HTTP 404
	Status_STATUS_NOT_FOUND Status = 4
	// HTTP 422
	Status_STATUS_UNPROCESSABLE Status = 3
	// HTTP 500
	Status_STATUS_INTERNAL Status = 2
	// HTTP 503
	Status_STATUS_UNAVAILABLE Status = 6
	// HTTP 504
	Status_STATUS_TIMEOUT Status = 7
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_OK",
		1: "STATUS_BAD_REQUEST",
		5: "STATUS_UNAUTHENTICATED",
		8: "STATUS_FORBIDDEN",
		4: "STATUS_NOT_FOUND",
		3: "STATUS_UNPROCESSABLE",
		2: "STATUS_INTERNAL",
		6: "STATUS_UNAVAILABLE",
		7: "STATUS_TIMEOUT",
	}
	Status_value = map[string]int32{
		"STATUS_OK":              0,
		"STATUS_BAD_REQUEST":     1,
		"STATUS_UNAUTHENTICATED": 5,
		"STATUS_FORBIDDEN":       8,
		"STATUS_NOT_FOUND":       4,
		"STATUS_UNPROCESSABLE":   3,
		"STATUS_INTERNAL":        2,
		"STATUS_UNAVAILABLE":     6,
		"STATUS_TIMEOUT":         7,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_match_v1_common_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_match_v1_common_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_match_v1_common_proto_rawDescGZIP(), []int{0}
}

type SortField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Desc bool   `protobuf:"varint,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *SortField) Reset() {
	*x = SortField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortField) ProtoMessage() {}

func (x *SortField) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortField.ProtoReflect.Descriptor instead.
func (*SortField) Descriptor() ([]byte, []int) {
	return file_match_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *SortField) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SortField) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

type FieldError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field   string            `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Params  map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Code    string            `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *FieldError) Reset() {
	*x = FieldError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldError) ProtoMessage() {}

func (x *FieldError) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldError.ProtoReflect.Descriptor instead.
func (*FieldError) Descriptor() ([]byte, []int) {
	return file_match_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *FieldError) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *FieldError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FieldError) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *FieldError) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type UnprocessableDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*FieldError `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *UnprocessableDetails) Reset() {
	*x = UnprocessableDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnprocessableDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnprocessableDetails) ProtoMessage() {}

func (x *UnprocessableDetails) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnprocessableDetails.ProtoReflect.Descriptor instead.
func (*UnprocessableDetails) Descriptor() ([]byte, []int) {
	return file_match_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *UnprocessableDetails) GetFields() []*FieldError {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Status `protobuf:"varint,1,opt,name=code,proto3,enum=match.v1.Status" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Types that are assignable to Details:
	//	*Error_Unprocessable
	Details isError_Details `protobuf_oneof:"details"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_match_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCode() Status {
	if x != nil {
		return x.Code
	}
	return Status_STATUS_OK
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (m *Error) GetDetails() isError_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (x *Error) GetUnprocessable() *UnprocessableDetails {
	if x, ok := x.GetDetails().(*Error_Unprocessable); ok {
		return x.Unprocessable
	}
	return nil
}

type isError_Details interface {
	isError_Details()
}

type Error_Unprocessable struct {
	Unprocessable *UnprocessableDetails `protobuf:"bytes,3,opt,name=unprocessable,proto3,oneof"`
}

func (*Error_Unprocessable) isError_Details() {}

var File_match_v1_common_proto protoreflect.FileDescriptor

var file_match_v1_common_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x22, 0x31, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x22, 0xc5, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x44, 0x0a, 0x14,
	0x55, 0x6e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0d,
	0x75, 0x6e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x6e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x6e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2a,
	0xd2, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x01, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x41, 0x55,
	0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45,
	0x4e, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x06,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x07, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_v1_common_proto_rawDescOnce sync.Once
	file_match_v1_common_proto_rawDescData = file_match_v1_common_proto_rawDesc
)

func file_match_v1_common_proto_rawDescGZIP() []byte {
	file_match_v1_common_proto_rawDescOnce.Do(func() {
		file_match_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_v1_common_proto_rawDescData)
	})
	return file_match_v1_common_proto_rawDescData
}

var file_match_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_match_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_match_v1_common_proto_goTypes = []interface{}{
	(Status)(0),                  // 0: match.v1.Status
	(*SortField)(nil),            // 1: match.v1.SortField
	(*FieldError)(nil),           // 2: match.v1.FieldError
	(*UnprocessableDetails)(nil), // 3: match.v1.UnprocessableDetails
	(*Error)(nil),                // 4: match.v1.Error
	nil,                          // 5: match.v1.FieldError.ParamsEntry
}
var file_match_v1_common_proto_depIdxs = []int32{
	5, // 0: match.v1.FieldError.params:type_name -> match.v1.FieldError.ParamsEntry
	2, // 1: match.v1.UnprocessableDetails.fields:type_name -> match.v1.FieldError
	0, // 2: match.v1.Error.code:type_name -> match.v1.Status
	3, // 3: match.v1.Error.unprocessable:type_name -> match.v1.UnprocessableDetails
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_match_v1_common_proto_init() }
func file_match_v1_common_proto_init() {
	if File_match_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortField); i {
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
		file_match_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldError); i {
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
		file_match_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnprocessableDetails); i {
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
		file_match_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
	file_match_v1_common_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Error_Unprocessable)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_match_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_match_v1_common_proto_goTypes,
		DependencyIndexes: file_match_v1_common_proto_depIdxs,
		EnumInfos:         file_match_v1_common_proto_enumTypes,
		MessageInfos:      file_match_v1_common_proto_msgTypes,
	}.Build()
	File_match_v1_common_proto = out.File
	file_match_v1_common_proto_rawDesc = nil
	file_match_v1_common_proto_goTypes = nil
	file_match_v1_common_proto_depIdxs = nil
}
