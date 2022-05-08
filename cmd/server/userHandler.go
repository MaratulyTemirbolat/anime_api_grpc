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
