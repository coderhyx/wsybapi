// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.2
// source: video.proto

package proto

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

type RequestChannelAdvert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (x *RequestChannelAdvert) Reset() {
	*x = RequestChannelAdvert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChannelAdvert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChannelAdvert) ProtoMessage() {}

func (x *RequestChannelAdvert) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChannelAdvert.ProtoReflect.Descriptor instead.
func (*RequestChannelAdvert) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *RequestChannelAdvert) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type ResponseChannelAdvert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string               `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Items []*ChannelAdvertData `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Count int64                `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ResponseChannelAdvert) Reset() {
	*x = ResponseChannelAdvert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChannelAdvert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChannelAdvert) ProtoMessage() {}

func (x *ResponseChannelAdvert) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseChannelAdvert.ProtoReflect.Descriptor instead.
func (*ResponseChannelAdvert) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseChannelAdvert) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseChannelAdvert) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ResponseChannelAdvert) GetItems() []*ChannelAdvertData {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ResponseChannelAdvert) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ChannelAdvertData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle string `protobuf:"bytes,3,opt,name=subTitle,proto3" json:"subTitle,omitempty"`
	AddTime  int64  `protobuf:"varint,4,opt,name=addTime,proto3" json:"addTime,omitempty"`
	Img      string `protobuf:"bytes,5,opt,name=img,proto3" json:"img,omitempty"`
	Url      string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ChannelAdvertData) Reset() {
	*x = ChannelAdvertData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelAdvertData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelAdvertData) ProtoMessage() {}

func (x *ChannelAdvertData) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelAdvertData.ProtoReflect.Descriptor instead.
func (*ChannelAdvertData) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelAdvertData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChannelAdvertData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChannelAdvertData) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *ChannelAdvertData) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *ChannelAdvertData) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *ChannelAdvertData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type RequestChannelHotList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (x *RequestChannelHotList) Reset() {
	*x = RequestChannelHotList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChannelHotList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChannelHotList) ProtoMessage() {}

func (x *RequestChannelHotList) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChannelHotList.ProtoReflect.Descriptor instead.
func (*RequestChannelHotList) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *RequestChannelHotList) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type ResponseChannelHotLists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64                     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string                    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Items []*ResponseChannelHotList `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Count int64                     `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ResponseChannelHotLists) Reset() {
	*x = ResponseChannelHotLists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChannelHotLists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChannelHotLists) ProtoMessage() {}

func (x *ResponseChannelHotLists) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseChannelHotLists.ProtoReflect.Descriptor instead.
func (*ResponseChannelHotLists) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseChannelHotLists) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseChannelHotLists) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ResponseChannelHotLists) GetItems() []*ResponseChannelHotList {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ResponseChannelHotLists) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ResponseChannelHotList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle      string `protobuf:"bytes,3,opt,name=subTitle,proto3" json:"subTitle,omitempty"`
	AddTime       int64  `protobuf:"varint,4,opt,name=addTime,proto3" json:"addTime,omitempty"`
	Img           string `protobuf:"bytes,5,opt,name=img,proto3" json:"img,omitempty"`
	Img1          string `protobuf:"bytes,6,opt,name=img1,proto3" json:"img1,omitempty"`
	EpisodesCount int64  `protobuf:"varint,7,opt,name=episodesCount,proto3" json:"episodesCount,omitempty"`
	IsEnd         int64  `protobuf:"varint,8,opt,name=isEnd,proto3" json:"isEnd,omitempty"`
	Comment       int64  `protobuf:"varint,9,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *ResponseChannelHotList) Reset() {
	*x = ResponseChannelHotList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseChannelHotList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChannelHotList) ProtoMessage() {}

func (x *ResponseChannelHotList) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseChannelHotList.ProtoReflect.Descriptor instead.
func (*ResponseChannelHotList) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseChannelHotList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResponseChannelHotList) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResponseChannelHotList) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *ResponseChannelHotList) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *ResponseChannelHotList) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *ResponseChannelHotList) GetImg1() string {
	if x != nil {
		return x.Img1
	}
	return ""
}

func (x *ResponseChannelHotList) GetEpisodesCount() int64 {
	if x != nil {
		return x.EpisodesCount
	}
	return 0
}

func (x *ResponseChannelHotList) GetIsEnd() int64 {
	if x != nil {
		return x.IsEnd
	}
	return 0
}

func (x *ResponseChannelHotList) GetComment() int64 {
	if x != nil {
		return x.Comment
	}
	return 0
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x15, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x93, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x6d, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x35, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x8a, 0x01,
	0x0a, 0x17, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf0, 0x01, 0x0a, 0x16, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x6f,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x6d, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x67, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x6d, 0x67, 0x31, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x73, 0x45, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x73,
	0x45, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xae, 0x01,
	0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x00, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_video_proto_goTypes = []interface{}{
	(*RequestChannelAdvert)(nil),    // 0: proto.RequestChannelAdvert
	(*ResponseChannelAdvert)(nil),   // 1: proto.ResponseChannelAdvert
	(*ChannelAdvertData)(nil),       // 2: proto.ChannelAdvertData
	(*RequestChannelHotList)(nil),   // 3: proto.RequestChannelHotList
	(*ResponseChannelHotLists)(nil), // 4: proto.ResponseChannelHotLists
	(*ResponseChannelHotList)(nil),  // 5: proto.ResponseChannelHotList
}
var file_video_proto_depIdxs = []int32{
	2, // 0: proto.ResponseChannelAdvert.items:type_name -> proto.ChannelAdvertData
	5, // 1: proto.ResponseChannelHotLists.items:type_name -> proto.ResponseChannelHotList
	0, // 2: proto.VideoService.ChannelAdvert:input_type -> proto.RequestChannelAdvert
	3, // 3: proto.VideoService.ChannelHotList:input_type -> proto.RequestChannelHotList
	1, // 4: proto.VideoService.ChannelAdvert:output_type -> proto.ResponseChannelAdvert
	4, // 5: proto.VideoService.ChannelHotList:output_type -> proto.ResponseChannelHotLists
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChannelAdvert); i {
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
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseChannelAdvert); i {
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
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelAdvertData); i {
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
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChannelHotList); i {
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
		file_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseChannelHotLists); i {
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
		file_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseChannelHotList); i {
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
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}
