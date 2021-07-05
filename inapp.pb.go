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
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0x49, 0x3f, 0xe8, 0x04, 0x21, 0x30, 0x08, 0x99, 0x0a, 0xa9, 0x55, 0x4e, 0x5c, 0xc8,
	0xa1, 0x15, 0x3f, 0x80, 0x52, 0x09, 0x55, 0x08, 0x54, 0xc2, 0x0d, 0x71, 0x49, 0x9c, 0xa1, 0x32,
	0x22, 0xb6, 0xb1, 0x1d, 0x09, 0xf6, 0x97, 0xec, 0x5f, 0xd9, 0xdb, 0xfe, 0xb4, 0x55, 0xec, 0xb4,
	0x4d, 0xda, 0xdd, 0xd5, 0x9e, 0x92, 0x79, 0xf3, 0xe6, 0xcd, 0xbc, 0x17, 0x05, 0x62, 0x2e, 0x72,
	0xa5, 0x52, 0xa5, 0xa5, 0x95, 0x24, 0x66, 0xb2, 0xaa, 0xa4, 0x70, 0xc5, 0x14, 0x98, 0xd4, 0xe8,
	0x1b, 0xc9, 0x75, 0x00, 0x4f, 0x36, 0xe2, 0x83, 0x52, 0x5f, 0xd0, 0x98, 0x7c, 0x87, 0xe4, 0x29,
	0x84, 0x9b, 0x35, 0x0d, 0xe6, 0xc1, 0xdb, 0x49, 0x16, 0x6e, 0xd6, 0x84, 0xc2, 0x58, 0x69, 0x59,
	0xd6, 0xcc, 0xd2, 0xd0, 0x81, 0xfb, 0x92, 0xbc, 0x82, 0x91, 0x41, 0x51, 0xa2, 0xa6, 0x91, 0x6b,
	0xb4, 0x15, 0x79, 0x03, 0x13, 0x8d, 0x8c, 0x2b, 0x8e, 0xc2, 0xd2, 0x81, 0x6b, 0x1d, 0x81, 0x46,
	0xcf, 0xd4, 0xc5, 0x6f, 0x64, 0x96, 0x0e, 0xbd, 0x5e, 0x5b, 0x36, 0x1d, 0x26, 0x85, 0x6d, 0xa6,
	0x46, 0xbe, 0xd3, 0x96, 0x8d, 0x22, 0xd3, 0x98, 0x5b, 0x2c, 0x8b, 0xff, 0x74, 0xec, 0x15, 0x0f,
	0x40, 0xb2, 0x03, 0xf2, 0x1d, 0x45, 0xd9, 0x1a, 0xc8, 0xf0, 0x6f, 0x8d, 0xc6, 0xf6, 0xaf, 0x08,
	0xe6, 0x51, 0xff, 0x8a, 0x25, 0x8c, 0x2b, 0xcf, 0x77, 0xae, 0xe2, 0xc5, 0xeb, 0xb4, 0x93, 0x50,
	0xda, 0x4d, 0x24, 0xdb, 0x33, 0x93, 0x9f, 0x40, 0x3e, 0xa1, 0x6d, 0x61, 0x73, 0xc7, 0xa2, 0x13,
	0xbb, 0x2f, 0x61, 0xf8, 0x87, 0x57, 0xdc, 0x87, 0x37, 0xcc, 0x7c, 0x41, 0x08, 0x0c, 0x54, 0xb3,
	0x3b, 0x72, 0xa0, 0x7b, 0x4f, 0x2e, 0x03, 0x20, 0x6e, 0xef, 0x57, 0x69, 0xf9, 0xaf, 0x0c, 0x8d,
	0x92, 0xc2, 0x20, 0xf9, 0x0c, 0x2f, 0x4c, 0xd7, 0x9d, 0x87, 0xdd, 0xa2, 0xd3, 0xab, 0x3f, 0xea,
	0xba, 0xdc, 0x13, 0xb2, 0xdb, 0xa6, 0xc8, 0x7b, 0x78, 0xdc, 0x9a, 0x31, 0x34, 0x9c, 0x47, 0xf7,
	0xfb, 0x3e, 0x50, 0x17, 0x57, 0x01, 0x3c, 0x3f, 0x9e, 0xc6, 0x59, 0x6e, 0xb9, 0x14, 0xe4, 0x1b,
	0xc4, 0x9d, 0xdc, 0xc9, 0xac, 0xa7, 0x74, 0xfe, 0x45, 0xa6, 0xb3, 0xf3, 0x55, 0x3d, 0xab, 0xc9,
	0xa3, 0x46, 0xb2, 0x93, 0xf0, 0x89, 0xe4, 0x79, 0xf6, 0x0f, 0x90, 0x5c, 0xbd, 0x83, 0x67, 0xbc,
	0x4c, 0x2f, 0xb0, 0xd0, 0xf9, 0xbf, 0x96, 0xbd, 0x02, 0xc7, 0xdc, 0x36, 0x43, 0xdb, 0xe0, 0x47,
	0xf7, 0xdf, 0x28, 0x46, 0xee, 0xb1, 0xbc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xcf, 0x1f, 0x8b,
	0x3e, 0x03, 0x00, 0x00,
}
