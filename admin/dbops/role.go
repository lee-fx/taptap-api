package dbops

import (
	"api/admin/defs"
	"log"
)

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
