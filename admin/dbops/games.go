package dbops

import (
	"api/admin/defs"
	"api/admin/utils"
	"database/sql"
	"fmt"
	"log"
)

// 获取游戏列表
func GetGameList(page, to int, name string) (*defs.GameList, error) {
	gameList := &defs.GameList{}
	gameList.PageNum = page // 1
	gameList.PageSize = to  // 10

	whereName := ""
	if name == "" {
		whereName = "'%'"
	} else {
		whereName = "'%" + name + "%'"
	}

	//whereUrl := ""
	//if url == "" {
	//	whereUrl = "'%'"
	//} else {
	//	whereUrl = "'%" + url + "%'"
	//}

	//whereCateGoryId := ""
	//if id != 0 {
	//	//println(id)
	//	whereCateGoryId = " AND category_id = " + strconv.Itoa(id)
	//}

	sqlQuery := "SELECT COUNT(*) FROM game WHERE name like " + whereName
	//fmt.Println(sqlQuery)
	totalRow, err := dbConn.Query(sqlQuery)
	if err != nil {
		fmt.Println("get total games sql error", err)
		return nil, err
	}
	total := 0
	for totalRow.Next() {
		err := totalRow.Scan(
			&total,
		)
		if err != nil {
			fmt.Println("get total game error", err)
			continue
		}
	}
	totalRow.Close()

	// 获取总页数
	maxpage := utils.GetPageLimit(total, to)
	gameList.Total = total
	gameList.TotalPage = maxpage

	stmtRole, err := dbConn.Prepare("SELECT id, icon, name, company, mana, attention, down_url, game_desc, game_size, game_version, update_time, company_tag, create_time FROM game WHERE name like " + whereName  + " LIMIT ?,?")
	defer stmtRole.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get game info err: %s", err)
		return gameList, err
	}
	stmtRows, err := stmtRole.Query((page-1)*to, to)
	for stmtRows.Next() {
		line := &defs.Game{}
		err = stmtRows.Scan(&line.Id, &line.Icon, &line.Name, &line.Company, &line.Mana, &line.Attention, &line.DownUrl, &line.GameDesc, &line.GameSize, &line.GameVersion, &line.UpdateTime, &line.CompanyTag, &line.CreateTime)
		if err != nil {
			log.Printf("game sql scan error: %s", err)
			return gameList, err
		}
		gameList.List = append(gameList.List, line)
	}

	return gameList, nil
}
