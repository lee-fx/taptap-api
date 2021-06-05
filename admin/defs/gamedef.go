package defs

// 游戏详细信息使用
type Game struct {
	Id          int    `json:"id"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Mana        string `json:"mana"`
	Attention   string `json:"attention"`
	Company     string `json:"company"`
	DownUrl     string `json:"downUrl"`
	GameDesc    string `json:"gameDesc"`
	GameSize    string `json:"gameSize"`
	GameVersion string `json:"gameVersion"`
	UpdateTime  string `json:"updateTime"`
	CreateTime  string `json:"CreateTime"`
	Status      int    `json:"status"`
}

type GameCreate struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Mana        string         `json:"mana"`
	Attention   int         `json:"attention,string"`
	Image       *FileUpload `json:"image"`
	File        *FileUpload `json:"file"`
	Description string      `json:"description"`
	GameSize    string      `json:"game_size"`
	GameVersion string      `json:"game_version"`
	Status      int         `json:"status"`

	CompanyId  int    `json:"company_id"`
	GameTagIds string `json:"game_tag_ids"`
}

type GameList struct {
	PageNum   int     `json:"pageNum"`
	PageSize  int     `json:"pageSize"`
	TotalPage int     `json:"totalPage"`
	Total     int     `json:"total"`
	List      []*Game `json:"list"`
}

type GameTag struct {
	Id      int    `json:"id"`
	TagName string `json:"tagName"`
}

// 公司
type Company struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ShortTag string `json:"shortTag"`
}

type CompanyList struct {
	PageNum   int        `json:"pageNum"`
	PageSize  int        `json:"pageSize"`
	TotalPage int        `json:"totalPage"`
	Total     int        `json:"total"`
	List      []*Company `json:"list"`
}

type FileUpload struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
