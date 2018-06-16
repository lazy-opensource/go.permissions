package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
)

func AddRoleForCurrentUserGroupByRoleId(roleId string) int64 {
	if util.IsEmpty(roleId) {
		return 0
	}

	currentGrouId := GetCurrentUserGroupId()
	num := AddRoleGroup(roleId, currentGrouId)

	return num

}

func AddRoleGroup(roleId, groupId string) int64 {
	if util.IsEmpty(roleId) || util.IsEmpty(groupId) {
		return 0
	}

	m := model.NewSysRoleGroup()
	m.GroupId = groupId
	m.RoleId = roleId
	SetUuid(m)

	mf := NewMF(m)

	ex := dao.NewExecute()
	ex.Insert(mf)
	return mf.Num
}

func DelRoleGroupInRoleIds(ids []string) int64 {
	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRoleGroup()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Id = "RoleId"
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

func DelRoleGroupInGroupIds(ids []string) int64 {
	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRoleGroup()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Id = "GroupId"
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

//get group ids by role id
func GetGroupIdsByRoleId(id string) []string {

	if !util.IsNotEmpty(id) {
		return make([]string, 0)
	}

	m := model.NewSysRoleGroup()
	m.RoleId = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Rule["RoleId"] = dao.EQ
	mf.IsPagin = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	groupIds := make([]string, len(mf.ResultSet))
	tag := 0

	for _, v := range mf.ResultSet {
		groupIds[tag] = v["GroupId"]
		tag++
	}

	fmt.Println("group ids by role id", groupIds[0:])

	return groupIds

}

//get role ids by group id
func GetRoleIdsByGroupId(id string) []string {

	if !util.IsNotEmpty(id) {
		return make([]string, 0)
	}

	m := model.NewSysRoleGroup()
	m.GroupId = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Rule["GroupId"] = dao.EQ
	mf.IsPagin = false
	mf.Field["RoleId"] = "RoleId"

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	roleIds := make([]string, len(mf.ResultSet))
	tag := 0

	for _, v := range mf.ResultSet {
		roleIds[tag] = v["RoleId"]
		tag++
	}

	fmt.Println("role ids by group id", roleIds[0:])

	return roleIds

}
