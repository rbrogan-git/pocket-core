// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/nodes/nodes.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_pokt_network_pocket_core_types "github.com/pokt-network/pocket-core/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorProto struct {
	Address                 github_com_pokt_network_pocket_core_types.Address     `protobuf:"bytes,1,opt,name=Address,proto3,casttype=github.com/pokt-network/pocket-core/types.Address" json:"address" yaml:"address"`
	PublicKey               string                                                `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"public_key" yaml:"public_key"`
	Jailed                  bool                                                  `protobuf:"varint,3,opt,name=jailed,proto3" json:"jailed,omitempty"`
	Status                  github_com_pokt_network_pocket_core_types.StakeStatus `protobuf:"varint,4,opt,name=status,proto3,casttype=github.com/pokt-network/pocket-core/types.StakeStatus" json:"status,omitempty"`
	Chains                  []string                                              `protobuf:"bytes,5,rep,name=Chains,proto3" json:"chains"`
	ServiceURL              string                                                `protobuf:"bytes,6,opt,name=ServiceURL,proto3" json:"service_url"`
	StakedTokens            github_com_pokt_network_pocket_core_types.Int         `protobuf:"bytes,7,opt,name=StakedTokens,proto3,customtype=github.com/pokt-network/pocket-core/types.Int" json:"tokens"`
	UnstakingCompletionTime time.Time                                             `protobuf:"bytes,8,opt,name=UnstakingCompletionTime,proto3,stdtime" json:"unstaking_time" yaml:"unstaking_time"`
}

func (m *ValidatorProto) Reset()         { *m = ValidatorProto{} }
func (m *ValidatorProto) String() string { return proto.CompactTextString(m) }
func (*ValidatorProto) ProtoMessage()    {}
func (*ValidatorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb49073b61e33a, []int{0}
}
func (m *ValidatorProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorProto.Merge(m, src)
}
func (m *ValidatorProto) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorProto.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorProto proto.InternalMessageInfo

// ValidatorSigningInfo defines the signing info for a validator
type ValidatorSigningInfo struct {
	Address github_com_pokt_network_pocket_core_types.Address `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/pokt-network/pocket-core/types.Address" json:"address"`
	// height at which validator was first a candidate OR was unjailed
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height" yaml:"start_height"`
	// index offset into signed block bit array
	Index int64 `protobuf:"varint,3,opt,name=Index,proto3" json:"index_offset" yaml:"index_offset"`
	// timestamp validator cannot be unjailed until
	JailedUntil time.Time `protobuf:"bytes,4,opt,name=jailed_until,json=jailedUntil,proto3,stdtime" json:"jailed_until" yaml:"jailed_until"`
	// missed blocks counter (to avoid scanning the array every time)
	MissedBlocksCounter int64 `protobuf:"varint,5,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter" yaml:"missed_blocks_counter"`
	JailedBlocksCounter int64 `protobuf:"varint,6,opt,name=jailed_blocks_counter,json=jailedBlocksCounter,proto3" json:"jailed_blocks_counter" yaml:"jailed_blocks_counter"`
}

func (m *ValidatorSigningInfo) Reset()      { *m = ValidatorSigningInfo{} }
func (*ValidatorSigningInfo) ProtoMessage() {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb49073b61e33a, []int{1}
}
func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() github_com_pokt_network_pocket_core_types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ValidatorSigningInfo) GetJailedUntil() time.Time {
	if m != nil {
		return m.JailedUntil
	}
	return time.Time{}
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

func (m *ValidatorSigningInfo) GetJailedBlocksCounter() int64 {
	if m != nil {
		return m.JailedBlocksCounter
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidatorProto)(nil), "x.nodes.ValidatorProto")
	proto.RegisterType((*ValidatorSigningInfo)(nil), "x.nodes.ValidatorSigningInfo")
}

func init() { proto.RegisterFile("x/nodes/nodes.proto", fileDescriptor_63cb49073b61e33a) }

var fileDescriptor_63cb49073b61e33a = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xbf, 0x6f, 0xd3, 0x4c,
	0x18, 0xc7, 0xe3, 0xb7, 0x4d, 0xd2, 0x5e, 0xa2, 0xbe, 0x7a, 0x9d, 0xf6, 0x7d, 0xad, 0xea, 0x55,
	0x2e, 0x32, 0x03, 0x59, 0x1a, 0x0b, 0x2a, 0x86, 0x56, 0x42, 0x50, 0x77, 0xa1, 0x94, 0xa1, 0x38,
	0x2d, 0x43, 0x17, 0xcb, 0xb1, 0x2f, 0xce, 0xe1, 0x1f, 0x17, 0xf9, 0xce, 0xd0, 0x6c, 0x8c, 0xb0,
	0x75, 0xec, 0x98, 0x3f, 0xa7, 0x63, 0x17, 0x24, 0xc4, 0x70, 0xa0, 0x76, 0x41, 0x19, 0x33, 0x76,
	0x42, 0xbe, 0x73, 0x48, 0x82, 0x02, 0xaa, 0x58, 0xa2, 0x3c, 0x9f, 0xbb, 0xe7, 0xfb, 0x7d, 0x4e,
	0xcf, 0x57, 0x06, 0xb5, 0x33, 0x23, 0x26, 0x1e, 0xa2, 0xf2, 0xb7, 0xd5, 0x4f, 0x08, 0x23, 0x6a,
	0xf9, 0xac, 0x25, 0xca, 0xcd, 0x75, 0x9f, 0xf8, 0x44, 0x30, 0x23, 0xfb, 0x27, 0x8f, 0x37, 0xa1,
	0x4f, 0x88, 0x1f, 0x22, 0x43, 0x54, 0x9d, 0xb4, 0x6b, 0x30, 0x1c, 0x21, 0xca, 0x9c, 0xa8, 0x2f,
	0x2f, 0xe8, 0xef, 0x8a, 0x60, 0xed, 0x95, 0x13, 0x62, 0xcf, 0x61, 0x24, 0x39, 0x12, 0x92, 0x21,
	0x28, 0xef, 0x79, 0x5e, 0x82, 0x28, 0xd5, 0x94, 0x86, 0xd2, 0xac, 0x9a, 0xd6, 0x88, 0xc3, 0xb2,
	0x23, 0xd1, 0x98, 0xc3, 0xb5, 0x81, 0x13, 0x85, 0xbb, 0x7a, 0x0e, 0xf4, 0x5b, 0x0e, 0x1f, 0xf8,
	0x98, 0xf5, 0xd2, 0x4e, 0xcb, 0x25, 0x91, 0xd1, 0x27, 0x01, 0xdb, 0x8a, 0x11, 0x7b, 0x4b, 0x92,
	0xc0, 0xe8, 0x13, 0x37, 0x40, 0x6c, 0xcb, 0x25, 0x09, 0x32, 0xd8, 0xa0, 0x8f, 0x68, 0x2b, 0x57,
	0xb6, 0x26, 0x16, 0xea, 0x1e, 0x58, 0x3d, 0x4a, 0x3b, 0x21, 0x76, 0x0f, 0xd1, 0x40, 0xfb, 0xab,
	0xa1, 0x34, 0x57, 0xcd, 0x7b, 0x23, 0x0e, 0x41, 0x5f, 0x40, 0x3b, 0x40, 0x83, 0x31, 0x87, 0xff,
	0x48, 0xcb, 0x29, 0xd3, 0xad, 0x69, 0x97, 0xfa, 0x2f, 0x28, 0xbd, 0x76, 0x70, 0x88, 0x3c, 0x6d,
	0xa9, 0xa1, 0x34, 0x57, 0xac, 0xbc, 0x52, 0x5f, 0x82, 0x12, 0x65, 0x0e, 0x4b, 0xa9, 0xb6, 0xdc,
	0x50, 0x9a, 0x45, 0x73, 0xe7, 0x96, 0xc3, 0x47, 0x77, 0x1f, 0xb5, 0xcd, 0x9c, 0x00, 0xb5, 0x85,
	0x80, 0x95, 0x0b, 0xa9, 0x3a, 0x28, 0xed, 0xf7, 0x1c, 0x1c, 0x53, 0xad, 0xd8, 0x58, 0x6a, 0xae,
	0x9a, 0x60, 0xc4, 0x61, 0xc9, 0x15, 0xc4, 0xca, 0x4f, 0x54, 0x03, 0x80, 0x36, 0x4a, 0xde, 0x60,
	0x17, 0x9d, 0x58, 0x2f, 0xb4, 0x92, 0x78, 0xd2, 0xdf, 0x23, 0x0e, 0x2b, 0x54, 0x52, 0x3b, 0x4d,
	0x42, 0x6b, 0xe6, 0x8a, 0xea, 0x82, 0xaa, 0xf0, 0xf2, 0x8e, 0x49, 0x80, 0x62, 0xaa, 0x95, 0x45,
	0xcb, 0x93, 0x4b, 0x0e, 0x0b, 0x9f, 0x39, 0xdc, 0xba, 0xfb, 0xc4, 0x07, 0x31, 0xcb, 0xe6, 0x61,
	0x42, 0xc6, 0x9a, 0x13, 0x55, 0x3f, 0x28, 0xe0, 0xbf, 0x93, 0x98, 0x32, 0x27, 0xc0, 0xb1, 0xbf,
	0x4f, 0xa2, 0x7e, 0x88, 0x18, 0x26, 0xf1, 0x31, 0x8e, 0x90, 0xb6, 0xd2, 0x50, 0x9a, 0x95, 0x87,
	0x9b, 0x2d, 0x19, 0x96, 0xd6, 0x24, 0x2c, 0xad, 0xe3, 0x49, 0x58, 0xcc, 0xed, 0x6c, 0x98, 0x11,
	0x87, 0x6b, 0xe9, 0x44, 0xc2, 0xce, 0x92, 0x34, 0xe6, 0x70, 0x43, 0xae, 0x66, 0x9e, 0xeb, 0xe7,
	0x5f, 0xa0, 0x62, 0xfd, 0xca, 0x6f, 0xb7, 0xfa, 0x7e, 0x08, 0x0b, 0x17, 0x43, 0xa8, 0x7c, 0x1b,
	0x42, 0x45, 0xff, 0xb8, 0x0c, 0xd6, 0x7f, 0x44, 0xb0, 0x8d, 0xfd, 0x18, 0xc7, 0xfe, 0x41, 0xdc,
	0x25, 0xea, 0x29, 0x98, 0xa4, 0x2e, 0x0f, 0xe2, 0xd3, 0x99, 0x20, 0xfe, 0x61, 0xec, 0xf2, 0x6e,
	0xf5, 0x39, 0xa8, 0x52, 0xe6, 0x24, 0xcc, 0xee, 0x21, 0xec, 0xf7, 0x98, 0x48, 0xde, 0x92, 0x79,
	0x7f, 0xc4, 0xe1, 0x1c, 0x1f, 0x73, 0x58, 0x93, 0x0f, 0x9c, 0xa5, 0xba, 0x55, 0x11, 0xe5, 0x33,
	0x51, 0xa9, 0x8f, 0x41, 0xf1, 0x20, 0xf6, 0xd0, 0x99, 0x88, 0x5f, 0x2e, 0x82, 0x33, 0x60, 0x93,
	0x6e, 0x97, 0xa2, 0x19, 0x91, 0x59, 0xaa, 0x5b, 0xb2, 0x4b, 0x8d, 0x41, 0x55, 0x06, 0xd6, 0x4e,
	0x63, 0x86, 0x43, 0x11, 0xd6, 0xdf, 0x6f, 0xc3, 0xc8, 0xb7, 0x31, 0xd7, 0x37, 0x75, 0x99, 0xa5,
	0x72, 0x13, 0x15, 0x89, 0x4e, 0x32, 0xa2, 0x46, 0x60, 0x23, 0xc2, 0x94, 0x22, 0xcf, 0xee, 0x84,
	0xc4, 0x0d, 0xa8, 0xed, 0x92, 0x34, 0x66, 0x28, 0xd1, 0x8a, 0x62, 0xfc, 0x9d, 0x11, 0x87, 0x8b,
	0x2f, 0x8c, 0x39, 0xfc, 0x5f, 0x3a, 0x2c, 0x3c, 0xd6, 0xad, 0x9a, 0xe4, 0xa6, 0xc0, 0xfb, 0x92,
	0x66, 0x76, 0xf9, 0x40, 0x3f, 0xd9, 0x95, 0xa6, 0x76, 0x0b, 0x2f, 0x4c, 0xed, 0x16, 0x1e, 0xeb,
	0x56, 0x4d, 0xf2, 0x39, 0xbb, 0xdd, 0x95, 0x8b, 0x21, 0x2c, 0x64, 0xb9, 0x32, 0x0f, 0x2f, 0xaf,
	0xeb, 0xca, 0xd5, 0x75, 0x5d, 0xf9, 0x7a, 0x5d, 0x57, 0xce, 0x6f, 0xea, 0x85, 0xab, 0x9b, 0x7a,
	0xe1, 0xd3, 0x4d, 0xbd, 0x70, 0x7a, 0xa7, 0xe0, 0x4c, 0xbe, 0xb6, 0x22, 0x40, 0x9d, 0x92, 0x58,
	0xc3, 0xf6, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xc0, 0x60, 0x38, 0x85, 0x05, 0x00, 0x00,
}

func (this *ValidatorProto) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorProto)
	if !ok {
		that2, ok := that.(ValidatorProto)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if this.PublicKey != that1.PublicKey {
		return false
	}
	if this.Jailed != that1.Jailed {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if len(this.Chains) != len(that1.Chains) {
		return false
	}
	for i := range this.Chains {
		if this.Chains[i] != that1.Chains[i] {
			return false
		}
	}
	if this.ServiceURL != that1.ServiceURL {
		return false
	}
	if !this.StakedTokens.Equal(that1.StakedTokens) {
		return false
	}
	if !this.UnstakingCompletionTime.Equal(that1.UnstakingCompletionTime) {
		return false
	}
	return true
}
func (this *ValidatorSigningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfo)
	if !ok {
		that2, ok := that.(ValidatorSigningInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if !this.JailedUntil.Equal(that1.JailedUntil) {
		return false
	}
	if this.MissedBlocksCounter != that1.MissedBlocksCounter {
		return false
	}
	if this.JailedBlocksCounter != that1.JailedBlocksCounter {
		return false
	}
	return true
}
func (m *ValidatorProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UnstakingCompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.UnstakingCompletionTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintNodes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	{
		size := m.StakedTokens.Size()
		i -= size
		if _, err := m.StakedTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNodes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.ServiceURL) > 0 {
		i -= len(m.ServiceURL)
		copy(dAtA[i:], m.ServiceURL)
		i = encodeVarintNodes(dAtA, i, uint64(len(m.ServiceURL)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintNodes(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Status != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintNodes(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNodes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.JailedBlocksCounter != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.JailedBlocksCounter))
		i--
		dAtA[i] = 0x30
	}
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x28
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.JailedUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintNodes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if m.Index != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if m.StartHeight != 0 {
		i = encodeVarintNodes(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNodes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNodes(dAtA []byte, offset int, v uint64) int {
	offset -= sovNodes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNodes(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovNodes(uint64(l))
	}
	if m.Jailed {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovNodes(uint64(m.Status))
	}
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovNodes(uint64(l))
		}
	}
	l = len(m.ServiceURL)
	if l > 0 {
		n += 1 + l + sovNodes(uint64(l))
	}
	l = m.StakedTokens.Size()
	n += 1 + l + sovNodes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UnstakingCompletionTime)
	n += 1 + l + sovNodes(uint64(l))
	return n
}

func (m *ValidatorSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNodes(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovNodes(uint64(m.StartHeight))
	}
	if m.Index != 0 {
		n += 1 + sovNodes(uint64(m.Index))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil)
	n += 1 + l + sovNodes(uint64(l))
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovNodes(uint64(m.MissedBlocksCounter))
	}
	if m.JailedBlocksCounter != 0 {
		n += 1 + sovNodes(uint64(m.JailedBlocksCounter))
	}
	return n
}

func sovNodes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNodes(x uint64) (n int) {
	return sovNodes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidatorProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= github_com_pokt_network_pocket_core_types.StakeStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakedTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakingCompletionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UnstakingCompletionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValidatorSigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNodes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.JailedUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedBlocksCounter", wireType)
			}
			m.JailedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JailedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNodes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNodes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNodes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNodes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNodes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNodes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNodes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNodes = fmt.Errorf("proto: unexpected end of group")
)
