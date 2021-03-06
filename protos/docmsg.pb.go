// Code generated by protoc-gen-go. DO NOT EDIT.
// source: docmsg.proto

package protos

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

type DocMsg struct {
	GetProtos            bool     `protobuf:"varint,1,opt,name=getProtos" json:"getProtos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocMsg) Reset()         { *m = DocMsg{} }
func (m *DocMsg) String() string { return proto.CompactTextString(m) }
func (*DocMsg) ProtoMessage()    {}
func (*DocMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_docmsg_6d699a2b0ab5525e, []int{0}
}
func (m *DocMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocMsg.Unmarshal(m, b)
}
func (m *DocMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocMsg.Marshal(b, m, deterministic)
}
func (dst *DocMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocMsg.Merge(dst, src)
}
func (m *DocMsg) XXX_Size() int {
	return xxx_messageInfo_DocMsg.Size(m)
}
func (m *DocMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DocMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DocMsg proto.InternalMessageInfo

func (m *DocMsg) GetGetProtos() bool {
	if m != nil {
		return m.GetProtos
	}
	return false
}

func init() {
	proto.RegisterType((*DocMsg)(nil), "protos.DocMsg")
}

func init() { proto.RegisterFile("docmsg.proto", fileDescriptor_docmsg_6d699a2b0ab5525e) }

var fileDescriptor_docmsg_6d699a2b0ab5525e = []byte{
	// 75 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xc9, 0x4f, 0xce,
	0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x6a, 0x5c,
	0x6c, 0x2e, 0xf9, 0xc9, 0xbe, 0xc5, 0xe9, 0x42, 0x32, 0x5c, 0x9c, 0xe9, 0xa9, 0x25, 0x01, 0x60,
	0x61, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x84, 0x40, 0x12, 0x44, 0xbd, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x77, 0x5a, 0xd8, 0xbe, 0x46, 0x00, 0x00, 0x00,
}
