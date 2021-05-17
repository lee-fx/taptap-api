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

	// user logout
	router.POST("/admin/logout", controllers.AdminUserLogout)

	// user curd
	router.GET("/admin/list", controllers.AdminUserList)
	router.POST("/admin/register", controllers.AdminUserRegister)
	router.POST("/admin/delete/:id", controllers.AdminUserDelete)

	// role curd
	router.GET("/role/listAll", controllers.AdminRoleList)

	// menu curd

	// resource curd

	return router
}
