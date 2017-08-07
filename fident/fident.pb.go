// Code generated by protoc-gen-go.
// source: fident.proto
// DO NOT EDIT!

/*
Package fident is a generated protocol buffer package.

It is generated from these files:
	fident.proto

It has these top-level messages:
	ServiceRegistrationResponse
	ServiceRegistrationRequest
	AuthResponse
	AuthChallengeResponse
	AuthChallengeRequest
	PerformAuthRequest
	AuthKeyRequest
	AuthKeyResponse
	LoginTimestampResponse
	LoginTimestampRequest
	AccountAttribute
	AccountDetail
	AccountDetailResponse
	AccountDetailRequest
*/
package fident

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

type ServiceRegistrationResponse struct {
	IdentityId string `protobuf:"bytes,1,opt,name=identity_id,json=identityId" json:"identity_id,omitempty"`
	ServiceKey string `protobuf:"bytes,2,opt,name=service_key,json=serviceKey" json:"service_key,omitempty"`
}

func (m *ServiceRegistrationResponse) Reset()                    { *m = ServiceRegistrationResponse{} }
func (m *ServiceRegistrationResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceRegistrationResponse) ProtoMessage()               {}
func (*ServiceRegistrationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceRegistrationResponse) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *ServiceRegistrationResponse) GetServiceKey() string {
	if m != nil {
		return m.ServiceKey
	}
	return ""
}

type ServiceRegistrationRequest struct {
	ServiceName   string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServiceVendor string `protobuf:"bytes,2,opt,name=service_vendor,json=serviceVendor" json:"service_vendor,omitempty"`
	SupportEmail  string `protobuf:"bytes,3,opt,name=support_email,json=supportEmail" json:"support_email,omitempty"`
}

func (m *ServiceRegistrationRequest) Reset()                    { *m = ServiceRegistrationRequest{} }
func (m *ServiceRegistrationRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceRegistrationRequest) ProtoMessage()               {}
func (*ServiceRegistrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceRegistrationRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceRegistrationRequest) GetServiceVendor() string {
	if m != nil {
		return m.ServiceVendor
	}
	return ""
}

func (m *ServiceRegistrationRequest) GetSupportEmail() string {
	if m != nil {
		return m.SupportEmail
	}
	return ""
}

type AuthResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthChallengeResponse struct {
	Challenge string `protobuf:"bytes,1,opt,name=challenge" json:"challenge,omitempty"`
}

func (m *AuthChallengeResponse) Reset()                    { *m = AuthChallengeResponse{} }
func (m *AuthChallengeResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthChallengeResponse) ProtoMessage()               {}
func (*AuthChallengeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AuthChallengeResponse) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

type AuthChallengeRequest struct {
	IdentityId string `protobuf:"bytes,1,opt,name=identity_id,json=identityId" json:"identity_id,omitempty"`
	ProjectId  string `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *AuthChallengeRequest) Reset()                    { *m = AuthChallengeRequest{} }
func (m *AuthChallengeRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthChallengeRequest) ProtoMessage()               {}
func (*AuthChallengeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AuthChallengeRequest) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *AuthChallengeRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

type PerformAuthRequest struct {
	IdentityId        string `protobuf:"bytes,1,opt,name=identity_id,json=identityId" json:"identity_id,omitempty"`
	KeyHandle         string `protobuf:"bytes,2,opt,name=key_handle,json=keyHandle" json:"key_handle,omitempty"`
	ProjectId         string `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	ChallengeResponse string `protobuf:"bytes,4,opt,name=challenge_response,json=challengeResponse" json:"challenge_response,omitempty"`
}

func (m *PerformAuthRequest) Reset()                    { *m = PerformAuthRequest{} }
func (m *PerformAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*PerformAuthRequest) ProtoMessage()               {}
func (*PerformAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PerformAuthRequest) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *PerformAuthRequest) GetKeyHandle() string {
	if m != nil {
		return m.KeyHandle
	}
	return ""
}

func (m *PerformAuthRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *PerformAuthRequest) GetChallengeResponse() string {
	if m != nil {
		return m.ChallengeResponse
	}
	return ""
}

type AuthKeyRequest struct {
	Username         string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password         string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	UsageDescription string `protobuf:"bytes,3,opt,name=usage_description,json=usageDescription" json:"usage_description,omitempty"`
}

func (m *AuthKeyRequest) Reset()                    { *m = AuthKeyRequest{} }
func (m *AuthKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthKeyRequest) ProtoMessage()               {}
func (*AuthKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AuthKeyRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthKeyRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthKeyRequest) GetUsageDescription() string {
	if m != nil {
		return m.UsageDescription
	}
	return ""
}

type AuthKeyResponse struct {
	Username   string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	IdentityId string `protobuf:"bytes,2,opt,name=identity_id,json=identityId" json:"identity_id,omitempty"`
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	KeyHandle  string `protobuf:"bytes,4,opt,name=key_handle,json=keyHandle" json:"key_handle,omitempty"`
}

func (m *AuthKeyResponse) Reset()                    { *m = AuthKeyResponse{} }
func (m *AuthKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthKeyResponse) ProtoMessage()               {}
func (*AuthKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AuthKeyResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthKeyResponse) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *AuthKeyResponse) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *AuthKeyResponse) GetKeyHandle() string {
	if m != nil {
		return m.KeyHandle
	}
	return ""
}

type LoginTimestampResponse struct {
	Results map[string]int64 `protobuf:"bytes,1,rep,name=results" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *LoginTimestampResponse) Reset()                    { *m = LoginTimestampResponse{} }
func (m *LoginTimestampResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginTimestampResponse) ProtoMessage()               {}
func (*LoginTimestampResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LoginTimestampResponse) GetResults() map[string]int64 {
	if m != nil {
		return m.Results
	}
	return nil
}

type LoginTimestampRequest struct {
	IdentityId []string `protobuf:"bytes,1,rep,name=identity_id,json=identityId" json:"identity_id,omitempty"`
}

func (m *LoginTimestampRequest) Reset()                    { *m = LoginTimestampRequest{} }
func (m *LoginTimestampRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginTimestampRequest) ProtoMessage()               {}
func (*LoginTimestampRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *LoginTimestampRequest) GetIdentityId() []string {
	if m != nil {
		return m.IdentityId
	}
	return nil
}

type AccountAttribute struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *AccountAttribute) Reset()                    { *m = AccountAttribute{} }
func (m *AccountAttribute) String() string            { return proto.CompactTextString(m) }
func (*AccountAttribute) ProtoMessage()               {}
func (*AccountAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AccountAttribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AccountAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AccountDetail struct {
	IdentityId string              `protobuf:"bytes,1,opt,name=identity_id,json=identityId" json:"identity_id,omitempty"`
	Username   string              `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Attributes []*AccountAttribute `protobuf:"bytes,3,rep,name=attributes" json:"attributes,omitempty"`
	Created    int64               `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
}

func (m *AccountDetail) Reset()                    { *m = AccountDetail{} }
func (m *AccountDetail) String() string            { return proto.CompactTextString(m) }
func (*AccountDetail) ProtoMessage()               {}
func (*AccountDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AccountDetail) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *AccountDetail) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountDetail) GetAttributes() []*AccountAttribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *AccountDetail) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type AccountDetailResponse struct {
	Results []*AccountDetail `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *AccountDetailResponse) Reset()                    { *m = AccountDetailResponse{} }
func (m *AccountDetailResponse) String() string            { return proto.CompactTextString(m) }
func (*AccountDetailResponse) ProtoMessage()               {}
func (*AccountDetailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AccountDetailResponse) GetResults() []*AccountDetail {
	if m != nil {
		return m.Results
	}
	return nil
}

type AccountDetailRequest struct {
	IdentityId []string `protobuf:"bytes,1,rep,name=identity_id,json=identityId" json:"identity_id,omitempty"`
}

func (m *AccountDetailRequest) Reset()                    { *m = AccountDetailRequest{} }
func (m *AccountDetailRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountDetailRequest) ProtoMessage()               {}
func (*AccountDetailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AccountDetailRequest) GetIdentityId() []string {
	if m != nil {
		return m.IdentityId
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceRegistrationResponse)(nil), "fident.ServiceRegistrationResponse")
	proto.RegisterType((*ServiceRegistrationRequest)(nil), "fident.ServiceRegistrationRequest")
	proto.RegisterType((*AuthResponse)(nil), "fident.AuthResponse")
	proto.RegisterType((*AuthChallengeResponse)(nil), "fident.AuthChallengeResponse")
	proto.RegisterType((*AuthChallengeRequest)(nil), "fident.AuthChallengeRequest")
	proto.RegisterType((*PerformAuthRequest)(nil), "fident.PerformAuthRequest")
	proto.RegisterType((*AuthKeyRequest)(nil), "fident.AuthKeyRequest")
	proto.RegisterType((*AuthKeyResponse)(nil), "fident.AuthKeyResponse")
	proto.RegisterType((*LoginTimestampResponse)(nil), "fident.LoginTimestampResponse")
	proto.RegisterType((*LoginTimestampRequest)(nil), "fident.LoginTimestampRequest")
	proto.RegisterType((*AccountAttribute)(nil), "fident.AccountAttribute")
	proto.RegisterType((*AccountDetail)(nil), "fident.AccountDetail")
	proto.RegisterType((*AccountDetailResponse)(nil), "fident.AccountDetailResponse")
	proto.RegisterType((*AccountDetailRequest)(nil), "fident.AccountDetailRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	// Pre-authentication methods
	GetAuthenticationChallenge(ctx context.Context, in *AuthChallengeRequest, opts ...grpc.CallOption) (*AuthChallengeResponse, error)
	PerformAuthentication(ctx context.Context, in *PerformAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// Account generation / management
	GetNewAuthenticationKey(ctx context.Context, in *AuthKeyRequest, opts ...grpc.CallOption) (*AuthKeyResponse, error)
	CreateServiceAccount(ctx context.Context, in *ServiceRegistrationRequest, opts ...grpc.CallOption) (*ServiceRegistrationResponse, error)
	// Identity queries
	GetLastLoginTimestamps(ctx context.Context, in *LoginTimestampRequest, opts ...grpc.CallOption) (*LoginTimestampResponse, error)
	GetAccountDetails(ctx context.Context, in *AccountDetailRequest, opts ...grpc.CallOption) (*AccountDetailResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetAuthenticationChallenge(ctx context.Context, in *AuthChallengeRequest, opts ...grpc.CallOption) (*AuthChallengeResponse, error) {
	out := new(AuthChallengeResponse)
	err := grpc.Invoke(ctx, "/fident.Auth/GetAuthenticationChallenge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) PerformAuthentication(ctx context.Context, in *PerformAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/fident.Auth/PerformAuthentication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetNewAuthenticationKey(ctx context.Context, in *AuthKeyRequest, opts ...grpc.CallOption) (*AuthKeyResponse, error) {
	out := new(AuthKeyResponse)
	err := grpc.Invoke(ctx, "/fident.Auth/GetNewAuthenticationKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateServiceAccount(ctx context.Context, in *ServiceRegistrationRequest, opts ...grpc.CallOption) (*ServiceRegistrationResponse, error) {
	out := new(ServiceRegistrationResponse)
	err := grpc.Invoke(ctx, "/fident.Auth/CreateServiceAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetLastLoginTimestamps(ctx context.Context, in *LoginTimestampRequest, opts ...grpc.CallOption) (*LoginTimestampResponse, error) {
	out := new(LoginTimestampResponse)
	err := grpc.Invoke(ctx, "/fident.Auth/GetLastLoginTimestamps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAccountDetails(ctx context.Context, in *AccountDetailRequest, opts ...grpc.CallOption) (*AccountDetailResponse, error) {
	out := new(AccountDetailResponse)
	err := grpc.Invoke(ctx, "/fident.Auth/GetAccountDetails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	// Pre-authentication methods
	GetAuthenticationChallenge(context.Context, *AuthChallengeRequest) (*AuthChallengeResponse, error)
	PerformAuthentication(context.Context, *PerformAuthRequest) (*AuthResponse, error)
	// Account generation / management
	GetNewAuthenticationKey(context.Context, *AuthKeyRequest) (*AuthKeyResponse, error)
	CreateServiceAccount(context.Context, *ServiceRegistrationRequest) (*ServiceRegistrationResponse, error)
	// Identity queries
	GetLastLoginTimestamps(context.Context, *LoginTimestampRequest) (*LoginTimestampResponse, error)
	GetAccountDetails(context.Context, *AccountDetailRequest) (*AccountDetailResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_GetAuthenticationChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAuthenticationChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fident.Auth/GetAuthenticationChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAuthenticationChallenge(ctx, req.(*AuthChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_PerformAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).PerformAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fident.Auth/PerformAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).PerformAuthentication(ctx, req.(*PerformAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetNewAuthenticationKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetNewAuthenticationKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fident.Auth/GetNewAuthenticationKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetNewAuthenticationKey(ctx, req.(*AuthKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fident.Auth/CreateServiceAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateServiceAccount(ctx, req.(*ServiceRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetLastLoginTimestamps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginTimestampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetLastLoginTimestamps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fident.Auth/GetLastLoginTimestamps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetLastLoginTimestamps(ctx, req.(*LoginTimestampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetAccountDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAccountDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fident.Auth/GetAccountDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAccountDetails(ctx, req.(*AccountDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fident.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthenticationChallenge",
			Handler:    _Auth_GetAuthenticationChallenge_Handler,
		},
		{
			MethodName: "PerformAuthentication",
			Handler:    _Auth_PerformAuthentication_Handler,
		},
		{
			MethodName: "GetNewAuthenticationKey",
			Handler:    _Auth_GetNewAuthenticationKey_Handler,
		},
		{
			MethodName: "CreateServiceAccount",
			Handler:    _Auth_CreateServiceAccount_Handler,
		},
		{
			MethodName: "GetLastLoginTimestamps",
			Handler:    _Auth_GetLastLoginTimestamps_Handler,
		},
		{
			MethodName: "GetAccountDetails",
			Handler:    _Auth_GetAccountDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fident.proto",
}

func init() { proto.RegisterFile("fident.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0x51, 0x4f, 0xdb, 0x48,
	0x10, 0xc6, 0x71, 0x0e, 0x2e, 0x43, 0xe0, 0x60, 0x95, 0x80, 0xe5, 0x83, 0x3b, 0x6e, 0xb9, 0x4a,
	0x48, 0xa8, 0x54, 0xa2, 0xaa, 0x8a, 0x78, 0x43, 0x80, 0x00, 0x05, 0xa1, 0xca, 0xad, 0x90, 0xaa,
	0x3e, 0x44, 0x8b, 0x3d, 0x04, 0x37, 0x89, 0xed, 0xee, 0xae, 0x83, 0xfc, 0x1b, 0x2a, 0x55, 0x7d,
	0xef, 0x4b, 0xd5, 0x5f, 0x5a, 0xd9, 0xde, 0x75, 0x1c, 0x93, 0x00, 0x6f, 0xde, 0x6f, 0xbf, 0x99,
	0x6f, 0x76, 0xbe, 0xd9, 0x35, 0x34, 0x6f, 0x7d, 0x0f, 0x03, 0xb9, 0x17, 0xf1, 0x50, 0x86, 0x64,
	0x3e, 0x5f, 0xd1, 0x2e, 0xfc, 0xfd, 0x1e, 0xf9, 0xc8, 0x77, 0xd1, 0xc1, 0x9e, 0x2f, 0x24, 0x67,
	0xd2, 0x0f, 0x03, 0x07, 0x45, 0x14, 0x06, 0x02, 0xc9, 0xbf, 0xb0, 0x98, 0xf1, 0x7c, 0x99, 0x74,
	0x7d, 0xcf, 0x32, 0xb6, 0x8c, 0x9d, 0x86, 0x03, 0x1a, 0xba, 0xf0, 0x52, 0x82, 0xc8, 0xe3, 0xbb,
	0x7d, 0x4c, 0xac, 0x5a, 0x4e, 0x50, 0x50, 0x07, 0x13, 0xfa, 0xd5, 0x00, 0x7b, 0xaa, 0xc2, 0x97,
	0x18, 0x85, 0x24, 0xff, 0x41, 0x53, 0xc7, 0x07, 0x6c, 0x88, 0x4a, 0x41, 0xe7, 0xbc, 0x62, 0x43,
	0x24, 0x2f, 0x60, 0x59, 0x53, 0x46, 0x18, 0x78, 0x21, 0x57, 0x2a, 0x4b, 0x0a, 0xbd, 0xce, 0x40,
	0xb2, 0x0d, 0x4b, 0x22, 0x8e, 0xa2, 0x90, 0xcb, 0x2e, 0x0e, 0x99, 0x3f, 0xb0, 0xcc, 0x8c, 0xd5,
	0x54, 0xe0, 0x69, 0x8a, 0xd1, 0xff, 0xa1, 0x79, 0x14, 0xcb, 0xbb, 0xe2, 0x7c, 0x2d, 0xf8, 0x43,
	0x86, 0x7d, 0x0c, 0x94, 0x6e, 0xbe, 0xa0, 0x6f, 0xa0, 0x9d, 0xb2, 0x8e, 0xef, 0xd8, 0x60, 0x80,
	0x41, 0x0f, 0x0b, 0xfa, 0x06, 0x34, 0x5c, 0x0d, 0xaa, 0x90, 0x31, 0x40, 0xaf, 0xa1, 0x55, 0x09,
	0xcb, 0xcf, 0xf8, 0x64, 0x13, 0x37, 0x01, 0x22, 0x1e, 0x7e, 0x46, 0x57, 0xa6, 0xfb, 0xf9, 0xe9,
	0x1a, 0x0a, 0xb9, 0xf0, 0xe8, 0x2f, 0x03, 0xc8, 0x3b, 0xe4, 0xb7, 0x21, 0x1f, 0xe6, 0xc5, 0x3f,
	0x3f, 0x6d, 0x1f, 0x93, 0xee, 0x1d, 0x0b, 0xbc, 0x01, 0xea, 0xb4, 0x7d, 0x4c, 0xce, 0x33, 0xa0,
	0xa2, 0x6a, 0x56, 0x54, 0xc9, 0x4b, 0x20, 0xc5, 0xd1, 0xba, 0x5c, 0x75, 0xc0, 0xaa, 0x67, 0xb4,
	0x55, 0xb7, 0xda, 0x1a, 0x1a, 0xc3, 0x72, 0x5a, 0x5c, 0x07, 0x13, 0x5d, 0x9f, 0x0d, 0x7f, 0xc6,
	0x02, 0x79, 0xc9, 0xd6, 0x62, 0x9d, 0xee, 0x45, 0x4c, 0x88, 0xfb, 0x90, 0xeb, 0xf3, 0x16, 0x6b,
	0xb2, 0x0b, 0xab, 0xb1, 0x60, 0x3d, 0xec, 0x7a, 0x28, 0x5c, 0xee, 0x47, 0xe9, 0xb8, 0xa8, 0xf2,
	0x56, 0xb2, 0x8d, 0x93, 0x31, 0x4e, 0xbf, 0x19, 0xf0, 0x57, 0xa1, 0xab, 0x5c, 0x7a, 0x4c, 0xb8,
	0xd2, 0xb4, 0xda, 0xb4, 0x81, 0x8e, 0xb8, 0x3f, 0x62, 0x32, 0x1f, 0xe8, 0x5c, 0x17, 0x14, 0xd4,
	0xc1, 0xa4, 0xd2, 0xd5, 0x7a, 0xa5, 0xab, 0xf4, 0x87, 0x01, 0x6b, 0x97, 0x61, 0xcf, 0x0f, 0x3e,
	0xf8, 0x43, 0x14, 0x92, 0x0d, 0xa3, 0xa2, 0xae, 0x53, 0x58, 0xe0, 0x28, 0xe2, 0x81, 0x14, 0x96,
	0xb1, 0x65, 0xee, 0x2c, 0xee, 0xef, 0xee, 0xa9, 0x3b, 0x39, 0x3d, 0x60, 0xcf, 0xc9, 0xd9, 0xa7,
	0x81, 0xe4, 0x89, 0xa3, 0x63, 0xed, 0x43, 0x68, 0x96, 0x37, 0xc8, 0x0a, 0x98, 0x69, 0xa5, 0xf9,
	0x49, 0xd3, 0xcf, 0x74, 0xaa, 0x47, 0x6c, 0x10, 0xe7, 0x9e, 0x9b, 0x4e, 0xbe, 0x38, 0xac, 0x1d,
	0x18, 0xf4, 0x00, 0xda, 0x55, 0xad, 0x19, 0xc3, 0x64, 0x4e, 0xf6, 0x85, 0x1e, 0xc2, 0xca, 0x91,
	0xeb, 0x86, 0x71, 0x20, 0x8f, 0xa4, 0xe4, 0xfe, 0x4d, 0x2c, 0xf1, 0x29, 0xe5, 0x86, 0x52, 0xa6,
	0x3f, 0x0d, 0x58, 0x52, 0xc1, 0x27, 0x28, 0x99, 0x3f, 0x78, 0x7a, 0x76, 0xcb, 0x1e, 0xd6, 0x2a,
	0x1e, 0x1e, 0x00, 0x30, 0x5d, 0x83, 0xb0, 0xcc, 0xac, 0x95, 0x96, 0x6e, 0x65, 0xb5, 0x48, 0xa7,
	0xc4, 0x25, 0x16, 0x2c, 0xb8, 0x1c, 0x99, 0x44, 0x2f, 0x33, 0xce, 0x74, 0xf4, 0x92, 0x9e, 0x43,
	0x7b, 0xa2, 0xc2, 0xc2, 0xb4, 0x57, 0x55, 0xd3, 0xda, 0x15, 0x25, 0xc5, 0xd7, 0x2c, 0xfa, 0x16,
	0x5a, 0x95, 0x4c, 0xcf, 0xeb, 0xf0, 0xfe, 0xf7, 0x3a, 0xd4, 0xd3, 0x51, 0x26, 0x9f, 0xc0, 0x3e,
	0x43, 0x99, 0x7e, 0xa6, 0x7b, 0x6e, 0xf6, 0x5e, 0x16, 0x8f, 0x0a, 0xd9, 0x28, 0xf4, 0xa7, 0xbc,
	0x35, 0xf6, 0xe6, 0x8c, 0x5d, 0x75, 0x4b, 0xe7, 0x48, 0x07, 0xda, 0xa5, 0xb7, 0x64, 0x2c, 0x40,
	0x6c, 0x1d, 0xf9, 0xf0, 0xa9, 0xb1, 0x5b, 0xe5, 0xac, 0xa5, 0x64, 0x97, 0xb0, 0x7e, 0x86, 0xf2,
	0x0a, 0xef, 0x27, 0x73, 0xa5, 0xd7, 0x64, 0xad, 0x1c, 0x32, 0x7e, 0x15, 0xec, 0xf5, 0x07, 0x78,
	0x91, 0x8d, 0x41, 0xeb, 0x38, 0xb3, 0x43, 0xfd, 0x2f, 0x54, 0x1b, 0x09, 0xd5, 0x21, 0xb3, 0xff,
	0x23, 0xf6, 0xf6, 0xa3, 0x9c, 0x42, 0xe2, 0x23, 0xac, 0x9d, 0xa1, 0xbc, 0x64, 0x42, 0x4e, 0x5e,
	0x03, 0x41, 0x36, 0x67, 0xdd, 0xc5, 0x3c, 0xff, 0x3f, 0x8f, 0x5f, 0x55, 0x3a, 0x47, 0x1c, 0x58,
	0x4d, 0x5d, 0x2b, 0x5b, 0x2f, 0x4a, 0x66, 0x4d, 0x19, 0x89, 0x92, 0x59, 0xd3, 0x46, 0x8f, 0xce,
	0xdd, 0xcc, 0x67, 0x3f, 0xeb, 0xd7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x51, 0x10, 0xc7, 0xc8,
	0xbc, 0x07, 0x00, 0x00,
}
