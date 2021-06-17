package dbops

import (
	"api/app/defs"
	"fmt"
	"log"
)

func GetConfigs(content string) ([]*defs.Global, error) {
	var res []*defs.Global
	stmtOut, err := dbConn.Prepare("SELECT id, global_value, name, global_supplement FROM global WHERE global_key = ?")
	if err != nil {
		log.Printf("get Global configs error: %s", err)
		return nil, err
	}
	rows, err := stmtOut.Query(content)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	for rows.Next() {
		line := defs.Global{}
		err = rows.Scan(&line.Id, &line.GlobalValue, &line.Name, &line.GlobalSupplement)
		if err != nil {
			log.Printf("sql exe err: %s", err)
			return nil, err
		}
		res = append(res, &line)
	}
	defer stmtOut.Close()

	return res, nil
}

func GetTypeGames(num int) ([]*defs.GameTagArr, error) {
	stmtOutNew, err := dbConn.Prepare("SELECT id, icon, name FROM game WHERE status=0 ORDER BY id DESC limit ?")
	if err != nil {
		log.Printf("get new games error: %s", err)
		return nil, err
	}
	stmtOutHot, err := dbConn.Prepare("SELECT id, icon, name FROM game WHERE status=0 ORDER BY attention DESC limit ?")
	if err != nil {
		log.Printf("get hot games error: %s", err)
		return nil, err
	}
	stmtOutGood, err := dbConn.Prepare("SELECT id, icon, name FROM game WHERE status=0 ORDER BY mana DESC limit ?")
	if err != nil {
		log.Printf("get good games error: %s", err)
		return nil, err
	}
	stmtOutNovel, err := dbConn.Prepare("SELECT id, icon, name FROM game WHERE status=0 ORDER BY create_time DESC limit ?")
	if err != nil {
		log.Printf("get novel games error: %s", err)
		return nil, err
	}

	newGames := defs.GameTagArr{}
	newGames.Type = 1
	newGames.Title = "新游"
	rowsNew, err := stmtOutNew.Query(num)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	for rowsNew.Next() {
		game := defs.Game{}
		err = rowsNew.Scan(&game.Id, &game.Icon, &game.Name)
		if err != nil {
			return nil, err
		}
		newGames.GameList = append(newGames.GameList, &game)
	}

	hotGames := defs.GameTagArr{}
	hotGames.Type = 2
	hotGames.Title = "热游"

	rowsHot, err := stmtOutHot.Query(num)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	for rowsHot.Next() {
		game := defs.Game{}
		err = rowsHot.Scan(&game.Id, &game.Icon, &game.Name)
		if err != nil {
			return nil, err
		}
		hotGames.GameList = append(hotGames.GameList, &game)
	}

	goodGames := defs.GameTagArr{}
	goodGames.Type = 3
	goodGames.Title = "优质推荐"

	rowsGood, err := stmtOutGood.Query(num)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	for rowsGood.Next() {
		game := defs.Game{}
		err = rowsGood.Scan(&game.Id, &game.Icon, &game.Name)
		if err != nil {
			return nil, err
		}
		goodGames.GameList = append(goodGames.GameList, &game)
	}

	novelGames := defs.GameTagArr{}
	novelGames.Type = 4
	novelGames.Title = "尝鲜特供"

	rowsNovel, err := stmtOutNovel.Query(num)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	for rowsNovel.Next() {
		game := defs.Game{}
		err = rowsNovel.Scan(&game.Id, &game.Icon, &game.Name)
		if err != nil {
			return nil, err
		}
		novelGames.GameList = append(novelGames.GameList, &game)
	}

	var resArr []*defs.GameTagArr
	resArr = append(resArr, &newGames, &hotGames, &goodGames, &novelGames)

	defer stmtOutNew.Close()
	defer stmtOutHot.Close()
	defer stmtOutGood.Close()
	defer stmtOutNovel.Close()

	return resArr, nil

}
