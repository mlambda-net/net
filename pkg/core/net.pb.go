// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: net.proto

package core

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Type    string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Kind    string `protobuf:"bytes,2,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	Payload []byte `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (m *Request) Reset()      { *m = Request{} }
func (*Request) ProtoMessage() {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Request) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Request) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Response struct {
	Status  int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (m *Response) Reset()      { *m = Response{} }
func (*Response) ProtoMessage() {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Check struct {
}

func (m *Check) Reset()      { *m = Check{} }
func (*Check) ProtoMessage() {}
func (*Check) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{2}
}
func (m *Check) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Check) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Check.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Check) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Check.Merge(m, src)
}
func (m *Check) XXX_Size() int {
	return m.Size()
}
func (m *Check) XXX_DiscardUnknown() {
	xxx_messageInfo_Check.DiscardUnknown(m)
}

var xxx_messageInfo_Check proto.InternalMessageInfo

type Status struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Status) Reset()      { *m = Status{} }
func (*Status) ProtoMessage() {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{3}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Error struct {
	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Error) Reset()      { *m = Error{} }
func (*Error) ProtoMessage() {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{4}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Done struct {
}

func (m *Done) Reset()      { *m = Done{} }
func (*Done) ProtoMessage() {}
func (*Done) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{5}
}
func (m *Done) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Done) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Done.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Done) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Done.Merge(m, src)
}
func (m *Done) XXX_Size() int {
	return m.Size()
}
func (m *Done) XXX_DiscardUnknown() {
	xxx_messageInfo_Done.DiscardUnknown(m)
}

var xxx_messageInfo_Done proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "core.Request")
	proto.RegisterType((*Response)(nil), "core.Response")
	proto.RegisterType((*Check)(nil), "core.Check")
	proto.RegisterType((*Status)(nil), "core.Status")
	proto.RegisterType((*Error)(nil), "core.Error")
	proto.RegisterType((*Done)(nil), "core.Done")
}

func init() { proto.RegisterFile("net.proto", fileDescriptor_a5b10ce944527a32) }

var fileDescriptor_a5b10ce944527a32 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xb1, 0x4e, 0x32, 0x41,
	0x10, 0xc7, 0x77, 0xf9, 0xee, 0x0e, 0x98, 0x8f, 0xef, 0x2b, 0x36, 0xc6, 0x5c, 0x28, 0x36, 0xe4,
	0x8c, 0x91, 0x8a, 0x42, 0x6d, 0x4c, 0xac, 0x44, 0x13, 0x13, 0x35, 0x31, 0x27, 0xb1, 0x3f, 0x8f,
	0x89, 0x10, 0xce, 0x5d, 0xbc, 0x5d, 0x4c, 0xe8, 0x78, 0x04, 0x1f, 0xc3, 0x47, 0xb1, 0xa4, 0xa4,
	0x94, 0xa5, 0xb1, 0xe4, 0x11, 0xcc, 0xdd, 0xde, 0x19, 0x28, 0xb4, 0x9b, 0xdf, 0xcc, 0xce, 0x7f,
	0x66, 0xfe, 0x59, 0xa8, 0x0b, 0xd4, 0x9d, 0x71, 0x2a, 0xb5, 0x64, 0x4e, 0x2c, 0x53, 0x0c, 0x22,
	0xa8, 0x86, 0xf8, 0x3c, 0x41, 0xa5, 0x19, 0x03, 0xa7, 0x37, 0x1d, 0xa3, 0x4f, 0x5b, 0xb4, 0x5d,
	0x0f, 0xf3, 0x38, 0xcb, 0x5d, 0x0d, 0x45, 0xdf, 0xaf, 0xd8, 0x5c, 0x16, 0xb3, 0x1d, 0x70, 0x7b,
	0x72, 0x84, 0xc2, 0xff, 0x93, 0x27, 0x2d, 0x30, 0x1f, 0xaa, 0xb7, 0xd1, 0x34, 0x91, 0x51, 0xdf,
	0x77, 0x5a, 0xb4, 0xdd, 0x08, 0x4b, 0x0c, 0xee, 0xa1, 0x16, 0xa2, 0x1a, 0x4b, 0xa1, 0x90, 0xed,
	0x82, 0x77, 0xa7, 0x23, 0x3d, 0x51, 0xf9, 0x14, 0x37, 0x2c, 0x68, 0xb3, 0xbb, 0xb2, 0xd5, 0x9d,
	0x55, 0x6e, 0x50, 0xa9, 0xe8, 0x11, 0x8b, 0x79, 0x25, 0x06, 0x55, 0x70, 0xbb, 0x03, 0x8c, 0x47,
	0xc1, 0x29, 0x6c, 0xc8, 0xa8, 0x49, 0x1c, 0xa3, 0xb2, 0xfa, 0xb5, 0xb0, 0xc4, 0xac, 0xf2, 0x54,
	0xc8, 0xd8, 0x5b, 0x4a, 0x0c, 0x4e, 0xc0, 0xbd, 0x48, 0x53, 0x99, 0x66, 0xbb, 0xa9, 0xad, 0xdd,
	0xd4, 0xb7, 0xe8, 0x0f, 0xad, 0x1e, 0x38, 0xe7, 0x52, 0xe0, 0xe1, 0x8c, 0x42, 0xbd, 0x2b, 0x85,
	0xc0, 0x58, 0xcb, 0x94, 0x1d, 0x80, 0xd3, 0x8d, 0x92, 0x84, 0xfd, 0xeb, 0x64, 0x0e, 0x77, 0x0a,
	0x7b, 0x9b, 0xff, 0x4b, 0xb4, 0x56, 0x04, 0x84, 0xed, 0x81, 0x73, 0x3d, 0x7c, 0x41, 0xf6, 0xd7,
	0x56, 0xf2, 0x63, 0x9a, 0x0d, 0x0b, 0xf6, 0xa0, 0x80, 0xb0, 0x7d, 0xf0, 0x2e, 0x31, 0x4a, 0xf4,
	0xe0, 0xd7, 0x67, 0x67, 0xc7, 0xf3, 0x25, 0x27, 0x8b, 0x25, 0x27, 0xeb, 0x25, 0xa7, 0x33, 0xc3,
	0xe9, 0x9b, 0xe1, 0xf4, 0xdd, 0x70, 0x3a, 0x37, 0x9c, 0x7e, 0x18, 0x4e, 0x3f, 0x0d, 0x27, 0x6b,
	0xc3, 0xe9, 0xeb, 0x8a, 0x93, 0xf9, 0x8a, 0x93, 0xc5, 0x8a, 0x93, 0x07, 0x2f, 0xff, 0x0a, 0x47,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x38, 0x6f, 0x68, 0x15, 0x17, 0x02, 0x00, 0x00,
}

func (this *Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Request)
	if !ok {
		that2, ok := that.(Request)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	return true
}
func (this *Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response)
	if !ok {
		that2, ok := that.(Response)
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
	if this.Status != that1.Status {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *Check) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Check)
	if !ok {
		that2, ok := that.(Check)
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
	return true
}
func (this *Status) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Status)
	if !ok {
		that2, ok := that.(Status)
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
	if this.Success != that1.Success {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *Error) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Error)
	if !ok {
		that2, ok := that.(Error)
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
	if this.Status != that1.Status {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *Done) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Done)
	if !ok {
		that2, ok := that.(Done)
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
	return true
}
func (this *Request) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&core.Request{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Kind: "+fmt.Sprintf("%#v", this.Kind)+",\n")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Response) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&core.Response{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Check) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&core.Check{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Status) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&core.Status{")
	s = append(s, "Success: "+fmt.Sprintf("%#v", this.Success)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Error) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&core.Error{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Done) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&core.Done{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringNet(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintNet(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintNet(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = encodeVarintNet(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintNet(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintNet(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintNet(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintNet(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Check) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Check) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Check) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintNet(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintNet(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintNet(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Done) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Done) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Done) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintNet(dAtA []byte, offset int, v uint64) int {
	offset -= sovNet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovNet(uint64(l))
	}
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovNet(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovNet(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovNet(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovNet(uint64(m.Status))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovNet(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovNet(uint64(l))
	}
	return n
}

func (m *Check) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovNet(uint64(l))
	}
	return n
}

func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovNet(uint64(m.Status))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovNet(uint64(l))
	}
	return n
}

func (m *Done) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovNet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNet(x uint64) (n int) {
	return sovNet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Request{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Response) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Response{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Check) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Check{`,
		`}`,
	}, "")
	return s
}
func (this *Status) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Status{`,
		`Success:` + fmt.Sprintf("%v", this.Success) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Error) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Error{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Done) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Done{`,
		`}`,
	}, "")
	return s
}
func valueToStringNet(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
				return ErrInvalidLengthNet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
				return ErrInvalidLengthNet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
				return ErrInvalidLengthNet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
				return ErrInvalidLengthNet
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNet
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
				return ErrInvalidLengthNet
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
				return ErrInvalidLengthNet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNet
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
func (m *Check) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
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
			return fmt.Errorf("proto: Check: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Check: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNet
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
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
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
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
				return ErrInvalidLengthNet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNet
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
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
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
				return ErrInvalidLengthNet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNet
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
func (m *Done) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
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
			return fmt.Errorf("proto: Done: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Done: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNet
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
func skipNet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNet
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
					return 0, ErrIntOverflowNet
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
					return 0, ErrIntOverflowNet
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
				return 0, ErrInvalidLengthNet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNet = fmt.Errorf("proto: unexpected end of group")
)
