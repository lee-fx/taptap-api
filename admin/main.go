package main

import (
	"github.com/julienschmidt/httprouter"
	"api/admin/controllers"
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

	// 跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	//check session
	//ValidateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	// user
	router.POST("/admin/login", controllers.AdminUserLogin)
	//router.POST("/user", controllers.CreateUser)
	//router.POST("/user/:user_name", controllers.UserLogin)

	return router
}

// listen -> RegisterHandelers -> handlers

func main() {
	r := RegisterHandlers()
	//mh := NewMiddleWareHandler(r)
	err := http.ListenAndServe(":7788", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
