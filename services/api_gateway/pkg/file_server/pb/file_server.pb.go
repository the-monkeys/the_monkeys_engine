// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: services/api_gateway/pkg/file_server/pb/file_server.proto

package pb

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

type UploadBlogFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId   string `protobuf:"bytes,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	FileName string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"` // int64 id = 4;
}

func (x *UploadBlogFileReq) Reset() {
	*x = UploadBlogFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBlogFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBlogFileReq) ProtoMessage() {}

func (x *UploadBlogFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBlogFileReq.ProtoReflect.Descriptor instead.
func (*UploadBlogFileReq) Descriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescGZIP(), []int{0}
}

func (x *UploadBlogFileReq) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

func (x *UploadBlogFileReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadBlogFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type UploadBlogFileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error       string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id          int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	NewFileName string `protobuf:"bytes,4,opt,name=newFileName,proto3" json:"newFileName,omitempty"`
}

func (x *UploadBlogFileRes) Reset() {
	*x = UploadBlogFileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBlogFileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBlogFileRes) ProtoMessage() {}

func (x *UploadBlogFileRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBlogFileRes.ProtoReflect.Descriptor instead.
func (*UploadBlogFileRes) Descriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescGZIP(), []int{1}
}

func (x *UploadBlogFileRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UploadBlogFileRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *UploadBlogFileRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UploadBlogFileRes) GetNewFileName() string {
	if x != nil {
		return x.NewFileName
	}
	return ""
}

type GetBlogFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId   string `protobuf:"bytes,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *GetBlogFileReq) Reset() {
	*x = GetBlogFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogFileReq) ProtoMessage() {}

func (x *GetBlogFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogFileReq.ProtoReflect.Descriptor instead.
func (*GetBlogFileReq) Descriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlogFileReq) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

func (x *GetBlogFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type GetBlogFileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetBlogFileRes) Reset() {
	*x = GetBlogFileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogFileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogFileRes) ProtoMessage() {}

func (x *GetBlogFileRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogFileRes.ProtoReflect.Descriptor instead.
func (*GetBlogFileRes) Descriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescGZIP(), []int{3}
}

func (x *GetBlogFileRes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetBlogFileRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetBlogFileRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteBlogFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId   string `protobuf:"bytes,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *DeleteBlogFileReq) Reset() {
	*x = DeleteBlogFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlogFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogFileReq) ProtoMessage() {}

func (x *DeleteBlogFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogFileReq.ProtoReflect.Descriptor instead.
func (*DeleteBlogFileReq) Descriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBlogFileReq) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

func (x *DeleteBlogFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type DeleteBlogFileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteBlogFileRes) Reset() {
	*x = DeleteBlogFileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlogFileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogFileRes) ProtoMessage() {}

func (x *DeleteBlogFileRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogFileRes.ProtoReflect.Descriptor instead.
func (*DeleteBlogFileRes) Descriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBlogFileRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteBlogFileRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteBlogFileRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_services_api_gateway_pkg_file_server_pb_file_server_proto protoreflect.FileDescriptor

var file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDesc = []byte{
	0x0a, 0x39, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x22, 0x5b, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x73,
	0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x47, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0xdd, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c,
	0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3d,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescOnce sync.Once
	file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescData = file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDesc
)

func file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescGZIP() []byte {
	file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescOnce.Do(func() {
		file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescData)
	})
	return file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDescData
}

var file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_api_gateway_pkg_file_server_pb_file_server_proto_goTypes = []interface{}{
	(*UploadBlogFileReq)(nil), // 0: auth.UploadBlogFileReq
	(*UploadBlogFileRes)(nil), // 1: auth.UploadBlogFileRes
	(*GetBlogFileReq)(nil),    // 2: auth.GetBlogFileReq
	(*GetBlogFileRes)(nil),    // 3: auth.GetBlogFileRes
	(*DeleteBlogFileReq)(nil), // 4: auth.DeleteBlogFileReq
	(*DeleteBlogFileRes)(nil), // 5: auth.DeleteBlogFileRes
}
var file_services_api_gateway_pkg_file_server_pb_file_server_proto_depIdxs = []int32{
	0, // 0: auth.UploadBlogFile.UploadBlogFile:input_type -> auth.UploadBlogFileReq
	2, // 1: auth.UploadBlogFile.GetBlogFile:input_type -> auth.GetBlogFileReq
	4, // 2: auth.UploadBlogFile.DeleteBlogFile:input_type -> auth.DeleteBlogFileReq
	1, // 3: auth.UploadBlogFile.UploadBlogFile:output_type -> auth.UploadBlogFileRes
	3, // 4: auth.UploadBlogFile.GetBlogFile:output_type -> auth.GetBlogFileRes
	5, // 5: auth.UploadBlogFile.DeleteBlogFile:output_type -> auth.DeleteBlogFileRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_api_gateway_pkg_file_server_pb_file_server_proto_init() }
func file_services_api_gateway_pkg_file_server_pb_file_server_proto_init() {
	if File_services_api_gateway_pkg_file_server_pb_file_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBlogFileReq); i {
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
		file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBlogFileRes); i {
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
		file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogFileReq); i {
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
		file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogFileRes); i {
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
		file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlogFileReq); i {
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
		file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlogFileRes); i {
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
			RawDescriptor: file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_api_gateway_pkg_file_server_pb_file_server_proto_goTypes,
		DependencyIndexes: file_services_api_gateway_pkg_file_server_pb_file_server_proto_depIdxs,
		MessageInfos:      file_services_api_gateway_pkg_file_server_pb_file_server_proto_msgTypes,
	}.Build()
	File_services_api_gateway_pkg_file_server_pb_file_server_proto = out.File
	file_services_api_gateway_pkg_file_server_pb_file_server_proto_rawDesc = nil
	file_services_api_gateway_pkg_file_server_pb_file_server_proto_goTypes = nil
	file_services_api_gateway_pkg_file_server_pb_file_server_proto_depIdxs = nil
}
