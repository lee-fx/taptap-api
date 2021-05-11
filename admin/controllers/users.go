package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
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
	ubody := &defs.UserCreadential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
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

func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

func AdminUserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	//uname := r.PostFormValue("username")
	//pwd := r.PostFormValue("password")
	//
	//fmt.Println(uname)
	//fmt.Println(pwd)
	//io.WriteString(w, uname)
	//io.WriteString(w, pwd)
	tmp := &defs.NormalResponse{
		Code:    "200",
		Message: "登陆成功",
		Data: "111",
	}

	utils.SendNormalResponse(w, *tmp, 200 )
}