// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	commonproto "github.com/zebraxid/common-services-proto"

	mock "github.com/stretchr/testify/mock"
)

// AuditTrailHandler is an autogenerated mock type for the AuditTrailHandler type
type AuditTrailHandler struct {
	mock.Mock
}

// ReadLog provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuditTrailHandler) ReadLog(_a0 context.Context, _a1 *commonproto.ReadLogRequest, _a2 *commonproto.AuditResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.ReadLogRequest, *commonproto.AuditResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendLog provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuditTrailHandler) SendLog(_a0 context.Context, _a1 *commonproto.SendLogRequest, _a2 *commonproto.AuditResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.SendLogRequest, *commonproto.AuditResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}