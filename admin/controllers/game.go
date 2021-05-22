package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"api/admin/utils"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

// 获取游戏列表
func GetGameList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取请求参数
	params := r.URL.Query()
	page, _ := strconv.Atoi(params["pageNum"][0])
	to, _ := strconv.Atoi(params["pageSize"][0])

	// 查询参数
	name := params.Get("gameName")
	//url := params.Get("urlKeyword")
	//id, _ := strconv.Atoi(params.Get("categoryId"))

	res, err := dbops.GetGameList(page, to, name)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "查询成功",
		Data:    res,
	}

	utils.SendNormalResponse(w, *resData, 200)
}
