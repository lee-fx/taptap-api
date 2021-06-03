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
	gameName := params.Get("gameName")
	cid, _ := strconv.Atoi(params.Get("companyId"))

	res, err := dbops.GetGameList(page, to, gameName, cid)
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
	id, _ := strconv.Atoi(p.ByName("id"))
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.TagNames{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	err := dbops.GameTagUpdateByGameId(id, ubody)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "修改成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 获取公司列表
func GetCompanyList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取请求参数
	params := r.URL.Query()
	page, _ := strconv.Atoi(params["pageNum"][0])
	to, _ := strconv.Atoi(params["pageSize"][0])

	res, err := dbops.GetCompanyList(page, to)
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

// 修改游戏状态
func GameUpdateStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取请求参数
	params := r.URL.Query()
	ids := params["ids"][0]
	gameStatus, _ := strconv.Atoi(params["gameStatus"][0])

	err := dbops.GameUpdateStatus(gameStatus, ids)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 上传icon
func GameUploadIcon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	GAME_ICON := "./static/image/icon/"
	MAX_UPLOAD_SIZE := 1024 * 1024 * 5

	r.Body = http.MaxBytesReader(w, r.Body, int64(MAX_UPLOAD_SIZE))
	if err := r.ParseMultipartForm(int64(MAX_UPLOAD_SIZE)); err != nil {
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Printf("Read file error: %v", err)
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	gid, _ := strconv.Atoi(p.ByName("id"))

	file_url := GAME_ICON + handler.Filename
	fmt.Printf("%s", file_url)
	err = ioutil.WriteFile(file_url, data, 0666)

	if err != nil {
		log.Printf("Write file error: %v", err)
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	// 如果存在gid则修改图片地址
	if gid != 0 {
		// 删除原图片
		// 修改图片地址
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    file_url,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 上传icon
func GameUploadApk(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	GAME_ICON := "./static/pkg/apk/"
	MAX_UPLOAD_SIZE := 1024 * 1024 * 200

	r.Body = http.MaxBytesReader(w, r.Body, int64(MAX_UPLOAD_SIZE))
	if err := r.ParseMultipartForm(int64(MAX_UPLOAD_SIZE)); err != nil {
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Printf("Read file error: %v", err)
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	gid, _ := strconv.Atoi(p.ByName("id"))

	file_url := GAME_ICON + handler.Filename
	fmt.Printf("%s", file_url)
	err = ioutil.WriteFile(file_url, data, 0666)

	if err != nil {
		log.Printf("Write file error: %v", err)
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	// 如果存在gid则修改图片地址
	if gid != 0 {
		// 删除原apk
		// 修改apk地址
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    file_url,
	}

	utils.SendNormalResponse(w, *resData, 200)
}
