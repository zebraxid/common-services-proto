// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: audit.proto

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

// Api Endpoints for AuditTrail service

func NewAuditTrailEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuditTrail service

type AuditTrailService interface {
	SendLog(ctx context.Context, in *SendLogRequest, opts ...client.CallOption) (*AuditResponse, error)
	ReadLog(ctx context.Context, in *ReadLogRequest, opts ...client.CallOption) (*AuditResponse, error)
}

type auditTrailService struct {
	c    client.Client
	name string
}

func NewAuditTrailService(name string, c client.Client) AuditTrailService {
	return &auditTrailService{
		c:    c,
		name: name,
	}
}

func (c *auditTrailService) SendLog(ctx context.Context, in *SendLogRequest, opts ...client.CallOption) (*AuditResponse, error) {
	req := c.c.NewRequest(c.name, "AuditTrail.SendLog", in)
	out := new(AuditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditTrailService) ReadLog(ctx context.Context, in *ReadLogRequest, opts ...client.CallOption) (*AuditResponse, error) {
	req := c.c.NewRequest(c.name, "AuditTrail.ReadLog", in)
	out := new(AuditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuditTrail service

type AuditTrailHandler interface {
	SendLog(context.Context, *SendLogRequest, *AuditResponse) error
	ReadLog(context.Context, *ReadLogRequest, *AuditResponse) error
}

func RegisterAuditTrailHandler(s server.Server, hdlr AuditTrailHandler, opts ...server.HandlerOption) error {
	type auditTrail interface {
		SendLog(ctx context.Context, in *SendLogRequest, out *AuditResponse) error
		ReadLog(ctx context.Context, in *ReadLogRequest, out *AuditResponse) error
	}
	type AuditTrail struct {
		auditTrail
	}
	h := &auditTrailHandler{hdlr}
	return s.Handle(s.NewHandler(&AuditTrail{h}, opts...))
}

type auditTrailHandler struct {
	AuditTrailHandler
}

func (h *auditTrailHandler) SendLog(ctx context.Context, in *SendLogRequest, out *AuditResponse) error {
	return h.AuditTrailHandler.SendLog(ctx, in, out)
}

func (h *auditTrailHandler) ReadLog(ctx context.Context, in *ReadLogRequest, out *AuditResponse) error {
	return h.AuditTrailHandler.ReadLog(ctx, in, out)
}
