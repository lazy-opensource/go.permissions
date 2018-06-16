package controller

import (
	"core/code/model/system"
	"core/code/service/system"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"lzy/framework/web"
	"net/http"
	"strconv"
)

const (
	//------------ route name  ---------------
	SYSTEM_USER_TOLIST            = "/system/user/toList"
	SYSTEM_USER_LIST              = "/system/user/list"
	SYSTEM_USER_DEL               = "/system/user/del"
	SYSTEM_USER_ADD               = "/system/user/add"
	SYSTEM_USER_EDIT              = "/system/user/edit"
	SYSTEM_USER_TOEDIT            = "/system/user/toEdit"
	SYSTEM_USER_CONTACTGROUP      = "/system/user/contactGroup"
	SYSTEM_USER_CONFIRM_USERGROUP = "/system/user/confirmUserGroup"
	SYSTEM_GET_USER_BY_ID         = "/system/user/getUserById"
	SYSTEM_USER_GETROLESBYGROUPID = "/system/user/getRoleListByGroupId"
	SYSTEM_USER_CONTACTROLE       = "/system/user/contactRole"
	SYSTEM_USER_DETAILUSER        = "/system/user/userDetail"

	//------------ html name -------------------
	USER_LIST = "system/user.html"
)

//user detail
func userDetail(w http.ResponseWriter, r *http.Request) {

	userId := r.FormValue("userId")
	groupId := r.FormValue("groupId")
	roleIds := service.GetRoleIdsByUserId(userId)
	roles := service.GetRolesInIds(roleIds)
	group := service.GetUserGroupById(groupId)

	rows := make(map[string]interface{})
	rows["group"] = group
	rows["roles"] = roles

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))
}

//contact role
func contactRole(w http.ResponseWriter, r *http.Request) {
	roleIds := r.FormValue("roleIds")
	userId := r.FormValue("userId")

	num := service.ContactRole(roleIds, userId)
	if num > 0 {
		service.ActivateUserStatusByUserId(userId)
		io.WriteString(w, "1")
	} else {
		io.WriteString(w, "false")
	}

}

//get roles by groupId
func getRoleListByGroupId(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userId")
	groupId := r.FormValue("groupId")

	groupRoleIds := service.GetRoleIdsByGroupId(groupId)
	userRoleIds := service.GetRoleIdsByUserId(userId)
	groupRoles := service.GetRolesInIds(groupRoleIds)
	userRoles := service.GetRolesInIds(userRoleIds)

	rows := make(map[string]interface{})
	rows["groupRoles"] = groupRoles
	rows["userRoles"] = userRoles

	data, _ := json.Marshal(rows)

	io.WriteString(w, string(data))
}

//get user by id
func getUserById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("userId")

	user := service.GetUserById(id)

	data, _ := json.Marshal(user)
	io.WriteString(w, string(data))
}

//关联用户组/部门
func userConfirmUserGroup(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userId")
	groupId := r.FormValue("groupId")

	m := model.NewSysUser()
	m.GroupId = groupId
	m.Id = userId
	m.Status = ACTIVATE

	num := service.UserConfirmUserGroup(m)

	if num > 0 {
		io.WriteString(w, strconv.FormatInt(num, 10))
	} else {
		io.WriteString(w, "false")
	}

}

//准备关联用户组/部门
func userToContactGroup(w http.ResponseWriter, r *http.Request) {
	rlt := service.GetUserById(r.FormValue("id"))

	id := rlt[0]["GroupId"]

	userG := service.GetUserGroupById(id)
	groups := service.GetUserGroupList()

	rows := make(map[string]interface{})
	rows["userG"] = userG
	rows["groups"] = groups

	data, _ := json.Marshal(rows)

	io.WriteString(w, string(data))

}

//to user list
func userToList(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("menuId")
	web.DevEnvExecuteHtml(w, USER_LIST, id)
}

func userAdd(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	email := r.FormValue("Email")
	password := r.FormValue("Password")
	remark := r.FormValue("Remark")

	sysUser := model.NewSysUser()
	sysUser.Name = name
	sysUser.Remark = remark
	h := md5.New()
	h.Write([]byte(password))
	sysUser.Password = hex.EncodeToString(h.Sum(nil))
	sysUser.Email = email
	sysUser.Status = ILLEGAL
	sysUser.GroupId = getCurrentUserGroupId()

	dispatcher = "userAdd"
	adapterAdd(w, sysUser)

}

//user edit
func userEdit(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	remark := r.FormValue("remark")
	id := r.FormValue("id")
	status := r.FormValue("status")

	sysUser := model.NewSysUser()
	sysUser.Name = name
	sysUser.Password = password
	sysUser.Remark = remark
	sysUser.Email = email
	sysUser.Status = status
	sysUser.Id = id

	dispatcher = "userEdit"
	adapterEdit(w, sysUser)
}

func userDel(w http.ResponseWriter, r *http.Request) {

	dispatcher = "userDel"
	adapterDel(w, r)
}

//user list
func userListByPagin(w http.ResponseWriter, r *http.Request) {

	dispatcher = "user"
	adapterList(w, r)
}
