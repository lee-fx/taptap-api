package dbops

import (
	"api/admin/defs"
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
