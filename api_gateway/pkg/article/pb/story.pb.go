// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: api_gateway/pkg/article/pb/story.proto

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

// CreateArticleRequest
type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title      string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content    string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Author     string                 `protobuf:"bytes,4,opt,name=Author,proto3" json:"Author,omitempty"`
	IsDraft    bool                   `protobuf:"varint,5,opt,name=isDraft,proto3" json:"isDraft,omitempty"`
	Tags       []string               `protobuf:"bytes,6,rep,name=Tags,proto3" json:"Tags,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	QuickRead  bool                   `protobuf:"varint,9,opt,name=quickRead,proto3" json:"quickRead,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_pkg_article_pb_story_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArticleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticleRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateArticleRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateArticleRequest) GetIsDraft() bool {
	if x != nil {
		return x.IsDraft
	}
	return false
}

func (x *CreateArticleRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateArticleRequest) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CreateArticleRequest) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *CreateArticleRequest) GetQuickRead() bool {
	if x != nil {
		return x.QuickRead
	}
	return false
}

type CreateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateArticleResponse) Reset() {
	*x = CreateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleResponse) ProtoMessage() {}

func (x *CreateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleResponse.ProtoReflect.Descriptor instead.
func (*CreateArticleResponse) Descriptor() ([]byte, []int) {
	return file_api_gateway_pkg_article_pb_story_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArticleResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateArticleResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateArticleResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetArticlesRequest) Reset() {
	*x = GetArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesRequest) ProtoMessage() {}

func (x *GetArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesRequest.ProtoReflect.Descriptor instead.
func (*GetArticlesRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_pkg_article_pb_story_proto_rawDescGZIP(), []int{2}
}

type GetArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title      string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author     string                 `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	QuickRead  bool                   `protobuf:"varint,5,opt,name=quickRead,proto3" json:"quickRead,omitempty"`
	ViewBy     []string               `protobuf:"bytes,6,rep,name=viewBy,proto3" json:"viewBy,omitempty"`
}

func (x *GetArticlesResponse) Reset() {
	*x = GetArticlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesResponse) ProtoMessage() {}

func (x *GetArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesResponse.ProtoReflect.Descriptor instead.
func (*GetArticlesResponse) Descriptor() ([]byte, []int) {
	return file_api_gateway_pkg_article_pb_story_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticlesResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetArticlesResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetArticlesResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetArticlesResponse) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GetArticlesResponse) GetQuickRead() bool {
	if x != nil {
		return x.QuickRead
	}
	return false
}

func (x *GetArticlesResponse) GetViewBy() []string {
	if x != nil {
		return x.ViewBy
	}
	return nil
}

type GetArticleByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetArticleByIdReq) Reset() {
	*x = GetArticleByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleByIdReq) ProtoMessage() {}

func (x *GetArticleByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleByIdReq.ProtoReflect.Descriptor instead.
func (*GetArticleByIdReq) Descriptor() ([]byte, []int) {
	return file_api_gateway_pkg_article_pb_story_proto_rawDescGZIP(), []int{4}
}

func (x *GetArticleByIdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetArticleByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title      string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author     string                 `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Content    string                 `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	QuickRead  bool                   `protobuf:"varint,7,opt,name=quickRead,proto3" json:"quickRead,omitempty"`
	NoOfViews  int64                  `protobuf:"varint,8,opt,name=noOfViews,proto3" json:"noOfViews,omitempty"`
}

func (x *GetArticleByIdResp) Reset() {
	*x = GetArticleByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleByIdResp) ProtoMessage() {}

func (x *GetArticleByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_pkg_article_pb_story_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleByIdResp.ProtoReflect.Descriptor instead.
func (*GetArticleByIdResp) Descriptor() ([]byte, []int) {
	return file_api_gateway_pkg_article_pb_story_proto_rawDescGZIP(), []int{5}
}

func (x *GetArticleByIdResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetArticleByIdResp) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetArticleByIdResp) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetArticleByIdResp) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetArticleByIdResp) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GetArticleByIdResp) GetQuickRead() bool {
	if x != nil {
		return x.QuickRead
	}
	return false
}

func (x *GetArticleByIdResp) GetNoOfViews() int64 {
	if x != nil {
		return x.NoOfViews
	}
	return 0
}

var File_api_gateway_pkg_article_pb_story_proto protoreflect.FileDescriptor

var file_api_gateway_pkg_article_pb_story_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb2, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x3a,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x52,
	0x65, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x71, 0x75, 0x69, 0x63, 0x6b,
	0x52, 0x65, 0x61, 0x64, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xc5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x65, 0x77, 0x42, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x42, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0xe4,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x69,
	0x63, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x71, 0x75,
	0x69, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f, 0x4f, 0x66, 0x56,
	0x69, 0x65, 0x77, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x6f, 0x4f, 0x66,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x32, 0xeb, 0x01, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_gateway_pkg_article_pb_story_proto_rawDescOnce sync.Once
	file_api_gateway_pkg_article_pb_story_proto_rawDescData = file_api_gateway_pkg_article_pb_story_proto_rawDesc
)

func file_api_gateway_pkg_article_pb_story_proto_rawDescGZIP() []byte {
	file_api_gateway_pkg_article_pb_story_proto_rawDescOnce.Do(func() {
		file_api_gateway_pkg_article_pb_story_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gateway_pkg_article_pb_story_proto_rawDescData)
	})
	return file_api_gateway_pkg_article_pb_story_proto_rawDescData
}

var file_api_gateway_pkg_article_pb_story_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_gateway_pkg_article_pb_story_proto_goTypes = []interface{}{
	(*CreateArticleRequest)(nil),  // 0: auth.CreateArticleRequest
	(*CreateArticleResponse)(nil), // 1: auth.CreateArticleResponse
	(*GetArticlesRequest)(nil),    // 2: auth.GetArticlesRequest
	(*GetArticlesResponse)(nil),   // 3: auth.GetArticlesResponse
	(*GetArticleByIdReq)(nil),     // 4: auth.GetArticleByIdReq
	(*GetArticleByIdResp)(nil),    // 5: auth.GetArticleByIdResp
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_api_gateway_pkg_article_pb_story_proto_depIdxs = []int32{
	6, // 0: auth.CreateArticleRequest.createTime:type_name -> google.protobuf.Timestamp
	6, // 1: auth.CreateArticleRequest.updateTime:type_name -> google.protobuf.Timestamp
	6, // 2: auth.GetArticlesResponse.createTime:type_name -> google.protobuf.Timestamp
	6, // 3: auth.GetArticleByIdResp.createTime:type_name -> google.protobuf.Timestamp
	0, // 4: auth.ArticleService.CreateArticle:input_type -> auth.CreateArticleRequest
	2, // 5: auth.ArticleService.GetArticles:input_type -> auth.GetArticlesRequest
	4, // 6: auth.ArticleService.GetArticleById:input_type -> auth.GetArticleByIdReq
	1, // 7: auth.ArticleService.CreateArticle:output_type -> auth.CreateArticleResponse
	3, // 8: auth.ArticleService.GetArticles:output_type -> auth.GetArticlesResponse
	5, // 9: auth.ArticleService.GetArticleById:output_type -> auth.GetArticleByIdResp
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_gateway_pkg_article_pb_story_proto_init() }
func file_api_gateway_pkg_article_pb_story_proto_init() {
	if File_api_gateway_pkg_article_pb_story_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gateway_pkg_article_pb_story_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_api_gateway_pkg_article_pb_story_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleResponse); i {
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
		file_api_gateway_pkg_article_pb_story_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticlesRequest); i {
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
		file_api_gateway_pkg_article_pb_story_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticlesResponse); i {
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
		file_api_gateway_pkg_article_pb_story_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleByIdReq); i {
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
		file_api_gateway_pkg_article_pb_story_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleByIdResp); i {
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
			RawDescriptor: file_api_gateway_pkg_article_pb_story_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_gateway_pkg_article_pb_story_proto_goTypes,
		DependencyIndexes: file_api_gateway_pkg_article_pb_story_proto_depIdxs,
		MessageInfos:      file_api_gateway_pkg_article_pb_story_proto_msgTypes,
	}.Build()
	File_api_gateway_pkg_article_pb_story_proto = out.File
	file_api_gateway_pkg_article_pb_story_proto_rawDesc = nil
	file_api_gateway_pkg_article_pb_story_proto_goTypes = nil
	file_api_gateway_pkg_article_pb_story_proto_depIdxs = nil
}
