package controllers

import (
	"admin/api/dbops"
	"admin/api/defs"
	"admin/api/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// 获取 globle config 的轮播图
func GetConfigs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	content := p.ByName("content")
	if res, err := dbops.GetConfigs(content); err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		lng, _ := json.Marshal(res)
		utils.SendNormalResponse(w, string(lng), 200)
		return
	}
}

// 获取首页各种类的游戏
func GetTypeGames(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	num, _ := strconv.Atoi(p.ByName("num"))
	if res, err := dbops.GetTypeGames(num); err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		lng, _ := json.Marshal(res)
		utils.SendNormalResponse(w, string(lng), 200)
		return
	}
}
