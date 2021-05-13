package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"fmt"
	"log"

	//"api/admin/session"
	"api/admin/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
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

	//id := session.GenerateNewSesssionId(ubody.Username)
	//su := &defs.SignedUp{
	//	Success:   true,
	//	SessionId: id,
	//}
	//
	//if resp, err := json.Marshal(su); err != nil {
	//	utils.SendErrorResponse(w, defs.ErrorInternalFaults)
	//	return
	//} else {
	//	utils.SendNormalResponse(w, string(resp), 201)
	//}
}

func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

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
	user.Name = ubody.UserName
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


func AdminUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// 从token中获取用户id
	claims := utils.GetTokenParseInfo(r)
	uid := claims.Id

	// 通过用户id组装用户操作menu信息
	fmt.Println(uid)

	if res, err := dbops.GetUserInfo(uid); err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		lng, _ := json.Marshal(res)
		utils.SendNormalResponse(w, string(lng), 200)
		return
	}

	roles := []string{"超级管理员"}
	menus := []defs.Menu{}
	menu1 := defs.Menu{
		Id:         1,
		ParentId:   0,
		CreateTime: "2020-02-02T06:50:36.000+00:00",
		Title:      "商品",
		Level:      0,
		Sort:       0,
		Name:       "pms",
		Icon:       "product",
		Hidden:     0,
	}

	menus = append(menus, menu1)

	data := defs.UserInfo{
		Roles: roles,
		Icon:  "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg.jpg",
		Menus: menus,
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "路由信息",
		Data:    data,
	}

	utils.SendNormalResponse(w, *tmp, 200)
}

func AdminUserLogout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "登出成功",
		Data:    "",
	}

	utils.SendNormalResponse(w, *tmp, 200)
}
