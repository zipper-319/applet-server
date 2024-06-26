// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: v2/applet/feedback.proto

package applet

import (
	common "applet-server/api/v2/applet/common"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type QAType int32

const (
	QAType_CommonQA QAType = 0
	QAType_UserQA   QAType = 1
)

// Enum value maps for QAType.
var (
	QAType_name = map[int32]string{
		0: "CommonQA",
		1: "UserQA",
	}
	QAType_value = map[string]int32{
		"CommonQA": 0,
		"UserQA":   1,
	}
)

func (x QAType) Enum() *QAType {
	p := new(QAType)
	*p = x
	return p
}

func (x QAType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QAType) Descriptor() protoreflect.EnumDescriptor {
	return file_v2_applet_feedback_proto_enumTypes[0].Descriptor()
}

func (QAType) Type() protoreflect.EnumType {
	return &file_v2_applet_feedback_proto_enumTypes[0]
}

func (x QAType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QAType.Descriptor instead.
func (QAType) EnumDescriptor() ([]byte, []int) {
	return file_v2_applet_feedback_proto_rawDescGZIP(), []int{0}
}

type BugType int32

const (
	BugType_unknown BugType = 0
	BugType_ASR     BugType = 1
	BugType_NLP     BugType = 2
	BugType_TTS     BugType = 3
	BugType_Other   BugType = 4
)

// Enum value maps for BugType.
var (
	BugType_name = map[int32]string{
		0: "unknown",
		1: "ASR",
		2: "NLP",
		3: "TTS",
		4: "Other",
	}
	BugType_value = map[string]int32{
		"unknown": 0,
		"ASR":     1,
		"NLP":     2,
		"TTS":     3,
		"Other":   4,
	}
)

func (x BugType) Enum() *BugType {
	p := new(BugType)
	*p = x
	return p
}

func (x BugType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BugType) Descriptor() protoreflect.EnumDescriptor {
	return file_v2_applet_feedback_proto_enumTypes[1].Descriptor()
}

func (BugType) Type() protoreflect.EnumType {
	return &file_v2_applet_feedback_proto_enumTypes[1]
}

func (x BugType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BugType.Descriptor instead.
func (BugType) EnumDescriptor() ([]byte, []int) {
	return file_v2_applet_feedback_proto_rawDescGZIP(), []int{1}
}

type CollectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId  int32          `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id"`
	Language string         `protobuf:"bytes,2,opt,name=language,proto3" json:"language"`
	Question string         `protobuf:"bytes,3,opt,name=question,proto3" json:"question"`
	Answer   string         `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer"`
	QaType   QAType         `protobuf:"varint,5,opt,name=qa_type,json=qaType,proto3,enum=applet.v2.QAType" json:"qa_type"`
	EnvType  common.EnvType `protobuf:"varint,6,opt,name=env_type,json=envType,proto3,enum=applet.v2.EnvType" json:"env_type"`
}

func (x *CollectReq) Reset() {
	*x = CollectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_feedback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectReq) ProtoMessage() {}

func (x *CollectReq) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_feedback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectReq.ProtoReflect.Descriptor instead.
func (*CollectReq) Descriptor() ([]byte, []int) {
	return file_v2_applet_feedback_proto_rawDescGZIP(), []int{0}
}

func (x *CollectReq) GetAgentId() int32 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *CollectReq) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CollectReq) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *CollectReq) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *CollectReq) GetQaType() QAType {
	if x != nil {
		return x.QaType
	}
	return QAType_CommonQA
}

func (x *CollectReq) GetEnvType() common.EnvType {
	if x != nil {
		return x.EnvType
	}
	return common.EnvType(0)
}

type CollectLikeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId    int32                    `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id"`
	Language   string                   `protobuf:"bytes,2,opt,name=language,proto3" json:"language"`
	Question   string                   `protobuf:"bytes,3,opt,name=question,proto3" json:"question"`
	Answer     string                   `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer"`
	Intent     string                   `protobuf:"bytes,5,opt,name=intent,proto3" json:"intent"`
	QuestionId string                   `protobuf:"bytes,6,opt,name=question_id,json=questionId,proto3" json:"question_id"`
	Entities   []*CollectLikeReq_Entity `protobuf:"bytes,7,rep,name=entities,proto3" json:"entities"`
	Source     string                   `protobuf:"bytes,8,opt,name=source,proto3" json:"source"`
	Domain     string                   `protobuf:"bytes,9,opt,name=domain,proto3" json:"domain"`
	EnvType    common.EnvType           `protobuf:"varint,10,opt,name=env_type,json=envType,proto3,enum=applet.v2.EnvType" json:"env_type"`
	SessionId  string                   `protobuf:"bytes,11,opt,name=session_id,json=sessionId,proto3" json:"session_id"`
}

func (x *CollectLikeReq) Reset() {
	*x = CollectLikeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_feedback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectLikeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectLikeReq) ProtoMessage() {}

func (x *CollectLikeReq) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_feedback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectLikeReq.ProtoReflect.Descriptor instead.
func (*CollectLikeReq) Descriptor() ([]byte, []int) {
	return file_v2_applet_feedback_proto_rawDescGZIP(), []int{1}
}

func (x *CollectLikeReq) GetAgentId() int32 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *CollectLikeReq) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CollectLikeReq) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *CollectLikeReq) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *CollectLikeReq) GetIntent() string {
	if x != nil {
		return x.Intent
	}
	return ""
}

func (x *CollectLikeReq) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *CollectLikeReq) GetEntities() []*CollectLikeReq_Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *CollectLikeReq) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CollectLikeReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CollectLikeReq) GetEnvType() common.EnvType {
	if x != nil {
		return x.EnvType
	}
	return common.EnvType(0)
}

func (x *CollectLikeReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type CollectDislikeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId  string         `protobuf:"bytes,1,opt,name=question_id,json=questionId,proto3" json:"question_id"`
	EnvType     common.EnvType `protobuf:"varint,2,opt,name=env_type,json=envType,proto3,enum=applet.v2.EnvType" json:"env_type"`
	AgentId     int32          `protobuf:"varint,3,opt,name=agent_id,json=agentId,proto3" json:"agent_id"`
	Language    string         `protobuf:"bytes,4,opt,name=language,proto3" json:"language"`
	Question    string         `protobuf:"bytes,5,opt,name=question,proto3" json:"question"`
	Answer      string         `protobuf:"bytes,6,opt,name=answer,proto3" json:"answer"`
	Intent      string         `protobuf:"bytes,7,opt,name=intent,proto3" json:"intent"`
	Expectation string         `protobuf:"bytes,8,opt,name=expectation,proto3" json:"expectation"`
	Reality     string         `protobuf:"bytes,9,opt,name=reality,proto3" json:"reality"`
	BugType     BugType        `protobuf:"varint,10,opt,name=bug_type,json=bugType,proto3,enum=applet.v2.BugType" json:"bug_type"`
	BugDesc     string         `protobuf:"bytes,11,opt,name=bug_desc,json=bugDesc,proto3" json:"bug_desc"`
	SessionId   string         `protobuf:"bytes,12,opt,name=session_id,json=sessionId,proto3" json:"session_id"`
}

func (x *CollectDislikeReq) Reset() {
	*x = CollectDislikeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_feedback_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectDislikeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectDislikeReq) ProtoMessage() {}

func (x *CollectDislikeReq) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_feedback_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectDislikeReq.ProtoReflect.Descriptor instead.
func (*CollectDislikeReq) Descriptor() ([]byte, []int) {
	return file_v2_applet_feedback_proto_rawDescGZIP(), []int{2}
}

func (x *CollectDislikeReq) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *CollectDislikeReq) GetEnvType() common.EnvType {
	if x != nil {
		return x.EnvType
	}
	return common.EnvType(0)
}

func (x *CollectDislikeReq) GetAgentId() int32 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *CollectDislikeReq) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CollectDislikeReq) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *CollectDislikeReq) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *CollectDislikeReq) GetIntent() string {
	if x != nil {
		return x.Intent
	}
	return ""
}

func (x *CollectDislikeReq) GetExpectation() string {
	if x != nil {
		return x.Expectation
	}
	return ""
}

func (x *CollectDislikeReq) GetReality() string {
	if x != nil {
		return x.Reality
	}
	return ""
}

func (x *CollectDislikeReq) GetBugType() BugType {
	if x != nil {
		return x.BugType
	}
	return BugType_unknown
}

func (x *CollectDislikeReq) GetBugDesc() string {
	if x != nil {
		return x.BugDesc
	}
	return ""
}

func (x *CollectDislikeReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type CollectLikeReq_Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Value       string `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	BeforeValue string `protobuf:"bytes,4,opt,name=before_value,json=beforeValue,proto3" json:"before_value"`
}

func (x *CollectLikeReq_Entity) Reset() {
	*x = CollectLikeReq_Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_applet_feedback_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectLikeReq_Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectLikeReq_Entity) ProtoMessage() {}

func (x *CollectLikeReq_Entity) ProtoReflect() protoreflect.Message {
	mi := &file_v2_applet_feedback_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectLikeReq_Entity.ProtoReflect.Descriptor instead.
func (*CollectLikeReq_Entity) Descriptor() ([]byte, []int) {
	return file_v2_applet_feedback_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CollectLikeReq_Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CollectLikeReq_Entity) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CollectLikeReq_Entity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CollectLikeReq_Entity) GetBeforeValue() string {
	if x != nil {
		return x.BeforeValue
	}
	return ""
}

var File_v2_applet_feedback_proto protoreflect.FileDescriptor

var file_v2_applet_feedback_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2f, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x70, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x76, 0x32, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x71, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x51, 0x41, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x71, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d,
	0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x76,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x54, 0x79, 0x70, 0x65, 0x22, 0xdb, 0x03,
	0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x45, 0x6e, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x1a, 0x69, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8b, 0x03, 0x0a, 0x11,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x45, 0x6e, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x2d, 0x0a, 0x08, 0x62, 0x75, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42,
	0x75, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x62, 0x75, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x75, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x75, 0x67, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x2a, 0x22, 0x0a, 0x06, 0x51, 0x41, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x51, 0x41, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x51, 0x41, 0x10, 0x01, 0x2a, 0x3c, 0x0a,
	0x07, 0x42, 0x75, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x52, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x4e, 0x4c, 0x50, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x54, 0x53, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x10, 0x04, 0x32, 0xca, 0x02, 0x0a, 0x08,
	0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x5d, 0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x6a, 0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6c,
	0x69, 0x6b, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x69,
	0x73, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x5f, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x61, 0x70, 0x70, 0x6c,
	0x65, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_applet_feedback_proto_rawDescOnce sync.Once
	file_v2_applet_feedback_proto_rawDescData = file_v2_applet_feedback_proto_rawDesc
)

func file_v2_applet_feedback_proto_rawDescGZIP() []byte {
	file_v2_applet_feedback_proto_rawDescOnce.Do(func() {
		file_v2_applet_feedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_applet_feedback_proto_rawDescData)
	})
	return file_v2_applet_feedback_proto_rawDescData
}

var file_v2_applet_feedback_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v2_applet_feedback_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v2_applet_feedback_proto_goTypes = []interface{}{
	(QAType)(0),                   // 0: applet.v2.QAType
	(BugType)(0),                  // 1: applet.v2.BugType
	(*CollectReq)(nil),            // 2: applet.v2.CollectReq
	(*CollectLikeReq)(nil),        // 3: applet.v2.CollectLikeReq
	(*CollectDislikeReq)(nil),     // 4: applet.v2.CollectDislikeReq
	(*CollectLikeReq_Entity)(nil), // 5: applet.v2.CollectLikeReq.Entity
	(common.EnvType)(0),           // 6: applet.v2.EnvType
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_v2_applet_feedback_proto_depIdxs = []int32{
	0, // 0: applet.v2.CollectReq.qa_type:type_name -> applet.v2.QAType
	6, // 1: applet.v2.CollectReq.env_type:type_name -> applet.v2.EnvType
	5, // 2: applet.v2.CollectLikeReq.entities:type_name -> applet.v2.CollectLikeReq.Entity
	6, // 3: applet.v2.CollectLikeReq.env_type:type_name -> applet.v2.EnvType
	6, // 4: applet.v2.CollectDislikeReq.env_type:type_name -> applet.v2.EnvType
	1, // 5: applet.v2.CollectDislikeReq.bug_type:type_name -> applet.v2.BugType
	2, // 6: applet.v2.Feedback.Collect:input_type -> applet.v2.CollectReq
	3, // 7: applet.v2.Feedback.CollectLike:input_type -> applet.v2.CollectLikeReq
	4, // 8: applet.v2.Feedback.CollectDislike:input_type -> applet.v2.CollectDislikeReq
	7, // 9: applet.v2.Feedback.Collect:output_type -> google.protobuf.Empty
	7, // 10: applet.v2.Feedback.CollectLike:output_type -> google.protobuf.Empty
	7, // 11: applet.v2.Feedback.CollectDislike:output_type -> google.protobuf.Empty
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v2_applet_feedback_proto_init() }
func file_v2_applet_feedback_proto_init() {
	if File_v2_applet_feedback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_applet_feedback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectReq); i {
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
		file_v2_applet_feedback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectLikeReq); i {
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
		file_v2_applet_feedback_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectDislikeReq); i {
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
		file_v2_applet_feedback_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectLikeReq_Entity); i {
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
			RawDescriptor: file_v2_applet_feedback_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v2_applet_feedback_proto_goTypes,
		DependencyIndexes: file_v2_applet_feedback_proto_depIdxs,
		EnumInfos:         file_v2_applet_feedback_proto_enumTypes,
		MessageInfos:      file_v2_applet_feedback_proto_msgTypes,
	}.Build()
	File_v2_applet_feedback_proto = out.File
	file_v2_applet_feedback_proto_rawDesc = nil
	file_v2_applet_feedback_proto_goTypes = nil
	file_v2_applet_feedback_proto_depIdxs = nil
}
