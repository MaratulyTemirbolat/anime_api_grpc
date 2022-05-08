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
// // UnimplementedUserServiceServer
// type UserServ struct {
// 	api.UnimplementedUserServiceServer
// }

// func (us *UserServ) BlockUser(ctx context.Context, req *api.UserAddBlockUserRequest) (*api.ActionResponse, error) {
// 	if err := is_id_provided(req.FromUserId); err != nil {
// 		return &api.ActionResponse{
// 			Success: false,
// 			Message: "You need to provide correct data.",
// 		}, err
// 	} else if err := is_id_provided(req.ToUserId); err != nil {
// 		return &api.ActionResponse{
// 			Success: false,
// 			Message: "The data of user that you want to block.",
// 		}, err
// 	} else if req.FromUserId == req.ToUserId {
// 		return &api.ActionResponse{
// 			Success: false,
// 			Message: "You can not block yourself",
// 		}, status.Errorf(codes.Unavailable, "You can not block yourself")
// 	}

// 	_, dbErr := MyDatabase.Exec("call block_unblock_user($1, $2, $3);", req.FromUserId, req.ToUserId, blockAction)
// 	if dbErr != nil {
// 		return &api.ActionResponse{
// 			Success: false,
// 			Message: "Blocking process canceled",
// 		}, status.Errorf(codes.Canceled, dbErr.Error())
// 	}
// 	message = fmt.Sprintf(
// 		"User with id '%d' blocked user with id '%d'", req.FromUserId, req.ToUserId,
// 	)
// 	return &api.ActionResponse{
// 		Success: true,
// 		Message: message,
// 	}, nil
// }

// func (us *UserServ) AddUser(ctx context.Context, req *api.UserAddBlockUserRequest) (*api.ActionResponse, error) {
// 	if errID := is_id_provided(req.FromUserId); errID != nil {
// 		return &api.ActionResponse{
// 			Success: false,
// 			Message: "Your ID must be provided!",
// 		}, errID
// 	} else if errID := is_id_provided(req.ToUserId); errID != nil {
// 		return &api.ActionResponse{
// 			Success: false,
// 			Message: "ID of the user that you want to ADD MUST BE PROVIDED",
// 		}, errID
// 	}

// 	_, dbErr := MyDatabase.Exec("call add_friend($1, $2);", req.FromUserId, req.ToUserId)
// 	if dbErr != nil {
// 		return &api.ActionResponse{
// 			Success: false,
// 			Message: fmt.Sprintf("During friend adding the problem appeared: \"%v\"", dbErr),
// 		}, dbErr
// 	}
// 	message = fmt.Sprintf(
// 		"User with id '%d' Added user with id '%d'", req.FromUserId, req.ToUserId,
// 	)
// 	return &api.ActionResponse{
// 		Success: true,
// 		Message: message,
// 	}, nil
// }

// func (us *UserServ) ViewUserPage(ctx context.Context, req *api.ViewUserPageRequest) (*api.ViewUserPageResponse, error) {
// 	if errID := is_id_provided(req.UserId); errID != nil {
// 		return nil, errID
// 	} else if errID := is_id_provided(req.VisitedUserId); errID != nil {
// 		return nil, errID
// 	}

// 	var isBlockedU bool
// 	isBlockRow, dbErr := MyDatabase.Query("SELECT is_blocked FROM friends WHERE user_a_id = $1 AND user_b_id = $2 LIMIT 1;", req.UserId, req.VisitedUserId)
// 	if dbErr != nil {
// 		return nil, dbErr
// 	}

// 	for isBlockRow.Next() {
// 		err := isBlockRow.Scan(&isBlockedU)
// 		if err != nil {
// 			log.Println("Appeared error during iteration in query: ", err)
// 			break
// 		}
// 	}
// 	resultRow, dbErr := MyDatabase.Query(
// 		"SELECT users.id, users.first_name, users.last_name, users.email, users.username, phones.phone FROM users LEFT JOIN phones ON users.id = phones.owner_id WHERE users.id = $1;", req.VisitedUserId)
// 	if dbErr != nil {
// 		return nil, dbErr
// 	}
// 	defer resultRow.Close()
// 	user := []api.ViewUserPageResponse{}

// 	for resultRow.Next() {
// 		us := api.ViewUserPageResponse{}
// 		var phonee string

// 		err := resultRow.Scan(&us.UserId, &us.FirstName, &us.LastName, &us.Email, &us.Username, &phonee)
// 		if err != nil {
// 			log.Println("Appeared error during iteration: ", err)
// 			break
// 		}

// 		us.Phones = append(us.Phones, phonee)
// 		user = append(user, us)
// 	}
// 	resultedUser := api.ViewUserPageResponse{}
// 	if len(user) != 0 {
// 		resultedUser.UserId = user[0].UserId
// 		resultedUser.FirstName = user[0].FirstName
// 		resultedUser.LastName = user[0].LastName
// 		resultedUser.Email = user[0].Email
// 		resultedUser.Username = user[0].Username
// 		resultedUser.IsBlocked = isBlockedU
// 		for _, v := range user {
// 			resultedUser.Phones = append(resultedUser.Phones, v.Phones[0])
// 		}

// 		return &resultedUser, nil
// 	}

// 	return nil, status.Errorf(
// 		codes.NotFound,
// 		("The user WAS NOT FOUND with ID " + string(req.VisitedUserId)))
// }

// // UnimplementedAnimeServiceServer
// type AnimeServ struct {
// 	api.UnimplementedAnimeServiceServer
// }

// func (as *AnimeServ) HandleAnime(ctx context.Context, req *api.UserAnimeActionRequest) (*api.ActionResponse, error) {
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
// 	} else if validErr := is_id_provided(req.ActionId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"ACTION ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	}

// 	_, dbErr := MyDatabase.Exec("call user_anime_action($1, $2, $3);", req.UserId, req.AnimeId, req.ActionId)
// 	if dbErr != nil {
// 		return &api.ActionResponse{
// 			Success: false,
// 			Message: "Anime Handling failed",
// 		}, status.Errorf(codes.Canceled, dbErr.Error())
// 	}
// 	message = fmt.Sprintf(
// 		"You '%d' successfully managed to Handle ANIME '%d'", req.UserId, req.AnimeId,
// 	)
// 	return &api.ActionResponse{
// 		Success: true,
// 		Message: message,
// 	}, nil
// }

// func (as *AnimeServ) RemoveAnime(ctx context.Context, req *api.RemoveAnimeRequest) (*api.AnimeResponse, error) {
// 	if validErr := is_id_provided(req.UserId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"USER ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	} else if validErr := is_id_provided(req.UserId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"USER ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	}

// 	_, dbErr := MyDatabase.Exec("call remove_anime_list($1, $2);", req.UserId, req.AnimeId)
// 	if dbErr != nil {
// 		return nil, status.Errorf(codes.Canceled, dbErr.Error())
// 	}
// 	message = fmt.Sprintf(
// 		"You '%d' successfully removed ANIME from your LIST '%d'", req.UserId, req.AnimeId,
// 	)
// 	return &api.AnimeResponse{
// 		Name:    "Need to customize",
// 		Success: true,
// 		Message: message,
// 	}, nil
// }

// func (as *AnimeServ) ViewAnime(ctx context.Context, req *api.ViewAnimeRequest) (*api.ViewAnimeResponse, error) {
// 	if validErr := is_id_provided(req.AnimeId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"ANIME ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	} else if validErr := is_id_provided(req.UserId); validErr != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			"ANIME ID ERROR: %s",
// 			validErr.Error(),
// 		)
// 	}
// 	resultRow, dbErr := MyDatabase.Query(
// 		"SELECT animes.id, animes.name, animes.description, animes.release_year, anime_groups.name, animes.rating, animes.views_number, studios.name, genres.name FROM animes LEFT JOIN anime_genres ON animes.id = anime_genres.anime_id LEFT JOIN genres ON genres.id = anime_genres.genre_id LEFT JOIN anime_groups ON animes.group_id = anime_groups.id LEFT JOIN studios ON animes.studio_id = studios.id WHERE animes.id = $1;", req.AnimeId)

// 	if dbErr != nil {
// 		return nil, dbErr
// 	}
// 	var releaseYear uint32 = 0
// 	var resultAnimeResponse api.ViewAnimeResponse = api.ViewAnimeResponse{}
// 	var animeGenres []string
// 	for resultRow.Next() {
// 		var curGenre string
// 		err := resultRow.Scan(
// 			&resultAnimeResponse.AnimeId,
// 			&resultAnimeResponse.Name,
// 			&resultAnimeResponse.Description,
// 			&releaseYear,
// 			&resultAnimeResponse.AnimeGroup,
// 			&resultAnimeResponse.Rating,
// 			&resultAnimeResponse.ViewsNumber,
// 			&resultAnimeResponse.StudioName,
// 			&curGenre,
// 		)
// 		if err != nil {
// 			log.Println("Appeared error during iteration in query: ", err)
// 			break
// 		}
// 		animeGenres = append(animeGenres, curGenre)
// 	}
// 	if releaseYear == uint32(zeroValue) {
// 		return nil, status.Errorf(codes.NotFound, "Anime with ID %d was not FOUND", req.AnimeId)
// 	}
// 	resultAnimeResponse.Genres = animeGenres
// 	resultAnimeResponse.ReleaseDate = string(releaseYear)

// 	return &resultAnimeResponse, nil
// }

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
