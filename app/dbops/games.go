package dbops

import (
	"api/app/defs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetAllGames(Type int, page int, to int) ([]defs.GameList, error) {
	start := (page - 1) * to
	var game []defs.GameList
	var orderBy string
	switch Type {
	case 0:
		orderBy = ""
	case 1:
		orderBy = "ORDER BY id"
	case 2:
		orderBy = "ORDER BY attention desc"
	case 3:
		orderBy = "ORDER BY mana desc"
	case 4:
		orderBy = "ORDER BY create_time desc"
	default:
		orderBy = "ORDER BY id desc"
	}

	stmtOut, err := dbConn.Prepare("SELECT id, icon,name, mana FROM game WHERE status = 0 " + orderBy + " LIMIT ?,?")

	if err != nil {
		log.Printf("get Games error: %s", err)
		return game, err
	}

	rows, err := stmtOut.Query(start, to)
	if err != nil {
		return game, err
	}
	for rows.Next() {
		line := defs.GameList{}
		err = rows.Scan(&line.Id, &line.Icon, &line.Name, &line.Mana)
		if err != nil {
			return game, err
		}

		// 组装game_tag
		var gameTag []defs.GameTag

		stmtOutTag, err := dbConn.Prepare("SELECT GT.tag_name FROM game_tag GT INNER JOIN game_tag_relation GTR " +
			"ON GT.id = GTR.tag_id WHERE GTR.game_id=? limit 3")
		if err != nil {
			log.Printf("get game tag error: %s", err)
			return game, err
		}

		tagRows, err := stmtOutTag.Query(&line.Id)
		for tagRows.Next() {
			tagLine := defs.GameTag{}
			err = tagRows.Scan(&tagLine.TagName)
			if err != nil {
				log.Printf("sql scan error: %s", err)
				return game, err
			}
			gameTag = append(gameTag, tagLine)
		}
		line.GameTag = gameTag
		//gameTag = []defs.GameTag{}

		game = append(game, line)

	}
	defer stmtOut.Close()

	return game, nil
}

func GetGameInfoById(id int) (*defs.GameInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, name, icon, mana, attention, down_url, game_desc, game_size, game_version, update_time FROM game WHERE status = 0 AND id=?")
	if err != nil {
		log.Printf("get Games error: %s", err)
		return nil, err
	}
	game := &defs.GameInfo{}
	err = stmtOut.QueryRow(id).Scan(&game.Id, &game.Name, &game.Icon, &game.Mana, &game.Attention, &game.DownUrl, &game.GameDesc, &game.GameSize, &game.GameVersion, &game.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("error: %s", err)

		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	stmtCompany, err := dbConn.Prepare("SELECT GC.name FROM game_company_relation GCR INNER JOIN game_company GC ON GCR.company_id = GC.id" +
		" WHERE GCR.game_id=?")
	if err != nil {
		log.Printf("get Game company error: %s", err)
		return nil, err
	}

	err = stmtCompany.QueryRow(id).Scan(&game.Company)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("error: %s", err)

		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtCompany.Close()
	return game, nil
}

func GetRecommends(Type int, page int, to int) ([]defs.RecommondGame, error) {
	start := (page - 1) * to
	var games []defs.RecommondGame

	stmtOut, err := dbConn.Prepare("SELECT id, icon,name, mana FROM game WHERE status = 0 LIMIT ?,?")

	if err != nil {
		log.Printf("get Games error: %s", err)
		return games, err
	}

	rows, err := stmtOut.Query(start, to)
	if err != nil {
		return games, err
	}
	for rows.Next() {
		line := defs.RecommondGame{}
		err = rows.Scan(&line.Id, &line.Icon, &line.Name, &line.Mana)
		if err != nil {
			return games, err
		}

		// 组装game_tag
		var gameTag []defs.GameTag

		stmtOutTag, err := dbConn.Prepare("SELECT GT.tag_name FROM game_tag GT INNER JOIN game_tag_relation GTR " +
			"ON GT.id = GTR.tag_id WHERE GTR.game_id=? limit 3")
		if err != nil {
			log.Printf("get game tag error: %s", err)
			return games, err
		}

		tagRows, err := stmtOutTag.Query(&line.Id)
		for tagRows.Next() {
			tagLine := defs.GameTag{}
			err = tagRows.Scan(&tagLine.TagName)
			if err != nil {
				log.Printf("sql scan error: %s", err)
				return games, err
			}
			gameTag = append(gameTag, tagLine)
		}
		line.GameTag = gameTag
		stmtOutTag.Close()

		stmtOutBanner, err := dbConn.Prepare("SELECT id, img_url FROM game_banner WHERE game_id=?")

		var id int
		var img_url string

		err = stmtOutBanner.QueryRow(&line.Id).Scan(&id, &img_url)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if err == sql.ErrNoRows {
			return nil, nil
		}

		defer stmtOutBanner.Close()

		res := &defs.GameBanner{Id: id, ImgUrl: img_url}
		line.GameBanner = *res

		games = append(games, line)

	}
	defer stmtOut.Close()

	return games, nil
}
