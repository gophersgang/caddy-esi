// Code generated by protoc-gen-go.
// source: grpc_data.proto
// DO NOT EDIT!

/*
Package esigrpc is a generated protocol buffer package.

It is generated from these files:
	grpc_data.proto

It has these top-level messages:
	ResourceArg
	HeaderBody
*/
package esigrpc

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

type ResourceArg struct {
	ExternalReq       *ResourceArg_ExternalReq `protobuf:"bytes,1,opt,name=external_req,json=externalReq" json:"external_req,omitempty"`
	Url               string                   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Timeout           int64                    `protobuf:"varint,3,opt,name=timeout" json:"timeout,omitempty"`
	MaxBodySize       uint64                   `protobuf:"varint,34,opt,name=max_body_size,json=maxBodySize" json:"max_body_size,omitempty"`
	Ttl               int64                    `protobuf:"varint,5,opt,name=ttl" json:"ttl,omitempty"`
	Key               string                   `protobuf:"bytes,6,opt,name=key" json:"key,omitempty"`
	ForwardHeaders    []string                 `protobuf:"bytes,7,rep,name=forward_headers,json=forwardHeaders" json:"forward_headers,omitempty"`
	ForwardHeadersAll bool                     `protobuf:"varint,8,opt,name=forward_headers_all,json=forwardHeadersAll" json:"forward_headers_all,omitempty"`
	ReturnHeaders     []string                 `protobuf:"bytes,9,rep,name=return_headers,json=returnHeaders" json:"return_headers,omitempty"`
	ReturnHeadersAll  bool                     `protobuf:"varint,10,opt,name=return_headers_all,json=returnHeadersAll" json:"return_headers_all,omitempty"`
}

func (m *ResourceArg) Reset()                    { *m = ResourceArg{} }
func (m *ResourceArg) String() string            { return proto.CompactTextString(m) }
func (*ResourceArg) ProtoMessage()               {}
func (*ResourceArg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResourceArg) GetExternalReq() *ResourceArg_ExternalReq {
	if m != nil {
		return m.ExternalReq
	}
	return nil
}

func (m *ResourceArg) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ResourceArg) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ResourceArg) GetMaxBodySize() uint64 {
	if m != nil {
		return m.MaxBodySize
	}
	return 0
}

func (m *ResourceArg) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *ResourceArg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ResourceArg) GetForwardHeaders() []string {
	if m != nil {
		return m.ForwardHeaders
	}
	return nil
}

func (m *ResourceArg) GetForwardHeadersAll() bool {
	if m != nil {
		return m.ForwardHeadersAll
	}
	return false
}

func (m *ResourceArg) GetReturnHeaders() []string {
	if m != nil {
		return m.ReturnHeaders
	}
	return nil
}

func (m *ResourceArg) GetReturnHeadersAll() bool {
	if m != nil {
		return m.ReturnHeadersAll
	}
	return false
}

type ResourceArg_ExternalReq struct {
	Method           string   `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Url              string   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Proto            string   `protobuf:"bytes,3,opt,name=proto" json:"proto,omitempty"`
	ProtoMajor       int32    `protobuf:"varint,4,opt,name=proto_major,json=protoMajor" json:"proto_major,omitempty"`
	ProtoMinor       int32    `protobuf:"varint,5,opt,name=proto_minor,json=protoMinor" json:"proto_minor,omitempty"`
	Header           []string `protobuf:"bytes,6,rep,name=header" json:"header,omitempty"`
	ContentLength    int64    `protobuf:"varint,7,opt,name=ContentLength" json:"ContentLength,omitempty"`
	TransferEncoding []string `protobuf:"bytes,8,rep,name=transfer_encoding,json=transferEncoding" json:"transfer_encoding,omitempty"`
	Close            bool     `protobuf:"varint,9,opt,name=close" json:"close,omitempty"`
	Host             string   `protobuf:"bytes,10,opt,name=host" json:"host,omitempty"`
	Form             []string `protobuf:"bytes,11,rep,name=form" json:"form,omitempty"`
	PostForm         []string `protobuf:"bytes,12,rep,name=post_form,json=postForm" json:"post_form,omitempty"`
	RemoteAddr       string   `protobuf:"bytes,13,opt,name=remote_addr,json=remoteAddr" json:"remote_addr,omitempty"`
	RequestUri       string   `protobuf:"bytes,14,opt,name=request_uri,json=requestUri" json:"request_uri,omitempty"`
}

func (m *ResourceArg_ExternalReq) Reset()                    { *m = ResourceArg_ExternalReq{} }
func (m *ResourceArg_ExternalReq) String() string            { return proto.CompactTextString(m) }
func (*ResourceArg_ExternalReq) ProtoMessage()               {}
func (*ResourceArg_ExternalReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ResourceArg_ExternalReq) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ResourceArg_ExternalReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ResourceArg_ExternalReq) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *ResourceArg_ExternalReq) GetProtoMajor() int32 {
	if m != nil {
		return m.ProtoMajor
	}
	return 0
}

func (m *ResourceArg_ExternalReq) GetProtoMinor() int32 {
	if m != nil {
		return m.ProtoMinor
	}
	return 0
}

func (m *ResourceArg_ExternalReq) GetHeader() []string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ResourceArg_ExternalReq) GetContentLength() int64 {
	if m != nil {
		return m.ContentLength
	}
	return 0
}

func (m *ResourceArg_ExternalReq) GetTransferEncoding() []string {
	if m != nil {
		return m.TransferEncoding
	}
	return nil
}

func (m *ResourceArg_ExternalReq) GetClose() bool {
	if m != nil {
		return m.Close
	}
	return false
}

func (m *ResourceArg_ExternalReq) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ResourceArg_ExternalReq) GetForm() []string {
	if m != nil {
		return m.Form
	}
	return nil
}

func (m *ResourceArg_ExternalReq) GetPostForm() []string {
	if m != nil {
		return m.PostForm
	}
	return nil
}

func (m *ResourceArg_ExternalReq) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *ResourceArg_ExternalReq) GetRequestUri() string {
	if m != nil {
		return m.RequestUri
	}
	return ""
}

// HeaderBody a return value like in the backend package but here the three
// return arguments are combined into one struct. If error is empty, then no
// error has happend.
type HeaderBody struct {
	Header []string `protobuf:"bytes,1,rep,name=header" json:"header,omitempty"`
	Body   []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HeaderBody) Reset()                    { *m = HeaderBody{} }
func (m *HeaderBody) String() string            { return proto.CompactTextString(m) }
func (*HeaderBody) ProtoMessage()               {}
func (*HeaderBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HeaderBody) GetHeader() []string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HeaderBody) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceArg)(nil), "esigrpc.ResourceArg")
	proto.RegisterType((*ResourceArg_ExternalReq)(nil), "esigrpc.ResourceArg.ExternalReq")
	proto.RegisterType((*HeaderBody)(nil), "esigrpc.HeaderBody")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HeaderBodyService service

type HeaderBodyServiceClient interface {
	GetHeaderBody(ctx context.Context, in *ResourceArg, opts ...grpc.CallOption) (*HeaderBody, error)
}

type headerBodyServiceClient struct {
	cc *grpc.ClientConn
}

func NewHeaderBodyServiceClient(cc *grpc.ClientConn) HeaderBodyServiceClient {
	return &headerBodyServiceClient{cc}
}

func (c *headerBodyServiceClient) GetHeaderBody(ctx context.Context, in *ResourceArg, opts ...grpc.CallOption) (*HeaderBody, error) {
	out := new(HeaderBody)
	err := grpc.Invoke(ctx, "/esigrpc.HeaderBodyService/GetHeaderBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HeaderBodyService service

type HeaderBodyServiceServer interface {
	GetHeaderBody(context.Context, *ResourceArg) (*HeaderBody, error)
}

func RegisterHeaderBodyServiceServer(s *grpc.Server, srv HeaderBodyServiceServer) {
	s.RegisterService(&_HeaderBodyService_serviceDesc, srv)
}

func _HeaderBodyService_GetHeaderBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeaderBodyServiceServer).GetHeaderBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/esigrpc.HeaderBodyService/GetHeaderBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeaderBodyServiceServer).GetHeaderBody(ctx, req.(*ResourceArg))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeaderBodyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "esigrpc.HeaderBodyService",
	HandlerType: (*HeaderBodyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHeaderBody",
			Handler:    _HeaderBodyService_GetHeaderBody_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_data.proto",
}

func init() { proto.RegisterFile("grpc_data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4d, 0x6f, 0xd3, 0x4e,
	0x10, 0xc6, 0xff, 0xfe, 0x27, 0x4d, 0xe2, 0x71, 0xd3, 0x97, 0x6d, 0x85, 0x96, 0x72, 0xa8, 0x15,
	0x81, 0xb0, 0x44, 0xe5, 0x43, 0xb9, 0x70, 0xe0, 0x92, 0x56, 0xe5, 0x45, 0x02, 0xa9, 0x6c, 0xc5,
	0xd9, 0xda, 0x7a, 0x27, 0xb1, 0xa9, 0xed, 0x6d, 0xd7, 0x6b, 0x68, 0xfa, 0x21, 0xb8, 0xf2, 0x75,
	0xd1, 0x8e, 0x9d, 0x26, 0x41, 0xbd, 0xcd, 0xfc, 0xe6, 0xf1, 0x93, 0xdd, 0x67, 0x36, 0xb0, 0x3b,
	0x37, 0xb7, 0x69, 0xa2, 0xa4, 0x95, 0xf1, 0xad, 0xd1, 0x56, 0xb3, 0x21, 0xd6, 0xb9, 0x63, 0x93,
	0xdf, 0x03, 0x08, 0x04, 0xd6, 0xba, 0x31, 0x29, 0x4e, 0xcd, 0x9c, 0x9d, 0xc3, 0x36, 0xde, 0x5b,
	0x34, 0x95, 0x2c, 0x12, 0x83, 0x77, 0xdc, 0x0b, 0xbd, 0x28, 0x38, 0x0d, 0xe3, 0x4e, 0x1f, 0xaf,
	0x69, 0xe3, 0x8b, 0x4e, 0x28, 0xf0, 0x4e, 0x04, 0xb8, 0x6a, 0xd8, 0x1e, 0xf4, 0x1a, 0x53, 0xf0,
	0xff, 0x43, 0x2f, 0xf2, 0x85, 0x2b, 0x19, 0x87, 0xa1, 0xcd, 0x4b, 0xd4, 0x8d, 0xe5, 0xbd, 0xd0,
	0x8b, 0x7a, 0x62, 0xd9, 0xb2, 0x09, 0x8c, 0x4b, 0x79, 0x9f, 0x5c, 0x6b, 0xb5, 0x48, 0xea, 0xfc,
	0x01, 0xf9, 0x24, 0xf4, 0xa2, 0xbe, 0x08, 0x4a, 0x79, 0x7f, 0xa6, 0xd5, 0xe2, 0x2a, 0x7f, 0x40,
	0xe7, 0x67, 0x6d, 0xc1, 0xb7, 0xe8, 0x4b, 0x57, 0x3a, 0x72, 0x83, 0x0b, 0x3e, 0x68, 0x7f, 0xe1,
	0x06, 0x17, 0xec, 0x35, 0xec, 0xce, 0xb4, 0xf9, 0x25, 0x8d, 0x4a, 0x32, 0x94, 0x0a, 0x4d, 0xcd,
	0x87, 0x61, 0x2f, 0xf2, 0xc5, 0x4e, 0x87, 0x3f, 0xb5, 0x94, 0xc5, 0x70, 0xf0, 0x8f, 0x30, 0x91,
	0x45, 0xc1, 0x47, 0xa1, 0x17, 0x8d, 0xc4, 0xfe, 0xa6, 0x78, 0x5a, 0x14, 0xec, 0x15, 0xec, 0x18,
	0xb4, 0x8d, 0xa9, 0x1e, 0x7d, 0x7d, 0xf2, 0x1d, 0xb7, 0x74, 0x69, 0x7b, 0x02, 0x6c, 0x53, 0x46,
	0xae, 0x40, 0xae, 0x7b, 0x1b, 0xd2, 0x69, 0x51, 0x1c, 0xfd, 0xe9, 0x41, 0xb0, 0x16, 0x1f, 0x7b,
	0x06, 0x83, 0x12, 0x6d, 0xa6, 0x15, 0x05, 0xee, 0x8b, 0xae, 0x7b, 0x22, 0xc9, 0x43, 0xd8, 0xa2,
	0x15, 0x52, 0x8e, 0xbe, 0x68, 0x1b, 0x76, 0x0c, 0x01, 0x15, 0x49, 0x29, 0x7f, 0x68, 0xc3, 0xfb,
	0xa1, 0x17, 0x6d, 0x09, 0x20, 0xf4, 0xd5, 0x91, 0x35, 0x41, 0x5e, 0x69, 0x43, 0x51, 0x3e, 0x0a,
	0x1c, 0x71, 0x27, 0x68, 0x0f, 0xce, 0x07, 0x74, 0xbd, 0xae, 0x63, 0x2f, 0x61, 0x7c, 0xae, 0x2b,
	0x8b, 0x95, 0xfd, 0x82, 0xd5, 0xdc, 0x66, 0x7c, 0x48, 0x5b, 0xd8, 0x84, 0xec, 0x0d, 0xec, 0x5b,
	0x23, 0xab, 0x7a, 0x86, 0x26, 0xc1, 0x2a, 0xd5, 0x2a, 0xaf, 0xe6, 0x7c, 0x44, 0x46, 0x7b, 0xcb,
	0xc1, 0x45, 0xc7, 0xdd, 0x15, 0xd2, 0x42, 0xd7, 0xc8, 0x7d, 0x4a, 0xa7, 0x6d, 0x18, 0x83, 0x7e,
	0xa6, 0x6b, 0x4b, 0x91, 0xf9, 0x82, 0x6a, 0xc7, 0x66, 0xda, 0x94, 0x3c, 0x20, 0x27, 0xaa, 0xd9,
	0x0b, 0xf0, 0x6f, 0x75, 0x6d, 0x13, 0x1a, 0x6c, 0xd3, 0x60, 0xe4, 0xc0, 0x07, 0x37, 0x3c, 0x86,
	0xc0, 0x60, 0xa9, 0x2d, 0x26, 0x52, 0x29, 0xc3, 0xc7, 0xe4, 0x05, 0x2d, 0x9a, 0x2a, 0x65, 0x5a,
	0xc1, 0x5d, 0x83, 0xb5, 0x4d, 0x1a, 0x93, 0xf3, 0x9d, 0xa5, 0x80, 0xd0, 0x77, 0x93, 0x4f, 0xde,
	0x01, 0xb4, 0x7b, 0x72, 0xaf, 0x6f, 0x2d, 0x15, 0x6f, 0x23, 0x15, 0x06, 0x7d, 0xf7, 0x62, 0x69,
	0x31, 0xdb, 0x82, 0xea, 0xd3, 0x6f, 0xb0, 0xbf, 0xfa, 0xf2, 0x0a, 0xcd, 0xcf, 0x3c, 0x45, 0xf6,
	0x1e, 0xc6, 0x1f, 0xd1, 0xae, 0x39, 0x1e, 0x3e, 0xf5, 0x57, 0x3a, 0x3a, 0x78, 0xa4, 0x2b, 0xe9,
	0xe4, 0xbf, 0xb3, 0x13, 0x78, 0x3e, 0x2b, 0xe3, 0x3a, 0xcd, 0x9a, 0x52, 0xa6, 0x19, 0x9a, 0x59,
	0x19, 0xa7, 0x52, 0xa9, 0x85, 0x53, 0x9f, 0x8d, 0xcf, 0x5d, 0x79, 0x71, 0xf5, 0xf9, 0xd2, 0x6d,
	0xf1, 0xd2, 0xbb, 0x1e, 0xd0, 0x3a, 0xdf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xc9, 0x6d,
	0xa0, 0xee, 0x03, 0x00, 0x00,
}
