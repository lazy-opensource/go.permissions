package service

import (
	"core/code/model/system"
	"lzy/framework/dao"
	"lzy/framework/util"
)

func AddUserGroupOper(userGroupId, operId string) int64 {
	if util.IsEmpty(userGroupId) || util.IsEmpty(operId) {
		return 0
	}

	m := model.NewSysGroupOper()
	m.GroupId = userGroupId
	m.OperId = operId

	SetUuid(m)

	mf := NewMF(m)

	ex := dao.NewExecute()
	ex.Insert(mf)

	return mf.Num
}

func DelGroupOperInOperIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysGroupOper()

	mf := NewMF(m)
	mf.Ids = ids
	mf.Id = "OperId"

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

func DelGroupOperInGroupIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysGroupOper()

	mf := NewMF(m)
	mf.Ids = ids
	mf.Id = "GroupId"

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

func GetOperIdsByGroupId(id string) []string {

	if !util.IsNotEmpty(id) {
		return make([]string, 0)
	}

	m := model.NewSysGroupOper()
	m.GroupId = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	mf.Rule["GroupId"] = dao.EQ
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	rows := mf.ResultSet
	operIds := make([]string, len(rows))
	tag := 0

	for _, v := range rows {
		operIds[tag] = v["OperId"]
		tag++
	}

	return operIds

}

//get oper ids in group ids
func GetOperIdsInGroupIds(ids []string) []string {
	if util.CheckSlick(ids) {
		return make([]string, 0)
	}

	m := model.NewSysGroupOper()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	mf.Ids = ids
	mf.Id = "GroupId"
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	rows := mf.ResultSet
	operIds := make([]string, len(rows))
	tag := 0

	for _, v := range rows {
		operIds[tag] = v["OperId"]
		tag++
	}

	return operIds
}
