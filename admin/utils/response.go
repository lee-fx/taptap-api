package utils

import (
	"api/admin/defs"
	"encoding/json"
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErroResponse) {
	w.WriteHeader(errResp.HttpSc)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp defs.NormalResponse, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(&resp)
	io.WriteString(w, string(resStr))
}