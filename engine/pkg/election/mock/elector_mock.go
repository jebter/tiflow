// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tiflow/engine/pkg/election (interfaces: Elector)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	election "github.com/pingcap/tiflow/engine/pkg/election"
)

// MockElector is a mock of Elector interface.
type MockElector struct {
	ctrl     *gomock.Controller
	recorder *MockElectorMockRecorder
}

// MockElectorMockRecorder is the mock recorder for MockElector.
type MockElectorMockRecorder struct {
	mock *MockElector
}

// NewMockElector creates a new mock instance.
func NewMockElector(ctrl *gomock.Controller) *MockElector {
	mock := &MockElector{ctrl: ctrl}
	mock.recorder = &MockElectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElector) EXPECT() *MockElectorMockRecorder {
	return m.recorder
}

// GetLeader mocks base method.
func (m *MockElector) GetLeader() (*election.Member, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeader")
	ret0, _ := ret[0].(*election.Member)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetLeader indicates an expected call of GetLeader.
func (mr *MockElectorMockRecorder) GetLeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeader", reflect.TypeOf((*MockElector)(nil).GetLeader))
}

// GetMembers mocks base method.
func (m *MockElector) GetMembers() []*election.Member {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembers")
	ret0, _ := ret[0].([]*election.Member)
	return ret0
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockElectorMockRecorder) GetMembers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockElector)(nil).GetMembers))
}

// IsLeader mocks base method.
func (m *MockElector) IsLeader() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLeader")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLeader indicates an expected call of IsLeader.
func (mr *MockElectorMockRecorder) IsLeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLeader", reflect.TypeOf((*MockElector)(nil).IsLeader))
}

// ResignLeader mocks base method.
func (m *MockElector) ResignLeader(arg0 context.Context, arg1 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResignLeader", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResignLeader indicates an expected call of ResignLeader.
func (mr *MockElectorMockRecorder) ResignLeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResignLeader", reflect.TypeOf((*MockElector)(nil).ResignLeader), arg0, arg1)
}

// Run mocks base method.
func (m *MockElector) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockElectorMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockElector)(nil).Run), arg0)
}