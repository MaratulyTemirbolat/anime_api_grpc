package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"proto/db"
	"proto/pkg/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port        string = ":8080"
	zeroValue   int    = 0
	blockAction uint32 = 1
)

var (
	message    string
	MyDatabase *sql.DB
)

func is_id_provided(userId uint32) error {
	if userId == uint32(zeroValue) {
		return status.Errorf(codes.InvalidArgument, "You can not send ID = 0")
	}
	return nil
}

// Описываем наши сервисы
// UnimplementedUserServiceServer
type UserServ struct {
	api.UnimplementedUserServiceServer
}

func (us *UserServ) BlockUser(ctx context.Context, req *api.UserAddBlockUserRequest) (*api.ActionResponse, error) {
	if err := is_id_provided(req.FromUserId); err != nil {
		return &api.ActionResponse{
			Success: false,
			Message: "You need to provide correct data.",
		}, err
	} else if err := is_id_provided(req.ToUserId); err != nil {
		return &api.ActionResponse{
			Success: false,
			Message: "The data of user that you want to block.",
		}, err
	} else if req.FromUserId == req.ToUserId {
		return &api.ActionResponse{
			Success: false,
			Message: "You can not block yourself",
		}, status.Errorf(codes.Unavailable, "You can not block yourself")
	}

	_, dbErr := MyDatabase.Exec("call block_unblock_user($1, $2, $3);", req.FromUserId, req.ToUserId, blockAction)
	if dbErr != nil {
		return &api.ActionResponse{
			Success: false,
			Message: "Blocking process canceled",
		}, status.Errorf(codes.Canceled, dbErr.Error())
	}
	message = fmt.Sprintf(
		"User with id '%d' blocked user with id '%d'", req.FromUserId, req.ToUserId,
	)
	return &api.ActionResponse{
		Success: true,
		Message: message,
	}, nil
}

func (us *UserServ) RegisterUser(ctx context.Context, req *api.UserRegisterRequest) (*api.UserLoginRegisterResponse, error) {

	_, dbErr := MyDatabase.Exec("CALL register_user($1, $2, $3, $4, $5);", req.Username, req.Email, req.Password, req.FirstName, req.LastName)
	if dbErr != nil {
		return nil, status.Errorf(codes.Canceled, "The process of registration was canceled because: ", dbErr.Error())
	}

	userVariableResponse := api.UserLoginRegisterResponse{}
	registeredUserRow, dbErr := MyDatabase.Query("SELECT id FROM users WHERE username = $1 OR email = $2;", req.Username, req.Email)
	if dbErr != nil {
		return nil, status.Errorf(codes.Canceled, "The process was canceled because: ", dbErr.Error())
	}
	defer registeredUserRow.Close()

	for registeredUserRow.Next() {
		err := registeredUserRow.Scan(&userVariableResponse.Id)
		if err != nil {
			log.Println("Appeared error during iteration: ", err)
			break
		}
	}

	return &userVariableResponse, nil
}

func (us *UserServ) LoginUser(ctx context.Context, req *api.UserLoginRequest) (*api.UserLoginRegisterResponse, error) {
	userVariableResponse := api.UserLoginRegisterResponse{}
	loginUserRow, dbErr := MyDatabase.Query("SELECT get_user_id_log_in($1, $2) AS id;", req.EmailLogin, req.Password)
	if dbErr != nil {
		return nil, status.Errorf(codes.Canceled, "The process of entering system was canceled because: ", dbErr.Error())
	}
	defer loginUserRow.Close()

	for loginUserRow.Next() {
		err := loginUserRow.Scan(&userVariableResponse.Id)
		if err != nil {
			log.Println("Appeared error during iteration: ", err)
			break
		}
	}
	return &userVariableResponse, nil
}

func (us *UserServ) ViewAllUsersInfo(ctx context.Context, req *api.ViewAllUsersInfoRequest) (*api.ViewAllUsersInfoResponse, error) {
	allUsersResponse := api.ViewAllUsersInfoResponse{}

	allUsersRow, dbErr := MyDatabase.Query("SELECT id AS userId,first_name AS firstName, last_name AS lastName, email, username FROM users WHERE deleted_at IS NULL;")
	if dbErr != nil {
		return nil, status.Errorf(codes.Canceled, "The process of getting all users was canceled because: ", dbErr.Error())
	}
	defer allUsersRow.Close()

	for allUsersRow.Next() {
		us := api.ViewUserPageResponse{}
		err := allUsersRow.Scan(&us.UserId, &us.FirstName, &us.LastName, &us.Email, &us.Username)
		if err != nil {
			log.Println("Appeared error during iteration: ", err)
			break
		}

		allUsersResponse.AllUsers = append(allUsersResponse.AllUsers, &us)
	}
	return &allUsersResponse, nil
}

func (us *UserServ) AddUser(ctx context.Context, req *api.UserAddBlockUserRequest) (*api.ActionResponse, error) {
	if errID := is_id_provided(req.FromUserId); errID != nil {
		return &api.ActionResponse{
			Success: false,
			Message: "Your ID must be provided!",
		}, errID
	} else if errID := is_id_provided(req.ToUserId); errID != nil {
		return &api.ActionResponse{
			Success: false,
			Message: "ID of the user that you want to ADD MUST BE PROVIDED",
		}, errID
	}

	_, dbErr := MyDatabase.Exec("call add_friend($1, $2);", req.FromUserId, req.ToUserId)
	if dbErr != nil {
		return &api.ActionResponse{
			Success: false,
			Message: fmt.Sprintf("During friend adding the problem appeared: \"%v\"", dbErr),
		}, dbErr
	}
	message = fmt.Sprintf(
		"User with id '%d' Added user with id '%d'", req.FromUserId, req.ToUserId,
	)
	return &api.ActionResponse{
		Success: true,
		Message: message,
	}, nil
}

func (us *UserServ) ViewUserPage(ctx context.Context, req *api.ViewUserPageRequest) (*api.ViewUserPageResponse, error) {
	if errID := is_id_provided(req.UserId); errID != nil {
		return nil, errID
	} else if errID := is_id_provided(req.VisitedUserId); errID != nil {
		return nil, errID
	}

	resultRow, dbErr := MyDatabase.Query(
		"SELECT users.id, users.first_name, users.last_name, users.email, users.username FROM users WHERE users.id = $1;", req.VisitedUserId)
	if dbErr != nil {
		return nil, dbErr
	}
	defer resultRow.Close()

	curUser := api.ViewUserPageResponse{}

	for resultRow.Next() {
		err := resultRow.Scan(&curUser.UserId, &curUser.FirstName, &curUser.LastName, &curUser.Email, &curUser.Username)
		if err != nil {
			log.Println("Appeared error during iteration: ", err)
			break
		}
	}
	if curUser.UserId != 0 {
		return &curUser, nil
	}

	return nil, status.Errorf(
		codes.NotFound,
		("The user WAS NOT FOUND with ID " + string(req.VisitedUserId)))
}

// UnimplementedAnimeServiceServer
type AnimeServ struct {
	api.UnimplementedAnimeServiceServer
}

func (as *AnimeServ) HandleAnime(ctx context.Context, req *api.UserAnimeActionRequest) (*api.ActionResponse, error) {
	if validErr := is_id_provided(req.UserId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"USER ID ERROR: %s",
			validErr.Error(),
		)
	} else if validErr := is_id_provided(req.AnimeId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ANIME ID ERROR: %s",
			validErr.Error(),
		)
	} else if validErr := is_id_provided(req.ActionId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ACTION ID ERROR: %s",
			validErr.Error(),
		)
	}

	_, dbErr := MyDatabase.Exec("call user_anime_action($1, $2, $3, $4);", req.UserId, req.AnimeId, req.ActionId, req.IsLike)
	if dbErr != nil {
		return &api.ActionResponse{
			Success: false,
			Message: "Anime Handling failed",
		}, status.Errorf(codes.Canceled, dbErr.Error())
	}
	message = fmt.Sprintf(
		"You '%d' successfully managed to Handle ANIME '%d'", req.UserId, req.AnimeId,
	)
	return &api.ActionResponse{
		Success: true,
		Message: message,
	}, nil
}

func (as *AnimeServ) RemoveAnime(ctx context.Context, req *api.RemoveAnimeRequest) (*api.AnimeResponse, error) {
	if validErr := is_id_provided(req.UserId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"USER ID ERROR: %s",
			validErr.Error(),
		)
	} else if validErr := is_id_provided(req.UserId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"USER ID ERROR: %s",
			validErr.Error(),
		)
	}

	_, dbErr := MyDatabase.Exec("call remove_anime_list($1, $2);", req.UserId, req.AnimeId)
	if dbErr != nil {
		return nil, status.Errorf(codes.Canceled, dbErr.Error())
	}
	message = fmt.Sprintf(
		"You '%d' successfully removed ANIME from your LIST '%d'", req.UserId, req.AnimeId,
	)
	return &api.AnimeResponse{
		Name:    "Need to customize",
		Success: true,
		Message: message,
	}, nil
}

func (as *AnimeServ) ViewAllAnimes(ctx context.Context, req *api.ViewAllAnimeRequest) (*api.ViewAllAnimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAllAnimes not implemented")
}

func (as *AnimeServ) ViewAnime(ctx context.Context, req *api.ViewAnimeRequest) (*api.ViewAnimeResponse, error) {
	if validErr := is_id_provided(req.AnimeId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ANIME ID ERROR: %s",
			validErr.Error(),
		)
	} else if validErr := is_id_provided(req.UserId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ANIME ID ERROR: %s",
			validErr.Error(),
		)
	}
	resultRow, dbErr := MyDatabase.Query(
		"SELECT animes.id, animes.name, animes.description, animes.release_year, anime_groups.name, animes.rating, animes.views_number, studios.name, genres.name FROM animes LEFT JOIN anime_genres ON animes.id = anime_genres.anime_id LEFT JOIN genres ON genres.id = anime_genres.genre_id LEFT JOIN anime_groups ON animes.group_id = anime_groups.id LEFT JOIN studios ON animes.studio_id = studios.id WHERE animes.id = $1;", req.AnimeId)

	if dbErr != nil {
		return nil, dbErr
	}
	var releaseYear uint32 = 0
	var resultAnimeResponse api.ViewAnimeResponse = api.ViewAnimeResponse{}
	var animeGenres []string
	for resultRow.Next() {
		var curGenre string
		err := resultRow.Scan(
			&resultAnimeResponse.AnimeId,
			&resultAnimeResponse.Name,
			&resultAnimeResponse.Description,
			&releaseYear,
			&resultAnimeResponse.AnimeGroup,
			&resultAnimeResponse.Rating,
			&resultAnimeResponse.ViewsNumber,
			&resultAnimeResponse.StudioName,
			&curGenre,
		)
		if err != nil {
			log.Println("Appeared error during iteration in query: ", err)
			break
		}
		animeGenres = append(animeGenres, curGenre)
	}
	if releaseYear == uint32(zeroValue) {
		return nil, status.Errorf(codes.NotFound, "Anime with ID %d was not FOUND", req.AnimeId)
	}
	resultAnimeResponse.Genres = animeGenres
	resultAnimeResponse.ReleaseDate = string(releaseYear)

	return &resultAnimeResponse, nil
}

// UnimplementedCommentServiceServer
type CommentServ struct {
	api.UnimplementedCommentServiceServer
}

func (cs *CommentServ) ShowAnimeComments(context.Context, *api.ShowAnimeCommentsRequest) (*api.ShowAnimeCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAnimeComments not implemented")
}

func (cs *CommentServ) CommentAnime(ctx context.Context, req *api.CommentAnimeRequest) (*api.ActionResponse, error) {
	if validErr := is_id_provided(req.UserId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"USER ID ERROR: %s",
			validErr.Error(),
		)
	} else if validErr := is_id_provided(req.AnimeId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ANIME ID ERROR: %s",
			validErr.Error(),
		)
	} else if req.Content == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"You cannot send empty comment",
		)
	}
	_, dbErr := MyDatabase.Exec("CALL leave_anime_comment($1, $2, NULL, $3);", req.UserId, req.Content, req.AnimeId)
	if dbErr != nil {
		return nil, status.Errorf(
			codes.Canceled,
			"The operation was cancelled due to: %v",
			dbErr,
		)
	}
	message = "You successfully left your comment"
	return &api.ActionResponse{
		Success: true,
		Message: message,
	}, nil
}

func (cs *CommentServ) ReplyUserCommentAnime(ctx context.Context, req *api.ReplyUserCommentAnimeRequest) (*api.ActionResponse, error) {
	if validErr := is_id_provided(req.UserId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"USER ID ERROR: %s",
			validErr.Error(),
		)
	} else if validErr := is_id_provided(req.AnimeId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ANIME ID ERROR: %s",
			validErr.Error(),
		)
	} else if req.Content == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"You cannot send empty comment",
		)
	} else if validErr := is_id_provided(req.RepliedCommentId); validErr != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"REPLIED COMMENT ID ERROR: %s",
			validErr.Error(),
		)
	}
	_, dbErr := MyDatabase.Exec("CALL leave_anime_comment($1, $2, $3, $4);", req.UserId, req.Content, req.RepliedCommentId, req.AnimeId)
	if dbErr != nil {
		return nil, status.Errorf(
			codes.Canceled,
			"The operation was cancelled due to: %v",
			dbErr,
		)
	}
	message = "You successfully left your comment"
	return &api.ActionResponse{
		Success: true,
		Message: message,
	}, nil
}

func main() {
	MyDatabase = db.ConnectGenDB()
	defer MyDatabase.Close()
	listener, err := net.Listen("tcp", port) // То, что мы используем для слушанья сервера
	if err != nil {
		log.Fatalf("Cannot listen to %s: %v", port, err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()         // Наш сервер, но он еще не слушает порт
	var userServ *UserServ = new(UserServ) // Делаем instance наших имплементированных сервисов
	var animeServ *AnimeServ = new(AnimeServ)
	var commentServ *CommentServ = new(CommentServ)

	// Связывание наших Сервисов с GRPC Server
	api.RegisterUserServiceServer(grpcServer, userServ)
	api.RegisterAnimeServiceServer(grpcServer, animeServ)
	api.RegisterCommentServiceServer(grpcServer, commentServ)

	log.Printf("Serving on %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil { // Serve - слушать порт
		log.Fatalf("Failed to serve on %v: %v", listener.Addr(), err)
	}
}
