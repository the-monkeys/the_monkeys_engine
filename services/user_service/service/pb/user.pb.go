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
	Id          int64  `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *SetMyProfileReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SetMyProfileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *SetMyProfileRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UploadProfilePicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Id   int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UploadProfilePicReq) Reset() {
	*x = UploadProfilePicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadProfilePicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadProfilePicReq) ProtoMessage() {}

func (x *UploadProfilePicReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UploadProfilePicReq.ProtoReflect.Descriptor instead.
func (*UploadProfilePicReq) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *UploadProfilePicReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadProfilePicReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UploadProfilePicRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UploadProfilePicRes) Reset() {
	*x = UploadProfilePicRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadProfilePicRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadProfilePicRes) ProtoMessage() {}

func (x *UploadProfilePicRes) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UploadProfilePicRes.ProtoReflect.Descriptor instead.
func (*UploadProfilePicRes) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *UploadProfilePicRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UploadProfilePicRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *UploadProfilePicRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProfilePicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProfilePicReq) Reset() {
	*x = GetProfilePicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfilePicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfilePicReq) ProtoMessage() {}

func (x *GetProfilePicReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfilePicReq.ProtoReflect.Descriptor instead.
func (*GetProfilePicReq) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetProfilePicReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProfilePicRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetProfilePicRes) Reset() {
	*x = GetProfilePicRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfilePicRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfilePicRes) ProtoMessage() {}

func (x *GetProfilePicRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfilePicRes.ProtoReflect.Descriptor instead.
func (*GetProfilePicRes) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetProfilePicRes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteMyAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMyAccountReq) Reset() {
	*x = DeleteMyAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMyAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMyAccountReq) ProtoMessage() {}

func (x *DeleteMyAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMyAccountReq.ProtoReflect.Descriptor instead.
func (*DeleteMyAccountReq) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMyAccountReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteMyAccountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMyAccountRes) Reset() {
	*x = DeleteMyAccountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_service_service_pb_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMyAccountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMyAccountRes) ProtoMessage() {}

func (x *DeleteMyAccountRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_service_service_pb_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMyAccountRes.ProtoReflect.Descriptor instead.
func (*DeleteMyAccountRes) Descriptor() ([]byte, []int) {
	return file_services_user_service_service_pb_user_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMyAccountRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteMyAccountRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *DeleteMyAccountRes) GetId() int64 {
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
	0x64, 0x22, 0xfd, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69,
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
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4f, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a,
	0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xe1, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x4d,
	0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x53, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x79, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22,
	0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_services_user_service_service_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_user_service_service_pb_user_proto_goTypes = []interface{}{
	(*GetMyProfileReq)(nil),       // 0: auth.GetMyProfileReq
	(*GetMyProfileRes)(nil),       // 1: auth.GetMyProfileRes
	(*SetMyProfileReq)(nil),       // 2: auth.SetMyProfileReq
	(*SetMyProfileRes)(nil),       // 3: auth.SetMyProfileRes
	(*UploadProfilePicReq)(nil),   // 4: auth.UploadProfilePicReq
	(*UploadProfilePicRes)(nil),   // 5: auth.UploadProfilePicRes
	(*GetProfilePicReq)(nil),      // 6: auth.GetProfilePicReq
	(*GetProfilePicRes)(nil),      // 7: auth.GetProfilePicRes
	(*DeleteMyAccountReq)(nil),    // 8: auth.DeleteMyAccountReq
	(*DeleteMyAccountRes)(nil),    // 9: auth.DeleteMyAccountRes
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_services_user_service_service_pb_user_proto_depIdxs = []int32{
	10, // 0: auth.GetMyProfileRes.createTime:type_name -> google.protobuf.Timestamp
	0,  // 1: auth.UserService.GetMyProfile:input_type -> auth.GetMyProfileReq
	2,  // 2: auth.UserService.SetMyProfile:input_type -> auth.SetMyProfileReq
	4,  // 3: auth.UserService.UploadProfile:input_type -> auth.UploadProfilePicReq
	6,  // 4: auth.UserService.Download:input_type -> auth.GetProfilePicReq
	8,  // 5: auth.UserService.DeleteMyProfile:input_type -> auth.DeleteMyAccountReq
	1,  // 6: auth.UserService.GetMyProfile:output_type -> auth.GetMyProfileRes
	3,  // 7: auth.UserService.SetMyProfile:output_type -> auth.SetMyProfileRes
	5,  // 8: auth.UserService.UploadProfile:output_type -> auth.UploadProfilePicRes
	7,  // 9: auth.UserService.Download:output_type -> auth.GetProfilePicRes
	9,  // 10: auth.UserService.DeleteMyProfile:output_type -> auth.DeleteMyAccountRes
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
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
			switch v := v.(*UploadProfilePicReq); i {
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
			switch v := v.(*UploadProfilePicRes); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfilePicReq); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfilePicRes); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMyAccountReq); i {
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
		file_services_user_service_service_pb_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMyAccountRes); i {
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
			NumMessages:   10,
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
