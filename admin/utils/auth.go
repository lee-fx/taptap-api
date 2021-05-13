package utils

import (
	"net/http"
	"strings"
)

var REQUEST_HEADER = "Authorization"
var HEADER_LTRIM = "Bearer "


// 校验token
func ValidateJwtToken(r *http.Request) bool {
	htoken := r.Header.Get(REQUEST_HEADER)
	//fmt.Println(htoken)
	token := strings.TrimPrefix(htoken, HEADER_LTRIM)
	//fmt.Println(token)
	_, err := ParseToken(token)
	if err != nil {
		//fmt.Printf("token过期:%v", err)
		return true
	}
	// 白名单|黑名单
	//fmt.Printf("解析出来用户数据:%v", claims)
	return false
}

// 获取token的信息
func GetTokenParseInfo(r *http.Request) *Claims {
	htoken := r.Header.Get(REQUEST_HEADER)
	//fmt.Println(htoken)
	token := strings.TrimPrefix(htoken, HEADER_LTRIM)
	//fmt.Println(token)
	claims, err := ParseToken(token)
	if err != nil {
		//fmt.Printf("token过期:%v", err)
		return nil
	}
	// 白名单|黑名单
	//fmt.Printf("解析出来用户数据:%v", claims)
	return claims
}