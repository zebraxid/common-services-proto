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
	Product              string   `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Subject              string   `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Createdby            string   `protobuf:"bytes,6,opt,name=createdby,proto3" json:"createdby,omitempty"`
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

type InAppNotifRequestResponse struct {
	SendMessageResponse  *CrudResponse   `protobuf:"bytes,1,opt,name=sendMessageResponse,proto3" json:"sendMessageResponse,omitempty"`
	Messages             []*InAppMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *InAppNotifRequestResponse) Reset()         { *m = InAppNotifRequestResponse{} }
func (m *InAppNotifRequestResponse) String() string { return proto.CompactTextString(m) }
func (*InAppNotifRequestResponse) ProtoMessage()    {}
func (*InAppNotifRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f1929be0d77a46, []int{3}
}

func (m *InAppNotifRequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InAppNotifRequestResponse.Unmarshal(m, b)
}
func (m *InAppNotifRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InAppNotifRequestResponse.Marshal(b, m, deterministic)
}
func (m *InAppNotifRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InAppNotifRequestResponse.Merge(m, src)
}
func (m *InAppNotifRequestResponse) XXX_Size() int {
	return xxx_messageInfo_InAppNotifRequestResponse.Size(m)
}
func (m *InAppNotifRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InAppNotifRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InAppNotifRequestResponse proto.InternalMessageInfo

func (m *InAppNotifRequestResponse) GetSendMessageResponse() *CrudResponse {
	if m != nil {
		return m.SendMessageResponse
	}
	return nil
}

func (m *InAppNotifRequestResponse) GetMessages() []*InAppMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*InAppMessage)(nil), "commonproto.InAppMessage")
	proto.RegisterType((*SendMessageRequest)(nil), "commonproto.SendMessageRequest")
	proto.RegisterType((*GetMessagesRequest)(nil), "commonproto.GetMessagesRequest")
	proto.RegisterType((*InAppNotifRequestResponse)(nil), "commonproto.InAppNotifRequestResponse")
}

func init() { proto.RegisterFile("inapp.proto", fileDescriptor_c3f1929be0d77a46) }

var fileDescriptor_c3f1929be0d77a46 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x6f, 0xda, 0x7b, 0xdb, 0xdb, 0x13, 0x17, 0x3a, 0x82, 0x4c, 0x8b, 0x60, 0xc9, 0x42,
	0xdc, 0x98, 0x45, 0x8b, 0x0f, 0x60, 0x5d, 0x88, 0x88, 0x52, 0xe2, 0x46, 0xdc, 0x25, 0x93, 0x63,
	0x19, 0xa1, 0x33, 0xe3, 0xcc, 0x04, 0xd4, 0x97, 0xf1, 0x15, 0x7c, 0x0b, 0x5f, 0x4b, 0x32, 0x99,
	0xb6, 0x49, 0x5b, 0xc5, 0x55, 0x38, 0x7f, 0xbe, 0xdf, 0x7c, 0xe7, 0x23, 0x10, 0x72, 0x91, 0x2a,
	0x15, 0x2b, 0x2d, 0xad, 0x24, 0x21, 0x93, 0xf3, 0xb9, 0x14, 0xae, 0x18, 0x00, 0x93, 0x1a, 0xab,
	0x41, 0xf4, 0x11, 0xc0, 0xce, 0x95, 0x38, 0x57, 0xea, 0x06, 0x8d, 0x49, 0x67, 0x48, 0x28, 0x74,
	0x95, 0x96, 0x79, 0xc1, 0x2c, 0x0d, 0x86, 0xc1, 0x49, 0x2f, 0x59, 0x94, 0xe4, 0x00, 0x3a, 0x06,
	0x45, 0x8e, 0x9a, 0xb6, 0xdc, 0xc0, 0x57, 0xe4, 0x10, 0x7a, 0x1a, 0x19, 0x57, 0x1c, 0x85, 0xa5,
	0x6d, 0x37, 0x5a, 0x35, 0x4a, 0x9e, 0x29, 0xb2, 0x27, 0x64, 0x96, 0xfe, 0xad, 0x78, 0xbe, 0x2c,
	0x27, 0x4c, 0x0a, 0x5b, 0xaa, 0xfe, 0x55, 0x13, 0x5f, 0x96, 0x44, 0xa6, 0x31, 0xb5, 0x98, 0x67,
	0xaf, 0xb4, 0x53, 0x11, 0x97, 0x8d, 0x68, 0x06, 0xe4, 0x0e, 0x45, 0xee, 0x0d, 0x27, 0xf8, 0x5c,
	0xa0, 0xb1, 0x4d, 0x17, 0xc1, 0xb0, 0xdd, 0x74, 0x31, 0x86, 0xee, 0xbc, 0xda, 0x77, 0xe6, 0xc3,
	0x51, 0x3f, 0xae, 0x25, 0x12, 0xd7, 0x13, 0x48, 0x16, 0x9b, 0xd1, 0x08, 0xc8, 0x25, 0x5a, 0xdf,
	0x36, 0xdf, 0x3c, 0xd4, 0x3c, 0x37, 0x7a, 0x0f, 0xa0, 0xef, 0x68, 0xb7, 0xd2, 0xf2, 0x47, 0xaf,
	0x49, 0xd0, 0x28, 0x29, 0x0c, 0x92, 0x6b, 0xd8, 0x37, 0x75, 0xeb, 0x55, 0xdb, 0x51, 0xd6, 0x2d,
	0x5d, 0xe8, 0x22, 0x5f, 0x2c, 0x24, 0xdb, 0x54, 0xe4, 0x0c, 0xfe, 0x7b, 0xa7, 0x86, 0xb6, 0x86,
	0xed, 0x9f, 0x8f, 0x5a, 0xae, 0x8e, 0x3e, 0x03, 0xd8, 0x5b, 0x39, 0xe4, 0x2c, 0xb5, 0x5c, 0x0a,
	0x72, 0x0f, 0x61, 0x2d, 0x54, 0x72, 0xd4, 0x20, 0x6d, 0xc6, 0x3d, 0x38, 0xde, 0x7c, 0x6a, 0xdb,
	0xc5, 0xd1, 0x9f, 0x92, 0x5c, 0x4b, 0x71, 0x8d, 0xbc, 0x99, 0xef, 0xef, 0xc9, 0x93, 0x53, 0xd8,
	0xe5, 0x79, 0xfc, 0x86, 0x99, 0x4e, 0x5f, 0xbc, 0x68, 0x02, 0x4e, 0x30, 0x2d, 0xb5, 0xd3, 0xe0,
	0xa1, 0xfe, 0xdb, 0x67, 0x1d, 0xf7, 0x19, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x52, 0x8f,
	0x34, 0x19, 0x03, 0x00, 0x00,
}
