// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/service.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Size struct {
	Paperclips           int32    `protobuf:"varint,1,opt,name=paperclips,proto3" json:"paperclips,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Size) Reset()         { *m = Size{} }
func (m *Size) String() string { return proto.CompactTextString(m) }
func (*Size) ProtoMessage()    {}
func (*Size) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec64d44e618a02a6, []int{0}
}

func (m *Size) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Size.Unmarshal(m, b)
}
func (m *Size) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Size.Marshal(b, m, deterministic)
}
func (m *Size) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Size.Merge(m, src)
}
func (m *Size) XXX_Size() int {
	return xxx_messageInfo_Size.Size(m)
}
func (m *Size) XXX_DiscardUnknown() {
	xxx_messageInfo_Size.DiscardUnknown(m)
}

var xxx_messageInfo_Size proto.InternalMessageInfo

func (m *Size) GetPaperclips() int32 {
	if m != nil {
		return m.Paperclips
	}
	return 0
}

type Paperclips struct {
	Paperclips           int32    `protobuf:"varint,1,opt,name=paperclips,proto3" json:"paperclips,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paperclips) Reset()         { *m = Paperclips{} }
func (m *Paperclips) String() string { return proto.CompactTextString(m) }
func (*Paperclips) ProtoMessage()    {}
func (*Paperclips) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec64d44e618a02a6, []int{1}
}

func (m *Paperclips) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paperclips.Unmarshal(m, b)
}
func (m *Paperclips) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paperclips.Marshal(b, m, deterministic)
}
func (m *Paperclips) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paperclips.Merge(m, src)
}
func (m *Paperclips) XXX_Size() int {
	return xxx_messageInfo_Paperclips.Size(m)
}
func (m *Paperclips) XXX_DiscardUnknown() {
	xxx_messageInfo_Paperclips.DiscardUnknown(m)
}

var xxx_messageInfo_Paperclips proto.InternalMessageInfo

func (m *Paperclips) GetPaperclips() int32 {
	if m != nil {
		return m.Paperclips
	}
	return 0
}

type Dread struct {
	Paperclips           int32    `protobuf:"varint,1,opt,name=paperclips,proto3" json:"paperclips,omitempty"`
	UniverseLifespan     string   `protobuf:"bytes,2,opt,name=universeLifespan,proto3" json:"universeLifespan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dread) Reset()         { *m = Dread{} }
func (m *Dread) String() string { return proto.CompactTextString(m) }
func (*Dread) ProtoMessage()    {}
func (*Dread) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec64d44e618a02a6, []int{2}
}

func (m *Dread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dread.Unmarshal(m, b)
}
func (m *Dread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dread.Marshal(b, m, deterministic)
}
func (m *Dread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dread.Merge(m, src)
}
func (m *Dread) XXX_Size() int {
	return xxx_messageInfo_Dread.Size(m)
}
func (m *Dread) XXX_DiscardUnknown() {
	xxx_messageInfo_Dread.DiscardUnknown(m)
}

var xxx_messageInfo_Dread proto.InternalMessageInfo

func (m *Dread) GetPaperclips() int32 {
	if m != nil {
		return m.Paperclips
	}
	return 0
}

func (m *Dread) GetUniverseLifespan() string {
	if m != nil {
		return m.UniverseLifespan
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec64d44e618a02a6, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Size)(nil), "twirp.paperclips.Size")
	proto.RegisterType((*Paperclips)(nil), "twirp.paperclips.Paperclips")
	proto.RegisterType((*Dread)(nil), "twirp.paperclips.Dread")
	proto.RegisterType((*Empty)(nil), "twirp.paperclips.Empty")
}

func init() { proto.RegisterFile("rpc/service.proto", fileDescriptor_ec64d44e618a02a6) }

var fileDescriptor_ec64d44e618a02a6 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2a, 0x48, 0xd6,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28,
	0x29, 0xcf, 0x2c, 0x2a, 0xd0, 0x2b, 0x48, 0x2c, 0x48, 0x2d, 0x4a, 0xce, 0xc9, 0x2c, 0x28, 0x56,
	0x52, 0xe3, 0x62, 0x09, 0xce, 0xac, 0x4a, 0x15, 0x92, 0xe3, 0xe2, 0x42, 0x88, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0xb0, 0x06, 0x21, 0x89, 0x28, 0xe9, 0x70, 0x71, 0x05, 0xc0, 0x79, 0x04, 0x55, 0x07,
	0x73, 0xb1, 0xba, 0x14, 0xa5, 0x26, 0xa6, 0x10, 0x52, 0x28, 0xa4, 0xc5, 0x25, 0x50, 0x9a, 0x97,
	0x59, 0x96, 0x5a, 0x54, 0x9c, 0xea, 0x93, 0x99, 0x96, 0x5a, 0x5c, 0x90, 0x98, 0x27, 0xc1, 0xa4,
	0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x21, 0xae, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50, 0x52, 0x69,
	0xf4, 0x85, 0x91, 0x4b, 0x38, 0x14, 0x22, 0x9b, 0x98, 0x83, 0xe4, 0x2a, 0x37, 0x2e, 0x5e, 0xf7,
	0xd4, 0x12, 0x24, 0x01, 0x71, 0x3d, 0x74, 0xff, 0xea, 0x81, 0x4d, 0x90, 0x92, 0xc1, 0x94, 0x40,
	0x31, 0x47, 0xd8, 0x33, 0x2f, 0xb9, 0x28, 0x35, 0x37, 0x35, 0x0f, 0xd9, 0x34, 0x31, 0x4c, 0x4d,
	0xa0, 0xa0, 0x93, 0xc2, 0x65, 0x8b, 0x90, 0x2f, 0x97, 0xa4, 0x73, 0x62, 0x4e, 0x72, 0x69, 0x4e,
	0x62, 0x49, 0x6a, 0x28, 0x9a, 0x6f, 0x70, 0xbb, 0x0d, 0x8b, 0x04, 0x38, 0x2c, 0x9d, 0x58, 0xa3,
	0x98, 0x8b, 0x0a, 0x92, 0x93, 0xd8, 0xc0, 0x51, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe2,
	0x8a, 0x3c, 0x9b, 0xdf, 0x01, 0x00, 0x00,
}
