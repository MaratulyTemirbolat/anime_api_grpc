// package main

// import (
// 	"fmt"
// 	"log"
// 	"proto/pkg/api"
// 	"strconv"
// )

// func LogInClentHandler() {
// 	fmt.Println("Sorry, but it is not ready yet. We will correct it in REST")
// }

// func RegisterClientHandler() {
// 	fmt.Println("Sorry, but it is not ready yet. We will correct it in REST")
// }

// func BlockUserClientHandler() {
// 	fmt.Println("Please, type the ID of USER that you want to block: ")
// 	Scanner.Scan()
// 	friendID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
// 	if err != nil {
// 		log.Println("You made a mistake as an input: ", err)
// 	}
// 	actionRespnse, errService := UserServClient.BlockUser(
// 		Ctx,
// 		&api.UserAddBlockUserRequest{
// 			FromUserId: CurrentUserID,
// 			ToUserId:   uint32(friendID),
// 		},
// 	)
// 	if errService != nil {
// 		log.Println("During process of blocking error appeared: ", errService)
// 	} else {
// 		fmt.Println(actionRespnse.Message)
// 	}
// }

// func AddUserFriendClientHandler() {
// 	fmt.Println("Please, type the ID of USER that you want to add to your friends: ")
// 	Scanner.Scan()
// 	friendID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
// 	if err != nil {
// 		log.Println("You made a mistake as an input: ", err)
// 	}
// 	actionResponse, errService := UserServClient.AddUser(
// 		Ctx,
// 		&api.UserAddBlockUserRequest{
// 			FromUserId: CurrentUserID,
// 			ToUserId:   uint32(friendID),
// 		},
// 	)
// 	if errService != nil {
// 		log.Println("During process adding error appeared: ", actionResponse.Message)
// 	} else {
// 		fmt.Println(actionResponse.Message)
// 	}
// }

// func ViewUserProfileClientHandler() {
// 	fmt.Println("Please, type the ID of USER that you want to VIEW: ")
// 	Scanner.Scan()
// 	visitedID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
// 	if err != nil {
// 		log.Println("You made a mistake as an input: ", err)
// 	}
// 	pageResponse, errService := UserServClient.ViewUserPage(
// 		Ctx,
// 		&api.ViewUserPageRequest{
// 			UserId:        CurrentUserID,
// 			VisitedUserId: uint32(visitedID),
// 		},
// 	)
// 	if errService != nil {
// 		log.Println("During process of page requesting error appeared: ", errService)
// 	} else {
// 		fmt.Println("User info: ")
// 		fmt.Println(
// 			"\tUser ID: ", pageResponse.UserId,
// 			"\n\tUser's firt name: ", pageResponse.FirstName,
// 			"\n\tUser's last name: ", pageResponse.LastName,
// 			"\n\tUser's email: ", pageResponse.Email,
// 			"\n\tUser's username: ", pageResponse.Username,
// 			"\n\tDid you block the user: ", pageResponse.IsBlocked,
// 			"\n\tUser's all phones: ", pageResponse.Phones,
// 		)
// 	}
// }
