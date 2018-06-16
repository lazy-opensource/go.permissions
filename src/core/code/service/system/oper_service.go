package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
	"lzy/framework/web"
)

func OperContactForCurrentUserPermissions(operId string) int64 {

	if util.IsEmpty(operId) {
		return 0
	}
	num = 1
	userId := GetCurrentUserId()
	roleIds := GetRoleIdsByUserId(userId)
	num = AddRolesOper(roleIds, operId)

	userGroupId := GetUserById(userId)[0]["GroupId"]
	num = AddUserGroupOper(userGroupId, operId)

	return num
}

func OperContactMenu(operId, menuId string) int64 {
	if !util.IsNotEmpty(operId) || !util.IsNotEmpty(menuId) {
		return 0
	}

	m := model.NewSysOper()
	m.Id = operId
	m.MenuId = menuId

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Rule["Id"] = dao.EQ
	mf.KV["MenuId"] = "MenuId"

	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num
}

func DelOperInMenuIds(ids []string) int64 {
	if util.CheckSlick(ids) {
		return 0
	}
	num = 1
	for _, v := range ids {

		if num < 1 {
			break
		}
		m := model.NewSysOper()
		m.MenuId = v

		mf := NewMF(m)
		mf.Rule["MenuId"] = dao.EQ

		ex := dao.NewExecute()
		ex.Delete(mf)
		num = mf.Num
	}

	return num

}

func GetOperIdsInMenuIds(ids []string) []string {

	if util.CheckSlick(ids) {
		return make([]string, 0)
	}

	m := model.NewSysOper()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Id = "MenuId"
	mf.Ids = ids
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	operIds := make([]string, len(mf.ResultSet))
	tag := 0

	for _, v := range mf.ResultSet {
		operIds[tag] = v["Id"]
		tag++
	}

	return operIds
}

func DelOperInIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysOper()

	mf := NewMF(m)
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	DelRoleOperInOperIds(ids)
	DelGroupOperInOperIds(ids)

	return mf.Num
}

func OperEdit(m *model.SysOper) int64 {

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
	mf.KV["Url"] = "Url"

	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num
}

func GetOperById(id string) map[int]map[string]string {

	if !util.IsNotEmpty(id) {
		return make(map[int]map[string]string)
	}

	m := model.NewSysOper()
	m.Id = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Rule["Id"] = dao.EQ
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}

//get current user opers
func GetCurrentUserOper() map[int]map[string]string {

	operIds := GetUserFinalOperIds()
	util.BubbleSortSlick(operIds)

	result := GetOpersByIds(operIds)

	return result
}

func GetUserFinalOperIds() []string {

	userId := web.GetSession().Get("userId").(string)
	roleId := GetRoleIdsByUserId(userId)
	roleOperIds := GetOperIdsInRoleIds(roleId)

	userGroupId := GetUserById(userId)[0]["GroupId"]
	groupOperIds := GetOperIdsByGroupId(userGroupId)

	if util.CheckSlick(roleOperIds) || util.CheckSlick(groupOperIds) {
		return make([]string, 0)
	}

	finalOperIds := make([]string, len(roleOperIds)+len(groupOperIds))
	tag := 0

	for _, v := range roleOperIds {
		for _, v2 := range groupOperIds {
			if v == v2 {
				finalOperIds[tag] = v
				tag++
				break
				//roleOperIds = append(roleOperIds[:k], roleOperIds[k+1:]...)
				//groupOperIds = append(groupOperIds[:k2], groupOperIds[k2+1:]...)
			}
		}
	}

	finalOperIds = finalOperIds[:tag+1]
	fmt.Println("current user operIds ", finalOperIds)
	return finalOperIds
}

func CheckAndGetCurrentUserOperPermissions(operIds []string) map[int]map[string]string {

	if util.CheckSlick(operIds) {
		return make(map[int]map[string]string)
	}

	currentUserOperIds := GetUserFinalOperIds()
	finalOperIds := make([]string, len(currentUserOperIds))
	tag := 0
	for _, v := range operIds {
		for _, v2 := range currentUserOperIds {
			if v == v2 {
				finalOperIds[tag] = v
				tag++
				break
			}
		}
	}

	finalOperIds = finalOperIds[:tag]

	rows := GetOpersByIds(finalOperIds)
	fmt.Printf("check  operids %s , check after operIds %s", operIds, finalOperIds)
	return rows
}

//check oper permissions
func CheckMenuOperPermissions(menuOper, userOper map[int]map[string]string) map[int]map[string]string {

	rlt := make(map[int]map[string]string)
	if menuOper == nil || userOper == nil {
		return rlt
	}

	fmt.Println("menu opers", menuOper)
	fmt.Println("user opers ", userOper)

	tag := 0
	for k, v := range menuOper {
		for k2, v2 := range userOper {
			if v["Id"] == v2["Id"] {
				rlt[tag] = v
				tag++
				delete(menuOper, k)
				delete(userOper, k2)
			}
		}
	}

	fmt.Println("final user oper ", rlt)
	return rlt
}

func GetOpersByMenuId(id string) map[int]map[string]string {

	if !util.IsNotEmpty(id) {
		return make(map[int]map[string]string)
	}

	m := model.NewSysOper()
	m.MenuId = id

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.Rule["MenuId"] = dao.EQ
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)
	return mf.ResultSet

}

//get oper by ids
func GetOpersByIds(operIds []string) map[int]map[string]string {

	if util.CheckSlick(operIds) {
		return make(map[int]map[string]string)
	}
	m := model.NewSysOper()

	mf := dao.NewMultiFunction(m)
	mf.Ids = operIds

	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	result := mf.ResultSet
	fmt.Println("opers -------->", result)
	return result
}

func GetOperList() map[int]map[string]string {
	m := model.NewSysOper()

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.IsQueryAll = true
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}

//操作列表
func GetOperListByPagin(param, param2 map[string]interface{}) (map[int]map[string]string, int) {
	m := model.NewSysOper()

	mf := handleQueryListParam(param, m)
	SetCondition(param2, m, mf)
	currentUserOperIds := GetUserFinalOperIds()
	mf.Ids = currentUserOperIds
	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet, mf.GetTotal()
}
