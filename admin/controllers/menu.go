package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"api/admin/utils"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// 目录树
func makeTree(Menu []*defs.Menu, node int) []*defs.Menu {
	MenuR := []*defs.Menu{} // 创建返回的数组
	for _, v := range Menu {
		childs, _ := haveChild(Menu, v) // 判断节点是否有子节点并返回
		if childs != nil {
			v.Children = append(v.Children, childs[0:]...) // 添加子节点
			for _, v := range childs {                     //查询子节点的子节点，并添加到子节点
				_, has := haveChild(Menu, v)
				if has {
					makeTree(Menu, v.ParentId) //递归添加节点
				}
			}
			MenuR = append(MenuR, v)
		}
	}
	return MenuR
}

func haveChild(Menu []*defs.Menu, node *defs.Menu) (childs []*defs.Menu, yes bool) {
	for _, v := range Menu {
		if v.ParentId == node.Id {
			childs = append(childs, v)
		}
	}
	if childs != nil {
		yes = true
	}
	return
}

// 获取菜单tree列表
func MenuTreeList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := dbops.MenuTreeList()
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	if len(res) != 0 {
		res = makeTree(res, 0) //调用生成tree
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "success",
		Data:    res,
	}
	utils.SendNormalResponse(w, *resData, 200)
}

// 通过pid获取菜单列表
func GetMenuListByPid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	pid, _ := strconv.Atoi(p.ByName("pid"))

	//获取请求参数
	params := r.URL.Query()
	page, _ := strconv.Atoi(params["pageNum"][0])
	to, _ := strconv.Atoi(params["pageSize"][0])

	res, err := dbops.GetMenuListByPid(page, to, pid)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "查询成功",
		Data:    res,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 菜单显示修改
func MenuUpdateHidden(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取用户id
	mid, _ := strconv.Atoi(p.ByName("mid"))
	// 状态
	params := r.URL.Query()
	hidden, _ := strconv.Atoi(params["hidden"][0])

	if err := dbops.MenuUpdateHidden(mid, hidden); err != nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "修改成功",
		Data:    nil,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 创建menu
func MenuCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.Menu{}
	if err := json.Unmarshal([]byte(res), ubody); err != nil {
		fmt.Printf("%v", err)
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.MenuCreate(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	resData := &defs.NormalResponse{
		Code:    200,
		Message: "添加菜单成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 201)
}

// 通过id删除菜单
func MenuDeleteByMid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mid, _ := strconv.Atoi(p.ByName("mid"))
	err := dbops.MenuDeleteByMid(mid)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "删除成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 200)
}

// 通过mid获取菜单信息
func GetMenuInfoById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 用户id
	id, _ := strconv.Atoi(p.ByName("id"))
	res, err := dbops.GetMenuInfoById(id)
	if err != nil {
		log.Printf("error: %v ", err)
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "操作成功",
		Data:    res,
	}
	utils.SendNormalResponse(w, *tmp, 200)
}

// 修改菜单
func MenuUpdateByMid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.Menu{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.MenuUpdateByMid(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	resData := &defs.NormalResponse{
		Code:    200,
		Message: "修改菜单成功",
		Data:    nil,
	}

	utils.SendNormalResponse(w, *resData, 200)
}
