// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: goadesign_goagen_character.proto

package characterpb

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

// CharacterClient is the client API for Character service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterClient interface {
	// List all stored bottles
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*StoredCharacterCollection, error)
	// Show character by Id
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
	// Add new character and return its ID.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Remove character
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	// update
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type characterClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterClient(cc grpc.ClientConnInterface) CharacterClient {
	return &characterClient{cc}
}

func (c *characterClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*StoredCharacterCollection, error) {
	out := new(StoredCharacterCollection)
	err := c.cc.Invoke(ctx, "/character.Character/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, "/character.Character/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/character.Character/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/character.Character/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/character.Character/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterServer is the server API for Character service.
// All implementations must embed UnimplementedCharacterServer
// for forward compatibility
type CharacterServer interface {
	// List all stored bottles
	List(context.Context, *ListRequest) (*StoredCharacterCollection, error)
	// Show character by Id
	Show(context.Context, *ShowRequest) (*ShowResponse, error)
	// Add new character and return its ID.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Remove character
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	// update
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	mustEmbedUnimplementedCharacterServer()
}

// UnimplementedCharacterServer must be embedded to have forward compatible implementations.
type UnimplementedCharacterServer struct {
}

func (UnimplementedCharacterServer) List(context.Context, *ListRequest) (*StoredCharacterCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCharacterServer) Show(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedCharacterServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCharacterServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedCharacterServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCharacterServer) mustEmbedUnimplementedCharacterServer() {}

// UnsafeCharacterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterServer will
// result in compilation errors.
type UnsafeCharacterServer interface {
	mustEmbedUnimplementedCharacterServer()
}

func RegisterCharacterServer(s grpc.ServiceRegistrar, srv CharacterServer) {
	s.RegisterService(&Character_ServiceDesc, srv)
}

func _Character_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.Character/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Character_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.Character/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Character_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.Character/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Character_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.Character/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Character_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.Character/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Character_ServiceDesc is the grpc.ServiceDesc for Character service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Character_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "character.Character",
	HandlerType: (*CharacterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Character_List_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Character_Show_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Character_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Character_Remove_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Character_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goadesign_goagen_character.proto",
}