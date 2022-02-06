package home

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/omar-bizreh/fcmCleaner/application/appservices"
	appservice "github.com/omar-bizreh/fcmCleaner/application/appservices"
	"github.com/omar-bizreh/fcmCleaner/data/repositories"
)

// HomeRouter contains home routes
type HomeRouter struct{}

func rootRotue(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request path: " + r.URL.Path)
	fmt.Fprintln(w, "Welcome to FCM Token Cleaner")
	strBuilder := strings.Builder{}
	strBuilder.WriteString("FCM Cleaning Service\n")
	tokenRepo := repositories.NewFCMTokensRepository()
	totalTokens := tokenRepo.GetTotalTokensCount()

	fmt.Fprintln(w, strings.Join([]string{"Total tokens found: ", strconv.Itoa(totalTokens)}, " "))
	strBuilder.WriteString(strings.Join([]string{"Total tokens found: ", strconv.Itoa(totalTokens), "\n"}, " "))
	if totalTokens == 0 {
		strBuilder.WriteString("No tokens found\n")
		fmt.Println("No tokens found")
		return
	}
	fcmService := appservice.NewFCMService()
	numOfChunks := fcmService.GetChunksForPageSize(1000, totalTokens)
	strBuilder.WriteString("Total chunks: " + strconv.Itoa(numOfChunks) + "\n")
	for i := 0; i < numOfChunks; i++ {
		tokenArray := tokenRepo.LoadTokens(1000, i)
		tokensToRemove, err := fcmService.ValidateAndGetNonValidTokens(tokenArray)
		if err != nil {
			strBuilder.WriteString("Unable to check validity: " + err.Error() + "\n")
			fmt.Fprintln(w, strings.Join([]string{"Unable to check validity:", err.Error()}, " "))
			return
		}
		strBuilder.WriteString(strings.Join([]string{"Tokens to remove: ", strconv.Itoa(len(tokensToRemove)), "\n"}, " "))
		fmt.Fprintln(w, strings.Join([]string{"Tokens to remove: ", strconv.Itoa(len(tokensToRemove))}, " "))

		removeError := tokenRepo.RemoveTokens(tokensToRemove)
		if removeError != nil {
			strBuilder.WriteString(strings.Join([]string{"Failed to remove tokens\n", removeError.Error()}, " "))
			fmt.Fprintln(w, strings.Join([]string{"Failed to remove tokens", removeError.Error()}, " "))
		}
	}
	strBuilder.WriteString("Service finished successfully\n")
	defer sendEmail(strBuilder.String())
}

func sendEmail(msg string) {
	fmt.Println("MESSAGE: " + msg)
	mailService := new(appservices.MailService)
	mailService.SendEmail(os.Getenv("recieverEmail"), "FCM Cleaning Service", msg)
}

// Init initailize home routes
func (router *HomeRouter) Init() {
	http.HandleFunc("/clean_tokens", rootRotue)
}
