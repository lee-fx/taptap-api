package defs

type Err struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErroResponse struct {
	HttpSc int
	Error  Err
}

var (
	// 请求错误
	ErrorRequestBodyParseFailed = ErroResponse{HttpSc: 200, Error: Err{Message: "请求参数错误.", Code: 100001, Data: nil}}

	// 校验错误
	ErrorNotAuthUser = ErroResponse{HttpSc: 200, Error: Err{Message: "用户登录校验失败.", Code: 200001, Data: nil}}

	ErrorJwtTokenValidateFaild = ErroResponse{HttpSc: 200, Error: Err{Message: "用户token校验失败.", Code: 200002, Data: nil}}

	// 数据错误
	ErrorDBError = ErroResponse{HttpSc: 200, Error: Err{Message: "DB ops failed.", Code: 300001, Data: nil}}

	// 内部错误
	ErrorInternalFaults = ErroResponse{HttpSc: 200, Error: Err{Message: "Internal service error.", Code: 400001, Data: nil}}
)
