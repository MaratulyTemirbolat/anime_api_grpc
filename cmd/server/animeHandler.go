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

