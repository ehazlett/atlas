// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/atlas/api/services/nameserver/v1/nameserver.proto

package nameserver

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RecordType int32

const (
	RecordType_UNKNOWN RecordType = 0
	RecordType_A       RecordType = 1
	RecordType_AAAA    RecordType = 2
	RecordType_CNAME   RecordType = 3
	RecordType_MX      RecordType = 4
	RecordType_TXT     RecordType = 5
	RecordType_SRV     RecordType = 6
)

var RecordType_name = map[int32]string{
	0: "UNKNOWN",
	1: "A",
	2: "AAAA",
	3: "CNAME",
	4: "MX",
	5: "TXT",
	6: "SRV",
}

var RecordType_value = map[string]int32{
	"UNKNOWN": 0,
	"A":       1,
	"AAAA":    2,
	"CNAME":   3,
	"MX":      4,
	"TXT":     5,
	"SRV":     6,
}

func (x RecordType) String() string {
	return proto.EnumName(RecordType_name, int32(x))
}

func (RecordType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2feef0e68dec6c99, []int{0}
}

type LookupRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2feef0e68dec6c99, []int{0}
}
func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (m *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(m, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

func (m *LookupRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Record struct {
	Type                 RecordType `protobuf:"varint,1,opt,name=type,proto3,enum=atlas.services.nameserver.v1.RecordType" json:"type,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                string     `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Options              *types.Any `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_2feef0e68dec6c99, []int{1}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_UNKNOWN
}

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Record) GetOptions() *types.Any {
	if m != nil {
		return m.Options
	}
	return nil
}

type LookupResponse struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Records              []*Record `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2feef0e68dec6c99, []int{2}
}
func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (m *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(m, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LookupResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2feef0e68dec6c99, []int{3}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2feef0e68dec6c99, []int{4}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type CreateRequest struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Records              []*Record `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2feef0e68dec6c99, []int{5}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type DeleteRequest struct {
	Type                 RecordType `protobuf:"varint,1,opt,name=type,proto3,enum=atlas.services.nameserver.v1.RecordType" json:"type,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2feef0e68dec6c99, []int{6}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_UNKNOWN
}

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("atlas.services.nameserver.v1.RecordType", RecordType_name, RecordType_value)
	proto.RegisterType((*LookupRequest)(nil), "atlas.services.nameserver.v1.LookupRequest")
	proto.RegisterType((*Record)(nil), "atlas.services.nameserver.v1.Record")
	proto.RegisterType((*LookupResponse)(nil), "atlas.services.nameserver.v1.LookupResponse")
	proto.RegisterType((*ListRequest)(nil), "atlas.services.nameserver.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "atlas.services.nameserver.v1.ListResponse")
	proto.RegisterType((*CreateRequest)(nil), "atlas.services.nameserver.v1.CreateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "atlas.services.nameserver.v1.DeleteRequest")
}

func init() {
	proto.RegisterFile("github.com/ehazlett/atlas/api/services/nameserver/v1/nameserver.proto", fileDescriptor_2feef0e68dec6c99)
}

var fileDescriptor_2feef0e68dec6c99 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x8e, 0x93, 0x40,
	0x14, 0x5d, 0x28, 0xa5, 0xee, 0xad, 0xdd, 0x90, 0x49, 0x63, 0x10, 0x7d, 0x68, 0x88, 0x26, 0x75,
	0x35, 0x43, 0xb6, 0x3e, 0x6a, 0x34, 0xb8, 0xf6, 0xc9, 0x2d, 0x9b, 0x60, 0xd5, 0xc6, 0xc4, 0x07,
	0xca, 0x5e, 0x59, 0x22, 0x65, 0x58, 0x18, 0x9a, 0xe0, 0xd7, 0xf8, 0x85, 0x7e, 0x83, 0x61, 0x00,
	0x69, 0x35, 0x6e, 0x9b, 0x8d, 0xfb, 0x36, 0x27, 0xf7, 0xdc, 0x73, 0xef, 0xe1, 0xcc, 0x00, 0xd3,
	0x20, 0xe4, 0x97, 0xf9, 0x92, 0xfa, 0x6c, 0x65, 0xe1, 0xa5, 0xf7, 0x3d, 0x42, 0xce, 0x2d, 0x8f,
	0x47, 0x5e, 0x66, 0x79, 0x49, 0x68, 0x65, 0x98, 0xae, 0x43, 0x1f, 0x33, 0x2b, 0xf6, 0x56, 0x58,
	0x02, 0x4c, 0xad, 0xf5, 0xc9, 0x06, 0xa2, 0x49, 0xca, 0x38, 0x23, 0x0f, 0x45, 0x0b, 0x6d, 0xe8,
	0x74, 0x83, 0xb0, 0x3e, 0x31, 0x86, 0x01, 0x0b, 0x98, 0x20, 0x5a, 0xe5, 0xa9, 0xea, 0x31, 0x1e,
	0x04, 0x8c, 0x05, 0x11, 0x5a, 0x02, 0x2d, 0xf3, 0xaf, 0x16, 0xae, 0x12, 0x5e, 0xd4, 0xc5, 0xfb,
	0x7f, 0x16, 0xbd, 0xb8, 0x2e, 0x99, 0x8f, 0x61, 0x70, 0xc6, 0xd8, 0xb7, 0x3c, 0x71, 0xf1, 0x2a,
	0xc7, 0x8c, 0x93, 0x21, 0x74, 0xaf, 0x72, 0x4c, 0x0b, 0x5d, 0x1a, 0x49, 0xe3, 0x43, 0xb7, 0x02,
	0xe6, 0x0f, 0x09, 0x54, 0x17, 0x7d, 0x96, 0x5e, 0x90, 0x97, 0xa0, 0xf0, 0x22, 0x41, 0x51, 0x3f,
	0x9a, 0x8c, 0xe9, 0x75, 0xcb, 0xd2, 0xaa, 0x67, 0x5e, 0x24, 0xe8, 0x8a, 0x2e, 0x42, 0x40, 0x29,
	0x19, 0xba, 0x2c, 0xd4, 0xc5, 0xb9, 0x1c, 0xb9, 0xf6, 0xa2, 0x1c, 0xf5, 0x4e, 0x35, 0x52, 0x00,
	0x42, 0xa1, 0xc7, 0x12, 0x1e, 0xb2, 0x38, 0xd3, 0x95, 0x91, 0x34, 0xee, 0x4f, 0x86, 0xb4, 0xb2,
	0x41, 0x1b, 0x1b, 0xd4, 0x8e, 0x0b, 0xb7, 0x21, 0x99, 0x17, 0x70, 0xd4, 0x38, 0xc9, 0x12, 0x16,
	0x67, 0xed, 0x2c, 0x69, 0x63, 0xd6, 0x2b, 0xe8, 0xa5, 0x62, 0xa7, 0x4c, 0x97, 0x47, 0x9d, 0x71,
	0x7f, 0xf2, 0x68, 0x1f, 0x03, 0x6e, 0xd3, 0x64, 0x0e, 0xa0, 0x7f, 0x16, 0x66, 0xbc, 0xfe, 0x5a,
	0xa6, 0x03, 0x77, 0x2b, 0x58, 0x8f, 0xdc, 0x90, 0x97, 0x6e, 0x22, 0xef, 0xc3, 0xe0, 0x34, 0x45,
	0x8f, 0x63, 0x13, 0xc7, 0x6d, 0x78, 0xf0, 0x60, 0xf0, 0x16, 0x23, 0x6c, 0x87, 0xfc, 0xf7, 0x48,
	0x8f, 0xcf, 0x01, 0x5a, 0x1e, 0xe9, 0x43, 0xef, 0x83, 0xf3, 0xce, 0x39, 0xff, 0xe4, 0x68, 0x07,
	0xa4, 0x0b, 0x92, 0xad, 0x49, 0xe4, 0x0e, 0x28, 0xb6, 0x6d, 0xdb, 0x9a, 0x4c, 0x0e, 0xa1, 0x7b,
	0xea, 0xd8, 0xb3, 0xa9, 0xd6, 0x21, 0x2a, 0xc8, 0xb3, 0x85, 0xa6, 0x90, 0x1e, 0x74, 0xe6, 0x8b,
	0xb9, 0xd6, 0x2d, 0x0f, 0xef, 0xdd, 0x8f, 0x9a, 0x3a, 0xf9, 0x29, 0x03, 0x38, 0xbf, 0xf7, 0x20,
	0x3e, 0xa8, 0x55, 0xd8, 0xe4, 0xe9, 0xf5, 0xdb, 0x6e, 0x5d, 0x6e, 0xe3, 0xd9, 0x7e, 0xe4, 0x3a,
	0xcc, 0x2f, 0xa0, 0x94, 0xe1, 0x92, 0x27, 0x3b, 0xba, 0xda, 0xfb, 0x60, 0x1c, 0xef, 0x43, 0xad,
	0xe5, 0x67, 0xa0, 0x56, 0x59, 0xef, 0xf2, 0xb0, 0x75, 0x23, 0x8c, 0x7b, 0x7f, 0x3d, 0x83, 0x69,
	0xf9, 0xd4, 0x4b, 0xb9, 0x2a, 0xd5, 0x5d, 0x72, 0x5b, 0xd9, 0xff, 0x4b, 0xee, 0x8d, 0xfd, 0xf9,
	0xf5, 0x4d, 0xfe, 0x66, 0x2f, 0x5a, 0xb4, 0x38, 0x58, 0xaa, 0x42, 0xf4, 0xf9, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0x0c, 0x2f, 0x34, 0x19, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NameserverClient is the client API for Nameserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NameserverClient interface {
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type nameserverClient struct {
	cc *grpc.ClientConn
}

func NewNameserverClient(cc *grpc.ClientConn) NameserverClient {
	return &nameserverClient{cc}
}

func (c *nameserverClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := c.cc.Invoke(ctx, "/atlas.services.nameserver.v1.Nameserver/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/atlas.services.nameserver.v1.Nameserver/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/atlas.services.nameserver.v1.Nameserver/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/atlas.services.nameserver.v1.Nameserver/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameserverServer is the server API for Nameserver service.
type NameserverServer interface {
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Create(context.Context, *CreateRequest) (*types.Empty, error)
	Delete(context.Context, *DeleteRequest) (*types.Empty, error)
}

// UnimplementedNameserverServer can be embedded to have forward compatible implementations.
type UnimplementedNameserverServer struct {
}

func (*UnimplementedNameserverServer) Lookup(ctx context.Context, req *LookupRequest) (*LookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (*UnimplementedNameserverServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedNameserverServer) Create(ctx context.Context, req *CreateRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNameserverServer) Delete(ctx context.Context, req *DeleteRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterNameserverServer(s *grpc.Server, srv NameserverServer) {
	s.RegisterService(&_Nameserver_serviceDesc, srv)
}

func _Nameserver_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.services.nameserver.v1.Nameserver/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.services.nameserver.v1.Nameserver/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.services.nameserver.v1.Nameserver/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.services.nameserver.v1.Nameserver/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nameserver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atlas.services.nameserver.v1.Nameserver",
	HandlerType: (*NameserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Nameserver_Lookup_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Nameserver_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Nameserver_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Nameserver_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/atlas/api/services/nameserver/v1/nameserver.proto",
}
