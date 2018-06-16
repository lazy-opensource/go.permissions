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
	SYSTEM_ROLE_TOLIST                 = "/system/role/toList"
	SYSTEM_ROLE_LIST                   = "/system/role/list"
	SYSTEM_ROLE_ADD                    = "/system/role/add"
	SYSTEM_ROLE_EDIT                   = "/system/role/edit"
	SYSTEM_ROLE_DEL                    = "/system/role/del"
	SYSTEM_ROLE_GETROLEBYID            = "/system/role/getRoleById"
	SYSTEM_ROLE_TOCONTACTOPER          = "/system/role/roleToContactOper"
	SYSTEM_ROLE_TOCONTACTGROUP         = "/system/role/roleToContactGroup"
	SYSTEM_ROLE_CONFIRMCONTACTGROUP    = "/system/role/confirmContactGroup"
	SYSTEM_ROLE_CONFIRMCONTACTMENU     = "/system/role/confirmContactMenu"
	SYSTEM_ROLE_TOCONTACTMENU          = "/system/role/roleToContactMenu"
	SYSTEM_ROLE_ROLECONFIRMCONTACTOPER = "/system/role/roleConfirmContactOper"

	//------------ html ------------------
	ROLE_LIST = "/system/role.html"
)

func roltToContactMenu(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	//currentUserMenus := service.GetCurrentUserMenus()
	roleMenuIds := service.GetMenuIdsByRoleId(id)
	rows, _ := service.GetMenusInIds(roleMenuIds)

	//	rows := make((map[string]interface{}))
	//	rows["currentMenus"] = currentUserMenus
	//	rows["roleMenus"] = roleMenus

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))

}

func roleConfirmContactMenu(w http.ResponseWriter, r *http.Request) {
	roleId := r.FormValue("roleId")
	menuIds := r.FormValue("menuIds")

	var num int64 = 0
	if util.IsNotEmpty(menuIds) {
		num = service.RoleConfirmContactMenu(roleId, strings.Split(menuIds, ","))
	} else {
		num = service.RoleConfirmContactMenu(roleId, nil)
	}

	if num > 0 {
		io.WriteString(w, strconv.FormatInt(num, 10))
	} else {
		io.WriteString(w, "")
	}
}

func roleConfirmContactGroup(w http.ResponseWriter, r *http.Request) {

	roleId := r.FormValue("roleId")
	groupIds := r.FormValue("groupIds")

	var num int64 = 0
	if util.IsNotEmpty(groupIds) {
		num = service.RoleConfirmContactGroup(roleId, strings.Split(groupIds, ","))
	} else {
		num = service.RoleConfirmContactGroup(roleId, nil)
	}

	if num > 0 {
		io.WriteString(w, strconv.FormatInt(num, 10))
	} else {
		io.WriteString(w, "")
	}
}

func roleToContactGroup(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	groups := service.GetUserGroupList()
	roleGroupIds := service.GetGroupIdsByRoleId(id)
	roleGroups := service.GetGroupsByIds(roleGroupIds)

	rows := make(map[string]interface{})

	rows["groups"] = groups
	rows["roleGroups"] = roleGroups
	rows["id"] = id

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))
}

//role add
func roleAdd(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	remark := r.FormValue("remark")

	m := model.NewSysRole()
	m.Name = name
	m.Remark = remark

	dispatcher = "roleAdd"
	adapterAdd(w, m)

}

//get rolr by id
func getRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := r.FormValue("roleId")

	rows := service.GetRoleById(roleId)

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))
}

//edit role
func roleEdit(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	remark := r.FormValue("remark")
	id := r.FormValue("roleId")

	m := model.NewSysRole()
	m.Name = name
	m.Remark = remark
	m.Id = id

	dispatcher = "roleEdit"
	adapterEdit(w, m)
}

func roleToContactOper(w http.ResponseWriter, r *http.Request) {
	userGroupId := r.FormValue("roleId")

	rows := service.RoleToContactOper(userGroupId)
	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))
}

func roleConfirmContactOper(w http.ResponseWriter, r *http.Request) {
	roleId := r.FormValue("roleId")
	operIds := r.FormValue("operIds")

	var num int64
	if util.IsEmpty(operIds) {
		num = service.RoleConfirmContactOper(roleId, nil)
	} else {
		num = service.RoleConfirmContactOper(roleId, strings.Split(operIds, ","))
	}

	io.WriteString(w, strconv.FormatInt(num, 10))

}

//delete role
func roleDel(w http.ResponseWriter, r *http.Request) {

	dispatcher = "roleDel"
	adapterDel(w, r)

}

//to role list
func roleToList(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("menuId")
	web.DevEnvExecuteHtml(w, ROLE_LIST, id)
}

//role list
func roleListByPagin(w http.ResponseWriter, r *http.Request) {

	dispatcher = "role"
	adapterList(w, r)
}
