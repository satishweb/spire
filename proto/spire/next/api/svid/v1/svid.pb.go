// Code generated by protoc-gen-go. DO NOT EDIT.
// source: svid.proto

package svid

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	svid "github.com/spiffe/spire/proto/spire/v2/types/svid"
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

type MintX509SVIDRequest struct {
	// Required. SPIFFE ID of the X509-SVID.
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Required. ASN.1 DER encoded CSR. The CSR is only used to convey the
	// public key and prove possession of the private key. The rest of the CSR
	// is ignored.
	Csr []byte `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	// TTL of the X509-SVID, in seconds. The server default will be used if
	// unset. The TTL is advisory only. The actual lifetime of the X509-SVID
	// may be lower depending on the remaining lifetime of the active SPIRE
	// Server CA.
	Ttl int32 `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// DNS names to include as DNS SANs in the X509-SVID. If set, the first
	// in the list is also set as the X509-SVID common name.
	DnsNames             []string `protobuf:"bytes,4,rep,name=dns_names,json=dnsNames,proto3" json:"dns_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MintX509SVIDRequest) Reset()         { *m = MintX509SVIDRequest{} }
func (m *MintX509SVIDRequest) String() string { return proto.CompactTextString(m) }
func (*MintX509SVIDRequest) ProtoMessage()    {}
func (*MintX509SVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{0}
}

func (m *MintX509SVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintX509SVIDRequest.Unmarshal(m, b)
}
func (m *MintX509SVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintX509SVIDRequest.Marshal(b, m, deterministic)
}
func (m *MintX509SVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintX509SVIDRequest.Merge(m, src)
}
func (m *MintX509SVIDRequest) XXX_Size() int {
	return xxx_messageInfo_MintX509SVIDRequest.Size(m)
}
func (m *MintX509SVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintX509SVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintX509SVIDRequest proto.InternalMessageInfo

func (m *MintX509SVIDRequest) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *MintX509SVIDRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *MintX509SVIDRequest) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *MintX509SVIDRequest) GetDnsNames() []string {
	if m != nil {
		return m.DnsNames
	}
	return nil
}

type MintX509SVIDResponse struct {
	// The newly issued X509-SVID.
	Svid                 *svid.X509SVID `protobuf:"bytes,1,opt,name=svid,proto3" json:"svid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MintX509SVIDResponse) Reset()         { *m = MintX509SVIDResponse{} }
func (m *MintX509SVIDResponse) String() string { return proto.CompactTextString(m) }
func (*MintX509SVIDResponse) ProtoMessage()    {}
func (*MintX509SVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{1}
}

func (m *MintX509SVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintX509SVIDResponse.Unmarshal(m, b)
}
func (m *MintX509SVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintX509SVIDResponse.Marshal(b, m, deterministic)
}
func (m *MintX509SVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintX509SVIDResponse.Merge(m, src)
}
func (m *MintX509SVIDResponse) XXX_Size() int {
	return xxx_messageInfo_MintX509SVIDResponse.Size(m)
}
func (m *MintX509SVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MintX509SVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MintX509SVIDResponse proto.InternalMessageInfo

func (m *MintX509SVIDResponse) GetSvid() *svid.X509SVID {
	if m != nil {
		return m.Svid
	}
	return nil
}

type MintJWTSVIDRequest struct {
	// Required. SPIFFE ID of the JWT-SVID.
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Required. List of audience claims to include in the JWT-SVID. At least one must
	// be set.
	Audience []string `protobuf:"bytes,2,rep,name=audience,proto3" json:"audience,omitempty"`
	// TTL of the JWT-SVID, in seconds. The server default will be used if
	// unset. The TTL is advisory only. The actual lifetime of the JWT-SVID may
	// be lower depending on the remaining lifetime of the active SPIRE Server
	// CA.
	Ttl                  int32    `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MintJWTSVIDRequest) Reset()         { *m = MintJWTSVIDRequest{} }
func (m *MintJWTSVIDRequest) String() string { return proto.CompactTextString(m) }
func (*MintJWTSVIDRequest) ProtoMessage()    {}
func (*MintJWTSVIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{2}
}

func (m *MintJWTSVIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintJWTSVIDRequest.Unmarshal(m, b)
}
func (m *MintJWTSVIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintJWTSVIDRequest.Marshal(b, m, deterministic)
}
func (m *MintJWTSVIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintJWTSVIDRequest.Merge(m, src)
}
func (m *MintJWTSVIDRequest) XXX_Size() int {
	return xxx_messageInfo_MintJWTSVIDRequest.Size(m)
}
func (m *MintJWTSVIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MintJWTSVIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MintJWTSVIDRequest proto.InternalMessageInfo

func (m *MintJWTSVIDRequest) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *MintJWTSVIDRequest) GetAudience() []string {
	if m != nil {
		return m.Audience
	}
	return nil
}

func (m *MintJWTSVIDRequest) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type MintJWTSVIDResponse struct {
	// The newly issued JWT-SVID.
	Svid                 *svid.JWTSVID `protobuf:"bytes,1,opt,name=svid,proto3" json:"svid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MintJWTSVIDResponse) Reset()         { *m = MintJWTSVIDResponse{} }
func (m *MintJWTSVIDResponse) String() string { return proto.CompactTextString(m) }
func (*MintJWTSVIDResponse) ProtoMessage()    {}
func (*MintJWTSVIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3513929f2a405cf, []int{3}
}

func (m *MintJWTSVIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MintJWTSVIDResponse.Unmarshal(m, b)
}
func (m *MintJWTSVIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MintJWTSVIDResponse.Marshal(b, m, deterministic)
}
func (m *MintJWTSVIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintJWTSVIDResponse.Merge(m, src)
}
func (m *MintJWTSVIDResponse) XXX_Size() int {
	return xxx_messageInfo_MintJWTSVIDResponse.Size(m)
}
func (m *MintJWTSVIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MintJWTSVIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MintJWTSVIDResponse proto.InternalMessageInfo

func (m *MintJWTSVIDResponse) GetSvid() *svid.JWTSVID {
	if m != nil {
		return m.Svid
	}
	return nil
}

func init() {
	proto.RegisterType((*MintX509SVIDRequest)(nil), "spire.api.svid.v1.MintX509SVIDRequest")
	proto.RegisterType((*MintX509SVIDResponse)(nil), "spire.api.svid.v1.MintX509SVIDResponse")
	proto.RegisterType((*MintJWTSVIDRequest)(nil), "spire.api.svid.v1.MintJWTSVIDRequest")
	proto.RegisterType((*MintJWTSVIDResponse)(nil), "spire.api.svid.v1.MintJWTSVIDResponse")
}

func init() { proto.RegisterFile("svid.proto", fileDescriptor_a3513929f2a405cf) }

var fileDescriptor_a3513929f2a405cf = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4f, 0x32, 0x31,
	0x10, 0xc7, 0xb3, 0xc0, 0xf3, 0x04, 0x0a, 0x07, 0xad, 0x1e, 0xd6, 0xf5, 0xb2, 0x59, 0x15, 0xb9,
	0xd8, 0x02, 0x86, 0x03, 0xe1, 0x66, 0x88, 0x09, 0x26, 0x7a, 0x58, 0x8d, 0x1a, 0x63, 0xb2, 0x59,
	0xd8, 0xa2, 0x35, 0xd2, 0xad, 0x4c, 0x41, 0xfc, 0x80, 0x7e, 0x2f, 0xd3, 0x96, 0xd7, 0x80, 0xc1,
	0xd3, 0x76, 0x66, 0xff, 0x33, 0xf3, 0x9b, 0x17, 0x84, 0x60, 0xcc, 0x13, 0x22, 0x87, 0xa9, 0x4a,
	0xf1, 0x2e, 0x48, 0x3e, 0x64, 0x24, 0x96, 0x9c, 0x18, 0xef, 0xb8, 0xe6, 0x1d, 0x19, 0x17, 0x15,
	0x6c, 0xa2, 0xa8, 0xfa, 0x92, 0x0c, 0xa8, 0xfe, 0x43, 0xdf, 0x3e, 0xd5, 0x22, 0xce, 0x3b, 0xde,
	0x2c, 0x9a, 0x34, 0xaa, 0xcd, 0x85, 0x2a, 0x00, 0xb4, 0x77, 0xcd, 0x85, 0x7a, 0x6c, 0x54, 0x9b,
	0xb7, 0xf7, 0x9d, 0x76, 0xc8, 0x3e, 0x46, 0x0c, 0x14, 0x3e, 0x44, 0x05, 0x90, 0xbc, 0xdf, 0x67,
	0x11, 0x4f, 0x5c, 0xc7, 0x77, 0x2a, 0x85, 0x30, 0x6f, 0x1d, 0x9d, 0x04, 0xef, 0xa0, 0x6c, 0x0f,
	0x86, 0x6e, 0xc6, 0x77, 0x2a, 0xa5, 0x50, 0x3f, 0xb5, 0x47, 0xa9, 0x77, 0x37, 0xeb, 0x3b, 0x95,
	0x7f, 0xa1, 0x7e, 0xea, 0x04, 0x89, 0x80, 0x48, 0xc4, 0x03, 0x06, 0x6e, 0xce, 0xcf, 0xea, 0x04,
	0x89, 0x80, 0x1b, 0x6d, 0x07, 0x97, 0x68, 0x7f, 0xb5, 0x28, 0xc8, 0x54, 0x00, 0xc3, 0x04, 0xe5,
	0x34, 0x9a, 0x29, 0x58, 0xac, 0x7b, 0xc4, 0x76, 0x6e, 0xe0, 0x6d, 0xef, 0xf3, 0x08, 0xa3, 0x0b,
	0x22, 0x84, 0x75, 0x9e, 0xab, 0x87, 0xbb, 0x3f, 0xb3, 0x7b, 0x28, 0x1f, 0x8f, 0x12, 0xce, 0x44,
	0x8f, 0xb9, 0x19, 0x8b, 0x35, 0xb3, 0xd7, 0xbb, 0x08, 0xda, 0x76, 0x3a, 0xf3, 0x02, 0x53, 0xce,
	0xb3, 0x15, 0xce, 0x83, 0x75, 0xce, 0x59, 0x80, 0x91, 0xd5, 0xbf, 0x1d, 0x94, 0xd3, 0x26, 0x8e,
	0x50, 0x69, 0xb9, 0x6f, 0x5c, 0x26, 0x6b, 0xbb, 0x25, 0x1b, 0xb6, 0xe1, 0x9d, 0x6e, 0xd5, 0x4d,
	0xc1, 0x9e, 0x51, 0x71, 0x89, 0x17, 0x9f, 0xfc, 0x12, 0xb7, 0x3a, 0x30, 0xaf, 0xbc, 0x4d, 0x66,
	0xb3, 0x5f, 0xb4, 0x9e, 0x9a, 0x2f, 0x5c, 0xbd, 0x8e, 0xba, 0xa4, 0x97, 0x0e, 0xa8, 0x1d, 0x29,
	0xb5, 0x57, 0x66, 0x8e, 0x89, 0x2e, 0x5d, 0x5c, 0x2c, 0xb9, 0xbd, 0xb7, 0x71, 0xad, 0xa5, 0xbf,
	0xdd, 0xff, 0x46, 0x72, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xcc, 0x32, 0x2f, 0xdb, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SVIDClient is the client API for SVID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SVIDClient interface {
	// MintX509SVID creates an X509-SVID.
	MintX509SVID(ctx context.Context, in *MintX509SVIDRequest, opts ...grpc.CallOption) (*MintX509SVIDResponse, error)
	// MintX509SVID creates an JWT-SVID.
	MintJWTSVID(ctx context.Context, in *MintJWTSVIDRequest, opts ...grpc.CallOption) (*MintJWTSVIDResponse, error)
}

type sVIDClient struct {
	cc *grpc.ClientConn
}

func NewSVIDClient(cc *grpc.ClientConn) SVIDClient {
	return &sVIDClient{cc}
}

func (c *sVIDClient) MintX509SVID(ctx context.Context, in *MintX509SVIDRequest, opts ...grpc.CallOption) (*MintX509SVIDResponse, error) {
	out := new(MintX509SVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.svid.v1.SVID/MintX509SVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sVIDClient) MintJWTSVID(ctx context.Context, in *MintJWTSVIDRequest, opts ...grpc.CallOption) (*MintJWTSVIDResponse, error) {
	out := new(MintJWTSVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.svid.v1.SVID/MintJWTSVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SVIDServer is the server API for SVID service.
type SVIDServer interface {
	// MintX509SVID creates an X509-SVID.
	MintX509SVID(context.Context, *MintX509SVIDRequest) (*MintX509SVIDResponse, error)
	// MintX509SVID creates an JWT-SVID.
	MintJWTSVID(context.Context, *MintJWTSVIDRequest) (*MintJWTSVIDResponse, error)
}

// UnimplementedSVIDServer can be embedded to have forward compatible implementations.
type UnimplementedSVIDServer struct {
}

func (*UnimplementedSVIDServer) MintX509SVID(ctx context.Context, req *MintX509SVIDRequest) (*MintX509SVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintX509SVID not implemented")
}
func (*UnimplementedSVIDServer) MintJWTSVID(ctx context.Context, req *MintJWTSVIDRequest) (*MintJWTSVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintJWTSVID not implemented")
}

func RegisterSVIDServer(s *grpc.Server, srv SVIDServer) {
	s.RegisterService(&_SVID_serviceDesc, srv)
}

func _SVID_MintX509SVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintX509SVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).MintX509SVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.svid.v1.SVID/MintX509SVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).MintX509SVID(ctx, req.(*MintX509SVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SVID_MintJWTSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintJWTSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SVIDServer).MintJWTSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.svid.v1.SVID/MintJWTSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SVIDServer).MintJWTSVID(ctx, req.(*MintJWTSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SVID_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.svid.v1.SVID",
	HandlerType: (*SVIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintX509SVID",
			Handler:    _SVID_MintX509SVID_Handler,
		},
		{
			MethodName: "MintJWTSVID",
			Handler:    _SVID_MintJWTSVID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svid.proto",
}
