// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: profession.proto

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

// ProfessionServiceClient is the client API for ProfessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfessionServiceClient interface {
	Create(ctx context.Context, in *CreateProfessionRequest, opts ...grpc.CallOption) (*Profession, error)
	GetAll(ctx context.Context, in *GetAllProfessionRequest, opts ...grpc.CallOption) (*GetAllProfessionResponse, error)
	GetById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Profession, error)
	Update(ctx context.Context, in *Profession, opts ...grpc.CallOption) (*Profession, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error)
}

type professionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfessionServiceClient(cc grpc.ClientConnInterface) ProfessionServiceClient {
	return &professionServiceClient{cc}
}

func (c *professionServiceClient) Create(ctx context.Context, in *CreateProfessionRequest, opts ...grpc.CallOption) (*Profession, error) {
	out := new(Profession)
	err := c.cc.Invoke(ctx, "/position_service.ProfessionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionServiceClient) GetAll(ctx context.Context, in *GetAllProfessionRequest, opts ...grpc.CallOption) (*GetAllProfessionResponse, error) {
	out := new(GetAllProfessionResponse)
	err := c.cc.Invoke(ctx, "/position_service.ProfessionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionServiceClient) GetById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Profession, error) {
	out := new(Profession)
	err := c.cc.Invoke(ctx, "/position_service.ProfessionService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionServiceClient) Update(ctx context.Context, in *Profession, opts ...grpc.CallOption) (*Profession, error) {
	out := new(Profession)
	err := c.cc.Invoke(ctx, "/position_service.ProfessionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/position_service.ProfessionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfessionServiceServer is the server API for ProfessionService service.
// All implementations must embed UnimplementedProfessionServiceServer
// for forward compatibility
type ProfessionServiceServer interface {
	Create(context.Context, *CreateProfessionRequest) (*Profession, error)
	GetAll(context.Context, *GetAllProfessionRequest) (*GetAllProfessionResponse, error)
	GetById(context.Context, *Id) (*Profession, error)
	Update(context.Context, *Profession) (*Profession, error)
	Delete(context.Context, *Id) (*Id, error)
	mustEmbedUnimplementedProfessionServiceServer()
}

// UnimplementedProfessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfessionServiceServer struct {
}

func (UnimplementedProfessionServiceServer) Create(context.Context, *CreateProfessionRequest) (*Profession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProfessionServiceServer) GetAll(context.Context, *GetAllProfessionRequest) (*GetAllProfessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedProfessionServiceServer) GetById(context.Context, *Id) (*Profession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedProfessionServiceServer) Update(context.Context, *Profession) (*Profession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProfessionServiceServer) Delete(context.Context, *Id) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProfessionServiceServer) mustEmbedUnimplementedProfessionServiceServer() {}

// UnsafeProfessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfessionServiceServer will
// result in compilation errors.
type UnsafeProfessionServiceServer interface {
	mustEmbedUnimplementedProfessionServiceServer()
}

func RegisterProfessionServiceServer(s grpc.ServiceRegistrar, srv ProfessionServiceServer) {
	s.RegisterService(&ProfessionService_ServiceDesc, srv)
}

func _ProfessionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.ProfessionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).Create(ctx, req.(*CreateProfessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProfessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.ProfessionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).GetAll(ctx, req.(*GetAllProfessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.ProfessionService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).GetById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.ProfessionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).Update(ctx, req.(*Profession))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_service.ProfessionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfessionService_ServiceDesc is the grpc.ServiceDesc for ProfessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "position_service.ProfessionService",
	HandlerType: (*ProfessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProfessionService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ProfessionService_GetAll_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ProfessionService_GetById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProfessionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProfessionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profession.proto",
}
