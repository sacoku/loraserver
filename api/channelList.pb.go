// Code generated by protoc-gen-go.
// source: channelList.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateChannelListRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CreateChannelListRequest) Reset()                    { *m = CreateChannelListRequest{} }
func (m *CreateChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateChannelListRequest) ProtoMessage()               {}
func (*CreateChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type CreateChannelListResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateChannelListResponse) Reset()                    { *m = CreateChannelListResponse{} }
func (m *CreateChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateChannelListResponse) ProtoMessage()               {}
func (*CreateChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type UpdateChannelListRequest struct {
	Id   int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *UpdateChannelListRequest) Reset()                    { *m = UpdateChannelListRequest{} }
func (m *UpdateChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateChannelListRequest) ProtoMessage()               {}
func (*UpdateChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type UpdateChannelListResponse struct {
}

func (m *UpdateChannelListResponse) Reset()                    { *m = UpdateChannelListResponse{} }
func (m *UpdateChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateChannelListResponse) ProtoMessage()               {}
func (*UpdateChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

type GetChannelListRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetChannelListRequest) Reset()                    { *m = GetChannelListRequest{} }
func (m *GetChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*GetChannelListRequest) ProtoMessage()               {}
func (*GetChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

type GetChannelListResponse struct {
	Id   int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *GetChannelListResponse) Reset()                    { *m = GetChannelListResponse{} }
func (m *GetChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*GetChannelListResponse) ProtoMessage()               {}
func (*GetChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

type ListChannelListRequest struct {
	Limit  int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListChannelListRequest) Reset()                    { *m = ListChannelListRequest{} }
func (m *ListChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListChannelListRequest) ProtoMessage()               {}
func (*ListChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

type ListChannelListResponse struct {
	TotalCount int64                     `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	Result     []*GetChannelListResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListChannelListResponse) Reset()                    { *m = ListChannelListResponse{} }
func (m *ListChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListChannelListResponse) ProtoMessage()               {}
func (*ListChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *ListChannelListResponse) GetResult() []*GetChannelListResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteChannelListRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteChannelListRequest) Reset()                    { *m = DeleteChannelListRequest{} }
func (m *DeleteChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteChannelListRequest) ProtoMessage()               {}
func (*DeleteChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

type DeleteChannelListResponse struct {
}

func (m *DeleteChannelListResponse) Reset()                    { *m = DeleteChannelListResponse{} }
func (m *DeleteChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteChannelListResponse) ProtoMessage()               {}
func (*DeleteChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func init() {
	proto.RegisterType((*CreateChannelListRequest)(nil), "api.CreateChannelListRequest")
	proto.RegisterType((*CreateChannelListResponse)(nil), "api.CreateChannelListResponse")
	proto.RegisterType((*UpdateChannelListRequest)(nil), "api.UpdateChannelListRequest")
	proto.RegisterType((*UpdateChannelListResponse)(nil), "api.UpdateChannelListResponse")
	proto.RegisterType((*GetChannelListRequest)(nil), "api.GetChannelListRequest")
	proto.RegisterType((*GetChannelListResponse)(nil), "api.GetChannelListResponse")
	proto.RegisterType((*ListChannelListRequest)(nil), "api.ListChannelListRequest")
	proto.RegisterType((*ListChannelListResponse)(nil), "api.ListChannelListResponse")
	proto.RegisterType((*DeleteChannelListRequest)(nil), "api.DeleteChannelListRequest")
	proto.RegisterType((*DeleteChannelListResponse)(nil), "api.DeleteChannelListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ChannelList service

type ChannelListClient interface {
	// Create creates the given channel-list.
	Create(ctx context.Context, in *CreateChannelListRequest, opts ...grpc.CallOption) (*CreateChannelListResponse, error)
	// Update updates the given channel-list.
	Update(ctx context.Context, in *UpdateChannelListRequest, opts ...grpc.CallOption) (*UpdateChannelListResponse, error)
	// Get returns the channel-list matching the given id.
	Get(ctx context.Context, in *GetChannelListRequest, opts ...grpc.CallOption) (*GetChannelListResponse, error)
	// List lists the channel-lists given an offset and limit.
	List(ctx context.Context, in *ListChannelListRequest, opts ...grpc.CallOption) (*ListChannelListResponse, error)
	// Delete deletes the channel-list matching the given id.
	Delete(ctx context.Context, in *DeleteChannelListRequest, opts ...grpc.CallOption) (*DeleteChannelListResponse, error)
}

type channelListClient struct {
	cc *grpc.ClientConn
}

func NewChannelListClient(cc *grpc.ClientConn) ChannelListClient {
	return &channelListClient{cc}
}

func (c *channelListClient) Create(ctx context.Context, in *CreateChannelListRequest, opts ...grpc.CallOption) (*CreateChannelListResponse, error) {
	out := new(CreateChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelListClient) Update(ctx context.Context, in *UpdateChannelListRequest, opts ...grpc.CallOption) (*UpdateChannelListResponse, error) {
	out := new(UpdateChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelListClient) Get(ctx context.Context, in *GetChannelListRequest, opts ...grpc.CallOption) (*GetChannelListResponse, error) {
	out := new(GetChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelListClient) List(ctx context.Context, in *ListChannelListRequest, opts ...grpc.CallOption) (*ListChannelListResponse, error) {
	out := new(ListChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelListClient) Delete(ctx context.Context, in *DeleteChannelListRequest, opts ...grpc.CallOption) (*DeleteChannelListResponse, error) {
	out := new(DeleteChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChannelList service

type ChannelListServer interface {
	// Create creates the given channel-list.
	Create(context.Context, *CreateChannelListRequest) (*CreateChannelListResponse, error)
	// Update updates the given channel-list.
	Update(context.Context, *UpdateChannelListRequest) (*UpdateChannelListResponse, error)
	// Get returns the channel-list matching the given id.
	Get(context.Context, *GetChannelListRequest) (*GetChannelListResponse, error)
	// List lists the channel-lists given an offset and limit.
	List(context.Context, *ListChannelListRequest) (*ListChannelListResponse, error)
	// Delete deletes the channel-list matching the given id.
	Delete(context.Context, *DeleteChannelListRequest) (*DeleteChannelListResponse, error)
}

func RegisterChannelListServer(s *grpc.Server, srv ChannelListServer) {
	s.RegisterService(&_ChannelList_serviceDesc, srv)
}

func _ChannelList_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).Create(ctx, req.(*CreateChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelList_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).Update(ctx, req.(*UpdateChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelList_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).Get(ctx, req.(*GetChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelList_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).List(ctx, req.(*ListChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelList_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).Delete(ctx, req.(*DeleteChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChannelList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ChannelList",
	HandlerType: (*ChannelListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChannelList_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ChannelList_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ChannelList_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ChannelList_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ChannelList_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor3,
}

func init() { proto.RegisterFile("channelList.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xda, 0x30,
	0x18, 0xc5, 0x45, 0xc2, 0x22, 0xed, 0x43, 0x9a, 0x34, 0x6f, 0x63, 0x21, 0xfc, 0x11, 0xf2, 0xa6,
	0x0d, 0xb1, 0x29, 0xd1, 0xe0, 0x6e, 0x9a, 0x76, 0xc3, 0x34, 0x6e, 0x7a, 0x15, 0xa9, 0xf7, 0x75,
	0x1b, 0x43, 0x2d, 0x05, 0x3b, 0x25, 0xa6, 0x37, 0x88, 0x9b, 0xbe, 0x42, 0xdf, 0xa0, 0xaf, 0xd4,
	0x57, 0xe8, 0x83, 0x54, 0xb1, 0x4d, 0x1b, 0x51, 0xbb, 0xed, 0x1d, 0xb1, 0x8f, 0xcf, 0xef, 0xf3,
	0xf1, 0x11, 0xf0, 0xfe, 0xec, 0x9c, 0x70, 0x4e, 0xf3, 0x23, 0x56, 0xca, 0xb8, 0x58, 0x0b, 0x29,
	0x90, 0x4f, 0x0a, 0x16, 0xf5, 0x96, 0x42, 0x2c, 0x73, 0x9a, 0x90, 0x82, 0x25, 0x84, 0x73, 0x21,
	0x89, 0x64, 0x82, 0x97, 0x5a, 0x82, 0x63, 0x08, 0x67, 0x6b, 0x4a, 0x24, 0x9d, 0x3d, 0x9e, 0x4e,
	0xe9, 0xc5, 0x86, 0x96, 0x12, 0x21, 0x68, 0x72, 0xb2, 0xa2, 0x61, 0x63, 0xd8, 0x18, 0xbd, 0x4d,
	0xd5, 0x6f, 0xfc, 0x03, 0x3a, 0x16, 0x7d, 0x59, 0x08, 0x5e, 0x52, 0xf4, 0x0e, 0x3c, 0x96, 0x29,
	0xb9, 0x9f, 0x7a, 0x2c, 0xc3, 0x7f, 0x21, 0x3c, 0x2e, 0x32, 0xbb, 0xf9, 0x81, 0xf6, 0x01, 0xe6,
	0xd5, 0x60, 0x5d, 0xe8, 0x58, 0xce, 0x6b, 0x18, 0xfe, 0x0e, 0x9f, 0xe6, 0x54, 0xbe, 0xec, 0x8c,
	0xff, 0x40, 0xfb, 0x50, 0x68, 0x9f, 0xd7, 0x3a, 0xc3, 0x7f, 0x68, 0x57, 0x67, 0x2c, 0x9c, 0x8f,
	0xf0, 0x26, 0x67, 0x2b, 0x26, 0x8d, 0x81, 0xfe, 0x40, 0x6d, 0x08, 0xc4, 0x62, 0x51, 0x52, 0xa9,
	0x5c, 0xfc, 0xd4, 0x7c, 0x61, 0x0e, 0x9f, 0x9f, 0xf8, 0x98, 0x31, 0x06, 0x00, 0x52, 0x48, 0x92,
	0xcf, 0xc4, 0x86, 0xef, 0xdd, 0x6a, 0x2b, 0x68, 0x0a, 0xc1, 0x9a, 0x96, 0x9b, 0xbc, 0xb2, 0xf4,
	0x47, 0xad, 0x49, 0x37, 0x26, 0x05, 0x8b, 0xed, 0x77, 0x4a, 0x8d, 0x14, 0x8f, 0x21, 0xfc, 0x47,
	0x73, 0xfa, 0x9a, 0xec, 0xab, 0x9c, 0x2d, 0x5a, 0x6d, 0x38, 0xb9, 0x69, 0x42, 0xab, 0xb6, 0x8e,
	0x96, 0x10, 0xe8, 0x06, 0xa0, 0xbe, 0x9a, 0xc3, 0x55, 0x9f, 0x68, 0xe0, 0xda, 0x36, 0x0f, 0x38,
	0xb8, 0xba, 0xbd, 0xbb, 0xf6, 0x42, 0xfc, 0x41, 0x55, 0xf3, 0xf2, 0x57, 0x52, 0x2b, 0xf0, 0xef,
	0xc6, 0x18, 0xe5, 0x10, 0xe8, 0xd7, 0x37, 0x20, 0x57, 0x95, 0x0c, 0xc8, 0xdd, 0x94, 0x2f, 0x0a,
	0xd4, 0x8f, 0x42, 0x0b, 0x28, 0xd9, 0xb2, 0x6c, 0x57, 0xd1, 0x4e, 0xc0, 0x9f, 0x53, 0x89, 0x22,
	0x6b, 0xb6, 0x9a, 0xf3, 0x5c, 0xee, 0x78, 0xa8, 0x20, 0x11, 0x72, 0x42, 0x90, 0x80, 0xa6, 0x0a,
	0x50, 0xdb, 0xd8, 0x4b, 0x15, 0xf5, 0xec, 0x9b, 0x06, 0xf2, 0x53, 0x41, 0xbe, 0xa1, 0xaf, 0x56,
	0x88, 0x6e, 0xda, 0x2e, 0xd9, 0xaa, 0x26, 0xee, 0x10, 0x83, 0x40, 0x3f, 0xab, 0x09, 0xd0, 0xd5,
	0x07, 0x13, 0xa0, 0xb3, 0x02, 0xfb, 0xbb, 0x8d, 0x9d, 0x77, 0x3b, 0x0d, 0xd4, 0xbf, 0xc9, 0xf4,
	0x3e, 0x00, 0x00, 0xff, 0xff, 0xff, 0x87, 0x8f, 0x4e, 0x85, 0x04, 0x00, 0x00,
}