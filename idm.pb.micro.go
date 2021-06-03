// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: idm.proto

package commonproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Auth service

func NewAuthEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Auth service

type AuthService interface {
	RequestToken(ctx context.Context, in *CoreRequest, opts ...client.CallOption) (*AuthResponse, error)
	RefreshToken(ctx context.Context, in *CoreRequest, opts ...client.CallOption) (*AuthResponse, error)
	AuthorizeToken(ctx context.Context, in *CoreRequest, opts ...client.CallOption) (*AuthResponse, error)
	LogoutToken(ctx context.Context, in *CoreRequest, opts ...client.CallOption) (*AuthResponse, error)
	// product
	GetProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	ListProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	CreateProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error)
	DeleteProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error)
	// product organization
	ListProductOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	CreateProductOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error)
	UpdateProductOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error)
	DeleteProductOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error)
	// user
	ListUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	CreateUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	DeleteUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	// role
	ListRole(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	CreateRole(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	UpdateRole(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	DeleteRole(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	// organization
	ListOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	CreateOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	UpdateOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	DeleteOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	// permission
	ListPermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	CreatePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	UpdatePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	DeletePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	// role permission
	ListRolePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	CreateRolePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	UpdateRolePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	DeleteRolePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	// change and reset password
	ChangePassword(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	ResetPassword(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	// one time password request
	OneTimePassword(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	//profile
	GetProfile(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	UpdateProfile(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	// ABAC (CasbinPolicy)
	AddCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, opts ...client.CallOption) (*Response, error)
	RemoveCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, opts ...client.CallOption) (*Response, error)
	ValidateCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, opts ...client.CallOption) (*Response, error)
	// ABAC (AttributePolicy)
	ListAttributePolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	CreateAttributePolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	UpdateAttributePolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	DeleteAttributePolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error)
	ListCasbinPolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
	// DMAA List User
	DMAAListUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) RequestToken(ctx context.Context, in *CoreRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.RequestToken", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) RefreshToken(ctx context.Context, in *CoreRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.RefreshToken", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) AuthorizeToken(ctx context.Context, in *CoreRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.AuthorizeToken", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) LogoutToken(ctx context.Context, in *CoreRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.LogoutToken", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.GetProduct", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListProduct", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateProduct", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateProduct", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteProduct(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteProduct", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListProductOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListProductOrganization", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateProductOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateProductOrganization", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateProductOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateProductOrganization", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteProductOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteProductOrganization", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListUser", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListRole(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListRole", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateRole(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateRole(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteRole(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListOrganization", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateOrganization", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateOrganization", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteOrganization(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteOrganization", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListPermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListPermission", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreatePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.CreatePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdatePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdatePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeletePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.DeletePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListRolePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListRolePermission", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateRolePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateRolePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateRolePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateRolePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteRolePermission(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteRolePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ChangePassword(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.ChangePassword", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ResetPassword(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.ResetPassword", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) OneTimePassword(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.OneTimePassword", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetProfile(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.GetProfile", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateProfile(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateProfile", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) AddCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.AddCasbinPolicy", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) RemoveCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.RemoveCasbinPolicy", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ValidateCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.ValidateCasbinPolicy", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListAttributePolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListAttributePolicy", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateAttributePolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateAttributePolicy", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UpdateAttributePolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.UpdateAttributePolicy", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeleteAttributePolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.DeleteAttributePolicy", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ListCasbinPolicy(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ListCasbinPolicy", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DMAAListUser(ctx context.Context, in *RequestPayload, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.DMAAListUser", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	RequestToken(context.Context, *CoreRequest, *AuthResponse) error
	RefreshToken(context.Context, *CoreRequest, *AuthResponse) error
	AuthorizeToken(context.Context, *CoreRequest, *AuthResponse) error
	LogoutToken(context.Context, *CoreRequest, *AuthResponse) error
	// product
	GetProduct(context.Context, *RequestPayload, *ListResponse) error
	ListProduct(context.Context, *RequestPayload, *ListResponse) error
	CreateProduct(context.Context, *RequestPayload, *ProductResponse) error
	UpdateProduct(context.Context, *RequestPayload, *ProductResponse) error
	DeleteProduct(context.Context, *RequestPayload, *ProductResponse) error
	// product organization
	ListProductOrganization(context.Context, *RequestPayload, *ListResponse) error
	CreateProductOrganization(context.Context, *RequestPayload, *ProductResponse) error
	UpdateProductOrganization(context.Context, *RequestPayload, *ProductResponse) error
	DeleteProductOrganization(context.Context, *RequestPayload, *ProductResponse) error
	// user
	ListUser(context.Context, *RequestPayload, *ListResponse) error
	CreateUser(context.Context, *RequestPayload, *Response) error
	UpdateUser(context.Context, *RequestPayload, *Response) error
	DeleteUser(context.Context, *RequestPayload, *Response) error
	// role
	ListRole(context.Context, *RequestPayload, *ListResponse) error
	CreateRole(context.Context, *RequestPayload, *Response) error
	UpdateRole(context.Context, *RequestPayload, *Response) error
	DeleteRole(context.Context, *RequestPayload, *Response) error
	// organization
	ListOrganization(context.Context, *RequestPayload, *ListResponse) error
	CreateOrganization(context.Context, *RequestPayload, *Response) error
	UpdateOrganization(context.Context, *RequestPayload, *Response) error
	DeleteOrganization(context.Context, *RequestPayload, *Response) error
	// permission
	ListPermission(context.Context, *RequestPayload, *ListResponse) error
	CreatePermission(context.Context, *RequestPayload, *Response) error
	UpdatePermission(context.Context, *RequestPayload, *Response) error
	DeletePermission(context.Context, *RequestPayload, *Response) error
	// role permission
	ListRolePermission(context.Context, *RequestPayload, *ListResponse) error
	CreateRolePermission(context.Context, *RequestPayload, *Response) error
	UpdateRolePermission(context.Context, *RequestPayload, *Response) error
	DeleteRolePermission(context.Context, *RequestPayload, *Response) error
	// change and reset password
	ChangePassword(context.Context, *RequestPayload, *Response) error
	ResetPassword(context.Context, *RequestPayload, *Response) error
	// one time password request
	OneTimePassword(context.Context, *RequestPayload, *Response) error
	//profile
	GetProfile(context.Context, *RequestPayload, *ListResponse) error
	UpdateProfile(context.Context, *RequestPayload, *Response) error
	// ABAC (CasbinPolicy)
	AddCasbinPolicy(context.Context, *CasbinPolicyPayload, *Response) error
	RemoveCasbinPolicy(context.Context, *CasbinPolicyPayload, *Response) error
	ValidateCasbinPolicy(context.Context, *CasbinPolicyPayload, *Response) error
	// ABAC (AttributePolicy)
	ListAttributePolicy(context.Context, *RequestPayload, *ListResponse) error
	CreateAttributePolicy(context.Context, *RequestPayload, *Response) error
	UpdateAttributePolicy(context.Context, *RequestPayload, *Response) error
	DeleteAttributePolicy(context.Context, *RequestPayload, *Response) error
	ListCasbinPolicy(context.Context, *RequestPayload, *ListResponse) error
	// DMAA List User
	DMAAListUser(context.Context, *RequestPayload, *ListResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		RequestToken(ctx context.Context, in *CoreRequest, out *AuthResponse) error
		RefreshToken(ctx context.Context, in *CoreRequest, out *AuthResponse) error
		AuthorizeToken(ctx context.Context, in *CoreRequest, out *AuthResponse) error
		LogoutToken(ctx context.Context, in *CoreRequest, out *AuthResponse) error
		GetProduct(ctx context.Context, in *RequestPayload, out *ListResponse) error
		ListProduct(ctx context.Context, in *RequestPayload, out *ListResponse) error
		CreateProduct(ctx context.Context, in *RequestPayload, out *ProductResponse) error
		UpdateProduct(ctx context.Context, in *RequestPayload, out *ProductResponse) error
		DeleteProduct(ctx context.Context, in *RequestPayload, out *ProductResponse) error
		ListProductOrganization(ctx context.Context, in *RequestPayload, out *ListResponse) error
		CreateProductOrganization(ctx context.Context, in *RequestPayload, out *ProductResponse) error
		UpdateProductOrganization(ctx context.Context, in *RequestPayload, out *ProductResponse) error
		DeleteProductOrganization(ctx context.Context, in *RequestPayload, out *ProductResponse) error
		ListUser(ctx context.Context, in *RequestPayload, out *ListResponse) error
		CreateUser(ctx context.Context, in *RequestPayload, out *Response) error
		UpdateUser(ctx context.Context, in *RequestPayload, out *Response) error
		DeleteUser(ctx context.Context, in *RequestPayload, out *Response) error
		ListRole(ctx context.Context, in *RequestPayload, out *ListResponse) error
		CreateRole(ctx context.Context, in *RequestPayload, out *Response) error
		UpdateRole(ctx context.Context, in *RequestPayload, out *Response) error
		DeleteRole(ctx context.Context, in *RequestPayload, out *Response) error
		ListOrganization(ctx context.Context, in *RequestPayload, out *ListResponse) error
		CreateOrganization(ctx context.Context, in *RequestPayload, out *Response) error
		UpdateOrganization(ctx context.Context, in *RequestPayload, out *Response) error
		DeleteOrganization(ctx context.Context, in *RequestPayload, out *Response) error
		ListPermission(ctx context.Context, in *RequestPayload, out *ListResponse) error
		CreatePermission(ctx context.Context, in *RequestPayload, out *Response) error
		UpdatePermission(ctx context.Context, in *RequestPayload, out *Response) error
		DeletePermission(ctx context.Context, in *RequestPayload, out *Response) error
		ListRolePermission(ctx context.Context, in *RequestPayload, out *ListResponse) error
		CreateRolePermission(ctx context.Context, in *RequestPayload, out *Response) error
		UpdateRolePermission(ctx context.Context, in *RequestPayload, out *Response) error
		DeleteRolePermission(ctx context.Context, in *RequestPayload, out *Response) error
		ChangePassword(ctx context.Context, in *RequestPayload, out *Response) error
		ResetPassword(ctx context.Context, in *RequestPayload, out *Response) error
		OneTimePassword(ctx context.Context, in *RequestPayload, out *Response) error
		GetProfile(ctx context.Context, in *RequestPayload, out *ListResponse) error
		UpdateProfile(ctx context.Context, in *RequestPayload, out *Response) error
		AddCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, out *Response) error
		RemoveCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, out *Response) error
		ValidateCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, out *Response) error
		ListAttributePolicy(ctx context.Context, in *RequestPayload, out *ListResponse) error
		CreateAttributePolicy(ctx context.Context, in *RequestPayload, out *Response) error
		UpdateAttributePolicy(ctx context.Context, in *RequestPayload, out *Response) error
		DeleteAttributePolicy(ctx context.Context, in *RequestPayload, out *Response) error
		ListCasbinPolicy(ctx context.Context, in *RequestPayload, out *ListResponse) error
		DMAAListUser(ctx context.Context, in *RequestPayload, out *ListResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) RequestToken(ctx context.Context, in *CoreRequest, out *AuthResponse) error {
	return h.AuthHandler.RequestToken(ctx, in, out)
}

func (h *authHandler) RefreshToken(ctx context.Context, in *CoreRequest, out *AuthResponse) error {
	return h.AuthHandler.RefreshToken(ctx, in, out)
}

func (h *authHandler) AuthorizeToken(ctx context.Context, in *CoreRequest, out *AuthResponse) error {
	return h.AuthHandler.AuthorizeToken(ctx, in, out)
}

func (h *authHandler) LogoutToken(ctx context.Context, in *CoreRequest, out *AuthResponse) error {
	return h.AuthHandler.LogoutToken(ctx, in, out)
}

func (h *authHandler) GetProduct(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.GetProduct(ctx, in, out)
}

func (h *authHandler) ListProduct(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListProduct(ctx, in, out)
}

func (h *authHandler) CreateProduct(ctx context.Context, in *RequestPayload, out *ProductResponse) error {
	return h.AuthHandler.CreateProduct(ctx, in, out)
}

func (h *authHandler) UpdateProduct(ctx context.Context, in *RequestPayload, out *ProductResponse) error {
	return h.AuthHandler.UpdateProduct(ctx, in, out)
}

func (h *authHandler) DeleteProduct(ctx context.Context, in *RequestPayload, out *ProductResponse) error {
	return h.AuthHandler.DeleteProduct(ctx, in, out)
}

func (h *authHandler) ListProductOrganization(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListProductOrganization(ctx, in, out)
}

func (h *authHandler) CreateProductOrganization(ctx context.Context, in *RequestPayload, out *ProductResponse) error {
	return h.AuthHandler.CreateProductOrganization(ctx, in, out)
}

func (h *authHandler) UpdateProductOrganization(ctx context.Context, in *RequestPayload, out *ProductResponse) error {
	return h.AuthHandler.UpdateProductOrganization(ctx, in, out)
}

func (h *authHandler) DeleteProductOrganization(ctx context.Context, in *RequestPayload, out *ProductResponse) error {
	return h.AuthHandler.DeleteProductOrganization(ctx, in, out)
}

func (h *authHandler) ListUser(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListUser(ctx, in, out)
}

func (h *authHandler) CreateUser(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.CreateUser(ctx, in, out)
}

func (h *authHandler) UpdateUser(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.UpdateUser(ctx, in, out)
}

func (h *authHandler) DeleteUser(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.DeleteUser(ctx, in, out)
}

func (h *authHandler) ListRole(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListRole(ctx, in, out)
}

func (h *authHandler) CreateRole(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.CreateRole(ctx, in, out)
}

func (h *authHandler) UpdateRole(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.UpdateRole(ctx, in, out)
}

func (h *authHandler) DeleteRole(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.DeleteRole(ctx, in, out)
}

func (h *authHandler) ListOrganization(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListOrganization(ctx, in, out)
}

func (h *authHandler) CreateOrganization(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.CreateOrganization(ctx, in, out)
}

func (h *authHandler) UpdateOrganization(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.UpdateOrganization(ctx, in, out)
}

func (h *authHandler) DeleteOrganization(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.DeleteOrganization(ctx, in, out)
}

func (h *authHandler) ListPermission(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListPermission(ctx, in, out)
}

func (h *authHandler) CreatePermission(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.CreatePermission(ctx, in, out)
}

func (h *authHandler) UpdatePermission(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.UpdatePermission(ctx, in, out)
}

func (h *authHandler) DeletePermission(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.DeletePermission(ctx, in, out)
}

func (h *authHandler) ListRolePermission(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListRolePermission(ctx, in, out)
}

func (h *authHandler) CreateRolePermission(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.CreateRolePermission(ctx, in, out)
}

func (h *authHandler) UpdateRolePermission(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.UpdateRolePermission(ctx, in, out)
}

func (h *authHandler) DeleteRolePermission(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.DeleteRolePermission(ctx, in, out)
}

func (h *authHandler) ChangePassword(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.ChangePassword(ctx, in, out)
}

func (h *authHandler) ResetPassword(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.ResetPassword(ctx, in, out)
}

func (h *authHandler) OneTimePassword(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.OneTimePassword(ctx, in, out)
}

func (h *authHandler) GetProfile(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.GetProfile(ctx, in, out)
}

func (h *authHandler) UpdateProfile(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.UpdateProfile(ctx, in, out)
}

func (h *authHandler) AddCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, out *Response) error {
	return h.AuthHandler.AddCasbinPolicy(ctx, in, out)
}

func (h *authHandler) RemoveCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, out *Response) error {
	return h.AuthHandler.RemoveCasbinPolicy(ctx, in, out)
}

func (h *authHandler) ValidateCasbinPolicy(ctx context.Context, in *CasbinPolicyPayload, out *Response) error {
	return h.AuthHandler.ValidateCasbinPolicy(ctx, in, out)
}

func (h *authHandler) ListAttributePolicy(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListAttributePolicy(ctx, in, out)
}

func (h *authHandler) CreateAttributePolicy(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.CreateAttributePolicy(ctx, in, out)
}

func (h *authHandler) UpdateAttributePolicy(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.UpdateAttributePolicy(ctx, in, out)
}

func (h *authHandler) DeleteAttributePolicy(ctx context.Context, in *RequestPayload, out *Response) error {
	return h.AuthHandler.DeleteAttributePolicy(ctx, in, out)
}

func (h *authHandler) ListCasbinPolicy(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.ListCasbinPolicy(ctx, in, out)
}

func (h *authHandler) DMAAListUser(ctx context.Context, in *RequestPayload, out *ListResponse) error {
	return h.AuthHandler.DMAAListUser(ctx, in, out)
}