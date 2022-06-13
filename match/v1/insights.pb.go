// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.21.1
// source: match/v1/insights.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Insights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComputedAt                   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=computed_at,json=computedAt,proto3" json:"computed_at,omitempty"`
	Emails                       int64                  `protobuf:"varint,2,opt,name=emails,proto3" json:"emails,omitempty"`
	AppleIdfas                   int64                  `protobuf:"varint,3,opt,name=apple_idfas,json=appleIdfas,proto3" json:"apple_idfas,omitempty"`
	GoogleGaids                  int64                  `protobuf:"varint,4,opt,name=google_gaids,json=googleGaids,proto3" json:"google_gaids,omitempty"`
	Ipv4S                        int64                  `protobuf:"varint,5,opt,name=ipv4s,proto3" json:"ipv4s,omitempty"`
	Ipv6S                        int64                  `protobuf:"varint,6,opt,name=ipv6s,proto3" json:"ipv6s,omitempty"`
	SamsungTifas                 int64                  `protobuf:"varint,7,opt,name=samsung_tifas,json=samsungTifas,proto3" json:"samsung_tifas,omitempty"`
	RokuRidas                    int64                  `protobuf:"varint,8,opt,name=roku_ridas,json=rokuRidas,proto3" json:"roku_ridas,omitempty"`
	AmazonAfais                  int64                  `protobuf:"varint,9,opt,name=amazon_afais,json=amazonAfais,proto3" json:"amazon_afais,omitempty"`
	PhoneNumbers                 int64                  `protobuf:"varint,10,opt,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	NetidTpids                   int64                  `protobuf:"varint,12,opt,name=netid_tpids,json=netidTpids,proto3" json:"netid_tpids,omitempty"`
	DifferentialPrivacyThreshold int32                  `protobuf:"varint,11,opt,name=differential_privacy_threshold,json=differentialPrivacyThreshold,proto3" json:"differential_privacy_threshold,omitempty"`
}

func (x *Insights) Reset() {
	*x = Insights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_insights_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insights) ProtoMessage() {}

func (x *Insights) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_insights_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insights.ProtoReflect.Descriptor instead.
func (*Insights) Descriptor() ([]byte, []int) {
	return file_match_v1_insights_proto_rawDescGZIP(), []int{0}
}

func (x *Insights) GetComputedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ComputedAt
	}
	return nil
}

func (x *Insights) GetEmails() int64 {
	if x != nil {
		return x.Emails
	}
	return 0
}

func (x *Insights) GetAppleIdfas() int64 {
	if x != nil {
		return x.AppleIdfas
	}
	return 0
}

func (x *Insights) GetGoogleGaids() int64 {
	if x != nil {
		return x.GoogleGaids
	}
	return 0
}

func (x *Insights) GetIpv4S() int64 {
	if x != nil {
		return x.Ipv4S
	}
	return 0
}

func (x *Insights) GetIpv6S() int64 {
	if x != nil {
		return x.Ipv6S
	}
	return 0
}

func (x *Insights) GetSamsungTifas() int64 {
	if x != nil {
		return x.SamsungTifas
	}
	return 0
}

func (x *Insights) GetRokuRidas() int64 {
	if x != nil {
		return x.RokuRidas
	}
	return 0
}

func (x *Insights) GetAmazonAfais() int64 {
	if x != nil {
		return x.AmazonAfais
	}
	return 0
}

func (x *Insights) GetPhoneNumbers() int64 {
	if x != nil {
		return x.PhoneNumbers
	}
	return 0
}

func (x *Insights) GetNetidTpids() int64 {
	if x != nil {
		return x.NetidTpids
	}
	return 0
}

func (x *Insights) GetDifferentialPrivacyThreshold() int32 {
	if x != nil {
		return x.DifferentialPrivacyThreshold
	}
	return 0
}

var File_match_v1_insights_proto protoreflect.FileDescriptor

var file_match_v1_insights_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x03, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x66, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x70, 0x70,
	0x6c, 0x65, 0x49, 0x64, 0x66, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x67, 0x61, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x47, 0x61, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x70,
	0x76, 0x34, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x70, 0x76, 0x34, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x70, 0x76, 0x36, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x69, 0x70, 0x76, 0x36, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x61, 0x6d, 0x73, 0x75, 0x6e,
	0x67, 0x5f, 0x74, 0x69, 0x66, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73,
	0x61, 0x6d, 0x73, 0x75, 0x6e, 0x67, 0x54, 0x69, 0x66, 0x61, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x6f, 0x6b, 0x75, 0x5f, 0x72, 0x69, 0x64, 0x61, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x72, 0x6f, 0x6b, 0x75, 0x52, 0x69, 0x64, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6d,
	0x61, 0x7a, 0x6f, 0x6e, 0x5f, 0x61, 0x66, 0x61, 0x69, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x41, 0x66, 0x61, 0x69, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x69, 0x64, 0x5f, 0x74, 0x70, 0x69, 0x64,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x69, 0x64, 0x54, 0x70,
	0x69, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x1e, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1c, 0x64, 0x69, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_v1_insights_proto_rawDescOnce sync.Once
	file_match_v1_insights_proto_rawDescData = file_match_v1_insights_proto_rawDesc
)

func file_match_v1_insights_proto_rawDescGZIP() []byte {
	file_match_v1_insights_proto_rawDescOnce.Do(func() {
		file_match_v1_insights_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_v1_insights_proto_rawDescData)
	})
	return file_match_v1_insights_proto_rawDescData
}

var file_match_v1_insights_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_match_v1_insights_proto_goTypes = []interface{}{
	(*Insights)(nil),              // 0: match.v1.Insights
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_match_v1_insights_proto_depIdxs = []int32{
	1, // 0: match.v1.Insights.computed_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_match_v1_insights_proto_init() }
func file_match_v1_insights_proto_init() {
	if File_match_v1_insights_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_v1_insights_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insights); i {
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
			RawDescriptor: file_match_v1_insights_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_match_v1_insights_proto_goTypes,
		DependencyIndexes: file_match_v1_insights_proto_depIdxs,
		MessageInfos:      file_match_v1_insights_proto_msgTypes,
	}.Build()
	File_match_v1_insights_proto = out.File
	file_match_v1_insights_proto_rawDesc = nil
	file_match_v1_insights_proto_goTypes = nil
	file_match_v1_insights_proto_depIdxs = nil
}
