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

func AdminRoleList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := dbops.AdminRoleList()
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
	uid,_ := strconv.Atoi(r.Form.Get("adminId"))
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
