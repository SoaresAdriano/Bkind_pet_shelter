// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: src/services/petCadaster_service/ports/grpc/protos/petCadaster.proto

package grpc

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

// PetCadasterClient is the client API for PetCadaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetCadasterClient interface {
	ListPets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (PetCadaster_ListPetsClient, error)
	GetPet(ctx context.Context, in *ID, opts ...grpc.CallOption) (*PetInfo, error)
	CreatePetCadaster(ctx context.Context, in *PetInfo, opts ...grpc.CallOption) (*ID, error)
	UpdatePetCadaster(ctx context.Context, in *PetInfo, opts ...grpc.CallOption) (*Status, error)
	RenamePet(ctx context.Context, in *RenameInfo, opts ...grpc.CallOption) (*Status, error)
	DeletePetCadaster(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error)
	CountPets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (PetCadaster_CountPetsClient, error)
}

type petCadasterClient struct {
	cc grpc.ClientConnInterface
}

func NewPetCadasterClient(cc grpc.ClientConnInterface) PetCadasterClient {
	return &petCadasterClient{cc}
}

func (c *petCadasterClient) ListPets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (PetCadaster_ListPetsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PetCadaster_ServiceDesc.Streams[0], "/grpc.petCadaster/ListPets", opts...)
	if err != nil {
		return nil, err
	}
	x := &petCadasterListPetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PetCadaster_ListPetsClient interface {
	Recv() (*PetInfo, error)
	grpc.ClientStream
}

type petCadasterListPetsClient struct {
	grpc.ClientStream
}

func (x *petCadasterListPetsClient) Recv() (*PetInfo, error) {
	m := new(PetInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *petCadasterClient) GetPet(ctx context.Context, in *ID, opts ...grpc.CallOption) (*PetInfo, error) {
	out := new(PetInfo)
	err := c.cc.Invoke(ctx, "/grpc.petCadaster/GetPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petCadasterClient) CreatePetCadaster(ctx context.Context, in *PetInfo, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/grpc.petCadaster/CreatePetCadaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petCadasterClient) UpdatePetCadaster(ctx context.Context, in *PetInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/grpc.petCadaster/UpdatePetCadaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petCadasterClient) RenamePet(ctx context.Context, in *RenameInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/grpc.petCadaster/RenamePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petCadasterClient) DeletePetCadaster(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/grpc.petCadaster/DeletePetCadaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petCadasterClient) CountPets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (PetCadaster_CountPetsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PetCadaster_ServiceDesc.Streams[1], "/grpc.petCadaster/CountPets", opts...)
	if err != nil {
		return nil, err
	}
	x := &petCadasterCountPetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PetCadaster_CountPetsClient interface {
	Recv() (*CountInfo, error)
	grpc.ClientStream
}

type petCadasterCountPetsClient struct {
	grpc.ClientStream
}

func (x *petCadasterCountPetsClient) Recv() (*CountInfo, error) {
	m := new(CountInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PetCadasterServer is the server API for PetCadaster service.
// All implementations must embed UnimplementedPetCadasterServer
// for forward compatibility
type PetCadasterServer interface {
	ListPets(*Empty, PetCadaster_ListPetsServer) error
	GetPet(context.Context, *ID) (*PetInfo, error)
	CreatePetCadaster(context.Context, *PetInfo) (*ID, error)
	UpdatePetCadaster(context.Context, *PetInfo) (*Status, error)
	RenamePet(context.Context, *RenameInfo) (*Status, error)
	DeletePetCadaster(context.Context, *ID) (*Status, error)
	CountPets(*Empty, PetCadaster_CountPetsServer) error
	mustEmbedUnimplementedPetCadasterServer()
}

// UnimplementedPetCadasterServer must be embedded to have forward compatible implementations.
type UnimplementedPetCadasterServer struct {
}

func (UnimplementedPetCadasterServer) ListPets(*Empty, PetCadaster_ListPetsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPets not implemented")
}
func (UnimplementedPetCadasterServer) GetPet(context.Context, *ID) (*PetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPet not implemented")
}
func (UnimplementedPetCadasterServer) CreatePetCadaster(context.Context, *PetInfo) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePetCadaster not implemented")
}
func (UnimplementedPetCadasterServer) UpdatePetCadaster(context.Context, *PetInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePetCadaster not implemented")
}
func (UnimplementedPetCadasterServer) RenamePet(context.Context, *RenameInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenamePet not implemented")
}
func (UnimplementedPetCadasterServer) DeletePetCadaster(context.Context, *ID) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePetCadaster not implemented")
}
func (UnimplementedPetCadasterServer) CountPets(*Empty, PetCadaster_CountPetsServer) error {
	return status.Errorf(codes.Unimplemented, "method CountPets not implemented")
}
func (UnimplementedPetCadasterServer) mustEmbedUnimplementedPetCadasterServer() {}

// UnsafePetCadasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetCadasterServer will
// result in compilation errors.
type UnsafePetCadasterServer interface {
	mustEmbedUnimplementedPetCadasterServer()
}

func RegisterPetCadasterServer(s grpc.ServiceRegistrar, srv PetCadasterServer) {
	s.RegisterService(&PetCadaster_ServiceDesc, srv)
}

func _PetCadaster_ListPets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PetCadasterServer).ListPets(m, &petCadasterListPetsServer{stream})
}

type PetCadaster_ListPetsServer interface {
	Send(*PetInfo) error
	grpc.ServerStream
}

type petCadasterListPetsServer struct {
	grpc.ServerStream
}

func (x *petCadasterListPetsServer) Send(m *PetInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _PetCadaster_GetPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetCadasterServer).GetPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.petCadaster/GetPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetCadasterServer).GetPet(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetCadaster_CreatePetCadaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetCadasterServer).CreatePetCadaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.petCadaster/CreatePetCadaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetCadasterServer).CreatePetCadaster(ctx, req.(*PetInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetCadaster_UpdatePetCadaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetCadasterServer).UpdatePetCadaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.petCadaster/UpdatePetCadaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetCadasterServer).UpdatePetCadaster(ctx, req.(*PetInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetCadaster_RenamePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetCadasterServer).RenamePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.petCadaster/RenamePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetCadasterServer).RenamePet(ctx, req.(*RenameInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetCadaster_DeletePetCadaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetCadasterServer).DeletePetCadaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.petCadaster/DeletePetCadaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetCadasterServer).DeletePetCadaster(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetCadaster_CountPets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PetCadasterServer).CountPets(m, &petCadasterCountPetsServer{stream})
}

type PetCadaster_CountPetsServer interface {
	Send(*CountInfo) error
	grpc.ServerStream
}

type petCadasterCountPetsServer struct {
	grpc.ServerStream
}

func (x *petCadasterCountPetsServer) Send(m *CountInfo) error {
	return x.ServerStream.SendMsg(m)
}

// PetCadaster_ServiceDesc is the grpc.ServiceDesc for PetCadaster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetCadaster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.petCadaster",
	HandlerType: (*PetCadasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPet",
			Handler:    _PetCadaster_GetPet_Handler,
		},
		{
			MethodName: "CreatePetCadaster",
			Handler:    _PetCadaster_CreatePetCadaster_Handler,
		},
		{
			MethodName: "UpdatePetCadaster",
			Handler:    _PetCadaster_UpdatePetCadaster_Handler,
		},
		{
			MethodName: "RenamePet",
			Handler:    _PetCadaster_RenamePet_Handler,
		},
		{
			MethodName: "DeletePetCadaster",
			Handler:    _PetCadaster_DeletePetCadaster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPets",
			Handler:       _PetCadaster_ListPets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CountPets",
			Handler:       _PetCadaster_CountPets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/services/petCadaster_service/ports/grpc/protos/petCadaster.proto",
}
