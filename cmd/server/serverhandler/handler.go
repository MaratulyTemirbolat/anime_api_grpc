package serverhandler

import (
	"context"
	"fmt"
	"proto/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Описываем наши сервисы
// UnimplementedUserServiceServer
type UserServ struct {
	api.UnimplementedUserServiceServer
}

func (us *UserServ) BlockUser(ctx context.Context, req *api.UserAddBlockUserRequest) (*api.ActionResponse, error) {
	if req.FromUserId == req.ToUserId {
		return &api.ActionResponse{
			Success: false,
			Message: "You can not block yourself",
		}, status.Errorf(codes.Unavailable, "You can not block yourself")
	}
	var message string = fmt.Sprintf(
		"User with id '%d' blocked user with id '%d'", req.FromUserId, req.ToUserId,
	)
	return &api.ActionResponse{
		Success: true,
		Message: message,
	}, nil
}

func (us *UserServ) AddUser(context.Context, *api.UserAddBlockUserRequest) (*api.ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}

func (us *UserServ) ViewUserPage(context.Context, *api.ViewUserPageRequest) (*api.ViewUserPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewUserPage not implemented")
}

// UnimplementedAnimeServiceServer
type AnimeServ struct {
	api.UnimplementedAnimeServiceServer
}

func (as *AnimeServ) HandleAnime(context.Context, *api.UserAnimeActionRequest) (*api.ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleAnime not implemented")
}

func (as *AnimeServ) RemoveAnime(context.Context, *api.RemoveAnimeRequest) (*api.AnimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAnime not implemented")
}

func (as *AnimeServ) ViewAnime(context.Context, *api.ViewAnimeRequest) (*api.ViewAnimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAnime not implemented")
}

// UnimplementedCommentServiceServer
type CommentServ struct {
	api.UnimplementedCommentServiceServer
}

func (cs *CommentServ) ShowAnimeComments(context.Context, *api.ShowAnimeCommentsRequest) (*api.ShowAnimeCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAnimeComments not implemented")
}

func (cs *CommentServ) CommentAnime(context.Context, *api.CommentAnimeRequest) (*api.ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAnime not implemented")
}

func (cs *CommentServ) ReplyUserCommentAnime(context.Context, *api.ReplyUserCommentAnimeRequest) (*api.ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyUserCommentAnime not implemented")
}
