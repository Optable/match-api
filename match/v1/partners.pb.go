// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.1
// source: match/v1/partners.proto

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

type SnowflakePartner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Token used to authenticate the external partner.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Required. The region of the partner's snowflake account
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Required. The account locator ID of the snowflake partner
	AccountLocatorId string `protobuf:"bytes,3,opt,name=account_locator_id,json=accountLocatorId,proto3" json:"account_locator_id,omitempty"`
}

func (x *SnowflakePartner) Reset() {
	*x = SnowflakePartner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_partners_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnowflakePartner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnowflakePartner) ProtoMessage() {}

func (x *SnowflakePartner) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_partners_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnowflakePartner.ProtoReflect.Descriptor instead.
func (*SnowflakePartner) Descriptor() ([]byte, []int) {
	return file_match_v1_partners_proto_rawDescGZIP(), []int{0}
}

func (x *SnowflakePartner) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SnowflakePartner) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *SnowflakePartner) GetAccountLocatorId() string {
	if x != nil {
		return x.AccountLocatorId
	}
	return ""
}

type DCNPartner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Token used to authenticate the external partner.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DCNPartner) Reset() {
	*x = DCNPartner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_partners_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DCNPartner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DCNPartner) ProtoMessage() {}

func (x *DCNPartner) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_partners_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DCNPartner.ProtoReflect.Descriptor instead.
func (*DCNPartner) Descriptor() ([]byte, []int) {
	return file_match_v1_partners_proto_rawDescGZIP(), []int{1}
}

func (x *DCNPartner) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type HeadlessPartner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Public Key of the external partner.
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Required. Token used to authenticate the external partner.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *HeadlessPartner) Reset() {
	*x = HeadlessPartner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_partners_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeadlessPartner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeadlessPartner) ProtoMessage() {}

func (x *HeadlessPartner) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_partners_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeadlessPartner.ProtoReflect.Descriptor instead.
func (*HeadlessPartner) Descriptor() ([]byte, []int) {
	return file_match_v1_partners_proto_rawDescGZIP(), []int{2}
}

func (x *HeadlessPartner) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *HeadlessPartner) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RegisterPartnerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key that will be used to authenticate this partner.
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Secret token used to identify the partner invitation.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// Optional. Partnership kind aware registration request
	//
	// Types that are assignable to PartnerInfo:
	//
	//	*RegisterPartnerReq_DcnPartner
	//	*RegisterPartnerReq_HeadlessPartner
	//	*RegisterPartnerReq_SnowflakePartner
	PartnerInfo isRegisterPartnerReq_PartnerInfo `protobuf_oneof:"PartnerInfo"`
}

func (x *RegisterPartnerReq) Reset() {
	*x = RegisterPartnerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_partners_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPartnerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPartnerReq) ProtoMessage() {}

func (x *RegisterPartnerReq) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_partners_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPartnerReq.ProtoReflect.Descriptor instead.
func (*RegisterPartnerReq) Descriptor() ([]byte, []int) {
	return file_match_v1_partners_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterPartnerReq) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *RegisterPartnerReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (m *RegisterPartnerReq) GetPartnerInfo() isRegisterPartnerReq_PartnerInfo {
	if m != nil {
		return m.PartnerInfo
	}
	return nil
}

func (x *RegisterPartnerReq) GetDcnPartner() *DCNPartner {
	if x, ok := x.GetPartnerInfo().(*RegisterPartnerReq_DcnPartner); ok {
		return x.DcnPartner
	}
	return nil
}

func (x *RegisterPartnerReq) GetHeadlessPartner() *HeadlessPartner {
	if x, ok := x.GetPartnerInfo().(*RegisterPartnerReq_HeadlessPartner); ok {
		return x.HeadlessPartner
	}
	return nil
}

func (x *RegisterPartnerReq) GetSnowflakePartner() *SnowflakePartner {
	if x, ok := x.GetPartnerInfo().(*RegisterPartnerReq_SnowflakePartner); ok {
		return x.SnowflakePartner
	}
	return nil
}

type isRegisterPartnerReq_PartnerInfo interface {
	isRegisterPartnerReq_PartnerInfo()
}

type RegisterPartnerReq_DcnPartner struct {
	DcnPartner *DCNPartner `protobuf:"bytes,3,opt,name=dcn_partner,json=dcnPartner,proto3,oneof"`
}

type RegisterPartnerReq_HeadlessPartner struct {
	HeadlessPartner *HeadlessPartner `protobuf:"bytes,4,opt,name=headless_partner,json=headlessPartner,proto3,oneof"`
}

type RegisterPartnerReq_SnowflakePartner struct {
	SnowflakePartner *SnowflakePartner `protobuf:"bytes,5,opt,name=snowflake_partner,json=snowflakePartner,proto3,oneof"`
}

func (*RegisterPartnerReq_DcnPartner) isRegisterPartnerReq_PartnerInfo() {}

func (*RegisterPartnerReq_HeadlessPartner) isRegisterPartnerReq_PartnerInfo() {}

func (*RegisterPartnerReq_SnowflakePartner) isRegisterPartnerReq_PartnerInfo() {}

type PartnerInitToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier used in the JWT token auth.
	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	// Url to the API the we can register to.
	SandboxInfo string `protobuf:"bytes,2,opt,name=sandbox_info,json=sandboxInfo,proto3" json:"sandbox_info,omitempty"`
	// Secret token used to identify the partner invitation.
	Secret string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *PartnerInitToken) Reset() {
	*x = PartnerInitToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_partners_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerInitToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerInitToken) ProtoMessage() {}

func (x *PartnerInitToken) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_partners_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerInitToken.ProtoReflect.Descriptor instead.
func (*PartnerInitToken) Descriptor() ([]byte, []int) {
	return file_match_v1_partners_proto_rawDescGZIP(), []int{4}
}

func (x *PartnerInitToken) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *PartnerInitToken) GetSandboxInfo() string {
	if x != nil {
		return x.SandboxInfo
	}
	return ""
}

func (x *PartnerInitToken) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

var File_match_v1_partners_proto protoreflect.FileDescriptor

var file_match_v1_partners_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x22, 0x6e, 0x0a, 0x10, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x0a, 0x44, 0x43, 0x4e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x0f, 0x48, 0x65, 0x61, 0x64, 0x6c,
	0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xa4, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x0b, 0x64,
	0x63, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x43, 0x4e, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x63, 0x6e, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x10, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x65, 0x73, 0x73,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x6c, 0x65,
	0x73, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x68, 0x65, 0x61,
	0x64, 0x6c, 0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x11,
	0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x48, 0x00, 0x52, 0x10, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x61, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x49, 0x6e, 0x69, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_v1_partners_proto_rawDescOnce sync.Once
	file_match_v1_partners_proto_rawDescData = file_match_v1_partners_proto_rawDesc
)

func file_match_v1_partners_proto_rawDescGZIP() []byte {
	file_match_v1_partners_proto_rawDescOnce.Do(func() {
		file_match_v1_partners_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_v1_partners_proto_rawDescData)
	})
	return file_match_v1_partners_proto_rawDescData
}

var file_match_v1_partners_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_match_v1_partners_proto_goTypes = []interface{}{
	(*SnowflakePartner)(nil),   // 0: match.v1.SnowflakePartner
	(*DCNPartner)(nil),         // 1: match.v1.DCNPartner
	(*HeadlessPartner)(nil),    // 2: match.v1.HeadlessPartner
	(*RegisterPartnerReq)(nil), // 3: match.v1.RegisterPartnerReq
	(*PartnerInitToken)(nil),   // 4: match.v1.PartnerInitToken
}
var file_match_v1_partners_proto_depIdxs = []int32{
	1, // 0: match.v1.RegisterPartnerReq.dcn_partner:type_name -> match.v1.DCNPartner
	2, // 1: match.v1.RegisterPartnerReq.headless_partner:type_name -> match.v1.HeadlessPartner
	0, // 2: match.v1.RegisterPartnerReq.snowflake_partner:type_name -> match.v1.SnowflakePartner
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_match_v1_partners_proto_init() }
func file_match_v1_partners_proto_init() {
	if File_match_v1_partners_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_v1_partners_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnowflakePartner); i {
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
		file_match_v1_partners_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DCNPartner); i {
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
		file_match_v1_partners_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeadlessPartner); i {
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
		file_match_v1_partners_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPartnerReq); i {
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
		file_match_v1_partners_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartnerInitToken); i {
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
	file_match_v1_partners_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RegisterPartnerReq_DcnPartner)(nil),
		(*RegisterPartnerReq_HeadlessPartner)(nil),
		(*RegisterPartnerReq_SnowflakePartner)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_match_v1_partners_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_match_v1_partners_proto_goTypes,
		DependencyIndexes: file_match_v1_partners_proto_depIdxs,
		MessageInfos:      file_match_v1_partners_proto_msgTypes,
	}.Build()
	File_match_v1_partners_proto = out.File
	file_match_v1_partners_proto_rawDesc = nil
	file_match_v1_partners_proto_goTypes = nil
	file_match_v1_partners_proto_depIdxs = nil
}
