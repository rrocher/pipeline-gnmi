// Code generated by protoc-gen-go.
// source: snmp_pdu_stats.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number is a generated protocol buffer package.

It is generated from these files:
	snmp_pdu_stats.proto

It has these top-level messages:
	SnmpPduStats_KEYS
	SnmpPduStats
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number

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

type SnmpPduStats_KEYS struct {
	Number string `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	ReqId  uint32 `protobuf:"varint,2,opt,name=req_id,json=reqId" json:"req_id,omitempty"`
	Port   uint32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
}

func (m *SnmpPduStats_KEYS) Reset()                    { *m = SnmpPduStats_KEYS{} }
func (m *SnmpPduStats_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpPduStats_KEYS) ProtoMessage()               {}
func (*SnmpPduStats_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SnmpPduStats_KEYS) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *SnmpPduStats_KEYS) GetReqId() uint32 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *SnmpPduStats_KEYS) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type SnmpPduStats struct {
	//  NMS address Rx PDU
	Nms string `protobuf:"bytes,50,opt,name=nms" json:"nms,omitempty"`
	//  SNMP request id per PDU
	RequestId uint32 `protobuf:"varint,51,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	// NMS port number
	Port uint32 `protobuf:"varint,52,opt,name=port" json:"port,omitempty"`
	//  PDU type
	PduType uint32 `protobuf:"varint,53,opt,name=pdu_type,json=pduType" json:"pdu_type,omitempty"`
	// Is reques dropped due to error
	ErrorStatus uint32 `protobuf:"varint,54,opt,name=error_status,json=errorStatus" json:"error_status,omitempty"`
	// Serial number per PDU processing
	SerialNum uint32 `protobuf:"varint,55,opt,name=serial_num,json=serialNum" json:"serial_num,omitempty"`
	// Request inserted into input queue
	InputQ string `protobuf:"bytes,56,opt,name=input_q,json=inputQ" json:"input_q,omitempty"`
	// De-queue the request from the input queue
	OutputQ uint32 `protobuf:"varint,57,opt,name=output_q,json=outputQ" json:"output_q,omitempty"`
	// Enqueue the request into pending queue
	PendingQ uint32 `protobuf:"varint,58,opt,name=pending_q,json=pendingQ" json:"pending_q,omitempty"`
	// Response sent
	ResponseOut uint32 `protobuf:"varint,59,opt,name=response_out,json=responseOut" json:"response_out,omitempty"`
}

func (m *SnmpPduStats) Reset()                    { *m = SnmpPduStats{} }
func (m *SnmpPduStats) String() string            { return proto.CompactTextString(m) }
func (*SnmpPduStats) ProtoMessage()               {}
func (*SnmpPduStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpPduStats) GetNms() string {
	if m != nil {
		return m.Nms
	}
	return ""
}

func (m *SnmpPduStats) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *SnmpPduStats) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *SnmpPduStats) GetPduType() uint32 {
	if m != nil {
		return m.PduType
	}
	return 0
}

func (m *SnmpPduStats) GetErrorStatus() uint32 {
	if m != nil {
		return m.ErrorStatus
	}
	return 0
}

func (m *SnmpPduStats) GetSerialNum() uint32 {
	if m != nil {
		return m.SerialNum
	}
	return 0
}

func (m *SnmpPduStats) GetInputQ() string {
	if m != nil {
		return m.InputQ
	}
	return ""
}

func (m *SnmpPduStats) GetOutputQ() uint32 {
	if m != nil {
		return m.OutputQ
	}
	return 0
}

func (m *SnmpPduStats) GetPendingQ() uint32 {
	if m != nil {
		return m.PendingQ
	}
	return 0
}

func (m *SnmpPduStats) GetResponseOut() uint32 {
	if m != nil {
		return m.ResponseOut
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpPduStats_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.serial_numbers.serial_number.snmp_pdu_stats_KEYS")
	proto.RegisterType((*SnmpPduStats)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.serial_numbers.serial_number.snmp_pdu_stats")
}

func init() { proto.RegisterFile("snmp_pdu_stats.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbb, 0x4e, 0x33, 0x31,
	0x10, 0x85, 0x95, 0xe4, 0xff, 0x73, 0x31, 0x17, 0x21, 0x73, 0x33, 0x42, 0x48, 0x21, 0x55, 0xaa,
	0x2d, 0x08, 0x77, 0x6a, 0x0a, 0x40, 0x02, 0x25, 0xa1, 0x80, 0xca, 0xda, 0x64, 0x87, 0xc8, 0x12,
	0x6b, 0x7b, 0xc7, 0xb6, 0x44, 0x5e, 0x88, 0xe7, 0x44, 0x1e, 0x07, 0x85, 0xed, 0xe6, 0x7c, 0x33,
	0x3e, 0x67, 0x6c, 0xb3, 0x3d, 0xa7, 0x4b, 0x2b, 0x6d, 0x11, 0xa4, 0xf3, 0xb9, 0x77, 0x99, 0x45,
	0xe3, 0x0d, 0x7f, 0x9c, 0x2b, 0x37, 0x37, 0x52, 0x19, 0x27, 0xbf, 0x50, 0xd2, 0x48, 0xbe, 0x00,
	0xed, 0xa5, 0xb1, 0x80, 0x59, 0xd4, 0x99, 0xd2, 0x1f, 0x06, 0xcb, 0xdc, 0x2b, 0xa3, 0x33, 0x07,
	0xa8, 0xf2, 0x4f, 0xa9, 0x43, 0x39, 0x03, 0x74, 0x75, 0x39, 0x78, 0x63, 0xbb, 0xf5, 0x0c, 0xf9,
	0x74, 0xff, 0x3e, 0xe5, 0x07, 0xac, 0x9d, 0x06, 0x44, 0xa3, 0xdf, 0x18, 0xf6, 0x26, 0x2b, 0xc5,
	0xf7, 0x59, 0x1b, 0xa1, 0x92, 0xaa, 0x10, 0xcd, 0x7e, 0x63, 0xb8, 0x35, 0xf9, 0x8f, 0x50, 0x3d,
	0x14, 0x9c, 0xb3, 0x7f, 0xd6, 0xa0, 0x17, 0x2d, 0x82, 0x54, 0x0f, 0xbe, 0x9b, 0x6c, 0xbb, 0x6e,
	0xcd, 0x77, 0x58, 0x4b, 0x97, 0x4e, 0x9c, 0x91, 0x65, 0x2c, 0xf9, 0x09, 0x63, 0x08, 0x55, 0x00,
	0xe7, 0xa3, 0xe7, 0x88, 0x8e, 0xf7, 0x56, 0xe4, 0x8f, 0xef, 0xf9, 0xda, 0x97, 0x1f, 0xb1, 0x6e,
	0x74, 0xf4, 0x4b, 0x0b, 0xe2, 0x82, 0x78, 0xc7, 0x16, 0xe1, 0x75, 0x69, 0x81, 0x9f, 0xb2, 0x4d,
	0x40, 0x34, 0x48, 0x71, 0xc1, 0x89, 0x4b, 0x6a, 0x6f, 0x10, 0x9b, 0x12, 0x8a, 0x81, 0xeb, 0x07,
	0x10, 0x57, 0x29, 0x30, 0x91, 0xe7, 0x50, 0xf2, 0x43, 0xd6, 0x51, 0xda, 0x06, 0x2f, 0x2b, 0x71,
	0x9d, 0x2e, 0x4e, 0x72, 0x1c, 0x53, 0x4d, 0xf0, 0xa9, 0x73, 0x93, 0x52, 0x93, 0x1e, 0xf3, 0x63,
	0xd6, 0xb3, 0xa0, 0x0b, 0xa5, 0x17, 0xb2, 0x12, 0xb7, 0xd4, 0xeb, 0xae, 0xc0, 0x38, 0xae, 0x84,
	0xe0, 0xac, 0xd1, 0x0e, 0xa4, 0x09, 0x5e, 0xdc, 0xa5, 0x95, 0x7e, 0xd9, 0x4b, 0xf0, 0xb3, 0x36,
	0xfd, 0xea, 0xe8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x97, 0x3f, 0x5e, 0x67, 0xed, 0x01, 0x00, 0x00,
}