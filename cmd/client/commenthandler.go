// package main

// import (
// 	"fmt"
// 	"log"
// 	"proto/pkg/api"
// 	"strconv"
// )

// func CommentAnimeClientHandler() {
// 	fmt.Println("Please, type the ID of Anime that you want to Comment: ")
// 	Scanner.Scan()
// 	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
// 	if err != nil {
// 		log.Println("You made a mistake as an input: ", err)
// 	}
// 	fmt.Println("Please, type the text of your comment: ")
// 	Scanner.Scan()
// 	var commentText string = Scanner.Text()
// 	actionResponse, errService := CommentServClient.CommentAnime(
// 		Ctx,
// 		&api.CommentAnimeRequest{
// 			UserId:  CurrentUserID,
// 			AnimeId: uint32(animeID),
// 			Content: commentText,
// 		},
// 	)
// 	if errService != nil {
// 		log.Println("During process of leaving comment error appeared: ", errService)
// 	} else {
// 		fmt.Println(actionResponse.Message)
// 	}
// }

// func RepplyUserCommentClientHandler() {
// 	fmt.Println("Please, type the ID of Anime that you want to Comment: ")
// 	Scanner.Scan()
// 	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
// 	if err != nil {
// 		log.Println("You made a mistake as an input: ", err)
// 	}

// 	fmt.Println("Please, type the text of your comment: ")
// 	Scanner.Scan()
// 	var commentText string = Scanner.Text()

// 	fmt.Println("Please, type the ID of comment that you want to reply: ")
// 	Scanner.Scan()
// 	commentRepplyID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
// 	if err != nil {
// 		log.Println("You made a mistake as an input: ", err)
// 	}
// 	actionResponse, errService := CommentServClient.ReplyUserCommentAnime(
// 		Ctx,
// 		&api.ReplyUserCommentAnimeRequest{
// 			UserId:           CurrentUserID,
// 			AnimeId:          uint32(animeID),
// 			Content:          commentText,
// 			RepliedCommentId: uint32(commentRepplyID),
// 		},
// 	)
// 	if errService != nil {
// 		log.Println("During process of leaving comment error appeared: ", errService)
// 	} else {
// 		fmt.Println(actionResponse.Message)
// 	}
// }
