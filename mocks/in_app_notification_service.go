// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	client "github.com/micro/go-micro/v2/client"
	commonproto "github.com/zebraxid/common-services-proto"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// InAppNotificationService is an autogenerated mock type for the InAppNotificationService type
type InAppNotificationService struct {
	mock.Mock
}

// GetMessages provides a mock function with given fields: ctx, in, opts
func (_m *InAppNotificationService) GetMessages(ctx context.Context, in *commonproto.GetMessagesRequest, opts ...client.CallOption) (*commonproto.InAppNotifResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *commonproto.InAppNotifResponse
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.GetMessagesRequest, ...client.CallOption) *commonproto.InAppNotifResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonproto.InAppNotifResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *commonproto.GetMessagesRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctx, in, opts
func (_m *InAppNotificationService) SendMessage(ctx context.Context, in *commonproto.SendMessageRequest, opts ...client.CallOption) (*commonproto.InAppNotifResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *commonproto.InAppNotifResponse
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.SendMessageRequest, ...client.CallOption) *commonproto.InAppNotifResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonproto.InAppNotifResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *commonproto.SendMessageRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
