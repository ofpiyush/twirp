// Code generated by protoc-gen-go. DO NOT EDIT.
// source: json_serialization.proto

package json_serialization

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Msg_FooBar int32

const (
	Msg_FOO Msg_FooBar = 0
	Msg_BAR Msg_FooBar = 1
)

var Msg_FooBar_name = map[int32]string{
	0: "FOO",
	1: "BAR",
}

var Msg_FooBar_value = map[string]int32{
	"FOO": 0,
	"BAR": 1,
}

func (x Msg_FooBar) String() string {
	return proto.EnumName(Msg_FooBar_name, int32(x))
}

func (Msg_FooBar) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bb56edd8c69db2c, []int{0, 0}
}

type Msg struct {
	Query                string     `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32      `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Hell                 float64    `protobuf:"fixed64,3,opt,name=hell,proto3" json:"hell,omitempty"`
	Foobar               Msg_FooBar `protobuf:"varint,4,opt,name=foobar,proto3,enum=Msg_FooBar" json:"foobar,omitempty"`
	Snippets             []string   `protobuf:"bytes,5,rep,name=snippets,proto3" json:"snippets,omitempty"`
	AllEmpty             bool       `protobuf:"varint,6,opt,name=all_empty,json=allEmpty,proto3" json:"all_empty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bb56edd8c69db2c, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Msg) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *Msg) GetHell() float64 {
	if m != nil {
		return m.Hell
	}
	return 0
}

func (m *Msg) GetFoobar() Msg_FooBar {
	if m != nil {
		return m.Foobar
	}
	return Msg_FOO
}

func (m *Msg) GetSnippets() []string {
	if m != nil {
		return m.Snippets
	}
	return nil
}

func (m *Msg) GetAllEmpty() bool {
	if m != nil {
		return m.AllEmpty
	}
	return false
}

type Result struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bb56edd8c69db2c, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("Msg_FooBar", Msg_FooBar_name, Msg_FooBar_value)
	proto.RegisterType((*Msg)(nil), "Msg")
	proto.RegisterType((*Result)(nil), "Result")
}

func init() { proto.RegisterFile("json_serialization.proto", fileDescriptor_2bb56edd8c69db2c) }

var fileDescriptor_2bb56edd8c69db2c = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xbb, 0xe6, 0x8f, 0xc9, 0x14, 0xa4, 0x0e, 0x3d, 0x2c, 0xf1, 0xe0, 0x12, 0x2f, 0x7b,
	0x31, 0x87, 0xfa, 0x09, 0x0c, 0xb4, 0x07, 0xa1, 0x0d, 0x6c, 0x6f, 0x5e, 0xc2, 0x46, 0xd6, 0x34,
	0xb2, 0xcd, 0xc6, 0xdd, 0xe4, 0x50, 0x3f, 0xa2, 0x9f, 0x4a, 0x92, 0x8a, 0x08, 0xbd, 0x0c, 0xef,
	0xf7, 0x1e, 0x0c, 0x33, 0x0f, 0xe8, 0x87, 0x33, 0x6d, 0xe9, 0x94, 0x6d, 0xa4, 0x6e, 0xbe, 0x64,
	0xdf, 0x98, 0x36, 0xeb, 0xac, 0xe9, 0x4d, 0xfa, 0x4d, 0xc0, 0xdb, 0xba, 0x1a, 0x97, 0x10, 0x7c,
	0x0e, 0xca, 0x9e, 0x28, 0x61, 0x84, 0xc7, 0xe2, 0x0c, 0x78, 0x0f, 0xf3, 0x4e, 0xd6, 0xaa, 0x6c,
	0x87, 0x63, 0xa5, 0x2c, 0xbd, 0x62, 0x84, 0x07, 0x02, 0x46, 0x6b, 0x37, 0x39, 0x88, 0xe0, 0x1f,
	0x94, 0xd6, 0xd4, 0x63, 0x84, 0x13, 0x31, 0x69, 0x7c, 0x80, 0xf0, 0xdd, 0x98, 0x4a, 0x5a, 0xea,
	0x33, 0xc2, 0x6f, 0x56, 0xf3, 0x6c, 0xeb, 0xea, 0x6c, 0x63, 0x4c, 0x2e, 0xad, 0xf8, 0x8d, 0x30,
	0x81, 0xc8, 0xb5, 0x4d, 0xd7, 0xa9, 0xde, 0xd1, 0x80, 0x79, 0x3c, 0x16, 0x7f, 0x8c, 0x77, 0x10,
	0x4b, 0xad, 0x4b, 0x75, 0xec, 0xfa, 0x13, 0x0d, 0x19, 0xe1, 0x91, 0x88, 0xa4, 0xd6, 0xeb, 0x91,
	0xd3, 0x04, 0xc2, 0xf3, 0x2a, 0xbc, 0x06, 0x6f, 0x53, 0x14, 0x8b, 0xd9, 0x28, 0xf2, 0x67, 0xb1,
	0x20, 0x69, 0x04, 0xa1, 0x50, 0x6e, 0xd0, 0xfd, 0xea, 0x11, 0x6e, 0x5f, 0xf6, 0xc5, 0x6e, 0xff,
	0xff, 0x63, 0xa4, 0x10, 0xad, 0xdf, 0x0e, 0x66, 0x0c, 0xd0, 0x1f, 0x8f, 0x4a, 0xa6, 0x99, 0xce,
	0xf2, 0xe5, 0x2b, 0x5e, 0x36, 0x54, 0x85, 0x53, 0x45, 0x4f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x53, 0x09, 0x7b, 0x23, 0x3e, 0x01, 0x00, 0x00,
}
