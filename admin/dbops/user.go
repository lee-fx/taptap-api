package dbops

import (
	"api/admin/defs"
	"api/admin/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func VerifyUserLogin(info *defs.UserLogin) (*defs.User, error) {
	username := info.UserName
	password := utils.GetMD5HashCode([]byte(info.UserName))
	//fmt.Println(password)
	userInfo := &defs.User{}
	stmtOut, err := dbConn.Prepare("SELECT id, email FROM admin_user WHERE username = ? AND password = ? AND del_flag = 0")
	if err != nil {
		log.Printf("verify user sql error: %s", err)
		return nil, err
	}
	err = stmtOut.QueryRow(username, password).Scan(&userInfo.Id, &userInfo.Email)
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

// 获取用户列表
func AdminUserList(page, to int) (*defs.UserList, error) {

	userList := &defs.UserList{}
	userList.PageNum = page // 1
	userList.PageSize = to  // 10

	totalRow, err := dbConn.Query("SELECT COUNT(*) FROM admin_user WHERE del_flag = 0")
	if err != nil {
		fmt.Println("get total users sql error", err)
		return nil, err
	}
	total := 0
	for totalRow.Next() {
		err := totalRow.Scan(
			&total,
		)
		if err != nil {
			fmt.Println("get total users error", err)
			continue
		}
	}
	totalRow.Close()

	// 获取总页数
	maxpage := utils.GetPageLimit(total, to)
	userList.Total = total
	userList.TotalPage = maxpage

	stmtUser, err := dbConn.Prepare("SELECT id, username, password, icon, email, nickname, note, create_time, login_time, status FROM admin_user WHERE del_flag = 0 LIMIT ?,?")
	defer stmtUser.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get user info err: %s", err)
		return userList, err
	}
	stmtRows, err := stmtUser.Query(page-1, to)
	for stmtRows.Next() {
		line := &defs.User{}
		err = stmtRows.Scan(&line.Id, &line.UserName, &line.PassWord, &line.Icon, &line.Email, &line.NickName, &line.Note, &line.CreateTime, &line.LoginTime, &line.Status)
		if err != nil {
			log.Printf("users sql scan error: %s", err)
			return userList, err
		}
		userList.List = append(userList.List, line)
	}

	return userList, nil
}

// 获取用户权限信息
func GetUserInfo(uid int64) (*defs.UserInfo, error) {

	// 根据uid获取用户信息
	user := &defs.User{}
	userInfo := &defs.UserInfo{}
	stmtUser, err := dbConn.Prepare("SELECT icon,username FROM admin_user WHERE del_flag = 0 AND id = ?")
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get user info err: %s", err)

		return userInfo, err
	}
	err = stmtUser.QueryRow(uid).Scan(&user.Icon, &user.UserName)
	defer stmtUser.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("select user info err: %s", err)

		return userInfo, err
	}
	userInfo.Icon = user.Icon
	//fmt.Printf("%v", userInfo.Icon)

	stmtRoles, err := dbConn.Prepare("SELECT ROLE.id,ROLE.name FROM admin_role_relation AS R RIGHT JOIN admin_role AS ROLE ON R.role_id = ROLE.id WHERE R.admin_id = ?")
	defer stmtRoles.Close()
	if err != nil {
		log.Printf("get user roles error: %s", err)
		return userInfo, err
	}

	RoleRows, err := stmtRoles.Query(uid)
	for RoleRows.Next() {
		var id int
		var name string
		err = RoleRows.Scan(&id, &name)
		if err != nil {
			log.Printf("roles sql scan error: %s", err)
			return userInfo, err
		}
		userInfo.Roles = append(userInfo.Roles, name)

		//fmt.Printf("%v", userInfo.Roles)

		stmtMenus, err := dbConn.Prepare("SELECT M.id, M.parent_id, M.create_time, M.title, M.level, M.sort, M.name, M.icon, M.hidden FROM admin_role_menu_relation AS R RIGHT JOIN admin_menu AS M ON R.menu_id = M.id WHERE R.role_id = ?")
		if err != nil {
			log.Printf("get user menus error: %s", err)
			return userInfo, err
		}

		menuRows, err := stmtMenus.Query(&id)
		for menuRows.Next() {
			menu := defs.Menu{}
			err = menuRows.Scan(&menu.Id, &menu.ParentId, &menu.CreateTime, &menu.Title, &menu.Level, &menu.Sort, &menu.Name, &menu.Icon, &menu.Hidden)
			if err != nil {
				log.Printf("menus sql scan error: %s", err)
				return userInfo, err
			}
			userInfo.Menus = append(userInfo.Menus, menu)
		}
		//gameTag = []defs.GameTag{}

	}

	return userInfo, nil

}
