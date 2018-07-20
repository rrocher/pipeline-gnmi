// Code generated by protoc-gen-go.
// source: ipv6_pfx_edm_ace.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_acl_oper_ipv6_acl_and_prefix_list_access_list_manager_prefixes_prefix_prefix_list_sequences_prefix_list_sequence is a generated protocol buffer package.

It is generated from these files:
	ipv6_pfx_edm_ace.proto

It has these top-level messages:
	Ipv6PfxEdmAce_KEYS
	Ipv6PfxEdmAce
*/
package cisco_ios_xr_ipv6_acl_oper_ipv6_acl_and_prefix_list_access_list_manager_prefixes_prefix_prefix_list_sequences_prefix_list_sequence

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

// Prefix list entry BAG
type Ipv6PfxEdmAce_KEYS struct {
	PrefixListName string `protobuf:"bytes,1,opt,name=prefix_list_name,json=prefixListName" json:"prefix_list_name,omitempty"`
	SequenceNumber uint32 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
}

func (m *Ipv6PfxEdmAce_KEYS) Reset()                    { *m = Ipv6PfxEdmAce_KEYS{} }
func (m *Ipv6PfxEdmAce_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6PfxEdmAce_KEYS) ProtoMessage()               {}
func (*Ipv6PfxEdmAce_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6PfxEdmAce_KEYS) GetPrefixListName() string {
	if m != nil {
		return m.PrefixListName
	}
	return ""
}

func (m *Ipv6PfxEdmAce_KEYS) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

type Ipv6PfxEdmAce struct {
	// ACE type (acl, remark)
	IsAceType string `protobuf:"bytes,50,opt,name=is_ace_type,json=isAceType" json:"is_ace_type,omitempty"`
	// ACLE sequence number
	IsAceSequenceNumber uint32 `protobuf:"varint,51,opt,name=is_ace_sequence_number,json=isAceSequenceNumber" json:"is_ace_sequence_number,omitempty"`
	// Grant value permit/deny
	IsPacketAllowOrDeny string `protobuf:"bytes,52,opt,name=is_packet_allow_or_deny,json=isPacketAllowOrDeny" json:"is_packet_allow_or_deny,omitempty"`
	// IPv6 prefix
	IsAddressInNumbers string `protobuf:"bytes,53,opt,name=is_address_in_numbers,json=isAddressInNumbers" json:"is_address_in_numbers,omitempty"`
	// Prefix length
	IsAddressMaskLength uint32 `protobuf:"varint,54,opt,name=is_address_mask_length,json=isAddressMaskLength" json:"is_address_mask_length,omitempty"`
	// Port Operator
	IsLengthOperator string `protobuf:"bytes,55,opt,name=is_length_operator,json=isLengthOperator" json:"is_length_operator,omitempty"`
	// Min length
	IsPacketMinimumLength uint32 `protobuf:"varint,56,opt,name=is_packet_minimum_length,json=isPacketMinimumLength" json:"is_packet_minimum_length,omitempty"`
	// Maximum length
	IsPacketMaximumLength uint32 `protobuf:"varint,57,opt,name=is_packet_maximum_length,json=isPacketMaximumLength" json:"is_packet_maximum_length,omitempty"`
	// Number of hits
	Hits uint32 `protobuf:"varint,58,opt,name=hits" json:"hits,omitempty"`
	// Remark String
	IsCommentForEntry string `protobuf:"bytes,59,opt,name=is_comment_for_entry,json=isCommentForEntry" json:"is_comment_for_entry,omitempty"`
	// ACL Name
	AclName string `protobuf:"bytes,60,opt,name=acl_name,json=aclName" json:"acl_name,omitempty"`
}

func (m *Ipv6PfxEdmAce) Reset()                    { *m = Ipv6PfxEdmAce{} }
func (m *Ipv6PfxEdmAce) String() string            { return proto.CompactTextString(m) }
func (*Ipv6PfxEdmAce) ProtoMessage()               {}
func (*Ipv6PfxEdmAce) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6PfxEdmAce) GetIsAceType() string {
	if m != nil {
		return m.IsAceType
	}
	return ""
}

func (m *Ipv6PfxEdmAce) GetIsAceSequenceNumber() uint32 {
	if m != nil {
		return m.IsAceSequenceNumber
	}
	return 0
}

func (m *Ipv6PfxEdmAce) GetIsPacketAllowOrDeny() string {
	if m != nil {
		return m.IsPacketAllowOrDeny
	}
	return ""
}

func (m *Ipv6PfxEdmAce) GetIsAddressInNumbers() string {
	if m != nil {
		return m.IsAddressInNumbers
	}
	return ""
}

func (m *Ipv6PfxEdmAce) GetIsAddressMaskLength() uint32 {
	if m != nil {
		return m.IsAddressMaskLength
	}
	return 0
}

func (m *Ipv6PfxEdmAce) GetIsLengthOperator() string {
	if m != nil {
		return m.IsLengthOperator
	}
	return ""
}

func (m *Ipv6PfxEdmAce) GetIsPacketMinimumLength() uint32 {
	if m != nil {
		return m.IsPacketMinimumLength
	}
	return 0
}

func (m *Ipv6PfxEdmAce) GetIsPacketMaximumLength() uint32 {
	if m != nil {
		return m.IsPacketMaximumLength
	}
	return 0
}

func (m *Ipv6PfxEdmAce) GetHits() uint32 {
	if m != nil {
		return m.Hits
	}
	return 0
}

func (m *Ipv6PfxEdmAce) GetIsCommentForEntry() string {
	if m != nil {
		return m.IsCommentForEntry
	}
	return ""
}

func (m *Ipv6PfxEdmAce) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv6PfxEdmAce_KEYS)(nil), "cisco_ios_xr_ipv6_acl_oper.ipv6_acl_and_prefix_list.access_list_manager.prefixes.prefix.prefix_list_sequences.prefix_list_sequence.ipv6_pfx_edm_ace_KEYS")
	proto.RegisterType((*Ipv6PfxEdmAce)(nil), "cisco_ios_xr_ipv6_acl_oper.ipv6_acl_and_prefix_list.access_list_manager.prefixes.prefix.prefix_list_sequences.prefix_list_sequence.ipv6_pfx_edm_ace")
}

func init() { proto.RegisterFile("ipv6_pfx_edm_ace.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcb, 0x6f, 0xd3, 0x40,
	0x10, 0x87, 0x15, 0x54, 0x01, 0x5d, 0x44, 0x09, 0x5b, 0x52, 0x96, 0x0b, 0xaa, 0x7a, 0x21, 0x07,
	0x64, 0x04, 0x29, 0x2d, 0xaf, 0x4b, 0x04, 0x45, 0x42, 0xf4, 0x81, 0x5a, 0x2e, 0x9c, 0x46, 0xdb,
	0xf5, 0xa4, 0x5d, 0xe2, 0xdd, 0x35, 0x3b, 0x1b, 0x88, 0xaf, 0x48, 0xfc, 0xdf, 0xc8, 0x63, 0x3b,
	0x4d, 0x72, 0x5b, 0xcf, 0x37, 0x3f, 0x7f, 0xb3, 0xf6, 0x88, 0x1d, 0x5b, 0xfe, 0x3e, 0x80, 0x72,
	0x32, 0x07, 0xcc, 0x1d, 0x68, 0x83, 0x59, 0x19, 0x43, 0x0a, 0xf2, 0x6f, 0xcf, 0x58, 0x32, 0x01,
	0x6c, 0x20, 0x98, 0x47, 0xe0, 0x2e, 0x6d, 0x0a, 0x08, 0x25, 0xc6, 0x6c, 0xf1, 0xa4, 0x7d, 0x0e,
	0x65, 0xc4, 0x89, 0x9d, 0x43, 0x61, 0x29, 0x65, 0xda, 0x18, 0x24, 0xe2, 0x33, 0x38, 0xed, 0xf5,
	0x15, 0xc6, 0xac, 0xe1, 0x48, 0xed, 0x21, 0x5b, 0xea, 0x07, 0xc2, 0x5f, 0x33, 0xf4, 0x66, 0x01,
	0x57, 0xab, 0x7b, 0x3f, 0xc5, 0x60, 0x7d, 0x3c, 0xf8, 0x7a, 0xf4, 0xe3, 0x42, 0x0e, 0x45, 0x7f,
	0x39, 0xe0, 0xb5, 0x43, 0xd5, 0xdb, 0xed, 0x0d, 0x37, 0xcf, 0xb7, 0x9a, 0xfa, 0xb1, 0xa5, 0x74,
	0xaa, 0x1d, 0xca, 0x67, 0xe2, 0x41, 0xf7, 0x3a, 0xf0, 0x33, 0x77, 0x89, 0x51, 0xdd, 0xda, 0xed,
	0x0d, 0xef, 0x9f, 0x6f, 0x75, 0xe5, 0x53, 0xae, 0xee, 0xfd, 0xdb, 0x10, 0xfd, 0x75, 0x99, 0x7c,
	0x2a, 0xee, 0x59, 0x62, 0x6d, 0xaa, 0x4a, 0x54, 0xaf, 0x58, 0xb1, 0x69, 0x69, 0x6c, 0xf0, 0x7b,
	0x55, 0xa2, 0x1c, 0x89, 0x9d, 0x96, 0xaf, 0x4b, 0x46, 0x2c, 0xd9, 0xe6, 0xd6, 0x8b, 0x15, 0x93,
	0xdc, 0x17, 0x8f, 0x2d, 0x41, 0xa9, 0xcd, 0x14, 0x13, 0xe8, 0xa2, 0x08, 0x7f, 0x20, 0x44, 0xc8,
	0xd1, 0x57, 0x6a, 0x9f, 0x05, 0xdb, 0x96, 0xbe, 0x31, 0x1d, 0xd7, 0xf0, 0x2c, 0x7e, 0x42, 0x5f,
	0xc9, 0x97, 0x62, 0x50, 0xab, 0xf2, 0x3c, 0xd6, 0x5f, 0xd8, 0xfa, 0x56, 0x44, 0xea, 0x35, 0x67,
	0xa4, 0xa5, 0x71, 0xc3, 0xbe, 0xf8, 0xc6, 0x43, 0xdd, 0x74, 0x6d, 0xc4, 0x69, 0x9a, 0x42, 0x81,
	0xfe, 0x2a, 0x5d, 0xab, 0x83, 0xc5, 0x74, 0x0d, 0x3c, 0xd1, 0x34, 0x3d, 0x66, 0x24, 0x9f, 0x0b,
	0x69, 0xa9, 0xed, 0xe3, 0xbf, 0xad, 0x53, 0x88, 0xea, 0x90, 0x25, 0x7d, 0x4b, 0x4d, 0xd7, 0x59,
	0x5b, 0x97, 0x87, 0x42, 0xdd, 0xdc, 0xc5, 0x59, 0x6f, 0xdd, 0xcc, 0x75, 0x92, 0x37, 0x2c, 0x19,
	0x74, 0x97, 0x39, 0x69, 0x68, 0xab, 0x59, 0x0d, 0xea, 0xf9, 0x72, 0xf0, 0xed, 0x5a, 0xb0, 0xa1,
	0x6d, 0x50, 0x8a, 0x8d, 0x6b, 0x9b, 0x48, 0xbd, 0xe3, 0x26, 0x3e, 0xcb, 0x17, 0xe2, 0x91, 0x25,
	0x30, 0xc1, 0x39, 0xf4, 0x09, 0x26, 0x21, 0x02, 0xfa, 0x14, 0x2b, 0xf5, 0x9e, 0xa7, 0x7e, 0x68,
	0xe9, 0x63, 0x83, 0x3e, 0x87, 0x78, 0x54, 0x03, 0xf9, 0x44, 0xdc, 0xad, 0xd7, 0x97, 0xf7, 0xe6,
	0x03, 0x37, 0xdd, 0xd1, 0xa6, 0xa8, 0x17, 0xe6, 0xf2, 0x36, 0xef, 0xff, 0xe8, 0x7f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x73, 0xb4, 0x22, 0x1d, 0x19, 0x03, 0x00, 0x00,
}
