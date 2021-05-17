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

	ErrorUserEmailIsEmpty = ErroResponse{HttpSc: 200, Error: Err{Message: "邮箱为空.", Code: 200003, Data: nil}}
	ErrorUserNikeNameIsEmpty = ErroResponse{HttpSc: 200, Error: Err{Message: "昵称为空.", Code: 200004, Data: nil}}
	ErrorUserPwdIsEmpty = ErroResponse{HttpSc: 200, Error: Err{Message: "密码不可为空.", Code: 200005, Data: nil}}
	ErrorUserUserNameIsEmpty = ErroResponse{HttpSc: 200, Error: Err{Message: "用户名不可为空.", Code: 200006, Data: nil}}
	ErrorUserEmailValidateFaild = ErroResponse{HttpSc: 200, Error: Err{Message: "邮箱格式错误.", Code: 200003, Data: nil}}

	ErrorUserIsHave = ErroResponse{HttpSc: 200, Error: Err{Message: "用户已存在.", Code: 200008, Data: nil}}

	// 数据错误
	ErrorDBError = ErroResponse{HttpSc: 200, Error: Err{Message: "DB ops failed.", Code: 300001, Data: nil}}

	// 内部错误
	ErrorInternalFaults = ErroResponse{HttpSc: 200, Error: Err{Message: "Internal service error.", Code: 400001, Data: nil}}
)
