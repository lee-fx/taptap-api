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
	router.POST("/admin/updateStatus/:id", controllers.AdminUpdateUserStatus)
	router.POST("/admin/update/:id", controllers.AdminUpdateUser)

	// role curd
	router.GET("/role/listAll", controllers.AdminRoleListAll)
	router.GET("/admin/role/:id", controllers.AdminRoles)
	router.POST("/admin/role/update", controllers.AdminRoleUpdate)
	router.GET("/role/list", controllers.AdminRoleList)
	router.POST("/role/updateStatus/:rid", controllers.RoleUpdateStatus)
	router.POST("/role/delete", controllers.RoleDelete)
	router.POST("/role/create", controllers.RoleCreate)
	router.POST("/role/update/:rid", controllers.RoleUpdate)
	router.GET("/role/listResource/:rid", controllers.RoleResourceByRoleId)
	router.POST("/role/allocResource", controllers.RoleAllocResource)
	router.GET("/role/listMenu/:rid", controllers.RoleListMenuByRid)
	router.POST("/role/allocMenu", controllers.RoleAllocMenu)

	// menu curd
	router.GET("/menu/treeList", controllers.MenuTreeList)
	router.GET("/menu/list/:pid", controllers.GetMenuListByPid)
	router.POST("/menu/updateHidden/:mid", controllers.MenuUpdateHidden)
	router.POST("/menu/create", controllers.MenuCreate)
	router.POST("/menu/delete/:mid", controllers.MenuDeleteByMid)
	router.GET("/menu/getMenu/:id", controllers.GetMenuInfoById)
	router.POST("/menu/update/:mid", controllers.MenuUpdateByMid)

	// resource curd
	router.GET("/resource/listAll", controllers.ResourceListAll)

	// resourceCategory curd
	router.GET("/resourceCategory/listAll", controllers.ResourceCategoryListAll)

	return router
}
