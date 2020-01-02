// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/audit.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Request_Mode int32

const (
	Request_READ_ONLY  Request_Mode = 0
	Request_READ_WRITE Request_Mode = 1
)

var Request_Mode_name = map[int32]string{
	0: "READ_ONLY",
	1: "READ_WRITE",
}

var Request_Mode_value = map[string]int32{
	"READ_ONLY":  0,
	"READ_WRITE": 1,
}

func (x Request_Mode) String() string {
	return proto.EnumName(Request_Mode_name, int32(x))
}

func (Request_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{3, 0}
}

// Audit log that is emitted on each user request.
type AuditLog struct {
	Principal            *Principal `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	Client               *Client    `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Request              *Request   `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Response             *Response  `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AuditLog) Reset()         { *m = AuditLog{} }
func (m *AuditLog) String() string { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()    {}
func (*AuditLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{0}
}

func (m *AuditLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditLog.Unmarshal(m, b)
}
func (m *AuditLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditLog.Marshal(b, m, deterministic)
}
func (m *AuditLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLog.Merge(m, src)
}
func (m *AuditLog) XXX_Size() int {
	return xxx_messageInfo_AuditLog.Size(m)
}
func (m *AuditLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditLog.DiscardUnknown(m)
}

var xxx_messageInfo_AuditLog proto.InternalMessageInfo

func (m *AuditLog) GetPrincipal() *Principal {
	if m != nil {
		return m.Principal
	}
	return nil
}

func (m *AuditLog) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *AuditLog) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AuditLog) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

// Details about the authenticated user who issued a service request.
type Principal struct {
	// Identifies authenticated end-user
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// The client that initiated the auth flow.
	ClientId             string               `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TokenIssuedAt        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=token_issued_at,json=tokenIssuedAt,proto3" json:"token_issued_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{1}
}

func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

func (m *Principal) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Principal) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Principal) GetTokenIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.TokenIssuedAt
	}
	return nil
}

// Details about the user's browser client.
type Client struct {
	ClientIp             string   `protobuf:"bytes,1,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{2}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

// Details about a specific request issued by a user.
type Request struct {
	// Service method endpoint e.g. GetWorkflowExecution
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// In the case of HTTP service calls.
	HttpVerb string `protobuf:"bytes,2,opt,name=http_verb,json=httpVerb,proto3" json:"http_verb,omitempty"`
	// Includes parameters submitted in the request.
	Parameters           map[string]string    `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Mode                 Request_Mode         `protobuf:"varint,4,opt,name=mode,proto3,enum=flyteidl.admin.Request_Mode" json:"mode,omitempty"`
	ReceivedAt           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetHttpVerb() string {
	if m != nil {
		return m.HttpVerb
	}
	return ""
}

func (m *Request) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Request) GetMode() Request_Mode {
	if m != nil {
		return m.Mode
	}
	return Request_READ_ONLY
}

func (m *Request) GetReceivedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ReceivedAt
	}
	return nil
}

// Summary of service response details.
type Response struct {
	// e.g. gRPC status code
	ResponseCode         string               `protobuf:"bytes,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	SentAt               *timestamp.Timestamp `protobuf:"bytes,2,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponseCode() string {
	if m != nil {
		return m.ResponseCode
	}
	return ""
}

func (m *Response) GetSentAt() *timestamp.Timestamp {
	if m != nil {
		return m.SentAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.admin.Request_Mode", Request_Mode_name, Request_Mode_value)
	proto.RegisterType((*AuditLog)(nil), "flyteidl.admin.AuditLog")
	proto.RegisterType((*Principal)(nil), "flyteidl.admin.Principal")
	proto.RegisterType((*Client)(nil), "flyteidl.admin.Client")
	proto.RegisterType((*Request)(nil), "flyteidl.admin.Request")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.Request.ParametersEntry")
	proto.RegisterType((*Response)(nil), "flyteidl.admin.Response")
}

func init() { proto.RegisterFile("flyteidl/admin/audit.proto", fileDescriptor_11f7dfb38018d590) }

var fileDescriptor_11f7dfb38018d590 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xdc, 0xb4, 0x49, 0x3c, 0xf9, 0x92, 0x46, 0x2b, 0x54, 0x4c, 0x40, 0xa2, 0x32, 0xaa,
	0xe8, 0x0d, 0x36, 0x4d, 0x90, 0x40, 0x20, 0x2e, 0xd2, 0x12, 0xa1, 0x48, 0x05, 0xaa, 0x55, 0x05,
	0x82, 0x1b, 0xcb, 0x3f, 0x13, 0x67, 0xa9, 0xed, 0x35, 0xf6, 0x3a, 0x52, 0x5e, 0x80, 0xa7, 0xec,
	0xc3, 0x20, 0xef, 0x7a, 0x43, 0x1b, 0x54, 0xb8, 0xf3, 0xec, 0x9c, 0x33, 0x7b, 0xce, 0xf1, 0x2c,
	0x8c, 0x16, 0xc9, 0x5a, 0x20, 0x8b, 0x12, 0xd7, 0x8f, 0x52, 0x96, 0xb9, 0x7e, 0x15, 0x31, 0xe1,
	0xe4, 0x05, 0x17, 0x9c, 0x0c, 0x74, 0xcf, 0x91, 0xbd, 0xd1, 0xe3, 0x98, 0xf3, 0x38, 0x41, 0x57,
	0x76, 0x83, 0x6a, 0xe1, 0x0a, 0x96, 0x62, 0x29, 0xfc, 0x34, 0x57, 0x04, 0xfb, 0xda, 0x80, 0xee,
	0xb4, 0x1e, 0x70, 0xce, 0x63, 0xf2, 0x12, 0xcc, 0xbc, 0x60, 0x59, 0xc8, 0x72, 0x3f, 0xb1, 0x8c,
	0x43, 0xe3, 0xb8, 0x37, 0x7e, 0xe0, 0xdc, 0x9e, 0xe8, 0x5c, 0x68, 0x00, 0xfd, 0x8d, 0x25, 0x0e,
	0xb4, 0xc3, 0x84, 0x61, 0x26, 0xac, 0x1d, 0xc9, 0x3a, 0xd8, 0x66, 0x9d, 0xc9, 0x2e, 0x6d, 0x50,
	0xe4, 0x04, 0x3a, 0x05, 0xfe, 0xa8, 0xb0, 0x14, 0x56, 0x4b, 0x12, 0xee, 0x6f, 0x13, 0xa8, 0x6a,
	0x53, 0x8d, 0x23, 0x2f, 0xa0, 0x5b, 0x60, 0x99, 0xf3, 0xac, 0x44, 0x6b, 0x57, 0x72, 0xac, 0x3f,
	0x39, 0xaa, 0x4f, 0x37, 0x48, 0xfb, 0xa7, 0x01, 0xe6, 0x46, 0x31, 0xb1, 0xa0, 0x53, 0x56, 0xc1,
	0x77, 0x0c, 0x85, 0x74, 0x67, 0x52, 0x5d, 0x92, 0x87, 0x60, 0x2a, 0x69, 0x1e, 0x8b, 0xa4, 0x07,
	0x93, 0x76, 0xd5, 0xc1, 0x3c, 0x22, 0xa7, 0xb0, 0x2f, 0xf8, 0x15, 0x66, 0x1e, 0x2b, 0xcb, 0x0a,
	0x23, 0xcf, 0xd7, 0xaa, 0x47, 0x8e, 0x8a, 0xd7, 0xd1, 0xf1, 0x3a, 0x97, 0x3a, 0x5e, 0xda, 0x97,
	0x94, 0xb9, 0x64, 0x4c, 0x85, 0x7d, 0x04, 0x6d, 0x95, 0xc1, 0xcd, 0xab, 0xf2, 0x46, 0x86, 0xbe,
	0x2a, 0xb7, 0xaf, 0x77, 0xa0, 0xd3, 0x58, 0x27, 0x07, 0xd0, 0x4e, 0x51, 0x2c, 0x79, 0xd4, 0xa0,
	0x9a, 0xaa, 0x1e, 0xb0, 0x14, 0x22, 0xf7, 0x56, 0x58, 0x04, 0x5a, 0x6b, 0x7d, 0xf0, 0x19, 0x8b,
	0x80, 0xbc, 0x07, 0xc8, 0xfd, 0xc2, 0x4f, 0x51, 0x60, 0x51, 0x5a, 0xad, 0xc3, 0xd6, 0x71, 0x6f,
	0xfc, 0xf4, 0x8e, 0x70, 0x9d, 0x8b, 0x0d, 0x72, 0x96, 0x89, 0x62, 0x4d, 0x6f, 0x50, 0xc9, 0x73,
	0xd8, 0x4d, 0x79, 0xa4, 0xb2, 0x1e, 0x8c, 0x1f, 0xdd, 0x35, 0xe2, 0x03, 0x8f, 0x90, 0x4a, 0x24,
	0x79, 0x03, 0xbd, 0x02, 0x43, 0x64, 0x2b, 0x15, 0xd1, 0xde, 0x3f, 0x23, 0x02, 0x0d, 0x9f, 0x8a,
	0xd1, 0x5b, 0xd8, 0xdf, 0x52, 0x43, 0x86, 0xd0, 0xba, 0xc2, 0x75, 0x63, 0xbe, 0xfe, 0x24, 0xf7,
	0x60, 0x6f, 0xe5, 0x27, 0x15, 0x36, 0xae, 0x55, 0xf1, 0x7a, 0xe7, 0x95, 0x61, 0x1f, 0xc1, 0x6e,
	0xad, 0x84, 0xf4, 0xc1, 0xa4, 0xb3, 0xe9, 0x3b, 0xef, 0xd3, 0xc7, 0xf3, 0xaf, 0xc3, 0xff, 0xc8,
	0x00, 0x40, 0x96, 0x5f, 0xe8, 0xfc, 0x72, 0x36, 0x34, 0xec, 0x08, 0xba, 0x7a, 0x49, 0xc8, 0x13,
	0xe8, 0xeb, 0x35, 0xf1, 0xc2, 0xda, 0xa9, 0xba, 0xe8, 0x7f, 0x7d, 0x78, 0x56, 0xcf, 0x9b, 0x40,
	0xa7, 0xac, 0x7f, 0x95, 0xaf, 0x37, 0xfb, 0x6f, 0x7e, 0xda, 0x35, 0x74, 0x2a, 0x4e, 0x27, 0xdf,
	0x4e, 0x62, 0x26, 0x96, 0x55, 0xe0, 0x84, 0x3c, 0x75, 0x93, 0xf5, 0x42, 0xb8, 0x9b, 0x27, 0x1b,
	0x63, 0xe6, 0xe6, 0xc1, 0xb3, 0x98, 0xbb, 0xb7, 0x5f, 0x71, 0xd0, 0x96, 0x03, 0x27, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xf8, 0x89, 0x5c, 0x3c, 0xde, 0x03, 0x00, 0x00,
}