package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
)

func UpdateUserStatusInGroupIds(ids []string) int64 {
	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysUser()
	m.Status = "2"

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Ids = ids
	mf.Id = "GroupId"
	mf.KV["Status"] = "Status"

	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num

}

func ActivateUserStatusByUserId(id string) int64 {
	if util.IsEmpty(id) {
		return 0
	}

	m := model.NewSysUser()
	m.Id = id
	m.Status = "1"

	mf := NewMF(m)
	mf.Rule["Id"] = dao.EQ
	mf.KV["Status"] = "Status"

	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num

}

func DelUserInGroupIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysUser()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Ids = ids
	mf.Id = "GroupId"

	ex := dao.NewExecute()
	ex.Delete(mf)

	return mf.Num
}

//关联用户组/部门
func UserConfirmUserGroup(m *model.SysUser) int64 {

	if m == nil {
		return 0
	}

	mf := NewMF(m)
	mf.KV["GroupId"] = "GroupId"
	mf.KV["Status"] = "Status"
	mf.Rule["Id"] = dao.EQ

	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num

}

//edit
func UserEdit(m *model.SysUser) int64 {

	if m == nil {
		return 0
	}

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Ids = make([]string, 1)
	mf.Ids[0] = m.Id
	//mf.Id = "Id"
	mf.KV["Name"] = "Name"
	mf.KV["Password"] = "Password"
	mf.KV["Remark"] = "Remark"
	mf.KV["Email"] = "Email"
	mf.KV["Status"] = "Status"
	mf.KV["UpdateTime"] = "UpdateTime"

	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num

}

//get user by id
func GetUserById(id string) map[int]map[string]string {
	if !util.IsNotEmpty(id) {
		return make(map[int]map[string]string)
	}

	sysUser := model.NewSysUser()
	sysUser.Id = id

	mf := dao.NewMultiFunction(sysUser)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	mf.IsSort = false
	mf.IsPagin = false

	mf.Rule["Id"] = dao.EQ

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}

//delete user by ids
func DelUserByIds(ids []string) int64 {
	sysUser := model.NewSysUser()

	mf := NewMF(sysUser)
	mf.Ids = ids
	mf.Id = "Id"

	ex := dao.NewExecute()
	ex.Delete(mf)

	DelRoleUserInUserIds(ids)
	return mf.Num
}

//用户列表
func GetUserListByPagin(param, param2 map[string]interface{}) (map[int]map[string]string, int) {
	m := model.NewSysUser()

	mf := handleQueryListParam(param, m)
	SetCondition(param2, m, mf)
	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	fmt.Println("user list by pagin ------>", mf.ResultSet)
	return mf.ResultSet, mf.GetTotal()
}
