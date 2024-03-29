// Code generated by protoc-gen-go. DO NOT EDIT.
// source: audit.proto

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

type SendLogRequest struct {
	Logs                 []byte   `protobuf:"bytes,1,opt,name=logs,proto3" json:"logs,omitempty"`
	BulkLog              [][]byte `protobuf:"bytes,2,rep,name=bulkLog,proto3" json:"bulkLog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendLogRequest) Reset()         { *m = SendLogRequest{} }
func (m *SendLogRequest) String() string { return proto.CompactTextString(m) }
func (*SendLogRequest) ProtoMessage()    {}
func (*SendLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{0}
}

func (m *SendLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendLogRequest.Unmarshal(m, b)
}
func (m *SendLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendLogRequest.Marshal(b, m, deterministic)
}
func (m *SendLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendLogRequest.Merge(m, src)
}
func (m *SendLogRequest) XXX_Size() int {
	return xxx_messageInfo_SendLogRequest.Size(m)
}
func (m *SendLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendLogRequest proto.InternalMessageInfo

func (m *SendLogRequest) GetLogs() []byte {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *SendLogRequest) GetBulkLog() [][]byte {
	if m != nil {
		return m.BulkLog
	}
	return nil
}

type AuditResponse struct {
	PageInfo             *PageInfo      `protobuf:"bytes,1,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	CrudResponse         *CrudResponse  `protobuf:"bytes,2,opt,name=crudResponse,proto3" json:"crudResponse,omitempty"`
	Logs                 []*AuditRecord `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty"`
	Count                int32          `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AuditResponse) Reset()         { *m = AuditResponse{} }
func (m *AuditResponse) String() string { return proto.CompactTextString(m) }
func (*AuditResponse) ProtoMessage()    {}
func (*AuditResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{1}
}

func (m *AuditResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditResponse.Unmarshal(m, b)
}
func (m *AuditResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditResponse.Marshal(b, m, deterministic)
}
func (m *AuditResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditResponse.Merge(m, src)
}
func (m *AuditResponse) XXX_Size() int {
	return xxx_messageInfo_AuditResponse.Size(m)
}
func (m *AuditResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuditResponse proto.InternalMessageInfo

func (m *AuditResponse) GetPageInfo() *PageInfo {
	if m != nil {
		return m.PageInfo
	}
	return nil
}

func (m *AuditResponse) GetCrudResponse() *CrudResponse {
	if m != nil {
		return m.CrudResponse
	}
	return nil
}

func (m *AuditResponse) GetLogs() []*AuditRecord {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *AuditResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReadLogRequest struct {
	ListParam            *AuditListParam `protobuf:"bytes,1,opt,name=listParam,proto3" json:"listParam,omitempty"`
	AggrParam            *AuditAggrParam `protobuf:"bytes,2,opt,name=aggrParam,proto3" json:"aggrParam,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReadLogRequest) Reset()         { *m = ReadLogRequest{} }
func (m *ReadLogRequest) String() string { return proto.CompactTextString(m) }
func (*ReadLogRequest) ProtoMessage()    {}
func (*ReadLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{2}
}

func (m *ReadLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadLogRequest.Unmarshal(m, b)
}
func (m *ReadLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadLogRequest.Marshal(b, m, deterministic)
}
func (m *ReadLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadLogRequest.Merge(m, src)
}
func (m *ReadLogRequest) XXX_Size() int {
	return xxx_messageInfo_ReadLogRequest.Size(m)
}
func (m *ReadLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadLogRequest proto.InternalMessageInfo

func (m *ReadLogRequest) GetListParam() *AuditListParam {
	if m != nil {
		return m.ListParam
	}
	return nil
}

func (m *ReadLogRequest) GetAggrParam() *AuditAggrParam {
	if m != nil {
		return m.AggrParam
	}
	return nil
}

type AuditListParam struct {
	Search               string   `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int32    `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditListParam) Reset()         { *m = AuditListParam{} }
func (m *AuditListParam) String() string { return proto.CompactTextString(m) }
func (*AuditListParam) ProtoMessage()    {}
func (*AuditListParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{3}
}

func (m *AuditListParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditListParam.Unmarshal(m, b)
}
func (m *AuditListParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditListParam.Marshal(b, m, deterministic)
}
func (m *AuditListParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditListParam.Merge(m, src)
}
func (m *AuditListParam) XXX_Size() int {
	return xxx_messageInfo_AuditListParam.Size(m)
}
func (m *AuditListParam) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditListParam.DiscardUnknown(m)
}

var xxx_messageInfo_AuditListParam proto.InternalMessageInfo

func (m *AuditListParam) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *AuditListParam) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *AuditListParam) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *AuditListParam) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type AuditAggrParam struct {
	ProductID            string   `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
	StartDateUnix        int64    `protobuf:"varint,2,opt,name=startDateUnix,proto3" json:"startDateUnix,omitempty"`
	EndDateUnix          int64    `protobuf:"varint,3,opt,name=endDateUnix,proto3" json:"endDateUnix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditAggrParam) Reset()         { *m = AuditAggrParam{} }
func (m *AuditAggrParam) String() string { return proto.CompactTextString(m) }
func (*AuditAggrParam) ProtoMessage()    {}
func (*AuditAggrParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{4}
}

func (m *AuditAggrParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditAggrParam.Unmarshal(m, b)
}
func (m *AuditAggrParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditAggrParam.Marshal(b, m, deterministic)
}
func (m *AuditAggrParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditAggrParam.Merge(m, src)
}
func (m *AuditAggrParam) XXX_Size() int {
	return xxx_messageInfo_AuditAggrParam.Size(m)
}
func (m *AuditAggrParam) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditAggrParam.DiscardUnknown(m)
}

var xxx_messageInfo_AuditAggrParam proto.InternalMessageInfo

func (m *AuditAggrParam) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *AuditAggrParam) GetStartDateUnix() int64 {
	if m != nil {
		return m.StartDateUnix
	}
	return 0
}

func (m *AuditAggrParam) GetEndDateUnix() int64 {
	if m != nil {
		return m.EndDateUnix
	}
	return 0
}

type AuditRecord struct {
	ID                   string              `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Timestamp            string              `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UserName             string              `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	Email                string              `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Role                 string              `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Action               string              `protobuf:"bytes,6,opt,name=action,proto3" json:"action,omitempty"`
	Status               string              `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Product              string              `protobuf:"bytes,8,opt,name=product,proto3" json:"product,omitempty"`
	Organization         string              `protobuf:"bytes,9,opt,name=organization,proto3" json:"organization,omitempty"`
	AuditDetails         *AuditRecordDetails `protobuf:"bytes,10,opt,name=auditDetails,proto3" json:"auditDetails,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AuditRecord) Reset()         { *m = AuditRecord{} }
func (m *AuditRecord) String() string { return proto.CompactTextString(m) }
func (*AuditRecord) ProtoMessage()    {}
func (*AuditRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{5}
}

func (m *AuditRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditRecord.Unmarshal(m, b)
}
func (m *AuditRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditRecord.Marshal(b, m, deterministic)
}
func (m *AuditRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditRecord.Merge(m, src)
}
func (m *AuditRecord) XXX_Size() int {
	return xxx_messageInfo_AuditRecord.Size(m)
}
func (m *AuditRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AuditRecord proto.InternalMessageInfo

func (m *AuditRecord) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AuditRecord) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *AuditRecord) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuditRecord) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuditRecord) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *AuditRecord) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AuditRecord) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AuditRecord) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *AuditRecord) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *AuditRecord) GetAuditDetails() *AuditRecordDetails {
	if m != nil {
		return m.AuditDetails
	}
	return nil
}

type AuditRecordDetails struct {
	Details              string   `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditRecordDetails) Reset()         { *m = AuditRecordDetails{} }
func (m *AuditRecordDetails) String() string { return proto.CompactTextString(m) }
func (*AuditRecordDetails) ProtoMessage()    {}
func (*AuditRecordDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{6}
}

func (m *AuditRecordDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditRecordDetails.Unmarshal(m, b)
}
func (m *AuditRecordDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditRecordDetails.Marshal(b, m, deterministic)
}
func (m *AuditRecordDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditRecordDetails.Merge(m, src)
}
func (m *AuditRecordDetails) XXX_Size() int {
	return xxx_messageInfo_AuditRecordDetails.Size(m)
}
func (m *AuditRecordDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditRecordDetails.DiscardUnknown(m)
}

var xxx_messageInfo_AuditRecordDetails proto.InternalMessageInfo

func (m *AuditRecordDetails) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func init() {
	proto.RegisterType((*SendLogRequest)(nil), "commonproto.SendLogRequest")
	proto.RegisterType((*AuditResponse)(nil), "commonproto.AuditResponse")
	proto.RegisterType((*ReadLogRequest)(nil), "commonproto.ReadLogRequest")
	proto.RegisterType((*AuditListParam)(nil), "commonproto.AuditListParam")
	proto.RegisterType((*AuditAggrParam)(nil), "commonproto.AuditAggrParam")
	proto.RegisterType((*AuditRecord)(nil), "commonproto.AuditRecord")
	proto.RegisterType((*AuditRecordDetails)(nil), "commonproto.AuditRecordDetails")
}

func init() { proto.RegisterFile("audit.proto", fileDescriptor_5594839dd8e38a1b) }

var fileDescriptor_5594839dd8e38a1b = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0xd9, 0x6e, 0x9a, 0x7a, 0x9c, 0xe6, 0x87, 0x56, 0x80, 0x96, 0x80, 0x44, 0x64, 0x71,
	0xc8, 0x01, 0x45, 0xa2, 0x9c, 0x38, 0x80, 0xd4, 0xd6, 0x07, 0x2a, 0x45, 0xa8, 0x5a, 0xe8, 0x85,
	0xdb, 0xc6, 0xde, 0x9a, 0x15, 0xb6, 0x37, 0xec, 0xae, 0xab, 0xaa, 0x0f, 0xc0, 0xf3, 0x70, 0xe0,
	0x15, 0x78, 0x2f, 0xb4, 0x7f, 0x9c, 0xd8, 0x94, 0x8a, 0x43, 0x4f, 0xd9, 0x6f, 0xf6, 0x9b, 0x6f,
	0x66, 0x3e, 0xef, 0x04, 0x12, 0xda, 0x16, 0x5c, 0x2f, 0x37, 0x52, 0x68, 0x81, 0x92, 0x5c, 0xd4,
	0xb5, 0x68, 0x2c, 0x98, 0x41, 0x2e, 0x24, 0x73, 0x17, 0xe9, 0x3b, 0x98, 0x7e, 0x64, 0x4d, 0xb1,
	0x12, 0x25, 0x61, 0xdf, 0x5a, 0xa6, 0x34, 0x42, 0xb0, 0x57, 0x89, 0x52, 0xe1, 0x60, 0x1e, 0x2c,
	0x26, 0xc4, 0x9e, 0x11, 0x86, 0xf1, 0xba, 0xad, 0xbe, 0xae, 0x44, 0x89, 0xc3, 0x79, 0xb4, 0x98,
	0x90, 0x0e, 0xa6, 0xbf, 0x02, 0x38, 0x3c, 0x36, 0x85, 0x08, 0x53, 0x1b, 0xd1, 0x28, 0x86, 0x5e,
	0xc1, 0xc1, 0x86, 0x96, 0xec, 0xac, 0xb9, 0x14, 0x56, 0x23, 0x39, 0x7a, 0xb4, 0xec, 0x55, 0x5f,
	0x9e, 0xfb, 0x4b, 0xb2, 0xa5, 0xa1, 0xb7, 0x30, 0xc9, 0x65, 0x5b, 0x74, 0x12, 0x38, 0xb4, 0x69,
	0x4f, 0x06, 0x69, 0xa7, 0x3d, 0x02, 0x19, 0xd0, 0xd1, 0x4b, 0xdf, 0x71, 0x34, 0x8f, 0x16, 0xc9,
	0x11, 0x1e, 0xa4, 0xf9, 0xde, 0x72, 0x21, 0x0b, 0x3f, 0xcb, 0x43, 0x18, 0xe5, 0xa2, 0x6d, 0x34,
	0xde, 0x9b, 0x07, 0x8b, 0x11, 0x71, 0x20, 0xfd, 0x1e, 0xc0, 0x94, 0x30, 0xda, 0x37, 0xe2, 0x0d,
	0xc4, 0x15, 0x57, 0xfa, 0x9c, 0x4a, 0x5a, 0xfb, 0x49, 0x9e, 0xde, 0xd6, 0x5e, 0x75, 0x14, 0xb2,
	0x63, 0x9b, 0x54, 0x5a, 0x96, 0xd2, 0xa5, 0x86, 0x77, 0xa5, 0x1e, 0x77, 0x14, 0xb2, 0x63, 0xa7,
	0x97, 0x30, 0x1d, 0xea, 0xa2, 0xc7, 0xb0, 0xaf, 0x18, 0x95, 0xf9, 0x17, 0xdb, 0x44, 0x4c, 0x3c,
	0x32, 0x1f, 0xaa, 0xa0, 0xda, 0xb9, 0x15, 0x13, 0x7b, 0x36, 0xc3, 0x55, 0xbc, 0xe6, 0x1a, 0x47,
	0x6e, 0x38, 0x0b, 0x0c, 0xd3, 0x78, 0xed, 0x27, 0xb6, 0xe7, 0xf4, 0xca, 0xd7, 0xd9, 0x36, 0x81,
	0x9e, 0x41, 0xbc, 0x91, 0xa2, 0x68, 0x73, 0x7d, 0x96, 0xf9, 0x52, 0xbb, 0x00, 0x7a, 0x01, 0x87,
	0x4a, 0x53, 0xa9, 0x33, 0xaa, 0xd9, 0x45, 0xc3, 0xaf, 0x6d, 0xd9, 0x88, 0x0c, 0x83, 0x68, 0x0e,
	0x09, 0x6b, 0x8a, 0x2d, 0x27, 0xb2, 0x9c, 0x7e, 0x28, 0xfd, 0x19, 0x42, 0xd2, 0xfb, 0x28, 0x68,
	0x0a, 0xe1, 0xb6, 0x5c, 0x78, 0x96, 0x99, 0x2e, 0x34, 0xaf, 0x99, 0xd2, 0xb4, 0xde, 0xf8, 0xd1,
	0x76, 0x01, 0x34, 0x83, 0x83, 0x56, 0x31, 0xf9, 0x81, 0xd6, 0xcc, 0x8a, 0xc7, 0x64, 0x8b, 0xcd,
	0xec, 0xac, 0xa6, 0xbc, 0xb2, 0x63, 0xc6, 0xc4, 0x01, 0x33, 0xbb, 0x14, 0x15, 0xc3, 0x23, 0xe7,
	0x92, 0x39, 0x1b, 0x47, 0x69, 0xae, 0xb9, 0x68, 0xf0, 0xbe, 0x73, 0xd4, 0x21, 0xeb, 0xb4, 0xa6,
	0xba, 0x55, 0x78, 0xec, 0x9d, 0xb6, 0xc8, 0x3c, 0x7f, 0x6f, 0x04, 0x3e, 0xb0, 0x17, 0x1d, 0x44,
	0x29, 0x4c, 0x84, 0x2c, 0x69, 0xc3, 0x6f, 0xa8, 0xd5, 0x8b, 0xed, 0xf5, 0x20, 0x86, 0x4e, 0x61,
	0x62, 0x57, 0x31, 0x63, 0x9a, 0xf2, 0x4a, 0x61, 0xb0, 0xef, 0xe1, 0xf9, 0x5d, 0xcf, 0xd4, 0xd3,
	0xc8, 0x20, 0x29, 0x5d, 0x02, 0xba, 0xcd, 0x31, 0x8d, 0x15, 0x5e, 0xd5, 0x39, 0xd8, 0xc1, 0xa3,
	0x1f, 0x21, 0x80, 0x4d, 0xf8, 0x24, 0x8d, 0x0b, 0x19, 0x8c, 0xfd, 0x9a, 0xa3, 0xe1, 0x43, 0x1c,
	0x2e, 0xff, 0x6c, 0xf6, 0xb7, 0xae, 0xdc, 0x9a, 0xa5, 0xff, 0xa1, 0xf7, 0x90, 0x18, 0xfe, 0x89,
	0xdb, 0xfd, 0xfb, 0x28, 0x65, 0x30, 0xf6, 0xdb, 0xf6, 0x87, 0xca, 0x70, 0x07, 0xff, 0xa1, 0xb2,
	0x82, 0xff, 0x8f, 0x73, 0xcd, 0xaf, 0xd8, 0x85, 0x62, 0xf2, 0xd4, 0xec, 0xf1, 0x3d, 0xd4, 0x4e,
	0x16, 0xf0, 0x80, 0x17, 0xcb, 0x1b, 0xb6, 0x96, 0xf4, 0xda, 0x13, 0x4f, 0x46, 0x96, 0x74, 0x1e,
	0x7c, 0xee, 0xff, 0x81, 0xae, 0xf7, 0xed, 0xcf, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x02,
	0x58, 0x86, 0x6f, 0x63, 0x05, 0x00, 0x00,
}
