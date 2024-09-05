// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.1
// source: match/v1/identifier.proto

package v1

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

type IdKind int32

const (
	IdKind_ID_KIND_UNKNOWN IdKind = 0
	// Normalize email (lowercased and white spaces removed) then hashed with
	// SHA256 checksum.
	IdKind_ID_KIND_EMAIL_HASH IdKind = 1
	// ID_KIND_APPLE_IDFA, Apple iOS IDFA, ID For Advertising, is a unique UUID
	// per physical iOS device.
	IdKind_ID_KIND_APPLE_IDFA IdKind = 2
	// ID_KIND_GOOGLE_GAID, Google Advertising ID (GAID) is a unique UUID per
	// physical android device.
	IdKind_ID_KIND_GOOGLE_GAID IdKind = 3
	// ID_KIND_IPV4 is the type representing Internet IPv4 addresses.
	IdKind_ID_KIND_IPV4 IdKind = 4
	// ID_KIND_IPV6 is the type representing Internet IPv6 addresses.
	IdKind_ID_KIND_IPV6 IdKind = 5
	// ID_KIND_SAMSUNG_TIFA, Samsung ID for a unique physical Samsung smart TV device.
	IdKind_ID_KIND_SAMSUNG_TIFA IdKind = 6
	// ID_KIND_ROKU_RIDA, Roku ID for Advertising.
	IdKind_ID_KIND_ROKU_RIDA IdKind = 7
	// ID_KIND_AMAZON_AFAI, Amazon Advertising ID.
	IdKind_ID_KIND_AMAZON_AFAI IdKind = 8
	// IdKind_ID_KIND_PHONE_NUMBER, Phone number.
	IdKind_ID_KIND_PHONE_NUMBER IdKind = 9
	// IdKind_ID_KIND_NETID, NetID is a unique TPID.
	IdKind_ID_KIND_NETID IdKind = 10
	// IdKind_ID_KIND_POSTAL_CODE, Postal code.
	IdKind_ID_KIND_POSTAL_CODE IdKind = 11
	// IdKind_ID_KIND_ID5, Decrypted ID5.
	IdKind_ID_KIND_ID5 IdKind = 12
	// IdKind_ID_KIND_UTIQ, UTIQ.
	IdKind_ID_KIND_UTIQ IdKind = 13
)

// Enum value maps for IdKind.
var (
	IdKind_name = map[int32]string{
		0:  "ID_KIND_UNKNOWN",
		1:  "ID_KIND_EMAIL_HASH",
		2:  "ID_KIND_APPLE_IDFA",
		3:  "ID_KIND_GOOGLE_GAID",
		4:  "ID_KIND_IPV4",
		5:  "ID_KIND_IPV6",
		6:  "ID_KIND_SAMSUNG_TIFA",
		7:  "ID_KIND_ROKU_RIDA",
		8:  "ID_KIND_AMAZON_AFAI",
		9:  "ID_KIND_PHONE_NUMBER",
		10: "ID_KIND_NETID",
		11: "ID_KIND_POSTAL_CODE",
		12: "ID_KIND_ID5",
		13: "ID_KIND_UTIQ",
	}
	IdKind_value = map[string]int32{
		"ID_KIND_UNKNOWN":      0,
		"ID_KIND_EMAIL_HASH":   1,
		"ID_KIND_APPLE_IDFA":   2,
		"ID_KIND_GOOGLE_GAID":  3,
		"ID_KIND_IPV4":         4,
		"ID_KIND_IPV6":         5,
		"ID_KIND_SAMSUNG_TIFA": 6,
		"ID_KIND_ROKU_RIDA":    7,
		"ID_KIND_AMAZON_AFAI":  8,
		"ID_KIND_PHONE_NUMBER": 9,
		"ID_KIND_NETID":        10,
		"ID_KIND_POSTAL_CODE":  11,
		"ID_KIND_ID5":          12,
		"ID_KIND_UTIQ":         13,
	}
)

func (x IdKind) Enum() *IdKind {
	p := new(IdKind)
	*p = x
	return p
}

func (x IdKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdKind) Descriptor() protoreflect.EnumDescriptor {
	return file_match_v1_identifier_proto_enumTypes[0].Descriptor()
}

func (IdKind) Type() protoreflect.EnumType {
	return &file_match_v1_identifier_proto_enumTypes[0]
}

func (x IdKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdKind.Descriptor instead.
func (IdKind) EnumDescriptor() ([]byte, []int) {
	return file_match_v1_identifier_proto_rawDescGZIP(), []int{0}
}

var File_match_v1_identifier_proto protoreflect.FileDescriptor

var file_match_v1_identifier_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2a, 0xbd, 0x02, 0x0a, 0x06, 0x49, 0x64, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x45, 0x5f, 0x49,
	0x44, 0x46, 0x41, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x47, 0x41, 0x49, 0x44, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x50, 0x56, 0x34, 0x10, 0x04,
	0x12, 0x10, 0x0a, 0x0c, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x50, 0x56, 0x36,
	0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x41,
	0x4d, 0x53, 0x55, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x46, 0x41, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11,
	0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x52, 0x4f, 0x4b, 0x55, 0x5f, 0x52, 0x49, 0x44,
	0x41, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41,
	0x4d, 0x41, 0x5a, 0x4f, 0x4e, 0x5f, 0x41, 0x46, 0x41, 0x49, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14,
	0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55,
	0x4d, 0x42, 0x45, 0x52, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x4e, 0x45, 0x54, 0x49, 0x44, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x44,
	0x35, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55,
	0x54, 0x49, 0x51, 0x10, 0x0d, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_v1_identifier_proto_rawDescOnce sync.Once
	file_match_v1_identifier_proto_rawDescData = file_match_v1_identifier_proto_rawDesc
)

func file_match_v1_identifier_proto_rawDescGZIP() []byte {
	file_match_v1_identifier_proto_rawDescOnce.Do(func() {
		file_match_v1_identifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_v1_identifier_proto_rawDescData)
	})
	return file_match_v1_identifier_proto_rawDescData
}

var file_match_v1_identifier_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_match_v1_identifier_proto_goTypes = []interface{}{
	(IdKind)(0), // 0: match.v1.IdKind
}
var file_match_v1_identifier_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_match_v1_identifier_proto_init() }
func file_match_v1_identifier_proto_init() {
	if File_match_v1_identifier_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_match_v1_identifier_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_match_v1_identifier_proto_goTypes,
		DependencyIndexes: file_match_v1_identifier_proto_depIdxs,
		EnumInfos:         file_match_v1_identifier_proto_enumTypes,
	}.Build()
	File_match_v1_identifier_proto = out.File
	file_match_v1_identifier_proto_rawDesc = nil
	file_match_v1_identifier_proto_goTypes = nil
	file_match_v1_identifier_proto_depIdxs = nil
}
