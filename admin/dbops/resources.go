package dbops

import (
	"api/admin/defs"
	"database/sql"
	"log"
)

func ResourceListAll() ([]*defs.Resource, error) {
	resourceList := []*defs.Resource{}
	stmtResource, err := dbConn.Prepare("SELECT id, name, url, description, category_id, create_time FROM admin_resource")
	defer stmtResource.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get resource err: %s", err)
		return resourceList, err
	}
	stmtRows, err := stmtResource.Query()
	for stmtRows.Next() {
		line := &defs.Resource{}
		err = stmtRows.Scan(&line.Id, &line.Name, &line.Url, &line.Description, &line.CategoryId, &line.CreateTime)
		if err != nil {
			log.Printf("resource sql scan error: %s", err)
			return resourceList, err
		}
		resourceList = append(resourceList, line)
	}
	return resourceList, nil
}

func ResourceCategoryListAll() ([]*defs.ResourceCategory, error) {
	rcList := []*defs.ResourceCategory{}
	stmtRC, err := dbConn.Prepare("SELECT id, name, sort, create_time FROM admin_resource_category ORDER BY sort DESC")
	defer stmtRC.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get resource category err: %s", err)
		return rcList, err
	}
	stmtRows, err := stmtRC.Query()
	for stmtRows.Next() {
		line := &defs.ResourceCategory{}
		err = stmtRows.Scan(&line.Id, &line.Name, &line.Sort, &line.CreateTime)
		if err != nil {
			log.Printf("resource category sql scan error: %s", err)
			return rcList, err
		}
		rcList = append(rcList, line)
	}
	return rcList, nil
}
