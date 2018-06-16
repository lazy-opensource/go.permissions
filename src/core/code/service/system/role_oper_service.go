package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
)

func AddRolesOper(roleIds []string, operId string) int64 {
	if util.CheckSlick(roleIds) || util.IsEmpty(operId) {
		return 0
	}

	num = 1
	for _, v := range roleIds {

		if num == 0 {
			break
		}
		m := model.NewSysRoleOper()
		m.RoleId = v
		m.OperId = operId
		SetUuid(m)

		mf := NewMF(m)

		ex := dao.NewExecute()
		ex.Insert(mf)

		num = mf.Num

	}

	return num
}

func DelRoleOperInOperIds(ids []string) int64 {
	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRoleOper()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Id = "OperId"
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

func DelRoleOperInRoleIds(ids []string) int64 {
	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysRoleOper()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Id = "RoleId"
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

//get operids by role id
func GetOperIdsByRoleId(id string) []string {
	if !util.IsNotEmpty(id) {
		return make([]string, 0)
	}

	m := model.NewSysRoleOper()
	m.RoleId = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	mf.Rule["RoleId"] = dao.EQ
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	ids := make([]string, len(mf.ResultSet))
	tag := 0

	for _, v := range mf.ResultSet {
		ids[tag] = v["OperId"]
		tag++
	}

	fmt.Println("operids by roleid", ids[0:])
	return ids
}

//get oper ids by role ids
func GetOperIdsInRoleIds(roleIds []string) []string {
	if util.CheckSlick(roleIds) {
		return nil
	}
	m := model.NewSysRoleOper()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsPagin = false
	mf.IsSort = false
	mf.Ids = roleIds
	mf.Id = "RoleId"
	mf.Field["OperId"] = "OperId" //自定义需要检索的字段
	mf.IsDistinct = true          //多个角色能有相同的菜单权限，需要去重

	ex := dao.NewExecute()
	msg := ex.QueryAllOrByCondition(mf)

	if msg.Code == 200 {
		rlt := mf.ResultSet
		if l := len(rlt); l > 0 {
			operIds := make([]string, len(mf.ResultSet))
			tag := 0
			for _, v := range rlt {
				operIds[tag] = v["OperId"]
				tag++
			}
			fmt.Println("operIds ----->", operIds)
			return operIds
		}
	}
	return nil
}
