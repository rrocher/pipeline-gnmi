// Code generated by protoc-gen-go.
// source: ospfv3_edm_protocol.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_summary_protocol is a generated protocol buffer package.

It is generated from these files:
	ospfv3_edm_protocol.proto

It has these top-level messages:
	Ospfv3EdmProtocol_KEYS
	Ospfv3EdmProtocol
*/
package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_summary_protocol

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

// OSPFv3 Protocol Information
type Ospfv3EdmProtocol_KEYS struct {
	ProcessName string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName     string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
}

func (m *Ospfv3EdmProtocol_KEYS) Reset()                    { *m = Ospfv3EdmProtocol_KEYS{} }
func (m *Ospfv3EdmProtocol_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmProtocol_KEYS) ProtoMessage()               {}
func (*Ospfv3EdmProtocol_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ospfv3EdmProtocol_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmProtocol_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type Ospfv3EdmProtocol struct {
	// Protocol router ID
	ProtocolRouterId string `protobuf:"bytes,50,opt,name=protocol_router_id,json=protocolRouterId" json:"protocol_router_id,omitempty"`
	// Administrative distance
	AdministrativeDistance uint32 `protobuf:"varint,51,opt,name=administrative_distance,json=administrativeDistance" json:"administrative_distance,omitempty"`
	// Administrative Distance for Inter Area routes
	AdministrativeDistanceInterArea uint32 `protobuf:"varint,52,opt,name=administrative_distance_inter_area,json=administrativeDistanceInterArea" json:"administrative_distance_inter_area,omitempty"`
	// Administrative Distance for External routes
	AdministrativeDistanceExternal uint32 `protobuf:"varint,53,opt,name=administrative_distance_external,json=administrativeDistanceExternal" json:"administrative_distance_external,omitempty"`
	// If true, Graceful restart is enabled
	IsGracefulRestart bool `protobuf:"varint,54,opt,name=is_graceful_restart,json=isGracefulRestart" json:"is_graceful_restart,omitempty"`
	// Distribute List In
	DistributeListIn string `protobuf:"bytes,55,opt,name=distribute_list_in,json=distributeListIn" json:"distribute_list_in,omitempty"`
}

func (m *Ospfv3EdmProtocol) Reset()                    { *m = Ospfv3EdmProtocol{} }
func (m *Ospfv3EdmProtocol) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmProtocol) ProtoMessage()               {}
func (*Ospfv3EdmProtocol) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ospfv3EdmProtocol) GetProtocolRouterId() string {
	if m != nil {
		return m.ProtocolRouterId
	}
	return ""
}

func (m *Ospfv3EdmProtocol) GetAdministrativeDistance() uint32 {
	if m != nil {
		return m.AdministrativeDistance
	}
	return 0
}

func (m *Ospfv3EdmProtocol) GetAdministrativeDistanceInterArea() uint32 {
	if m != nil {
		return m.AdministrativeDistanceInterArea
	}
	return 0
}

func (m *Ospfv3EdmProtocol) GetAdministrativeDistanceExternal() uint32 {
	if m != nil {
		return m.AdministrativeDistanceExternal
	}
	return 0
}

func (m *Ospfv3EdmProtocol) GetIsGracefulRestart() bool {
	if m != nil {
		return m.IsGracefulRestart
	}
	return false
}

func (m *Ospfv3EdmProtocol) GetDistributeListIn() string {
	if m != nil {
		return m.DistributeListIn
	}
	return ""
}

func init() {
	proto.RegisterType((*Ospfv3EdmProtocol_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.summary.protocol.ospfv3_edm_protocol_KEYS")
	proto.RegisterType((*Ospfv3EdmProtocol)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.summary.protocol.ospfv3_edm_protocol")
}

func init() { proto.RegisterFile("ospfv3_edm_protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x91, 0x0f, 0xad, 0xbb, 0x6d, 0xa1, 0x5d, 0x43, 0xbb, 0xbe, 0xb4, 0x8a, 0x4f, 0x3e,
	0x04, 0x1d, 0xe2, 0xc4, 0x3e, 0x07, 0x62, 0x12, 0xe3, 0x10, 0x82, 0x72, 0x49, 0x4e, 0xc3, 0x5a,
	0x1a, 0x85, 0x01, 0x49, 0x2b, 0x66, 0x57, 0xc2, 0x79, 0xec, 0xbc, 0x41, 0xd0, 0x4a, 0x4e, 0x08,
	0xc8, 0x97, 0x65, 0xd8, 0xef, 0x9f, 0xef, 0x87, 0x11, 0x53, 0x63, 0xab, 0xac, 0x59, 0x00, 0xa6,
	0x05, 0x54, 0x6c, 0x9c, 0x49, 0x4c, 0x1e, 0xf9, 0x41, 0xde, 0x27, 0x64, 0x13, 0x03, 0x64, 0x2c,
	0xec, 0x19, 0xa8, 0x6a, 0x96, 0xd0, 0x87, 0x4d, 0x85, 0x1c, 0x75, 0x73, 0x9b, 0x4d, 0xd0, 0x5a,
	0xb4, 0x87, 0x29, 0x6a, 0x38, 0xf3, 0x4f, 0x64, 0xeb, 0xa2, 0xd0, 0xfc, 0x12, 0x1d, 0xbc, 0xb3,
	0x47, 0xa1, 0x06, 0xea, 0x60, 0xbb, 0x7e, 0x7a, 0x90, 0x27, 0xe2, 0x47, 0x2f, 0x80, 0x52, 0x17,
	0xa8, 0x82, 0x30, 0x98, 0x7f, 0x8b, 0xbf, 0xf7, 0x7f, 0x77, 0xba, 0x40, 0x39, 0x15, 0xe3, 0x86,
	0xb3, 0x0e, 0x8f, 0x3c, 0xfe, 0xda, 0x70, 0xd6, 0xa2, 0xd9, 0xeb, 0x48, 0x4c, 0x06, 0xd4, 0xf2,
	0x54, 0xc8, 0xf7, 0x1a, 0x36, 0xb5, 0x43, 0x06, 0x4a, 0xd5, 0x99, 0x5f, 0xfe, 0x75, 0x20, 0xb1,
	0x07, 0x9b, 0x54, 0xae, 0xc4, 0x5f, 0x9d, 0x16, 0x54, 0x92, 0x75, 0xac, 0x1d, 0x35, 0x08, 0x29,
	0x59, 0xa7, 0xcb, 0x04, 0xd5, 0x22, 0x0c, 0xe6, 0x3f, 0xe3, 0x3f, 0x9f, 0xf1, 0x55, 0x4f, 0xe5,
	0x56, 0xcc, 0x8e, 0x2c, 0x02, 0x95, 0x6d, 0xa9, 0x66, 0xd4, 0xea, 0xdc, 0x3b, 0xfe, 0x0f, 0x3b,
	0x36, 0x6d, 0xee, 0x92, 0x51, 0xcb, 0x1b, 0x11, 0x1e, 0x93, 0xe1, 0xde, 0x21, 0x97, 0x3a, 0x57,
	0x17, 0x5e, 0xf5, 0x6f, 0x58, 0xb5, 0xee, 0x53, 0x32, 0x12, 0x13, 0xb2, 0xf0, 0xcc, 0x3a, 0xc1,
	0xac, 0xce, 0x81, 0xd1, 0x3a, 0xcd, 0x4e, 0x2d, 0xc3, 0x60, 0x3e, 0x8e, 0x7f, 0x93, 0xbd, 0xee,
	0x49, 0xdc, 0x81, 0xf6, 0x5a, 0x6d, 0x15, 0xd3, 0xae, 0x76, 0x08, 0x39, 0x59, 0x07, 0x54, 0xaa,
	0x55, 0x77, 0xad, 0x0f, 0x72, 0x4b, 0xd6, 0x6d, 0xca, 0xdd, 0x17, 0x7f, 0xbf, 0xc5, 0x5b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xba, 0xae, 0x36, 0xae, 0x43, 0x02, 0x00, 0x00,
}
