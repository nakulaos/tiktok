// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: comment.proto

package comment

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

const (
	CommentSrv_CommentAction_FullMethodName        = "/comment.CommentSrv/CommentAction"
	CommentSrv_CommentList_FullMethodName          = "/comment.CommentSrv/CommentList"
	CommentSrv_CommentActionRevert_FullMethodName  = "/comment.CommentSrv/CommentActionRevert"
	CommentSrv_DanMuAction_FullMethodName          = "/comment.CommentSrv/DanMuAction"
	CommentSrv_DanMuList_FullMethodName            = "/comment.CommentSrv/DanMuList"
	CommentSrv_PrepareCommentAction_FullMethodName = "/comment.CommentSrv/PrepareCommentAction"
	CommentSrv_FindComment_FullMethodName          = "/comment.CommentSrv/FindComment"
)

// CommentSrvClient is the client API for CommentSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentSrvClient interface {
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
	CommentActionRevert(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	DanMuAction(ctx context.Context, in *DanmuActionRequest, opts ...grpc.CallOption) (*DanmuActionResponse, error)
	DanMuList(ctx context.Context, in *DanmuListRequest, opts ...grpc.CallOption) (*DanmuListResponse, error)
	PrepareCommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*PrepareCommentAction, error)
	FindComment(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*FindCommentResp, error)
}

type commentSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentSrvClient(cc grpc.ClientConnInterface) CommentSrvClient {
	return &commentSrvClient{cc}
}

func (c *commentSrvClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, CommentSrv_CommentAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentSrvClient) CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	out := new(CommentListResponse)
	err := c.cc.Invoke(ctx, CommentSrv_CommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentSrvClient) CommentActionRevert(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, CommentSrv_CommentActionRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentSrvClient) DanMuAction(ctx context.Context, in *DanmuActionRequest, opts ...grpc.CallOption) (*DanmuActionResponse, error) {
	out := new(DanmuActionResponse)
	err := c.cc.Invoke(ctx, CommentSrv_DanMuAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentSrvClient) DanMuList(ctx context.Context, in *DanmuListRequest, opts ...grpc.CallOption) (*DanmuListResponse, error) {
	out := new(DanmuListResponse)
	err := c.cc.Invoke(ctx, CommentSrv_DanMuList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentSrvClient) PrepareCommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*PrepareCommentAction, error) {
	out := new(PrepareCommentAction)
	err := c.cc.Invoke(ctx, CommentSrv_PrepareCommentAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentSrvClient) FindComment(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*FindCommentResp, error) {
	out := new(FindCommentResp)
	err := c.cc.Invoke(ctx, CommentSrv_FindComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentSrvServer is the server API for CommentSrv service.
// All implementations must embed UnimplementedCommentSrvServer
// for forward compatibility
type CommentSrvServer interface {
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error)
	CommentActionRevert(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	DanMuAction(context.Context, *DanmuActionRequest) (*DanmuActionResponse, error)
	DanMuList(context.Context, *DanmuListRequest) (*DanmuListResponse, error)
	PrepareCommentAction(context.Context, *CommentActionRequest) (*PrepareCommentAction, error)
	FindComment(context.Context, *CommentActionRequest) (*FindCommentResp, error)
	mustEmbedUnimplementedCommentSrvServer()
}

// UnimplementedCommentSrvServer must be embedded to have forward compatible implementations.
type UnimplementedCommentSrvServer struct {
}

func (UnimplementedCommentSrvServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedCommentSrvServer) CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedCommentSrvServer) CommentActionRevert(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentActionRevert not implemented")
}
func (UnimplementedCommentSrvServer) DanMuAction(context.Context, *DanmuActionRequest) (*DanmuActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DanMuAction not implemented")
}
func (UnimplementedCommentSrvServer) DanMuList(context.Context, *DanmuListRequest) (*DanmuListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DanMuList not implemented")
}
func (UnimplementedCommentSrvServer) PrepareCommentAction(context.Context, *CommentActionRequest) (*PrepareCommentAction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareCommentAction not implemented")
}
func (UnimplementedCommentSrvServer) FindComment(context.Context, *CommentActionRequest) (*FindCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindComment not implemented")
}
func (UnimplementedCommentSrvServer) mustEmbedUnimplementedCommentSrvServer() {}

// UnsafeCommentSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentSrvServer will
// result in compilation errors.
type UnsafeCommentSrvServer interface {
	mustEmbedUnimplementedCommentSrvServer()
}

func RegisterCommentSrvServer(s grpc.ServiceRegistrar, srv CommentSrvServer) {
	s.RegisterService(&CommentSrv_ServiceDesc, srv)
}

func _CommentSrv_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_CommentAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentSrv_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_CommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).CommentList(ctx, req.(*CommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentSrv_CommentActionRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).CommentActionRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_CommentActionRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).CommentActionRevert(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentSrv_DanMuAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DanmuActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).DanMuAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_DanMuAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).DanMuAction(ctx, req.(*DanmuActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentSrv_DanMuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DanmuListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).DanMuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_DanMuList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).DanMuList(ctx, req.(*DanmuListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentSrv_PrepareCommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).PrepareCommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_PrepareCommentAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).PrepareCommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentSrv_FindComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentSrvServer).FindComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentSrv_FindComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentSrvServer).FindComment(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentSrv_ServiceDesc is the grpc.ServiceDesc for CommentSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentSrv",
	HandlerType: (*CommentSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentAction",
			Handler:    _CommentSrv_CommentAction_Handler,
		},
		{
			MethodName: "CommentList",
			Handler:    _CommentSrv_CommentList_Handler,
		},
		{
			MethodName: "CommentActionRevert",
			Handler:    _CommentSrv_CommentActionRevert_Handler,
		},
		{
			MethodName: "DanMuAction",
			Handler:    _CommentSrv_DanMuAction_Handler,
		},
		{
			MethodName: "DanMuList",
			Handler:    _CommentSrv_DanMuList_Handler,
		},
		{
			MethodName: "PrepareCommentAction",
			Handler:    _CommentSrv_PrepareCommentAction_Handler,
		},
		{
			MethodName: "FindComment",
			Handler:    _CommentSrv_FindComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
