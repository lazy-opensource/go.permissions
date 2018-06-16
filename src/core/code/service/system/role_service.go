package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
)

func RoleToContactOper(roleId string) map[string]interface{} {

	if util.IsEmpty(roleId) {
		return make(map[string]interface{})
	}

	roleOperIds := GetOperIdsByRoleId(roleId)
	roleOpers := CheckAndGetCurrentUserOperPermissions(roleOperIds)

	currentUserMenus := GetCurrentUserMenus()
	currentUserOperIds := GetUserFinalOperIds()
	currentUserOpers := make(map[string]string)
	for _, v := range currentUserOperIds {
		currentUserOpers[v] = v
	}

	menuUnitOper := make(map[string]interface{})
	menuNames := make(map[string]string)
	var id string = ""
	for _, v := range currentUserMenus {

		id = v["Id"]
		menuOpers := GetOpersByMenuId(id)

		if len(menuOpers) > 0 {
			for k2, v2 := range menuOpers {
				if !util.IsExistsMap(currentUserOpers, v2["Id"]) {
					delete(menuOpers, k2)
				}
			}

			if len(menuOpers) > 0 {
				menuUnitOper[id] = menuOpers
				menuNames[id] = v["Name"]
			}
		}

	}

	rows := make(map[string]interface{})

	rows["roleOpers"] = roleOpers
	rows["menuUnitOper"] = menuUnitOper
	rows["menuNames"] = menuNames

	return rows
}

func RoleConfirmContactOper(roleId string, operIds []string) int64 {

	if util.IsEmpty(roleId) {
		return 0
	}

	var num int64 = 1
	ids := make([]string, 1)
	ids[0] = roleId
	if util.CheckSlick(operIds) {

		DelRoleOperInRoleIds(ids)
	} else {

		DelRoleOperInRoleIds(ids)
		for _, v := range operIds {

			if num == 0 {
				break
			}
			m := model.NewSysRoleOper()
			m.RoleId = roleId
			m.OperId = v
			SetUuid(m)

			mf := NewMF(m)

			ex := dao.NewExecute()
			ex.Insert(mf)

			num = mf.Num
		}
	}

	return num
}

func RoleConfirmContactMenu(roleId string, menuIds []string) int64 {
	if !util.IsNotEmpty(roleId) {
		return 0
	}

	var num int64 = 1
	ids := make([]string, 1)
	ids[0] = roleId
	if util.CheckSlick(menuIds) {

		DelRoleMenuInRoleIds(ids)

	} else {

		DelRoleMenuInRoleIds(ids)
		for _, v := range menuIds {

			if num == 0 {
				break
			}
			m := model.NewSysRoleMenu()
			m.RoleId = roleId
			m.MenuId = v
			SetUuid(m)

			mf := dao.NewMultiFunction(m)
			mf.DbAddr = getDbAddr()
			mf.DbType = getDbType()

			ex := dao.NewExecute()
			ex.Insert(mf)

			num = mf.Num
		}

	}
	return num
}

func RoleConfirmContactGroup(roleId string, groupIds []string) int64 {
	if !util.IsNotEmpty(roleId) {
		return 0
	}

	var num int64 = 1
	ids := make([]string, 1)
	ids[0] = roleId
	if util.CheckSlick(groupIds) {

		DelRoleGroupInRoleIds(ids)

	} else {

		DelRoleGroupInRoleIds(ids)
		for _, v := range groupIds {

			if num == 0 {
				break
			}
			m := model.NewSysRoleGroup()
			m.RoleId = roleId
			m.GroupId = v
			SetUuid(m)

			mf := dao.NewMultiFunction(m)
			mf.DbAddr = getDbAddr()
			mf.DbType = getDbType()

			ex := dao.NewExecute()
			ex.Insert(mf)

			num = mf.Num
		}

	}
	return num
}

//delete role by ids
func DelRoleInIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRole()
	mf := NewMF(m)
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	DelRoleOperInRoleIds(ids)
	DelRoleUserInRoleIds(ids)
	DelRoleMenuInRoleIds(ids)
	DelRoleGroupInRoleIds(ids)
	return mf.Num
}

//role edit
func RoleEdit(m *model.SysRole) int64 {

	if m == nil {
		return 0
	}

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Rule["Id"] = dao.EQ
	mf.KV["Name"] = "Name"
	mf.KV["Remark"] = "Remark"
	mf.KV["UpdateTime"] = "UpdateTime"

	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num

}

//get role by lid
func GetRoleById(id string) map[int]map[string]string {

	if !util.IsNotEmpty(id) {
		return make(map[int]map[string]string)
	}

	m := model.NewSysRole()
	m.Id = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	mf.Rule["Id"] = dao.EQ

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet

}

//get role by ids
func GetRolesInIds(ids []string) map[int]map[string]string {

	if util.CheckSlick(ids) {
		return make(map[int]map[string]string)
	}

	m := model.NewSysRole()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsPagin = false
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet

}

//角色列表
func GetRoleListByPagin(param, param2 map[string]interface{}) (map[int]map[string]string, int) {
	m := model.NewSysRole()

	mf := handleQueryListParam(param, m)
	SetCondition(param2, m, mf)
	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	fmt.Println("role list by pagin ------>", mf.ResultSet)
	return mf.ResultSet, mf.GetTotal()
}

//get roleIds by userId
func GetRoleIdsByUserId(userId string) []string {
	if !util.IsNotEmpty(userId) {
		return make([]string, 0)
	}
	m := model.NewSysRoleUser()
	m.UserId = userId

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsSort = false
	mf.IsPagin = false
	mf.Rule["UserId"] = dao.EQ    //检索条件及规则
	mf.Field["RoleId"] = "RoleId" //自定义检索字段

	ex := dao.NewExecute()
	msg := ex.QueryAllOrByCondition(mf)
	if msg.Code == 200 {
		if len(mf.ResultSet) > 0 {
			roleIds := make([]string, len(mf.ResultSet))
			rlt := mf.ResultSet
			tag := 0
			for _, v := range rlt {
				roleIds[tag] = v["RoleId"]
				tag++
			}
			fmt.Println("roleIds ----------->", roleIds)
			return roleIds
		}
	}
	return make([]string, 0)
}
