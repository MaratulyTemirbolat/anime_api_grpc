package main

import (
	"context"
	"log"
	"proto/pkg/api"
	"time"

	"google.golang.org/grpc"
)

const (
	port string = ":8080"
)

func main() {
	ctx := context.Background()
	var connectionStartTime time.Time = time.Now()
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", port, err)
	}
	log.Printf("Connected in %d microsec", time.Now().Sub(connectionStartTime).Microseconds())
	var userServClient api.UserServiceClient = api.NewUserServiceClient(conn)
	// var animeServClient api.AnimeServiceClient = api.NewAnimeServiceClient(conn)
	// var commentServClient api.CommentServiceClient = api.NewCommentServiceClient(conn)

	var userServStartTime time.Time = time.Now()
	actionResponse, err := userServClient.BlockUser(ctx, &api.UserAddBlockUserRequest{
		FromUserId: 1,
		ToUserId:   2,
	})
	if actionResponse != nil {
		log.Printf("The result of actionRequest is message: %v. success state: %v", actionResponse.Message, actionResponse.Success)
	}
	if err != nil {
		log.Fatalf("Could not block user: %v", err)
	}
	log.Printf("Blocked user in %d microsec", time.Now().Sub(userServStartTime).Microseconds())
}
