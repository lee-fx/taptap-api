package dbops

import (
	"api/admin/defs"
	"api/admin/utils"
	"fmt"
	"log"
	"strings"
)

// 查看所有role
func AdminRoleListAll() ([]*defs.Role, error) {
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

	if ids == "" {
		return nil
	}

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
		defer stmtIns.Close()

	}

	return nil
}

// 修改角色状态
func RoleUpdateStatus(rid, status int) error {
	stmtUpdate, err := dbConn.Prepare("UPDATE admin_role SET status = ? WHERE id = ?")
	if err != nil {
		log.Printf("update role status error: %s", err)
	}
	_, err = stmtUpdate.Exec(status, rid)
	defer stmtUpdate.Close()
	return nil
}

// 删除角色
func RoleDelete(id int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM admin_role WHERE id=?")
	if err != nil {
		log.Printf("delete role error: %s", err)
	}
	_, err = stmtDel.Exec(id)
	defer stmtDel.Close()
	return nil
}

// 创建角色
func RoleCreate(role *defs.Role) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO admin_role (name , description, admin_count, create_time, status, sort) VALUES(?,?,?,?,?,?)")
	if err != nil {
		fmt.Printf("ins role error: %v", err)
		return err
	}
	timeNow := utils.GetTimeNowFormatDate()
	_, err = stmtIns.Exec(role.Name, role.Description, role.AdminCount, timeNow, role.Status, role.Sort)
	if err != nil {
		fmt.Printf("ins role exe error: %v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 修改角色
func RoleUpdate(role *defs.Role) error {

	stmtUpdate, err := dbConn.Prepare("UPDATE admin_role SET name=?, description=?, admin_count=?, create_time=?, status=?, sort=? WHERE id = ?")
	if err != nil {
		fmt.Printf("update role error: %v", err)
		return err
	}

	_, err = stmtUpdate.Exec(role.Name, role.Description, role.AdminCount, role.CreateTime, role.Status, role.Sort, role.Id)
	if err != nil {
		fmt.Printf("update role exec error: %v", err)
		return err
	}
	defer stmtUpdate.Close()
	return nil
}

// 获取角色资源
func RoleResourceByRoleId(rid int) ([]*defs.Resource, error) {
	resources := []*defs.Resource{}
	stmtOut, err := dbConn.Prepare("SELECT R.id, R.name, R.url, R.description, R.category_id, R.create_time FROM admin_role_resource_relation AS R2 RIGHT JOIN admin_resource AS R ON R2.resource_id = R.id WHERE R2.role_id = ?")
	if err != nil {
		log.Printf("get resources by roleid error: %s", err)
		return resources, err
	}
	rows, err := stmtOut.Query(rid)
	if err != nil {
		return resources, err
	}
	for rows.Next() {
		line := defs.Resource{}
		err = rows.Scan(&line.Id, &line.Name, &line.Url, &line.Description, &line.CategoryId, &line.CreateTime)
		if err != nil {
			return resources, err
		}
		resources = append(resources, &line)
	}
	defer stmtOut.Close()
	return resources, nil
}

// 角色分配权限
func RoleAllocResource(rid int, ids string) error {
	// 删除当前用户角色
	stmtDel, err := dbConn.Prepare("DELETE FROM admin_role_resource_relation WHERE role_id = ?")
	if err != nil {
		log.Printf("delete admin_role_resource_relation by rid error: %s", err)
	}
	_, err = stmtDel.Exec(rid)
	defer stmtDel.Close()

	if ids == "" {
		return nil
	}

	idsArr := strings.Split(ids, ",")
	for _, res_id := range idsArr {
		// 增加关系
		stmtIns, err := dbConn.Prepare("INSERT INTO admin_role_resource_relation (role_id, resource_id) VALUES(?,?)")
		if err != nil {
			fmt.Printf("insert admin role resource relation error: %v", err)
			return err
		}
		_, err = stmtIns.Exec(rid, res_id)
		if err != nil {
			fmt.Printf("insert admin role resource relation exe error: %v", err)
			return err
		}
		defer stmtIns.Close()
	}
	return nil
}

// 获取角色菜单
func RoleListMenuByRid(rid int) ([]*defs.Menu, error) {
	menus := []*defs.Menu{}
	stmtOut, err := dbConn.Prepare("SELECT M.id, M.parent_id, M.title, M.sort, M.Name, M.icon, M.hidden, M.create_time FROM admin_role_menu_relation AS R RIGHT JOIN admin_menu AS M ON R.menu_id = M.id WHERE R.role_id = ?")
	if err != nil {
		log.Printf("get Role list menu by rid error: %s", err)
		return menus, err
	}
	rows, err := stmtOut.Query(rid)
	if err != nil {
		return menus, err
	}
	for rows.Next() {
		line := defs.Menu{}
		err = rows.Scan(&line.Id, &line.ParentId, &line.Title, &line.Sort, &line.Name, &line.Icon, &line.Hidden, &line.CreateTime)
		if err != nil {
			return menus, err
		}
		menus = append(menus, &line)
	}
	defer stmtOut.Close()
	return menus, nil
}

// 分配角色菜单
func RoleAllocMenu(rid int, ids string) error {
	// 删除当前用户角色
	stmtDel, err := dbConn.Prepare("DELETE FROM admin_role_menu_relation WHERE role_id = ?")
	if err != nil {
		log.Printf("delete admin_role_menu_relation by rid error: %s", err)
	}
	_, err = stmtDel.Exec(rid)
	defer stmtDel.Close()

	if ids == "" {
		return nil
	}

	idsArr := strings.Split(ids, ",")
	for _, m_id := range idsArr {
		// 增加关系
		stmtIns, err := dbConn.Prepare("INSERT INTO admin_role_menu_relation (role_id, menu_id) VALUES(?,?)")
		if err != nil {
			fmt.Printf("insert admin role menu relation error: %v", err)
			return err
		}
		_, err = stmtIns.Exec(rid, m_id)
		if err != nil {
			fmt.Printf("insert admin role menu relation exe error: %v", err)
			return err
		}
		defer stmtIns.Close()
	}
	return nil
}