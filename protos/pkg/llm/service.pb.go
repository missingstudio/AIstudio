// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: llm/service.proto

package llmv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FinishReason int32

const (
	FinishReason_NULL   FinishReason = 0
	FinishReason_LENGTH FinishReason = 1
	FinishReason_STOP   FinishReason = 2
	FinishReason_ERROR  FinishReason = 3
)

// Enum value maps for FinishReason.
var (
	FinishReason_name = map[int32]string{
		0: "NULL",
		1: "LENGTH",
		2: "STOP",
		3: "ERROR",
	}
	FinishReason_value = map[string]int32{
		"NULL":   0,
		"LENGTH": 1,
		"STOP":   2,
		"ERROR":  3,
	}
)

func (x FinishReason) Enum() *FinishReason {
	p := new(FinishReason)
	*p = x
	return p
}

func (x FinishReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FinishReason) Descriptor() protoreflect.EnumDescriptor {
	return file_llm_service_proto_enumTypes[0].Descriptor()
}

func (FinishReason) Type() protoreflect.EnumType {
	return &file_llm_service_proto_enumTypes[0]
}

func (x FinishReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FinishReason.Descriptor instead.
func (FinishReason) EnumDescriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{0}
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Role:
	//
	//	*Role_System
	//	*Role_User
	//	*Role_Assistant
	Role isRole_Role `protobuf_oneof:"role"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{0}
}

func (m *Role) GetRole() isRole_Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (x *Role) GetSystem() string {
	if x, ok := x.GetRole().(*Role_System); ok {
		return x.System
	}
	return ""
}

func (x *Role) GetUser() string {
	if x, ok := x.GetRole().(*Role_User); ok {
		return x.User
	}
	return ""
}

func (x *Role) GetAssistant() string {
	if x, ok := x.GetRole().(*Role_Assistant); ok {
		return x.Assistant
	}
	return ""
}

type isRole_Role interface {
	isRole_Role()
}

type Role_System struct {
	System string `protobuf:"bytes,1,opt,name=system,proto3,oneof"`
}

type Role_User struct {
	User string `protobuf:"bytes,2,opt,name=user,proto3,oneof"`
}

type Role_Assistant struct {
	Assistant string `protobuf:"bytes,3,opt,name=assistant,proto3,oneof"`
}

func (*Role_System) isRole_Role() {}

func (*Role_User) isRole_Role() {}

func (*Role_Assistant) isRole_Role() {}

type ResponseFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ResponseFormat) Reset() {
	*x = ResponseFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFormat) ProtoMessage() {}

func (x *ResponseFormat) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFormat.ProtoReflect.Descriptor instead.
func (*ResponseFormat) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseFormat) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// role of the message author. One of "system", "user", "assistant".
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// content of the message
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{2}
}

func (x *ChatMessage) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type TextCompletionParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// temperature of the sampling, between [0, 2]. default = 1.0
	Temperature *float32 `protobuf:"fixed32,1,opt,name=temperature,proto3,oneof" json:"temperature,omitempty"`
	// whether to stream partial completions back as they are generated. default = false
	Stream *bool    `protobuf:"varint,2,opt,name=stream,proto3,oneof" json:"stream,omitempty"`
	TopK   *uint32  `protobuf:"varint,3,opt,name=top_k,json=topK,proto3,oneof" json:"top_k,omitempty"`
	TopP   *float32 `protobuf:"fixed32,4,opt,name=top_p,json=topP,proto3,oneof" json:"top_p,omitempty"`
	// number of chat completion choices to generate for each input message. default = 1
	N                *uint32         `protobuf:"varint,5,opt,name=n,proto3,oneof" json:"n,omitempty"`
	Stop             []string        `protobuf:"bytes,6,rep,name=stop,proto3" json:"stop,omitempty"`
	MaxTokens        *uint32         `protobuf:"varint,7,opt,name=max_tokens,json=maxTokens,proto3,oneof" json:"max_tokens,omitempty"`
	PresencePenalty  *float32        `protobuf:"fixed32,8,opt,name=presence_penalty,json=presencePenalty,proto3,oneof" json:"presence_penalty,omitempty"`
	FrequencyPenalty *float32        `protobuf:"fixed32,9,opt,name=frequency_penalty,json=frequencyPenalty,proto3,oneof" json:"frequency_penalty,omitempty"`
	ResponseFormat   *ResponseFormat `protobuf:"bytes,10,opt,name=response_format,json=responseFormat,proto3,oneof" json:"response_format,omitempty"`
	Seed             *uint32         `protobuf:"varint,11,opt,name=seed,proto3,oneof" json:"seed,omitempty"`
}

func (x *TextCompletionParameters) Reset() {
	*x = TextCompletionParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextCompletionParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextCompletionParameters) ProtoMessage() {}

func (x *TextCompletionParameters) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextCompletionParameters.ProtoReflect.Descriptor instead.
func (*TextCompletionParameters) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{3}
}

func (x *TextCompletionParameters) GetTemperature() float32 {
	if x != nil && x.Temperature != nil {
		return *x.Temperature
	}
	return 0
}

func (x *TextCompletionParameters) GetStream() bool {
	if x != nil && x.Stream != nil {
		return *x.Stream
	}
	return false
}

func (x *TextCompletionParameters) GetTopK() uint32 {
	if x != nil && x.TopK != nil {
		return *x.TopK
	}
	return 0
}

func (x *TextCompletionParameters) GetTopP() float32 {
	if x != nil && x.TopP != nil {
		return *x.TopP
	}
	return 0
}

func (x *TextCompletionParameters) GetN() uint32 {
	if x != nil && x.N != nil {
		return *x.N
	}
	return 0
}

func (x *TextCompletionParameters) GetStop() []string {
	if x != nil {
		return x.Stop
	}
	return nil
}

func (x *TextCompletionParameters) GetMaxTokens() uint32 {
	if x != nil && x.MaxTokens != nil {
		return *x.MaxTokens
	}
	return 0
}

func (x *TextCompletionParameters) GetPresencePenalty() float32 {
	if x != nil && x.PresencePenalty != nil {
		return *x.PresencePenalty
	}
	return 0
}

func (x *TextCompletionParameters) GetFrequencyPenalty() float32 {
	if x != nil && x.FrequencyPenalty != nil {
		return *x.FrequencyPenalty
	}
	return 0
}

func (x *TextCompletionParameters) GetResponseFormat() *ResponseFormat {
	if x != nil {
		return x.ResponseFormat
	}
	return nil
}

func (x *TextCompletionParameters) GetSeed() uint32 {
	if x != nil && x.Seed != nil {
		return *x.Seed
	}
	return 0
}

type CompletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	// a list of messages comprising all the conversation so far
	Messages   []*ChatMessage            `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Parameters *TextCompletionParameters `protobuf:"bytes,3,opt,name=parameters,proto3,oneof" json:"parameters,omitempty"`
}

func (x *CompletionRequest) Reset() {
	*x = CompletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionRequest) ProtoMessage() {}

func (x *CompletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionRequest.ProtoReflect.Descriptor instead.
func (*CompletionRequest) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{4}
}

func (x *CompletionRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CompletionRequest) GetMessages() []*ChatMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *CompletionRequest) GetParameters() *TextCompletionParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type CompletionChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// index of the choice in the list of choices.
	Index uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// message generated by the model.
	Message      *ChatMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FinishReason string       `protobuf:"bytes,3,opt,name=finish_reason,json=finishReason,proto3" json:"finish_reason,omitempty"`
}

func (x *CompletionChoice) Reset() {
	*x = CompletionChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionChoice) ProtoMessage() {}

func (x *CompletionChoice) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionChoice.ProtoReflect.Descriptor instead.
func (*CompletionChoice) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{5}
}

func (x *CompletionChoice) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CompletionChoice) GetMessage() *ChatMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *CompletionChoice) GetFinishReason() string {
	if x != nil {
		return x.FinishReason
	}
	return ""
}

type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number of tokens in the prompt.
	PromptTokens *int32 `protobuf:"varint,1,opt,name=prompt_tokens,json=promptTokens,proto3,oneof" json:"prompt_tokens,omitempty"`
	// number of tokens in the generated completion.
	CompletionTokens *int32 `protobuf:"varint,2,opt,name=completion_tokens,json=completionTokens,proto3,oneof" json:"completion_tokens,omitempty"`
	// total number of tokens used in the request (prompt + completion).
	TotalTokens *int32 `protobuf:"varint,3,opt,name=total_tokens,json=totalTokens,proto3,oneof" json:"total_tokens,omitempty"`
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{6}
}

func (x *Usage) GetPromptTokens() int32 {
	if x != nil && x.PromptTokens != nil {
		return *x.PromptTokens
	}
	return 0
}

func (x *Usage) GetCompletionTokens() int32 {
	if x != nil && x.CompletionTokens != nil {
		return *x.CompletionTokens
	}
	return 0
}

func (x *Usage) GetTotalTokens() int32 {
	if x != nil && x.TotalTokens != nil {
		return *x.TotalTokens
	}
	return 0
}

type CompletionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique id for the chat completion.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// object type, which is always "chat.completion[.chunk]".
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// unix timestamp (in seconds) of when the chat completion was created.
	Created uint64 `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	// model used for the completion
	Model string `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	// list of generated completion choices for the input prompt
	Choices []*CompletionChoice `protobuf:"bytes,5,rep,name=choices,proto3" json:"choices,omitempty"`
	// usage statistics for the completion request.
	Usage             *Usage `protobuf:"bytes,6,opt,name=usage,proto3" json:"usage,omitempty"`
	SystemFingerprint string `protobuf:"bytes,7,opt,name=system_fingerprint,json=systemFingerprint,proto3" json:"system_fingerprint,omitempty"`
}

func (x *CompletionResponse) Reset() {
	*x = CompletionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionResponse) ProtoMessage() {}

func (x *CompletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_llm_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionResponse.ProtoReflect.Descriptor instead.
func (*CompletionResponse) Descriptor() ([]byte, []int) {
	return file_llm_service_proto_rawDescGZIP(), []int{7}
}

func (x *CompletionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompletionResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *CompletionResponse) GetCreated() uint64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *CompletionResponse) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CompletionResponse) GetChoices() []*CompletionChoice {
	if x != nil {
		return x.Choices
	}
	return nil
}

func (x *CompletionResponse) GetUsage() *Usage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *CompletionResponse) GetSystemFingerprint() string {
	if x != nil {
		return x.SystemFingerprint
	}
	return ""
}

var File_llm_service_proto protoreflect.FileDescriptor

var file_llm_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x6c, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x09, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xcc, 0x04, 0x0a, 0x18, 0x54, 0x65, 0x78, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x4b, 0x88, 0x01,
	0x01, 0x12, 0x18, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x48, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x01, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74,
	0x6f, 0x70, 0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02,
	0x42, 0x0f, 0xba, 0x48, 0x0c, 0x0a, 0x0a, 0x15, 0x00, 0x00, 0x00, 0x40, 0x25, 0x00, 0x00, 0x00,
	0xc0, 0x48, 0x06, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x65, 0x6e,
	0x61, 0x6c, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x11, 0x66, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x02, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0x0a, 0x0a, 0x15, 0x00, 0x00, 0x00, 0x40, 0x25, 0x00,
	0x00, 0x00, 0xc0, 0x48, 0x07, 0x52, 0x10, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x48, 0x08, 0x52, 0x0e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x09,
	0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x6e, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x13, 0x0a,
	0x11, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c,
	0x74, 0x79, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x73, 0x65, 0x65, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x6c,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x7c, 0x0a, 0x10, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xc4, 0x01, 0x0a, 0x05, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22,
	0xf4, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x32,
	0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2a, 0x39, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x53, 0x54, 0x4f, 0x50, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x32, 0xf1, 0x01, 0x0a, 0x0a, 0x4c, 0x4c, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x69, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x78, 0x0a, 0x15, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x30, 0x01, 0x42, 0x89, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x6c,
	0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x73,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6c, 0x6c, 0x6d, 0x3b, 0x6c, 0x6c, 0x6d, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4c, 0x58, 0x58,
	0xaa, 0x02, 0x06, 0x4c, 0x6c, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x4c, 0x6c, 0x6d, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x12, 0x4c, 0x6c, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4c, 0x6c, 0x6d, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_llm_service_proto_rawDescOnce sync.Once
	file_llm_service_proto_rawDescData = file_llm_service_proto_rawDesc
)

func file_llm_service_proto_rawDescGZIP() []byte {
	file_llm_service_proto_rawDescOnce.Do(func() {
		file_llm_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_llm_service_proto_rawDescData)
	})
	return file_llm_service_proto_rawDescData
}

var file_llm_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_llm_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_llm_service_proto_goTypes = []interface{}{
	(FinishReason)(0),                // 0: llm.v1.FinishReason
	(*Role)(nil),                     // 1: llm.v1.Role
	(*ResponseFormat)(nil),           // 2: llm.v1.ResponseFormat
	(*ChatMessage)(nil),              // 3: llm.v1.ChatMessage
	(*TextCompletionParameters)(nil), // 4: llm.v1.TextCompletionParameters
	(*CompletionRequest)(nil),        // 5: llm.v1.CompletionRequest
	(*CompletionChoice)(nil),         // 6: llm.v1.CompletionChoice
	(*Usage)(nil),                    // 7: llm.v1.Usage
	(*CompletionResponse)(nil),       // 8: llm.v1.CompletionResponse
}
var file_llm_service_proto_depIdxs = []int32{
	2, // 0: llm.v1.TextCompletionParameters.response_format:type_name -> llm.v1.ResponseFormat
	3, // 1: llm.v1.CompletionRequest.messages:type_name -> llm.v1.ChatMessage
	4, // 2: llm.v1.CompletionRequest.parameters:type_name -> llm.v1.TextCompletionParameters
	3, // 3: llm.v1.CompletionChoice.message:type_name -> llm.v1.ChatMessage
	6, // 4: llm.v1.CompletionResponse.choices:type_name -> llm.v1.CompletionChoice
	7, // 5: llm.v1.CompletionResponse.usage:type_name -> llm.v1.Usage
	5, // 6: llm.v1.LLMService.ChatCompletions:input_type -> llm.v1.CompletionRequest
	5, // 7: llm.v1.LLMService.StreamChatCompletions:input_type -> llm.v1.CompletionRequest
	8, // 8: llm.v1.LLMService.ChatCompletions:output_type -> llm.v1.CompletionResponse
	8, // 9: llm.v1.LLMService.StreamChatCompletions:output_type -> llm.v1.CompletionResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_llm_service_proto_init() }
func file_llm_service_proto_init() {
	if File_llm_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_llm_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_llm_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFormat); i {
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
		file_llm_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_llm_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextCompletionParameters); i {
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
		file_llm_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletionRequest); i {
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
		file_llm_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletionChoice); i {
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
		file_llm_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
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
		file_llm_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletionResponse); i {
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
	file_llm_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Role_System)(nil),
		(*Role_User)(nil),
		(*Role_Assistant)(nil),
	}
	file_llm_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_llm_service_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_llm_service_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_llm_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_llm_service_proto_goTypes,
		DependencyIndexes: file_llm_service_proto_depIdxs,
		EnumInfos:         file_llm_service_proto_enumTypes,
		MessageInfos:      file_llm_service_proto_msgTypes,
	}.Build()
	File_llm_service_proto = out.File
	file_llm_service_proto_rawDesc = nil
	file_llm_service_proto_goTypes = nil
	file_llm_service_proto_depIdxs = nil
}
