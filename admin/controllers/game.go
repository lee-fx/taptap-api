package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"api/admin/utils"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
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

// 获取所有标签
func GetGameTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := dbops.GetGameTag()
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

// 获取游戏标签
func GetGameTagByGameId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	params := r.URL.Query()
	gid, _ := strconv.Atoi(params["game_id"][0])

	res, err := dbops.GetGameTagByGameId(gid)
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

// 修改游戏标签
func GameTagUpdateByGameId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//id, _ := strconv.Atoi(p.ByName("id"))
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserLogin{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	fmt.Printf("%v", tagNames)
	//
	//res, err := dbops.GameTagUpdateByGameId(id, tagNames)
	//if err != nil {
	//	log.Printf("error: %v ", err)
	//	utils.SendErrorResponse(w, defs.ErrorInternalFaults)
	//	return
	//}
	//
	//resData := &defs.NormalResponse{
	//	Code:    200,
	//	Message: "查询成功",
	//	Data:    res,
	//}
	//
	//utils.SendNormalResponse(w, *resData, 200)
}
