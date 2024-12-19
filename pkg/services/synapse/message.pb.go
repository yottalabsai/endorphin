// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: services/synapse/message.proto

package synapse

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

type BaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SenderId  string `protobuf:"bytes,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
}

func (x *BaseMessage) Reset() {
	*x = BaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseMessage) ProtoMessage() {}

func (x *BaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseMessage.ProtoReflect.Descriptor instead.
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return file_services_synapse_message_proto_rawDescGZIP(), []int{0}
}

func (x *BaseMessage) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *BaseMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BaseMessage) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

type StreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *BaseMessage      `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Payload:
	//
	//	*StreamMessage_Heartbeat
	//	*StreamMessage_RunModel
	//	*StreamMessage_RunModelResponse
	//	*StreamMessage_Inference
	//	*StreamMessage_InferenceResponse
	Payload isStreamMessage_Payload `protobuf_oneof:"payload"`
}

func (x *StreamMessage) Reset() {
	*x = StreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessage) ProtoMessage() {}

func (x *StreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessage.ProtoReflect.Descriptor instead.
func (*StreamMessage) Descriptor() ([]byte, []int) {
	return file_services_synapse_message_proto_rawDescGZIP(), []int{1}
}

func (x *StreamMessage) GetBase() *BaseMessage {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *StreamMessage) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (m *StreamMessage) GetPayload() isStreamMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *StreamMessage) GetHeartbeat() *HeartbeatRequest {
	if x, ok := x.GetPayload().(*StreamMessage_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (x *StreamMessage) GetRunModel() *RunModelRequest {
	if x, ok := x.GetPayload().(*StreamMessage_RunModel); ok {
		return x.RunModel
	}
	return nil
}

func (x *StreamMessage) GetRunModelResponse() *RunModelResponse {
	if x, ok := x.GetPayload().(*StreamMessage_RunModelResponse); ok {
		return x.RunModelResponse
	}
	return nil
}

func (x *StreamMessage) GetInference() *InferenceRequest {
	if x, ok := x.GetPayload().(*StreamMessage_Inference); ok {
		return x.Inference
	}
	return nil
}

func (x *StreamMessage) GetInferenceResponse() *InferenceResponse {
	if x, ok := x.GetPayload().(*StreamMessage_InferenceResponse); ok {
		return x.InferenceResponse
	}
	return nil
}

type isStreamMessage_Payload interface {
	isStreamMessage_Payload()
}

type StreamMessage_Heartbeat struct {
	Heartbeat *HeartbeatRequest `protobuf:"bytes,3,opt,name=heartbeat,proto3,oneof"`
}

type StreamMessage_RunModel struct {
	RunModel *RunModelRequest `protobuf:"bytes,4,opt,name=run_model,json=runModel,proto3,oneof"`
}

type StreamMessage_RunModelResponse struct {
	RunModelResponse *RunModelResponse `protobuf:"bytes,5,opt,name=run_model_response,json=runModelResponse,proto3,oneof"`
}

type StreamMessage_Inference struct {
	Inference *InferenceRequest `protobuf:"bytes,6,opt,name=inference,proto3,oneof"`
}

type StreamMessage_InferenceResponse struct {
	InferenceResponse *InferenceResponse `protobuf:"bytes,7,opt,name=inference_response,json=inferenceResponse,proto3,oneof"`
}

func (*StreamMessage_Heartbeat) isStreamMessage_Payload() {}

func (*StreamMessage_RunModel) isStreamMessage_Payload() {}

func (*StreamMessage_RunModelResponse) isStreamMessage_Payload() {}

func (*StreamMessage_Inference) isStreamMessage_Payload() {}

func (*StreamMessage_InferenceResponse) isStreamMessage_Payload() {}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_services_synapse_message_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type RunModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelId string `protobuf:"bytes,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	Prompt  string `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *RunModelRequest) Reset() {
	*x = RunModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunModelRequest) ProtoMessage() {}

func (x *RunModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunModelRequest.ProtoReflect.Descriptor instead.
func (*RunModelRequest) Descriptor() ([]byte, []int) {
	return file_services_synapse_message_proto_rawDescGZIP(), []int{3}
}

func (x *RunModelRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *RunModelRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

type RunModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunModelMessageId string `protobuf:"bytes,1,opt,name=run_model_message_id,json=runModelMessageId,proto3" json:"run_model_message_id,omitempty"`
	Code              int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message           string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RunModelResponse) Reset() {
	*x = RunModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunModelResponse) ProtoMessage() {}

func (x *RunModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunModelResponse.ProtoReflect.Descriptor instead.
func (*RunModelResponse) Descriptor() ([]byte, []int) {
	return file_services_synapse_message_proto_rawDescGZIP(), []int{4}
}

func (x *RunModelResponse) GetRunModelMessageId() string {
	if x != nil {
		return x.RunModelMessageId
	}
	return ""
}

func (x *RunModelResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RunModelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type InferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelId  string              `protobuf:"bytes,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	Contents []*InferenceContent `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents,omitempty"`
}

func (x *InferenceRequest) Reset() {
	*x = InferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceRequest) ProtoMessage() {}

func (x *InferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceRequest.ProtoReflect.Descriptor instead.
func (*InferenceRequest) Descriptor() ([]byte, []int) {
	return file_services_synapse_message_proto_rawDescGZIP(), []int{5}
}

func (x *InferenceRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *InferenceRequest) GetContents() []*InferenceContent {
	if x != nil {
		return x.Contents
	}
	return nil
}

type InferenceContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *InferenceContent) Reset() {
	*x = InferenceContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceContent) ProtoMessage() {}

func (x *InferenceContent) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceContent.ProtoReflect.Descriptor instead.
func (*InferenceContent) Descriptor() ([]byte, []int) {
	return file_services_synapse_message_proto_rawDescGZIP(), []int{6}
}

func (x *InferenceContent) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *InferenceContent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type InferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InferenceMessageId string `protobuf:"bytes,1,opt,name=inference_message_id,json=inferenceMessageId,proto3" json:"inference_message_id,omitempty"`
	IsStream           bool   `protobuf:"varint,2,opt,name=is_stream,json=isStream,proto3" json:"is_stream,omitempty"`
	Result             string `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"` // stringify json result
}

func (x *InferenceResponse) Reset() {
	*x = InferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceResponse) ProtoMessage() {}

func (x *InferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceResponse.ProtoReflect.Descriptor instead.
func (*InferenceResponse) Descriptor() ([]byte, []int) {
	return file_services_synapse_message_proto_rawDescGZIP(), []int{7}
}

func (x *InferenceResponse) GetInferenceMessageId() string {
	if x != nil {
		return x.InferenceMessageId
	}
	return ""
}

func (x *InferenceResponse) GetIsStream() bool {
	if x != nil {
		return x.IsStream
	}
	return false
}

func (x *InferenceResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_services_synapse_message_proto protoreflect.FileDescriptor

var file_services_synapse_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70,
	0x73, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x22, 0x67, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xf3, 0x04, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x48, 0x0a,
	0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x46, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79, 0x6f, 0x74,
	0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61,
	0x70, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x58, 0x0a, 0x12, 0x72, 0x75, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x6f,
	0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e,
	0x61, 0x70, 0x73, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x10, 0x72, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x69, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79,
	0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79,
	0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x12, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x11, 0x69, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2f, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x71,
	0x0a, 0x10, 0x52, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x75, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x73, 0x0a, 0x10, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x44, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x40, 0x0a, 0x10, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x11, 0x49, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x73, 0x61, 0x69, 0x2f, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_synapse_message_proto_rawDescOnce sync.Once
	file_services_synapse_message_proto_rawDescData = file_services_synapse_message_proto_rawDesc
)

func file_services_synapse_message_proto_rawDescGZIP() []byte {
	file_services_synapse_message_proto_rawDescOnce.Do(func() {
		file_services_synapse_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_synapse_message_proto_rawDescData)
	})
	return file_services_synapse_message_proto_rawDescData
}

var file_services_synapse_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_services_synapse_message_proto_goTypes = []interface{}{
	(*BaseMessage)(nil),       // 0: yotta.services.synapse.BaseMessage
	(*StreamMessage)(nil),     // 1: yotta.services.synapse.StreamMessage
	(*HeartbeatRequest)(nil),  // 2: yotta.services.synapse.HeartbeatRequest
	(*RunModelRequest)(nil),   // 3: yotta.services.synapse.RunModelRequest
	(*RunModelResponse)(nil),  // 4: yotta.services.synapse.RunModelResponse
	(*InferenceRequest)(nil),  // 5: yotta.services.synapse.InferenceRequest
	(*InferenceContent)(nil),  // 6: yotta.services.synapse.InferenceContent
	(*InferenceResponse)(nil), // 7: yotta.services.synapse.InferenceResponse
	nil,                       // 8: yotta.services.synapse.StreamMessage.MetadataEntry
}
var file_services_synapse_message_proto_depIdxs = []int32{
	0, // 0: yotta.services.synapse.StreamMessage.base:type_name -> yotta.services.synapse.BaseMessage
	8, // 1: yotta.services.synapse.StreamMessage.metadata:type_name -> yotta.services.synapse.StreamMessage.MetadataEntry
	2, // 2: yotta.services.synapse.StreamMessage.heartbeat:type_name -> yotta.services.synapse.HeartbeatRequest
	3, // 3: yotta.services.synapse.StreamMessage.run_model:type_name -> yotta.services.synapse.RunModelRequest
	4, // 4: yotta.services.synapse.StreamMessage.run_model_response:type_name -> yotta.services.synapse.RunModelResponse
	5, // 5: yotta.services.synapse.StreamMessage.inference:type_name -> yotta.services.synapse.InferenceRequest
	7, // 6: yotta.services.synapse.StreamMessage.inference_response:type_name -> yotta.services.synapse.InferenceResponse
	6, // 7: yotta.services.synapse.InferenceRequest.contents:type_name -> yotta.services.synapse.InferenceContent
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_services_synapse_message_proto_init() }
func file_services_synapse_message_proto_init() {
	if File_services_synapse_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_synapse_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseMessage); i {
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
		file_services_synapse_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMessage); i {
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
		file_services_synapse_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_services_synapse_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunModelRequest); i {
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
		file_services_synapse_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunModelResponse); i {
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
		file_services_synapse_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceRequest); i {
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
		file_services_synapse_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceContent); i {
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
		file_services_synapse_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceResponse); i {
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
	file_services_synapse_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StreamMessage_Heartbeat)(nil),
		(*StreamMessage_RunModel)(nil),
		(*StreamMessage_RunModelResponse)(nil),
		(*StreamMessage_Inference)(nil),
		(*StreamMessage_InferenceResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_synapse_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_synapse_message_proto_goTypes,
		DependencyIndexes: file_services_synapse_message_proto_depIdxs,
		MessageInfos:      file_services_synapse_message_proto_msgTypes,
	}.Build()
	File_services_synapse_message_proto = out.File
	file_services_synapse_message_proto_rawDesc = nil
	file_services_synapse_message_proto_goTypes = nil
	file_services_synapse_message_proto_depIdxs = nil
}
