package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"api/admin/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
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
