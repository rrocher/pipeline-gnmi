// Code generated by protoc-gen-go.
// source: isis_sh_te_tunnel.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_te_tunnels_te_tunnel is a generated protocol buffer package.

It is generated from these files:
	isis_sh_te_tunnel.proto

It has these top-level messages:
	IsisShTeTunnel_KEYS
	IsisShTeTunnel
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_te_tunnels_te_tunnel

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

// MPLS TE tunnel
type IsisShTeTunnel_KEYS struct {
	InstanceName  string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName        string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName       string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	TopologyName  string `protobuf:"bytes,4,opt,name=topology_name,json=topologyName" json:"topology_name,omitempty"`
	Level         string `protobuf:"bytes,5,opt,name=level" json:"level,omitempty"`
	SystemId      string `protobuf:"bytes,6,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
	InterfaceName string `protobuf:"bytes,7,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *IsisShTeTunnel_KEYS) Reset()                    { *m = IsisShTeTunnel_KEYS{} }
func (m *IsisShTeTunnel_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeTunnel_KEYS) ProtoMessage()               {}
func (*IsisShTeTunnel_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShTeTunnel_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *IsisShTeTunnel_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IsisShTeTunnel struct {
	// Destination system ID
	TeSystemId string `protobuf:"bytes,50,opt,name=te_system_id,json=teSystemId" json:"te_system_id,omitempty"`
	// Tunnel interface
	TeInterface string `protobuf:"bytes,51,opt,name=te_interface,json=teInterface" json:"te_interface,omitempty"`
	// Tunnel bandwidth
	TeBandwidth uint32 `protobuf:"varint,52,opt,name=te_bandwidth,json=teBandwidth" json:"te_bandwidth,omitempty"`
	// Tunnel metric
	TeigpMetric int32 `protobuf:"zigzag32,53,opt,name=teigp_metric,json=teigpMetric" json:"teigp_metric,omitempty"`
	// Tunnel next-hop IP address
	TeNextHopIpAddress string `protobuf:"bytes,54,opt,name=te_next_hop_ip_address,json=teNextHopIpAddress" json:"te_next_hop_ip_address,omitempty"`
	// Tunnel metric mode
	TeModeType string `protobuf:"bytes,55,opt,name=te_mode_type,json=teModeType" json:"te_mode_type,omitempty"`
	// Indicates whether MPLS TE IPv4 forwarding adjacency is enabled
	TeIpv4FaEnabled bool `protobuf:"varint,56,opt,name=te_ipv4_fa_enabled,json=teIpv4FaEnabled" json:"te_ipv4_fa_enabled,omitempty"`
	// Indicates whether MPLS TE IPv6 forwarding adjacency is enabled
	TeIpv6FaEnabled bool `protobuf:"varint,57,opt,name=te_ipv6_fa_enabled,json=teIpv6FaEnabled" json:"te_ipv6_fa_enabled,omitempty"`
	// Indicates whether MPLS TE IPv4 autoroute announce is enabled
	TeIpv4AaEnabled bool `protobuf:"varint,58,opt,name=te_ipv4_aa_enabled,json=teIpv4AaEnabled" json:"te_ipv4_aa_enabled,omitempty"`
	// Indicates whether MPLS TE IPv6 autoroute announce is enabled
	TeIpv6AaEnabled bool `protobuf:"varint,59,opt,name=te_ipv6_aa_enabled,json=teIpv6AaEnabled" json:"te_ipv6_aa_enabled,omitempty"`
	// Tunnel checkpoint object ID
	TeCheckpointObjectId uint32 `protobuf:"varint,60,opt,name=te_checkpoint_object_id,json=teCheckpointObjectId" json:"te_checkpoint_object_id,omitempty"`
	// Indicates whether MPLS TE segment routing is enabled
	TeSegmentRoutingEnabled bool `protobuf:"varint,61,opt,name=te_segment_routing_enabled,json=teSegmentRoutingEnabled" json:"te_segment_routing_enabled,omitempty"`
}

func (m *IsisShTeTunnel) Reset()                    { *m = IsisShTeTunnel{} }
func (m *IsisShTeTunnel) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeTunnel) ProtoMessage()               {}
func (*IsisShTeTunnel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShTeTunnel) GetTeSystemId() string {
	if m != nil {
		return m.TeSystemId
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeInterface() string {
	if m != nil {
		return m.TeInterface
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeBandwidth() uint32 {
	if m != nil {
		return m.TeBandwidth
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeigpMetric() int32 {
	if m != nil {
		return m.TeigpMetric
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeNextHopIpAddress() string {
	if m != nil {
		return m.TeNextHopIpAddress
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeModeType() string {
	if m != nil {
		return m.TeModeType
	}
	return ""
}

func (m *IsisShTeTunnel) GetTeIpv4FaEnabled() bool {
	if m != nil {
		return m.TeIpv4FaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeIpv6FaEnabled() bool {
	if m != nil {
		return m.TeIpv6FaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeIpv4AaEnabled() bool {
	if m != nil {
		return m.TeIpv4AaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeIpv6AaEnabled() bool {
	if m != nil {
		return m.TeIpv6AaEnabled
	}
	return false
}

func (m *IsisShTeTunnel) GetTeCheckpointObjectId() uint32 {
	if m != nil {
		return m.TeCheckpointObjectId
	}
	return 0
}

func (m *IsisShTeTunnel) GetTeSegmentRoutingEnabled() bool {
	if m != nil {
		return m.TeSegmentRoutingEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*IsisShTeTunnel_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_tunnels.te_tunnel.isis_sh_te_tunnel_KEYS")
	proto.RegisterType((*IsisShTeTunnel)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_tunnels.te_tunnel.isis_sh_te_tunnel")
}

func init() { proto.RegisterFile("isis_sh_te_tunnel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0x95, 0xf7, 0xdd, 0xda, 0xce, 0xac, 0xa0, 0x59, 0xd3, 0x1a, 0xe0, 0xa6, 0x6c, 0x42,
	0xaa, 0x84, 0xd4, 0x8b, 0xad, 0x2b, 0x7f, 0x06, 0x17, 0x03, 0x0d, 0x51, 0xa1, 0x0d, 0xa9, 0xe5,
	0x86, 0xab, 0x23, 0x37, 0x3e, 0x6d, 0x0d, 0x89, 0x6d, 0xc5, 0x67, 0xa5, 0x15, 0x5f, 0x85, 0x0f,
	0xc9, 0x47, 0x40, 0x71, 0xea, 0xa4, 0x62, 0xbb, 0x3b, 0x7e, 0x9e, 0x5f, 0x9e, 0xf3, 0x24, 0x56,
	0x58, 0x47, 0x39, 0xe5, 0xc0, 0x2d, 0x80, 0x10, 0xe8, 0x56, 0x6b, 0x4c, 0xfb, 0x36, 0x37, 0x64,
	0xf8, 0xaf, 0x44, 0xb9, 0xc4, 0x80, 0x32, 0x0e, 0x56, 0x39, 0x24, 0xa9, 0x76, 0xe0, 0x51, 0x63,
	0x31, 0xef, 0x17, 0x53, 0x5f, 0x69, 0x47, 0x42, 0x27, 0x58, 0x4f, 0x7d, 0x32, 0xd6, 0xa4, 0x66,
	0xae, 0xd0, 0x85, 0x71, 0x5d, 0x0d, 0x90, 0xe2, 0x12, 0x53, 0xf7, 0xcf, 0xb9, 0x5f, 0xed, 0x75,
	0xf5, 0x78, 0xfc, 0x27, 0x62, 0x47, 0x77, 0x8a, 0xc1, 0xe7, 0xab, 0x6f, 0x13, 0x7e, 0xc2, 0xda,
	0x61, 0x1d, 0x68, 0x91, 0x61, 0x1c, 0x75, 0xa3, 0xde, 0xde, 0x78, 0x3f, 0x88, 0x37, 0x22, 0x43,
	0xde, 0x61, 0x4d, 0x31, 0x2b, 0xed, 0xff, 0xbc, 0xdd, 0x10, 0x33, 0x6f, 0x3c, 0x66, 0x2d, 0x17,
	0x9c, 0xff, 0xbd, 0xd3, 0x74, 0x1b, 0xeb, 0x84, 0xb5, 0xab, 0x6a, 0xde, 0xdf, 0x29, 0x83, 0x83,
	0xe8, 0xa1, 0x43, 0xb6, 0xeb, 0x6b, 0xc7, 0xbb, 0xde, 0x2c, 0x0f, 0xfc, 0x29, 0xdb, 0x73, 0x6b,
	0x47, 0x98, 0x81, 0x92, 0x71, 0xc3, 0x3b, 0xad, 0x52, 0x18, 0x49, 0xfe, 0x9c, 0x3d, 0x54, 0x9a,
	0x30, 0x9f, 0x89, 0xd0, 0xb8, 0xe9, 0x89, 0x76, 0xa5, 0x16, 0xc9, 0xc7, 0xbf, 0x77, 0xd8, 0xc1,
	0x9d, 0x57, 0xe6, 0x5d, 0xb6, 0x4f, 0x08, 0x75, 0xf8, 0xa9, 0x7f, 0x94, 0x11, 0x4e, 0x42, 0xfc,
	0x33, 0x4f, 0x54, 0x59, 0xf1, 0x99, 0x27, 0x1e, 0x10, 0x8e, 0x82, 0xb4, 0x41, 0xa6, 0x42, 0xcb,
	0x9f, 0x4a, 0xd2, 0x22, 0x1e, 0x74, 0xa3, 0x5e, 0xbb, 0x40, 0xde, 0x07, 0xa9, 0x44, 0xd4, 0xdc,
	0x42, 0x86, 0x94, 0xab, 0x24, 0x3e, 0xef, 0x46, 0xbd, 0x83, 0x02, 0x51, 0x73, 0x7b, 0xed, 0x25,
	0x7e, 0xca, 0x8e, 0x08, 0x41, 0xe3, 0x8a, 0x60, 0x61, 0x2c, 0x28, 0x0b, 0x42, 0xca, 0x1c, 0x9d,
	0x8b, 0x87, 0x7e, 0x25, 0x27, 0xbc, 0xc1, 0x15, 0x7d, 0x32, 0x76, 0x64, 0x2f, 0x4b, 0x67, 0x53,
	0x3f, 0x33, 0x12, 0x81, 0xd6, 0x16, 0xe3, 0x97, 0xa1, 0xfe, 0xb5, 0x91, 0xf8, 0x75, 0x6d, 0x91,
	0xbf, 0x60, 0xbc, 0xa8, 0x6f, 0x97, 0x03, 0x98, 0x09, 0x40, 0x2d, 0xa6, 0x29, 0xca, 0xf8, 0x55,
	0x37, 0xea, 0xb5, 0xc6, 0x8f, 0x08, 0x47, 0x76, 0x39, 0xf8, 0x28, 0xae, 0x4a, 0xb9, 0x86, 0x87,
	0xdb, 0xf0, 0xeb, 0x2d, 0x78, 0x78, 0x0f, 0x3c, 0x00, 0x51, 0xc3, 0x6f, 0xb6, 0x93, 0x2f, 0xef,
	0x4b, 0xde, 0x82, 0x2f, 0xb6, 0x93, 0x6b, 0xf8, 0x9c, 0x75, 0x08, 0x21, 0x59, 0x60, 0xf2, 0xc3,
	0x1a, 0xa5, 0x09, 0xcc, 0xf4, 0x3b, 0x26, 0x54, 0xdc, 0xcf, 0x5b, 0xff, 0x69, 0x0f, 0x09, 0x3f,
	0x54, 0xee, 0x17, 0x6f, 0x8e, 0x24, 0xbf, 0x60, 0x4f, 0x8a, 0xbb, 0xc4, 0x79, 0x86, 0x9a, 0x20,
	0x37, 0xb7, 0xa4, 0xf4, 0xbc, 0xda, 0xf5, 0xce, 0xef, 0xea, 0x10, 0x4e, 0x4a, 0x60, 0x5c, 0xfa,
	0x9b, 0x9d, 0xd3, 0x86, 0xff, 0x2b, 0xcf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x18, 0xb8, 0xb1,
	0xc9, 0xb0, 0x03, 0x00, 0x00,
}
