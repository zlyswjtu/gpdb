// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command.proto

/*
Package idl is a generated protocol buffer package.

It is generated from these files:
	command.proto

It has these top-level messages:
	CheckUpgradeStatusRequest
	CheckUpgradeStatusReply
	FileSysUsage
	CheckDiskUsageRequest
	CheckDiskUsageReply
*/
package idl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckUpgradeStatusRequest struct {
}

func (m *CheckUpgradeStatusRequest) Reset()                    { *m = CheckUpgradeStatusRequest{} }
func (m *CheckUpgradeStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckUpgradeStatusRequest) ProtoMessage()               {}
func (*CheckUpgradeStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CheckUpgradeStatusReply struct {
	ProcessList string `protobuf:"bytes,1,opt,name=process_list,json=processList" json:"process_list,omitempty"`
}

func (m *CheckUpgradeStatusReply) Reset()                    { *m = CheckUpgradeStatusReply{} }
func (m *CheckUpgradeStatusReply) String() string            { return proto.CompactTextString(m) }
func (*CheckUpgradeStatusReply) ProtoMessage()               {}
func (*CheckUpgradeStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckUpgradeStatusReply) GetProcessList() string {
	if m != nil {
		return m.ProcessList
	}
	return ""
}

type FileSysUsage struct {
	Filesystem string  `protobuf:"bytes,1,opt,name=filesystem" json:"filesystem,omitempty"`
	Usage      float64 `protobuf:"fixed64,2,opt,name=usage" json:"usage,omitempty"`
}

func (m *FileSysUsage) Reset()                    { *m = FileSysUsage{} }
func (m *FileSysUsage) String() string            { return proto.CompactTextString(m) }
func (*FileSysUsage) ProtoMessage()               {}
func (*FileSysUsage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FileSysUsage) GetFilesystem() string {
	if m != nil {
		return m.Filesystem
	}
	return ""
}

func (m *FileSysUsage) GetUsage() float64 {
	if m != nil {
		return m.Usage
	}
	return 0
}

type CheckDiskUsageRequest struct {
}

func (m *CheckDiskUsageRequest) Reset()                    { *m = CheckDiskUsageRequest{} }
func (m *CheckDiskUsageRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckDiskUsageRequest) ProtoMessage()               {}
func (*CheckDiskUsageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CheckDiskUsageReply struct {
	ListOfFileSysUsage []*FileSysUsage `protobuf:"bytes,1,rep,name=list_of_file_sys_usage,json=listOfFileSysUsage" json:"list_of_file_sys_usage,omitempty"`
}

func (m *CheckDiskUsageReply) Reset()                    { *m = CheckDiskUsageReply{} }
func (m *CheckDiskUsageReply) String() string            { return proto.CompactTextString(m) }
func (*CheckDiskUsageReply) ProtoMessage()               {}
func (*CheckDiskUsageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CheckDiskUsageReply) GetListOfFileSysUsage() []*FileSysUsage {
	if m != nil {
		return m.ListOfFileSysUsage
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckUpgradeStatusRequest)(nil), "idl.CheckUpgradeStatusRequest")
	proto.RegisterType((*CheckUpgradeStatusReply)(nil), "idl.CheckUpgradeStatusReply")
	proto.RegisterType((*FileSysUsage)(nil), "idl.FileSysUsage")
	proto.RegisterType((*CheckDiskUsageRequest)(nil), "idl.CheckDiskUsageRequest")
	proto.RegisterType((*CheckDiskUsageReply)(nil), "idl.CheckDiskUsageReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CommandListener service

type CommandListenerClient interface {
	CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error)
	CheckDiskUsage(ctx context.Context, in *CheckDiskUsageRequest, opts ...grpc.CallOption) (*CheckDiskUsageReply, error)
}

type commandListenerClient struct {
	cc *grpc.ClientConn
}

func NewCommandListenerClient(cc *grpc.ClientConn) CommandListenerClient {
	return &commandListenerClient{cc}
}

func (c *commandListenerClient) CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error) {
	out := new(CheckUpgradeStatusReply)
	err := grpc.Invoke(ctx, "/idl.CommandListener/CheckUpgradeStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandListenerClient) CheckDiskUsage(ctx context.Context, in *CheckDiskUsageRequest, opts ...grpc.CallOption) (*CheckDiskUsageReply, error) {
	out := new(CheckDiskUsageReply)
	err := grpc.Invoke(ctx, "/idl.CommandListener/CheckDiskUsage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommandListener service

type CommandListenerServer interface {
	CheckUpgradeStatus(context.Context, *CheckUpgradeStatusRequest) (*CheckUpgradeStatusReply, error)
	CheckDiskUsage(context.Context, *CheckDiskUsageRequest) (*CheckDiskUsageReply, error)
}

func RegisterCommandListenerServer(s *grpc.Server, srv CommandListenerServer) {
	s.RegisterService(&_CommandListener_serviceDesc, srv)
}

func _CommandListener_CheckUpgradeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUpgradeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandListenerServer).CheckUpgradeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CommandListener/CheckUpgradeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandListenerServer).CheckUpgradeStatus(ctx, req.(*CheckUpgradeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandListener_CheckDiskUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDiskUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandListenerServer).CheckDiskUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CommandListener/CheckDiskUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandListenerServer).CheckDiskUsage(ctx, req.(*CheckDiskUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommandListener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.CommandListener",
	HandlerType: (*CommandListenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUpgradeStatus",
			Handler:    _CommandListener_CheckUpgradeStatus_Handler,
		},
		{
			MethodName: "CheckDiskUsage",
			Handler:    _CommandListener_CheckDiskUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "command.proto",
}

func init() { proto.RegisterFile("command.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x4f, 0x4f, 0xbb, 0x40,
	0x14, 0xfc, 0xed, 0xaf, 0xd1, 0xc4, 0xd7, 0xaa, 0xf1, 0xf9, 0xa7, 0x88, 0xa6, 0x41, 0x4e, 0x9c,
	0x38, 0xd4, 0xab, 0xb7, 0x56, 0xe3, 0xc1, 0xc4, 0x84, 0xda, 0x9b, 0x09, 0x41, 0x78, 0xd4, 0x4d,
	0x97, 0x82, 0xbc, 0xe5, 0xc0, 0xc7, 0xf2, 0x1b, 0x9a, 0x85, 0x1a, 0xb1, 0xb6, 0xc7, 0x9d, 0xd9,
	0x99, 0xcc, 0xcc, 0x83, 0xc3, 0x38, 0xcf, 0xb2, 0x68, 0x95, 0xf8, 0x45, 0x99, 0xeb, 0x1c, 0x7b,
	0x32, 0x51, 0xee, 0x15, 0x5c, 0x4e, 0xde, 0x29, 0x5e, 0xce, 0x8b, 0x45, 0x19, 0x25, 0x34, 0xd3,
	0x91, 0xae, 0x38, 0xa0, 0x8f, 0x8a, 0x58, 0xbb, 0x77, 0x30, 0xdc, 0x46, 0x16, 0xaa, 0xc6, 0x1b,
	0x18, 0x14, 0x65, 0x1e, 0x13, 0x73, 0xa8, 0x24, 0x6b, 0x4b, 0x38, 0xc2, 0x3b, 0x08, 0xfa, 0x6b,
	0xec, 0x49, 0xb2, 0x76, 0xa7, 0x30, 0x78, 0x90, 0x8a, 0x66, 0x35, 0xcf, 0x39, 0x5a, 0x10, 0x8e,
	0x00, 0x52, 0xa9, 0x88, 0x6b, 0xd6, 0x94, 0xad, 0x05, 0x1d, 0x04, 0xcf, 0x60, 0xaf, 0x32, 0x1f,
	0xad, 0xff, 0x8e, 0xf0, 0x44, 0xd0, 0x3e, 0xdc, 0x21, 0x9c, 0x37, 0x19, 0xa6, 0x92, 0x97, 0x8d,
	0xcf, 0x77, 0xb8, 0x57, 0x38, 0xdd, 0x24, 0x4c, 0xb0, 0x7b, 0xb8, 0x30, 0x81, 0xc2, 0x3c, 0x0d,
	0x8d, 0x77, 0xc8, 0x35, 0x87, 0xad, 0xad, 0x70, 0x7a, 0x5e, 0x7f, 0x7c, 0xe2, 0xcb, 0x44, 0xf9,
	0xdd, 0x60, 0x01, 0x1a, 0xc1, 0x73, 0xda, 0xc5, 0xc6, 0x9f, 0x02, 0x8e, 0x27, 0xed, 0x5c, 0xa6,
	0x0c, 0xad, 0xa8, 0xc4, 0x17, 0xc0, 0xbf, 0x73, 0xe0, 0xa8, 0x31, 0xdc, 0x39, 0xa2, 0x7d, 0xbd,
	0x93, 0x2f, 0x54, 0xed, 0xfe, 0xc3, 0x47, 0x38, 0xfa, 0xdd, 0x03, 0xed, 0x1f, 0xc5, 0x66, 0x6b,
	0xdb, 0xda, 0xca, 0x35, 0x4e, 0x6f, 0xfb, 0xcd, 0x5d, 0x6f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x48, 0xd5, 0xb1, 0xff, 0xe8, 0x01, 0x00, 0x00,
}
