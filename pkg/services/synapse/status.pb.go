// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: services/synapse/status.proto

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

type AgentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId   string         `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Cpu       *CpuInfo       `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory    *ResourceUsage `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk      *ResourceUsage `protobuf:"bytes,4,opt,name=disk,proto3" json:"disk,omitempty"`
	Gpu       *GpuInfo       `protobuf:"bytes,5,opt,name=gpu,proto3" json:"gpu,omitempty"`
	Timestamp string         `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AgentStatusRequest) Reset() {
	*x = AgentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentStatusRequest) ProtoMessage() {}

func (x *AgentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentStatusRequest.ProtoReflect.Descriptor instead.
func (*AgentStatusRequest) Descriptor() ([]byte, []int) {
	return file_services_synapse_status_proto_rawDescGZIP(), []int{0}
}

func (x *AgentStatusRequest) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentStatusRequest) GetCpu() *CpuInfo {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *AgentStatusRequest) GetMemory() *ResourceUsage {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *AgentStatusRequest) GetDisk() *ResourceUsage {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *AgentStatusRequest) GetGpu() *GpuInfo {
	if x != nil {
		return x.Gpu
	}
	return nil
}

func (x *AgentStatusRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type CpuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName    string       `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`                     // CPU 型号名称（例如 Intel Core i7-10700K）
	Usage        float64      `protobuf:"fixed64,2,opt,name=usage,proto3" json:"usage,omitempty"`                                            // 当前总 CPU 利用率（百分比）
	PerCoreUsage []float64    `protobuf:"fixed64,3,rep,packed,name=per_core_usage,json=perCoreUsage,proto3" json:"per_core_usage,omitempty"` // 每个核心的利用率（百分比，可选）
	LoadAverage  *LoadAverage `protobuf:"bytes,4,opt,name=load_average,json=loadAverage,proto3" json:"load_average,omitempty"`               // 系统负载
	Threads      int32        `protobuf:"varint,5,opt,name=threads,proto3" json:"threads,omitempty"`                                         // 当前线程数
}

func (x *CpuInfo) Reset() {
	*x = CpuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuInfo) ProtoMessage() {}

func (x *CpuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuInfo.ProtoReflect.Descriptor instead.
func (*CpuInfo) Descriptor() ([]byte, []int) {
	return file_services_synapse_status_proto_rawDescGZIP(), []int{1}
}

func (x *CpuInfo) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *CpuInfo) GetUsage() float64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *CpuInfo) GetPerCoreUsage() []float64 {
	if x != nil {
		return x.PerCoreUsage
	}
	return nil
}

func (x *CpuInfo) GetLoadAverage() *LoadAverage {
	if x != nil {
		return x.LoadAverage
	}
	return nil
}

func (x *CpuInfo) GetThreads() int32 {
	if x != nil {
		return x.Threads
	}
	return 0
}

type LoadAverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OneMin     float64 `protobuf:"fixed64,1,opt,name=one_min,json=oneMin,proto3" json:"one_min,omitempty"`             // 过去 1 分钟的负载
	FiveMin    float64 `protobuf:"fixed64,2,opt,name=five_min,json=fiveMin,proto3" json:"five_min,omitempty"`          // 过去 5 分钟的负载
	FifteenMin float64 `protobuf:"fixed64,3,opt,name=fifteen_min,json=fifteenMin,proto3" json:"fifteen_min,omitempty"` // 过去 15 分钟的负载
}

func (x *LoadAverage) Reset() {
	*x = LoadAverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadAverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadAverage) ProtoMessage() {}

func (x *LoadAverage) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadAverage.ProtoReflect.Descriptor instead.
func (*LoadAverage) Descriptor() ([]byte, []int) {
	return file_services_synapse_status_proto_rawDescGZIP(), []int{2}
}

func (x *LoadAverage) GetOneMin() float64 {
	if x != nil {
		return x.OneMin
	}
	return 0
}

func (x *LoadAverage) GetFiveMin() float64 {
	if x != nil {
		return x.FiveMin
	}
	return 0
}

func (x *LoadAverage) GetFifteenMin() float64 {
	if x != nil {
		return x.FifteenMin
	}
	return 0
}

type ResourceUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total float64 `protobuf:"fixed64,1,opt,name=total,proto3" json:"total,omitempty"`
	Used  float64 `protobuf:"fixed64,2,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *ResourceUsage) Reset() {
	*x = ResourceUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceUsage) ProtoMessage() {}

func (x *ResourceUsage) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceUsage.ProtoReflect.Descriptor instead.
func (*ResourceUsage) Descriptor() ([]byte, []int) {
	return file_services_synapse_status_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceUsage) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ResourceUsage) GetUsed() float64 {
	if x != nil {
		return x.Used
	}
	return 0
}

type GpuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName   string  `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	TotalMemory float64 `protobuf:"fixed64,2,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	UsedMemory  float64 `protobuf:"fixed64,3,opt,name=used_memory,json=usedMemory,proto3" json:"used_memory,omitempty"`
}

func (x *GpuInfo) Reset() {
	*x = GpuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GpuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuInfo) ProtoMessage() {}

func (x *GpuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuInfo.ProtoReflect.Descriptor instead.
func (*GpuInfo) Descriptor() ([]byte, []int) {
	return file_services_synapse_status_proto_rawDescGZIP(), []int{4}
}

func (x *GpuInfo) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *GpuInfo) GetTotalMemory() float64 {
	if x != nil {
		return x.TotalMemory
	}
	return 0
}

func (x *GpuInfo) GetUsedMemory() float64 {
	if x != nil {
		return x.UsedMemory
	}
	return 0
}

type ReportAckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReportAckResponse) Reset() {
	*x = ReportAckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_synapse_status_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportAckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportAckResponse) ProtoMessage() {}

func (x *ReportAckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_synapse_status_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportAckResponse.ProtoReflect.Descriptor instead.
func (*ReportAckResponse) Descriptor() ([]byte, []int) {
	return file_services_synapse_status_proto_rawDescGZIP(), []int{5}
}

func (x *ReportAckResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ReportAckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_services_synapse_status_proto protoreflect.FileDescriptor

var file_services_synapse_status_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70,
	0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x22, 0xad, 0x02, 0x0a, 0x12, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x03, 0x63, 0x70, 0x75,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e,
	0x43, 0x70, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x3d, 0x0a, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x79,
	0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79,
	0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x04, 0x64,
	0x69, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x79, 0x6f, 0x74, 0x74,
	0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x31, 0x0a, 0x03, 0x67, 0x70, 0x75, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x47, 0x70, 0x75,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x67, 0x70, 0x75, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xc6, 0x01, 0x0a, 0x07, 0x43, 0x70, 0x75, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01,
	0x52, 0x0c, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46,
	0x0a, 0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73,
	0x22, 0x62, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x6f, 0x6e, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x6f, 0x6e, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x76, 0x65,
	0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x69, 0x76, 0x65,
	0x4d, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x66, 0x74, 0x65, 0x65, 0x6e, 0x5f, 0x6d,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x66, 0x69, 0x66, 0x74, 0x65, 0x65,
	0x6e, 0x4d, 0x69, 0x6e, 0x22, 0x39, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x22,
	0x6c, 0x0a, 0x07, 0x47, 0x70, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x47, 0x0a,
	0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x73, 0x61, 0x69,
	0x2f, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_synapse_status_proto_rawDescOnce sync.Once
	file_services_synapse_status_proto_rawDescData = file_services_synapse_status_proto_rawDesc
)

func file_services_synapse_status_proto_rawDescGZIP() []byte {
	file_services_synapse_status_proto_rawDescOnce.Do(func() {
		file_services_synapse_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_synapse_status_proto_rawDescData)
	})
	return file_services_synapse_status_proto_rawDescData
}

var file_services_synapse_status_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_synapse_status_proto_goTypes = []interface{}{
	(*AgentStatusRequest)(nil), // 0: yotta.services.synapse.AgentStatusRequest
	(*CpuInfo)(nil),            // 1: yotta.services.synapse.CpuInfo
	(*LoadAverage)(nil),        // 2: yotta.services.synapse.LoadAverage
	(*ResourceUsage)(nil),      // 3: yotta.services.synapse.ResourceUsage
	(*GpuInfo)(nil),            // 4: yotta.services.synapse.GpuInfo
	(*ReportAckResponse)(nil),  // 5: yotta.services.synapse.ReportAckResponse
}
var file_services_synapse_status_proto_depIdxs = []int32{
	1, // 0: yotta.services.synapse.AgentStatusRequest.cpu:type_name -> yotta.services.synapse.CpuInfo
	3, // 1: yotta.services.synapse.AgentStatusRequest.memory:type_name -> yotta.services.synapse.ResourceUsage
	3, // 2: yotta.services.synapse.AgentStatusRequest.disk:type_name -> yotta.services.synapse.ResourceUsage
	4, // 3: yotta.services.synapse.AgentStatusRequest.gpu:type_name -> yotta.services.synapse.GpuInfo
	2, // 4: yotta.services.synapse.CpuInfo.load_average:type_name -> yotta.services.synapse.LoadAverage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_synapse_status_proto_init() }
func file_services_synapse_status_proto_init() {
	if File_services_synapse_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_synapse_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentStatusRequest); i {
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
		file_services_synapse_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuInfo); i {
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
		file_services_synapse_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadAverage); i {
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
		file_services_synapse_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceUsage); i {
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
		file_services_synapse_status_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GpuInfo); i {
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
		file_services_synapse_status_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportAckResponse); i {
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
			RawDescriptor: file_services_synapse_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_synapse_status_proto_goTypes,
		DependencyIndexes: file_services_synapse_status_proto_depIdxs,
		MessageInfos:      file_services_synapse_status_proto_msgTypes,
	}.Build()
	File_services_synapse_status_proto = out.File
	file_services_synapse_status_proto_rawDesc = nil
	file_services_synapse_status_proto_goTypes = nil
	file_services_synapse_status_proto_depIdxs = nil
}
