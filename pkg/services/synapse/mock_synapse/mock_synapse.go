// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/synapse/synapse_grpc.pb.go

// Package mock_synapse is a generated GoMock package.
package mock_synapse

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	synapse "github.com/yottalabsai/endorphin/pkg/services/synapse"
	grpc "google.golang.org/grpc"
)

// MockSynapseServiceClient is a mock of SynapseServiceClient interface.
type MockSynapseServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockSynapseServiceClientMockRecorder
}

// MockSynapseServiceClientMockRecorder is the mock recorder for MockSynapseServiceClient.
type MockSynapseServiceClientMockRecorder struct {
	mock *MockSynapseServiceClient
}

// NewMockSynapseServiceClient creates a new mock instance.
func NewMockSynapseServiceClient(ctrl *gomock.Controller) *MockSynapseServiceClient {
	mock := &MockSynapseServiceClient{ctrl: ctrl}
	mock.recorder = &MockSynapseServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSynapseServiceClient) EXPECT() *MockSynapseServiceClientMockRecorder {
	return m.recorder
}

// ReportAgentStatus mocks base method.
func (m *MockSynapseServiceClient) ReportAgentStatus(ctx context.Context, in *synapse.AgentStatusRequest, opts ...grpc.CallOption) (*synapse.ReportAckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReportAgentStatus", varargs...)
	ret0, _ := ret[0].(*synapse.ReportAckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportAgentStatus indicates an expected call of ReportAgentStatus.
func (mr *MockSynapseServiceClientMockRecorder) ReportAgentStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportAgentStatus", reflect.TypeOf((*MockSynapseServiceClient)(nil).ReportAgentStatus), varargs...)
}

// SayHello mocks base method.
func (m *MockSynapseServiceClient) SayHello(ctx context.Context, in *synapse.HelloRequest, opts ...grpc.CallOption) (*synapse.HelloResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SayHello", varargs...)
	ret0, _ := ret[0].(*synapse.HelloResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello.
func (mr *MockSynapseServiceClientMockRecorder) SayHello(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockSynapseServiceClient)(nil).SayHello), varargs...)
}

// MockSynapseServiceServer is a mock of SynapseServiceServer interface.
type MockSynapseServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockSynapseServiceServerMockRecorder
}

// MockSynapseServiceServerMockRecorder is the mock recorder for MockSynapseServiceServer.
type MockSynapseServiceServerMockRecorder struct {
	mock *MockSynapseServiceServer
}

// NewMockSynapseServiceServer creates a new mock instance.
func NewMockSynapseServiceServer(ctrl *gomock.Controller) *MockSynapseServiceServer {
	mock := &MockSynapseServiceServer{ctrl: ctrl}
	mock.recorder = &MockSynapseServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSynapseServiceServer) EXPECT() *MockSynapseServiceServerMockRecorder {
	return m.recorder
}

// ReportAgentStatus mocks base method.
func (m *MockSynapseServiceServer) ReportAgentStatus(arg0 context.Context, arg1 *synapse.AgentStatusRequest) (*synapse.ReportAckResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportAgentStatus", arg0, arg1)
	ret0, _ := ret[0].(*synapse.ReportAckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportAgentStatus indicates an expected call of ReportAgentStatus.
func (mr *MockSynapseServiceServerMockRecorder) ReportAgentStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportAgentStatus", reflect.TypeOf((*MockSynapseServiceServer)(nil).ReportAgentStatus), arg0, arg1)
}

// SayHello mocks base method.
func (m *MockSynapseServiceServer) SayHello(arg0 context.Context, arg1 *synapse.HelloRequest) (*synapse.HelloResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SayHello", arg0, arg1)
	ret0, _ := ret[0].(*synapse.HelloResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello.
func (mr *MockSynapseServiceServerMockRecorder) SayHello(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockSynapseServiceServer)(nil).SayHello), arg0, arg1)
}

// MockUnsafeSynapseServiceServer is a mock of UnsafeSynapseServiceServer interface.
type MockUnsafeSynapseServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSynapseServiceServerMockRecorder
}

// MockUnsafeSynapseServiceServerMockRecorder is the mock recorder for MockUnsafeSynapseServiceServer.
type MockUnsafeSynapseServiceServerMockRecorder struct {
	mock *MockUnsafeSynapseServiceServer
}

// NewMockUnsafeSynapseServiceServer creates a new mock instance.
func NewMockUnsafeSynapseServiceServer(ctrl *gomock.Controller) *MockUnsafeSynapseServiceServer {
	mock := &MockUnsafeSynapseServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSynapseServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSynapseServiceServer) EXPECT() *MockUnsafeSynapseServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSynapseServiceServer mocks base method.
func (m *MockUnsafeSynapseServiceServer) mustEmbedUnimplementedSynapseServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSynapseServiceServer")
}

// mustEmbedUnimplementedSynapseServiceServer indicates an expected call of mustEmbedUnimplementedSynapseServiceServer.
func (mr *MockUnsafeSynapseServiceServerMockRecorder) mustEmbedUnimplementedSynapseServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSynapseServiceServer", reflect.TypeOf((*MockUnsafeSynapseServiceServer)(nil).mustEmbedUnimplementedSynapseServiceServer))
}
