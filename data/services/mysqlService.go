package services

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLService handles calls to database
type MySQLService struct {
	IDBService
	DB            *sql.DB
	isInitialized bool
}

// NewMySQLService init new db service
func NewMySQLService() MySQLService {
	service := new(MySQLService)
	service.InitDBInstance()
	return *service
}

// InitDBInstance Gets a new db instance
func (dbService *MySQLService) InitDBInstance() error {
	db, err := sql.Open("mysql", os.Getenv("db_user")+":"+os.Getenv("db_pass")+"@tcp("+os.Getenv("db_ip")+")/"+os.Getenv("db_name")+"?parseTime=true")
	if err != nil {
		return err
	}
	dbService.DB = db
	dbService.isInitialized = true
	return nil
}

// CloseConnection Close connection to database
func (dbService *MySQLService) CloseConnection() error {
	if dbService.isInitialized {
		err := dbService.DB.Close()
		return err
	}
	return nil
}
