package controllers

import (
	"api/admin/dbops"
	"api/admin/defs"
	"fmt"
	//"api/admin/session"
	"api/admin/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserLogin{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.UserName, ubody.PassWord); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	//id := session.GenerateNewSesssionId(ubody.Username)
	//su := &defs.SignedUp{
	//	Success:   true,
	//	SessionId: id,
	//}
	//
	//if resp, err := json.Marshal(su); err != nil {
	//	utils.SendErrorResponse(w, defs.ErrorInternalFaults)
	//	return
	//} else {
	//	utils.SendNormalResponse(w, string(resp), 201)
	//}
}

func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

func AdminUserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 获取请求参数
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserLogin{}
	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 数据库校验用户名密码
	if res, err := dbops.VerifyUserLogin(ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	} else if res{
		utils.SendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}


	fmt.Printf("%v", ubody)

	//if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
	//	utils.SendErrorResponse(w, defs.ErrorDBError)
	//	return
	//}




	// 生成jwt-token并返回


	data := defs.UserToken{
		TokenHead: "Bearer ",
		Token:     "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJhZG1pbiIsImNyZWF0ZWQiOjE2MjA3ODgwMzg3MzgsImV4cCI6MTYyMTM5MjgzOH0.GsmYYQRW86duj3W_1RZfN8SBo0KPKP52JysgyuH3VMa9bNVVk8n4bw34tVY6D7zUDnUQMQlajhQhsCxHJOkXgQ",
	}
	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "登陆成功",
		Data:    data,
	}

	utils.SendNormalResponse(w, *tmp, 200)
}

// {"code":200,"message":"操作成功","data":{"roles":["超级管理员"],"icon":"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg.jpg","menus":[{"id":1,"parentId":0,"createTime":"2020-02-02T06:50:36.000+00:00","title":"商品","level":0,"sort":0,"name":"pms","icon":"product","hidden":0},{"id":2,"parentId":1,"createTime":"2020-02-02T06:51:50.000+00:00","title":"商品列表","level":1,"sort":0,"name":"product","icon":"product-list","hidden":0},{"id":3,"parentId":1,"createTime":"2020-02-02T06:52:44.000+00:00","title":"添加商品","level":1,"sort":0,"name":"addProduct","icon":"product-add","hidden":0},{"id":4,"parentId":1,"createTime":"2020-02-02T06:53:51.000+00:00","title":"商品分类","level":1,"sort":0,"name":"productCate","icon":"product-cate","hidden":0},{"id":5,"parentId":1,"createTime":"2020-02-02T06:54:51.000+00:00","title":"商品类型","level":1,"sort":0,"name":"productAttr","icon":"product-attr","hidden":0},{"id":6,"parentId":1,"createTime":"2020-02-02T06:56:29.000+00:00","title":"品牌管理","level":1,"sort":0,"name":"brand","icon":"product-brand","hidden":0},{"id":7,"parentId":0,"createTime":"2020-02-02T08:54:07.000+00:00","title":"订单","level":0,"sort":0,"name":"oms","icon":"order","hidden":0},{"id":8,"parentId":7,"createTime":"2020-02-02T08:55:18.000+00:00","title":"订单列表","level":1,"sort":0,"name":"order","icon":"product-list","hidden":0},{"id":9,"parentId":7,"createTime":"2020-02-02T08:56:46.000+00:00","title":"订单设置","level":1,"sort":0,"name":"orderSetting","icon":"order-setting","hidden":0},{"id":10,"parentId":7,"createTime":"2020-02-02T08:57:39.000+00:00","title":"退货申请处理","level":1,"sort":0,"name":"returnApply","icon":"order-return","hidden":0},{"id":11,"parentId":7,"createTime":"2020-02-02T08:59:40.000+00:00","title":"退货原因设置","level":1,"sort":0,"name":"returnReason","icon":"order-return-reason","hidden":0},{"id":12,"parentId":0,"createTime":"2020-02-04T08:18:00.000+00:00","title":"营销","level":0,"sort":0,"name":"sms","icon":"sms","hidden":0},{"id":13,"parentId":12,"createTime":"2020-02-04T08:19:22.000+00:00","title":"秒杀活动列表","level":1,"sort":0,"name":"flash","icon":"sms-flash","hidden":0},{"id":14,"parentId":12,"createTime":"2020-02-04T08:20:16.000+00:00","title":"优惠券列表","level":1,"sort":0,"name":"coupon","icon":"sms-coupon","hidden":0},{"id":16,"parentId":12,"createTime":"2020-02-07T08:22:38.000+00:00","title":"品牌推荐","level":1,"sort":0,"name":"homeBrand","icon":"product-brand","hidden":0},{"id":17,"parentId":12,"createTime":"2020-02-07T08:23:14.000+00:00","title":"新品推荐","level":1,"sort":0,"name":"homeNew","icon":"sms-new","hidden":0},{"id":18,"parentId":12,"createTime":"2020-02-07T08:26:38.000+00:00","title":"人气推荐","level":1,"sort":0,"name":"homeHot","icon":"sms-hot","hidden":0},{"id":19,"parentId":12,"createTime":"2020-02-07T08:28:16.000+00:00","title":"专题推荐","level":1,"sort":0,"name":"homeSubject","icon":"sms-subject","hidden":0},{"id":20,"parentId":12,"createTime":"2020-02-07T08:28:42.000+00:00","title":"广告列表","level":1,"sort":0,"name":"homeAdvertise","icon":"sms-ad","hidden":0},{"id":21,"parentId":0,"createTime":"2020-02-07T08:29:13.000+00:00","title":"权限","level":0,"sort":0,"name":"ums","icon":"ums","hidden":0},{"id":22,"parentId":21,"createTime":"2020-02-07T08:29:51.000+00:00","title":"用户列表","level":1,"sort":0,"name":"admin","icon":"ums-admin","hidden":0},{"id":23,"parentId":21,"createTime":"2020-02-07T08:30:13.000+00:00","title":"角色列表","level":1,"sort":0,"name":"role","icon":"ums-role","hidden":0},{"id":24,"parentId":21,"createTime":"2020-02-07T08:30:53.000+00:00","title":"菜单列表","level":1,"sort":0,"name":"menu","icon":"ums-menu","hidden":0},{"id":25,"parentId":21,"createTime":"2020-02-07T08:31:13.000+00:00","title":"资源列表","level":1,"sort":0,"name":"resource","icon":"ums-resource","hidden":0}],"username":"admin"}}

func AdminUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//uname := r.PostFormValue("username")
	//pwd := r.PostFormValue("password")
	//fmt.Println(uname)
	//fmt.Println(pwd)
	//io.WriteString(w, uname)
	//io.WriteString(w, pwd)
	//Authorization := r.Header.Get("Authorization")
	//log.Println(Authorization)

	roles := []string{"超级管理员"}
	menus := []defs.Menu{}
	menu1 := defs.Menu{
		Id:         1,
		ParentId:   0,
		CreateTime: "2020-02-02T06:50:36.000+00:00",
		Title:      "商品",
		Level:      0,
		Sort:       0,
		Name:       "pms",
		Icon:       "product",
		Hidden:     0,
	}

	menus = append(menus, menu1)

	menu2 := defs.Menu{
		Id:         2,
		ParentId:   1,
		CreateTime: "2020-02-02T06:51:50.000+00:00",
		Title:      "商品列表",
		Level:      1,
		Sort:       0,
		Name:       "product",
		Icon:       "product-list",
		Hidden:     0,
	}

	menus = append(menus, menu2)

	menu3 := defs.Menu{
		Id:         3,
		ParentId:   1,
		CreateTime: "2020-02-02T06:51:50.000+00:00",
		Title:      "添加商品",
		Level:      1,
		Sort:       0,
		Name:       "addProduct",
		Icon:       "product-add",
		Hidden:     0,
	}

	menus = append(menus, menu3)

	menu4 := defs.Menu{
		Id:         4,
		ParentId:   1,
		CreateTime: "2020-02-02T06:51:50.000+00:00",
		Title:      "商品分类",
		Level:      1,
		Sort:       0,
		Name:       "productCate",
		Icon:       "product-cate",
		Hidden:     0,
	}

	menus = append(menus, menu4)

	menu5 := defs.Menu{
		Id:         21,
		ParentId:   0,
		CreateTime: "2020-02-02T06:51:50.000+00:00",
		Title:      "权限",
		Level:      0,
		Sort:       0,
		Name:       "ums",
		Icon:       "ums",
		Hidden:     0,
	}

	menus = append(menus, menu5)

	menu6 := defs.Menu{
		Id:         22,
		ParentId:   21,
		CreateTime: "2020-02-02T06:51:50.000+00:00",
		Title:      "用户列表",
		Level:      1,
		Sort:       0,
		Name:       "admin",
		Icon:       "ums-admin",
		Hidden:     0,
	}

	menus = append(menus, menu6)

	menu7 := defs.Menu{
		Id:         23,
		ParentId:   21,
		CreateTime: "2020-02-02T06:51:50.000+00:00",
		Title:      "角色列表",
		Level:      1,
		Sort:       0,
		Name:       "role",
		Icon:       "ums-role",
		Hidden:     0,
	}

	menus = append(menus, menu7)

	menu8 := defs.Menu{
		Id:         24,
		ParentId:   21,
		CreateTime: "2020-02-02T06:51:50.000+00:00",
		Title:      "菜单列表",
		Level:      1,
		Sort:       0,
		Name:       "menu",
		Icon:       "ums-menu",
		Hidden:     0,
	}

	menus = append(menus, menu8)

	menu9 := defs.Menu{
		Id:         25,
		ParentId:   21,
		CreateTime: "2020-02-02T06:51:50.000+00:00",
		Title:      "资源列表",
		Level:      1,
		Sort:       0,
		Name:       "resource",
		Icon:       "ums-resource",
		Hidden:     0,
	}

	menus = append(menus, menu9)

	data := defs.UserInfo{
		Roles: roles,
		Icon:  "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg.jpg",
		Menus: menus,
	}

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "路由信息",
		Data:    data,
	}

	utils.SendNormalResponse(w, *tmp, 200)
}

func AdminUserLogout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	tmp := &defs.NormalResponse{
		Code:    200,
		Message: "登出成功",
		Data:    "",
	}

	utils.SendNormalResponse(w, *tmp, 200)
}
