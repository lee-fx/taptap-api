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

// 资源列表
func ResourceListAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := dbops.ResourceListAll()
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

// 资源列表
func ResourceList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取请求参数
	params := r.URL.Query()
	page, _ := strconv.Atoi(params["pageNum"][0])
	to, _ := strconv.Atoi(params["pageSize"][0])

	// 查询参数
	name := params.Get("nameKeyword")
	url := params.Get("urlKeyword")
	id, _ := strconv.Atoi(params.Get("categoryId"))

	res, err := dbops.ResourceList(page, to, name, url, id)
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

// 资源创建
func ResourceCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.Resource{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 是否为空
	if ubody.Name == "" {
		utils.SendErrorResponse(w, defs.ErrorResourceNameIsEmpty)
		return
	}

	if ubody.Url == "" {
		utils.SendErrorResponse(w, defs.ErrorResourceUrlIsEmpty)
		return
	}

	if err := dbops.ResourceCreate(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	resData := &defs.NormalResponse{
		Code:    200,
		Message: "添加资源成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 201)
}

// 资源修改
func ResourUpdateById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.Resource{}
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(err)
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 是否为空
	if ubody.Name == "" {
		utils.SendErrorResponse(w, defs.ErrorResourceNameIsEmpty)
		return
	}

	if ubody.Url == "" {
		utils.SendErrorResponse(w, defs.ErrorResourceUrlIsEmpty)
		return
	}

	if err := dbops.ResourUpdateById(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "修改资源成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 资源删除
func ResourDeleteById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	err := dbops.ResourDeleteById(id)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "删除成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 资源分类列表
func ResourceCategoryListAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	res, err := dbops.ResourceCategoryListAll()
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

// 分类创建
func ResourceCategoryCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.ResourceCategory{}
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Printf("%s", err)
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 是否为空
	if ubody.Name == "" {
		utils.SendErrorResponse(w, defs.ErrorResourceCategoryIsEmpty)
		return
	}

	if strconv.Itoa(ubody.Sort) == "" {
		ubody.Sort = 0
	}

	if err := dbops.ResourceCategoryCreate(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	resData := &defs.NormalResponse{
		Code:    200,
		Message: "添加成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 201)
}

// 通过id删除资源分类
func ResourceCategoryDeleteById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	err := dbops.ResourceCategoryDeleteById(id)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "删除成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 通过id修改资源分类
func ResourceCategoryUpdateById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.ResourceCategory{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 是否为空
	if ubody.Name == "" {
		utils.SendErrorResponse(w, defs.ErrorResourceCategoryIsEmpty)
		return
	}

	if strconv.Itoa(ubody.Sort) == "" {
		ubody.Sort = 0
	}

	if err := dbops.ResourceCategoryUpdateById(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "修改角色成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 200)
}
