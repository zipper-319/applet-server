// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: v2/applet/voiceData.proto

package applet

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VoiceType int32

const (
	VoiceType_Primary VoiceType = 0
	VoiceType_Middle  VoiceType = 1
	VoiceType_Senior  VoiceType = 2
	VoiceType_Custom  VoiceType = 3
)

// Enum value maps for VoiceType.
var (
	VoiceType_name = map[int32]string{
		0: "Primary",
		1: "Middle",
		2: "Senior",
		3: "Custom",
	}
	VoiceType_value = map[string]int32{
		"Primary": 0,
		"Middle":  1,
		"Senior":  2,
		"Custom":  3,
	}
)

func (x VoiceType) Enum() *VoiceType {
	p := new(VoiceType)
	*p = x
	return p
}

func (x VoiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_v2_applet_voiceData_proto_enumTypes[0].Descriptor()
}

func (VoiceType) Type() protoreflect.EnumType {
	return &file_v2_applet_voiceData_proto_enumTypes[0]
}

func (x VoiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VoiceType.Descriptor instead.
func (VoiceType) EnumDescriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{0}
}

type VoiceDataReqData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceType VoiceType `protobuf:"varint,1,opt,name=voice_type,json=voiceType,proto3,enum=applet.v2.VoiceType" json:"voice_type,omitempty"`
	// 音频,base64编码
	Voice string `protobuf:"bytes,2,opt,name=voice,proto3" json:"voice,omitempty"`
	// 序号
	Sequence int32 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *VoiceDataReqData) Reset() {
	*x = VoiceDataReqData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceDataReqData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceDataReqData) ProtoMessage() {}

func (x *VoiceDataReqData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceDataReqData.ProtoReflect.Descriptor instead.
func (*VoiceDataReqData) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{0}
}

func (x *VoiceDataReqData) GetVoiceType() VoiceType {
	if x != nil {
		return x.VoiceType
	}
	return VoiceType_Primary
}

func (x *VoiceDataReqData) GetVoice() string {
	if x != nil {
		return x.Voice
	}
	return ""
}

func (x *VoiceDataReqData) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type VoiceDataResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 下一条录制音频序号
	NextSequence int32 `protobuf:"varint,3,opt,name=next_sequence,json=nextSequence,proto3" json:"next_sequence,omitempty"`
}

func (x *VoiceDataResData) Reset() {
	*x = VoiceDataResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceDataResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceDataResData) ProtoMessage() {}

func (x *VoiceDataResData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceDataResData.ProtoReflect.Descriptor instead.
func (*VoiceDataResData) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{1}
}

func (x *VoiceDataResData) GetNextSequence() int32 {
	if x != nil {
		return x.NextSequence
	}
	return 0
}

type ProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceType VoiceType `protobuf:"varint,1,opt,name=voice_type,json=voiceType,proto3,enum=applet.v2.VoiceType" json:"voice_type,omitempty"`
}

func (x *ProgressRequest) Reset() {
	*x = ProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressRequest) ProtoMessage() {}

func (x *ProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressRequest.ProtoReflect.Descriptor instead.
func (*ProgressRequest) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{2}
}

func (x *ProgressRequest) GetVoiceType() VoiceType {
	if x != nil {
		return x.VoiceType
	}
	return VoiceType_Primary
}

type ProgressResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 已录制的数量
	CurrentNumber int32 `protobuf:"varint,2,opt,name=current_number,json=currentNumber,proto3" json:"current_number,omitempty"`
	// 时间戳
	FinishedTime int64 `protobuf:"varint,3,opt,name=finished_time,json=finishedTime,proto3" json:"finished_time,omitempty"`
}

func (x *ProgressResData) Reset() {
	*x = ProgressResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressResData) ProtoMessage() {}

func (x *ProgressResData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressResData.ProtoReflect.Descriptor instead.
func (*ProgressResData) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{3}
}

func (x *ProgressResData) GetCurrentNumber() int32 {
	if x != nil {
		return x.CurrentNumber
	}
	return 0
}

func (x *ProgressResData) GetFinishedTime() int64 {
	if x != nil {
		return x.FinishedTime
	}
	return 0
}

type DownloadReqData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 序号
	Sequence int32 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *DownloadReqData) Reset() {
	*x = DownloadReqData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadReqData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadReqData) ProtoMessage() {}

func (x *DownloadReqData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadReqData.ProtoReflect.Descriptor instead.
func (*DownloadReqData) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadReqData) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type DownloadResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 音频数据，base64编码
	VoiceData string `protobuf:"bytes,2,opt,name=voiceData,proto3" json:"voiceData,omitempty"`
}

func (x *DownloadResData) Reset() {
	*x = DownloadResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResData) ProtoMessage() {}

func (x *DownloadResData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResData.ProtoReflect.Descriptor instead.
func (*DownloadResData) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadResData) GetVoiceData() string {
	if x != nil {
		return x.VoiceData
	}
	return ""
}

type CommitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceType VoiceType `protobuf:"varint,1,opt,name=voice_type,json=voiceType,proto3,enum=applet.v2.VoiceType" json:"voice_type,omitempty"`
	Speaker   string    `protobuf:"bytes,2,opt,name=speaker,proto3" json:"speaker,omitempty"` // 发音人
}

func (x *CommitRequest) Reset() {
	*x = CommitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitRequest) ProtoMessage() {}

func (x *CommitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitRequest.ProtoReflect.Descriptor instead.
func (*CommitRequest) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{6}
}

func (x *CommitRequest) GetVoiceType() VoiceType {
	if x != nil {
		return x.VoiceType
	}
	return VoiceType_Primary
}

func (x *CommitRequest) GetSpeaker() string {
	if x != nil {
		return x.Speaker
	}
	return ""
}

type CommitResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 完成时间点  时间格式"2006-01-02"
	FinishedTime string `protobuf:"bytes,2,opt,name=finished_time,json=finishedTime,proto3" json:"finished_time,omitempty"`
	// 需要等待训练的时间  单位小时
	AwaitTrain int32 `protobuf:"varint,3,opt,name=await_train,json=awaitTrain,proto3" json:"await_train,omitempty"`
}

func (x *CommitResData) Reset() {
	*x = CommitResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitResData) ProtoMessage() {}

func (x *CommitResData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitResData.ProtoReflect.Descriptor instead.
func (*CommitResData) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{7}
}

func (x *CommitResData) GetFinishedTime() string {
	if x != nil {
		return x.FinishedTime
	}
	return ""
}

func (x *CommitResData) GetAwaitTrain() int32 {
	if x != nil {
		return x.AwaitTrain
	}
	return 0
}

type GetTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceType VoiceType `protobuf:"varint,1,opt,name=voice_type,json=voiceType,proto3,enum=applet.v2.VoiceType" json:"voice_type,omitempty"`
}

func (x *GetTextRequest) Reset() {
	*x = GetTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextRequest) ProtoMessage() {}

func (x *GetTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextRequest.ProtoReflect.Descriptor instead.
func (*GetTextRequest) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{8}
}

func (x *GetTextRequest) GetVoiceType() VoiceType {
	if x != nil {
		return x.VoiceType
	}
	return VoiceType_Primary
}

type GetTextResData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文本数据
	Text []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
}

func (x *GetTextResData) Reset() {
	*x = GetTextResData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextResData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextResData) ProtoMessage() {}

func (x *GetTextResData) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextResData.ProtoReflect.Descriptor instead.
func (*GetTextResData) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{9}
}

func (x *GetTextResData) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

type UploadFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speaker string `protobuf:"bytes,1,opt,name=speaker,proto3" json:"speaker,omitempty"`
	File    string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UploadFilesRequest) Reset() {
	*x = UploadFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_voiceData_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFilesRequest) ProtoMessage() {}

func (x *UploadFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_voiceData_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFilesRequest.ProtoReflect.Descriptor instead.
func (*UploadFilesRequest) Descriptor() ([]byte, []int) {
	return file_v2_applet_voiceData_proto_rawDescGZIP(), []int{10}
}

func (x *UploadFilesRequest) GetSpeaker() string {
	if x != nil {
		return x.Speaker
	}
	return ""
}

func (x *UploadFilesRequest) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

var File_v2_applet_voiceData_proto protoreflect.FileDescriptor

var file_v2_applet_voiceData_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x70,
	0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79,
	0x0a, 0x10, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x33, 0x0a, 0x0a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x37, 0x0a, 0x10, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x22, 0x46, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5d, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x0f, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2f, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5e, 0x0a, 0x0d, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x0d, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x77, 0x61, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x22, 0x45, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x42, 0x0a,
	0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x2a, 0x3c, 0x0a, 0x09, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x69, 0x6f,
	0x72, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x10, 0x03, 0x32,
	0x9a, 0x05, 0x0a, 0x12, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x93, 0x01, 0x0a, 0x0c, 0x70, 0x75, 0x74, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x49, 0x92, 0x41, 0x1c, 0x12, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x12,
	0xe4, 0xb8, 0x8a, 0xe4, 0xbc, 0xa0, 0xe9, 0x9f, 0xb3, 0xe9, 0xa2, 0x91, 0xe6, 0x95, 0xb0, 0xe6,
	0x8d, 0xae, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x6e, 0x0a, 0x0b,
	0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x67, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x6c, 0x0a, 0x0d,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x2e,
	0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x62, 0x0a, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x64,
	0x0a, 0x07, 0x67, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x78, 0x74,
	0x2f, 0x67, 0x65, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b,
	0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v2_applet_voiceData_proto_rawDescOnce sync.Once
	file_v2_applet_voiceData_proto_rawDescData = file_v2_applet_voiceData_proto_rawDesc
)

func file_v2_applet_voiceData_proto_rawDescGZIP() []byte {
	file_v2_applet_voiceData_proto_rawDescOnce.Do(func() {
		file_v2_applet_voiceData_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_applet_voiceData_proto_rawDescData)
	})
	return file_v2_applet_voiceData_proto_rawDescData
}

var file_v2_applet_voiceData_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v2_applet_voiceData_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_v2_applet_voiceData_proto_goTypes = []interface{}{
	(VoiceType)(0),             // 0: applet.v2.VoiceType
	(*VoiceDataReqData)(nil),   // 1: applet.v2.VoiceDataReqData
	(*VoiceDataResData)(nil),   // 2: applet.v2.VoiceDataResData
	(*ProgressRequest)(nil),    // 3: applet.v2.ProgressRequest
	(*ProgressResData)(nil),    // 4: applet.v2.ProgressResData
	(*DownloadReqData)(nil),    // 5: applet.v2.DownloadReqData
	(*DownloadResData)(nil),    // 6: applet.v2.DownloadResData
	(*CommitRequest)(nil),      // 7: applet.v2.CommitRequest
	(*CommitResData)(nil),      // 8: applet.v2.CommitResData
	(*GetTextRequest)(nil),     // 9: applet.v2.GetTextRequest
	(*GetTextResData)(nil),     // 10: applet.v2.GetTextResData
	(*UploadFilesRequest)(nil), // 11: applet.v2.UploadFilesRequest
	(*emptypb.Empty)(nil),      // 12: google.protobuf.Empty
}
var file_v2_applet_voiceData_proto_depIdxs = []int32{
	0,  // 0: applet.v2.VoiceDataReqData.voice_type:type_name -> applet.v2.VoiceType
	0,  // 1: applet.v2.ProgressRequest.voice_type:type_name -> applet.v2.VoiceType
	0,  // 2: applet.v2.CommitRequest.voice_type:type_name -> applet.v2.VoiceType
	0,  // 3: applet.v2.GetTextRequest.voice_type:type_name -> applet.v2.VoiceType
	1,  // 4: applet.v2.VoiceDataOperation.putVoiceData:input_type -> applet.v2.VoiceDataReqData
	3,  // 5: applet.v2.VoiceDataOperation.getProgress:input_type -> applet.v2.ProgressRequest
	5,  // 6: applet.v2.VoiceDataOperation.downloadVoice:input_type -> applet.v2.DownloadReqData
	7,  // 7: applet.v2.VoiceDataOperation.commit:input_type -> applet.v2.CommitRequest
	9,  // 8: applet.v2.VoiceDataOperation.getText:input_type -> applet.v2.GetTextRequest
	11, // 9: applet.v2.VoiceDataOperation.UploadFiles:input_type -> applet.v2.UploadFilesRequest
	2,  // 10: applet.v2.VoiceDataOperation.putVoiceData:output_type -> applet.v2.VoiceDataResData
	4,  // 11: applet.v2.VoiceDataOperation.getProgress:output_type -> applet.v2.ProgressResData
	6,  // 12: applet.v2.VoiceDataOperation.downloadVoice:output_type -> applet.v2.DownloadResData
	8,  // 13: applet.v2.VoiceDataOperation.commit:output_type -> applet.v2.CommitResData
	10, // 14: applet.v2.VoiceDataOperation.getText:output_type -> applet.v2.GetTextResData
	12, // 15: applet.v2.VoiceDataOperation.UploadFiles:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_v2_applet_voiceData_proto_init() }
func file_v2_applet_voiceData_proto_init() {
	if File_v2_applet_voiceData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_applet_voiceData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceDataReqData); i {
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
		file_v2_applet_voiceData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceDataResData); i {
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
		file_v2_applet_voiceData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressRequest); i {
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
		file_v2_applet_voiceData_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressResData); i {
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
		file_v2_applet_voiceData_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadReqData); i {
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
		file_v2_applet_voiceData_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadResData); i {
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
		file_v2_applet_voiceData_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitRequest); i {
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
		file_v2_applet_voiceData_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitResData); i {
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
		file_v2_applet_voiceData_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextRequest); i {
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
		file_v2_applet_voiceData_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextResData); i {
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
		file_v2_applet_voiceData_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFilesRequest); i {
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
			RawDescriptor: file_v2_applet_voiceData_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v2_applet_voiceData_proto_goTypes,
		DependencyIndexes: file_v2_applet_voiceData_proto_depIdxs,
		EnumInfos:         file_v2_applet_voiceData_proto_enumTypes,
		MessageInfos:      file_v2_applet_voiceData_proto_msgTypes,
	}.Build()
	File_v2_applet_voiceData_proto = out.File
	file_v2_applet_voiceData_proto_rawDesc = nil
	file_v2_applet_voiceData_proto_goTypes = nil
	file_v2_applet_voiceData_proto_depIdxs = nil
}
