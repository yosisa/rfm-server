// Code generated by protoc-gen-go.
// source: rfm.proto
// DO NOT EDIT!

/*
Package rfm is a generated protocol buffer package.

It is generated from these files:
	rfm.proto

It has these top-level messages:
	Request
	FindRequest
	DirInfo
	DiskUsage
	FileInfo
*/
package rfm

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
const _ = proto.ProtoPackageIsVersion1

type Request struct {
	BaseDir string `protobuf:"bytes,1,opt,name=base_dir" json:"base_dir,omitempty"`
	Target  string `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FindRequest struct {
	BaseDir string `protobuf:"bytes,1,opt,name=base_dir" json:"base_dir,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *FindRequest) Reset()                    { *m = FindRequest{} }
func (m *FindRequest) String() string            { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()               {}
func (*FindRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DirInfo struct {
	Path      string      `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	DiskUsage *DiskUsage  `protobuf:"bytes,2,opt,name=disk_usage" json:"disk_usage,omitempty"`
	Items     []*FileInfo `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
}

func (m *DirInfo) Reset()                    { *m = DirInfo{} }
func (m *DirInfo) String() string            { return proto.CompactTextString(m) }
func (*DirInfo) ProtoMessage()               {}
func (*DirInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DirInfo) GetDiskUsage() *DiskUsage {
	if m != nil {
		return m.DiskUsage
	}
	return nil
}

func (m *DirInfo) GetItems() []*FileInfo {
	if m != nil {
		return m.Items
	}
	return nil
}

type DiskUsage struct {
	Size uint64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Free uint64 `protobuf:"varint,2,opt,name=free" json:"free,omitempty"`
}

func (m *DiskUsage) Reset()                    { *m = DiskUsage{} }
func (m *DiskUsage) String() string            { return proto.CompactTextString(m) }
func (*DiskUsage) ProtoMessage()               {}
func (*DiskUsage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type FileInfo struct {
	Name    string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size    int64   `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Mode    uint32  `protobuf:"varint,3,opt,name=mode" json:"mode,omitempty"`
	ModTime float64 `protobuf:"fixed64,4,opt,name=mod_time" json:"mod_time,omitempty"`
	IsDir   bool    `protobuf:"varint,5,opt,name=is_dir" json:"is_dir,omitempty"`
	Owner   string  `protobuf:"bytes,6,opt,name=owner" json:"owner,omitempty"`
}

func (m *FileInfo) Reset()                    { *m = FileInfo{} }
func (m *FileInfo) String() string            { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()               {}
func (*FileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Request)(nil), "rfm.Request")
	proto.RegisterType((*FindRequest)(nil), "rfm.FindRequest")
	proto.RegisterType((*DirInfo)(nil), "rfm.DirInfo")
	proto.RegisterType((*DiskUsage)(nil), "rfm.DiskUsage")
	proto.RegisterType((*FileInfo)(nil), "rfm.FileInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for FS service

type FSClient interface {
	ReadDir(ctx context.Context, in *Request, opts ...grpc.CallOption) (*DirInfo, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*DirInfo, error)
}

type fSClient struct {
	cc *grpc.ClientConn
}

func NewFSClient(cc *grpc.ClientConn) FSClient {
	return &fSClient{cc}
}

func (c *fSClient) ReadDir(ctx context.Context, in *Request, opts ...grpc.CallOption) (*DirInfo, error) {
	out := new(DirInfo)
	err := grpc.Invoke(ctx, "/rfm.FS/ReadDir", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fSClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*DirInfo, error) {
	out := new(DirInfo)
	err := grpc.Invoke(ctx, "/rfm.FS/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FS service

type FSServer interface {
	ReadDir(context.Context, *Request) (*DirInfo, error)
	Find(context.Context, *FindRequest) (*DirInfo, error)
}

func RegisterFSServer(s *grpc.Server, srv FSServer) {
	s.RegisterService(&_FS_serviceDesc, srv)
}

func _FS_ReadDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FSServer).ReadDir(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FS_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FSServer).Find(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _FS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rfm.FS",
	HandlerType: (*FSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadDir",
			Handler:    _FS_ReadDir_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _FS_Find_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x51, 0xcd, 0x6a, 0xf3, 0x30,
	0x10, 0xfc, 0x1c, 0xff, 0xc4, 0xde, 0xd8, 0x21, 0xe8, 0x64, 0x3e, 0x7a, 0x08, 0xba, 0xc4, 0x50,
	0x9a, 0x43, 0xfa, 0x0a, 0xc1, 0xd0, 0x6b, 0x4b, 0x29, 0x3d, 0x19, 0x05, 0xaf, 0x53, 0xd1, 0xda,
	0x4e, 0x25, 0x85, 0x42, 0x9f, 0xbe, 0x5a, 0x45, 0x75, 0xa1, 0xd0, 0x9b, 0x67, 0x77, 0x76, 0x3c,
	0x33, 0x82, 0x4c, 0x75, 0xfd, 0xf6, 0xa4, 0x46, 0x33, 0xb2, 0xd0, 0x7e, 0xf2, 0x6b, 0x98, 0xdf,
	0xe3, 0xfb, 0x19, 0xb5, 0x61, 0x2b, 0x48, 0x0f, 0x42, 0x63, 0xd3, 0x4a, 0x55, 0x06, 0xeb, 0xa0,
	0xca, 0xd8, 0x12, 0x12, 0x23, 0xd4, 0x11, 0x4d, 0x39, 0x23, 0xcc, 0x6f, 0x60, 0x51, 0xcb, 0xa1,
	0xfd, 0xfb, 0x20, 0x87, 0x68, 0x10, 0x3d, 0x7a, 0xfa, 0x33, 0xcc, 0xf7, 0x52, 0xdd, 0x0d, 0xdd,
	0x48, 0x8b, 0x93, 0x30, 0x2f, 0x9e, 0xc6, 0x01, 0x5a, 0xa9, 0x5f, 0x9b, 0xb3, 0x16, 0xc7, 0x0b,
	0x79, 0xb1, 0x5b, 0x6e, 0xc9, 0xd9, 0xde, 0x8e, 0x1f, 0x69, 0xca, 0xae, 0x20, 0x96, 0x06, 0x7b,
	0x5d, 0x86, 0xeb, 0xd0, 0xae, 0x0b, 0xb7, 0xae, 0xe5, 0x1b, 0x92, 0x1e, 0xdf, 0x40, 0xf6, 0x43,
	0xb5, 0xe2, 0x5a, 0x7e, 0xa2, 0x13, 0x8f, 0x08, 0x75, 0x0a, 0x2f, 0xb2, 0x11, 0x47, 0x48, 0xbf,
	0x8f, 0x26, 0x77, 0x93, 0x57, 0x77, 0x45, 0xbc, 0x90, 0x50, 0x3f, 0xb6, 0x68, 0xff, 0x16, 0x54,
	0x05, 0x25, 0xb3, 0xa8, 0x31, 0xd2, 0xb2, 0x23, 0x3b, 0x09, 0xa8, 0x0a, 0xa9, 0x5d, 0xd2, 0xd8,
	0xe2, 0x94, 0x15, 0x10, 0x8f, 0x1f, 0x03, 0xaa, 0x32, 0x21, 0xb1, 0xdd, 0x13, 0xcc, 0xea, 0x07,
	0xb6, 0xa1, 0x32, 0x45, 0x6b, 0x43, 0xb3, 0xdc, 0xf9, 0xf5, 0x4d, 0xfd, 0xcf, 0x7d, 0x38, 0x57,
	0x06, 0xff, 0xc7, 0x2a, 0x88, 0xa8, 0x48, 0xb6, 0xf2, 0xa9, 0xa6, 0x4e, 0x7f, 0x33, 0x0f, 0x89,
	0x7b, 0xab, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x4a, 0x30, 0x11, 0xb8, 0x01, 0x00,
	0x00,
}
