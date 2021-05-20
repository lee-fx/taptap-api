package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"api/admin/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// 所有角色信息
func AdminRoleListAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := dbops.AdminRoleListAll()
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    res,
	}
	utils.SendNormalResponse(w, *resData, 200)
}

// 角色信息列表
func AdminRoleList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取请求参数
	params := r.URL.Query()
	page, _ := strconv.Atoi(params["pageNum"][0])
	to, _ := strconv.Atoi(params["pageSize"][0])

	// 查询参数
	keyword := params.Get("keyword")

	res, err := dbops.AdminRoleList(page, to, keyword)
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

// 查询用户的角色
func AdminRoles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 用户id
	uid, _ := strconv.Atoi(p.ByName("id"))
	res, err := dbops.AdminRoles(uid)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    res,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 修改用户角色
func AdminRoleUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取请求参数

	r.ParseMultipartForm(32 << 20)
	uid, _ := strconv.Atoi(r.Form.Get("adminId"))
	ids := r.Form.Get("roleIds")

	err := dbops.AdminRoleUpdate(uid, ids)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    nil,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 修改角色状态
func RoleUpdateStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取用户id
	rid, _ := strconv.Atoi(p.ByName("rid"))
	// 状态
	params := r.URL.Query()
	status, _ := strconv.Atoi(params["status"][0])

	if err := dbops.RoleUpdateStatus(rid, status); err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "修改成功",
		Data:    nil,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 删除角色
func RoleDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseMultipartForm(32 << 20)
	ids, _ := strconv.Atoi(r.Form.Get("ids"))
	err := dbops.RoleDelete(ids)
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

// 创建角色
func RoleCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.Role{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 是否为空
	if ubody.Name == "" {
		utils.SendErrorResponse(w, defs.ErrorRoleNameIsEmpty)
		return
	}

	if err := dbops.RoleCreate(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	resData := &defs.NormalResponse{
		Code:    200,
		Message: "添加角色成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 201)
}

// 修改角色
func RoleUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.Role{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	// 是否为空
	if ubody.Name == "" {
		utils.SendErrorResponse(w, defs.ErrorRoleNameIsEmpty)
		return
	}

	if err := dbops.RoleUpdate(ubody); err != nil {
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

// 获取角色资源
func RoleResourceByRoleId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 角色id
	rid, _ := strconv.Atoi(p.ByName("rid"))
	res, err := dbops.RoleResourceByRoleId(rid)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "success",
		Data:    res,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 角色重新分配权限
func RoleAllocResource(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取请求参数
	r.ParseMultipartForm(32 << 20)
	rid, _ := strconv.Atoi(r.Form.Get("roleId"))
	ids := r.Form.Get("resourceIds")

	err := dbops.RoleAllocResource(rid, ids)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    nil,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 获取角色菜单
func RoleListMenuByRid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 角色id
	rid, _ := strconv.Atoi(p.ByName("rid"))
	res, err := dbops.RoleListMenuByRid(rid)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "success",
		Data:    res,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 角色分配菜单
func RoleAllocMenu(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取请求参数
	r.ParseMultipartForm(32 << 20)
	rid, _ := strconv.Atoi(r.Form.Get("roleId"))
	ids := r.Form.Get("menuIds")

	err := dbops.RoleAllocMenu(rid, ids)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    nil,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}