package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"proto/pkg/api"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	port          string = ":8080"
	lineDelimeter string = "\n---------------------------\n"
	zeroValue     int    = 0
)

var (
	CurrentUserID uint32   = 1
	AllApiOptions []string = []string{
		"1. Log in (not ready)",
		"2. Register (not ready)",
		"3. Block user",
		"4. Add user to friends",
		"5. View User Profile",
		"6. Handle Anime",
		"7. Exit",
	}
	AllAnimeOptions []string = []string{
		"1. Like Anime",
		"2. Add to 'Will watch later' list",
		"3. Add to 'Currently watching' list",
		"4. Add to 'Through away' list",
		"5. Add to 'Already watched' list",
		"6. Remove Anime from your list",
		"7. View Anime info",
		"8. Comment Anime",
		"9. Repply to User Comment",
		"10. Go back",
	}
	UserNumberChoice        int64
	ApiMaxPossibleOptions   int             = len(AllApiOptions)
	AnimeMaxPossibleOptions int             = len(AllAnimeOptions)
	Ctx                     context.Context = context.Background()
	UserServClient          api.UserServiceClient
	AnimeServClient         api.AnimeServiceClient
	CommentServClient       api.CommentServiceClient
	Scanner                 *bufio.Scanner = bufio.NewScanner(os.Stdin)
	ApiFunctions            []func()       = []func(){
		LogInClentHandler,
		RegisterClientHandler,
		BlockUserClientHandler,
		AddUserFriendClientHandler,
		ViewUserProfileClientHandler,
		startAnimeOptions,
	}
	AnimeFunctions []func() = []func(){
		LikeAnimeClientHandler,
		AddWatchLaterClientHanler,
		AddCurrentlyWatchingClientHandler,
		AddThroughAwayClientHandler,
		AddAlreadyWatchedClientHandler,
		RemoveAnimeListClientHandler,
		ViewAnimeClientHandler,
		CommentAnimeClientHandler,
	}
)

func LogInClentHandler() {
	fmt.Println("Sorry, but it is not ready yet. We will correct it in REST")
}

func RegisterClientHandler() {
	fmt.Println("Sorry, but it is not ready yet. We will correct it in REST")
}

func BlockUserClientHandler() {
	fmt.Println("Please, type the ID of USER that you want to block: ")
	Scanner.Scan()
	friendID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionRespnse, errService := UserServClient.BlockUser(
		Ctx,
		&api.UserAddBlockUserRequest{
			FromUserId: CurrentUserID,
			ToUserId:   uint32(friendID),
		},
	)
	if errService != nil {
		log.Println("During process of blocking error appeared: ", errService)
	} else {
		fmt.Println(actionRespnse.Message)
	}
}

func AddUserFriendClientHandler() {
	fmt.Println("Please, type the ID of USER that you want to add to your friends: ")
	Scanner.Scan()
	friendID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionResponse, errService := UserServClient.AddUser(
		Ctx,
		&api.UserAddBlockUserRequest{
			FromUserId: CurrentUserID,
			ToUserId:   uint32(friendID),
		},
	)
	if errService != nil {
		log.Println("During process adding error appeared: ", actionResponse.Message)
	} else {
		fmt.Println(actionResponse.Message)
	}
}

func ViewUserProfileClientHandler() {
	fmt.Println("Please, type the ID of USER that you want to VIEW: ")
	Scanner.Scan()
	visitedID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	pageResponse, errService := UserServClient.ViewUserPage(
		Ctx,
		&api.ViewUserPageRequest{
			UserId:        CurrentUserID,
			VisitedUserId: uint32(visitedID),
		},
	)
	if errService != nil {
		log.Println("During process of page requesting error appeared: ", errService)
	} else {
		fmt.Println("User info: ")
		fmt.Println(
			"\tUser ID: ", pageResponse.UserId,
			"\n\tUser's firt name: ", pageResponse.FirstName,
			"\n\tUser's last name: ", pageResponse.LastName,
			"\n\tUser's email: ", pageResponse.Email,
			"\n\tUser's username: ", pageResponse.Username,
			"\n\tDid you block the user: ", pageResponse.IsBlocked,
			"\n\tUser's all phones: ", pageResponse.Phones,
		)
	}
}

func LikeAnimeClientHandler() {
	fmt.Println("Please, type the ID of Anime that you want to HANDLE: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionRespone, errService := AnimeServClient.HandleAnime(
		Ctx,
		&api.UserAnimeActionRequest{
			UserId:   CurrentUserID,
			AnimeId:  uint32(animeID),
			ActionId: 5,
		},
	)
	if errService != nil {
		log.Println("During process of liking error appeared: ", errService)
	} else {
		fmt.Println(actionRespone.Message)
	}
}

func AddWatchLaterClientHanler() {
	fmt.Println("Please, type the ID of Anime that you want to HANDLE: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionRespone, errService := AnimeServClient.HandleAnime(
		Ctx,
		&api.UserAnimeActionRequest{
			UserId:   CurrentUserID,
			AnimeId:  uint32(animeID),
			ActionId: 1,
		},
	)
	if errService != nil {
		log.Println("During process of adding anime to 'Watch later' list error appeared: ", errService)
	} else {
		fmt.Println(actionRespone.Message)
	}
}

func AddCurrentlyWatchingClientHandler() {
	fmt.Println("Please, type the ID of Anime that you want to add to currently watching list: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionRespone, errService := AnimeServClient.HandleAnime(
		Ctx,
		&api.UserAnimeActionRequest{
			UserId:   CurrentUserID,
			AnimeId:  uint32(animeID),
			ActionId: 2,
		},
	)
	if errService != nil {
		log.Println("During process of adding anime to 'Currently watching' list error appeared: ", errService)
	} else {
		fmt.Println(actionRespone.Message)
	}
}

func AddThroughAwayClientHandler() {
	fmt.Println("Please, type the ID of Anime that you want to Trhough away: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionRespone, errService := AnimeServClient.HandleAnime(
		Ctx,
		&api.UserAnimeActionRequest{
			UserId:   CurrentUserID,
			AnimeId:  uint32(animeID),
			ActionId: 3,
		},
	)
	if errService != nil {
		log.Println("During process of adding anime to 'Through Away' list error appeared: ", errService)
	} else {
		fmt.Println(actionRespone.Message)
	}
}

func AddAlreadyWatchedClientHandler() {
	fmt.Println("Please, type the ID of Anime that you want to Watch later: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionRespone, errService := AnimeServClient.HandleAnime(
		Ctx,
		&api.UserAnimeActionRequest{
			UserId:   CurrentUserID,
			AnimeId:  uint32(animeID),
			ActionId: 3,
		},
	)
	if errService != nil {
		log.Println("During process of adding anime to 'Already watched' list error appeared: ", errService)
	} else {
		fmt.Println(actionRespone.Message)
	}
}

func RemoveAnimeListClientHandler() {
	fmt.Println("Please, type the ID of Anime that you want to remove from your list: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	animeResponse, errService := AnimeServClient.RemoveAnime(
		Ctx,
		&api.RemoveAnimeRequest{
			UserId:  CurrentUserID,
			AnimeId: uint32(animeID),
		},
	)
	if errService != nil {
		log.Println("During process of removing anime from your list error appeared: ", errService)
	} else {
		fmt.Println(animeResponse.Message)
	}
}

func ViewAnimeClientHandler() {
	fmt.Println("Please, type the ID of Anime that you want to View: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	viewAnimeResponse, errService := AnimeServClient.ViewAnime(
		Ctx,
		&api.ViewAnimeRequest{
			UserId:  CurrentUserID,
			AnimeId: uint32(animeID),
		},
	)
	if errService != nil {
		log.Println("During process of removing anime from your list error appeared: ", errService)
	} else {
		fmt.Println("Anime info: ")
		fmt.Println(
			"\tAnime ID: ", viewAnimeResponse.AnimeId,
			"\n\tAnime's name: ", viewAnimeResponse.Name,
			"\n\tAnime's description: ", viewAnimeResponse.Description,
			"\n\tAnime's anime group: ", viewAnimeResponse.AnimeGroup,
			"\n\tAnime's release year: ", viewAnimeResponse.ReleaseDate,
			"\n\tAnime's rating (out of 10): ", viewAnimeResponse.Rating,
			"\n\tAnime's views number: ", viewAnimeResponse.ViewsNumber,
			"\n\tAnime's studio release name: ", viewAnimeResponse.StudioName,
			"\n\tAnime's genres: ", viewAnimeResponse.Genres,
		)
	}
}

func CommentAnimeClientHandler() {
	fmt.Println("Please, type the ID of Anime that you want to Comment: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	fmt.Println("Please, type the text of your comment: ")
	Scanner.Scan()
	var commentText string = Scanner.Text()
	actionResponse, errService := CommentServClient.CommentAnime(
		Ctx,
		&api.CommentAnimeRequest{
			UserId:  CurrentUserID,
			AnimeId: uint32(animeID),
			Content: commentText,
		},
	)
	if errService != nil {
		log.Println("During process of leaving comment error appeared: ", errService)
	} else {
		fmt.Println(actionResponse.Message)
	}
}

func RepplyUserCommentClientHandler() {
	fmt.Println("Please, type the ID of Anime that you want to Comment: ")
	Scanner.Scan()
	animeID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}

	fmt.Println("Please, type the text of your comment: ")
	Scanner.Scan()
	var commentText string = Scanner.Text()

	fmt.Println("Please, type the ID of comment that you want to reply: ")
	Scanner.Scan()
	commentRepplyID, err := strconv.ParseInt(Scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionResponse, errService := CommentServClient.ReplyUserCommentAnime(
		Ctx,
		&api.ReplyUserCommentAnimeRequest{
			UserId:           CurrentUserID,
			AnimeId:          uint32(animeID),
			Content:          commentText,
			RepliedCommentId: uint32(commentRepplyID),
		},
	)
	if errService != nil {
		log.Println("During process of leaving comment error appeared: ", errService)
	} else {
		fmt.Println(actionResponse.Message)
	}
}

func ShowPreviewMessage() {
	fmt.Println("Please, choose one of the possible provided actions:")
}

func GetApiOptions() string {
	var stringApiOptions string
	for _, v := range AllApiOptions {
		stringApiOptions += v + "\n"
	}
	return stringApiOptions
}

func GetAnimeOptions() string {
	var stringAnimeOptions string
	for _, v := range AllAnimeOptions {
		stringAnimeOptions += v + "\n"
	}
	return stringAnimeOptions
}

// func startConnection() {
// 	fmt.Println("Making an attempt to set the connection with server")
// 	var connectionStartTime time.Time = time.Now()
// 	conn, err := grpc.Dial("localhost "+port, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("Could not connect to %s: %v", port, err)
// 	}
// 	log.Printf("Connected in %d microsec", time.Now().Sub(connectionStartTime).Microseconds())
// 	UserServClient = api.NewUserServiceClient(conn)
// 	AnimeServClient = api.NewAnimeServiceClient(conn)
// 	CommentServClient = api.NewCommentServiceClient(conn)
// }

func startAnimeOptions() {
	var animeOptions = GetAnimeOptions()
	for {
		fmt.Println(lineDelimeter)
		ShowPreviewMessage()
		fmt.Println(animeOptions)
		fmt.Println("Your choice(type only number from list): ")
		Scanner.Scan()
		userNumberChoice, err := strconv.ParseInt(Scanner.Text(), 10, 0)
		if err != nil {
			log.Println("You made a mistake as an input: ", err)
		} else {
			if userNumberChoice <= int64(zeroValue) || userNumberChoice > int64(AnimeMaxPossibleOptions) {
				log.Printf("The only opions can be between 1 and %d", AnimeMaxPossibleOptions)
			} else if userNumberChoice == int64(AnimeMaxPossibleOptions) {
				break
			} else {
				AnimeFunctions[userNumberChoice-1]()
			}
		}
	}
}

func startAnimeApi() {
	var apiOptions = GetApiOptions()
	for {
		fmt.Println(lineDelimeter)
		ShowPreviewMessage()
		fmt.Println(apiOptions)
		fmt.Println("Your choice(type only number from list): ")
		Scanner.Scan()
		userNumberChoice, err := strconv.ParseInt(Scanner.Text(), 10, 0)
		if err != nil {
			log.Println("You made a mistake as an input: ", err)
		} else {
			if userNumberChoice <= int64(zeroValue) || userNumberChoice > int64(ApiMaxPossibleOptions) {
				log.Printf("The only opions can be between 1 and %d", ApiMaxPossibleOptions)
			} else if userNumberChoice == int64(ApiMaxPossibleOptions) {
				break
			} else {
				ApiFunctions[userNumberChoice-1]()
			}
		}
	}
}

func main() {
	fmt.Println("Hello!")
	fmt.Println("Making an attempt to set the connection with server")
	var connectionStartTime time.Time = time.Now()
	conn, err := grpc.Dial("localhost "+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", port, err)
	}
	log.Printf("Connected in %d microsec", time.Now().Sub(connectionStartTime).Microseconds())
	UserServClient = api.NewUserServiceClient(conn)
	// AnimeServClient = api.NewAnimeServiceClient(conn)
	// CommentServClient = api.NewCommentServiceClient(conn)
	// startConnection()
	fmt.Print(
		"Welcome to Anime Api prepared by Maratuly Temirbolat.",
	)
	startAnimeApi()
	fmt.Println("Thank for being with us! See you!")
	// var userServStartTime time.Time = time.Now()
	// actionResponse, err := UserServClient.BlockUser(ctx, &api.UserAddBlockUserRequest{
	// 	FromUserId: 1,
	// 	ToUserId:   2,
	// })
	// if actionResponse != nil {
	// 	log.Printf("The result of actionRequest is message: %v. success state: %v", actionResponse.Message, actionResponse.Success)
	// }
	// if err != nil {
	// 	log.Fatalf("Could not block user: %v", err)
	// }
	// log.Printf("Blocked user in %d microsec", time.Now().Sub(userServStartTime).Microseconds())
}
