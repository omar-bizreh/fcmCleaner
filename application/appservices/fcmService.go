package appservices

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/omar-bizreh/fcmCleaner/data/models"
)

// FCMService Handles requests for FCM
type FCMService struct{}

// NewFCMService Creates new instance of FCMService
func NewFCMService() FCMService {
	service := new(FCMService)
	return *service
}

// checkTokensValidity Checks if tokens are valid or not
func (service *FCMService) checkTokensValidity(tokens []models.PushNotificationToken) (*models.PushResponse, error) {
	reqBody := new(models.PushRequest)
	reqBody.DryRun = true
	reqBody.PushTokens = make([]string, len(tokens))
	for i, token := range tokens {
		reqBody.PushTokens[i] = token.Token
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(jsonBody))

	client := new(http.Client)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "key="+os.Getenv("fcmAuthKey"))

	resp, _ := client.Do(req)

	pushResponse := new(models.PushResponse)
	responseBytes, _ := ioutil.ReadAll(resp.Body)

	parseError := json.Unmarshal(responseBytes, &pushResponse)
	if parseError != nil {
		return nil, parseError
	}
	return pushResponse, nil
}

// extractNonValidTokens extracts non valid tokens and return them in array
func (service *FCMService) extractNonValidTokens(validatedTokens []models.PushNotificationToken, pushResponse models.PushResponse) []models.PushNotificationToken {
	tokensToRemove := make([]models.PushNotificationToken, 0)

	for i, result := range pushResponse.Results {
		if len(result.Error) > 0 {
			tokensToRemove = append(tokensToRemove, validatedTokens[i])
		}
	}

	return tokensToRemove
}

// ValidateAndGetNonValidTokens validates and returns non valid tokens
func (service *FCMService) ValidateAndGetNonValidTokens(tokens []models.PushNotificationToken) ([]models.PushNotificationToken, error) {
	validationResult, validationError := service.checkTokensValidity(tokens)
	if validationError != nil {
		return make([]models.PushNotificationToken, 0), validationError
	}

	nonValidTokens := service.extractNonValidTokens(tokens, *validationResult)
	return nonValidTokens, nil
}

func (service *FCMService) GetChunksForPageSize(size int, numOfTokens int) int {
	numOfChunks := numOfTokens / size
	remainder := numOfTokens % size
	if remainder > 0 {
		numOfChunks += 1
	}
	return numOfChunks
}
