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
	currentUserID uint32 = 1
	ctx           context.Context
	allApiOptions []string = []string{
		"1. Log in (not ready)",
		"2. Register (not ready)",
		"3. Block user",
		"4. Add user to friends",
		"5. View User Profile",
		"6. Handle Anime",
		"7. Exit",
	}
	allAnimeOptions []string = []string{
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
	apiMaxPossibleOptions   int            = len(allApiOptions)
	animeMaxPossibleOptions int            = len(allAnimeOptions)
	scanner                 *bufio.Scanner = bufio.NewScanner(os.Stdin)
	userServClient          api.UserServiceClient
	animeServClient         api.AnimeServiceClient
	commentServClient       api.CommentServiceClient
	apiFunctions            []func() = []func(){
		logInClentHandler,
		registerClientHandler,
		blockUserClientHandler,
		addUserFriendClientHandler,
		viewUserProfileClientHandler,
		startAnimeOptions,
	}
	animeFunctions []func() = []func(){
		likeAnimeClientHandler,
		addWatchLaterClientHanler,
		addCurrentlyWatchingClientHandler,
		addThroughAwayClientHandler,
		addAlreadyWatchedClientHandler,
		removeAnimeListClientHandler,
		viewAnimeClientHandler,
		commentAnimeClientHandler,
	}
)

func showPreviewMessage() {
	fmt.Println("Please, choose one of the possible provided actions:")
}

func getAnimeOptions() string {
	var stringAnimeOptions string
	for _, v := range allAnimeOptions {
		stringAnimeOptions += v + "\n"
	}
	return stringAnimeOptions
}

func getApiOptions() string {
	var stringApiOptions string
	for _, v := range allApiOptions {
		stringApiOptions += v + "\n"
	}
	return stringApiOptions
}

func startAnimeOptions() {
	var animeOptions = getAnimeOptions()
	for {
		fmt.Println(lineDelimeter)
		showPreviewMessage()
		fmt.Print(animeOptions)
		fmt.Print("\nYour choice (type only number from list): ")
		scanner.Scan()
		userNumberChoice, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			log.Println("You made a mistake as an input: ", err)
		} else {
			if userNumberChoice <= int64(zeroValue) || userNumberChoice > int64(animeMaxPossibleOptions) {
				log.Printf("The only opions can be between 1 and %d", animeMaxPossibleOptions)
			} else if userNumberChoice == int64(animeMaxPossibleOptions) {
				break
			} else {
				// fmt.Println(userNumberChoice, allAnimeOptions[userNumberChoice-1])
				animeFunctions[userNumberChoice-1]()
			}
		}
	}
}

func startAnimeApi() {
	var apiOptions = getApiOptions()
	for {
		fmt.Println(lineDelimeter)
		showPreviewMessage()
		fmt.Println(apiOptions)
		fmt.Print("Your choice(type only number from list): ")
		scanner.Scan()
		fmt.Println()
		userNumberChoice, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			log.Println("You made a mistake as an input: ", err)
		} else {
			if userNumberChoice <= int64(zeroValue) || userNumberChoice > int64(apiMaxPossibleOptions) {
				log.Printf("The only opions can be between 1 and %d", apiMaxPossibleOptions)
			} else if userNumberChoice == int64(apiMaxPossibleOptions) {
				break
			} else {
				// fmt.Println(userNumberChoice, allApiOptions[userNumberChoice-1])
				apiFunctions[userNumberChoice-1]()
			}
		}
	}
}

func main() {
	ctx = context.Background()
	var connectionStartTime time.Time = time.Now()
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", port, err)
	}
	log.Printf("Connected in %d microsec", time.Now().Sub(connectionStartTime).Microseconds())
	userServClient = api.NewUserServiceClient(conn)
	animeServClient = api.NewAnimeServiceClient(conn)
	commentServClient = api.NewCommentServiceClient(conn)
	startAnimeApi()
	fmt.Println("Thank you! Good buy!")
	// var userServStartTime time.Time = time.Now()
	// actionResponse, err := userServClient.BlockUser(ctx, &api.UserAddBlockUserRequest{
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

func logInClentHandler() {
	fmt.Println("Sorry, but it is not ready yet. We will correct it in REST")
}

func registerClientHandler() {
	fmt.Println("Sorry, but it is not ready yet. We will correct it in REST")
}

func blockUserClientHandler() {
	fmt.Print("\nPlease, type the ID of USER that you want to block: ")
	scanner.Scan()
	friendID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	actionRespnse, errService := userServClient.BlockUser(
		ctx,
		&api.UserAddBlockUserRequest{
			FromUserId: currentUserID,
			ToUserId:   uint32(friendID),
		},
	)
	if errService != nil {
		log.Println("During process of blocking error appeared: ", errService)
	} else {
		fmt.Println(actionRespnse.Message)
	}
}

func addUserFriendClientHandler() {
	fmt.Print("Please, type the ID of USER that you want to add to your friends: ")
	scanner.Scan()
	friendID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	actionResponse, errService := userServClient.AddUser(
		ctx,
		&api.UserAddBlockUserRequest{
			FromUserId: currentUserID,
			ToUserId:   uint32(friendID),
		},
	)
	if errService != nil {
		log.Println("During process adding error appeared: ", actionResponse.Message)
	} else {
		fmt.Println(actionResponse.Message)
	}
}

func viewUserProfileClientHandler() {
	fmt.Print("\nPlease, type the ID of USER that you want to VIEW: ")
	scanner.Scan()
	visitedID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	pageResponse, errService := userServClient.ViewUserPage(
		ctx,
		&api.ViewUserPageRequest{
			UserId:        currentUserID,
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

func likeAnimeClientHandler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to LIKE: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	actionRespone, errService := animeServClient.HandleAnime(
		ctx,
		&api.UserAnimeActionRequest{
			UserId:   currentUserID,
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

func addWatchLaterClientHanler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to HANDLE: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	actionRespone, errService := animeServClient.HandleAnime(
		ctx,
		&api.UserAnimeActionRequest{
			UserId:   currentUserID,
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

func addCurrentlyWatchingClientHandler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to add to currently watching list: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	actionRespone, errService := animeServClient.HandleAnime(
		ctx,
		&api.UserAnimeActionRequest{
			UserId:   currentUserID,
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

func addThroughAwayClientHandler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to Trhough away: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	actionRespone, errService := animeServClient.HandleAnime(
		ctx,
		&api.UserAnimeActionRequest{
			UserId:   currentUserID,
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

func addAlreadyWatchedClientHandler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to Watch later: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	actionRespone, errService := animeServClient.HandleAnime(
		ctx,
		&api.UserAnimeActionRequest{
			UserId:   currentUserID,
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

func removeAnimeListClientHandler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to remove from your list: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
	}
	animeResponse, errService := animeServClient.RemoveAnime(
		ctx,
		&api.RemoveAnimeRequest{
			UserId:  currentUserID,
			AnimeId: uint32(animeID),
		},
	)
	if errService != nil {
		log.Println("During process of removing anime from your list error appeared: ", errService)
	} else {
		fmt.Println(animeResponse.Message)
	}
}

func viewAnimeClientHandler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to View: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	viewAnimeResponse, errService := animeServClient.ViewAnime(
		ctx,
		&api.ViewAnimeRequest{
			UserId:  currentUserID,
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

func commentAnimeClientHandler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to Comment: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	fmt.Println("Please, type the text of your comment: ")
	scanner.Scan()
	var commentText string = scanner.Text()
	actionResponse, errService := commentServClient.CommentAnime(
		ctx,
		&api.CommentAnimeRequest{
			UserId:  currentUserID,
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

func repplyUserCommentClientHandler() {
	fmt.Print("\nPlease, type the ID of Anime that you want to Comment: ")
	scanner.Scan()
	animeID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}

	fmt.Print("\nPlease, type the text of your comment: ")
	scanner.Scan()
	var commentText string = scanner.Text()

	fmt.Print("\nPlease, type the ID of comment that you want to reply: ")
	scanner.Scan()
	commentRepplyID, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		log.Println("You made a mistake as an input: ", err)
		return
	}
	actionResponse, errService := commentServClient.ReplyUserCommentAnime(
		ctx,
		&api.ReplyUserCommentAnimeRequest{
			UserId:           currentUserID,
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
