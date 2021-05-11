package main

import (
	"api/admin/defs"
	"api/admin/utils"
	//"api/admin/session"
	"net/http"
)

var HEADER_FIELD_SESSIOON = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSIOON)
	if len(sid) == 0 {
		return false
	}
	//uname, ok := session.IsSessionExpired(sid)
	//if ok {
	//	return false
	//}
	//r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		utils.SendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
