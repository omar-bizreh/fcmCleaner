package router

import "github.com/omar-bizreh/fcmCleaner/application/router/home"

// AppRouter handles init all routes in service
type AppRouter struct{}

// InitRoutes init all routes in service
func (router *AppRouter) InitRoutes() {
	homeRouter := new(home.HomeRouter)
	homeRouter.Init()
}
