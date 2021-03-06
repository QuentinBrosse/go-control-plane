// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_http1_bridge/v2/config.proto

package envoy_config_filter_http_grpc_http1_bridge_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Config struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_e529266ab36980e8, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Config)(nil), "envoy.config.filter.http.grpc_http1_bridge.v2.Config")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_http1_bridge/v2/config.proto", fileDescriptor_e529266ab36980e8)
}

var fileDescriptor_e529266ab36980e8 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x07, 0xb1, 0x0c, 0xe3, 0x93, 0x8a,
	0x32, 0x53, 0xd2, 0x53, 0xf5, 0xcb, 0x8c, 0xa0, 0xea, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x74, 0xc1, 0x7a, 0xf5, 0xa0, 0x62, 0x10, 0xbd, 0x7a, 0x20, 0x1d, 0x7a, 0x18, 0x7a, 0xf5, 0xca,
	0x8c, 0xa4, 0xe4, 0x4a, 0x53, 0x0a, 0x12, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0x8a, 0xf5, 0x73, 0x33, 0xd3, 0x8b, 0x12, 0x4b, 0x52, 0x21, 0xc6, 0x29, 0x71, 0x70,
	0xb1, 0x39, 0x83, 0x8d, 0x72, 0x6a, 0x63, 0xfc, 0x34, 0xe3, 0x5f, 0x3f, 0xab, 0x89, 0x90, 0x11,
	0xc4, 0x86, 0xd4, 0x8a, 0x92, 0xd4, 0xbc, 0x62, 0x90, 0x16, 0xa8, 0x2d, 0xc5, 0x38, 0xad, 0x31,
	0xe6, 0xb2, 0xce, 0xcc, 0xd7, 0x03, 0x6b, 0x2b, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0xc9, 0x8d,
	0x4e, 0xdc, 0x10, 0x17, 0x04, 0x80, 0x1c, 0x14, 0xc0, 0x98, 0xc4, 0x06, 0x76, 0x99, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x79, 0xda, 0xfd, 0xd7, 0x26, 0x01, 0x00, 0x00,
}
