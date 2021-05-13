package dbops

import (
	"api/admin/defs"
	"api/admin/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func VerifyUserLogin(info *defs.UserLogin) (*defs.User, error) {
	username := info.UserName
	password := utils.GetMD5HashCode([]byte(info.UserName))
	//fmt.Println(password)
	userInfo := &defs.User{}
	stmtOut, err := dbConn.Prepare("SELECT id, iphone FROM admin_user WHERE username = ? AND password = ? AND del_flag = 0")
	if err != nil{
		log.Printf("verify user sql error: %s", err)

		return nil, err
	}
	err = stmtOut.QueryRow(username, password).Scan(&userInfo.Id, &userInfo.Iphone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user not found: %s", err)
			return userInfo, nil
		} else {
			log.Printf("select db err: %s", err)
			return nil, err
		}
	}
	stmtOut.Close()
	return userInfo, nil

}

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close() // 函数栈回收的时候会调用
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("get user error: %s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("delete user error: %s", err)
	}
	_, err = stmtDel.Exec(loginName, pwd)
	defer stmtDel.Close()
	return nil
}

func GetUserInfo() ([]defs.GameList, error) {
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

	stmtOut, err := dbConn.Prepare("SELECT id, icon,name, mana FROM game WHERE del_flag = 0 " + orderBy + " LIMIT ?,?")

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

		stmtOutTag, err := dbConn.Prepare("SELECT tag_name FROM game_tag WHERE game_id=? limit 3")
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
