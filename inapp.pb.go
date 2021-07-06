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
	Read                 bool     `protobuf:"varint,7,opt,name=read,proto3" json:"read,omitempty"`
	Createdby            string   `protobuf:"bytes,8,opt,name=createdby,proto3" json:"createdby,omitempty"`
	Createdat            string   `protobuf:"bytes,9,opt,name=createdat,proto3" json:"createdat,omitempty"`
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

func (m *InAppMessage) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
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

type ReadMessageRequest struct {
	ReadMessage          []*ReadMessage `protobuf:"bytes,1,rep,name=readMessage,proto3" json:"readMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReadMessageRequest) Reset()         { *m = ReadMessageRequest{} }
func (m *ReadMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ReadMessageRequest) ProtoMessage()    {}
func (*ReadMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f1929be0d77a46, []int{4}
}

func (m *ReadMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMessageRequest.Unmarshal(m, b)
}
func (m *ReadMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMessageRequest.Marshal(b, m, deterministic)
}
func (m *ReadMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMessageRequest.Merge(m, src)
}
func (m *ReadMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ReadMessageRequest.Size(m)
}
func (m *ReadMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMessageRequest proto.InternalMessageInfo

func (m *ReadMessageRequest) GetReadMessage() []*ReadMessage {
	if m != nil {
		return m.ReadMessage
	}
	return nil
}

type ReadMessage struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Read                 bool     `protobuf:"varint,2,opt,name=read,proto3" json:"read,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMessage) Reset()         { *m = ReadMessage{} }
func (m *ReadMessage) String() string { return proto.CompactTextString(m) }
func (*ReadMessage) ProtoMessage()    {}
func (*ReadMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f1929be0d77a46, []int{5}
}

func (m *ReadMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMessage.Unmarshal(m, b)
}
func (m *ReadMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMessage.Marshal(b, m, deterministic)
}
func (m *ReadMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMessage.Merge(m, src)
}
func (m *ReadMessage) XXX_Size() int {
	return xxx_messageInfo_ReadMessage.Size(m)
}
func (m *ReadMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMessage proto.InternalMessageInfo

func (m *ReadMessage) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ReadMessage) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func init() {
	proto.RegisterType((*InAppMessage)(nil), "commonproto.InAppMessage")
	proto.RegisterType((*SendMessageRequest)(nil), "commonproto.SendMessageRequest")
	proto.RegisterType((*GetMessagesRequest)(nil), "commonproto.GetMessagesRequest")
	proto.RegisterType((*InAppNotifResponse)(nil), "commonproto.InAppNotifResponse")
	proto.RegisterType((*ReadMessageRequest)(nil), "commonproto.ReadMessageRequest")
	proto.RegisterType((*ReadMessage)(nil), "commonproto.ReadMessage")
}

func init() { proto.RegisterFile("inapp.proto", fileDescriptor_c3f1929be0d77a46) }

var fileDescriptor_c3f1929be0d77a46 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x76, 0xa6, 0xdb, 0xbf, 0x33, 0x22, 0x1a, 0x45, 0xe2, 0x22, 0xb4, 0xcc, 0xd5, 0xde, 0x58,
	0xb0, 0x8b, 0x37, 0xde, 0xb9, 0x2e, 0x48, 0x11, 0xa5, 0xc6, 0x3b, 0xf1, 0x26, 0xcd, 0x1c, 0x97,
	0x88, 0x4d, 0x62, 0x92, 0x82, 0xfa, 0x0c, 0x3e, 0x80, 0xcf, 0xe9, 0x13, 0x48, 0x92, 0xe9, 0x36,
	0xd3, 0xba, 0x4b, 0xaf, 0x26, 0xe7, 0x7c, 0xe7, 0x7c, 0x5f, 0xce, 0x97, 0xc3, 0x40, 0x25, 0x15,
	0x37, 0x66, 0x66, 0xac, 0xf6, 0x9a, 0x54, 0x42, 0xaf, 0xd7, 0x5a, 0xc5, 0xe0, 0x14, 0x84, 0xb6,
	0x98, 0x80, 0xfa, 0x6f, 0x01, 0x77, 0x17, 0xea, 0x95, 0x31, 0xef, 0xd0, 0x39, 0x7e, 0x85, 0xe4,
	0x1e, 0x94, 0x8b, 0x4b, 0x5a, 0x4c, 0x8b, 0xb3, 0x31, 0x2b, 0x17, 0x97, 0x84, 0xc2, 0xd0, 0x58,
	0xdd, 0x6c, 0x84, 0xa7, 0x65, 0x4c, 0x6e, 0x43, 0xf2, 0x18, 0x06, 0x0e, 0x55, 0x83, 0x96, 0xf6,
	0x22, 0xd0, 0x46, 0xe4, 0x29, 0x8c, 0x2d, 0x0a, 0x69, 0x24, 0x2a, 0x4f, 0x4f, 0x22, 0xb4, 0x4b,
	0x04, 0x3e, 0xb7, 0x59, 0x7d, 0x45, 0xe1, 0x69, 0x3f, 0xf1, 0xb5, 0x61, 0x40, 0x84, 0x56, 0x3e,
	0x74, 0x0d, 0x12, 0xd2, 0x86, 0x84, 0xc0, 0x89, 0x45, 0xde, 0xd0, 0xe1, 0xb4, 0x38, 0x1b, 0xb1,
	0x78, 0x0e, 0x2a, 0xc2, 0x22, 0xf7, 0xd8, 0xac, 0x7e, 0xd2, 0x51, 0x52, 0xb9, 0x4e, 0x64, 0x28,
	0xf7, 0x74, 0xdc, 0x41, 0xb9, 0xaf, 0xaf, 0x80, 0x7c, 0x44, 0xd5, 0xb4, 0x23, 0x33, 0xfc, 0xbe,
	0x41, 0xe7, 0xbb, 0xf7, 0x2e, 0xa6, 0xbd, 0xee, 0xbd, 0xcf, 0x61, 0xb8, 0x4e, 0xf5, 0xd1, 0x87,
	0x6a, 0xfe, 0x64, 0x96, 0x79, 0x3a, 0xcb, 0x3d, 0x64, 0xdb, 0xca, 0xfa, 0x33, 0x90, 0x37, 0xe8,
	0xdb, 0xb4, 0xbb, 0x41, 0x68, 0xcf, 0xa0, 0x47, 0xd0, 0xff, 0x26, 0xd7, 0x32, 0xd9, 0xdd, 0x67,
	0x29, 0x08, 0x16, 0x98, 0xa0, 0xdd, 0x8b, 0xc9, 0x78, 0xae, 0xff, 0x14, 0x40, 0xa2, 0xee, 0x7b,
	0xed, 0xe5, 0x17, 0x86, 0xce, 0x68, 0xe5, 0x90, 0xbc, 0x85, 0x87, 0x2e, 0x9f, 0x2e, 0xa5, 0xa3,
	0xd0, 0xfe, 0xad, 0x5f, 0xdb, 0x4d, 0xb3, 0x2d, 0x60, 0xff, 0xeb, 0x22, 0x2f, 0x60, 0xd4, 0x0e,
	0xe3, 0x68, 0x39, 0xed, 0xdd, 0x3e, 0xf7, 0x75, 0x69, 0xbd, 0x04, 0xc2, 0x90, 0xef, 0x3b, 0xfc,
	0x12, 0x2a, 0xbb, 0xcb, 0x46, 0x8f, 0xab, 0x39, 0xed, 0xf0, 0xe5, 0x5d, 0x79, 0x71, 0xfd, 0x1c,
	0xaa, 0x0c, 0x3b, 0x58, 0xd3, 0xed, 0x8a, 0x94, 0xbb, 0x15, 0x99, 0xff, 0x2e, 0xe1, 0xc1, 0xce,
	0x1f, 0x29, 0xb8, 0x97, 0x5a, 0x91, 0x0f, 0x50, 0x65, 0x8f, 0x4f, 0x26, 0x1d, 0xf9, 0xc3, 0xb5,
	0x38, 0x9d, 0x1c, 0xce, 0xdb, 0xf1, 0xbb, 0xbe, 0x13, 0x28, 0xb3, 0x67, 0xde, 0xa3, 0x3c, 0x5c,
	0x80, 0x23, 0x29, 0xf3, 0x71, 0x27, 0x37, 0x9a, 0x74, 0x34, 0xe5, 0xc5, 0x33, 0xb8, 0x2f, 0x9b,
	0xd9, 0x2f, 0x5c, 0x59, 0xfe, 0xa3, 0xad, 0xbe, 0x80, 0x58, 0xb9, 0x0c, 0x4d, 0xcb, 0xe2, 0x53,
	0xfe, 0x97, 0x58, 0x0d, 0xe2, 0xe7, 0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xc3, 0x42,
	0x2c, 0x48, 0x04, 0x00, 0x00,
}
