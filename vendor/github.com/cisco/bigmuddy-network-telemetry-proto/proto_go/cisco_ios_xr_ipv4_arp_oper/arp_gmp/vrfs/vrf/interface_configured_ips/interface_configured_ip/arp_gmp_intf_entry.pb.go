// Code generated by protoc-gen-go.
// source: arp_gmp_intf_entry.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_arp_oper_arp_gmp_vrfs_vrf_interface_configured_ips_interface_configured_ip is a generated protocol buffer package.

It is generated from these files:
	arp_gmp_intf_entry.proto

It has these top-level messages:
	ArpGmpIntfEntry_KEYS
	ArpGmpIntfEntry
	ArpGmpConfigEntry
*/
package cisco_ios_xr_ipv4_arp_oper_arp_gmp_vrfs_vrf_interface_configured_ips_interface_configured_ip

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

// ARP GMP entries associated with an interface
type ArpGmpIntfEntry_KEYS struct {
	VrfName       string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	InterfaceName string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	Address       string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
}

func (m *ArpGmpIntfEntry_KEYS) Reset()                    { *m = ArpGmpIntfEntry_KEYS{} }
func (m *ArpGmpIntfEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*ArpGmpIntfEntry_KEYS) ProtoMessage()               {}
func (*ArpGmpIntfEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ArpGmpIntfEntry_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *ArpGmpIntfEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *ArpGmpIntfEntry_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ArpGmpIntfEntry struct {
	// Interface name
	InterfaceName string `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Route reference count
	ReferenceCount uint32 `protobuf:"varint,51,opt,name=reference_count,json=referenceCount" json:"reference_count,omitempty"`
	// Associated configuration entry
	AssociatedConfigurationEntry *ArpGmpConfigEntry `protobuf:"bytes,52,opt,name=associated_configuration_entry,json=associatedConfigurationEntry" json:"associated_configuration_entry,omitempty"`
}

func (m *ArpGmpIntfEntry) Reset()                    { *m = ArpGmpIntfEntry{} }
func (m *ArpGmpIntfEntry) String() string            { return proto.CompactTextString(m) }
func (*ArpGmpIntfEntry) ProtoMessage()               {}
func (*ArpGmpIntfEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ArpGmpIntfEntry) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *ArpGmpIntfEntry) GetReferenceCount() uint32 {
	if m != nil {
		return m.ReferenceCount
	}
	return 0
}

func (m *ArpGmpIntfEntry) GetAssociatedConfigurationEntry() *ArpGmpConfigEntry {
	if m != nil {
		return m.AssociatedConfigurationEntry
	}
	return nil
}

// ARP GMP configuration entry
type ArpGmpConfigEntry struct {
	// IP address
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
	// Hardware address
	HardwareAddress string `protobuf:"bytes,2,opt,name=hardware_address,json=hardwareAddress" json:"hardware_address,omitempty"`
	// Encap type
	EncapsulationType string `protobuf:"bytes,3,opt,name=encapsulation_type,json=encapsulationType" json:"encapsulation_type,omitempty"`
	// Entry type static/alias
	EntryType string `protobuf:"bytes,4,opt,name=entry_type,json=entryType" json:"entry_type,omitempty"`
}

func (m *ArpGmpConfigEntry) Reset()                    { *m = ArpGmpConfigEntry{} }
func (m *ArpGmpConfigEntry) String() string            { return proto.CompactTextString(m) }
func (*ArpGmpConfigEntry) ProtoMessage()               {}
func (*ArpGmpConfigEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ArpGmpConfigEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ArpGmpConfigEntry) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

func (m *ArpGmpConfigEntry) GetEncapsulationType() string {
	if m != nil {
		return m.EncapsulationType
	}
	return ""
}

func (m *ArpGmpConfigEntry) GetEntryType() string {
	if m != nil {
		return m.EntryType
	}
	return ""
}

func init() {
	proto.RegisterType((*ArpGmpIntfEntry_KEYS)(nil), "cisco_ios_xr_ipv4_arp_oper.arp_gmp.vrfs.vrf.interface_configured_ips.interface_configured_ip.arp_gmp_intf_entry_KEYS")
	proto.RegisterType((*ArpGmpIntfEntry)(nil), "cisco_ios_xr_ipv4_arp_oper.arp_gmp.vrfs.vrf.interface_configured_ips.interface_configured_ip.arp_gmp_intf_entry")
	proto.RegisterType((*ArpGmpConfigEntry)(nil), "cisco_ios_xr_ipv4_arp_oper.arp_gmp.vrfs.vrf.interface_configured_ips.interface_configured_ip.arp_gmp_config_entry")
}

func init() { proto.RegisterFile("arp_gmp_intf_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0x4a, 0x2b, 0x31,
	0x18, 0x65, 0x7a, 0x2f, 0xb7, 0xb7, 0xb9, 0xb4, 0xbd, 0x06, 0xc1, 0x11, 0x54, 0x4a, 0x41, 0xac,
	0x0b, 0x67, 0xd1, 0xf6, 0x05, 0xa4, 0x74, 0x25, 0xb8, 0xa8, 0x6e, 0x04, 0x21, 0xc4, 0x99, 0x2f,
	0x35, 0x60, 0x93, 0xf0, 0x25, 0x33, 0xda, 0xa7, 0xf1, 0x0d, 0x5c, 0xfb, 0x78, 0x32, 0xc9, 0xcc,
	0x54, 0x69, 0x5d, 0xba, 0x19, 0x98, 0x73, 0x4e, 0xce, 0xf9, 0xfe, 0x48, 0xcc, 0xd1, 0xb0, 0xe5,
	0xca, 0x30, 0xa9, 0x9c, 0x60, 0xa0, 0x1c, 0xae, 0x13, 0x83, 0xda, 0x69, 0x7a, 0x9f, 0x4a, 0x9b,
	0x6a, 0x26, 0xb5, 0x65, 0x2f, 0xc8, 0xa4, 0x29, 0xa6, 0xac, 0xd4, 0x6a, 0x03, 0x98, 0x54, 0x8f,
	0x92, 0x02, 0x85, 0x2d, 0x3f, 0x89, 0x54, 0x0e, 0x50, 0xf0, 0x14, 0x58, 0xaa, 0x95, 0x90, 0xcb,
	0x1c, 0x21, 0x63, 0xd2, 0xd8, 0xef, 0x88, 0x61, 0x4e, 0x0e, 0xb6, 0x93, 0xd9, 0xd5, 0xfc, 0xee,
	0x86, 0x1e, 0x92, 0xbf, 0x05, 0x0a, 0xa6, 0xf8, 0x0a, 0xe2, 0x68, 0x10, 0x8d, 0x3a, 0x8b, 0x76,
	0x81, 0xe2, 0x9a, 0xaf, 0x80, 0x9e, 0x92, 0xde, 0xc6, 0xd0, 0x0b, 0x5a, 0x5e, 0xd0, 0x6d, 0x50,
	0x2f, 0x8b, 0x49, 0x9b, 0x67, 0x19, 0x82, 0xb5, 0xf1, 0xaf, 0x60, 0x50, 0xfd, 0x0e, 0x5f, 0x5b,
	0x84, 0x6e, 0xe7, 0xee, 0xf0, 0x1d, 0xef, 0xf2, 0x3d, 0x23, 0x7d, 0x04, 0x01, 0x08, 0xca, 0xf7,
	0x93, 0x2b, 0x17, 0x4f, 0x06, 0xd1, 0xa8, 0xbb, 0xe8, 0x35, 0xf0, 0xac, 0x44, 0xe9, 0x7b, 0x44,
	0x4e, 0xb8, 0xb5, 0x3a, 0x95, 0xdc, 0x41, 0xd6, 0xb4, 0xce, 0x9d, 0xd4, 0x2a, 0x44, 0xc6, 0xd3,
	0x41, 0x34, 0xfa, 0x37, 0xc6, 0xe4, 0x27, 0xa7, 0x5c, 0x3b, 0x54, 0x68, 0x48, 0x5e, 0x1c, 0x6d,
	0x2a, 0x9b, 0x7d, 0x2e, 0x6c, 0x5e, 0xb2, 0xc3, 0xb7, 0x88, 0xec, 0xef, 0x7a, 0x46, 0x8f, 0x09,
	0x91, 0x86, 0xd5, 0x73, 0x0d, 0x8b, 0xe9, 0x48, 0x73, 0x19, 0x00, 0x7a, 0x4e, 0xfe, 0x3f, 0x72,
	0xcc, 0x9e, 0x39, 0x42, 0x23, 0x0a, 0xcb, 0xe9, 0xd7, 0x78, 0x2d, 0xbd, 0x20, 0x14, 0x54, 0xca,
	0x8d, 0xcd, 0x9f, 0xc2, 0x44, 0xdc, 0xda, 0x40, 0xb5, 0xa9, 0xbd, 0x2f, 0xcc, 0xed, 0xda, 0x40,
	0x19, 0x1c, 0xae, 0xc3, 0xcb, 0x7e, 0x87, 0x60, 0x8f, 0x94, 0xf4, 0xc3, 0x1f, 0x7f, 0xae, 0x93,
	0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x1b, 0x98, 0xad, 0xca, 0x02, 0x00, 0x00,
}
