package utils

import (
	"api/admin/defs"
	"encoding/json"
	"io"
	"net/http"
)

// 发送错误响应
func SendErrorResponse(w http.ResponseWriter, errResp defs.ErroResponse) {
	w.WriteHeader(errResp.HttpSc)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

// 发送正常响应
func SendNormalResponse(w http.ResponseWriter, resp defs.NormalResponse, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(&resp)
	io.WriteString(w, string(resStr))
}