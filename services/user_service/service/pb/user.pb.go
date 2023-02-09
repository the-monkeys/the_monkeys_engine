// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: services/user_service/service/pb/user.proto

package pb

import (
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

type GetMyProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMyProfileReq) Reset() {
	*x = GetMyProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyProfileReq) ProtoMessage() {}

func (x *GetMyProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyProfileReq.ProtoReflect.Descriptor instead.
func (*GetMyProfileReq) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetMyProfileReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMyProfileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	IsActive      bool                   `protobuf:"varint,6,opt,name=isActive,proto3" json:"isActive,omitempty"`
	CountryCode   string                 `protobuf:"bytes,7,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	Mobile        string                 `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	About         string                 `protobuf:"bytes,9,opt,name=about,proto3" json:"about,omitempty"`
	Instagram     string                 `protobuf:"bytes,10,opt,name=instagram,proto3" json:"instagram,omitempty"`
	Twitter       string                 `protobuf:"bytes,11,opt,name=twitter,proto3" json:"twitter,omitempty"`
	EmailVerified bool                   `protobuf:"varint,12,opt,name=emailVerified,proto3" json:"emailVerified,omitempty"`
}

func (x *GetMyProfileRes) Reset() {
	*x = GetMyProfileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyProfileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyProfileRes) ProtoMessage() {}

func (x *GetMyProfileRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyProfileRes.ProtoReflect.Descriptor instead.
func (*GetMyProfileRes) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetMyProfileRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetMyProfileRes) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetMyProfileRes) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetMyProfileRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetMyProfileRes) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GetMyProfileRes) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *GetMyProfileRes) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GetMyProfileRes) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *GetMyProfileRes) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

func (x *GetMyProfileRes) GetInstagram() string {
	if x != nil {
		return x.Instagram
	}
	return ""
}

func (x *GetMyProfileRes) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *GetMyProfileRes) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

type SetMyProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	CountryCode string `protobuf:"bytes,3,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	MobileNo    string `protobuf:"bytes,4,opt,name=mobileNo,proto3" json:"mobileNo,omitempty"`
	About       string `protobuf:"bytes,5,opt,name=about,proto3" json:"about,omitempty"`
	Instagram   string `protobuf:"bytes,6,opt,name=instagram,proto3" json:"instagram,omitempty"`
	Twitter     string `protobuf:"bytes,7,opt,name=twitter,proto3" json:"twitter,omitempty"`
	Email       string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SetMyProfileReq) Reset() {
	*x = SetMyProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMyProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMyProfileReq) ProtoMessage() {}

func (x *SetMyProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMyProfileReq.ProtoReflect.Descriptor instead.
func (*SetMyProfileReq) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *SetMyProfileReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SetMyProfileReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SetMyProfileReq) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *SetMyProfileReq) GetMobileNo() string {
	if x != nil {
		return x.MobileNo
	}
	return ""
}

func (x *SetMyProfileReq) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

func (x *SetMyProfileReq) GetInstagram() string {
	if x != nil {
		return x.Instagram
	}
	return ""
}

func (x *SetMyProfileReq) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *SetMyProfileReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SetMyProfileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SetMyProfileRes) Reset() {
	*x = SetMyProfileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMyProfileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMyProfileRes) ProtoMessage() {}

func (x *SetMyProfileRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMyProfileRes.ProtoReflect.Descriptor instead.
func (*SetMyProfileRes) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *SetMyProfileRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SetMyProfileRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SetMyProfileRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProfilePicChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProfilePicChunk) Reset() {
	*x = ProfilePicChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfilePicChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfilePicChunk) ProtoMessage() {}

func (x *ProfilePicChunk) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfilePicChunk.ProtoReflect.Descriptor instead.
func (*ProfilePicChunk) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *ProfilePicChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProfileId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProfileId) Reset() {
	*x = ProfileId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileId) ProtoMessage() {}

func (x *ProfileId) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileId.ProtoReflect.Descriptor instead.
func (*ProfileId) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *ProfileId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_services_user_service_service_pb_user_proto protoreflect.FileDescriptor

var file_services_user_service_service_pb_user_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf7, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3a, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x62,
	0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x22, 0xed, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x62,
	0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x4f, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x82, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65,
	0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0f,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x36, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x69, 0x63, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x2e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_user_service_service_pb_user_proto_rawDescOnce sync.Once
	file_services_user_service_service_pb_user_proto_rawDescData = file_services_user_service_service_pb_user_proto_rawDesc
)

func file_services_user_service_service_pb_user_proto_rawDescGZIP() []byte {
	file_services_user_service_service_pb_user_proto_rawDescOnce.Do(func() {
		file_services_user_service_service_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_user_service_service_pb_user_proto_rawDescData)
	})
	return file_services_user_service_service_pb_user_proto_rawDescData
}

var file_services_user_service_service_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_user_service_service_pb_user_proto_goTypes = []interface{}{
	(*GetMyProfileReq)(nil),       // 0: auth.GetMyProfileReq
	(*GetMyProfileRes)(nil),       // 1: auth.GetMyProfileRes
	(*SetMyProfileReq)(nil),       // 2: auth.SetMyProfileReq
	(*SetMyProfileRes)(nil),       // 3: auth.SetMyProfileRes
	(*ProfilePicChunk)(nil),       // 4: auth.ProfilePicChunk
	(*ProfileId)(nil),             // 5: auth.ProfileId
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_services_user_service_service_pb_user_proto_depIdxs = []int32{
	6, // 0: auth.GetMyProfileRes.createTime:type_name -> google.protobuf.Timestamp
	0, // 1: auth.UserService.GetMyProfile:input_type -> auth.GetMyProfileReq
	2, // 2: auth.UserService.SetMyProfile:input_type -> auth.SetMyProfileReq
	4, // 3: auth.UserService.UploadProfile:input_type -> auth.ProfilePicChunk
	5, // 4: auth.UserService.Download:input_type -> auth.ProfileId
	1, // 5: auth.UserService.GetMyProfile:output_type -> auth.GetMyProfileRes
	3, // 6: auth.UserService.SetMyProfile:output_type -> auth.SetMyProfileRes
	5, // 7: auth.UserService.UploadProfile:output_type -> auth.ProfileId
	4, // 8: auth.UserService.Download:output_type -> auth.ProfilePicChunk
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_user_service_service_pb_user_proto_init() }
func file_services_user_service_service_pb_user_proto_init() {
	if File_services_user_service_service_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_user_service_service_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyProfileReq); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyProfileRes); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMyProfileReq); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMyProfileRes); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfilePicChunk); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileId); i {
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
			RawDescriptor: file_services_user_service_service_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_user_service_service_pb_user_proto_goTypes,
		DependencyIndexes: file_services_user_service_service_pb_user_proto_depIdxs,
		MessageInfos:      file_services_user_service_service_pb_user_proto_msgTypes,
	}.Build()
	File_services_user_service_service_pb_user_proto = out.File
	file_services_user_service_service_pb_user_proto_rawDesc = nil
	file_services_user_service_service_pb_user_proto_goTypes = nil
	file_services_user_service_service_pb_user_proto_depIdxs = nil
}
