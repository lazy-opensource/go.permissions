package controller

import (
	"core/code/model/system"
	"core/code/service/system"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"lzy/framework/util"
	"lzy/framework/uuid"
	"lzy/framework/web"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (

	//------- user status ---------
	ACTIVATE = "1" //激活
	ILLEGAL  = "2" //不合法
	DISABLE  = "3" //停用
	FREEZE   = "4" //冻结

	//------------ route name  ---------------
	SYSTEM_TOLOGIN = "/system/toLogin"
	//页面Ajax
	SYSTEM_GETONEMENU = "/system/getOneMenu"
	//页面load
	SYSTEM_GETNAVIGATION = "/system/getNavigation"
	//导航菜单链接
	SYSTEM_GETCHILDMENUS = "/system/getChildMenus"
	//去首页
	SYSTEM_TOINDEX = "/system/toIndex"
	//退出
	SYSTEM_LOGOUT             = "/system/logout"
	SYSTEM_GETCURRENTMENUOPER = "/system/getCurrentMenuOper"
	SYSTEM_FUNCTION_UNDEVELOP = "/system/function/undevelop"

	//------------ html name -------------------
	PORT            = ":8688"
	DBCONFIGPATH    = "../conf/db.properties"
	LOGIN_HTML      = "login.html"
	MAIN_HTML       = "main.html"
	NAVIGATION_HTML = "common/navigation.html"
	INDEX_HTML      = "index.html"
	UNDEVELOP       = "/common/undevelop.html"
)

var dispatcher string = ""
var num int64 = 0
var userId string = ""
var dbConfig string = "../conf/db.properties"

func functionUndevelop(w http.ResponseWriter, r *http.Request) {

	web.DevEnvExecuteHtml(w, UNDEVELOP, nil)
}

//ajax get cuurent menu oper
func getCurrentMenuOper(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("parentMenuId")

	menuOper := service.GetOpersByMenuId(id)
	userOper := service.GetCurrentUserOper()

	rlt := service.CheckMenuOperPermissions(menuOper, userOper)
	util.BubbleSortMap(rlt)
	b, _ := json.Marshal(rlt)
	io.WriteString(w, string(b))

}

//adapter list
func adapterList(w http.ResponseWriter, r *http.Request) {

	aoData := r.FormValue("aoData")
	condition := r.FormValue("condition")
	fmt.Println("condition -------->", condition)
	fmt.Println("aoData --------->", aoData)

	data := make([]model.BoostrapAoData, 51)
	conditions := make([]model.BoostrapAoData, 10)

	param := make(map[string]interface{})
	param2 := make(map[string]interface{})
	json.Unmarshal([]byte(aoData), &data)
	json.Unmarshal([]byte(condition), &conditions)

	fmt.Println("conditions --------->", conditions)
	for _, v := range data {
		param[v.Name] = v.Value
	}
	for _, v := range conditions {
		param2[v.Name] = v.Value
	}
	fmt.Println("param2 +++++++++++++++", param2)

	rows := make(map[string]interface{})
	param3 := make(map[int]map[string]string)
	total := 0
	switch dispatcher {
	case "user":
		param3, total = service.GetUserListByPagin(param, param2)
	case "role":
		param3, total = service.GetRoleListByPagin(param, param2)
	case "menu":
		param3, total = service.GetMenuListByPagin(param, param2)
	case "oper":
		param3, total = service.GetOperListByPagin(param, param2)
	case "userGroup":
		param3, total = service.GetUserGroupListByPagin(param, param2)
	default:
		io.WriteString(w, "")
	}

	if len(param) > 0 {
		rows["sEcho"] = param["sEcho"].(float64)
		rows["iTotalRecords"] = total
		rows["iTotalDisplayRecords"] = total
	}

	rows["aaData"] = param3

	result, _ := json.Marshal(rows)
	io.WriteString(w, string(result))
}

//adapter edit
func adapterEdit(w http.ResponseWriter, m interface{}) {

	setUpdateTime(m)

	switch dispatcher {
	case "userEdit":
		num = service.UserEdit(m.(*model.SysUser))
	case "roleEdit":
		num = service.RoleEdit(m.(*model.SysRole))
	case "userGroupEdit":
		num = service.UserGroupEdit(m.(*model.SysUserGroup))
	case "operEdit":
		num = service.OperEdit(m.(*model.SysOper))
	case "menuEdit":
		num = service.MenuEdit(m.(*model.SysMenu))
	default:
		io.WriteString(w, "false")
	}

	if num > 0 {
		io.WriteString(w, strconv.FormatInt(num, 10))
	} else {
		io.WriteString(w, "false")
	}
}

func setCreateTimeAndUuid(m interface{}) {

	o := reflect.ValueOf(m).Elem()
	o.FieldByName("CreateTime").SetString(time.Now().Format("2006-01-02 15:04:05"))
	o.FieldByName("Id").SetString(uuid.Rand().Hex())

}

func setUuid(m interface{}) {
	reflect.ValueOf(m).Elem().FieldByName("Id").SetString(uuid.Rand().Hex())
}

func setUpdateTime(m interface{}) {
	reflect.ValueOf(m).Elem().FieldByName("UpdateTime").SetString(time.Now().Format("2006-01-02 15:04:05"))
}

//adapter edit
func adapterAdd(w http.ResponseWriter, m interface{}) {

	setCreateTimeAndUuid(m)

	switch dispatcher {
	case "userAdd":
		num = service.Add("user", m)
	case "roleAdd":
		num = service.Add("role", m)
	case "groupAdd":
		num = service.Add("group", m)
	case "operAdd":
		num = service.Add("oper", m)
	case "menuAdd":
		num = service.Add("menu", m)
	default:
		io.WriteString(w, "false")
	}

	if num > 0 {
		io.WriteString(w, strconv.FormatInt(num, 10))
	} else {
		io.WriteString(w, "false")
	}
}

//adapter del
func adapterDel(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("ids")
	if util.IsEmpty(id) {
		return
	}
	ids := strings.Split(id, ",")
	switch dispatcher {
	case "userDel":
		num = service.DelUserByIds(ids)
	case "roleDel":
		num = service.DelRoleInIds(ids)
	case "operDel":
		num = service.DelOperInIds(ids)
	case "menuDel":
		num = service.DelMenuInIds(ids)
	default:
		io.WriteString(w, "false")
	}

	if num > 0 {
		io.WriteString(w, strconv.FormatInt(num, 10))
	} else {
		io.WriteString(w, "false")
	}

}

//初始化路由控制器
func InitRoute() {
	web.SetLoginRoute(SYSTEM_TOLOGIN)
	web.SetPort(PORT)
	web.SetSessionMaxAge(30)
	web.SetDbConfig(DBCONFIGPATH)
	//------ common ------------
	web.AddHandler(SYSTEM_TOLOGIN, toLogin)
	web.AddHandler(SYSTEM_GETONEMENU, getNavigation)
	web.AddHandler(SYSTEM_GETCHILDMENUS, getChildMenusByParentId)
	web.AddHandler(SYSTEM_GETNAVIGATION, toNavigation)
	web.AddHandler(SYSTEM_TOINDEX, toIndex)
	web.AddHandler(SYSTEM_LOGOUT, logout)
	web.AddHandler(SYSTEM_GETCURRENTMENUOPER, getCurrentMenuOper)
	web.AddHandler(SYSTEM_FUNCTION_UNDEVELOP, functionUndevelop)
	//-------- user ------------
	web.AddHandler(SYSTEM_USER_TOLIST, userToList)
	web.AddHandler(SYSTEM_USER_LIST, userListByPagin)
	web.AddHandler(SYSTEM_USER_ADD, userAdd)
	web.AddHandler(SYSTEM_USER_EDIT, userEdit)
	web.AddHandler(SYSTEM_USER_DEL, userDel)
	web.AddHandler(SYSTEM_USER_CONTACTGROUP, userToContactGroup)
	web.AddHandler(SYSTEM_USER_CONFIRM_USERGROUP, userConfirmUserGroup)
	web.AddHandler(SYSTEM_GET_USER_BY_ID, getUserById)
	web.AddHandler(SYSTEM_USER_GETROLESBYGROUPID, getRoleListByGroupId)
	web.AddHandler(SYSTEM_USER_CONTACTROLE, contactRole)
	web.AddHandler(SYSTEM_USER_DETAILUSER, userDetail)
	//----------- menu -------------
	web.AddHandler(SYSTEM_MENU_TOLIST, menuToList)
	web.AddHandler(SYSTEM_MENU_LIST, menuList)
	web.AddHandler(SYSTEM_MENU_ADD, menuAdd)
	web.AddHandler(SYSTEM_MENU_EDIT, menuEdit)
	web.AddHandler(SYSTEM_MENU_DEL, menuDel)
	web.AddHandler(SYSTEM_MENU_LISTBYCONDITION, menuListByCondition)
	web.AddHandler(SYSTEM_MENU_MENUCONFIRMCONTACTPARENTMENU, menuConfirmContactParentMenu)
	web.AddHandler(SYSTEM_MENU_MENUBYID, getMenuById)
	//------------ role -----------------
	web.AddHandler(SYSTEM_ROLE_TOLIST, roleToList)
	web.AddHandler(SYSTEM_ROLE_LIST, roleListByPagin)
	web.AddHandler(SYSTEM_ROLE_ADD, roleAdd)
	web.AddHandler(SYSTEM_ROLE_EDIT, roleEdit)
	web.AddHandler(SYSTEM_ROLE_DEL, roleDel)
	web.AddHandler(SYSTEM_ROLE_GETROLEBYID, getRoleById)
	web.AddHandler(SYSTEM_ROLE_TOCONTACTOPER, roleToContactOper)
	web.AddHandler(SYSTEM_ROLE_TOCONTACTGROUP, roleToContactGroup)
	web.AddHandler(SYSTEM_ROLE_CONFIRMCONTACTGROUP, roleConfirmContactGroup)
	web.AddHandler(SYSTEM_ROLE_CONFIRMCONTACTMENU, roleConfirmContactMenu)
	web.AddHandler(SYSTEM_ROLE_ROLECONFIRMCONTACTOPER, roleConfirmContactOper)
	web.AddHandler(SYSTEM_ROLE_TOCONTACTMENU, roltToContactMenu)
	//------------ oper -------------------
	web.AddHandler(SYSTEM_OPER_TOLIST, operToList)
	web.AddHandler(SYSTEM_OPER_LIST, operListByPagin)
	web.AddHandler(SYSTEM_OPER_ADD, operAdd)
	web.AddHandler(SYSTEM_OPER_EDIT, operEdit)
	web.AddHandler(SYSTEM_OPER_DEL, operDel)
	web.AddHandler(SYSTEM_OPER_GETOPERBYID, getOperById)
	web.AddHandler(SYSTEM_OPER_CONTACTMENU, operConfirmContactMenu)
	//----------- userGroup -----------------
	web.AddHandler(SYSTEM_USERGROUP_TOLIST, userGroupToList)
	web.AddHandler(SYSTEM_USERGROUP_LIST, userGroupList)
	web.AddHandler(SYSTEM_USERGROUP_ADD, userGroupAdd)
	web.AddHandler(SYSTEM_USERGROUP_EDIT, userGroupEdit)
	web.AddHandler(SYSTEM_USERGROUP_DEL, userGroupDel)
	web.AddHandler(SYSTEM_USERGROUP_BYID, getUserGroupById)
	web.AddHandler(SYSTEM_USERGROUP_LIST2, groupList)
	web.AddHandler(SYSTEM_USERGROUP_USERGROUPTOCONTACTMENU, userGroupToContactMenu)
	web.AddHandler(SYSTEM_USERGROUP_USERGROUPCONFIRMCONTACTMENU, userGroupConfirmContactMenu)
	web.AddHandler(SYSTEM_USERGROUP_USERGROUPTOCONTACTOPER, userGroupToContactOper)
	web.AddHandler(SYSTEM_USERGROUP_USERGROUPCONFIRMCONTACTOPER, userGroupConfirmContactOper)
}

//get current user
func GetCurrentUser(w http.ResponseWriter, r *http.Request) map[int]map[string]string {
	username := r.FormValue("username")
	password := r.FormValue("password")

	m := service.GetUserByNameAndPassword(username, password)

	return m
}

//login route
func toLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := time.Now().Unix() //系统当前时间
		h := md5.New()         //进行md5
		io.WriteString(h, strconv.FormatInt(t, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		//	web.ExecuteHtml(w, TOLOGIN_FILE_NAME, nil)
		web.DevEnvExecuteHtml(w, LOGIN_HTML, token)
	} else {

		dbType, dbAdrr := service.GetDbInfo(dbConfig)

		web.GetSession().Set("DbType", dbType)
		web.GetSession().Set("DbAdrr", dbAdrr)

		userMap := GetCurrentUser(w, r)
		if userMap != nil {

			fmt.Println("login user userMap", userMap)
			userStatus := userMap[0]["Status"]
			if userStatus != ACTIVATE {
				msg := util.NewMessage()
				msg.Message = "账号不合法或停用或冻结 !"
				web.DevEnvExecuteHtml(w, LOGIN_HTML, msg)
			} else {
				userId := userMap[0]["Id"]
				web.GetSession().Set("userId", userId)
				web.GetSession().Set("username", userMap[0]["Name"])

				fmt.Println("web session ", web.GetSession())
				web.DevEnvExecuteHtml(w, MAIN_HTML, nil)
			}

		} else {
			msg := util.NewMessage()
			msg.Message = "用户名或密码错误 !"
			web.DevEnvExecuteHtml(w, LOGIN_HTML, msg)
		}

	}
}

//-----------------route----------------------------

//退出
func logout(w http.ResponseWriter, r *http.Request) {
	web.Logout(w, r)
}

//get navigation route
func toNavigation(w http.ResponseWriter, r *http.Request) {
	web.DevEnvExecuteHtml(w, NAVIGATION_HTML, nil)
}

//ajax获得一级菜单
func getNavigation(w http.ResponseWriter, r *http.Request) {

	menuIds := service.GetUserFinalMenuIds()

	util.BubbleSortSlick(menuIds)
	rlt, _ := service.GetOneMenusInIds(menuIds)
	fmt.Println("menus -------->", rlt)
	resp := make(map[string]interface{})
	resp["rows"] = rlt
	resp["name"] = web.GetSession().Get("username").(string)

	fmt.Println("login data ", resp)
	result, err := json.Marshal(resp)
	if err != nil {
		util.HandleErr(err)
	} else {
		io.WriteString(w, string(result))
	}

}

//获得子菜单 返回模板
func toIndex(w http.ResponseWriter, r *http.Request) {
	parentId := r.FormValue("parentId")
	parentName := r.FormValue("parentName")

	resp := make(map[string]interface{})
	resp["parentId"] = parentId
	resp["parentName"] = parentName
	web.DevEnvExecuteHtml(w, INDEX_HTML, resp)
}

//ajax获得子菜单
func getChildMenusByParentId(w http.ResponseWriter, r *http.Request) {

	parentId := r.FormValue("parentId")
	//parentName := r.FormValue("parentName")

	result := service.GetChildMenuByParentId(parentId)
	resp := make(map[string]interface{})
	if result != nil {
		resp["rows"] = result
	}
	//resp["parentName"] = parentName
	rows, _ := json.Marshal(resp)

	io.WriteString(w, string(rows))
	//web.DevEnvExecuteHtml(w, MENUS_HTML, resp)

}

func getCurrentUserId() string {

	userId = web.GetSession().Get("userId").(string)

	return userId

}

func getCurrentUserGroupId() string {
	userId := getCurrentUserId()

	user := service.GetUserById(userId)
	return user[0]["GroupId"]
}
