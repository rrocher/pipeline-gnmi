// Code generated by protoc-gen-go.
// source: cdp_stats.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_cdp_oper_cdp_nodes_node_statistics is a generated protocol buffer package.

It is generated from these files:
	cdp_stats.proto

It has these top-level messages:
	CdpStats_KEYS
	CdpStats
*/
package cisco_ios_xr_cdp_oper_cdp_nodes_node_statistics

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

// CDP statistics
type CdpStats_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *CdpStats_KEYS) Reset()                    { *m = CdpStats_KEYS{} }
func (m *CdpStats_KEYS) String() string            { return proto.CompactTextString(m) }
func (*CdpStats_KEYS) ProtoMessage()               {}
func (*CdpStats_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CdpStats_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type CdpStats struct {
	// Received packets
	ReceivedPackets uint32 `protobuf:"varint,50,opt,name=received_packets,json=receivedPackets" json:"received_packets,omitempty"`
	// Received v1 packets
	ReceivedPacketsV1 uint32 `protobuf:"varint,51,opt,name=received_packets_v1,json=receivedPacketsV1" json:"received_packets_v1,omitempty"`
	// Received v2 packets
	ReceivedPacketsV2 uint32 `protobuf:"varint,52,opt,name=received_packets_v2,json=receivedPacketsV2" json:"received_packets_v2,omitempty"`
	// Transmitted packets
	TransmittedPackets uint32 `protobuf:"varint,53,opt,name=transmitted_packets,json=transmittedPackets" json:"transmitted_packets,omitempty"`
	// Transmitted v1 packets
	TransmittedPacketsV1 uint32 `protobuf:"varint,54,opt,name=transmitted_packets_v1,json=transmittedPacketsV1" json:"transmitted_packets_v1,omitempty"`
	// Transmitted v2 packets
	TransmittedPacketsV2 uint32 `protobuf:"varint,55,opt,name=transmitted_packets_v2,json=transmittedPacketsV2" json:"transmitted_packets_v2,omitempty"`
	// Header syntax errors
	HeaderErrors uint32 `protobuf:"varint,56,opt,name=header_errors,json=headerErrors" json:"header_errors,omitempty"`
	// Checksum errors
	ChecksumErrors uint32 `protobuf:"varint,57,opt,name=checksum_errors,json=checksumErrors" json:"checksum_errors,omitempty"`
	// Transmission errors
	EncapsulationErrors uint32 `protobuf:"varint,58,opt,name=encapsulation_errors,json=encapsulationErrors" json:"encapsulation_errors,omitempty"`
	// Bad packet received and dropped
	BadPacketErrors uint32 `protobuf:"varint,59,opt,name=bad_packet_errors,json=badPacketErrors" json:"bad_packet_errors,omitempty"`
	// Out-of-memory conditions
	OutOfMemoryErrors uint32 `protobuf:"varint,60,opt,name=out_of_memory_errors,json=outOfMemoryErrors" json:"out_of_memory_errors,omitempty"`
	// Truncated messages
	TruncatedPacketErrors uint32 `protobuf:"varint,61,opt,name=truncated_packet_errors,json=truncatedPacketErrors" json:"truncated_packet_errors,omitempty"`
	// Can't handle receive version
	HeaderVersionErrors uint32 `protobuf:"varint,62,opt,name=header_version_errors,json=headerVersionErrors" json:"header_version_errors,omitempty"`
	// Cannot open file
	OpenFileErrors uint32 `protobuf:"varint,63,opt,name=open_file_errors,json=openFileErrors" json:"open_file_errors,omitempty"`
}

func (m *CdpStats) Reset()                    { *m = CdpStats{} }
func (m *CdpStats) String() string            { return proto.CompactTextString(m) }
func (*CdpStats) ProtoMessage()               {}
func (*CdpStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CdpStats) GetReceivedPackets() uint32 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *CdpStats) GetReceivedPacketsV1() uint32 {
	if m != nil {
		return m.ReceivedPacketsV1
	}
	return 0
}

func (m *CdpStats) GetReceivedPacketsV2() uint32 {
	if m != nil {
		return m.ReceivedPacketsV2
	}
	return 0
}

func (m *CdpStats) GetTransmittedPackets() uint32 {
	if m != nil {
		return m.TransmittedPackets
	}
	return 0
}

func (m *CdpStats) GetTransmittedPacketsV1() uint32 {
	if m != nil {
		return m.TransmittedPacketsV1
	}
	return 0
}

func (m *CdpStats) GetTransmittedPacketsV2() uint32 {
	if m != nil {
		return m.TransmittedPacketsV2
	}
	return 0
}

func (m *CdpStats) GetHeaderErrors() uint32 {
	if m != nil {
		return m.HeaderErrors
	}
	return 0
}

func (m *CdpStats) GetChecksumErrors() uint32 {
	if m != nil {
		return m.ChecksumErrors
	}
	return 0
}

func (m *CdpStats) GetEncapsulationErrors() uint32 {
	if m != nil {
		return m.EncapsulationErrors
	}
	return 0
}

func (m *CdpStats) GetBadPacketErrors() uint32 {
	if m != nil {
		return m.BadPacketErrors
	}
	return 0
}

func (m *CdpStats) GetOutOfMemoryErrors() uint32 {
	if m != nil {
		return m.OutOfMemoryErrors
	}
	return 0
}

func (m *CdpStats) GetTruncatedPacketErrors() uint32 {
	if m != nil {
		return m.TruncatedPacketErrors
	}
	return 0
}

func (m *CdpStats) GetHeaderVersionErrors() uint32 {
	if m != nil {
		return m.HeaderVersionErrors
	}
	return 0
}

func (m *CdpStats) GetOpenFileErrors() uint32 {
	if m != nil {
		return m.OpenFileErrors
	}
	return 0
}

func init() {
	proto.RegisterType((*CdpStats_KEYS)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.statistics.cdp_stats_KEYS")
	proto.RegisterType((*CdpStats)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.statistics.cdp_stats")
}

func init() { proto.RegisterFile("cdp_stats.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4d, 0xaf, 0xd2, 0x40,
	0x14, 0x86, 0xc3, 0x42, 0x23, 0x27, 0xf2, 0x35, 0x80, 0x36, 0x71, 0x43, 0x70, 0x21, 0x9a, 0x58,
	0xd2, 0x82, 0xf8, 0xad, 0x2b, 0xdc, 0x18, 0x3f, 0x82, 0x09, 0x89, 0xab, 0xc9, 0x30, 0x3d, 0x84,
	0x09, 0xb4, 0xd3, 0xcc, 0x4c, 0x89, 0xfe, 0xa2, 0xfb, 0x37, 0x6f, 0x3a, 0xd3, 0xe1, 0x02, 0x17,
	0x36, 0x5d, 0xbc, 0xef, 0xf3, 0xb4, 0x67, 0x4e, 0x33, 0xd0, 0xe2, 0x49, 0x4e, 0xb5, 0x61, 0x46,
	0x87, 0xb9, 0x92, 0x46, 0x92, 0x31, 0x17, 0x9a, 0x4b, 0x2a, 0xa4, 0xa6, 0xff, 0x14, 0x2d, 0x5b,
	0x99, 0xa3, 0x0a, 0x79, 0x92, 0x87, 0x99, 0x4c, 0x50, 0xdb, 0x67, 0x58, 0x1a, 0x42, 0x1b, 0xc1,
	0xf5, 0xf0, 0x35, 0x34, 0x0f, 0xef, 0xa0, 0xdf, 0xe7, 0x7f, 0xff, 0x90, 0x67, 0x50, 0x2f, 0x21,
	0x9a, 0xb1, 0x14, 0x83, 0xda, 0xa0, 0x36, 0xaa, 0x2f, 0x1e, 0x95, 0xc1, 0x4f, 0x96, 0xe2, 0xf0,
	0xe6, 0x01, 0xd4, 0x0f, 0x3c, 0x79, 0x09, 0x6d, 0x85, 0x1c, 0xc5, 0x1e, 0x13, 0x9a, 0x33, 0xbe,
	0x45, 0xa3, 0x83, 0x78, 0x50, 0x1b, 0x35, 0x16, 0x2d, 0x9f, 0xff, 0x76, 0x31, 0x09, 0xa1, 0x7b,
	0x8e, 0xd2, 0x7d, 0x14, 0x4c, 0x2c, 0xdd, 0x39, 0xa3, 0x97, 0xd1, 0x65, 0x3e, 0x0e, 0xa6, 0x97,
	0xf9, 0x98, 0x8c, 0xa1, 0x6b, 0x14, 0xcb, 0x74, 0x2a, 0x8c, 0x39, 0x9a, 0xe6, 0x8d, 0xe5, 0xc9,
	0x51, 0xe5, 0x07, 0x9a, 0xc2, 0x93, 0x0b, 0x42, 0x39, 0xd3, 0xcc, 0x3a, 0xbd, 0xfb, 0xce, 0x32,
	0xba, 0x6a, 0xc5, 0xc1, 0xdb, 0xab, 0x56, 0x4c, 0x9e, 0x43, 0x63, 0x83, 0x2c, 0x41, 0x45, 0x51,
	0x29, 0xa9, 0x74, 0xf0, 0xce, 0xc2, 0x8f, 0x5d, 0x38, 0xb7, 0x19, 0x79, 0x01, 0x2d, 0xbe, 0x41,
	0xbe, 0xd5, 0x45, 0xea, 0xb1, 0xf7, 0x16, 0x6b, 0xfa, 0xb8, 0x02, 0x23, 0xe8, 0x61, 0xc6, 0x59,
	0xae, 0x8b, 0x1d, 0x33, 0x42, 0x66, 0x9e, 0xfe, 0x60, 0xe9, 0xee, 0x49, 0x57, 0x29, 0xaf, 0xa0,
	0xb3, 0x62, 0x7e, 0x5c, 0xcf, 0x7f, 0x74, 0x7f, 0x6a, 0xc5, 0xaa, 0x49, 0x2b, 0x76, 0x0c, 0x3d,
	0x59, 0x18, 0x2a, 0xd7, 0x34, 0xc5, 0x54, 0xaa, 0xff, 0x1e, 0xff, 0xe4, 0x56, 0x2f, 0x0b, 0xf3,
	0x6b, 0xfd, 0xc3, 0x36, 0x95, 0x30, 0x83, 0xa7, 0x46, 0x15, 0x19, 0x67, 0x77, 0x1b, 0xf1, 0xce,
	0x67, 0xeb, 0xf4, 0x0f, 0xf5, 0xc9, 0x87, 0x62, 0xe8, 0x57, 0x5b, 0xd9, 0xa3, 0xd2, 0x47, 0x07,
	0xf9, 0xe2, 0x0e, 0xe2, 0xca, 0xa5, 0xeb, 0x2a, 0x67, 0x04, 0x6d, 0x99, 0x63, 0x46, 0xd7, 0x62,
	0x87, 0x1e, 0xff, 0xea, 0xb6, 0x54, 0xe6, 0xdf, 0xc4, 0x0e, 0x1d, 0xb9, 0x7a, 0x68, 0x2f, 0xc4,
	0xe4, 0x36, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x1c, 0x6d, 0x24, 0x23, 0x03, 0x00, 0x00,
}
