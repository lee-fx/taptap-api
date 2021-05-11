package controllers

import (
	"api/app/dbops"
	"api/app/defs"
	"api/app/session"
	"api/app/utils"
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

	id := session.GenerateNewSesssionId(ubody.Username)
	su := &defs.SignedUp{
		Success:   true,
		SessionId: id,
	}

	if resp, err := json.Marshal(su); err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		utils.SendNormalResponse(w, string(resp), 201)
	}
}

func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
