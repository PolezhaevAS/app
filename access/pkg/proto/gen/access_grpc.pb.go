// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: access.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccessClient is the client API for Access service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessClient interface {
	// Get list groups
	// Request: last id, limit
	// Response: list groups
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Get group by id
	// Request: id
	// Response: group
	Group(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
	// Create new group
	// Request: name, description
	// Response:
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Update group by id
	// Request: id, name, description
	// Response:
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete group by id
	// Request: id
	// Response:
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get users from group
	// Request: group id
	// Response: list users id
	Users(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	// Add users into group by group id
	// Request: group id, user id
	// Response:
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Remove user from group by group id and user id
	// Request: group id, user id
	// Response:
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get list services
	// Request:
	// Response: list services
	ListServices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// Add method into group
	// Request: group id, method id
	// Response:
	AddMethod(ctx context.Context, in *AddMethodRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Remove method from group
	// Request: group id, method id
	// Response:
	RemoveMethod(ctx context.Context, in *RemoveMethodRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type accessClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessClient(cc grpc.ClientConnInterface) AccessClient {
	return &accessClient{cc}
}

func (c *accessClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) Group(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/Group", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) Users(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/Users", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) ListServices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) AddMethod(ctx context.Context, in *AddMethodRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/AddMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) RemoveMethod(ctx context.Context, in *RemoveMethodRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/access.grpc.Access/RemoveMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessServer is the server API for Access service.
// All implementations must embed UnimplementedAccessServer
// for forward compatibility
type AccessServer interface {
	// Get list groups
	// Request: last id, limit
	// Response: list groups
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Get group by id
	// Request: id
	// Response: group
	Group(context.Context, *GroupRequest) (*GroupResponse, error)
	// Create new group
	// Request: name, description
	// Response:
	CreateGroup(context.Context, *CreateGroupRequest) (*empty.Empty, error)
	// Update group by id
	// Request: id, name, description
	// Response:
	UpdateGroup(context.Context, *UpdateGroupRequest) (*empty.Empty, error)
	// Delete group by id
	// Request: id
	// Response:
	DeleteGroup(context.Context, *DeleteGroupRequest) (*empty.Empty, error)
	// Get users from group
	// Request: group id
	// Response: list users id
	Users(context.Context, *UsersRequest) (*UsersResponse, error)
	// Add users into group by group id
	// Request: group id, user id
	// Response:
	AddUser(context.Context, *AddUserRequest) (*empty.Empty, error)
	// Remove user from group by group id and user id
	// Request: group id, user id
	// Response:
	RemoveUser(context.Context, *RemoveUserRequest) (*empty.Empty, error)
	// Get list services
	// Request:
	// Response: list services
	ListServices(context.Context, *empty.Empty) (*ListServicesResponse, error)
	// Add method into group
	// Request: group id, method id
	// Response:
	AddMethod(context.Context, *AddMethodRequest) (*empty.Empty, error)
	// Remove method from group
	// Request: group id, method id
	// Response:
	RemoveMethod(context.Context, *RemoveMethodRequest) (*empty.Empty, error)
	mustEmbedUnimplementedAccessServer()
}

// UnimplementedAccessServer must be embedded to have forward compatible implementations.
type UnimplementedAccessServer struct {
}

func (UnimplementedAccessServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccessServer) Group(context.Context, *GroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Group not implemented")
}
func (UnimplementedAccessServer) CreateGroup(context.Context, *CreateGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedAccessServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedAccessServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedAccessServer) Users(context.Context, *UsersRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Users not implemented")
}
func (UnimplementedAccessServer) AddUser(context.Context, *AddUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedAccessServer) RemoveUser(context.Context, *RemoveUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedAccessServer) ListServices(context.Context, *empty.Empty) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedAccessServer) AddMethod(context.Context, *AddMethodRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMethod not implemented")
}
func (UnimplementedAccessServer) RemoveMethod(context.Context, *RemoveMethodRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMethod not implemented")
}
func (UnimplementedAccessServer) mustEmbedUnimplementedAccessServer() {}

// UnsafeAccessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessServer will
// result in compilation errors.
type UnsafeAccessServer interface {
	mustEmbedUnimplementedAccessServer()
}

func RegisterAccessServer(s grpc.ServiceRegistrar, srv AccessServer) {
	s.RegisterService(&Access_ServiceDesc, srv)
}

func _Access_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_Group_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).Group(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/Group",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).Group(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_Users_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).Users(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/Users",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).Users(ctx, req.(*UsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).ListServices(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_AddMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).AddMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/AddMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).AddMethod(ctx, req.(*AddMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_RemoveMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).RemoveMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.grpc.Access/RemoveMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).RemoveMethod(ctx, req.(*RemoveMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Access_ServiceDesc is the grpc.ServiceDesc for Access service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Access_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "access.grpc.Access",
	HandlerType: (*AccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Access_List_Handler,
		},
		{
			MethodName: "Group",
			Handler:    _Access_Group_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _Access_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _Access_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _Access_DeleteGroup_Handler,
		},
		{
			MethodName: "Users",
			Handler:    _Access_Users_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _Access_AddUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _Access_RemoveUser_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _Access_ListServices_Handler,
		},
		{
			MethodName: "AddMethod",
			Handler:    _Access_AddMethod_Handler,
		},
		{
			MethodName: "RemoveMethod",
			Handler:    _Access_RemoveMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access.proto",
}
