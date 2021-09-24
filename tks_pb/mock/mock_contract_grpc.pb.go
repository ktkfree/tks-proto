// Code generated by MockGen. DO NOT EDIT.
// Source: tks_pb/contract_grpc.pb.go

// Package mock_tks_pb is a generated GoMock package.
package mock_tks_pb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tks_pb "github.com/openinfradev/tks-proto/tks_pb"
	grpc "google.golang.org/grpc"
)

// MockContractServiceClient is a mock of ContractServiceClient interface.
type MockContractServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockContractServiceClientMockRecorder
}

// MockContractServiceClientMockRecorder is the mock recorder for MockContractServiceClient.
type MockContractServiceClientMockRecorder struct {
	mock *MockContractServiceClient
}

// NewMockContractServiceClient creates a new mock instance.
func NewMockContractServiceClient(ctrl *gomock.Controller) *MockContractServiceClient {
	mock := &MockContractServiceClient{ctrl: ctrl}
	mock.recorder = &MockContractServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractServiceClient) EXPECT() *MockContractServiceClientMockRecorder {
	return m.recorder
}

// CreateContract mocks base method.
func (m *MockContractServiceClient) CreateContract(ctx context.Context, in *tks_pb.CreateContractRequest, opts ...grpc.CallOption) (*tks_pb.CreateContractResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateContract", varargs...)
	ret0, _ := ret[0].(*tks_pb.CreateContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContract indicates an expected call of CreateContract.
func (mr *MockContractServiceClientMockRecorder) CreateContract(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContract", reflect.TypeOf((*MockContractServiceClient)(nil).CreateContract), varargs...)
}

// GetAvailableServices mocks base method.
func (m *MockContractServiceClient) GetAvailableServices(ctx context.Context, in *tks_pb.GetAvailableServicesRequest, opts ...grpc.CallOption) (*tks_pb.GetAvailableServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAvailableServices", varargs...)
	ret0, _ := ret[0].(*tks_pb.GetAvailableServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableServices indicates an expected call of GetAvailableServices.
func (mr *MockContractServiceClientMockRecorder) GetAvailableServices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableServices", reflect.TypeOf((*MockContractServiceClient)(nil).GetAvailableServices), varargs...)
}

// GetContract mocks base method.
func (m *MockContractServiceClient) GetContract(ctx context.Context, in *tks_pb.GetContractRequest, opts ...grpc.CallOption) (*tks_pb.GetContractResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContract", varargs...)
	ret0, _ := ret[0].(*tks_pb.GetContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract.
func (mr *MockContractServiceClientMockRecorder) GetContract(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockContractServiceClient)(nil).GetContract), varargs...)
}

// GetContracts mocks base method.
func (m *MockContractServiceClient) GetContracts(ctx context.Context, in *tks_pb.GetContractsRequest, opts ...grpc.CallOption) (*tks_pb.GetContractsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContracts", varargs...)
	ret0, _ := ret[0].(*tks_pb.GetContractsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContracts indicates an expected call of GetContracts.
func (mr *MockContractServiceClientMockRecorder) GetContracts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContracts", reflect.TypeOf((*MockContractServiceClient)(nil).GetContracts), varargs...)
}

// GetQuota mocks base method.
func (m *MockContractServiceClient) GetQuota(ctx context.Context, in *tks_pb.GetQuotaRequest, opts ...grpc.CallOption) (*tks_pb.GetQuotaResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQuota", varargs...)
	ret0, _ := ret[0].(*tks_pb.GetQuotaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuota indicates an expected call of GetQuota.
func (mr *MockContractServiceClientMockRecorder) GetQuota(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuota", reflect.TypeOf((*MockContractServiceClient)(nil).GetQuota), varargs...)
}

// UpdateQuota mocks base method.
func (m *MockContractServiceClient) UpdateQuota(ctx context.Context, in *tks_pb.UpdateQuotaRequest, opts ...grpc.CallOption) (*tks_pb.UpdateQuotaResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateQuota", varargs...)
	ret0, _ := ret[0].(*tks_pb.UpdateQuotaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQuota indicates an expected call of UpdateQuota.
func (mr *MockContractServiceClientMockRecorder) UpdateQuota(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuota", reflect.TypeOf((*MockContractServiceClient)(nil).UpdateQuota), varargs...)
}

// UpdateServices mocks base method.
func (m *MockContractServiceClient) UpdateServices(ctx context.Context, in *tks_pb.UpdateServicesRequest, opts ...grpc.CallOption) (*tks_pb.UpdateServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateServices", varargs...)
	ret0, _ := ret[0].(*tks_pb.UpdateServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServices indicates an expected call of UpdateServices.
func (mr *MockContractServiceClientMockRecorder) UpdateServices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServices", reflect.TypeOf((*MockContractServiceClient)(nil).UpdateServices), varargs...)
}

// MockContractServiceServer is a mock of ContractServiceServer interface.
type MockContractServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockContractServiceServerMockRecorder
}

// MockContractServiceServerMockRecorder is the mock recorder for MockContractServiceServer.
type MockContractServiceServerMockRecorder struct {
	mock *MockContractServiceServer
}

// NewMockContractServiceServer creates a new mock instance.
func NewMockContractServiceServer(ctrl *gomock.Controller) *MockContractServiceServer {
	mock := &MockContractServiceServer{ctrl: ctrl}
	mock.recorder = &MockContractServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractServiceServer) EXPECT() *MockContractServiceServerMockRecorder {
	return m.recorder
}

// CreateContract mocks base method.
func (m *MockContractServiceServer) CreateContract(arg0 context.Context, arg1 *tks_pb.CreateContractRequest) (*tks_pb.CreateContractResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContract", arg0, arg1)
	ret0, _ := ret[0].(*tks_pb.CreateContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContract indicates an expected call of CreateContract.
func (mr *MockContractServiceServerMockRecorder) CreateContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContract", reflect.TypeOf((*MockContractServiceServer)(nil).CreateContract), arg0, arg1)
}

// GetAvailableServices mocks base method.
func (m *MockContractServiceServer) GetAvailableServices(arg0 context.Context, arg1 *tks_pb.GetAvailableServicesRequest) (*tks_pb.GetAvailableServicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableServices", arg0, arg1)
	ret0, _ := ret[0].(*tks_pb.GetAvailableServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableServices indicates an expected call of GetAvailableServices.
func (mr *MockContractServiceServerMockRecorder) GetAvailableServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableServices", reflect.TypeOf((*MockContractServiceServer)(nil).GetAvailableServices), arg0, arg1)
}

// GetContract mocks base method.
func (m *MockContractServiceServer) GetContract(arg0 context.Context, arg1 *tks_pb.GetContractRequest) (*tks_pb.GetContractResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract", arg0, arg1)
	ret0, _ := ret[0].(*tks_pb.GetContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract.
func (mr *MockContractServiceServerMockRecorder) GetContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockContractServiceServer)(nil).GetContract), arg0, arg1)
}

// GetContracts mocks base method.
func (m *MockContractServiceServer) GetContracts(arg0 context.Context, arg1 *tks_pb.GetContractsRequest) (*tks_pb.GetContractsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContracts", arg0, arg1)
	ret0, _ := ret[0].(*tks_pb.GetContractsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContracts indicates an expected call of GetContracts.
func (mr *MockContractServiceServerMockRecorder) GetContracts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContracts", reflect.TypeOf((*MockContractServiceServer)(nil).GetContracts), arg0, arg1)
}

// GetQuota mocks base method.
func (m *MockContractServiceServer) GetQuota(arg0 context.Context, arg1 *tks_pb.GetQuotaRequest) (*tks_pb.GetQuotaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuota", arg0, arg1)
	ret0, _ := ret[0].(*tks_pb.GetQuotaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuota indicates an expected call of GetQuota.
func (mr *MockContractServiceServerMockRecorder) GetQuota(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuota", reflect.TypeOf((*MockContractServiceServer)(nil).GetQuota), arg0, arg1)
}

// UpdateQuota mocks base method.
func (m *MockContractServiceServer) UpdateQuota(arg0 context.Context, arg1 *tks_pb.UpdateQuotaRequest) (*tks_pb.UpdateQuotaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuota", arg0, arg1)
	ret0, _ := ret[0].(*tks_pb.UpdateQuotaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQuota indicates an expected call of UpdateQuota.
func (mr *MockContractServiceServerMockRecorder) UpdateQuota(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuota", reflect.TypeOf((*MockContractServiceServer)(nil).UpdateQuota), arg0, arg1)
}

// UpdateServices mocks base method.
func (m *MockContractServiceServer) UpdateServices(arg0 context.Context, arg1 *tks_pb.UpdateServicesRequest) (*tks_pb.UpdateServicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServices", arg0, arg1)
	ret0, _ := ret[0].(*tks_pb.UpdateServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServices indicates an expected call of UpdateServices.
func (mr *MockContractServiceServerMockRecorder) UpdateServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServices", reflect.TypeOf((*MockContractServiceServer)(nil).UpdateServices), arg0, arg1)
}

// mustEmbedUnimplementedContractServiceServer mocks base method.
func (m *MockContractServiceServer) mustEmbedUnimplementedContractServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedContractServiceServer")
}

// mustEmbedUnimplementedContractServiceServer indicates an expected call of mustEmbedUnimplementedContractServiceServer.
func (mr *MockContractServiceServerMockRecorder) mustEmbedUnimplementedContractServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedContractServiceServer", reflect.TypeOf((*MockContractServiceServer)(nil).mustEmbedUnimplementedContractServiceServer))
}

// MockUnsafeContractServiceServer is a mock of UnsafeContractServiceServer interface.
type MockUnsafeContractServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeContractServiceServerMockRecorder
}

// MockUnsafeContractServiceServerMockRecorder is the mock recorder for MockUnsafeContractServiceServer.
type MockUnsafeContractServiceServerMockRecorder struct {
	mock *MockUnsafeContractServiceServer
}

// NewMockUnsafeContractServiceServer creates a new mock instance.
func NewMockUnsafeContractServiceServer(ctrl *gomock.Controller) *MockUnsafeContractServiceServer {
	mock := &MockUnsafeContractServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeContractServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeContractServiceServer) EXPECT() *MockUnsafeContractServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedContractServiceServer mocks base method.
func (m *MockUnsafeContractServiceServer) mustEmbedUnimplementedContractServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedContractServiceServer")
}

// mustEmbedUnimplementedContractServiceServer indicates an expected call of mustEmbedUnimplementedContractServiceServer.
func (mr *MockUnsafeContractServiceServerMockRecorder) mustEmbedUnimplementedContractServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedContractServiceServer", reflect.TypeOf((*MockUnsafeContractServiceServer)(nil).mustEmbedUnimplementedContractServiceServer))
}