// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authorize.proto

package authorize

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

type IsAuthorizedRequest struct {
	// User Context
	//
	UserToken string `protobuf:"bytes,1,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
	// Request Context
	//
	// Method specifies the HTTP method (GET, POST, PUT, etc.).
	RequestMethod string `protobuf:"bytes,2,opt,name=request_method,json=requestMethod,proto3" json:"request_method,omitempty"`
	// URL specifies either the URI being requested
	RequestUrl string `protobuf:"bytes,3,opt,name=request_url,json=requestUrl,proto3" json:"request_url,omitempty"`
	// host specifies the host on which the URL per RFC 7230, section 5.4
	RequestHost string `protobuf:"bytes,4,opt,name=request_host,json=requestHost,proto3" json:"request_host,omitempty"`
	// request_uri is the unmodified request-target of the
	// Request-Line (RFC 7230, Section 3.1.1) as sent by the client
	RequestRequestUri string `protobuf:"bytes,5,opt,name=request_request_uri,json=requestRequestUri,proto3" json:"request_request_uri,omitempty"`
	// RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	RequestRemoteAddr    string                                  `protobuf:"bytes,6,opt,name=request_remote_addr,json=requestRemoteAddr,proto3" json:"request_remote_addr,omitempty"`
	RequestHeaders       map[string]*IsAuthorizedRequest_Headers `protobuf:"bytes,7,rep,name=request_headers,json=requestHeaders,proto3" json:"request_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *IsAuthorizedRequest) Reset()         { *m = IsAuthorizedRequest{} }
func (m *IsAuthorizedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedRequest) ProtoMessage()    {}
func (*IsAuthorizedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffbc3c71370bee9a, []int{0}
}

func (m *IsAuthorizedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedRequest.Unmarshal(m, b)
}
func (m *IsAuthorizedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedRequest.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedRequest.Merge(m, src)
}
func (m *IsAuthorizedRequest) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedRequest.Size(m)
}
func (m *IsAuthorizedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedRequest proto.InternalMessageInfo

func (m *IsAuthorizedRequest) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestMethod() string {
	if m != nil {
		return m.RequestMethod
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestUrl() string {
	if m != nil {
		return m.RequestUrl
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestHost() string {
	if m != nil {
		return m.RequestHost
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestRequestUri() string {
	if m != nil {
		return m.RequestRequestUri
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestRemoteAddr() string {
	if m != nil {
		return m.RequestRemoteAddr
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestHeaders() map[string]*IsAuthorizedRequest_Headers {
	if m != nil {
		return m.RequestHeaders
	}
	return nil
}

// headers represents key-value pairs in an HTTP header; map[string][]string
type IsAuthorizedRequest_Headers struct {
	Value                []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorizedRequest_Headers) Reset()         { *m = IsAuthorizedRequest_Headers{} }
func (m *IsAuthorizedRequest_Headers) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedRequest_Headers) ProtoMessage()    {}
func (*IsAuthorizedRequest_Headers) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffbc3c71370bee9a, []int{0, 0}
}

func (m *IsAuthorizedRequest_Headers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedRequest_Headers.Unmarshal(m, b)
}
func (m *IsAuthorizedRequest_Headers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedRequest_Headers.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedRequest_Headers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedRequest_Headers.Merge(m, src)
}
func (m *IsAuthorizedRequest_Headers) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedRequest_Headers.Size(m)
}
func (m *IsAuthorizedRequest_Headers) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedRequest_Headers.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedRequest_Headers proto.InternalMessageInfo

func (m *IsAuthorizedRequest_Headers) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type IsAuthorizedReply struct {
	Allow                bool     `protobuf:"varint,1,opt,name=allow,proto3" json:"allow,omitempty"`
	SessionExpired       bool     `protobuf:"varint,2,opt,name=session_expired,json=sessionExpired,proto3" json:"session_expired,omitempty"`
	DenyReasons          []string `protobuf:"bytes,3,rep,name=deny_reasons,json=denyReasons,proto3" json:"deny_reasons,omitempty"`
	SignedJwt            string   `protobuf:"bytes,4,opt,name=signed_jwt,json=signedJwt,proto3" json:"signed_jwt,omitempty"`
	User                 string   `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Groups               []string `protobuf:"bytes,7,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorizedReply) Reset()         { *m = IsAuthorizedReply{} }
func (m *IsAuthorizedReply) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedReply) ProtoMessage()    {}
func (*IsAuthorizedReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffbc3c71370bee9a, []int{1}
}

func (m *IsAuthorizedReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedReply.Unmarshal(m, b)
}
func (m *IsAuthorizedReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedReply.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedReply.Merge(m, src)
}
func (m *IsAuthorizedReply) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedReply.Size(m)
}
func (m *IsAuthorizedReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedReply.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedReply proto.InternalMessageInfo

func (m *IsAuthorizedReply) GetAllow() bool {
	if m != nil {
		return m.Allow
	}
	return false
}

func (m *IsAuthorizedReply) GetSessionExpired() bool {
	if m != nil {
		return m.SessionExpired
	}
	return false
}

func (m *IsAuthorizedReply) GetDenyReasons() []string {
	if m != nil {
		return m.DenyReasons
	}
	return nil
}

func (m *IsAuthorizedReply) GetSignedJwt() string {
	if m != nil {
		return m.SignedJwt
	}
	return ""
}

func (m *IsAuthorizedReply) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *IsAuthorizedReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *IsAuthorizedReply) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*IsAuthorizedRequest)(nil), "authorize.IsAuthorizedRequest")
	proto.RegisterMapType((map[string]*IsAuthorizedRequest_Headers)(nil), "authorize.IsAuthorizedRequest.RequestHeadersEntry")
	proto.RegisterType((*IsAuthorizedRequest_Headers)(nil), "authorize.IsAuthorizedRequest.Headers")
	proto.RegisterType((*IsAuthorizedReply)(nil), "authorize.IsAuthorizedReply")
}

func init() {
	proto.RegisterFile("authorize.proto", fileDescriptor_ffbc3c71370bee9a)
}

var fileDescriptor_ffbc3c71370bee9a = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xd9, 0x6e, 0x9b, 0x76, 0x27, 0xa5, 0xa1, 0x13, 0x84, 0xac, 0x08, 0x68, 0x88, 0x04,
	0xe4, 0x94, 0x43, 0xb8, 0x20, 0xc4, 0xa5, 0x87, 0x4a, 0x05, 0x09, 0x0e, 0x16, 0x9c, 0x40, 0x5a,
	0x2d, 0xf2, 0xa8, 0x59, 0xea, 0xac, 0x17, 0xdb, 0x4b, 0x58, 0x1e, 0x94, 0x67, 0xe0, 0x31, 0x90,
	0xff, 0xa5, 0x14, 0x15, 0x38, 0xad, 0xe7, 0xe7, 0xcf, 0xe3, 0x99, 0xf9, 0xbc, 0x30, 0xaa, 0x3a,
	0xbb, 0x52, 0xba, 0xfe, 0x4e, 0x8b, 0x56, 0x2b, 0xab, 0xb0, 0xd8, 0x82, 0xd9, 0xcf, 0x1c, 0xc6,
	0xaf, 0xcc, 0x69, 0x8a, 0x05, 0xa7, 0x2f, 0x1d, 0x19, 0x8b, 0x0f, 0x00, 0x3a, 0x43, 0xba, 0xb4,
	0xea, 0x92, 0x1a, 0x96, 0x4d, 0xb3, 0x79, 0xc1, 0x0b, 0x47, 0xde, 0x39, 0x80, 0x8f, 0xe1, 0x48,
	0x07, 0x65, 0xb9, 0x26, 0xbb, 0x52, 0x82, 0xed, 0x78, 0xc9, 0xed, 0x48, 0xdf, 0x78, 0x88, 0x27,
	0x30, 0x4c, 0xb2, 0x4e, 0x4b, 0x96, 0x7b, 0x0d, 0x44, 0xf4, 0x5e, 0x4b, 0x7c, 0x04, 0x87, 0x49,
	0xb0, 0x52, 0xc6, 0xb2, 0x5d, 0xaf, 0x48, 0x87, 0xce, 0x95, 0xb1, 0xb8, 0x80, 0x71, 0x92, 0x5c,
	0xe5, 0xaa, 0xd9, 0x9e, 0x57, 0x1e, 0x47, 0xc4, 0x53, 0xca, 0xfa, 0xba, 0x7e, 0xad, 0x2c, 0x95,
	0x95, 0x10, 0x9a, 0x0d, 0xfe, 0xd0, 0xbb, 0x9d, 0x53, 0x21, 0x34, 0x7e, 0x80, 0xd1, 0xb6, 0x04,
	0xaa, 0x04, 0x69, 0xc3, 0xf6, 0xa7, 0xf9, 0x7c, 0xb8, 0x5c, 0x2e, 0xae, 0xe6, 0x76, 0xc3, 0x88,
	0x16, 0xf1, 0x7b, 0x1e, 0x0e, 0x9d, 0x35, 0x56, 0xf7, 0x3c, 0x4d, 0x25, 0xc2, 0xc9, 0x09, 0xec,
	0xc7, 0x25, 0xde, 0x85, 0xbd, 0xaf, 0x95, 0xec, 0x88, 0x65, 0xd3, 0x7c, 0x5e, 0xf0, 0x10, 0x4c,
	0x6a, 0x18, 0xdf, 0x90, 0x07, 0xef, 0x40, 0x7e, 0x49, 0x7d, 0x9c, 0xbb, 0x5b, 0xe2, 0xcb, 0x74,
	0xdc, 0x0d, 0x7a, 0xb8, 0x7c, 0xf2, 0x9f, 0xe2, 0x62, 0xb6, 0x78, 0xcd, 0x8b, 0x9d, 0xe7, 0xd9,
	0xec, 0x47, 0x06, 0xc7, 0xd7, 0xa5, 0xad, 0xec, 0x5d, 0x59, 0x95, 0x94, 0x6a, 0xe3, 0xef, 0x3a,
	0xe0, 0x21, 0xc0, 0xa7, 0x30, 0x32, 0x64, 0x4c, 0xad, 0x9a, 0x92, 0xbe, 0xb5, 0xb5, 0xa6, 0x60,
	0xf0, 0x01, 0x3f, 0x8a, 0xf8, 0x2c, 0x50, 0x67, 0xa0, 0xa0, 0xa6, 0x2f, 0x35, 0x55, 0x46, 0x35,
	0x86, 0xe5, 0xbe, 0xb9, 0xa1, 0x63, 0x3c, 0x20, 0xf7, 0x94, 0x4c, 0x7d, 0xd1, 0x90, 0x28, 0x3f,
	0x6f, 0x92, 0xc3, 0x45, 0x20, 0xaf, 0x37, 0x16, 0x11, 0x76, 0xdd, 0xbb, 0x8a, 0x86, 0xfa, 0xb5,
	0x2b, 0x8a, 0xd6, 0x55, 0x2d, 0xa3, 0x6b, 0x21, 0xc0, 0x7b, 0x30, 0xb8, 0xd0, 0xaa, 0x6b, 0x83,
	0x41, 0x05, 0x8f, 0xd1, 0xf2, 0x23, 0xc0, 0xb6, 0x2b, 0x8d, 0x6f, 0xe1, 0xf0, 0xf7, 0x2e, 0xf1,
	0xe1, 0xbf, 0x27, 0x35, 0xb9, 0xff, 0xd7, 0xfd, 0x56, 0xf6, 0xb3, 0x5b, 0x9f, 0x06, 0xfe, 0x9f,
	0x79, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x10, 0x59, 0xee, 0x46, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorizerClient is the client API for Authorizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizerClient interface {
	IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedReply, error)
}

type authorizerClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizerClient(cc grpc.ClientConnInterface) AuthorizerClient {
	return &authorizerClient{cc}
}

func (c *authorizerClient) IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedReply, error) {
	out := new(IsAuthorizedReply)
	err := c.cc.Invoke(ctx, "/authorize.Authorizer/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizerServer is the server API for Authorizer service.
type AuthorizerServer interface {
	IsAuthorized(context.Context, *IsAuthorizedRequest) (*IsAuthorizedReply, error)
}

// UnimplementedAuthorizerServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizerServer struct {
}

func (*UnimplementedAuthorizerServer) IsAuthorized(ctx context.Context, req *IsAuthorizedRequest) (*IsAuthorizedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthorized not implemented")
}

func RegisterAuthorizerServer(s *grpc.Server, srv AuthorizerServer) {
	s.RegisterService(&_Authorizer_serviceDesc, srv)
}

func _Authorizer_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizerServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorize.Authorizer/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizerServer).IsAuthorized(ctx, req.(*IsAuthorizedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authorize.Authorizer",
	HandlerType: (*AuthorizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthorized",
			Handler:    _Authorizer_IsAuthorized_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorize.proto",
}
