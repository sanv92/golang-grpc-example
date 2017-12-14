// Code generated by protoc-gen-go. DO NOT EDIT.
// source: posts/posts.proto

/*
Package posts is a generated protocol buffer package.

It is generated from these files:
	posts/posts.proto

It has these top-level messages:
	UserRequest
	UserResponse
*/
package posts

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type UserRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "posts.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "posts.UserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	PostTest(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Post(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) PostTest(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/posts.User/PostTest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Post(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/posts.User/Post", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	PostTest(context.Context, *UserRequest) (*UserResponse, error)
	Post(context.Context, *UserRequest) (*UserResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_PostTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).PostTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts.User/PostTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).PostTest(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts.User/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Post(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "posts.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostTest",
			Handler:    _User_PostTest_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _User_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posts/posts.proto",
}

func init() { proto.RegisterFile("posts/posts.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xc8, 0x2f, 0x2e,
	0x29, 0xd6, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x94, 0x4c,
	0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49,
	0x62, 0x49, 0x66, 0x7e, 0x1e, 0x54, 0x91, 0x92, 0x22, 0x17, 0x77, 0x68, 0x71, 0x6a, 0x51, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0xa4, 0xc1, 0xc5, 0x03, 0x51, 0x52, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0x2a, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x0e, 0x53, 0x06, 0xe3,
	0x1a, 0x75, 0x30, 0x72, 0xb1, 0x80, 0x94, 0x0a, 0x99, 0x72, 0x71, 0x04, 0xe4, 0x17, 0x97, 0x84,
	0x80, 0x8d, 0xd4, 0x83, 0x38, 0x0a, 0xc9, 0x1a, 0x29, 0x61, 0x14, 0x31, 0x88, 0xb9, 0x4a, 0x0c,
	0x42, 0x9e, 0x5c, 0x2c, 0x20, 0x6d, 0xc4, 0x6b, 0x91, 0x68, 0xba, 0xfc, 0x64, 0x32, 0x93, 0x90,
	0x90, 0x00, 0xd8, 0x77, 0xa5, 0xc5, 0xa9, 0x45, 0xfa, 0xd5, 0x20, 0x37, 0xd7, 0x26, 0xb1, 0x81,
	0xbd, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x32, 0xf1, 0xe6, 0xac, 0x18, 0x01, 0x00, 0x00,
}