package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
)

func AddRolesMenu(roleIds []string, menuId string) int64 {
	if util.CheckSlick(roleIds) || util.IsEmpty(menuId) {
		return 0
	}

	num = 1
	for _, v := range roleIds {

		if num == 0 {
			break
		}
		m := model.NewSysRoleMenu()
		m.RoleId = v
		m.MenuId = menuId
		SetUuid(m)

		mf := NewMF(m)

		ex := dao.NewExecute()
		ex.Insert(mf)

		num = mf.Num

	}

	return num
}

func DelRoleMenuInRoleIds(ids []string) int64 {
	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRoleMenu()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Id = "RoleId"
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

func DelRoleMenuInMenuIds(ids []string) int64 {
	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRoleMenu()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Id = "MenuId"
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

//get menuIds by roleIds
func GetMenuIdsByRoleId(roleId string) []string {

	if !util.IsNotEmpty(roleId) {
		return make([]string, 0)
	}
	m := model.NewSysRoleMenu()
	m.RoleId = roleId

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsPagin = false
	mf.Rule["RoleId"] = dao.EQ

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	menuIds := make([]string, len(mf.ResultSet))
	tag := 0

	for _, v := range mf.ResultSet {
		menuIds[tag] = v["MenuId"]
		tag++
	}

	fmt.Println("menudis by role id", menuIds[0:])
	return menuIds
}

//get menuIds by roleIds
func GetMenuIdsInRoleIds(roleIds []string) []string {

	if util.CheckSlick(roleIds) {
		return make([]string, 0)
	}
	m := model.NewSysRoleMenu()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsPagin = false
	mf.Ids = roleIds
	mf.Id = "RoleId"
	mf.Field["MenuId"] = "MenuId" //自定义需要检索的字段
	mf.IsDistinct = true          //多个角色能有相同的菜单权限，需要去重

	ex := dao.NewExecute()
	fmt.Println("get menuIds by roleIds--------------")
	ex.QueryAllOrByCondition(mf)

	rlt := mf.ResultSet
	if l := len(rlt); l > 0 {
		menuIds := make([]string, len(mf.ResultSet))
		tag := 0
		for _, v := range rlt {
			menuIds[tag] = v["MenuId"]
			tag++
		}
		fmt.Println("menuids ----->", menuIds)
		return menuIds
	}

	return make([]string, 0)
}
