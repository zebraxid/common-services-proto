// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: inapp.proto

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

// Api Endpoints for InAppNotification service

func NewInAppNotificationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for InAppNotification service

type InAppNotificationService interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...client.CallOption) (*InAppNotifResponse, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...client.CallOption) (*InAppNotifResponse, error)
	ReadMessage(ctx context.Context, in *ReadMessageRequest, opts ...client.CallOption) (*InAppNotifResponse, error)
}

type inAppNotificationService struct {
	c    client.Client
	name string
}

func NewInAppNotificationService(name string, c client.Client) InAppNotificationService {
	return &inAppNotificationService{
		c:    c,
		name: name,
	}
}

func (c *inAppNotificationService) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...client.CallOption) (*InAppNotifResponse, error) {
	req := c.c.NewRequest(c.name, "InAppNotification.SendMessage", in)
	out := new(InAppNotifResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inAppNotificationService) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...client.CallOption) (*InAppNotifResponse, error) {
	req := c.c.NewRequest(c.name, "InAppNotification.GetMessages", in)
	out := new(InAppNotifResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inAppNotificationService) ReadMessage(ctx context.Context, in *ReadMessageRequest, opts ...client.CallOption) (*InAppNotifResponse, error) {
	req := c.c.NewRequest(c.name, "InAppNotification.ReadMessage", in)
	out := new(InAppNotifResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InAppNotification service

type InAppNotificationHandler interface {
	SendMessage(context.Context, *SendMessageRequest, *InAppNotifResponse) error
	GetMessages(context.Context, *GetMessagesRequest, *InAppNotifResponse) error
	ReadMessage(context.Context, *ReadMessageRequest, *InAppNotifResponse) error
}

func RegisterInAppNotificationHandler(s server.Server, hdlr InAppNotificationHandler, opts ...server.HandlerOption) error {
	type inAppNotification interface {
		SendMessage(ctx context.Context, in *SendMessageRequest, out *InAppNotifResponse) error
		GetMessages(ctx context.Context, in *GetMessagesRequest, out *InAppNotifResponse) error
		ReadMessage(ctx context.Context, in *ReadMessageRequest, out *InAppNotifResponse) error
	}
	type InAppNotification struct {
		inAppNotification
	}
	h := &inAppNotificationHandler{hdlr}
	return s.Handle(s.NewHandler(&InAppNotification{h}, opts...))
}

type inAppNotificationHandler struct {
	InAppNotificationHandler
}

func (h *inAppNotificationHandler) SendMessage(ctx context.Context, in *SendMessageRequest, out *InAppNotifResponse) error {
	return h.InAppNotificationHandler.SendMessage(ctx, in, out)
}

func (h *inAppNotificationHandler) GetMessages(ctx context.Context, in *GetMessagesRequest, out *InAppNotifResponse) error {
	return h.InAppNotificationHandler.GetMessages(ctx, in, out)
}

func (h *inAppNotificationHandler) ReadMessage(ctx context.Context, in *ReadMessageRequest, out *InAppNotifResponse) error {
	return h.InAppNotificationHandler.ReadMessage(ctx, in, out)
}