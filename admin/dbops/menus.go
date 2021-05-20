package dbops

import (
	"api/admin/defs"
	"api/admin/utils"
	"database/sql"
	"fmt"
	"log"
)

// 查询菜单树列表
func MenuTreeList() ([]*defs.Menu, error) {
	menus := []*defs.Menu{}
	stmtOut, err := dbConn.Prepare("SELECT id, parent_id, title, level, sort, name, icon, hidden, create_time FROM admin_menu WHERE hidden = ?")
	if err != nil {
		log.Printf("get menus error: %s", err)
		return menus, err
	}

	hidden := 0

	rows, err := stmtOut.Query(hidden)
	if err != nil {
		return menus, err
	}
	for rows.Next() {
		line := defs.Menu{}
		err = rows.Scan(&line.Id, &line.ParentId, &line.Title, &line.Level, &line.Sort, &line.Name, &line.Icon, &line.Hidden, &line.CreateTime)
		if err != nil {
			return menus, err
		}
		menus = append(menus, &line)
	}

	defer stmtOut.Close()
	return menus, nil
}

// 获取相应pid的菜单列表
func GetMenuListByPid(page, to int, pid int) (*defs.MenuList, error) {
	menuList := &defs.MenuList{}
	menuList.PageNum = page // 1
	menuList.PageSize = to  // 5

	totalRow, err := dbConn.Query("SELECT COUNT(*) FROM admin_menu WHERE parent_id = ?", pid)
	if err != nil {
		fmt.Println("get total menus sql error", err)
		return nil, err
	}
	total := 0
	for totalRow.Next() {
		err := totalRow.Scan(
			&total,
		)
		if err != nil {
			fmt.Println("get total menus error", err)
			continue
		}
	}
	totalRow.Close()

	// 获取总页数
	maxpage := utils.GetPageLimit(total, to)
	menuList.Total = total
	menuList.TotalPage = maxpage

	stmtMenu, err := dbConn.Prepare("SELECT id, parent_id, title, level, sort, name, icon, hidden, create_time FROM admin_menu WHERE parent_id = ? LIMIT ?,?")
	defer stmtMenu.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get role info err: %s", err)
		return menuList, err
	}
	stmtRows, err := stmtMenu.Query(pid, (page-1)*to, to)
	for stmtRows.Next() {
		line := &defs.Menu{}
		err = stmtRows.Scan(&line.Id, &line.ParentId, &line.Title, &line.Level, &line.Sort, &line.Name, &line.Icon, &line.Hidden, &line.CreateTime)
		if err != nil {
			log.Printf("menus sql scan error: %s", err)
			return menuList, err
		}
		menuList.List = append(menuList.List, line)
	}

	return menuList, nil
}

// 修改菜单显示状态
func MenuUpdateHidden(mid, hidden int) error {
	stmtUpdate, err := dbConn.Prepare("UPDATE admin_menu SET hidden = ? WHERE id = ?")
	if err != nil {
		log.Printf("update menu hidden error: %s", err)
	}
	_, err = stmtUpdate.Exec(hidden, mid)
	defer stmtUpdate.Close()
	return nil
}

// 创建菜单
func MenuCreate(menu *defs.Menu) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO admin_menu (parent_id, title, level, sort, name, icon, hidden, create_time) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	timeNow := utils.GetTimeNowFormatDate()
	level_v := 0
	if menu.ParentId != 0 {
		level_v = 1
	}
	_, err = stmtIns.Exec(menu.ParentId, menu.Title, level_v, menu.Sort, menu.Name, menu.Icon, menu.Hidden, timeNow)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 通过mid删除菜单
func MenuDeleteByMid(mid int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM admin_menu WHERE id=?")
	if err != nil {
		log.Printf("delete user error: %s", err)
	}
	_, err = stmtDel.Exec(mid)
	defer stmtDel.Close()
	return nil
}

// 通过id查询菜单信息
func GetMenuInfoById(id int) (*defs.Menu, error) {
	menu := &defs.Menu{}
	stmtOut, err := dbConn.Prepare("SELECT id, parent_id, title, level, sort, name, icon, hidden, create_time FROM admin_menu WHERE id = ?")
	if err != nil {
		log.Printf("get menu error: %s", err)
		return menu, err
	}
	err = stmtOut.QueryRow(id).Scan(&menu.Id, &menu.ParentId, &menu.Title, &menu.Level, &menu.Sort, &menu.Name, &menu.Icon, &menu.Hidden, &menu.CreateTime)
	if err != nil {
		log.Printf("get menu scan error: %s", err)
		return menu, err
	}

	if err == sql.ErrNoRows {
		return menu, err
	}
	stmtOut.Close()

	return menu, nil
}

// 修改菜单
func MenuUpdateByMid(menu *defs.Menu) error {

	stmtUpdate, err := dbConn.Prepare("UPDATE admin_menu SET parent_id=?, title=?, level=?, sort=?, name=?, icon=?, hidden=?, create_time=? WHERE id = ?")
	if err != nil {
		fmt.Printf("update menu error: %v", err)
		return err
	}

	_, err = stmtUpdate.Exec(menu.ParentId, menu.Title, menu.Level, menu.Sort, menu.Name, menu.Icon, menu.Hidden, menu.CreateTime, menu.Id)
	if err != nil {
		fmt.Printf("update menu exec error: %v", err)
		return err
	}
	defer stmtUpdate.Close()
	return nil
}
