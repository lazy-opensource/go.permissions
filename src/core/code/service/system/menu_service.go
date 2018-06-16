package service

import (
	"core/code/model/system"
	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
	"strings"
)

func MenuConfirmContactParentMenu(pids, mainIds string) int64 {

	if util.IsEmpty(mainIds) {
		return 0
	}

	num = 1
	mainIdsSli := strings.Split(mainIds, ",")

	if util.IsEmpty(pids) {
		for _, v := range mainIdsSli {
			if num < 1 {
				return 0
			}
			m := model.NewSysMenu()
			m.Id = v
			m.ParentId = "0"

			mf := NewMF(m)
			mf.Rule["Id"] = dao.EQ
			mf.KV["ParentId"] = "ParentId"

			ex := dao.NewExecute()
			ex.Update(mf)
			num = mf.Num
		}
	} else {
		pidsSli := strings.Split(pids, ",")

		for _, v := range pidsSli {
			for _, v2 := range mainIdsSli {
				if num < 1 {
					return 0
				}
				m := model.NewSysMenu()
				m.Id = v2
				m.ParentId = v

				mf := NewMF(m)
				mf.Rule["Id"] = dao.EQ
				mf.KV["ParentId"] = "ParentId"

				ex := dao.NewExecute()
				ex.Update(mf)
				num = mf.Num
			}
		}
	}

	return num

}

func MenuContactForCurrentUserPermissions(menuId string) int64 {
	if util.IsEmpty(menuId) {
		return 0
	}

	num = 1
	userId := GetCurrentUserId()
	roleIds := GetRoleIdsByUserId(userId)
	num = AddRolesMenu(roleIds, menuId)

	userGroupId := GetUserById(userId)[0]["GroupId"]
	num = AddUserGroupMenu(userGroupId, menuId)

	return num

}

func DelMenuInIds(ids []string) int64 {

	if util.CheckSlick(ids) {
		return 0
	}

	m := model.NewSysMenu()

	mf := NewMF(m)
	mf.Ids = ids

	ex := dao.NewExecute()
	ex.Delete(mf)

	DelRoleMenuInMenuIds(ids)
	DelMenuInParentIds(ids)
	DelOperInMenuIds(ids)
	DelGroupMenuInMenuIds(ids)

	return mf.Num
}

func DelMenuInParentIds(pids []string) int64 {
	if util.CheckSlick(pids) {
		return 0
	}

	m := model.NewSysMenu()

	mf := NewMF(m)
	mf.Ids = pids
	mf.Id = "ParentId"

	ex := dao.NewExecute()
	ex.Delete(mf)
	return mf.Num
}

func MenuEdit(m *model.SysMenu) int64 {

	if m == nil {
		return 0
	}

	mf := NewMF(m)
	mf.Rule["Id"] = dao.EQ
	mf.KV["Name"] = "Name"
	mf.KV["Remark"] = "Remark"
	mf.KV["UpdateTime"] = "UpdateTime"
	mf.KV["Icon"] = "Icon"

	fmt.Println("begin edit menu ...")
	fmt.Println(m)
	ex := dao.NewExecute()
	ex.Update(mf)

	return mf.Num
}

func GetMenuList() map[int]map[string]string {

	m := model.NewSysMenu()

	mf := NewMF(m)
	mf.IsPagin = false
	mf.IsSort = false
	mf.IsQueryAll = true

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}

//获得当前用户的所有的菜单权限
func GetCurrentUserMenus() map[int]map[string]string {

	menuIds := GetUserFinalMenuIds()
	util.BubbleSortSlick(menuIds)
	result, _ := GetMenusInIds(menuIds)
	fmt.Println("current user menus +++++++++", result)

	return result
}

func GetUserFinalMenuIds() []string {

	userId := GetCurrentUserId()

	roleIds := GetRoleIdsByUserId(userId)
	roleMenuIds := GetMenuIdsInRoleIds(roleIds)

	groupId := GetUserById(userId)[0]["GroupId"]
	userGroupMenuIds := GetMenuIdsByGroupId(groupId)
	if util.CheckSlick(roleMenuIds) || util.CheckSlick(userGroupMenuIds) {
		return make([]string, 0)
	}

	finalMenuIds := make([]string, len(roleMenuIds)+len(userGroupMenuIds))
	tag := 0

	for _, v := range userGroupMenuIds {
		for _, v2 := range roleMenuIds {
			if v == v2 { //某个角色菜单属于该组
				finalMenuIds[tag] = v
				tag++
				break
				//roleMenuIds = append(roleMenuIds[:k], roleMenuIds[k+1:]...)
				//groupMenuIds = append(groupMenuIds[:k2], groupMenuIds[k2+1:]...)
			}
		}
	}

	finalMenuIds = finalMenuIds[0:tag]

	fmt.Println("final menu ids ", finalMenuIds[0:])
	return finalMenuIds
}

//检查当前用户是否拥有该菜单的权限
func CheckMenuPermissions(menuIds []string) map[int]map[string]string {

	//当前用户拥有的菜单Id
	menuMap := GetCurrentUserMenus()

	if menuMap != nil && len(menuMap) > 0 {
		var fv string
		var tag int
		result := make(map[int]map[string]string)
		for _, v := range menuIds { //遍历需要检查的菜单Id
			for k2, v2 := range menuMap { //遍历当前用户拥有的所有菜单结果集
				fv = v2["Id"]
				if fv == v {
					result[tag] = v2
					tag++
					delete(menuMap, k2)
					break
				}
			}
		}
		return result
	}
	return nil
}

//菜单列表
func GetMenuListByPagin(param, param2 map[string]interface{}) (map[int]map[string]string, int) {
	m := model.NewSysMenu()

	mf := handleQueryListParam(param, m)
	SetCondition(param2, m, mf)
	currentMenuIds := GetUserFinalMenuIds()
	mf.Ids = currentMenuIds
	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	fmt.Println("menu list by pagin ------>", mf.ResultSet)

	menuIds := make([]string, len(mf.ResultSet))
	tag := 0

	for _, v := range mf.ResultSet {
		menuIds[tag] = v["Id"]
		tag++
	}

	return CheckMenuPermissions(menuIds), 0
}

//get child menus by parentId
func GetChildMenuByParentId(id string) map[int]map[string]string {

	if !util.IsNotEmpty(id) {
		return nil
	}
	m := model.NewSysMenu()
	m.ParentId = id

	mf := NewMF(m)
	mf.Rule["ParentId"] = dao.EQ
	mf.IsPagin = false

	ex := dao.NewExecute()
	fmt.Println("++++++++++get child menus by parentId +++++++++++++++++")
	ex.QueryAllOrByCondition(mf)

	childIds := make([]string, len(mf.ResultSet))
	result := mf.ResultSet
	tag := 0

	for _, v := range result {
		id := v["Id"]
		childIds[tag] = id
		tag++
	}
	fmt.Printf("child menu is ------->", childIds)
	util.BubbleSortSlick(childIds)
	result = CheckMenuPermissions(childIds)
	return result

	return nil
}

func GetMenuById(id string) map[int]map[string]string {

	if !util.IsNotEmpty(id) {
		return make(map[int]map[string]string)
	}

	m := model.NewSysMenu()
	m.Id = id

	mf := NewMF(m)
	mf.Rule["Id"] = dao.EQ
	mf.IsSort = false
	mf.IsPagin = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}

//get menus by ids
func GetMenusInIds(menuIds []string) (map[int]map[string]string, *util.Message) {

	if util.CheckSlick(menuIds) {
		return nil, nil
	}
	m := model.NewSysMenu()

	mf := NewMF(m)
	mf.Ids = menuIds
	mf.IsPagin = false

	ex := dao.NewExecute()

	msg := ex.QueryAllOrByCondition(mf)

	result := mf.ResultSet
	fmt.Println("menus -------->", result)
	return result, msg
}

//get one menus by ids
func GetOneMenusInIds(menuIds []string) (map[int]map[string]string, *util.Message) {

	if util.CheckSlick(menuIds) {
		return nil, nil
	}

	m := model.NewSysMenu()
	m.ParentId = "0"

	mf := NewMF(m)
	mf.Ids = menuIds
	mf.Rule["ParentId"] = dao.EQ
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	fmt.Println("get one menus by ids--------------")
	msg := ex.QueryAllOrByCondition(mf)

	result := mf.ResultSet
	return result, msg
}
