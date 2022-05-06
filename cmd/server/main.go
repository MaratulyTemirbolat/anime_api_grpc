package main

import (
	"database/sql"
	"log"
	"net"
	"proto/cmd/server/serverhandler"
	"proto/db"
	"proto/pkg/api"

	"google.golang.org/grpc"
)

const (
	port string = ":8080"
)

func main() {
	// Подключаемся к нашей базе данных черезе стандартную библиотеку sql
	var MyDatabase *sql.DB = db.ConnectGenDB()

	// Не забываем закрыть базу данных в конце
	defer MyDatabase.Close()

	listener, err := net.Listen("tcp", port) // То, что мы используем для слушанья сервера
	if err != nil {
		log.Fatalf("Cannot listen to %s: %v", port, err)
	}
	defer listener.Close()

	// Наш сервер, но он еще не слушает порт
	grpcServer := grpc.NewServer()
	// Делаем instance наших имплементированных сервисов
	var userServ *serverhandler.UserServ = new(serverhandler.UserServ)
	var animeServ *serverhandler.AnimeServ = new(serverhandler.AnimeServ)
	var commentServ *serverhandler.CommentServ = new(serverhandler.CommentServ)

	// Связывание наших Сервисов с GRPC Server
	api.RegisterUserServiceServer(grpcServer, userServ)
	api.RegisterAnimeServiceServer(grpcServer, animeServ)
	api.RegisterCommentServiceServer(grpcServer, commentServ)

	log.Printf("Serving ont %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil { // Serve - слушать порт
		log.Fatalf("Failed to serve on %v: %v", listener.Addr(), err)
	}
}
