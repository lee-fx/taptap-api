package defs

// User
type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	PassWord   string `json:"passWord"`
	Icon       string `json:"icon"`
	Iphone     string `json:"iphone"`
	ShowTime   string `json:"showTime"`
	CreateTime string `json:"createTime"`
}

// login requests
type UserLogin struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

//

// response
type NormalResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Menu struct {
	Id         int    `json:"id"`
	ParentId   int    `json:"parentId"`
	CreateTime string `json:"createTime"`
	Title      string `json:"title"`
	Level      int    `json:"level"`
	Sort       int    `json:"sort"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Hidden     int    `json:"hidden"`
}

// user info
type UserInfo struct {
	Roles []string `json:"roles"`
	Icon  string   `json:"icon"`
	Menus []Menu   `json:"menus"`
}

// user Token
type UserToken struct {
	TokenHead string `json:"tokenHead"`
	Token     string `json:"token"`
}

// role
type Role struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AdminCount  int    `json:"adminCount"`
	CreateTime  string `json:"createTime"`
	Status      int    `json:"status"`
	Sort        int    `json:"sort"`
}
