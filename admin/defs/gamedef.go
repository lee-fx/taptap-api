package defs

// 游戏详细信息使用
type Game struct {
	Id          int    `json:"id"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Mana        string `json:"mana"`
	Attention   string `json:"attention"`
	DownUrl     string `json:"downUrl"`
	GameDesc    string `json:"gameDesc"`
	GameSize    string `json:"gameSize"`
	GameVersion string `json:"gameVersion"`
	CompanyTag  string `json:"companyTag"`
	UpdateTime  string `json:"updateTime"`
	CreateTime  string `json:"CreateTime"`

}

type GameList struct {
	PageNum   int     `json:"pageNum"`
	PageSize  int     `json:"pageSize"`
	TotalPage int     `json:"totalPage"`
	Total     int     `json:"total"`
	List      []*Game `json:"list"`
}
