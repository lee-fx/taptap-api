package dbops

import (
	"api/admin/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	mysql_name := config.GetMysqlName()
	mysql_password := config.GetMysqlPassword()
	mysql_host := config.GetMysqlHost()
	mysql_port := config.GetMysqlPort()
	mysql_admin_db := config.GetMysqlAdminDb()

	//	dbConn, err = sql.Open("mysql", "root:lx123321@tcp(localhost:3306)/taptap?charset=UTF8")

	dbConn, err = sql.Open("mysql", mysql_name+":"+mysql_password+"@tcp("+mysql_host+":"+mysql_port+")/"+mysql_admin_db+"?charset=UTF8")
	if err != nil {
		panic(err.Error())
	}
}
