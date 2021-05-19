package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"api/admin/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// 获取菜单tree列表
func MenuTreeList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := dbops.MenuTreeList()
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	// 组装tree


	resData := &defs.NormalResponse{
		Code:    200,
		Message: "success",
		Data:    res,
	}
	utils.SendNormalResponse(w, *resData, 200)
}
