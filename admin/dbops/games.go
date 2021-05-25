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

	stmtRole, err := dbConn.Prepare("SELECT id, icon, name, mana, attention, down_url, game_desc, game_size, game_version, update_time, create_time, status FROM game WHERE name like " + whereName + " LIMIT ?,?")
	defer stmtRole.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get game info err: %s", err)
		return gameList, err
	}
	stmtRows, err := stmtRole.Query((page-1)*to, to)
	for stmtRows.Next() {
		line := &defs.Game{}
		err = stmtRows.Scan(&line.Id, &line.Icon, &line.Name, &line.Mana, &line.Attention, &line.DownUrl, &line.GameDesc, &line.GameSize, &line.GameVersion, &line.UpdateTime, &line.CreateTime, &line.Status)
		if err != nil {
			log.Printf("game sql scan error: %s", err)
			return gameList, err
		}
		gameList.List = append(gameList.List, line)
	}

	return gameList, nil
}

// 获取所有标签
func GetGameTag() ([]string, error) {

	gameTags := []string{}
	stmtTags, err := dbConn.Prepare("SELECT tag_name FROM game_tag")
	defer stmtTags.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get tags info err: %s", err)
		return gameTags, err
	}
	stmtRows, err := stmtTags.Query()
	for stmtRows.Next() {
		line := ""
		err = stmtRows.Scan(&line)
		if err != nil {
			log.Printf("gettags sql scan error: %s", err)
			return gameTags, err
		}
		gameTags = append(gameTags, line)
	}

	return gameTags, nil
}



// 游戏标签获取
func GetGameTagByGameId(gid int) ([]string, error) {

	gameTags := []string{}
	stmtTags, err := dbConn.Prepare("SELECT tag_id FROM game_tag_relation WHERE game_id = ?")
	defer stmtTags.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get game tags info err: %s", err)
		return gameTags, err
	}
	stmtRows, err := stmtTags.Query(gid)
	for stmtRows.Next() {
		line := ""
		err = stmtRows.Scan(&line)
		if err != nil {
			log.Printf("tagnames sql scan error: %s", err)
			return gameTags, err
		}
		gameTags = append(gameTags, line)
	}

	return gameTags, nil
}
