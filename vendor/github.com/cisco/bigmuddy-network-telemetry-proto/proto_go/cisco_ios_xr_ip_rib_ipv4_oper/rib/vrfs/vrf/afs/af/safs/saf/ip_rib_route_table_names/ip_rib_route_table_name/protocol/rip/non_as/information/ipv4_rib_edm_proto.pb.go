// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_proto.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_rip_non_as_information is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_proto.proto

It has these top-level messages:
	Ipv4RibEdmProto_KEYS
	Ipv4RibEdmProto
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_rip_non_as_information

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

// Information of a rib protocol
type Ipv4RibEdmProto_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
}

func (m *Ipv4RibEdmProto_KEYS) Reset()                    { *m = Ipv4RibEdmProto_KEYS{} }
func (m *Ipv4RibEdmProto_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

type Ipv4RibEdmProto struct {
	// Name
	ProtocolNames string `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames" json:"protocol_names,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance" json:"instance,omitempty"`
	// Proto version
	Version uint32 `protobuf:"varint,52,opt,name=version" json:"version,omitempty"`
	// Number of redist clients
	RedistributionClientCount uint32 `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount" json:"redistribution_client_count,omitempty"`
	// Number of proto clients
	ProtocolClientsCount uint32 `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount" json:"protocol_clients_count,omitempty"`
	// Number of routes (including active, backup and deleted), where, number of backup routes = RoutesCounts - ActiveRoutesCount - DeletedRoutesCount
	RoutesCounts uint32 `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts" json:"routes_counts,omitempty"`
	// Number of active routes (not deleted)
	ActiveRoutesCount uint32 `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount" json:"active_routes_count,omitempty"`
	// Number of deleted routes
	DeletedRoutesCount uint32 `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount" json:"deleted_routes_count,omitempty"`
	// Number of paths for all routes
	PathsCount uint32 `protobuf:"varint,58,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// Memory for proto's routes and paths in bytes
	ProtocolRouteMemory uint32 `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory" json:"protocol_route_memory,omitempty"`
	// Number of backup routes
	BackupRoutesCount uint32 `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount" json:"backup_routes_count,omitempty"`
}

func (m *Ipv4RibEdmProto) Reset()                    { *m = Ipv4RibEdmProto{} }
func (m *Ipv4RibEdmProto) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto) ProtoMessage()               {}
func (*Ipv4RibEdmProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rip.non_as.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.rip.non_as.information.ipv4_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x14, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42, 0xca,
	0xc9, 0x42, 0x6d, 0xf8, 0x8f, 0xb8, 0x54, 0x9c, 0x10, 0x1c, 0x02, 0x17, 0x4e, 0x96, 0xd7, 0xf1,
	0x0a, 0x8b, 0x5d, 0x7b, 0xe5, 0xf1, 0xae, 0xe0, 0x19, 0xb8, 0xf1, 0xaa, 0xbc, 0x00, 0xf2, 0x78,
	0x37, 0x24, 0x44, 0x5c, 0x1c, 0x79, 0xbe, 0xdf, 0xcf, 0xf9, 0x26, 0x01, 0x6e, 0xdb, 0x7e, 0x29,
	0x83, 0x2d, 0xa5, 0x59, 0x37, 0xb2, 0x0d, 0x3e, 0x7a, 0x41, 0x27, 0xfb, 0x59, 0x68, 0x8b, 0xda,
	0x4b, 0xeb, 0x51, 0x7e, 0x0f, 0xd2, 0xb6, 0x44, 0x11, 0xee, 0x5b, 0x13, 0x44, 0xb0, 0xa5, 0xe8,
	0x43, 0x85, 0xe9, 0x10, 0xaa, 0x42, 0xa1, 0x2a, 0x81, 0xe9, 0x13, 0x55, 0x25, 0x06, 0x3a, 0xf8,
	0x2e, 0x1a, 0x19, 0x55, 0x59, 0x1b, 0xe9, 0x54, 0x63, 0xf0, 0x7f, 0x41, 0xfe, 0x4e, 0xed, 0x6b,
	0x11, 0x6c, 0x2b, 0x9c, 0x77, 0x52, 0xa1, 0xb0, 0xae, 0xf2, 0xa1, 0x51, 0xd1, 0x7a, 0x77, 0xfe,
	0xab, 0x80, 0xd3, 0xfd, 0xaa, 0xf2, 0xfd, 0xbb, 0x2f, 0x9f, 0xd8, 0x1c, 0xa6, 0x7d, 0xa8, 0xe8,
	0x1d, 0x5e, 0x9c, 0x15, 0x8b, 0x1b, 0xab, 0x49, 0x1f, 0xaa, 0x8f, 0xaa, 0x31, 0xec, 0x14, 0x26,
	0x6a, 0x48, 0xae, 0x51, 0x72, 0xa8, 0x72, 0x30, 0x87, 0x29, 0x8e, 0xc9, 0x41, 0x76, 0x70, 0x88,
	0x16, 0x70, 0xfb, 0xdf, 0x7a, 0xfc, 0x3a, 0x21, 0x47, 0x34, 0xff, 0x9c, 0xc6, 0x89, 0x3c, 0xff,
	0x7d, 0x00, 0x6c, 0xbf, 0x14, 0x7b, 0x0c, 0x47, 0xe3, 0x3a, 0x79, 0x6b, 0x7e, 0x41, 0xfa, 0x6c,
	0x9c, 0x26, 0x19, 0xd9, 0x7d, 0x98, 0x5a, 0x87, 0x51, 0x39, 0x6d, 0xf8, 0x25, 0x01, 0x9b, 0x3b,
	0xe3, 0x30, 0xe9, 0x4d, 0x40, 0xeb, 0x1d, 0x5f, 0x9e, 0x15, 0x8b, 0xd9, 0x6a, 0xbc, 0xb2, 0xb7,
	0xf0, 0x20, 0x98, 0xb5, 0xc5, 0x18, 0x6c, 0xd9, 0xa5, 0x9f, 0x46, 0xea, 0xda, 0x1a, 0x17, 0xa5,
	0xf6, 0x9d, 0x8b, 0xfc, 0x29, 0xd1, 0xf3, 0x5d, 0xe4, 0x8a, 0x88, 0xab, 0x04, 0xb0, 0x25, 0xdc,
	0xdb, 0x94, 0xcb, 0x26, 0x0e, 0xea, 0x33, 0x52, 0x4f, 0xc6, 0x34, 0x4b, 0x98, 0xad, 0x47, 0x30,
	0xa3, 0xdd, 0x07, 0x16, 0xf9, 0x73, 0x82, 0x6f, 0xe5, 0x21, 0x31, 0xc8, 0x04, 0x1c, 0x2b, 0x1d,
	0x6d, 0x6f, 0xe4, 0x36, 0xcb, 0x5f, 0x10, 0x7a, 0x27, 0x47, 0xab, 0xbf, 0x02, 0x7b, 0x02, 0x27,
	0x6b, 0x53, 0x9b, 0x68, 0xd6, 0xbb, 0xc2, 0x4b, 0x12, 0xd8, 0x90, 0x6d, 0x1b, 0x0f, 0xe1, 0x66,
	0xab, 0xe2, 0xd7, 0x11, 0x7c, 0x45, 0x20, 0xd0, 0x28, 0x03, 0x17, 0x70, 0x77, 0xb3, 0x5d, 0xfe,
	0x13, 0x1b, 0xd3, 0xf8, 0xf0, 0x83, 0xbf, 0x26, 0xf4, 0x78, 0x0c, 0xe9, 0xd1, 0x0f, 0x14, 0xa5,
	0xda, 0xa5, 0xd2, 0xdf, 0xba, 0x76, 0xb7, 0xc5, 0x9b, 0x5c, 0x3b, 0x47, 0x5b, 0x25, 0xca, 0x43,
	0x7a, 0xe4, 0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xef, 0x63, 0x36, 0x3b, 0x03, 0x00,
	0x00,
}