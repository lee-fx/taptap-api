package dbops

import (
	"api/admin/defs"
	"api/admin/utils"
	"database/sql"
	"fmt"
	"log"
	"strconv"
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

// 获取资源列表
func ResourceList(page, to int, name, url string, id int) (*defs.ResourceList, error) {
	resourceList := &defs.ResourceList{}
	resourceList.PageNum = page // 1
	resourceList.PageSize = to  // 10

	whereName := ""
	if name == "" {
		whereName = "'%'"
	} else {
		whereName = "'%" + name + "%'"
	}

	whereUrl := ""
	if url == "" {
		whereUrl = "'%'"
	} else {
		whereUrl = "'%" + url + "%'"
	}

	whereCateGoryId := ""
	if id != 0 {
		//println(id)
		whereCateGoryId = " AND category_id = " + strconv.Itoa(id)
	}

	sqlQuery := "SELECT COUNT(*) FROM admin_resource WHERE name like " + whereName + " AND url like " + whereUrl + whereCateGoryId
	//fmt.Println(sqlQuery)
	totalRow, err := dbConn.Query(sqlQuery)
	if err != nil {
		fmt.Println("get total resource sql error", err)
		return nil, err
	}
	total := 0
	for totalRow.Next() {
		err := totalRow.Scan(
			&total,
		)
		if err != nil {
			fmt.Println("get total resource error", err)
			continue
		}
	}
	totalRow.Close()

	// 获取总页数
	maxpage := utils.GetPageLimit(total, to)
	resourceList.Total = total
	resourceList.TotalPage = maxpage

	stmtRole, err := dbConn.Prepare("SELECT id, name, url, description, category_id, create_time FROM admin_resource WHERE name like " + whereName + " AND url like " + whereUrl + whereCateGoryId + " LIMIT ?,?")
	defer stmtRole.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Printf("get resource info err: %s", err)
		return resourceList, err
	}
	stmtRows, err := stmtRole.Query((page-1)*to, to)
	for stmtRows.Next() {
		line := &defs.Resource{}
		err = stmtRows.Scan(&line.Id, &line.Name, &line.Url, &line.Description, &line.CategoryId, &line.CreateTime)
		if err != nil {
			log.Printf("resourse sql scan error: %s", err)
			return resourceList, err
		}
		resourceList.List = append(resourceList.List, line)
	}

	return resourceList, nil
}

// 新建资源
func ResourceCreate(resource *defs.Resource) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO admin_resource (name, url, description, category_id, create_time) VALUES(?,?,?,?,?)")
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	timeNow := utils.GetTimeNowFormatDate()

	_, err = stmtIns.Exec(resource.Name, resource.Url, resource.Description, resource.CategoryId, timeNow)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 修改资源
func ResourUpdateById(resource *defs.Resource) error {

	stmtUpdate, err := dbConn.Prepare("UPDATE admin_resource SET name=?, url=?, description=?, category_id=?, create_time=? WHERE id = ?")
	if err != nil {
		fmt.Printf("update resource error: %v", err)
		return err
	}

	_, err = stmtUpdate.Exec(resource.Name, resource.Url, resource.Description, resource.CategoryId, resource.CreateTime, resource.Id)
	if err != nil {
		fmt.Printf("update resource exec error: %v", err)
		return err
	}
	defer stmtUpdate.Close()
	return nil
}

// 删除资源
func ResourDeleteById(id int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM admin_resource WHERE id=?")
	if err != nil {
		log.Printf("delete admin_resource error: %s", err)
	}
	_, err = stmtDel.Exec(id)
	defer stmtDel.Close()
	return nil
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

// 资源分类创建
func ResourceCategoryCreate(rc *defs.ResourceCategory) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO admin_resource_category (name, sort, create_time) VALUES(?,?,?)")
	if err != nil {
		fmt.Printf("ins resource category error: %v", err)
		return err
	}
	timeNow := utils.GetTimeNowFormatDate()
	_, err = stmtIns.Exec(rc.Name, rc.Sort, timeNow)
	if err != nil {
		fmt.Printf("ins resource category exe error: %v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

// id删除资源分类
func ResourceCategoryDeleteById(id int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM admin_resource_category WHERE id=?")
	if err != nil {
		log.Printf("delete  resource category error: %s", err)
	}
	_, err = stmtDel.Exec(id)
	defer stmtDel.Close()
	return nil
}

// 修改资源分类
func ResourceCategoryUpdateById(rc *defs.ResourceCategory) error {

	stmtUpdate, err := dbConn.Prepare("UPDATE admin_resource_category SET name=?, sort=?, create_time=?  WHERE id = ?")
	if err != nil {
		fmt.Printf("update resource category error: %v", err)
		return err
	}

	_, err = stmtUpdate.Exec(rc.Name, rc.Sort, rc.CreateTime, rc.Id)
	if err != nil {
		fmt.Printf("update resource category exec error: %v", err)
		return err
	}
	defer stmtUpdate.Close()
	return nil
}
