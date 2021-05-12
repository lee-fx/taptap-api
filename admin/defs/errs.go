package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct {
	HttpSc int
	Error  Err
}

var (
	// internal

	// 请求参数错误
	ErrorRequestBodyParseFailed = ErroResponse{HttpSc: 400, Error: Err{Error: "Request body is not correct.", ErrorCode: "1001"}}
	// 登录用户验证失败
	ErrorNotAuthUser            = ErroResponse{HttpSc: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "1002"}}

	// server

	// 数据库错误
	ErrorDBError        = ErroResponse{HttpSc: 500, Error: Err{Error: "DB ops failed.", ErrorCode: "1003"}}

	// 服务器内部故障
	ErrorInternalFaults = ErroResponse{HttpSc: 500, Error: Err{Error: "Internal service error.", ErrorCode: "1004"}}
)
