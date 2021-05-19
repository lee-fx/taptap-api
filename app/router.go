package main

import (
	"api/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	// games
	router.POST("/game/getAllGames/:type/:page/:to", controllers.GetAllGames)
	router.POST("/game/getGameInfoById/:id", controllers.GetGameInfoById)

	// home
	router.POST("/home/getConfigs/:content", controllers.GetConfigs)
	router.POST("/home/getTypeGames/:num", controllers.GetTypeGames)

	// recommend
	router.POST("/game/getRecommends/:type/:page/:to", controllers.GetRecommends)

	// mine

	router.POST("/user", controllers.CreateUser)
	router.POST("/user/:user_name", controllers.UserLogin)

	return router
}
