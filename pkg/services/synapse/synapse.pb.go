// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: services/synapse/synapse.proto

package synapse

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_synapse_synapse_proto protoreflect.FileDescriptor

var file_services_synapse_synapse_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70,
	0x73, 0x65, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x1a, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd9, 0x01, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x61, 0x70, 0x73,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x24, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x79, 0x6f, 0x74,
	0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61,
	0x70, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73,
	0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x79, 0x6f, 0x74, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x6f, 0x74, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x73, 0x61, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x70, 0x68, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_services_synapse_synapse_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),       // 0: yotta.services.synapse.HelloRequest
	(*AgentStatusRequest)(nil), // 1: yotta.services.synapse.AgentStatusRequest
	(*HelloResponse)(nil),      // 2: yotta.services.synapse.HelloResponse
	(*ReportAckResponse)(nil),  // 3: yotta.services.synapse.ReportAckResponse
}
var file_services_synapse_synapse_proto_depIdxs = []int32{
	0, // 0: yotta.services.synapse.SynapseService.SayHello:input_type -> yotta.services.synapse.HelloRequest
	1, // 1: yotta.services.synapse.SynapseService.ReportAgentStatus:input_type -> yotta.services.synapse.AgentStatusRequest
	2, // 2: yotta.services.synapse.SynapseService.SayHello:output_type -> yotta.services.synapse.HelloResponse
	3, // 3: yotta.services.synapse.SynapseService.ReportAgentStatus:output_type -> yotta.services.synapse.ReportAckResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_synapse_synapse_proto_init() }
func file_services_synapse_synapse_proto_init() {
	if File_services_synapse_synapse_proto != nil {
		return
	}
	file_services_synapse_hello_proto_init()
	file_services_synapse_status_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_synapse_synapse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_synapse_synapse_proto_goTypes,
		DependencyIndexes: file_services_synapse_synapse_proto_depIdxs,
	}.Build()
	File_services_synapse_synapse_proto = out.File
	file_services_synapse_synapse_proto_rawDesc = nil
	file_services_synapse_synapse_proto_goTypes = nil
	file_services_synapse_synapse_proto_depIdxs = nil
}
