// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	commonproto "github.com/zebraxid/common-services-proto"

	mock "github.com/stretchr/testify/mock"
)

// NotificationHandler is an autogenerated mock type for the NotificationHandler type
type NotificationHandler struct {
	mock.Mock
}

// BulkSendNotification provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) BulkSendNotification(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateNotificationGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) CreateNotificationGroup(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateNotificationGroupMember provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) CreateNotificationGroupMember(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNotificationGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) DeleteNotificationGroup(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNotificationGroupMember provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) DeleteNotificationGroupMember(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListNotificationGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) ListNotificationGroup(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListNotificationGroupMember provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) ListNotificationGroupMember(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListNotificationLog provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) ListNotificationLog(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListNotificationPermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) ListNotificationPermission(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendNotification provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) SendNotification(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNotificationPermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotificationHandler) UpdateNotificationPermission(_a0 context.Context, _a1 *commonproto.NotificationRequest, _a2 *commonproto.NotificationResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.NotificationRequest, *commonproto.NotificationResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}