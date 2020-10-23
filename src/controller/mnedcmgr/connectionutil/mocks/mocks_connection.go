/*******************************************************************************
 * Copyright 2020 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

// Code generated by MockGen. DO NOT EDIT.
// Source: controller/mnedcmgr/connectionutil (interfaces: NetworkUtil)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	net "net"
	reflect "reflect"
)

// MockNetworkUtil is a mock of NetworkUtil interface.
type MockNetworkUtil struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkUtilMockRecorder
}

// MockNetworkUtilMockRecorder is the mock recorder for MockNetworkUtil.
type MockNetworkUtilMockRecorder struct {
	mock *MockNetworkUtil
}

// NewMockNetworkUtil creates a new mock instance.
func NewMockNetworkUtil(ctrl *gomock.Controller) *MockNetworkUtil {
	mock := &MockNetworkUtil{ctrl: ctrl}
	mock.recorder = &MockNetworkUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkUtil) EXPECT() *MockNetworkUtilMockRecorder {
	return m.recorder
}

// ConnectToHost mocks base method.
func (m *MockNetworkUtil) ConnectToHost(arg0, arg1 string, arg2 bool) (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToHost", arg0, arg1, arg2)
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectToHost indicates an expected call of ConnectToHost.
func (mr *MockNetworkUtilMockRecorder) ConnectToHost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToHost", reflect.TypeOf((*MockNetworkUtil)(nil).ConnectToHost), arg0, arg1, arg2)
}

// ListenIP mocks base method.
func (m *MockNetworkUtil) ListenIP(arg0 string, arg1 bool) (net.Listener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenIP", arg0, arg1)
	ret0, _ := ret[0].(net.Listener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenIP indicates an expected call of ListenIP.
func (mr *MockNetworkUtilMockRecorder) ListenIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenIP", reflect.TypeOf((*MockNetworkUtil)(nil).ListenIP), arg0, arg1)
}

// ReadFrom mocks base method.
func (m *MockNetworkUtil) ReadFrom(arg0 net.Conn) (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFrom", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadFrom indicates an expected call of ReadFrom.
func (mr *MockNetworkUtilMockRecorder) ReadFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFrom", reflect.TypeOf((*MockNetworkUtil)(nil).ReadFrom), arg0)
}

// WriteTo mocks base method.
func (m *MockNetworkUtil) WriteTo(arg0 net.Conn, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTo indicates an expected call of WriteTo.
func (mr *MockNetworkUtilMockRecorder) WriteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockNetworkUtil)(nil).WriteTo), arg0, arg1)
}
