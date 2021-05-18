package dbops

import (
	"api/admin/defs"
	"fmt"
	"log"
	"strings"
)

// 查看所有role
func AdminRoleList() ([]*defs.Role, error) {
	roles := []*defs.Role{}
	stmtOut, err := dbConn.Prepare("SELECT id, name, description, admin_count, create_time, status, sort FROM admin_role WHERE status = 1")
	if err != nil {
		log.Printf("get roles error: %s", err)
		return roles, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		return roles, err
	}
	for rows.Next() {
		line := defs.Role{}
		err = rows.Scan(&line.Id, &line.Name, &line.Description, &line.AdminCount, &line.CreateTime, &line.Status, &line.Sort)
		if err != nil {
			return roles, err
		}
		roles = append(roles, &line)
	}
	defer stmtOut.Close()
	return roles, nil
}

// 查看用户规则获取
func AdminRoles(uid int) ([]*defs.Role, error) {
	roles := []*defs.Role{}
	stmtOut, err := dbConn.Prepare("SELECT R.id, R.name, R.description, R.admin_count, R.create_time, R.status, R.sort FROM admin_role_relation AS R2 RIGHT JOIN admin_role AS R ON R2.role_id = R.id WHERE R2.admin_id = ?")
	if err != nil {
		log.Printf("get roles error: %s", err)
		return roles, err
	}
	rows, err := stmtOut.Query(uid)
	if err != nil {
		return roles, err
	}
	for rows.Next() {
		line := defs.Role{}
		err = rows.Scan(&line.Id, &line.Name, &line.Description, &line.AdminCount, &line.CreateTime, &line.Status, &line.Sort)
		if err != nil {
			return roles, err
		}
		roles = append(roles, &line)
	}
	defer stmtOut.Close()
	return roles, nil
}

// 修改用户角色信息
func AdminRoleUpdate(uid int, ids string) error {
	// 删除当前用户角色

	stmtDel, err := dbConn.Prepare("DELETE FROM admin_role_relation WHERE admin_id = ?")
	if err != nil {
		log.Printf("delete admin_role_relation error: %s", err)
	}
	_, err = stmtDel.Exec(uid)
	defer stmtDel.Close()

	idsArr := strings.Split(ids, ",")
	for _, rid := range idsArr {
		// 增加关系
		stmtIns, err := dbConn.Prepare("INSERT INTO admin_role_relation (admin_id, role_id) VALUES(?,?)")
		if err != nil {
			fmt.Printf("insert admin role relation error: %v", err)
			return err
		}
		_, err = stmtIns.Exec(uid, rid)
		if err != nil {
			fmt.Printf("insert admin role relation exe error: %v", err)
			return err
		}
		defer stmtIns.Close() // 函数栈回收的时候会调用

	}

	return nil
}
