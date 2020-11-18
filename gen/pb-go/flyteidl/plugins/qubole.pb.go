// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/qubole.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
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

// Defines a query to execute on a hive cluster.
type HiveQuery struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	TimeoutSec           uint32   `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	RetryCount           uint32   `protobuf:"varint,3,opt,name=retryCount,proto3" json:"retryCount,omitempty"`
	StagingQuery         string   `protobuf:"bytes,4,opt,name=staging_query,json=stagingQuery,proto3" json:"staging_query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiveQuery) Reset()         { *m = HiveQuery{} }
func (m *HiveQuery) String() string { return proto.CompactTextString(m) }
func (*HiveQuery) ProtoMessage()    {}
func (*HiveQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cb86d766c12ee2e, []int{0}
}

func (m *HiveQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQuery.Unmarshal(m, b)
}
func (m *HiveQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQuery.Marshal(b, m, deterministic)
}
func (m *HiveQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQuery.Merge(m, src)
}
func (m *HiveQuery) XXX_Size() int {
	return xxx_messageInfo_HiveQuery.Size(m)
}
func (m *HiveQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQuery.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQuery proto.InternalMessageInfo

func (m *HiveQuery) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *HiveQuery) GetTimeoutSec() uint32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *HiveQuery) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *HiveQuery) GetStagingQuery() string {
	if m != nil {
		return m.StagingQuery
	}
	return ""
}

// Defines a collection of hive queries.
type HiveQueryCollection struct {
	Queries              []*HiveQuery `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HiveQueryCollection) Reset()         { *m = HiveQueryCollection{} }
func (m *HiveQueryCollection) String() string { return proto.CompactTextString(m) }
func (*HiveQueryCollection) ProtoMessage()    {}
func (*HiveQueryCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cb86d766c12ee2e, []int{1}
}

func (m *HiveQueryCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQueryCollection.Unmarshal(m, b)
}
func (m *HiveQueryCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQueryCollection.Marshal(b, m, deterministic)
}
func (m *HiveQueryCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQueryCollection.Merge(m, src)
}
func (m *HiveQueryCollection) XXX_Size() int {
	return xxx_messageInfo_HiveQueryCollection.Size(m)
}
func (m *HiveQueryCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQueryCollection.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQueryCollection proto.InternalMessageInfo

func (m *HiveQueryCollection) GetQueries() []*HiveQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

// This message works with the 'hive' task type in the SDK and is the object that will be in the 'custom' field
// of a hive task's TaskTemplate
type QuboleHiveJob struct {
	ClusterLabel         string               `protobuf:"bytes,1,opt,name=cluster_label,json=clusterLabel,proto3" json:"cluster_label,omitempty"`
	QueryCollection      *HiveQueryCollection `protobuf:"bytes,2,opt,name=query_collection,json=queryCollection,proto3" json:"query_collection,omitempty"` // Deprecated: Do not use.
	Tags                 []string             `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Query                *HiveQuery           `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QuboleHiveJob) Reset()         { *m = QuboleHiveJob{} }
func (m *QuboleHiveJob) String() string { return proto.CompactTextString(m) }
func (*QuboleHiveJob) ProtoMessage()    {}
func (*QuboleHiveJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cb86d766c12ee2e, []int{2}
}

func (m *QuboleHiveJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuboleHiveJob.Unmarshal(m, b)
}
func (m *QuboleHiveJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuboleHiveJob.Marshal(b, m, deterministic)
}
func (m *QuboleHiveJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuboleHiveJob.Merge(m, src)
}
func (m *QuboleHiveJob) XXX_Size() int {
	return xxx_messageInfo_QuboleHiveJob.Size(m)
}
func (m *QuboleHiveJob) XXX_DiscardUnknown() {
	xxx_messageInfo_QuboleHiveJob.DiscardUnknown(m)
}

var xxx_messageInfo_QuboleHiveJob proto.InternalMessageInfo

func (m *QuboleHiveJob) GetClusterLabel() string {
	if m != nil {
		return m.ClusterLabel
	}
	return ""
}

// Deprecated: Do not use.
func (m *QuboleHiveJob) GetQueryCollection() *HiveQueryCollection {
	if m != nil {
		return m.QueryCollection
	}
	return nil
}

func (m *QuboleHiveJob) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *QuboleHiveJob) GetQuery() *HiveQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func init() {
	proto.RegisterType((*HiveQuery)(nil), "flyteidl.plugins.HiveQuery")
	proto.RegisterType((*HiveQueryCollection)(nil), "flyteidl.plugins.HiveQueryCollection")
	proto.RegisterType((*QuboleHiveJob)(nil), "flyteidl.plugins.QuboleHiveJob")
}

func init() { proto.RegisterFile("flyteidl/plugins/qubole.proto", fileDescriptor_7cb86d766c12ee2e) }

var fileDescriptor_7cb86d766c12ee2e = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x4e, 0xc2, 0x30,
	0x14, 0xce, 0x18, 0x6a, 0x28, 0x10, 0x49, 0xf5, 0x62, 0x6a, 0xd4, 0x05, 0x63, 0xb2, 0x1b, 0xd7,
	0x08, 0xe1, 0x05, 0xe0, 0xc6, 0x18, 0x6e, 0x98, 0x5e, 0x79, 0xb3, 0x6c, 0xf5, 0x50, 0x1b, 0xcb,
	0x0a, 0x6b, 0x6b, 0xb2, 0x27, 0xf0, 0xf9, 0x7c, 0x23, 0xb3, 0x6e, 0x8c, 0x64, 0x17, 0xdc, 0x9d,
	0x7e, 0xdf, 0xf9, 0xf9, 0xbe, 0x73, 0x8a, 0x6e, 0xd7, 0xa2, 0xd0, 0xc0, 0x3f, 0x05, 0xd9, 0x0a,
	0xc3, 0x78, 0xa6, 0xc8, 0xce, 0xa4, 0x52, 0x40, 0xb8, 0xcd, 0xa5, 0x96, 0x78, 0xb4, 0xa7, 0xc3,
	0x9a, 0xbe, 0xbe, 0x6a, 0x0a, 0xa8, 0xcc, 0x81, 0xe8, 0x44, 0x7d, 0xab, 0x2a, 0x79, 0xfc, 0xeb,
	0xa0, 0xde, 0x0b, 0xff, 0x81, 0x95, 0x81, 0xbc, 0xc0, 0x97, 0xe8, 0x64, 0x57, 0x06, 0x9e, 0xe3,
	0x3b, 0x41, 0x2f, 0xaa, 0x1e, 0xf8, 0x1e, 0xf5, 0x35, 0xdf, 0x80, 0x34, 0x3a, 0x56, 0x40, 0xbd,
	0x8e, 0xef, 0x04, 0xc3, 0x08, 0xd5, 0xd0, 0x1b, 0x50, 0x7c, 0x87, 0x50, 0x0e, 0x3a, 0x2f, 0x16,
	0xd2, 0x64, 0xda, 0x73, 0x2b, 0xfe, 0x80, 0xe0, 0x07, 0x34, 0x54, 0x3a, 0x61, 0x3c, 0x63, 0x71,
	0xd5, 0xbe, 0x6b, 0xdb, 0x0f, 0x6a, 0xd0, 0xce, 0x1e, 0x2f, 0xd1, 0x45, 0x23, 0x64, 0x21, 0x85,
	0x00, 0xaa, 0xb9, 0xcc, 0xf0, 0x0c, 0x9d, 0x95, 0x35, 0x1c, 0x94, 0xd7, 0xf1, 0xdd, 0xa0, 0x3f,
	0xb9, 0x09, 0xdb, 0xfe, 0xc2, 0xa6, 0x2e, 0xda, 0xe7, 0x8e, 0xff, 0x1c, 0x34, 0x5c, 0xd9, 0xad,
	0x94, 0xe4, 0xab, 0x4c, 0x4b, 0x11, 0x54, 0x18, 0xa5, 0x21, 0x8f, 0x45, 0x92, 0x82, 0xa8, 0x3d,
	0x0e, 0x6a, 0x70, 0x59, 0x62, 0xf8, 0x1d, 0x8d, 0xac, 0xc2, 0x98, 0x36, 0x0a, 0xac, 0xdf, 0xfe,
	0xe4, 0xf1, 0xc8, 0xd8, 0x83, 0xdc, 0x79, 0xc7, 0x73, 0xa2, 0xf3, 0x5d, 0xcb, 0x03, 0x46, 0x5d,
	0x9d, 0x30, 0xe5, 0xb9, 0xbe, 0x1b, 0xf4, 0x22, 0x1b, 0xe3, 0xe7, 0xfd, 0xaa, 0xbb, 0xb6, 0xfd,
	0x51, 0x57, 0x55, 0xe6, 0x7c, 0xf6, 0x31, 0x65, 0x5c, 0x7f, 0x99, 0x34, 0xa4, 0x72, 0x43, 0x44,
	0xb1, 0xd6, 0xa4, 0x39, 0x2c, 0x83, 0x8c, 0x6c, 0xd3, 0x27, 0x26, 0x49, 0xfb, 0x73, 0xa4, 0xa7,
	0xf6, 0xd2, 0xd3, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x53, 0x85, 0x44, 0x37, 0x02, 0x00,
	0x00,
}
