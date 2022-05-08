// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"proto/pkg/api"

// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// const (
// 	blockAction uint32 = 1
// )

// var message string

// // Описываем наши сервисы

// // UnimplementedCommentServiceServer
// type CommentServ struct {
// 	api.UnimplementedCommentServiceServer
// }

// func (cs *CommentServ) ShowAnimeComments(context.Context, *api.ShowAnimeCommentsRequest) (*api.ShowAnimeCommentsResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method ShowAnimeComments not implemented")
// }

// func (cs *CommentServ) CommentAnime(ctx context.Context, req *api.CommentAnimeRequest) (*api.ActionResponse, error) {
// 	if validErr := is_id_provided(req.UserId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"USER ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	} else if validErr := is_id_provided(req.AnimeId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"ANIME ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	} else if req.Content == "" {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"You cannot send empty comment",
// 		)
// 	}
// 	_, dbErr := MyDatabase.Exec("CALL leave_anime_comment($1, $2, NULL, $3);", req.UserId, req.Content, req.AnimeId)
// 	if dbErr != nil {
// 		return nil, status.Errorf(
// 			codes.Canceled,
// 			"The operation was cancelled due to: %v",
// 			dbErr,
// 		)
// 	}
// 	message = "You successfully left your comment"
// 	return &api.ActionResponse{
// 		Success: true,
// 		Message: message,
// 	}, nil
// }

// func (cs *CommentServ) ReplyUserCommentAnime(ctx context.Context, req *api.ReplyUserCommentAnimeRequest) (*api.ActionResponse, error) {
// 	if validErr := is_id_provided(req.UserId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"USER ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	} else if validErr := is_id_provided(req.AnimeId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"ANIME ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	} else if req.Content == "" {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"You cannot send empty comment",
// 		)
// 	} else if validErr := is_id_provided(req.RepliedCommentId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"REPLIED COMMENT ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	}
// 	_, dbErr := MyDatabase.Exec("CALL leave_anime_comment($1, $2, $3, $4);", req.UserId, req.Content, req.RepliedCommentId, req.AnimeId)
// 	if dbErr != nil {
// 		return nil, status.Errorf(
// 			codes.Canceled,
// 			"The operation was cancelled due to: %v",
// 			dbErr,
// 		)
// 	}
// 	message = "You successfully left your comment"
// 	return &api.ActionResponse{
// 		Success: true,
// 		Message: message,
// 	}, nil
// }
