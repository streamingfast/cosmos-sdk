// Code generated by MockGen. DO NOT EDIT.
// Source: types/router.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockQueryRouter is a mock of QueryRouter interface.
type MockQueryRouter struct {
	ctrl     *gomock.Controller
	recorder *MockQueryRouterMockRecorder
}

// MockQueryRouterMockRecorder is the mock recorder for MockQueryRouter.
type MockQueryRouterMockRecorder struct {
	mock *MockQueryRouter
}

// NewMockQueryRouter creates a new mock instance.
func NewMockQueryRouter(ctrl *gomock.Controller) *MockQueryRouter {
	mock := &MockQueryRouter{ctrl: ctrl}
	mock.recorder = &MockQueryRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryRouter) EXPECT() *MockQueryRouterMockRecorder {
	return m.recorder
}

// AddRoute mocks base method.
func (m *MockQueryRouter) AddRoute(r string, h types.Querier) types.QueryRouter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", r, h)
	ret0, _ := ret[0].(types.QueryRouter)
	return ret0
}

// AddRoute indicates an expected call of AddRoute.
func (mr *MockQueryRouterMockRecorder) AddRoute(r, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockQueryRouter)(nil).AddRoute), r, h)
}

// Route mocks base method.
func (m *MockQueryRouter) Route(path string) types.Querier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", path)
	ret0, _ := ret[0].(types.Querier)
	return ret0
}

// Route indicates an expected call of Route.
func (mr *MockQueryRouterMockRecorder) Route(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockQueryRouter)(nil).Route), path)
}
