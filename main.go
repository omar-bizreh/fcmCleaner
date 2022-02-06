package main

import (
	"log"
	"net/http"
	"os"

	"github.com/omar-bizreh/fcmCleaner/application/router"
)

func initRoutes() {
	appRouter := new(router.AppRouter)
	appRouter.InitRoutes()
}

func initService() {
	// envLoader := appservices.NewEnvLoader()
	// envLoader.Load()
	initRoutes()
}

func startServer() {
	initService()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), nil))
}

func main() {
	startServer()
}
