package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
)

func AddUserGroupMenu(userGroupId, menuId string) int64 {
	if util.IsEmpty(userGroupId) || util.IsEmpty(menuId) {
		return 0
	}

	m := model.NewSysGroupMenu()
	m.GroupId = userGroupId
	m.MenuId = menuId

	SetUuid(m)

	mf := NewMF(m)

	ex := dao.NewExecute()
	ex.Insert(mf)

	return mf.Num
}

func GetMenuIdsByGroupId(id string) []string {

	if !util.IsNotEmpty(id) {
		return make([]string, 0)
	}

	m := model.NewSysGroupMenu()
	m.GroupId = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Rule["GroupId"] = dao.EQ
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	l := len(mf.ResultSet)
	menuIds := make([]string, l)
	tag := 0

	for _, v := range mf.ResultSet {
		menuIds[tag] = v["MenuId"]
		tag++
	}

	return menuIds
}

func DelGroupMenuInMenuIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysGroupMenu()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Ids = ids
	mf.Id = "MenuId"

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

func DelGroupMenuInGroupIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysGroupMenu()

	mf := NewMF(m)
	mf.Ids = ids
	mf.Id = "GroupId"

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

//get menuids in grouids
func GetMenuIdsInGroupIds(ids []string) []string {

	if util.CheckSlick(ids) {
		return make([]string, 0)
	}

	m := model.NewSysGroupMenu()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Id = "GroupId"
	mf.Ids = ids
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	l := len(mf.ResultSet)
	menuIds := make([]string, l)
	tag := 0

	for _, v := range mf.ResultSet {
		id := v["MenuId"]
		menuIds[tag] = id
		tag++
	}

	fmt.Println("menuids in groupids", menuIds[0:])

	return menuIds
}
