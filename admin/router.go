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

	// user
	router.POST("/admin/register", controllers.AdminUserRegister)
	router.POST("/admin/update/:id", controllers.AdminUpdateUser)
	router.POST("/admin/delete/:id", controllers.AdminUserDelete)
	router.GET("/admin/list", controllers.AdminUserList)
	router.POST("/admin/updateStatus/:id", controllers.AdminUpdateUserStatus)

	// role
	router.POST("/role/create", controllers.RoleCreate)
	router.POST("/role/update/:rid", controllers.RoleUpdate)
	router.POST("/role/delete", controllers.RoleDelete)
	router.GET("/admin/role/:id", controllers.AdminRoles)
	router.GET("/role/listAll", controllers.AdminRoleListAll)
	router.GET("/role/list", controllers.AdminRoleList)
	router.POST("/role/updateStatus/:rid", controllers.RoleUpdateStatus)
	router.POST("/admin/role/update", controllers.AdminRoleUpdate)
	router.GET("/role/listResource/:rid", controllers.RoleResourceByRoleId)
	router.POST("/role/allocResource", controllers.RoleAllocResource)
	router.GET("/role/listMenu/:rid", controllers.RoleListMenuByRid)
	router.POST("/role/allocMenu", controllers.RoleAllocMenu)

	// menu
	router.POST("/menu/create", controllers.MenuCreate)
	router.POST("/menu/update/:mid", controllers.MenuUpdateByMid)
	router.POST("/menu/delete/:mid", controllers.MenuDeleteByMid)
	router.GET("/menu/getMenu/:id", controllers.GetMenuInfoById)
	router.GET("/menu/list/:pid", controllers.GetMenuListByPid)
	router.GET("/menu/treeList", controllers.MenuTreeList)
	router.POST("/menu/updateHidden/:mid", controllers.MenuUpdateHidden)

	// resource
	router.GET("/resource/listAll", controllers.ResourceListAll)
	router.GET("/resource/list", controllers.ResourceList)

	router.POST("/resource/create", controllers.ResourceCreate)
	router.POST("/resource/update/:id", controllers.ResourUpdateById)
	router.POST("/resource/delete/:id", controllers.ResourDeleteById)

	// resourceCategory
	router.GET("/resourceCategory/listAll", controllers.ResourceCategoryListAll)
	router.POST("/resourceCategory/create", controllers.ResourceCategoryCreate)
	router.POST("/resourceCategory/delete/:id", controllers.ResourceCategoryDeleteById)
	router.POST("/resourceCategory/update/:id", controllers.ResourceCategoryUpdateById)

	// game
	router.GET("/game/list", controllers.GetGameList)
	router.GET("/game/gameTag", controllers.GetGameTag)
	router.GET("/game/gameTagByGameId", controllers.GetGameTagByGameId)

	return router
}
