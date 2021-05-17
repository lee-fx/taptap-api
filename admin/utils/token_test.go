package utils

import (
	"api/admin/defs"
	"fmt"
	"testing"
)


func TestToken(t *testing.T) {
	t.Run("token create", cToken)
	t.Run("token parse", pToken)

}

// 生成token
func cToken(t *testing.T) {
	user := defs.User{
		Id:         1,
		UserName:       "阿飞",
		PassWord:   "1234578945",
		Icon:     "sadf",
		Email:     "1325132780@qq.com",
		LoginTime:   "污了",
		CreateTime: "污了",
	}
	token, err := GenerateToken(&user)
	if err != nil {
		fmt.Printf("get token err : %v", err)
	}
	fmt.Println(token)
}

// 解析token
func pToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlck5hbWUiOiLpmL_po54iLCJpcGhvbmUiOiIxNTcxODgxNTIzMSIsImV4cCI6MTYyMDgxNDI4NSwiaXNzIjoibGVlLWZ4In0.bJVb7VrTHa3zF8P2m0eu4chNQHUO7SxdBAUNw7HuB_Y"
	res, err := ParseToken(token)
	if err != nil {
		fmt.Printf("parser token err : %v", err)
	}
	fmt.Printf("iphone: %v\n", res.Email)
}