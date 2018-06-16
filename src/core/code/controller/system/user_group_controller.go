package controller

import (
	"core/code/model/system"
	"core/code/service/system"
	"encoding/json"
	"io"
	"lzy/framework/util"
	"lzy/framework/web"
	"net/http"
	"strconv"
	"strings"
)

const (
	//------------ route name  ---------------
	SYSTEM_USERGROUP_TOLIST                      = "/system/userGroup/toList"
	SYSTEM_USERGROUP_LIST                        = "/system/userGroup/list"
	SYSTEM_USERGROUP_ADD                         = "/system/userGroup/add"
	SYSTEM_USERGROUP_EDIT                        = "/system/userGroup/edit"
	SYSTEM_USERGROUP_DEL                         = "/system/userGroup/del"
	SYSTEM_USERGROUP_BYID                        = "/system/userGroup/getUserGroupById"
	SYSTEM_USERGROUP_LIST2                       = "/system/userGroup/list2"
	SYSTEM_USERGROUP_USERGROUPTOCONTACTMENU      = "/system/userGroup/userGroupToContactMenu"
	SYSTEM_USERGROUP_USERGROUPCONFIRMCONTACTMENU = "/system/userGroup/userGroupConfirmContactMenu"
	SYSTEM_USERGROUP_USERGROUPTOCONTACTOPER      = "/system/userGroup/userGroupToContactOper"
	SYSTEM_USERGROUP_USERGROUPCONFIRMCONTACTOPER = "/system/userGroup/userGroupConfirmOper"

	//------------ html name -------------------
	USERGROUP_LIST = "system/user_group.html"
)

func userGroupConfirmContactOper(w http.ResponseWriter, r *http.Request) {
	userGroupId := r.FormValue("userGroupId")
	operIds := r.FormValue("operIds")

	var num int64
	if util.IsEmpty(operIds) {
		num = service.UserGroupConfirmContactOper(userGroupId, nil)
	} else {
		num = service.UserGroupConfirmContactOper(userGroupId, strings.Split(operIds, ","))
	}

	io.WriteString(w, strconv.FormatInt(num, 10))

}

func userGroupToContactOper(w http.ResponseWriter, r *http.Request) {
	userGroupId := r.FormValue("userGroupId")

	rows := service.UserGroupToContactOper(userGroupId)
	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))
}

func userGroupConfirmContactMenu(w http.ResponseWriter, r *http.Request) {
	userGroupId := r.FormValue("userGroupId")
	menuIds := r.FormValue("menuIds")

	var num int64 = 0
	if util.IsNotEmpty(menuIds) {
		num = service.UserGroupConfirmContactMenu(userGroupId, strings.Split(menuIds, ","))
	} else {
		num = service.UserGroupConfirmContactMenu(userGroupId, nil)
	}

	if num > 0 {
		io.WriteString(w, strconv.FormatInt(num, 10))
	} else {
		io.WriteString(w, "")
	}
}

func userGroupToContactMenu(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	//currentUserMenus := service.GetCurrentUserMenus()
	userGroupMenuIds := service.GetMenuIdsByGroupId(id)
	rows, _ := service.GetMenusInIds(userGroupMenuIds)

	//	rows := make((map[string]interface{}))
	//	rows["currentMenus"] = currentUserMenus
	//	rows["roleMenus"] = roleMenus

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))

}

//get user group id
func getUserGroupById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	rows := service.GetUserGroupById(id)
	data, _ := json.Marshal(rows)

	io.WriteString(w, string(data))
}

func userGroupAdd(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	remark := r.FormValue("remark")

	m := model.NewSysUserGroup()
	m.Name = name
	m.Remark = remark

	dispatcher = "groupAdd"
	adapterAdd(w, m)
}

func userGroupDel(w http.ResponseWriter, r *http.Request) {

	ids := r.FormValue("ids")
	isDelUser := r.FormValue("isDelUser")

	var delGroupN, delUserN int64 = 0, 0
	if isDelUser == "0" {
		delGroupN, delUserN = service.DelUserGroupInIds(strings.Split(ids, ","), false)
	} else if isDelUser == "1" {
		delGroupN, delUserN = service.DelUserGroupInIds(strings.Split(ids, ","), true)
	}

	rows := make(map[string]int64)

	rows["delGroupN"] = delGroupN
	rows["delUserN"] = delUserN

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))

}

func userGroupEdit(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	remark := r.FormValue("remark")
	id := r.FormValue("id")

	m := model.NewSysUserGroup()
	m.Id = id
	m.Name = name
	m.Remark = remark

	dispatcher = "userGroupEdit"
	adapterEdit(w, m)
}

//to userGroup list
func userGroupToList(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("menuId")
	web.DevEnvExecuteHtml(w, USERGROUP_LIST, id)
}

func groupList(w http.ResponseWriter, r *http.Request) {

	rows := service.GetUserGroupList()

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))
}

//userGroup list
func userGroupList(w http.ResponseWriter, r *http.Request) {

	dispatcher = "userGroup"
	adapterList(w, r)
}
