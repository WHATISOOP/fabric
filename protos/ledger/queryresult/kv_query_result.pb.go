// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger/queryresult/kv_query_result.proto

/*
Package queryresult is a generated protocol buffer package.

It is generated from these files:
	ledger/queryresult/kv_query_result.proto

It has these top-level messages:
	KV
	KeyModification
*/
package queryresult

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// KV -- QueryResult for range/execute query. Holds a key and corresponding value.
type KV struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KV) Reset()                    { *m = KV{} }
func (m *KV) String() string            { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()               {}
func (*KV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KV) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// KeyModification -- QueryResult for history query. Holds a transaction ID, value,
// timestamp, and delete marker which resulted from a history query.
type KeyModification struct {
	TxId      string                     `protobuf:"bytes,1,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
	Value     []byte                     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	IsDelete  bool                       `protobuf:"varint,4,opt,name=is_delete,json=isDelete" json:"is_delete,omitempty"`
}

func (m *KeyModification) Reset()                    { *m = KeyModification{} }
func (m *KeyModification) String() string            { return proto.CompactTextString(m) }
func (*KeyModification) ProtoMessage()               {}
func (*KeyModification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeyModification) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *KeyModification) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyModification) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *KeyModification) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func init() {
	proto.RegisterType((*KV)(nil), "queryresult.KV")
	proto.RegisterType((*KeyModification)(nil), "queryresult.KeyModification")
}

func init() { proto.RegisterFile("ledger/queryresult/kv_query_result.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0x24, 0xfd, 0xf3, 0xa3, 0xd9, 0xfe, 0x40, 0x59, 0x3d, 0x84, 0x2a, 0x58, 0x7a, 0xca, 0x69,
	0x57, 0x14, 0xc1, 0xb3, 0x78, 0xd1, 0xe2, 0x25, 0x88, 0x07, 0x2f, 0x61, 0x93, 0xbc, 0x24, 0x8f,
	0x26, 0xdd, 0xb8, 0xbb, 0xa9, 0xcd, 0xe7, 0xf0, 0x0b, 0x8b, 0xbb, 0xad, 0x09, 0x78, 0xdb, 0x99,
	0x37, 0x33, 0x0c, 0xb3, 0x24, 0xac, 0x20, 0x2b, 0x40, 0xf1, 0x8f, 0x16, 0x54, 0xa7, 0x40, 0xb7,
	0x95, 0xe1, 0x9b, 0x5d, 0x6c, 0x61, 0xec, 0x30, 0x6b, 0x94, 0x34, 0x92, 0xce, 0x07, 0x92, 0xc5,
	0x55, 0x21, 0x65, 0x51, 0x01, 0xb7, 0xa7, 0xa4, 0xcd, 0xb9, 0xc1, 0x1a, 0xb4, 0x11, 0x75, 0xe3,
	0xd4, 0xab, 0x67, 0x32, 0x5a, 0xbf, 0xd1, 0x4b, 0xe2, 0x6f, 0x45, 0x0d, 0xba, 0x11, 0x29, 0x04,
	0xde, 0xd2, 0x0b, 0xfd, 0xa8, 0x27, 0xe8, 0x29, 0x19, 0x6f, 0xa0, 0x0b, 0x46, 0x96, 0xff, 0x79,
	0xd2, 0x73, 0x32, 0xdd, 0x89, 0xaa, 0x85, 0x60, 0xbc, 0xf4, 0xc2, 0xff, 0x91, 0x03, 0xab, 0x2f,
	0x8f, 0x9c, 0xac, 0xa1, 0x7b, 0x91, 0x19, 0xe6, 0x98, 0x0a, 0x83, 0x72, 0x4b, 0xcf, 0xc8, 0xd4,
	0xec, 0x63, 0xcc, 0x0e, 0xa9, 0x13, 0xb3, 0x7f, 0xca, 0x7a, 0xfb, 0x68, 0x60, 0xa7, 0xf7, 0xc4,
	0xff, 0x6d, 0x67, 0x83, 0xe7, 0x37, 0x0b, 0xe6, 0xfa, 0xb3, 0x63, 0x7f, 0xf6, 0x7a, 0x54, 0x44,
	0xbd, 0x98, 0x5e, 0x10, 0x1f, 0x75, 0x9c, 0x41, 0x05, 0x06, 0x82, 0xc9, 0xd2, 0x0b, 0x67, 0xd1,
	0x0c, 0xf5, 0xa3, 0xc5, 0x0f, 0x48, 0xae, 0xa5, 0x2a, 0x58, 0xd9, 0x35, 0xa0, 0xdc, 0x88, 0x2c,
	0x17, 0x89, 0xc2, 0xd4, 0x85, 0x6a, 0x76, 0x20, 0x07, 0xb3, 0xbd, 0xdf, 0x15, 0x68, 0xca, 0x36,
	0x61, 0xa9, 0xac, 0xf9, 0x67, 0x29, 0x0c, 0x6a, 0x29, 0x1b, 0xee, 0x6c, 0x6e, 0x4b, 0xcd, 0xff,
	0x7e, 0x48, 0xf2, 0xcf, 0x9e, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x06, 0x94, 0x0c, 0x95,
	0xad, 0x01, 0x00, 0x00,
}
