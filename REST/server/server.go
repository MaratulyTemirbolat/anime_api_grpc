package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"tools/db"
	pb "tools/gen/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	MyDatabase *sql.DB
	zeroValue  int = 0
)

func is_id_provided(userId uint32) error {
	if userId == uint32(zeroValue) {
		return status.Errorf(codes.InvalidArgument, "You can not send ID = 0")
	}
	return nil
}

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.Response, error) {
	if validErr := is_id_provided(req.Id); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ANIME ID ERROR: %s",
			validErr.Error(),
		)
	} else if req.Message == "" {
		return nil, status.Error(
			codes.Canceled,
			"You cannot send EMPTY COMMENT",
		)
	}
	_, dbErr := MyDatabase.Exec("CALL leave_anime_comment_random_user($1,$2);", req.Message, req.Id)
	if dbErr != nil {
		return nil, status.Errorf(codes.Canceled, "Appeared error during comment leaving: %v ", dbErr.Error())
	}

	return nil, status.Errorf(codes.OK, "The comment is left")
}
func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	resultRow, dbErr := MyDatabase.Query("SELECT username, email from users WHERE id = $1 LIMIT 1;", req.Id)
	if dbErr != nil {
		return nil, status.Errorf(codes.Canceled, "The process was canceled because: ", dbErr.Error())
	}
	defer resultRow.Close()

	var foundUser pb.UserResponse = pb.UserResponse{}

	for resultRow.Next() {
		err := resultRow.Scan(&foundUser.Username, &foundUser.Email)
		if err != nil {
			log.Println("Appeared error during iteration: ", err)
			break
		}
	}

	if len(foundUser.Username) == 0 {
		return nil, status.Errorf(codes.NotFound, "The user with ID: %d not found", req.Id)
	}
	return &foundUser, nil
}

func main() {
	MyDatabase = db.ConnectGenDB()
	defer MyDatabase.Close()

	go func() {
		// mux
		mux := runtime.NewServeMux()

		// register
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})

		// http server
		log.Fatalln(http.ListenAndServe("localhost:8082", mux))
	}()

	listner, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("Cannot listen to 8081: %v", err)
	}
	defer listner.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	log.Printf("Serving on %v", listner.Addr())
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Println(err)
	}
}
