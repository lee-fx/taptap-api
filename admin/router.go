package main

import (
	"api/admin/controllers"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {

	router := httprouter.New()

	// user login
	router.POST("/admin/login", controllers.AdminUserLogin)

	// user info
	router.GET("/admin/info", controllers.AdminUserInfo)

	// user logout verify user error:
	router.POST("/admin/logout", controllers.AdminUserLogout)

	//router.POST("/user", controllers.CreateUser)
	//router.POST("/user/:user_name", controllers.UserLogin)

	return router
}
