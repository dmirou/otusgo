// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request/request.proto

package request

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

type ByID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByID) Reset()         { *m = ByID{} }
func (m *ByID) String() string { return proto.CompactTextString(m) }
func (*ByID) ProtoMessage()    {}
func (*ByID) Descriptor() ([]byte, []int) {
	return fileDescriptor_30d73066fed3fbb2, []int{0}
}

func (m *ByID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByID.Unmarshal(m, b)
}
func (m *ByID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByID.Marshal(b, m, deterministic)
}
func (m *ByID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByID.Merge(m, src)
}
func (m *ByID) XXX_Size() int {
	return xxx_messageInfo_ByID.Size(m)
}
func (m *ByID) XXX_DiscardUnknown() {
	xxx_messageInfo_ByID.DiscardUnknown(m)
}

var xxx_messageInfo_ByID proto.InternalMessageInfo

func (m *ByID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ByDate struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ByDate) Reset()         { *m = ByDate{} }
func (m *ByDate) String() string { return proto.CompactTextString(m) }
func (*ByDate) ProtoMessage()    {}
func (*ByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_30d73066fed3fbb2, []int{1}
}

func (m *ByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByDate.Unmarshal(m, b)
}
func (m *ByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByDate.Marshal(b, m, deterministic)
}
func (m *ByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByDate.Merge(m, src)
}
func (m *ByDate) XXX_Size() int {
	return xxx_messageInfo_ByDate.Size(m)
}
func (m *ByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_ByDate.DiscardUnknown(m)
}

var xxx_messageInfo_ByDate proto.InternalMessageInfo

func (m *ByDate) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterType((*ByID)(nil), "request.ByID")
	proto.RegisterType((*ByDate)(nil), "request.ByDate")
}

func init() {
	proto.RegisterFile("request/request.proto", fileDescriptor_30d73066fed3fbb2)
}

var fileDescriptor_30d73066fed3fbb2 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x05, 0x60, 0x6c, 0x8c, 0x4b, 0x55, 0xe8, 0x60, 0x68, 0x29, 0x5e, 0x5a, 0x3c, 0x75, 0xd2,
	0x41, 0x3b, 0xb4, 0x59, 0x8d, 0x97, 0xac, 0x26, 0x53, 0x36, 0x59, 0x52, 0x14, 0x11, 0xcb, 0xe7,
	0x48, 0xa7, 0xc1, 0xff, 0x3e, 0x20, 0xdb, 0xd3, 0x71, 0x77, 0x8f, 0x8f, 0xc7, 0xde, 0xbc, 0xbe,
	0x47, 0x1d, 0x08, 0xb6, 0xc9, 0x67, 0x8f, 0x84, 0xd5, 0xd3, 0xb6, 0xd6, 0x9f, 0x06, 0xd1, 0x8c,
	0x1a, 0xd2, 0x79, 0x88, 0x17, 0x20, 0xeb, 0x74, 0x20, 0xe1, 0xe6, 0x35, 0xd9, 0xbc, 0xb3, 0xa2,
	0x5d, 0x8e, 0x5d, 0xf5, 0xca, 0x72, 0xab, 0x3e, 0xb2, 0xaf, 0xec, 0xfb, 0xb9, 0xcf, 0xad, 0x6a,
	0xfe, 0x59, 0xd9, 0x2e, 0x9d, 0x20, 0x5d, 0x71, 0x56, 0x28, 0x41, 0x3a, 0xfd, 0x5e, 0x7e, 0x6a,
	0xbe, 0x8a, 0x7c, 0x17, 0xf9, 0x69, 0x17, 0xfb, 0x94, 0x6b, 0x0f, 0xe7, 0x3f, 0x63, 0xe9, 0x1a,
	0x07, 0x2e, 0xd1, 0x81, 0x72, 0xd6, 0x63, 0x04, 0xa4, 0x18, 0x0c, 0x82, 0x14, 0xa3, 0x9e, 0x94,
	0xf0, 0x30, 0xdf, 0x0c, 0x48, 0x9c, 0xc8, 0x0b, 0x49, 0x61, 0x2f, 0x3f, 0x94, 0x09, 0xfd, 0x7d,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x7f, 0x9d, 0x5e, 0xd6, 0x00, 0x00, 0x00,
}
