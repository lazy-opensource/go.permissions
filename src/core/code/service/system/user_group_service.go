package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
	"lzy/framework/web"
)

func UserGroupConfirmContactOper(userGroupId string, operIds []string) int64 {

	if util.IsEmpty(userGroupId) {
		return 0
	}

	var num int64 = 1
	ids := make([]string, 1)
	ids[0] = userGroupId
	if util.CheckSlick(operIds) {

		DelGroupOperInGroupIds(ids)
	} else {

		DelGroupOperInGroupIds(ids)
		for _, v := range operIds {

			if num == 0 {
				break
			}
			m := model.NewSysGroupOper()
			m.GroupId = userGroupId
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

func UserGroupToContactOper(userGroupId string) map[string]interface{} {

	if util.IsEmpty(userGroupId) {
		return make(map[string]interface{})
	}

	userGroupOperIds := GetOperIdsByGroupId(userGroupId)
	userGroupOpers := CheckAndGetCurrentUserOperPermissions(userGroupOperIds)

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

	rows["userGroupOpers"] = userGroupOpers
	rows["menuUnitOper"] = menuUnitOper
	rows["menuNames"] = menuNames

	return rows
}

func UserGroupConfirmContactMenu(userGroupId string, menuIds []string) int64 {
	if !util.IsNotEmpty(userGroupId) {
		return 0
	}

	fmt.Println("userGroup contact menu menuIds", menuIds)
	fmt.Println("userGroup contact menu useGroupId", userGroupId)

	var num int64 = 1
	ids := make([]string, 1)
	ids[0] = userGroupId
	if util.CheckSlick(menuIds) {

		DelGroupMenuInGroupIds(ids)

	} else {

		DelGroupMenuInGroupIds(ids)
		for _, v := range menuIds {

			if num == 0 {
				break
			}
			m := model.NewSysGroupMenu()
			m.GroupId = userGroupId
			m.MenuId = v
			SetUuid(m)

			mf := NewMF(m)

			ex := dao.NewExecute()
			ex.Insert(mf)

			num = mf.Num
		}

	}
	return num
}

func GetUserGroupList() map[int]map[string]string {
	m := model.NewSysUserGroup()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsQueryAll = true
	mf.IsPagin = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}

func DelUserGroupInIds(ids []string, isDelUser bool) (groupN, userN int64) {

	if util.CheckSlick(ids) {
		return 0, 0
	}

	m := model.NewSysUserGroup()

	mf := dao.NewMultiFunction(m)
	mf.Ids = ids
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	ex := dao.NewExecute()
	ex.Delete(mf)

	DelGroupMenuInGroupIds(ids)
	DelGroupOperInGroupIds(ids)
	DelRoleGroupInGroupIds(ids)

	var num int64 = 0
	if isDelUser {
		num = DelUserInGroupIds(ids)
	} else {
		num = UpdateUserStatusInGroupIds(ids)
	}

	return mf.Num, num
}

//edit
func UserGroupEdit(m *model.SysUserGroup) int64 {

	if m == nil {
		return 0
	}

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	//mf.Id = "Id"
	mf.Rule["Id"] = dao.EQ
	mf.KV["Name"] = "Name"
	mf.KV["Remark"] = "Remark"
	mf.KV["UpdateTime"] = "UpdateTime"

	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num

}

//get user group list by pagin
func GetUserGroupListByPagin(param, param2 map[string]interface{}) (map[int]map[string]string, int) {
	m := model.NewSysUserGroup()

	mf := handleQueryListParam(param, m)
	SetCondition(param2, m, mf)
	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	fmt.Println("user group list by pagin ------>", mf.ResultSet)
	return mf.ResultSet, mf.GetTotal()
}

func GetGroupsByIds(ids []string) map[int]map[string]string {
	if util.CheckSlick(ids) {
		return make(map[int]map[string]string)
	}

	m := model.NewSysUserGroup()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsPagin = false
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}

//get userGroup by id
func GetUserGroupById(id string) map[int]map[string]string {
	if !util.IsNotEmpty(id) {
		return make(map[int]map[string]string)
	}

	m := model.NewSysUserGroup()
	m.Id = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = web.GetSession().Get("DbAdrr").(string)
	mf.DbType = web.GetSession().Get("DbType").(string)
	mf.IsSort = false
	mf.IsPagin = false
	mf.Rule["Id"] = dao.EQ

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}
