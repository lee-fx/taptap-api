package controllers

import (
	"api/app/dbops"
	"api/app/defs"
	"api/app/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func GetAllGames(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Type, _ := strconv.Atoi(p.ByName("type"))
	page, _ := strconv.Atoi(p.ByName("page"))
	to, _ := strconv.Atoi(p.ByName("to"))
	if res, err := dbops.GetAllGames(Type, page, to); err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		lng, _ := json.Marshal(res)
		utils.SendNormalResponse(w, string(lng), 200)
		return
	}
}

func GetGameInfoById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	if res, err := dbops.GetGameInfoById(id); err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		lng, _ := json.Marshal(res)
		utils.SendNormalResponse(w, string(lng), 200)
		return
	}
}

func GetRecommends(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Type, _ := strconv.Atoi(p.ByName("type"))
	page, _ := strconv.Atoi(p.ByName("page"))
	to, _ := strconv.Atoi(p.ByName("to"))
	if res, err := dbops.GetRecommends(Type, page, to); err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		lng, _ := json.Marshal(res)
		utils.SendNormalResponse(w, string(lng), 200)
		return
	}
}
