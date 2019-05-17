// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic/error.proto

package basic

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

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6cc83d7bb3af054, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "protobuf.basic.Error")
}

func init() { proto.RegisterFile("basic/error.proto", fileDescriptor_d6cc83d7bb3af054) }

var fileDescriptor_d6cc83d7bb3af054 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4a, 0x2c, 0xce,
	0x4c, 0xd6, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03,
	0x53, 0x49, 0xa5, 0x69, 0x7a, 0x60, 0x39, 0x25, 0x5d, 0x2e, 0x56, 0x57, 0x90, 0xb4, 0x90, 0x10,
	0x17, 0x4b, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x24,
	0xc0, 0xc5, 0x9c, 0x5b, 0x9c, 0x2e, 0xc1, 0x04, 0x16, 0x02, 0x31, 0x9d, 0x5c, 0xa3, 0x9c, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x73, 0x33, 0x93, 0x8b, 0xf2, 0x75,
	0x33, 0xf3, 0x74, 0x93, 0xf3, 0xf4, 0x0b, 0x72, 0x12, 0x4b, 0xd2, 0xf2, 0x8b, 0x72, 0x75, 0xcb,
	0x53, 0x93, 0xf4, 0x13, 0x8b, 0x8b, 0x53, 0x73, 0x93, 0x72, 0x2a, 0x75, 0x73, 0x32, 0xf3, 0x52,
	0xf5, 0x61, 0x56, 0xea, 0xa7, 0xe7, 0xeb, 0x83, 0x6d, 0x4d, 0x62, 0x03, 0x0b, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xe1, 0x70, 0xad, 0x3a, 0xa1, 0x00, 0x00, 0x00,
}
