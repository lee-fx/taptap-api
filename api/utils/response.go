package utils

import (
	"admin/api/defs"
	"encoding/json"
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErroResponse) {
	w.WriteHeader(errResp.HttpSc)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}