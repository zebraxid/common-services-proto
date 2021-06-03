// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: notification.proto

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

// Api Endpoints for Notification service

func NewNotificationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Notification service

type NotificationService interface {
	// create
	CreateNotificationGroup(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error)
	CreateNotificationGroupMember(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error)
	SendNotification(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error)
	// update
	UpdateNotificationPermission(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error)
	// delete
	DeleteNotificationGroup(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error)
	DeleteNotificationGroupMember(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error)
	// list
	ListNotificationPermission(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationListResponse, error)
	ListNotificationLog(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationListResponse, error)
	ListNotificationGroup(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationListResponse, error)
	ListNotificationGroupMember(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationListResponse, error)
}

type notificationService struct {
	c    client.Client
	name string
}

func NewNotificationService(name string, c client.Client) NotificationService {
	return &notificationService{
		c:    c,
		name: name,
	}
}

func (c *notificationService) CreateNotificationGroup(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.CreateNotificationGroup", in)
	out := new(NotificationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) CreateNotificationGroupMember(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.CreateNotificationGroupMember", in)
	out := new(NotificationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) SendNotification(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.SendNotification", in)
	out := new(NotificationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) UpdateNotificationPermission(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.UpdateNotificationPermission", in)
	out := new(NotificationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) DeleteNotificationGroup(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.DeleteNotificationGroup", in)
	out := new(NotificationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) DeleteNotificationGroupMember(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.DeleteNotificationGroupMember", in)
	out := new(NotificationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) ListNotificationPermission(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationListResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.ListNotificationPermission", in)
	out := new(NotificationListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) ListNotificationLog(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationListResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.ListNotificationLog", in)
	out := new(NotificationListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) ListNotificationGroup(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationListResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.ListNotificationGroup", in)
	out := new(NotificationListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) ListNotificationGroupMember(ctx context.Context, in *NotificationRequest, opts ...client.CallOption) (*NotificationListResponse, error) {
	req := c.c.NewRequest(c.name, "Notification.ListNotificationGroupMember", in)
	out := new(NotificationListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notification service

type NotificationHandler interface {
	// create
	CreateNotificationGroup(context.Context, *NotificationRequest, *NotificationResponse) error
	CreateNotificationGroupMember(context.Context, *NotificationRequest, *NotificationResponse) error
	SendNotification(context.Context, *NotificationRequest, *NotificationResponse) error
	// update
	UpdateNotificationPermission(context.Context, *NotificationRequest, *NotificationResponse) error
	// delete
	DeleteNotificationGroup(context.Context, *NotificationRequest, *NotificationResponse) error
	DeleteNotificationGroupMember(context.Context, *NotificationRequest, *NotificationResponse) error
	// list
	ListNotificationPermission(context.Context, *NotificationRequest, *NotificationListResponse) error
	ListNotificationLog(context.Context, *NotificationRequest, *NotificationListResponse) error
	ListNotificationGroup(context.Context, *NotificationRequest, *NotificationListResponse) error
	ListNotificationGroupMember(context.Context, *NotificationRequest, *NotificationListResponse) error
}

func RegisterNotificationHandler(s server.Server, hdlr NotificationHandler, opts ...server.HandlerOption) error {
	type notification interface {
		CreateNotificationGroup(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error
		CreateNotificationGroupMember(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error
		SendNotification(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error
		UpdateNotificationPermission(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error
		DeleteNotificationGroup(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error
		DeleteNotificationGroupMember(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error
		ListNotificationPermission(ctx context.Context, in *NotificationRequest, out *NotificationListResponse) error
		ListNotificationLog(ctx context.Context, in *NotificationRequest, out *NotificationListResponse) error
		ListNotificationGroup(ctx context.Context, in *NotificationRequest, out *NotificationListResponse) error
		ListNotificationGroupMember(ctx context.Context, in *NotificationRequest, out *NotificationListResponse) error
	}
	type Notification struct {
		notification
	}
	h := &notificationHandler{hdlr}
	return s.Handle(s.NewHandler(&Notification{h}, opts...))
}

type notificationHandler struct {
	NotificationHandler
}

func (h *notificationHandler) CreateNotificationGroup(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error {
	return h.NotificationHandler.CreateNotificationGroup(ctx, in, out)
}

func (h *notificationHandler) CreateNotificationGroupMember(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error {
	return h.NotificationHandler.CreateNotificationGroupMember(ctx, in, out)
}

func (h *notificationHandler) SendNotification(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error {
	return h.NotificationHandler.SendNotification(ctx, in, out)
}

func (h *notificationHandler) UpdateNotificationPermission(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error {
	return h.NotificationHandler.UpdateNotificationPermission(ctx, in, out)
}

func (h *notificationHandler) DeleteNotificationGroup(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error {
	return h.NotificationHandler.DeleteNotificationGroup(ctx, in, out)
}

func (h *notificationHandler) DeleteNotificationGroupMember(ctx context.Context, in *NotificationRequest, out *NotificationResponse) error {
	return h.NotificationHandler.DeleteNotificationGroupMember(ctx, in, out)
}

func (h *notificationHandler) ListNotificationPermission(ctx context.Context, in *NotificationRequest, out *NotificationListResponse) error {
	return h.NotificationHandler.ListNotificationPermission(ctx, in, out)
}

func (h *notificationHandler) ListNotificationLog(ctx context.Context, in *NotificationRequest, out *NotificationListResponse) error {
	return h.NotificationHandler.ListNotificationLog(ctx, in, out)
}

func (h *notificationHandler) ListNotificationGroup(ctx context.Context, in *NotificationRequest, out *NotificationListResponse) error {
	return h.NotificationHandler.ListNotificationGroup(ctx, in, out)
}

func (h *notificationHandler) ListNotificationGroupMember(ctx context.Context, in *NotificationRequest, out *NotificationListResponse) error {
	return h.NotificationHandler.ListNotificationGroupMember(ctx, in, out)
}