package defs

// game 标签
type GameTag struct {
	Id int
	TagName string
}

// 全部游戏列表使用
type GameList struct {
	Id int
	Icon string
	Name string
	GameTag []GameTag
	Mana string
}

// 游戏详细信息使用
type GameInfo struct {
	Id int
	Icon string
	Name string
	Company string
	Mana string
	Attention string
	DownUrl string
	GameDesc string
	GameSize string
	GameVersion string
	UpdateTime string
	CompanyTag string
}

// 与GameTagArr 结合构造主页推荐
type Game struct {
	Id int
	Icon string
	Name string
}

// 与GameTag结合
type GameTagArr struct {
	Type int
	Title string
	GameList []*Game
}

// Global获取结构
type Global struct {
	Id int
	GlobalValue string
	GlobalSupplement string // 补充字段
	Name string
}

// GameBanner
type GameBanner struct {
	Id int
	ImgUrl string
}

// RecommondGame
type RecommondGame struct {
	Id int
	Icon string
	Name string
	GameTag  []GameTag
	Mana string
	GameBanner GameBanner
}


// requests
type UserCreadential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

// response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

// data model

// 视频信息
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

// 评论
type Comment struct {
	Id string
	VideoId string
	Author  string
	Content string
}

// session
type SimpleSession struct {
	Username string  // login name
	TTL		 int64
}