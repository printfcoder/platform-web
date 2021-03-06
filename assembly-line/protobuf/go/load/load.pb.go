// Code generated by protoc-gen-go. DO NOT EDIT.
// source: load/load.proto

package load

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	basic "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/basic"
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

type LoadAvgStat struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Load1                float64              `protobuf:"fixed64,2,opt,name=load1,proto3" json:"load1,omitempty"`
	Load5                float64              `protobuf:"fixed64,3,opt,name=load5,proto3" json:"load5,omitempty"`
	Load15               float64              `protobuf:"fixed64,4,opt,name=load15,proto3" json:"load15,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoadAvgStat) Reset()         { *m = LoadAvgStat{} }
func (m *LoadAvgStat) String() string { return proto.CompactTextString(m) }
func (*LoadAvgStat) ProtoMessage()    {}
func (*LoadAvgStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_750441b24f6866ff, []int{0}
}

func (m *LoadAvgStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadAvgStat.Unmarshal(m, b)
}
func (m *LoadAvgStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadAvgStat.Marshal(b, m, deterministic)
}
func (m *LoadAvgStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadAvgStat.Merge(m, src)
}
func (m *LoadAvgStat) XXX_Size() int {
	return xxx_messageInfo_LoadAvgStat.Size(m)
}
func (m *LoadAvgStat) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadAvgStat.DiscardUnknown(m)
}

var xxx_messageInfo_LoadAvgStat proto.InternalMessageInfo

func (m *LoadAvgStat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LoadAvgStat) GetLoad1() float64 {
	if m != nil {
		return m.Load1
	}
	return 0
}

func (m *LoadAvgStat) GetLoad5() float64 {
	if m != nil {
		return m.Load5
	}
	return 0
}

func (m *LoadAvgStat) GetLoad15() float64 {
	if m != nil {
		return m.Load15
	}
	return 0
}

type LoadRequest struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IP                   string               `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	NodeName             string               `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	LoadAvgStat          []*LoadAvgStat       `protobuf:"bytes,4,rep,name=loadAvgStat,proto3" json:"loadAvgStat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoadRequest) Reset()         { *m = LoadRequest{} }
func (m *LoadRequest) String() string { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()    {}
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_750441b24f6866ff, []int{1}
}

func (m *LoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadRequest.Unmarshal(m, b)
}
func (m *LoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadRequest.Marshal(b, m, deterministic)
}
func (m *LoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadRequest.Merge(m, src)
}
func (m *LoadRequest) XXX_Size() int {
	return xxx_messageInfo_LoadRequest.Size(m)
}
func (m *LoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadRequest proto.InternalMessageInfo

func (m *LoadRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LoadRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *LoadRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LoadRequest) GetLoadAvgStat() []*LoadAvgStat {
	if m != nil {
		return m.LoadAvgStat
	}
	return nil
}

type LoadResponse struct {
	Error                *basic.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LoadResponse) Reset()         { *m = LoadResponse{} }
func (m *LoadResponse) String() string { return proto.CompactTextString(m) }
func (*LoadResponse) ProtoMessage()    {}
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_750441b24f6866ff, []int{2}
}

func (m *LoadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadResponse.Unmarshal(m, b)
}
func (m *LoadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadResponse.Marshal(b, m, deterministic)
}
func (m *LoadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadResponse.Merge(m, src)
}
func (m *LoadResponse) XXX_Size() int {
	return xxx_messageInfo_LoadResponse.Size(m)
}
func (m *LoadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadResponse proto.InternalMessageInfo

func (m *LoadResponse) GetError() *basic.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadAvgStat)(nil), "protobuf.pb.load.LoadAvgStat")
	proto.RegisterType((*LoadRequest)(nil), "protobuf.pb.load.LoadRequest")
	proto.RegisterType((*LoadResponse)(nil), "protobuf.pb.load.LoadResponse")
}

func init() { proto.RegisterFile("load/load.proto", fileDescriptor_750441b24f6866ff) }

var fileDescriptor_750441b24f6866ff = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xdf, 0xee, 0x1f, 0xef, 0x52, 0x71, 0x1a, 0x54, 0x4a, 0x41, 0x1d, 0xbd, 0x1a, 0x48,
	0x53, 0x9c, 0x0c, 0x04, 0x2f, 0xc4, 0xa1, 0x17, 0x03, 0x91, 0x91, 0x79, 0xe5, 0x5d, 0xda, 0x65,
	0x5d, 0xa1, 0x69, 0x6a, 0x92, 0x4e, 0xfc, 0x0c, 0x7e, 0x1d, 0x3f, 0xa0, 0x34, 0xe9, 0xd6, 0x22,
	0xec, 0xca, 0x9b, 0x96, 0xf3, 0x9c, 0xa7, 0xe9, 0x2f, 0xcf, 0x39, 0x60, 0x90, 0x72, 0xb2, 0x0c,
	0xca, 0x07, 0xca, 0x05, 0x57, 0x1c, 0x1e, 0xe9, 0x57, 0x58, 0xac, 0x50, 0x1e, 0xa2, 0x52, 0x77,
	0x8f, 0x43, 0x22, 0x93, 0x28, 0xa0, 0x42, 0x70, 0x61, 0x4c, 0xee, 0x65, 0xcc, 0x79, 0x9c, 0xd2,
	0x60, 0xeb, 0x0d, 0x54, 0xc2, 0xa8, 0x54, 0x84, 0xe5, 0xc6, 0xe0, 0x7d, 0x59, 0xc0, 0x7e, 0xe6,
	0x64, 0xf9, 0xb0, 0x89, 0x17, 0x8a, 0x28, 0x78, 0x0b, 0xfa, 0x3b, 0x8b, 0x63, 0x0d, 0xad, 0x91,
	0x3d, 0x76, 0x91, 0x39, 0x04, 0xed, 0x7e, 0xf8, 0xba, 0x75, 0xe0, 0xda, 0x0c, 0x4f, 0x40, 0xb7,
	0xa4, 0xb8, 0x76, 0x5a, 0x43, 0x6b, 0x64, 0x61, 0x53, 0x6c, 0xd5, 0x89, 0xd3, 0xae, 0xd5, 0x09,
	0x3c, 0x03, 0x3d, 0xdd, 0x9e, 0x38, 0x1d, 0x2d, 0x57, 0x95, 0xf7, 0x5d, 0xd1, 0x60, 0xfa, 0x5e,
	0x50, 0xf9, 0x17, 0x9a, 0x43, 0xd0, 0x9a, 0xcd, 0x35, 0x4a, 0x1f, 0xb7, 0x66, 0x73, 0xe8, 0x82,
	0xff, 0x19, 0x5f, 0xd2, 0x17, 0xc2, 0xa8, 0x46, 0xe9, 0xe3, 0x5d, 0x0d, 0xef, 0x81, 0x9d, 0xd6,
	0x11, 0x38, 0x9d, 0x61, 0x7b, 0x64, 0x8f, 0xcf, 0xd1, 0xef, 0x7c, 0x51, 0x23, 0x27, 0xdc, 0xfc,
	0xc2, 0xbb, 0x03, 0x07, 0x86, 0x5a, 0xe6, 0x3c, 0x93, 0x14, 0x5e, 0x81, 0xae, 0x1e, 0x42, 0x85,
	0x7c, 0x5a, 0x1f, 0xa5, 0x27, 0x84, 0x9e, 0xca, 0x26, 0x36, 0x9e, 0x31, 0x31, 0x57, 0x5e, 0x50,
	0xb1, 0x49, 0x22, 0x0a, 0x31, 0x18, 0xcc, 0x0b, 0xb9, 0x6e, 0xce, 0x64, 0x0f, 0x4a, 0x15, 0x92,
	0x7b, 0xb1, 0xaf, 0x6d, 0x68, 0xbc, 0x7f, 0xd3, 0xc7, 0xb7, 0x69, 0x9c, 0xa8, 0x75, 0x11, 0xa2,
	0x88, 0xb3, 0x80, 0x25, 0x91, 0xe0, 0x7e, 0x92, 0xf9, 0x51, 0x16, 0xe4, 0x29, 0x51, 0x2b, 0x2e,
	0x98, 0xff, 0x41, 0xc3, 0x80, 0x48, 0x49, 0x59, 0x98, 0x7e, 0xfa, 0x69, 0x92, 0x35, 0x56, 0x26,
	0xe6, 0x7a, 0xed, 0xc2, 0x9e, 0x56, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x71, 0x93,
	0xac, 0x8a, 0x02, 0x00, 0x00,
}
