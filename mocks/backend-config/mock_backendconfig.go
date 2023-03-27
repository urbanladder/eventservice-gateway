// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/backend-config (interfaces: BackendConfig)

// Package mock_backendconfig is a generated GoMock package.
package mock_backendconfig

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	identity "github.com/rudderlabs/rudder-server/services/controlplane/identity"
	pubsub "github.com/rudderlabs/rudder-server/utils/pubsub"
)

// MockBackendConfig is a mock of BackendConfig interface.
type MockBackendConfig struct {
	ctrl     *gomock.Controller
	recorder *MockBackendConfigMockRecorder
}

// MockBackendConfigMockRecorder is the mock recorder for MockBackendConfig.
type MockBackendConfigMockRecorder struct {
	mock *MockBackendConfig
}

// NewMockBackendConfig creates a new mock instance.
func NewMockBackendConfig(ctrl *gomock.Controller) *MockBackendConfig {
	mock := &MockBackendConfig{ctrl: ctrl}
	mock.recorder = &MockBackendConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendConfig) EXPECT() *MockBackendConfigMockRecorder {
	return m.recorder
}

// AccessToken mocks base method.
func (m *MockBackendConfig) AccessToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// AccessToken indicates an expected call of AccessToken.
func (mr *MockBackendConfigMockRecorder) AccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessToken", reflect.TypeOf((*MockBackendConfig)(nil).AccessToken))
}

// Get mocks base method.
func (m *MockBackendConfig) Get(arg0 context.Context) (map[string]backendconfig.ConfigT, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(map[string]backendconfig.ConfigT)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBackendConfigMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBackendConfig)(nil).Get), arg0)
}

// Identity mocks base method.
func (m *MockBackendConfig) Identity() identity.Identifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identity")
	ret0, _ := ret[0].(identity.Identifier)
	return ret0
}

// Identity indicates an expected call of Identity.
func (mr *MockBackendConfigMockRecorder) Identity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockBackendConfig)(nil).Identity))
}

// SetUp mocks base method.
func (m *MockBackendConfig) SetUp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUp")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUp indicates an expected call of SetUp.
func (mr *MockBackendConfigMockRecorder) SetUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUp", reflect.TypeOf((*MockBackendConfig)(nil).SetUp))
}

// StartWithIDs mocks base method.
func (m *MockBackendConfig) StartWithIDs(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartWithIDs", arg0, arg1)
}

// StartWithIDs indicates an expected call of StartWithIDs.
func (mr *MockBackendConfigMockRecorder) StartWithIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWithIDs", reflect.TypeOf((*MockBackendConfig)(nil).StartWithIDs), arg0, arg1)
}

// Stop mocks base method.
func (m *MockBackendConfig) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockBackendConfigMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBackendConfig)(nil).Stop))
}

// Subscribe mocks base method.
func (m *MockBackendConfig) Subscribe(arg0 context.Context, arg1 backendconfig.Topic) pubsub.DataChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(pubsub.DataChannel)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockBackendConfigMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockBackendConfig)(nil).Subscribe), arg0, arg1)
}

// WaitForConfig mocks base method.
func (m *MockBackendConfig) WaitForConfig(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WaitForConfig", arg0)
}

// WaitForConfig indicates an expected call of WaitForConfig.
func (mr *MockBackendConfigMockRecorder) WaitForConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForConfig", reflect.TypeOf((*MockBackendConfig)(nil).WaitForConfig), arg0)
}
