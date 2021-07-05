// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inapp.proto

package commonproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type InAppMessage struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Product              string   `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Sender               string   `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Subject              string   `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Content              string   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Createdby            string   `protobuf:"bytes,7,opt,name=createdby,proto3" json:"createdby,omitempty"`
	Createdat            string   `protobuf:"bytes,8,opt,name=createdat,proto3" json:"createdat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InAppMessage) Reset()         { *m = InAppMessage{} }
func (m *InAppMessage) String() string { return proto.CompactTextString(m) }
func (*InAppMessage) ProtoMessage()    {}
func (*InAppMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f1929be0d77a46, []int{0}
}

func (m *InAppMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InAppMessage.Unmarshal(m, b)
}
func (m *InAppMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InAppMessage.Marshal(b, m, deterministic)
}
func (m *InAppMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InAppMessage.Merge(m, src)
}
func (m *InAppMessage) XXX_Size() int {
	return xxx_messageInfo_InAppMessage.Size(m)
}
func (m *InAppMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InAppMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InAppMessage proto.InternalMessageInfo

func (m *InAppMessage) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *InAppMessage) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *InAppMessage) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *InAppMessage) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *InAppMessage) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *InAppMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *InAppMessage) GetCreatedby() string {
	if m != nil {
		return m.Createdby
	}
	return ""
}

func (m *InAppMessage) GetCreatedat() string {
	if m != nil {
		return m.Createdat
	}
	return ""
}

type SendMessageRequest struct {
	Recipient            []string      `protobuf:"bytes,1,rep,name=recipient,proto3" json:"recipient,omitempty"`
	Message              *InAppMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f1929be0d77a46, []int{1}
}

func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetRecipient() []string {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *SendMessageRequest) GetMessage() *InAppMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type GetMessagesRequest struct {
	Recipient            string   `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessagesRequest) Reset()         { *m = GetMessagesRequest{} }
func (m *GetMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*GetMessagesRequest) ProtoMessage()    {}
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f1929be0d77a46, []int{2}
}

func (m *GetMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessagesRequest.Unmarshal(m, b)
}
func (m *GetMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessagesRequest.Marshal(b, m, deterministic)
}
func (m *GetMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessagesRequest.Merge(m, src)
}
func (m *GetMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_GetMessagesRequest.Size(m)
}
func (m *GetMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessagesRequest proto.InternalMessageInfo

func (m *GetMessagesRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *GetMessagesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetMessagesRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type InAppNotifResponse struct {
	SendMessageResponse  *CrudResponse   `protobuf:"bytes,1,opt,name=sendMessageResponse,proto3" json:"sendMessageResponse,omitempty"`
	Messages             []*InAppMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *InAppNotifResponse) Reset()         { *m = InAppNotifResponse{} }
func (m *InAppNotifResponse) String() string { return proto.CompactTextString(m) }
func (*InAppNotifResponse) ProtoMessage()    {}
func (*InAppNotifResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f1929be0d77a46, []int{3}
}

func (m *InAppNotifResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InAppNotifResponse.Unmarshal(m, b)
}
func (m *InAppNotifResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InAppNotifResponse.Marshal(b, m, deterministic)
}
func (m *InAppNotifResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InAppNotifResponse.Merge(m, src)
}
func (m *InAppNotifResponse) XXX_Size() int {
	return xxx_messageInfo_InAppNotifResponse.Size(m)
}
func (m *InAppNotifResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InAppNotifResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InAppNotifResponse proto.InternalMessageInfo

func (m *InAppNotifResponse) GetSendMessageResponse() *CrudResponse {
	if m != nil {
		return m.SendMessageResponse
	}
	return nil
}

func (m *InAppNotifResponse) GetMessages() []*InAppMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*InAppMessage)(nil), "commonproto.InAppMessage")
	proto.RegisterType((*SendMessageRequest)(nil), "commonproto.SendMessageRequest")
	proto.RegisterType((*GetMessagesRequest)(nil), "commonproto.GetMessagesRequest")
	proto.RegisterType((*InAppNotifResponse)(nil), "commonproto.InAppNotifResponse")
}

func init() { proto.RegisterFile("inapp.proto", fileDescriptor_c3f1929be0d77a46) }

var fileDescriptor_c3f1929be0d77a46 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x26, 0x49, 0x7f, 0x27, 0x08, 0x81, 0x41, 0xc8, 0x54, 0x48, 0xad, 0x72, 0xe2, 0x42, 0x0e,
	0xad, 0x78, 0x00, 0x4a, 0x25, 0x54, 0x21, 0x50, 0x09, 0x37, 0xc4, 0x25, 0x71, 0x86, 0xca, 0x88,
	0xd8, 0xc6, 0x76, 0x24, 0xe0, 0x49, 0x78, 0x15, 0x5e, 0x66, 0x9f, 0x65, 0x15, 0x3b, 0x6d, 0x93,
	0x76, 0x77, 0xb5, 0xa7, 0x64, 0xe6, 0xfb, 0xe6, 0x9b, 0xf9, 0xbe, 0x28, 0x10, 0x73, 0x91, 0x2b,
	0x95, 0x2a, 0x2d, 0xad, 0x24, 0x31, 0x93, 0x55, 0x25, 0x85, 0x2b, 0x66, 0xc0, 0xa4, 0x46, 0x0f,
	0x24, 0x57, 0x01, 0x3c, 0xdc, 0x8a, 0xb7, 0x4a, 0x7d, 0x44, 0x63, 0xf2, 0x3d, 0x92, 0x47, 0x10,
	0x6e, 0x37, 0x34, 0x58, 0x04, 0xaf, 0xa6, 0x59, 0xb8, 0xdd, 0x10, 0x0a, 0x63, 0xa5, 0x65, 0x59,
	0x33, 0x4b, 0x43, 0xd7, 0x3c, 0x94, 0xe4, 0x39, 0x8c, 0x0c, 0x8a, 0x12, 0x35, 0x8d, 0x1c, 0xd0,
	0x56, 0xe4, 0x25, 0x4c, 0x35, 0x32, 0xae, 0x38, 0x0a, 0x4b, 0x07, 0x0e, 0x3a, 0x35, 0x1a, 0x3d,
	0x53, 0x17, 0x3f, 0x90, 0x59, 0x3a, 0xf4, 0x7a, 0x6d, 0xd9, 0x20, 0x4c, 0x0a, 0xdb, 0x4c, 0x8d,
	0x3c, 0xd2, 0x96, 0x8d, 0x22, 0xd3, 0x98, 0x5b, 0x2c, 0x8b, 0x3f, 0x74, 0xec, 0x15, 0x8f, 0x8d,
	0x0e, 0x9a, 0x5b, 0x3a, 0xe9, 0xa1, 0xb9, 0x4d, 0xf6, 0x40, 0xbe, 0xa0, 0x28, 0x5b, 0x7b, 0x19,
	0xfe, 0xaa, 0xd1, 0xd8, 0xfe, 0x8d, 0xc1, 0x22, 0xea, 0xdf, 0xb8, 0x82, 0x71, 0xe5, 0xf9, 0xce,
	0x73, 0xbc, 0x7c, 0x91, 0x76, 0xf2, 0x4b, 0xbb, 0x79, 0x65, 0x07, 0x66, 0xf2, 0x0d, 0xc8, 0x7b,
	0xb4, 0x6d, 0xdb, 0xdc, 0xb2, 0xe8, 0x2c, 0x8c, 0x67, 0x30, 0xfc, 0xc9, 0x2b, 0xee, 0xa3, 0x1d,
	0x66, 0xbe, 0x20, 0x04, 0x06, 0xaa, 0xd9, 0x1d, 0xb9, 0xa6, 0x7b, 0x4f, 0xfe, 0x05, 0x40, 0xdc,
	0xde, 0x4f, 0xd2, 0xf2, 0xef, 0x19, 0x1a, 0x25, 0x85, 0x41, 0xf2, 0x01, 0x9e, 0x9a, 0xae, 0x3b,
	0xdf, 0x76, 0x8b, 0xce, 0xaf, 0x7e, 0xa7, 0xeb, 0xf2, 0x40, 0xc8, 0x6e, 0x9a, 0x22, 0x6f, 0x60,
	0xd2, 0x9a, 0x31, 0x34, 0x5c, 0x44, 0x77, 0xfb, 0x3e, 0x52, 0x97, 0xff, 0x03, 0x78, 0x72, 0x3a,
	0x8d, 0xb3, 0xdc, 0x72, 0x29, 0xc8, 0x67, 0x88, 0x3b, 0xb9, 0x93, 0x79, 0x4f, 0xe9, 0xf2, 0x8b,
	0xcc, 0xe6, 0x97, 0xab, 0x7a, 0x56, 0x93, 0x07, 0x8d, 0x64, 0x27, 0xe1, 0x33, 0xc9, 0xcb, 0xec,
	0xef, 0x21, 0xb9, 0x7e, 0x0d, 0x8f, 0x79, 0x99, 0xfe, 0xc5, 0x42, 0xe7, 0xbf, 0x5b, 0xf6, 0x1a,
	0x1c, 0x73, 0xd7, 0x0c, 0xed, 0x82, 0xaf, 0xdd, 0x3f, 0xa7, 0x18, 0xb9, 0xc7, 0xea, 0x3a, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0xa8, 0x4f, 0x38, 0x5c, 0x03, 0x00, 0x00,
}
