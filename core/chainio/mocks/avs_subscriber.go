// Code generated by MockGen. DO NOT EDIT.
// Source: anzen-avs/core/chainio (interfaces: AvsSubscriberer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avs_subscriber.go -package=mocks anzen-avs/core/chainio AvsSubscriberer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	contractAnzenTaskManager "anzen-avs/contracts/bindings/AnzenTaskManager"
	reflect "reflect"

	types "github.com/ethereum/go-ethereum/core/types"
	event "github.com/ethereum/go-ethereum/event"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsSubscriberer is a mock of AvsSubscriberer interface.
type MockAvsSubscriberer struct {
	ctrl     *gomock.Controller
	recorder *MockAvsSubscribererMockRecorder
}

// MockAvsSubscribererMockRecorder is the mock recorder for MockAvsSubscriberer.
type MockAvsSubscribererMockRecorder struct {
	mock *MockAvsSubscriberer
}

// NewMockAvsSubscriberer creates a new mock instance.
func NewMockAvsSubscriberer(ctrl *gomock.Controller) *MockAvsSubscriberer {
	mock := &MockAvsSubscriberer{ctrl: ctrl}
	mock.recorder = &MockAvsSubscribererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsSubscriberer) EXPECT() *MockAvsSubscribererMockRecorder {
	return m.recorder
}

// ParseTaskResponded mocks base method.
func (m *MockAvsSubscriberer) ParseTaskResponded(arg0 types.Log) (*contractAnzenTaskManager.ContractAnzenTaskManagerTaskResponded, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseTaskResponded", arg0)
	ret0, _ := ret[0].(*contractAnzenTaskManager.ContractAnzenTaskManagerTaskResponded)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseTaskResponded indicates an expected call of ParseTaskResponded.
func (mr *MockAvsSubscribererMockRecorder) ParseTaskResponded(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseTaskResponded", reflect.TypeOf((*MockAvsSubscriberer)(nil).ParseTaskResponded), arg0)
}

// SubcribeToNewOraclePullTasks mocks base method.
func (m *MockAvsSubscriberer) SubcribeToNewOraclePullTasks(arg0 chan *contractAnzenTaskManager.ContractAnzenTaskManagerNewOraclePullTaskCreated) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubcribeToNewOraclePullTasks", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubcribeToNewOraclePullTasks indicates an expected call of SubcribeToNewOraclePullTasks.
func (mr *MockAvsSubscribererMockRecorder) SubcribeToNewOraclePullTasks(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubcribeToNewOraclePullTasks", reflect.TypeOf((*MockAvsSubscriberer)(nil).SubcribeToNewOraclePullTasks), arg0)
}

// SubscribeToNewTasks mocks base method.
func (m *MockAvsSubscriberer) SubscribeToNewTasks(arg0 chan *contractAnzenTaskManager.ContractAnzenTaskManagerNewTaskCreated) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToNewTasks", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeToNewTasks indicates an expected call of SubscribeToNewTasks.
func (mr *MockAvsSubscribererMockRecorder) SubscribeToNewTasks(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNewTasks", reflect.TypeOf((*MockAvsSubscriberer)(nil).SubscribeToNewTasks), arg0)
}

// SubscribeToTaskResponses mocks base method.
func (m *MockAvsSubscriberer) SubscribeToTaskResponses(arg0 chan *contractAnzenTaskManager.ContractAnzenTaskManagerTaskResponded) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToTaskResponses", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeToTaskResponses indicates an expected call of SubscribeToTaskResponses.
func (mr *MockAvsSubscribererMockRecorder) SubscribeToTaskResponses(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToTaskResponses", reflect.TypeOf((*MockAvsSubscriberer)(nil).SubscribeToTaskResponses), arg0)
}
