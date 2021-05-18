package main

import (
	"api/admin/defs"
	"api/admin/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/wxnacy/wgo/arrays"
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

	var urlList []string
	urlList = append(urlList, "/admin/info")
	urlList = append(urlList, "/admin/logout")
	urlList = append(urlList, "/admin/list")
	urlList = append(urlList, "/admin/register")
	urlList = append(urlList, "/admin/delete/")
	urlList = append(urlList, "/admin/updateStatus/")
	urlList = append(urlList, "/admin/update/")
	urlList = append(urlList, "/role/listAll")
	urlList = append(urlList, "/admin/role/")
	urlList = append(urlList, "/admin/role/update")

	if arrays.ContainsString(urlList, r.RequestURI) > -1 {
		//check jwt-token
		if utils.ValidateJwtToken(r) {
			utils.SendErrorResponse(w, defs.ErrorJwtTokenValidateFaild)
			return
		}

	}
	m.r.ServeHTTP(w, r)
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	err := http.ListenAndServe(":7788", mh)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
