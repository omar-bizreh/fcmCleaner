package services

// IDBService service methods
type IDBService interface {
	InitDatabase() error
	CloseConnection() error
}
