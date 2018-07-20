// Code generated by protoc-gen-go.
// source: ipv6_dhcpv6d_proxy_stats.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_proxy_statistics is a generated protocol buffer package.

It is generated from these files:
	ipv6_dhcpv6d_proxy_stats.proto

It has these top-level messages:
	Ipv6Dhcpv6DProxyStats_KEYS
	Ipv6Dhcpv6DProxyStats
	Ipv6Dhcpv6DProxyStatsItem
	Ipv6Dhcpv6DFilteredStats
*/
package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_proxy_statistics

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

// DHCPv6 proxy statistics
type Ipv6Dhcpv6DProxyStats_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *Ipv6Dhcpv6DProxyStats_KEYS) Reset()                    { *m = Ipv6Dhcpv6DProxyStats_KEYS{} }
func (m *Ipv6Dhcpv6DProxyStats_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyStats_KEYS) ProtoMessage()               {}
func (*Ipv6Dhcpv6DProxyStats_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6Dhcpv6DProxyStats_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Ipv6Dhcpv6DProxyStats struct {
	Ipv6Dhcpv6DProxyStats []*Ipv6Dhcpv6DProxyStatsItem `protobuf:"bytes,50,rep,name=ipv6_dhcpv6d_proxy_stats,json=ipv6Dhcpv6dProxyStats" json:"ipv6_dhcpv6d_proxy_stats,omitempty"`
}

func (m *Ipv6Dhcpv6DProxyStats) Reset()                    { *m = Ipv6Dhcpv6DProxyStats{} }
func (m *Ipv6Dhcpv6DProxyStats) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyStats) ProtoMessage()               {}
func (*Ipv6Dhcpv6DProxyStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6Dhcpv6DProxyStats) GetIpv6Dhcpv6DProxyStats() []*Ipv6Dhcpv6DProxyStatsItem {
	if m != nil {
		return m.Ipv6Dhcpv6DProxyStats
	}
	return nil
}

type Ipv6Dhcpv6DProxyStatsItem struct {
	// DHCPv6 L3 VRF name
	VrfName string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// Proxy statistics
	Statistics *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,2,opt,name=statistics" json:"statistics,omitempty"`
}

func (m *Ipv6Dhcpv6DProxyStatsItem) Reset()                    { *m = Ipv6Dhcpv6DProxyStatsItem{} }
func (m *Ipv6Dhcpv6DProxyStatsItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyStatsItem) ProtoMessage()               {}
func (*Ipv6Dhcpv6DProxyStatsItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv6Dhcpv6DProxyStatsItem) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyStatsItem) GetStatistics() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Statistics
	}
	return nil
}

// DHCPv6 filtered statistics
type Ipv6Dhcpv6DFilteredStats struct {
	// Received packets
	ReceivedPackets uint64 `protobuf:"varint,1,opt,name=received_packets,json=receivedPackets" json:"received_packets,omitempty"`
	// Transmitted packets
	TransmittedPackets uint64 `protobuf:"varint,2,opt,name=transmitted_packets,json=transmittedPackets" json:"transmitted_packets,omitempty"`
	// Dropped packets
	DroppedPackets uint64 `protobuf:"varint,3,opt,name=dropped_packets,json=droppedPackets" json:"dropped_packets,omitempty"`
}

func (m *Ipv6Dhcpv6DFilteredStats) Reset()                    { *m = Ipv6Dhcpv6DFilteredStats{} }
func (m *Ipv6Dhcpv6DFilteredStats) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DFilteredStats) ProtoMessage()               {}
func (*Ipv6Dhcpv6DFilteredStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv6Dhcpv6DFilteredStats) GetReceivedPackets() uint64 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *Ipv6Dhcpv6DFilteredStats) GetTransmittedPackets() uint64 {
	if m != nil {
		return m.TransmittedPackets
	}
	return 0
}

func (m *Ipv6Dhcpv6DFilteredStats) GetDroppedPackets() uint64 {
	if m != nil {
		return m.DroppedPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DProxyStats_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.statistics.ipv6_dhcpv6d_proxy_stats_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DProxyStats)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.statistics.ipv6_dhcpv6d_proxy_stats")
	proto.RegisterType((*Ipv6Dhcpv6DProxyStatsItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.statistics.ipv6_dhcpv6d_proxy_stats_item")
	proto.RegisterType((*Ipv6Dhcpv6DFilteredStats)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.statistics.ipv6_dhcpv6d_filtered_stats")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_proxy_stats.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x26, 0xba, 0x65, 0xe0, 0x24, 0x22, 0x54, 0x86, 0x32, 0x7a, 0x71, 0x5e, 0x22,
	0x54, 0xd8, 0xc9, 0xa3, 0x3b, 0x09, 0x32, 0xba, 0x93, 0xa7, 0xd0, 0x25, 0xaf, 0x18, 0xb4, 0x4d,
	0x48, 0x42, 0x9d, 0x57, 0xff, 0x0f, 0x2f, 0xfe, 0x19, 0x5e, 0xfc, 0xd7, 0x24, 0xe9, 0xa6, 0xbd,
	0xb4, 0x27, 0xbd, 0x14, 0xfa, 0xfd, 0xf1, 0xde, 0xfb, 0x40, 0xf0, 0xb9, 0xd4, 0xd5, 0x9c, 0x89,
	0x47, 0xae, 0xab, 0xb9, 0x60, 0xda, 0xa8, 0xcd, 0x2b, 0xb3, 0x2e, 0x73, 0x96, 0x6a, 0xa3, 0x9c,
	0x22, 0x0b, 0x2e, 0x2d, 0x57, 0x4c, 0x2a, 0xcb, 0x36, 0x86, 0x85, 0x70, 0x09, 0x2f, 0x3f, 0x05,
	0xa5, 0xc1, 0xd0, 0xfa, 0x87, 0x96, 0x4a, 0x80, 0x0d, 0x5f, 0x1a, 0xe6, 0x50, 0x3f, 0x47, 0x5a,
	0x27, 0xb9, 0x8d, 0x6f, 0xf0, 0x59, 0xdb, 0x22, 0x76, 0xb7, 0x78, 0x58, 0x91, 0x09, 0x1e, 0xfa,
	0x26, 0x2b, 0xb3, 0x02, 0x22, 0x34, 0x45, 0xb3, 0x61, 0x3a, 0xf0, 0xc2, 0x7d, 0x56, 0x40, 0xfc,
	0x89, 0x70, 0xd4, 0x56, 0x27, 0xef, 0x1d, 0x66, 0x94, 0x4c, 0xfb, 0xb3, 0x51, 0x22, 0xe8, 0x9f,
	0x50, 0xd0, 0x56, 0x04, 0xe9, 0xa0, 0x48, 0x4f, 0xbc, 0x7d, 0x5b, 0xbb, 0x4b, 0x6f, 0xae, 0xbc,
	0x17, 0x7f, 0xa1, 0x0e, 0x76, 0x5f, 0x24, 0xa7, 0x78, 0x50, 0x99, 0xbc, 0x89, 0x7e, 0x50, 0x99,
	0xdc, 0x93, 0x93, 0x37, 0x84, 0xf1, 0xef, 0x01, 0x51, 0x6f, 0x8a, 0x66, 0xa3, 0x64, 0xfd, 0x1f,
	0x38, 0xb9, 0x7c, 0x76, 0x60, 0x40, 0xd4, 0x87, 0xa5, 0x8d, 0xad, 0xf1, 0x07, 0xc2, 0x93, 0x8e,
	0x2c, 0xb9, 0xc4, 0x47, 0x06, 0x38, 0xc8, 0x0a, 0x04, 0xd3, 0x19, 0x7f, 0x02, 0x67, 0x03, 0xc7,
	0x5e, 0x3a, 0xde, 0xe9, 0xcb, 0x5a, 0x26, 0x57, 0xf8, 0xd8, 0x99, 0xac, 0xb4, 0x85, 0x74, 0xae,
	0x91, 0xee, 0x85, 0x34, 0x69, 0x58, 0xbb, 0xc2, 0x05, 0x1e, 0x0b, 0xa3, 0xb4, 0x6e, 0x84, 0xfb,
	0x21, 0x7c, 0xb8, 0x95, 0xb7, 0xc1, 0xf5, 0x7e, 0x78, 0xaf, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x4a, 0x15, 0xb4, 0xed, 0xd1, 0x02, 0x00, 0x00,
}
