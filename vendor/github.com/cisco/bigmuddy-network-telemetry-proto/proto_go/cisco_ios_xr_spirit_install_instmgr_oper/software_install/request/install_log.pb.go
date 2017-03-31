// Code generated by protoc-gen-go.
// source: install_log.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_spirit_install_instmgr_oper_software_install_request is a generated protocol buffer package.

It is generated from these files:
	install_log.proto

It has these top-level messages:
	InstallLog_KEYS
	InstallLog
*/
package cisco_ios_xr_spirit_install_instmgr_oper_software_install_request

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

// Install Log
type InstallLog_KEYS struct {
}

func (m *InstallLog_KEYS) Reset()                    { *m = InstallLog_KEYS{} }
func (m *InstallLog_KEYS) String() string            { return proto.CompactTextString(m) }
func (*InstallLog_KEYS) ProtoMessage()               {}
func (*InstallLog_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type InstallLog struct {
	// log
	Log string `protobuf:"bytes,50,opt,name=log" json:"log,omitempty"`
}

func (m *InstallLog) Reset()                    { *m = InstallLog{} }
func (m *InstallLog) String() string            { return proto.CompactTextString(m) }
func (*InstallLog) ProtoMessage()               {}
func (*InstallLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InstallLog) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func init() {
	proto.RegisterType((*InstallLog_KEYS)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.request.install_log_KEYS")
	proto.RegisterType((*InstallLog)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.request.install_log")
}

func init() { proto.RegisterFile("install_log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xcc, 0x2b, 0x2e,
	0x49, 0xcc, 0xc9, 0x89, 0xcf, 0xc9, 0x4f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x72, 0x4c,
	0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x2f, 0x2e, 0xc8, 0x2c,
	0xca, 0x2c, 0x89, 0x87, 0x29, 0x03, 0xd1, 0xb9, 0xe9, 0x45, 0xf1, 0xf9, 0x05, 0xa9, 0x45, 0x7a,
	0xc5, 0xf9, 0x69, 0x25, 0xe5, 0x89, 0x45, 0xa9, 0x30, 0x59, 0xbd, 0xa2, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x25, 0x21, 0x2e, 0x01, 0x24, 0x73, 0xe3, 0xbd, 0x5d, 0x23, 0x83, 0x95, 0xe4, 0xb9,
	0xb8, 0x91, 0xc4, 0x84, 0x04, 0xb8, 0x98, 0x73, 0xf2, 0xd3, 0x25, 0x8c, 0x14, 0x18, 0x35, 0x38,
	0x83, 0x40, 0xcc, 0x24, 0x36, 0xb0, 0xf5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xca,
	0x9e, 0x4b, 0x93, 0x00, 0x00, 0x00,
}
