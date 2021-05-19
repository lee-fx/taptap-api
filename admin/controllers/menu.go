package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"api/admin/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func makeTree(Allnode []*defs.Menu, node *defs.Menu) { //参数为父节点，添加父节点的子节点指针切片
	childs, _ := haveChild(Allnode, node) //判断节点是否有子节点并返回
	if childs != nil {

		node.Children = append(node.Children, childs[0:]...) //添加子节点
		for _, v := range childs {                     //查询子节点的子节点，并添加到子节点
			_, has := haveChild(Allnode, v)
			if has {
				makeTree(Allnode, v) //递归添加节点
			}
		}
	}
}

func haveChild(Allnode []*defs.Menu, node *defs.Menu) (childs []*defs.Menu, yes bool) {
	for _, v := range Allnode {
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

	}

	makeTree(res, res[0]) //调用生成tree



	resData := &defs.NormalResponse{
		Code:    200,
		Message: "success",
		Data:    res,
	}
	utils.SendNormalResponse(w, *resData, 200)
}
