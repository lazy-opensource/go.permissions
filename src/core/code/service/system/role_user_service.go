package service

import (
	"core/code/model/system"
	"lzy/framework/dao"
	"lzy/framework/util"
	"lzy/framework/web"
	"strings"
)

func DelRoleUserInRoleIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRoleUser()

	mf := dao.NewMultiFunction(m)
	mf.Ids = ids
	mf.Id = "RoleId"
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	ex2 := dao.NewExecute()
	ex2.Delete(mf)

	return mf.Num
}

func DelRoleUserInUserIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRoleUser()

	mf := dao.NewMultiFunction(m)
	mf.Ids = ids
	mf.Id = "UserId"
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	ex2 := dao.NewExecute()
	ex2.Delete(mf)

	return mf.Num
}

//del role by userId
func DelRoleUserbyUserId(userId string) int64 {
	if !util.IsNotEmpty(userId) {
		return 0
	}

	m := model.NewSysRoleUser()
	m.UserId = userId

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Rule["UserId"] = dao.EQ

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

//contact role
func ContactRole(roleIds string, userId string) int64 {

	DelRoleUserbyUserId(userId)

	var num int64 = 1

	if !util.IsNotEmpty(roleIds) || !util.IsNotEmpty(userId) {
		return num
	}

	roleIdsSli := strings.Split(roleIds, ",")
	for _, v := range roleIdsSli {
		m := model.NewSysRoleUser()
		SetCreateTimeAndUuid(m)
		m.RoleId = v
		m.UserId = userId

		mf := dao.NewMultiFunction(m)
		mf.DbAddr = getDbAddr()
		mf.DbType = getDbType()

		ex := dao.NewExecute()
		ex.Insert(mf)

		if mf.Num < 1 {
			num = 0
			break
		}
	}

	return num

}

//delete role_user by user ids
func DelUserFromRoleUserByIds(userIds []string) {
	m := model.NewSysRoleUser()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = web.GetSession().Get("DbAdrr").(string)
	mf.DbType = web.GetSession().Get("DbType").(string)
	mf.Ids = userIds
	mf.Id = "UserId"

	ex := dao.NewExecute()
	ex.Delete(mf)
}
