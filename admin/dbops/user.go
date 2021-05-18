package dbops

import (
	"api/admin/defs"
	"api/admin/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 登录校验
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

// 新建用户
func AdminUserRegister(user *defs.User) error {
	password := utils.GetMD5HashCode([]byte(user.PassWord))
	stmtIns, err := dbConn.Prepare("INSERT INTO admin_user (username, nickname, password, email, note, status, create_time, login_time) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	timeNow := utils.GetTimeNowFormatDate()
	_, err = stmtIns.Exec(user.UserName, user.NickName, password, user.Email, user.Note, user.Status, timeNow, timeNow)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	defer stmtIns.Close() // 函数栈回收的时候会调用
	return nil
}

// 查看用户名是否存在
func GetUserByUserName(username string, uid int64) bool {
	whereTem := ""
	if uid != 99999 {
		whereTem = " AND id != ?"
	}
	stmtOut, err := dbConn.Prepare("SELECT id FROM admin_user WHERE username = ? " + whereTem)
	if err != nil {
		log.Printf("get userByUserName error: %s", err)
		return true
	}
	var id string
	err = stmtOut.QueryRow(username, uid).Scan(&id)
	if err != nil {
		log.Printf("get userByUserName scan error: %s", err)
		return true
	}

	if err == sql.ErrNoRows {
		return true
	}
	stmtOut.Close()

	return false
}

// 用户删除
func AdminUserDelete(id int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM admin_user WHERE id=?")
	if err != nil {
		log.Printf("delete user error: %s", err)
	}
	_, err = stmtDel.Exec(id)
	defer stmtDel.Close()
	return nil
}

// 获取用户列表
func AdminUserList(page, to int, keyword string) (*defs.UserList, error) {

	userList := &defs.UserList{}
	userList.PageNum = page // 1
	userList.PageSize = to  // 10

	// 多字段模糊查询构造
	whereKeyWord := ""
	if keyword == "" {
		whereKeyWord = "%"
	} else {
		whereKeyWord = "%" + keyword + "%"
	}

	totalRow, err := dbConn.Query("SELECT COUNT(*) FROM admin_user WHERE CONCAT(username, nickname) like '" + whereKeyWord + "' AND del_flag = 0")
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

	// 模糊查询 CONCAT
	stmtUser, err := dbConn.Prepare("SELECT id, username, password, icon, email, nickname, note, create_time, login_time, status FROM admin_user WHERE CONCAT(username, nickname) like '" + whereKeyWord + "' AND del_flag = 0 LIMIT ?,?")
	defer stmtUser.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get user info err: %s", err)
		return userList, err
	}
	stmtRows, err := stmtUser.Query((page-1)*to, to)
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

// 修改用户状态
func AdminUpdateUserStatus(uid, status int) error {
	stmtUpdate, err := dbConn.Prepare("UPDATE admin_user SET status = ? WHERE id = ?")
	if err != nil {
		log.Printf("update user status error: %s", err)
	}
	_, err = stmtUpdate.Exec(status, uid)
	defer stmtUpdate.Close()
	return nil
}

// 通过id获取用户密码
func GetUserPwdById(uid string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT password FROM admin_user WHERE id = ?")
	if err != nil {
		log.Printf("get user pwd error: %s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(uid).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get user pwd scan error: %s", err)
		return "", err
	}

	stmtOut.Close()

	return pwd, nil
}

// 修改用户信息
func AdminUpdateUser(user *defs.User) error {
	pwd, err := GetUserPwdById(string(user.Id))
	if err != nil {
		fmt.Printf("get pwd err: %v", err)
		return err
	}
	password := ""
	if pwd == user.PassWord {
		password = user.PassWord
	} else {
		password = utils.GetMD5HashCode([]byte(user.PassWord))

	}
	stmtUpdate, err := dbConn.Prepare("UPDATE admin_user SET username=?, nickname=?, password=?, email=?, note=?, status=?, create_time=?, login_time=? WHERE id = ?")
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	_, err = stmtUpdate.Exec(user.UserName, user.NickName, password, user.Email, user.Note, user.Status, user.CreateTime, user.LoginTime, user.Id)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	defer stmtUpdate.Close() // 函数栈回收的时候会调用
	return nil
}
