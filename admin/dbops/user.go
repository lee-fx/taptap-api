package dbops

import (
	"api/admin/defs"
	"api/admin/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func VerifyUserLogin(info *defs.UserLogin) (bool, error) {
	username := info.UserName
	password := utils.GetMD5HashCode([]byte(info.UserName))
	stmtOut, err := dbConn.Prepare("SELECT count(*) FROM admin_user WHERE username = ? AND password = ?")
	if err != nil {
		log.Printf("verify user error: %s", err)
		return false, err
	}
	var id string
	err = stmtOut.QueryRow(username, password).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return true, nil
	}
	stmtOut.Close()
	return false, nil

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
