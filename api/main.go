package main

import (
	"admin/api/controllers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	// check session
	//ValidateUserSession(r)

	m.r.ServeHTTP(w, r)
}

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

// listen -> RegisterHandelers -> handlers

func main() {
	r := RegisterHandlers()
	//mh := NewMiddleWareHandler(r)
	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// handler -> validation{1.request, 2.user} -> business logic -> response.
// 1.data model 2.error handling
