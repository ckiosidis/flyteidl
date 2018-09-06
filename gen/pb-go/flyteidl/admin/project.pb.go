// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/project.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Domain struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_3a260572ca4c894f, []int{0}
}
func (m *Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain.Unmarshal(m, b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
}
func (dst *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(dst, src)
}
func (m *Domain) XXX_Size() int {
	return xxx_messageInfo_Domain.Size(m)
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Domain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Project struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name.
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Domains              []*Domain `protobuf:"bytes,3,rep,name=domains,proto3" json:"domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_3a260572ca4c894f, []int{1}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (dst *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(dst, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetDomains() []*Domain {
	if m != nil {
		return m.Domains
	}
	return nil
}

type Projects struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Projects) Reset()         { *m = Projects{} }
func (m *Projects) String() string { return proto.CompactTextString(m) }
func (*Projects) ProtoMessage()    {}
func (*Projects) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_3a260572ca4c894f, []int{2}
}
func (m *Projects) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Projects.Unmarshal(m, b)
}
func (m *Projects) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Projects.Marshal(b, m, deterministic)
}
func (dst *Projects) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Projects.Merge(dst, src)
}
func (m *Projects) XXX_Size() int {
	return xxx_messageInfo_Projects.Size(m)
}
func (m *Projects) XXX_DiscardUnknown() {
	xxx_messageInfo_Projects.DiscardUnknown(m)
}

var xxx_messageInfo_Projects proto.InternalMessageInfo

func (m *Projects) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ProjectListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectListRequest) Reset()         { *m = ProjectListRequest{} }
func (m *ProjectListRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectListRequest) ProtoMessage()    {}
func (*ProjectListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_3a260572ca4c894f, []int{3}
}
func (m *ProjectListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectListRequest.Unmarshal(m, b)
}
func (m *ProjectListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectListRequest.Marshal(b, m, deterministic)
}
func (dst *ProjectListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectListRequest.Merge(dst, src)
}
func (m *ProjectListRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectListRequest.Size(m)
}
func (m *ProjectListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectListRequest proto.InternalMessageInfo

type ProjectRegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectRegisterResponse) Reset()         { *m = ProjectRegisterResponse{} }
func (m *ProjectRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectRegisterResponse) ProtoMessage()    {}
func (*ProjectRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_3a260572ca4c894f, []int{4}
}
func (m *ProjectRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectRegisterResponse.Unmarshal(m, b)
}
func (m *ProjectRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectRegisterResponse.Marshal(b, m, deterministic)
}
func (dst *ProjectRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectRegisterResponse.Merge(dst, src)
}
func (m *ProjectRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectRegisterResponse.Size(m)
}
func (m *ProjectRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectRegisterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Domain)(nil), "flyteidl.admin.Domain")
	proto.RegisterType((*Project)(nil), "flyteidl.admin.Project")
	proto.RegisterType((*Projects)(nil), "flyteidl.admin.Projects")
	proto.RegisterType((*ProjectListRequest)(nil), "flyteidl.admin.ProjectListRequest")
	proto.RegisterType((*ProjectRegisterResponse)(nil), "flyteidl.admin.ProjectRegisterResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/project.proto", fileDescriptor_project_3a260572ca4c894f)
}

var fileDescriptor_project_3a260572ca4c894f = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x14, 0xb5, 0xe5, 0x90, 0x3a, 0x58, 0x88, 0x06, 0x89, 0xa1, 0xf2, 0xd4, 0x01,
	0x6c, 0x20, 0x3f, 0x00, 0x09, 0x31, 0x32, 0x20, 0x8f, 0x2c, 0x28, 0xa9, 0xaf, 0xc1, 0x28, 0xb1,
	0x43, 0xee, 0x3a, 0xf4, 0xdf, 0x23, 0xb9, 0x4e, 0xa5, 0x32, 0x75, 0xb3, 0xee, 0x7d, 0xef, 0x3d,
	0xeb, 0xc1, 0xdd, 0xb6, 0xdd, 0x33, 0x3a, 0xdb, 0xea, 0xca, 0x76, 0xce, 0xeb, 0x7e, 0x08, 0x3f,
	0xb8, 0x61, 0xd5, 0x0f, 0x81, 0x83, 0x58, 0x8c, 0xaa, 0x8a, 0xaa, 0xbc, 0x87, 0xe9, 0x5b, 0xe8,
	0x2a, 0xe7, 0xc5, 0x02, 0x72, 0x67, 0x8b, 0x6c, 0x95, 0xad, 0x2f, 0x4d, 0xee, 0xac, 0x10, 0x70,
	0xe1, 0xab, 0x0e, 0x8b, 0x3c, 0x5e, 0xe2, 0x5b, 0x7e, 0xc1, 0xec, 0xe3, 0x10, 0x77, 0x0e, 0x2e,
	0x1e, 0x61, 0x66, 0x63, 0x38, 0x15, 0x93, 0xd5, 0x64, 0x7d, 0xf5, 0x7c, 0xa3, 0x4e, 0xeb, 0xd5,
	0xa1, 0xdb, 0x8c, 0x98, 0x7c, 0x81, 0x79, 0x2a, 0x20, 0x51, 0xc2, 0x3c, 0xfd, 0x9d, 0x8a, 0x2c,
	0xda, 0x97, 0xff, 0xed, 0x89, 0x35, 0x47, 0x50, 0x5e, 0x83, 0x48, 0xc7, 0x77, 0x47, 0x6c, 0xf0,
	0x77, 0x87, 0xc4, 0xf2, 0x16, 0x96, 0x23, 0x8a, 0x8d, 0x23, 0xc6, 0xc1, 0x20, 0xf5, 0xc1, 0x13,
	0xbe, 0x96, 0x9f, 0x4f, 0x8d, 0xe3, 0xef, 0x5d, 0xad, 0x36, 0xa1, 0xd3, 0xed, 0x7e, 0xcb, 0xfa,
	0x38, 0x60, 0x83, 0x5e, 0xf7, 0xf5, 0x43, 0x13, 0xf4, 0xe9, 0xa6, 0xf5, 0x34, 0x8e, 0x59, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x97, 0x5b, 0xa4, 0x6c, 0x01, 0x00, 0x00,
}
