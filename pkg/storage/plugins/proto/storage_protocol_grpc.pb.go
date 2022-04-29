// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pkg/storage/plugins/proto/storage_protocol.proto

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageProtocolClient is the client API for StorageProtocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageProtocolClient interface {
	EnsureIndex(ctx context.Context, in *EnsureIndexRequest, opts ...grpc.CallOption) (*EnsureIndexResponse, error)
	Aggregate(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*AggregateResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*PatchResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type storageProtocolClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageProtocolClient(cc grpc.ClientConnInterface) StorageProtocolClient {
	return &storageProtocolClient{cc}
}

func (c *storageProtocolClient) EnsureIndex(ctx context.Context, in *EnsureIndexRequest, opts ...grpc.CallOption) (*EnsureIndexResponse, error) {
	out := new(EnsureIndexResponse)
	err := c.cc.Invoke(ctx, "/plugins.StorageProtocol/EnsureIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProtocolClient) Aggregate(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*AggregateResponse, error) {
	out := new(AggregateResponse)
	err := c.cc.Invoke(ctx, "/plugins.StorageProtocol/Aggregate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProtocolClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/plugins.StorageProtocol/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProtocolClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/plugins.StorageProtocol/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProtocolClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/plugins.StorageProtocol/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProtocolClient) Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*PatchResponse, error) {
	out := new(PatchResponse)
	err := c.cc.Invoke(ctx, "/plugins.StorageProtocol/Patch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProtocolClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/plugins.StorageProtocol/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageProtocolClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/plugins.StorageProtocol/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageProtocolServer is the server API for StorageProtocol service.
// All implementations must embed UnimplementedStorageProtocolServer
// for forward compatibility
type StorageProtocolServer interface {
	EnsureIndex(context.Context, *EnsureIndexRequest) (*EnsureIndexResponse, error)
	Aggregate(context.Context, *AggregateRequest) (*AggregateResponse, error)
	Count(context.Context, *CountRequest) (*CountResponse, error)
	Find(context.Context, *FindRequest) (*FindResponse, error)
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)
	Patch(context.Context, *PatchRequest) (*PatchResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	mustEmbedUnimplementedStorageProtocolServer()
}

// UnimplementedStorageProtocolServer must be embedded to have forward compatible implementations.
type UnimplementedStorageProtocolServer struct {
}

func (UnimplementedStorageProtocolServer) EnsureIndex(context.Context, *EnsureIndexRequest) (*EnsureIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnsureIndex not implemented")
}
func (UnimplementedStorageProtocolServer) Aggregate(context.Context, *AggregateRequest) (*AggregateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Aggregate not implemented")
}
func (UnimplementedStorageProtocolServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedStorageProtocolServer) Find(context.Context, *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedStorageProtocolServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedStorageProtocolServer) Patch(context.Context, *PatchRequest) (*PatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedStorageProtocolServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedStorageProtocolServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStorageProtocolServer) mustEmbedUnimplementedStorageProtocolServer() {}

// UnsafeStorageProtocolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageProtocolServer will
// result in compilation errors.
type UnsafeStorageProtocolServer interface {
	mustEmbedUnimplementedStorageProtocolServer()
}

func RegisterStorageProtocolServer(s grpc.ServiceRegistrar, srv StorageProtocolServer) {
	s.RegisterService(&StorageProtocol_ServiceDesc, srv)
}

func _StorageProtocol_EnsureIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnsureIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProtocolServer).EnsureIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.StorageProtocol/EnsureIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProtocolServer).EnsureIndex(ctx, req.(*EnsureIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProtocol_Aggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProtocolServer).Aggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.StorageProtocol/Aggregate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProtocolServer).Aggregate(ctx, req.(*AggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProtocol_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProtocolServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.StorageProtocol/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProtocolServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProtocol_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProtocolServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.StorageProtocol/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProtocolServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProtocol_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProtocolServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.StorageProtocol/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProtocolServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProtocol_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProtocolServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.StorageProtocol/Patch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProtocolServer).Patch(ctx, req.(*PatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProtocol_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProtocolServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.StorageProtocol/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProtocolServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageProtocol_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageProtocolServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.StorageProtocol/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageProtocolServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageProtocol_ServiceDesc is the grpc.ServiceDesc for StorageProtocol service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageProtocol_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.StorageProtocol",
	HandlerType: (*StorageProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnsureIndex",
			Handler:    _StorageProtocol_EnsureIndex_Handler,
		},
		{
			MethodName: "Aggregate",
			Handler:    _StorageProtocol_Aggregate_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _StorageProtocol_Count_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _StorageProtocol_Find_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _StorageProtocol_Insert_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _StorageProtocol_Patch_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _StorageProtocol_Remove_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StorageProtocol_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/storage/plugins/proto/storage_protocol.proto",
}
