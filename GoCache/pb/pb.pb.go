// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package geecachepb

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

type Request struct {
	Group                string   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Request) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Response struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "geecachepb.Request")
	proto.RegisterType((*Response)(nil), "geecachepb.Response")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x48, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x4f, 0x4d, 0x4d, 0x4e, 0x4c, 0xce, 0x48, 0x2d, 0x48,
	0x52, 0x32, 0xe4, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x4d,
	0x2f, 0xca, 0x2f, 0x2d, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x04, 0xb8,
	0x98, 0xb3, 0x53, 0x2b, 0x25, 0x98, 0xc0, 0x62, 0x20, 0xa6, 0x92, 0x02, 0x17, 0x47, 0x50, 0x6a,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x4f, 0x59, 0x62, 0x4e, 0x69, 0x2a, 0x58, 0x0f, 0x4f,
	0x10, 0x84, 0x63, 0x64, 0xc7, 0xc5, 0xe5, 0x0e, 0xd2, 0xec, 0x0c, 0xb2, 0x44, 0xc8, 0x80, 0x8b,
	0xd9, 0x3d, 0xb5, 0x44, 0x48, 0x58, 0x0f, 0x61, 0xad, 0x1e, 0xd4, 0x4e, 0x29, 0x11, 0x54, 0x41,
	0x88, 0xa9, 0x49, 0x6c, 0x60, 0x77, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x87, 0x7f, 0x5e,
	0x65, 0xb3, 0x00, 0x00, 0x00,
}
