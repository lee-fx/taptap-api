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
	"strings"
	"api/admin/config"
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

// 游戏删除
func GameDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取请求参数
	params := r.URL.Query()
	ids := params["ids"][0]

	err := dbops.GameDeleteByIds(ids)
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
	GAME_ICON := config.GetGameIcon()
	MAX_UPLOAD_SIZE := config.GetMaxUploadImageSize()

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

	//gid, _ := strconv.Atoi(p.ByName("id"))

	rand_num := utils.GetRandNumByNumber(100)

	file_url := GAME_ICON + strconv.Itoa(int(rand_num)) +  handler.Filename
	err = ioutil.WriteFile(file_url, data, 0666)

	file_tmp := strings.TrimPrefix(file_url, ".")

	res_file_url := config.GetFileServerDirUrl() + file_tmp
	if err != nil {
		log.Printf("Write file error: %v", err)
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}


	resObj := defs.FileUpload{
		Name: handler.Filename,
		Url:  res_file_url,
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    resObj,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 上传icon
func GameUploadApk(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	GAME_FILE_APK := config.GetGameFileApk()
	MAX_UPLOAD_SIZE := config.GetMaxUploadFileSize()

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

	//gid, _ := strconv.Atoi(p.ByName("id"))

	rand_num := utils.GetRandNumByNumber(100)

	file_url := GAME_FILE_APK + strconv.Itoa(int(rand_num)) +  handler.Filename

	err = ioutil.WriteFile(file_url, data, 0666)

	if err != nil {
		log.Printf("Write file error: %v", err)
		utils.SendErrorResponse(w, defs.ErrorUploadFaults)
		return
	}

	file_tmp := strings.TrimPrefix(file_url, ".")

	res_file_url := config.GetFileServerDirUrl() + file_tmp

	resObj := defs.FileUpload{
		Name: handler.Filename,
		Url:  res_file_url,
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    &resObj,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 游戏添加
func GameCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.GameCreate{}
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(err)
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 是否存在游戏名称相同
	if b := dbops.GetGameByGameName(ubody.Name); !b {
		utils.SendErrorResponse(w, defs.ErrorGameNameIsHave)
		return
	}

	if err := dbops.GameCreate(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "添加游戏成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 201)
}

// 游戏修改
func GameUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.GameCreate{}
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(err)
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 是否存在游戏名称相同
	if b := dbops.GameNameIsHave(ubody.Id, ubody.Name); !b {
		utils.SendErrorResponse(w, defs.ErrorGameNameIsHave)
		return
	}

	if err := dbops.GameUpdate(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "修改成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 201)
}

// 获取修改游戏初始信息
func GameUpdateInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	res, err := dbops.GameUpdateInfo(id)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "修改成功",
		Data:    res,
	}

	utils.SendNormalResponse(w, *resData, 200)
}