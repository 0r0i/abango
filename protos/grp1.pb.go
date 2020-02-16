// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grp1.proto

package grp1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StdAsk struct {
	AskMsg               []byte   `protobuf:"bytes,1,opt,name=AskMsg,proto3" json:"AskMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StdAsk) Reset()         { *m = StdAsk{} }
func (m *StdAsk) String() string { return proto.CompactTextString(m) }
func (*StdAsk) ProtoMessage()    {}
func (*StdAsk) Descriptor() ([]byte, []int) {
	return fileDescriptor_7082fe03f9b760fe, []int{0}
}

func (m *StdAsk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StdAsk.Unmarshal(m, b)
}
func (m *StdAsk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StdAsk.Marshal(b, m, deterministic)
}
func (m *StdAsk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StdAsk.Merge(m, src)
}
func (m *StdAsk) XXX_Size() int {
	return xxx_messageInfo_StdAsk.Size(m)
}
func (m *StdAsk) XXX_DiscardUnknown() {
	xxx_messageInfo_StdAsk.DiscardUnknown(m)
}

var xxx_messageInfo_StdAsk proto.InternalMessageInfo

func (m *StdAsk) GetAskMsg() []byte {
	if m != nil {
		return m.AskMsg
	}
	return nil
}

type StdRet struct {
	RetSta               []byte   `protobuf:"bytes,1,opt,name=RetSta,proto3" json:"RetSta,omitempty"`
	RetMsg               []byte   `protobuf:"bytes,2,opt,name=RetMsg,proto3" json:"RetMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StdRet) Reset()         { *m = StdRet{} }
func (m *StdRet) String() string { return proto.CompactTextString(m) }
func (*StdRet) ProtoMessage()    {}
func (*StdRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_7082fe03f9b760fe, []int{1}
}

func (m *StdRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StdRet.Unmarshal(m, b)
}
func (m *StdRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StdRet.Marshal(b, m, deterministic)
}
func (m *StdRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StdRet.Merge(m, src)
}
func (m *StdRet) XXX_Size() int {
	return xxx_messageInfo_StdRet.Size(m)
}
func (m *StdRet) XXX_DiscardUnknown() {
	xxx_messageInfo_StdRet.DiscardUnknown(m)
}

var xxx_messageInfo_StdRet proto.InternalMessageInfo

func (m *StdRet) GetRetSta() []byte {
	if m != nil {
		return m.RetSta
	}
	return nil
}

func (m *StdRet) GetRetMsg() []byte {
	if m != nil {
		return m.RetMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*StdAsk)(nil), "grp1.StdAsk")
	proto.RegisterType((*StdRet)(nil), "grp1.StdRet")
}

func init() { proto.RegisterFile("grp1.proto", fileDescriptor_7082fe03f9b760fe) }

var fileDescriptor_7082fe03f9b760fe = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x30,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x14, 0xb8, 0xd8, 0x82, 0x4b,
	0x52, 0x1c, 0x8b, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x1c, 0x8b, 0xb3, 0x7d, 0x8b, 0xd3, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x78, 0x82, 0xa0, 0x3c, 0x25, 0x0b, 0xb0, 0x8a, 0xa0, 0xd4, 0x12, 0x90, 0x8a,
	0xa0, 0xd4, 0x92, 0xe0, 0x92, 0x44, 0x98, 0x0a, 0x08, 0x0f, 0x2a, 0x0e, 0xd2, 0xc9, 0x04, 0x17,
	0xf7, 0x2d, 0x4e, 0x37, 0xd2, 0xe3, 0x62, 0x71, 0x2f, 0x2a, 0x30, 0x14, 0x52, 0x83, 0x98, 0x50,
	0x90, 0x2c, 0xc4, 0xa3, 0x07, 0x76, 0x00, 0xc4, 0x46, 0x29, 0x04, 0x2f, 0x28, 0xb5, 0x44, 0x89,
	0x21, 0x89, 0x0d, 0xec, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xbe, 0x5c, 0xb6,
	0xa6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Grp1Client is the client API for Grp1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Grp1Client interface {
	StdRpc(ctx context.Context, in *StdAsk, opts ...grpc.CallOption) (*StdRet, error)
}

type grp1Client struct {
	cc grpc.ClientConnInterface
}

func NewGrp1Client(cc grpc.ClientConnInterface) Grp1Client {
	return &grp1Client{cc}
}

func (c *grp1Client) StdRpc(ctx context.Context, in *StdAsk, opts ...grpc.CallOption) (*StdRet, error) {
	out := new(StdRet)
	err := c.cc.Invoke(ctx, "/grp1.Grp1/StdRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Grp1Server is the server API for Grp1 service.
type Grp1Server interface {
	StdRpc(context.Context, *StdAsk) (*StdRet, error)
}

// UnimplementedGrp1Server can be embedded to have forward compatible implementations.
type UnimplementedGrp1Server struct {
}

func (*UnimplementedGrp1Server) StdRpc(ctx context.Context, req *StdAsk) (*StdRet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StdRpc not implemented")
}

func RegisterGrp1Server(s *grpc.Server, srv Grp1Server) {
	s.RegisterService(&_Grp1_serviceDesc, srv)
}

func _Grp1_StdRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StdAsk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Grp1Server).StdRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grp1.Grp1/StdRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Grp1Server).StdRpc(ctx, req.(*StdAsk))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grp1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grp1.Grp1",
	HandlerType: (*Grp1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StdRpc",
			Handler:    _Grp1_StdRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grp1.proto",
}
