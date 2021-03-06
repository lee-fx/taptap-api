package defs

// User
type User struct {
	Id         int64  `json:"id"`
	UserName   string `json:"username"`
	NickName   string `json:"nickName"`
	PassWord   string `json:"password"`
	Icon       string `json:"icon"`
	Email      string `json:"email"`
	Note       string `json:"note"`
	CreateTime string `json:"createTime"`
	LoginTime  string `json:"loginTime"`
	Status     int    `json:"status"`
}

// User List
type UserList struct {
	PageNum   int     `json:"pageNum"`
	PageSize  int     `json:"pageSize"`
	TotalPage int     `json:"totalPage"`
	Total     int     `json:"total"`
	List      []*User `json:"list"`
}

// Role List
type RoleList struct {
	PageNum   int     `json:"pageNum"`
	PageSize  int     `json:"pageSize"`
	TotalPage int     `json:"totalPage"`
	Total     int     `json:"total"`
	List      []*Role `json:"list"`
}

// Resource
type ResourceList struct {
	PageNum   int         `json:"pageNum"`
	PageSize  int         `json:"pageSize"`
	TotalPage int         `json:"totalPage"`
	Total     int         `json:"total"`
	List      []*Resource `json:"list"`
}

// Menu List
type MenuList struct {
	PageNum   int     `json:"pageNum"`
	PageSize  int     `json:"pageSize"`
	TotalPage int     `json:"totalPage"`
	Total     int     `json:"total"`
	List      []*Menu `json:"list"`
}

// login requests
type UserLogin struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// 修改游戏tag
type TagNames struct {
	TagNames string `json:"tagNames"`
}

// 分页参数
type PageParams struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

// response
type NormalResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Menu struct {
	Id         int     `json:"id"`
	ParentId   int     `json:"parentId"`
	CreateTime string  `json:"createTime"`
	Title      string  `json:"title"`
	Level      int     `json:"level"`
	Sort       int     `json:"sort,string"`
	Name       string  `json:"name"`
	Icon       string  `json:"icon"`
	Hidden     int     `json:"hidden"`
	Children   []*Menu `json:"children"`
}

// user info
type UserInfo struct {
	Roles []string `json:"roles"`
	Icon  string   `json:"icon"`
	Menus []*Menu  `json:"menus"`
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

// resourceCategory
type ResourceCategory struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Sort       int    `json:"sort,string"`
	CreateTime string `json:"createTime"`
}

// resource 
type Resource struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	CategoryId  int    `json:"categoryId"`
	CreateTime  string `json:"createTime"`
}
