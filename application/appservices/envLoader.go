package appservices

import (
	"log"

	"github.com/joho/godotenv"
)

// EnvLoader loads environment variable
type EnvLoader struct{}

// Load environment variables
func (env *EnvLoader) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func NewEnvLoader() EnvLoader {
	loader := new(EnvLoader)
	return *loader
}
