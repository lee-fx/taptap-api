package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"log"
	"strconv"

	//"strconv"

	//"api/admin/session"
	"api/admin/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserLogin{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.UserName, ubody.PassWord); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
}

// 查询所有用户
func AdminUserList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//获取请求参数
	params := r.URL.Query()
	page, _ := strconv.Atoi(params["pageNum"][0])
	to, _ := strconv.Atoi(params["pageSize"][0])

	res, err := dbops.AdminUserList(page, to)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "登陆成功",
		Data:    res,
	}

	utils.SendNormalResponse(w, *resData, 200)

}

// 用户登录
func AdminUserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取请求参数
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserLogin{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 数据库校验用户名密码
	user, err := dbops.VerifyUserLogin(ubody)
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	if user.Id == 0 {
		utils.SendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	// 生成jwt-token并返回
	user.UserName = ubody.UserName
	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "登陆成功",
		Data: defs.UserToken{
			TokenHead: "Bearer ",
			Token:     token,
		},
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 用户路由信息
func AdminUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// 从token中获取用户id
	claims := utils.GetTokenParseInfo(r)
	uid := claims.Id

	// 通过用户id组装用户操作menu信息
	//fmt.Println(uid)
	res, err := dbops.GetUserInfo(uid)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "登录路由",
		Data:    res,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 登出
func AdminUserLogout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "登出成功",
		Data:    "",
	}
	utils.SendNormalResponse(w, *tmp, 200)
}
