// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: position.proto

package position_service

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

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	GetAllPosition(ctx context.Context, in *GetAllPositionRequest, opts ...grpc.CallOption) (*GetAllPositionResponse, error)
	GetByIdPosition(ctx context.Context, in *PositionId, opts ...grpc.CallOption) (*Position, error)
	CreatePosition(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*PositionId, error)
	UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*PositionId, error)
	DeletePosition(ctx context.Context, in *PositionId, opts ...grpc.CallOption) (*PositionId, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) GetAllPosition(ctx context.Context, in *GetAllPositionRequest, opts ...grpc.CallOption) (*GetAllPositionResponse, error) {
	out := new(GetAllPositionResponse)
	err := c.cc.Invoke(ctx, "/position_service.PositionService/GetAllPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetByIdPosition(ctx context.Context, in *PositionId, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, "/position_service.PositionService/GetByIdPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) CreatePosition(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*PositionId, error) {
	out := new(PositionId)
	err := c.cc.Invoke(ctx, "/position_service.PositionService/CreatePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*PositionId, error) {
	out := new(PositionId)
	err := c.cc.Invoke(ctx, "/position_service.PositionService/UpdatePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) DeletePosition(ctx context.Context, in *PositionId, opts ...grpc.CallOption) (*PositionId, error) {
	out := new(PositionId)
	err := c.cc.Invoke(ctx, "/position_service.PositionService/DeletePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations must embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	GetAllPosition(context.Context, *GetAllPositionRequest) (*GetAllPositionResponse, error)
	GetByIdPosition(context.Context, *PositionId) (*Position, error)
	CreatePosition(context.Context, *CreatePositionRequest) (*PositionId, error)
	UpdatePosition(context.Context, *UpdatePositionRequest) (*PositionId, error)
	DeletePosition(context.Context, *PositionId) (*PositionId, error)
	mustEmbedUnimplementedPositionServiceServer()
}

// UnimplementedPositionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) GetAllPosition(context.Context, *GetAllPositionRequest) (*GetAllPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPosition not implemented")
}
func (UnimplementedPositionServiceServer) GetByIdPosition(context.Context, *PositionId) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdPosition not implemented")
}
func (UnimplementedPositionServiceServer) CreatePosition(context.Context, *CreatePositionRequest) (*PositionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosition not implemented")
}
func (UnimplementedPositionServiceServer) UpdatePosition(context.Context, *UpdatePositionRequest) (*PositionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosition not implemented")
}
func (UnimplementedPositionServiceServer) DeletePosition(context.Context, *PositionId) (*PositionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosition not implemented")
}
func (UnimplementedPositionServiceServer) mustEmbedUnimplementedPositionServiceServer() {}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_GetAllPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetAllPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.PositionService/GetAllPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetAllPosition(ctx, req.(*GetAllPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_GetByIdPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetByIdPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.PositionService/GetByIdPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetByIdPosition(ctx, req.(*PositionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_CreatePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).CreatePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.PositionService/CreatePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).CreatePosition(ctx, req.(*CreatePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_UpdatePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).UpdatePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.PositionService/UpdatePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).UpdatePosition(ctx, req.(*UpdatePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_DeletePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).DeletePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.PositionService/DeletePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).DeletePosition(ctx, req.(*PositionId))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "position_service.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPosition",
			Handler:    _PositionService_GetAllPosition_Handler,
		},
		{
			MethodName: "GetByIdPosition",
			Handler:    _PositionService_GetByIdPosition_Handler,
		},
		{
			MethodName: "CreatePosition",
			Handler:    _PositionService_CreatePosition_Handler,
		},
		{
			MethodName: "UpdatePosition",
			Handler:    _PositionService_UpdatePosition_Handler,
		},
		{
			MethodName: "DeletePosition",
			Handler:    _PositionService_DeletePosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "position.proto",
}
