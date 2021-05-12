package main

import (
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

	// 跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")                            //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization") //header的类型
	w.Header().Set("content-type", "application/json")                            //返回数据格式是json

	//check session
	//ValidateUserSession(r)

	m.r.ServeHTTP(w, r)
}

// listen -> RegisterHandelers -> handlers

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	err := http.ListenAndServe(":7788", mh)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
