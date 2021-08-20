// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	commonproto "github.com/zebraxid/common-services-proto"

	mock "github.com/stretchr/testify/mock"
)

// AuthHandler is an autogenerated mock type for the AuthHandler type
type AuthHandler struct {
	mock.Mock
}

// AddCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) AddCasbinPolicy(_a0 context.Context, _a1 *commonproto.CasbinPolicyPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.CasbinPolicyPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthorizeToken provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) AuthorizeToken(_a0 context.Context, _a1 *commonproto.CoreRequest, _a2 *commonproto.AuthResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.CoreRequest, *commonproto.AuthResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BulkAddCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) BulkAddCasbinPolicy(_a0 context.Context, _a1 *commonproto.BulkCasbinPolicyPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.BulkCasbinPolicyPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BulkAddUserGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) BulkAddUserGroup(_a0 context.Context, _a1 *commonproto.BulkAddUserGroupPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.BulkAddUserGroupPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BulkListCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) BulkListCasbinPolicy(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BulkRemoveCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) BulkRemoveCasbinPolicy(_a0 context.Context, _a1 *commonproto.BulkCasbinPolicyPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.BulkCasbinPolicyPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BulkReplaceCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) BulkReplaceCasbinPolicy(_a0 context.Context, _a1 *commonproto.BulkReplaceCasbinPolicyPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.BulkReplaceCasbinPolicyPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangePassword provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ChangePassword(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAttributePolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) CreateAttributePolicy(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOrganization provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) CreateOrganization(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) CreatePermission(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProduct provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) CreateProduct(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ProductResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ProductResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProductOrganization provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) CreateProductOrganization(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ProductResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ProductResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRole provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) CreateRole(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRolePermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) CreateRolePermission(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) CreateUser(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DMAAListUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DMAAListUser(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAttributePolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeleteAttributePolicy(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeleteGroup(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrganization provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeleteOrganization(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeletePermission(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProduct provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeleteProduct(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ProductResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ProductResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProductOrganization provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeleteProductOrganization(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ProductResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ProductResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRole provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeleteRole(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRolePermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeleteRolePermission(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) DeleteUser(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProduct provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) GetProduct(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProfile provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) GetProfile(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListAttributePolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListAttributePolicy(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListCasbinPolicy(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListOrganization provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListOrganization(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListPermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListPermission(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListProduct provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListProduct(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListProductOrganization provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListProductOrganization(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListRole provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListRole(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListRolePermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListRolePermission(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ListUser(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ListResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ListResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogoutToken provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) LogoutToken(_a0 context.Context, _a1 *commonproto.CoreRequest, _a2 *commonproto.AuthResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.CoreRequest, *commonproto.AuthResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OneTimePassword provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) OneTimePassword(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RefreshToken provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) RefreshToken(_a0 context.Context, _a1 *commonproto.CoreRequest, _a2 *commonproto.AuthResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.CoreRequest, *commonproto.AuthResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) RemoveCasbinPolicy(_a0 context.Context, _a1 *commonproto.CasbinPolicyPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.CasbinPolicyPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReplaceCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ReplaceCasbinPolicy(_a0 context.Context, _a1 *commonproto.ReplaceCasbinPolicyPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.ReplaceCasbinPolicyPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestToken provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) RequestToken(_a0 context.Context, _a1 *commonproto.CoreRequest, _a2 *commonproto.AuthResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.CoreRequest, *commonproto.AuthResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetPassword provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ResetPassword(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAttributePolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdateAttributePolicy(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrganization provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdateOrganization(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdatePermission(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProduct provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdateProduct(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ProductResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ProductResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProductOrganization provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdateProductOrganization(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.ProductResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.ProductResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfile provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdateProfile(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRole provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdateRole(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRolePermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdateRolePermission(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) UpdateUser(_a0 context.Context, _a1 *commonproto.RequestPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.RequestPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateCasbinPolicy provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthHandler) ValidateCasbinPolicy(_a0 context.Context, _a1 *commonproto.CasbinPolicyPayload, _a2 *commonproto.Response) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *commonproto.CasbinPolicyPayload, *commonproto.Response) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
