package repositories

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/omar-bizreh/fcmCleaner/data/models"
	"github.com/omar-bizreh/fcmCleaner/data/services"
)

// FCMTokensRepository fcm token repo
type FCMTokensRepository struct {
	dbService     services.MySQLService
	isInitialized bool
}

// NewFCMTokensRepository create new instance of repo
func NewFCMTokensRepository() FCMTokensRepository {
	repo := new(FCMTokensRepository)
	repo.dbService = services.NewMySQLService()
	repo.isInitialized = true
	return *repo
}

// LoadTokens load tokens from database
func (repo *FCMTokensRepository) LoadTokens(pageSize int, pageIndex int) []models.PushNotificationToken {
	tokenArray := make([]models.PushNotificationToken, 0)
	statement, err := repo.dbService.DB.Prepare("SELECT Id, Token FROM PushNotificationTokens LIMIT ?,?")
	if err != nil {
		return make([]models.PushNotificationToken, 0)
	}

	query, statementErr := statement.Query(pageIndex*pageSize, pageSize)
	if statementErr != nil {
		return tokenArray
	}
	defer query.Close()
	rowIndex := 0
	for query.Next() {
		token := new(models.PushNotificationToken)
		query.Scan(&token.ID, &token.Token)
		tokenArray = append(tokenArray, *token)
		rowIndex++
	}
	return tokenArray
}

// GetTotalTokensCount Gets total number of available tokens
func (repo *FCMTokensRepository) GetTotalTokensCount() int {
	var counter int
	repo.dbService.DB.QueryRow("SELECT count(*) FROM PushNotificationTokens").Scan(&counter)
	return counter
}

// RemoveTokens remove tokens from database
func (repo *FCMTokensRepository) RemoveTokens(tokenArray []models.PushNotificationToken) error {
	if len(tokenArray) == 0 {
		fmt.Println("Nothing to remove")
		return nil
	}
	tokenIDsArray := make([]interface{}, len(tokenArray))
	for i, x := range tokenArray {
		tokenIDsArray[i] = x.ID
	}
	result, stmtErr := repo.dbService.DB.Exec(`DELETE FROM PushNotificationTokens WHERE Id IN (?`+strings.Repeat(",?", len(tokenIDsArray)-1)+`)`, tokenIDsArray...)

	if stmtErr != nil {
		return stmtErr
	}
	numRows, numRowsErr := result.RowsAffected()
	if numRowsErr == nil {
		rowsRemoved := strconv.FormatInt(numRows, 10)
		fmt.Println(strings.Join([]string{"Removed", rowsRemoved, "rows"}, " "))
	}
	return nil
}
