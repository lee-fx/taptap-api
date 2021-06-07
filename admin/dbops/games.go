package dbops

import "C"
import (
	"api/admin/defs"
	"api/admin/utils"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// 获取游戏列表
func GetGameList(page, to int, gameName string, cid int) (*defs.GameList, error) {
	gameList := &defs.GameList{}
	gameList.PageNum = page // 1
	gameList.PageSize = to  // 10

	whereName := ""
	if gameName == "" {
		whereName = "'%'"
	} else {
		whereName = "'%" + gameName + "%'"
	}

	sqlQuery := ""

	if cid == -1 {
		sqlQuery = "SELECT COUNT(*) FROM game AS G RIGHT JOIN game_company_relation AS C ON G.id = C.game_id WHERE C.company_id <> ? AND G.name like " + whereName
	} else {
		sqlQuery = "SELECT COUNT(*) FROM game AS G RIGHT JOIN game_company_relation AS C ON G.id = C.game_id WHERE C.company_id = ? AND G.name like " + whereName
	}

	totalRow, err := dbConn.Query(sqlQuery, cid)
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

	sqlQueryList := ""

	if cid == -1 {
		sqlQueryList = "SELECT G.id, G.icon, G.name, G.mana, G.attention, G.down_url, G.game_desc, G.game_size, G.game_version, G.update_time, G.create_time, G.status FROM game AS G RIGHT JOIN game_company_relation AS C ON G.id = C.game_id WHERE C.company_id <> ? AND G.name like " + whereName + " ORDER BY id desc LIMIT ?,?"
	} else {
		sqlQueryList = "SELECT G.id, G.icon, G.name, G.mana, G.attention, G.down_url, G.game_desc, G.game_size, G.game_version, G.update_time, G.create_time, G.status FROM game AS G RIGHT JOIN game_company_relation AS C ON G.id = C.game_id WHERE C.company_id = ? AND G.name like " + whereName + " ORDER BY id desc LIMIT ?,?"
	}

	//fmt.Println(sqlQueryList)

	stmtGame, err := dbConn.Prepare(sqlQueryList)
	defer stmtGame.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get game info err: %s", err)
		return gameList, err
	}
	stmtRows, err := stmtGame.Query(cid, (page-1)*to, to)
	for stmtRows.Next() {
		line := &defs.Game{}
		err = stmtRows.Scan(&line.Id, &line.Icon, &line.Name, &line.Mana, &line.Attention, &line.DownUrl, &line.GameDesc, &line.GameSize, &line.GameVersion, &line.UpdateTime, &line.CreateTime, &line.Status)
		if err != nil {
			log.Printf("game sql scan error: %s", err)
			return gameList, err
		}
		// 查找游戏的关联公司 没有找到就为空
		sqlCompany := "SELECT C.name, C.short_tag FROM game_company AS C RIGHT JOIN game_company_relation AS R ON R.company_id = C.id WHERE R.game_id = ?"
		stmtGame, err := dbConn.Prepare(sqlCompany)
		defer stmtGame.Close()
		if err != nil && err != sql.ErrNoRows {
			log.Printf("get company info err: %s", err)
			return gameList, err
		}
		var name = ""
		var short_tag = ""
		err = stmtGame.QueryRow(line.Id).Scan(&name, &short_tag)
		if err != nil {
			log.Printf("get menu scan error: %s", err)
			return gameList, err
		}
		line.Company = name + "(" + short_tag + ")"
		gameList.List = append(gameList.List, line)
	}

	return gameList, nil
}

// 获取所有标签
func GetGameTag() ([]*defs.GameTag, error) {

	gameTags := []*defs.GameTag{}
	stmtTags, err := dbConn.Prepare("SELECT id, tag_name FROM game_tag")
	defer stmtTags.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get tags info err: %s", err)
		return gameTags, err
	}
	stmtRows, err := stmtTags.Query()
	for stmtRows.Next() {
		line := &defs.GameTag{}
		err = stmtRows.Scan(&line.Id, &line.TagName)
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
		log.Printf("get game tag_id info err: %s", err)
		return gameTags, err
	}
	stmtRows, err := stmtTags.Query(gid)
	for stmtRows.Next() {
		line := ""
		err = stmtRows.Scan(&line)
		if err != nil {
			log.Printf("tagids sql scan error: %s", err)
			return gameTags, err
		}
		gameTags = append(gameTags, line)
	}
	return gameTags, nil
}

// 修改游戏标签
func GameTagUpdateByGameId(gid int, tag_names *defs.TagNames) error {
	// 删除原有id关联tag
	stmtDel, err := dbConn.Prepare("DELETE FROM game_tag_relation WHERE game_id=?")
	if err != nil {
		log.Printf("delete game tag relation error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(gid)
	defer stmtDel.Close()

	// 判断是否为空串
	if tag_names.TagNames == "" {
		return nil
	}

	ids := strings.Split(tag_names.TagNames, ",")

	for _, tag_id := range ids {
		// 增加关系
		stmtIns, err := dbConn.Prepare("INSERT INTO game_tag_relation (game_id, tag_id) VALUES(?,?)")
		if err != nil {
			fmt.Printf("insert game tag relation error: %v", err)
			return err
		}
		_, err = stmtIns.Exec(gid, tag_id)
		if err != nil {
			fmt.Printf("insert game tag relation exe error: %v", err)
			return err
		}
		defer stmtIns.Close()
	}

	return nil
}

// 获取公司列表
func GetCompanyList(page, to int) (*defs.CompanyList, error) {
	companyList := &defs.CompanyList{}
	companyList.PageNum = page // 1
	companyList.PageSize = to  // 100

	sqlQuery := "SELECT COUNT(*) FROM game "
	totalRow, err := dbConn.Query(sqlQuery)
	if err != nil {
		fmt.Println("get total companys sql error", err)
		return nil, err
	}
	total := 0
	for totalRow.Next() {
		err := totalRow.Scan(
			&total,
		)
		if err != nil {
			fmt.Println("get total company error", err)
			continue
		}
	}
	totalRow.Close()

	// 获取总页数
	maxpage := utils.GetPageLimit(total, to)
	companyList.Total = total
	companyList.TotalPage = maxpage

	stmtRole, err := dbConn.Prepare("SELECT id, name, short_tag FROM game_company LIMIT ?,?")
	defer stmtRole.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get company info err: %s", err)
		return companyList, err
	}
	stmtRows, err := stmtRole.Query((page-1)*to, to)
	for stmtRows.Next() {
		line := &defs.Company{}
		err = stmtRows.Scan(&line.Id, &line.Name, &line.ShortTag)
		if err != nil {
			log.Printf("company sql scan error: %s", err)
			return companyList, err
		}
		companyList.List = append(companyList.List, line)
	}

	return companyList, nil
}

// 修改游戏状态
func GameUpdateStatus(gs int, ids string) error {
	// 修改game的状态
	sql := "UPDATE game SET status  = ? WHERE id IN " + "(" + ids + ")"
	//fmt.Println(sql)
	stmtUpdate, err := dbConn.Prepare(sql)
	if err != nil {
		log.Printf("update game status error: %s", err)
	}
	_, err = stmtUpdate.Exec(gs)
	defer stmtUpdate.Close()
	return nil
}

// 删除游戏
func GameDeleteByIds(ids string) error {
	// 修改game的状态
	sql := "DELETE FROM game WHERE id IN " + "(" + ids + ")"
	//fmt.Println(sql)
	stmtDel, err := dbConn.Prepare(sql)
	if err != nil {
		log.Printf("delete game error: %s", err)
	}
	_, err = stmtDel.Exec()
	defer stmtDel.Close()
	return nil
}

// 查看游戏名称是否存在
func GetGameByGameName(game_name string) bool {

	stmtOut, err := dbConn.Prepare("SELECT id FROM game WHERE name = ? ")
	if err != nil {
		log.Printf("get game name error: %s", err)
		return true
	}
	var id string
	err = stmtOut.QueryRow(game_name).Scan(&id)
	if err != nil {
		log.Printf("get game name scan error: %s", err)
		return true
	}

	if err == sql.ErrNoRows {
		return true
	}
	stmtOut.Close()

	return false
}

// 除了修改的游戏，是否存在同名
func GameNameIsHave(gid int, game_name string) bool {

	stmtOut, err := dbConn.Prepare("SELECT id FROM game WHERE name = ? AND id != ?")
	if err != nil {
		log.Printf("get game name error: %s", err)
		return true
	}
	var id string
	err = stmtOut.QueryRow(game_name, gid).Scan(&id)
	if err != nil {
		log.Printf("get game name scan error: %s", err)
		return true
	}

	if err == sql.ErrNoRows {
		return true
	}
	stmtOut.Close()

	return false
}

// 游戏添加
func GameCreate(game *defs.GameCreate) error {

	//log.Printf("%v", game)
	stmtIns, err := dbConn.Prepare("INSERT INTO game (name, mana, icon, attention, down_url, game_desc, game_size, game_version, update_time, create_time, status) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	timeNow := utils.GetTimeNowFormatDate()

	_, err = stmtIns.Exec(game.Name, game.Mana, game.Image.Url, game.Attention, game.File.Url, game.Description, game.GameSize, game.GameVersion, timeNow, timeNow, game.Status)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	defer stmtIns.Close()

	stmtOut, err := dbConn.Prepare("SELECT id FROM game WHERE name = ? ")
	if err != nil {
		log.Printf("get game name error: %s", err)
		return err
	}

	err = stmtOut.QueryRow(game.Name).Scan(&game.Id)

	if err != nil {
		log.Printf("get game name scan error: %s", err)
		return err
	}

	stmtOut.Close()

	// 处理tag_ids
	if game.GameTagIds != "" {
		idsArr := strings.Split(game.GameTagIds, ",")
		for _, tid := range idsArr {
			// 增加关系
			stmtInsTag, err := dbConn.Prepare("INSERT INTO game_tag_relation (game_id, tag_id) VALUES(?,?)")
			if err != nil {
				fmt.Printf("insert game tag relation error: %v", err)
				return err
			}
			_, err = stmtInsTag.Exec(game.Id, tid)
			if err != nil {
				fmt.Printf("insert game tag relation exe error: %v", err)
				return err
			}
			defer stmtInsTag.Close()

		}
	}

	// 处理company_id
	stmtInsCompany, err := dbConn.Prepare("INSERT INTO game_company_relation (game_id, company_id) VALUES(?,?)")
	if err != nil {
		fmt.Printf("insert game company relation error: %v", err)
		return err
	}
	_, err = stmtInsCompany.Exec(game.Id, game.CompanyId)
	if err != nil {
		fmt.Printf("insert game company relation exe error: %v", err)
		return err
	}
	defer stmtInsCompany.Close()
	//

	return nil
}

// 修改游戏
func GameUpdate(game *defs.GameCreate) error {

	stmtUpdate, err := dbConn.Prepare("UPDATE game SET name=?, mana=?, icon=?, attention=?, down_url=?, game_desc=?, game_size=?, game_version=?, update_time=?, status=? WHERE id=?")
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	timeNow := utils.GetTimeNowFormatDate()

	_, err = stmtUpdate.Exec(game.Name, game.Mana, game.Image.Url, game.Attention, game.File.Url, game.Description, game.GameSize, game.GameVersion, timeNow, game.Status, game.Id)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	defer stmtUpdate.Close()

	// 处理tag_ids
	if game.GameTagIds != "" {
		// 删除以前的
		stmtDel, err := dbConn.Prepare("DELETE FROM game_tag_relation WHERE game_id=?")
		if err != nil {
			log.Printf("delete update game tag of update game error: %s", err)
			return err
		}
		_, err = stmtDel.Exec(game.Id)
		if err != nil {
			log.Printf("delete update game tag of update game error: %s", err)
			return err
		}
		defer stmtDel.Close()

		// 新建现有的
		idsArr := strings.Split(game.GameTagIds, ",")
		for _, tid := range idsArr {
			// 增加关系
			stmtInsTag, err := dbConn.Prepare("INSERT INTO game_tag_relation (game_id, tag_id) VALUES(?,?)")
			if err != nil {
				fmt.Printf("insert game tag relation error: %v", err)
				return err
			}
			_, err = stmtInsTag.Exec(game.Id, tid)
			if err != nil {
				fmt.Printf("insert game tag relation exe error: %v", err)
				return err
			}
			defer stmtInsTag.Close()
		}
	}

	// 处理company_id
	stmtInsCompany, err := dbConn.Prepare("UPDATE game_company_relation SET company_id = ? WHERE game_id = ?")
	if err != nil {
		fmt.Printf("update game company relation error: %v", err)
		return err
	}
	_, err = stmtInsCompany.Exec(game.CompanyId, game.Id)
	if err != nil {
		fmt.Printf("update game company relation exe error: %v", err)
		return err
	}
	defer stmtInsCompany.Close()

	return nil
}

// 修改游戏获取游戏信息
func GameUpdateInfo(gid int) (*defs.GameCreate, error) {
	gameCreate := &defs.GameCreate{}
	gameCreate.Id = gid

	// 查询游戏数据并组装
	sqlGame := "SELECT name, mana, attention, icon, down_url, game_desc, game_size, game_version, status  FROM game WHERE id = ?"
	stmtGame, err := dbConn.Prepare(sqlGame)
	defer stmtGame.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get game info err: %s", err)
		return gameCreate, err
	}
	var Name string
	var Mana string
	var Attention = 0
	var Image = ""
	var File = ""
	var Description = ""
	var GameSize = ""
	var GameVersion = ""
	var Status = 0

	err = stmtGame.QueryRow(gid).Scan(&Name, &Mana, &Attention, &Image, &File, &Description, &GameSize, &GameVersion, &Status)
	if err != nil {
		log.Printf("get game info scan error: %s", err)
		return gameCreate, err
	}

	imag_struct := defs.FileUpload{
		Name: Image,
		Url:  Image,
	}

	file_struct := defs.FileUpload{
		Name: File,
		Url:  File,
	}

	gameCreate.Name = Name
	gameCreate.Mana = Mana
	gameCreate.Attention = Attention
	gameCreate.Image = &imag_struct
	gameCreate.File = &file_struct
	gameCreate.Description = Description
	gameCreate.GameSize = GameSize
	gameCreate.GameVersion = GameVersion
	gameCreate.Status = Status

	// 查找游戏的company id
	sqlCompany := "SELECT company_id FROM game_company_relation WHERE game_id = ?"
	stmtCompany, err := dbConn.Prepare(sqlCompany)
	defer stmtCompany.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get company id err: %s", err)
		return gameCreate, err
	}
	err = stmtCompany.QueryRow(gid).Scan(&gameCreate.CompanyId)
	if err != nil {
		log.Printf("get company id scan error: %s", err)
		return gameCreate, err
	}

	// tag_ids
	ids := []string{}
	stmtTags, err := dbConn.Prepare("SELECT tag_id FROM game_tag_relation WHERE game_id=?")
	defer stmtTags.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get tag ids err: %s", err)
		return gameCreate, err
	}
	if err == sql.ErrNoRows {
		gameCreate.GameTagIds = ""
	} else {
		stmtRows, err := stmtTags.Query(gid)
		for stmtRows.Next() {
			var tag_id int
			err = stmtRows.Scan(&tag_id)
			if err != nil {
				log.Printf("get tag_id sql scan error: %s", err)
				return gameCreate, err
			}
			ids = append(ids, strconv.Itoa(tag_id))
		}

		ids_s := strings.Join(ids, ",")
		gameCreate.GameTagIds = ids_s
	}

	return gameCreate, nil
}
