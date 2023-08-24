// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: content.proto

package pb

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
	Content_AddComment_FullMethodName                    = "/pb.content/AddComment"
	Content_DelComment_FullMethodName                    = "/pb.content/DelComment"
	Content_GetCommentById_FullMethodName                = "/pb.content/GetCommentById"
	Content_AddFavorite_FullMethodName                   = "/pb.content/AddFavorite"
	Content_UpdateFavorite_FullMethodName                = "/pb.content/UpdateFavorite"
	Content_DelFavorite_FullMethodName                   = "/pb.content/DelFavorite"
	Content_SearchFavorite_FullMethodName                = "/pb.content/SearchFavorite"
	Content_GetVideoById_FullMethodName                  = "/pb.content/GetVideoById"
	Content_GetFeedList_FullMethodName                   = "/pb.content/GetFeedList"
	Content_GetPublishList_FullMethodName                = "/pb.content/GetPublishList"
	Content_GetWorkCountByUserId_FullMethodName          = "/pb.content/GetWorkCountByUserId"
	Content_GetFavoriteCountByUserId_FullMethodName      = "/pb.content/GetFavoriteCountByUserId"
	Content_GetUserFavoritedCnt_FullMethodName           = "/pb.content/GetUserFavoritedCnt"
	Content_GetUserPublishAndLikedCntById_FullMethodName = "/pb.content/GetUserPublishAndLikedCntById"
	Content_GetVideoListByIdList_FullMethodName          = "/pb.content/GetVideoListByIdList"
)

// ContentClient is the client API for Content service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentClient interface {
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
	DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error)
	GetCommentById(ctx context.Context, in *GetCommentByIdReq, opts ...grpc.CallOption) (*GetCommentByIdResp, error)
	AddFavorite(ctx context.Context, in *AddFavoriteReq, opts ...grpc.CallOption) (*AddFavoriteResp, error)
	UpdateFavorite(ctx context.Context, in *UpdateFavoriteReq, opts ...grpc.CallOption) (*UpdateFavoriteResp, error)
	DelFavorite(ctx context.Context, in *DelFavoriteReq, opts ...grpc.CallOption) (*DelFavoriteResp, error)
	SearchFavorite(ctx context.Context, in *SearchFavoriteReq, opts ...grpc.CallOption) (*SearchFavoriteResp, error)
	GetVideoById(ctx context.Context, in *GetVideoByIdReq, opts ...grpc.CallOption) (*GetVideoByIdResp, error)
	GetFeedList(ctx context.Context, in *FeedListReq, opts ...grpc.CallOption) (*FeedListResp, error)
	GetPublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error)
	GetWorkCountByUserId(ctx context.Context, in *GetWorkCountByUserIdReq, opts ...grpc.CallOption) (*GetWorkCountByUserIdResp, error)
	GetFavoriteCountByUserId(ctx context.Context, in *GetFavoriteCountByUserIdReq, opts ...grpc.CallOption) (*GetFavoriteCountByUserIdResp, error)
	GetUserFavoritedCnt(ctx context.Context, in *GetUserFavoritedCntByIdReq, opts ...grpc.CallOption) (*GetUserFavoritedCntByIdResp, error)
	GetUserPublishAndLikedCntById(ctx context.Context, in *GetUserPublishAndLikedCntByIdReq, opts ...grpc.CallOption) (*GetUserPublishAndLikedCntByIdResp, error)
	GetVideoListByIdList(ctx context.Context, in *GetVideoListByIdListReq, opts ...grpc.CallOption) (*GetVideoListByIdListResp, error)
}

type contentClient struct {
	cc grpc.ClientConnInterface
}

func NewContentClient(cc grpc.ClientConnInterface) ContentClient {
	return &contentClient{cc}
}

func (c *contentClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, Content_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error) {
	out := new(DelCommentResp)
	err := c.cc.Invoke(ctx, Content_DelComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetCommentById(ctx context.Context, in *GetCommentByIdReq, opts ...grpc.CallOption) (*GetCommentByIdResp, error) {
	out := new(GetCommentByIdResp)
	err := c.cc.Invoke(ctx, Content_GetCommentById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) AddFavorite(ctx context.Context, in *AddFavoriteReq, opts ...grpc.CallOption) (*AddFavoriteResp, error) {
	out := new(AddFavoriteResp)
	err := c.cc.Invoke(ctx, Content_AddFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) UpdateFavorite(ctx context.Context, in *UpdateFavoriteReq, opts ...grpc.CallOption) (*UpdateFavoriteResp, error) {
	out := new(UpdateFavoriteResp)
	err := c.cc.Invoke(ctx, Content_UpdateFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) DelFavorite(ctx context.Context, in *DelFavoriteReq, opts ...grpc.CallOption) (*DelFavoriteResp, error) {
	out := new(DelFavoriteResp)
	err := c.cc.Invoke(ctx, Content_DelFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) SearchFavorite(ctx context.Context, in *SearchFavoriteReq, opts ...grpc.CallOption) (*SearchFavoriteResp, error) {
	out := new(SearchFavoriteResp)
	err := c.cc.Invoke(ctx, Content_SearchFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetVideoById(ctx context.Context, in *GetVideoByIdReq, opts ...grpc.CallOption) (*GetVideoByIdResp, error) {
	out := new(GetVideoByIdResp)
	err := c.cc.Invoke(ctx, Content_GetVideoById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetFeedList(ctx context.Context, in *FeedListReq, opts ...grpc.CallOption) (*FeedListResp, error) {
	out := new(FeedListResp)
	err := c.cc.Invoke(ctx, Content_GetFeedList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetPublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	out := new(PublishListResp)
	err := c.cc.Invoke(ctx, Content_GetPublishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetWorkCountByUserId(ctx context.Context, in *GetWorkCountByUserIdReq, opts ...grpc.CallOption) (*GetWorkCountByUserIdResp, error) {
	out := new(GetWorkCountByUserIdResp)
	err := c.cc.Invoke(ctx, Content_GetWorkCountByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetFavoriteCountByUserId(ctx context.Context, in *GetFavoriteCountByUserIdReq, opts ...grpc.CallOption) (*GetFavoriteCountByUserIdResp, error) {
	out := new(GetFavoriteCountByUserIdResp)
	err := c.cc.Invoke(ctx, Content_GetFavoriteCountByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetUserFavoritedCnt(ctx context.Context, in *GetUserFavoritedCntByIdReq, opts ...grpc.CallOption) (*GetUserFavoritedCntByIdResp, error) {
	out := new(GetUserFavoritedCntByIdResp)
	err := c.cc.Invoke(ctx, Content_GetUserFavoritedCnt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetUserPublishAndLikedCntById(ctx context.Context, in *GetUserPublishAndLikedCntByIdReq, opts ...grpc.CallOption) (*GetUserPublishAndLikedCntByIdResp, error) {
	out := new(GetUserPublishAndLikedCntByIdResp)
	err := c.cc.Invoke(ctx, Content_GetUserPublishAndLikedCntById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetVideoListByIdList(ctx context.Context, in *GetVideoListByIdListReq, opts ...grpc.CallOption) (*GetVideoListByIdListResp, error) {
	out := new(GetVideoListByIdListResp)
	err := c.cc.Invoke(ctx, Content_GetVideoListByIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServer is the server API for Content service.
// All implementations must embed UnimplementedContentServer
// for forward compatibility
type ContentServer interface {
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
	DelComment(context.Context, *DelCommentReq) (*DelCommentResp, error)
	GetCommentById(context.Context, *GetCommentByIdReq) (*GetCommentByIdResp, error)
	AddFavorite(context.Context, *AddFavoriteReq) (*AddFavoriteResp, error)
	UpdateFavorite(context.Context, *UpdateFavoriteReq) (*UpdateFavoriteResp, error)
	DelFavorite(context.Context, *DelFavoriteReq) (*DelFavoriteResp, error)
	SearchFavorite(context.Context, *SearchFavoriteReq) (*SearchFavoriteResp, error)
	GetVideoById(context.Context, *GetVideoByIdReq) (*GetVideoByIdResp, error)
	GetFeedList(context.Context, *FeedListReq) (*FeedListResp, error)
	GetPublishList(context.Context, *PublishListReq) (*PublishListResp, error)
	GetWorkCountByUserId(context.Context, *GetWorkCountByUserIdReq) (*GetWorkCountByUserIdResp, error)
	GetFavoriteCountByUserId(context.Context, *GetFavoriteCountByUserIdReq) (*GetFavoriteCountByUserIdResp, error)
	GetUserFavoritedCnt(context.Context, *GetUserFavoritedCntByIdReq) (*GetUserFavoritedCntByIdResp, error)
	GetUserPublishAndLikedCntById(context.Context, *GetUserPublishAndLikedCntByIdReq) (*GetUserPublishAndLikedCntByIdResp, error)
	GetVideoListByIdList(context.Context, *GetVideoListByIdListReq) (*GetVideoListByIdListResp, error)
	mustEmbedUnimplementedContentServer()
}

// UnimplementedContentServer must be embedded to have forward compatible implementations.
type UnimplementedContentServer struct {
}

func (UnimplementedContentServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedContentServer) DelComment(context.Context, *DelCommentReq) (*DelCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelComment not implemented")
}
func (UnimplementedContentServer) GetCommentById(context.Context, *GetCommentByIdReq) (*GetCommentByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedContentServer) AddFavorite(context.Context, *AddFavoriteReq) (*AddFavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFavorite not implemented")
}
func (UnimplementedContentServer) UpdateFavorite(context.Context, *UpdateFavoriteReq) (*UpdateFavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFavorite not implemented")
}
func (UnimplementedContentServer) DelFavorite(context.Context, *DelFavoriteReq) (*DelFavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFavorite not implemented")
}
func (UnimplementedContentServer) SearchFavorite(context.Context, *SearchFavoriteReq) (*SearchFavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFavorite not implemented")
}
func (UnimplementedContentServer) GetVideoById(context.Context, *GetVideoByIdReq) (*GetVideoByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoById not implemented")
}
func (UnimplementedContentServer) GetFeedList(context.Context, *FeedListReq) (*FeedListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedList not implemented")
}
func (UnimplementedContentServer) GetPublishList(context.Context, *PublishListReq) (*PublishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishList not implemented")
}
func (UnimplementedContentServer) GetWorkCountByUserId(context.Context, *GetWorkCountByUserIdReq) (*GetWorkCountByUserIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkCountByUserId not implemented")
}
func (UnimplementedContentServer) GetFavoriteCountByUserId(context.Context, *GetFavoriteCountByUserIdReq) (*GetFavoriteCountByUserIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteCountByUserId not implemented")
}
func (UnimplementedContentServer) GetUserFavoritedCnt(context.Context, *GetUserFavoritedCntByIdReq) (*GetUserFavoritedCntByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavoritedCnt not implemented")
}
func (UnimplementedContentServer) GetUserPublishAndLikedCntById(context.Context, *GetUserPublishAndLikedCntByIdReq) (*GetUserPublishAndLikedCntByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPublishAndLikedCntById not implemented")
}
func (UnimplementedContentServer) GetVideoListByIdList(context.Context, *GetVideoListByIdListReq) (*GetVideoListByIdListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoListByIdList not implemented")
}
func (UnimplementedContentServer) mustEmbedUnimplementedContentServer() {}

// UnsafeContentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServer will
// result in compilation errors.
type UnsafeContentServer interface {
	mustEmbedUnimplementedContentServer()
}

func RegisterContentServer(s grpc.ServiceRegistrar, srv ContentServer) {
	s.RegisterService(&Content_ServiceDesc, srv)
}

func _Content_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_DelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).DelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_DelComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).DelComment(ctx, req.(*DelCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetCommentById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetCommentById(ctx, req.(*GetCommentByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_AddFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_AddFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddFavorite(ctx, req.(*AddFavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_UpdateFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).UpdateFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_UpdateFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).UpdateFavorite(ctx, req.(*UpdateFavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_DelFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).DelFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_DelFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).DelFavorite(ctx, req.(*DelFavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_SearchFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).SearchFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_SearchFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).SearchFavorite(ctx, req.(*SearchFavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetVideoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetVideoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetVideoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetVideoById(ctx, req.(*GetVideoByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetFeedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetFeedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetFeedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetFeedList(ctx, req.(*FeedListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetPublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetPublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetPublishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetPublishList(ctx, req.(*PublishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetWorkCountByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkCountByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetWorkCountByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetWorkCountByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetWorkCountByUserId(ctx, req.(*GetWorkCountByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetFavoriteCountByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoriteCountByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetFavoriteCountByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetFavoriteCountByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetFavoriteCountByUserId(ctx, req.(*GetFavoriteCountByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetUserFavoritedCnt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFavoritedCntByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetUserFavoritedCnt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetUserFavoritedCnt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetUserFavoritedCnt(ctx, req.(*GetUserFavoritedCntByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetUserPublishAndLikedCntById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPublishAndLikedCntByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetUserPublishAndLikedCntById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetUserPublishAndLikedCntById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetUserPublishAndLikedCntById(ctx, req.(*GetUserPublishAndLikedCntByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetVideoListByIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoListByIdListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetVideoListByIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetVideoListByIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetVideoListByIdList(ctx, req.(*GetVideoListByIdListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Content_ServiceDesc is the grpc.ServiceDesc for Content service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Content_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.content",
	HandlerType: (*ContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddComment",
			Handler:    _Content_AddComment_Handler,
		},
		{
			MethodName: "DelComment",
			Handler:    _Content_DelComment_Handler,
		},
		{
			MethodName: "GetCommentById",
			Handler:    _Content_GetCommentById_Handler,
		},
		{
			MethodName: "AddFavorite",
			Handler:    _Content_AddFavorite_Handler,
		},
		{
			MethodName: "UpdateFavorite",
			Handler:    _Content_UpdateFavorite_Handler,
		},
		{
			MethodName: "DelFavorite",
			Handler:    _Content_DelFavorite_Handler,
		},
		{
			MethodName: "SearchFavorite",
			Handler:    _Content_SearchFavorite_Handler,
		},
		{
			MethodName: "GetVideoById",
			Handler:    _Content_GetVideoById_Handler,
		},
		{
			MethodName: "GetFeedList",
			Handler:    _Content_GetFeedList_Handler,
		},
		{
			MethodName: "GetPublishList",
			Handler:    _Content_GetPublishList_Handler,
		},
		{
			MethodName: "GetWorkCountByUserId",
			Handler:    _Content_GetWorkCountByUserId_Handler,
		},
		{
			MethodName: "GetFavoriteCountByUserId",
			Handler:    _Content_GetFavoriteCountByUserId_Handler,
		},
		{
			MethodName: "GetUserFavoritedCnt",
			Handler:    _Content_GetUserFavoritedCnt_Handler,
		},
		{
			MethodName: "GetUserPublishAndLikedCntById",
			Handler:    _Content_GetUserPublishAndLikedCntById_Handler,
		},
		{
			MethodName: "GetVideoListByIdList",
			Handler:    _Content_GetVideoListByIdList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}
