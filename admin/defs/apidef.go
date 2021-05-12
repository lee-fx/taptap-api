package defs

// User
type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	PassWord string `json:"passWord"`
	Avatar string `json:"avatar"`
	Iphone string `json:"iphone"`
	ShowTime string `json:"showTime"`
	CreateTime string `json:"createTime"`
}



// requests
type UserLogin struct {
	UserName string `json:"username"`
	PassWord      string `json:"password"`
}

//

// response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// session
type SimpleSession struct {
	Username string // login name
	TTL      int64
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
	Hidden     int `json:"hidden"`
}

// user login
type UserToken struct {
	TokenHead string `json:"tokenHead"`
	Token     string `json:"token"`
}

// user info
type UserInfo struct {
	Roles []string `json:"roles"`
	Icon  string   `json:"icon"`
	Menus []Menu `json:"menus"`
}
