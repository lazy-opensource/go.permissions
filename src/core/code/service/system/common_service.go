package service

import (
	"core/code/model/system"
	"crypto/md5"
	"encoding/hex"

	"reflect"
	"strconv"
	"strings"
	"time"

	"fmt"
	"lzy/framework/dao"
	"lzy/framework/util"
	"lzy/framework/uuid"
	"lzy/framework/web"
)

var dbAddr, dbType, userId string
var num int64

func NewMF(m interface{}) *dao.MultiFunction {
	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()

	return mf
}

//get db addr
func getDbAddr() string {
	if util.IsNotEmpty(dbAddr) {
		return dbAddr
	}
	dbAddr = web.GetSession().Get("DbAdrr").(string)
	return dbAddr
}

func SetCreateTimeAndUuid(m interface{}) {

	o := reflect.ValueOf(m).Elem()
	o.FieldByName("CreateTime").SetString(time.Now().Format("2006-01-02 15:04:05"))
	o.FieldByName("Id").SetString(uuid.Rand().Hex())

}

func Add(action string, m interface{}) int64 {

	if m == nil {
		return 0
	}

	var mf *dao.MultiFunction
	switch action {
	case "user":
		mf = dao.NewMultiFunction(m.(*model.SysUser))
	case "role":
		m2 := m.(*model.SysRole)
		AddRoleForCurrentUserGroupByRoleId(m2.Id)
		mf = dao.NewMultiFunction(m.(*model.SysRole))
	case "group":
		m2 := m.(*model.SysUserGroup)
		AddUserGroupMenu(m2.Id, "0")
		mf = dao.NewMultiFunction(m2)
	case "oper":
		m2 := m.(*model.SysOper)
		OperContactForCurrentUserPermissions(m2.Id)
		mf = dao.NewMultiFunction(m2)
	case "menu":
		m2 := m.(*model.SysMenu)
		MenuContactForCurrentUserPermissions(m2.Id)
		mf = dao.NewMultiFunction(m.(*model.SysMenu))
	default:
		return 0
	}

	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	ex := dao.NewExecute()
	ex.Insert(mf)

	return mf.Num

}

func SetUuid(m interface{}) {
	reflect.ValueOf(m).Elem().FieldByName("Id").SetString(uuid.Rand().Hex())
}

func SetUpdateTime(m interface{}) {
	reflect.ValueOf(m).Elem().FieldByName("UpdateTime").SetString(time.Now().Format("2006-01-02 15:04:05"))
}

//get db type
func getDbType() string {
	if util.IsNotEmpty(dbType) {
		return dbType
	}

	dbType = web.GetSession().Get("DbType").(string)
	return dbType
}

//set condition
func SetCondition(condition map[string]interface{}, m interface{}, mf *dao.MultiFunction) {

	if m == nil {
		return
	}

	if len(condition) < 1 {

		return
	}
	o := reflect.ValueOf(m).Elem()

	for k, v := range condition {
		if strings.Index(k, "_q") == -1 {
			fmt.Println(k)
			f := o.FieldByName(k)
			ft := f.Type()

			if util.IsIntT(ft) {
				v2, _ := strconv.ParseInt(v.(string), 10, 64)
				f.SetInt(v2)
			} else {
				f.SetString(v.(string))
			}
		} else {

			l := len(k)
			v3 := ""
			v4 := v.(string)
			if strings.Index(v4, "@_data_@") != -1 {
				v2 := strings.Split(v4, "@_data_@")
				v3 = conversionCondition(v2[0]) + "@_data_@" + conversionCondition(v2[1])
			} else {
				v3 = conversionCondition(v4)
			}

			mf.Rule[k[0:l-2]] = v3
		}
	}

	return
}

func conversionCondition(v string) string {
	v2 := ""
	switch v {
	case "LIKE":
		v2 = dao.LIKE
	case "EQ":
		v2 = dao.EQ
	case "GT":
		v2 = dao.GT
	case "LT":
		v2 = dao.LT
	case "GE":
		v2 = dao.GE
	case "LE":
		v2 = dao.LE
	default:
		v2 = dao.LIKE
	}
	return v2
}

//handle request param
func handleQueryListParam(param map[string]interface{}, m interface{}) *dao.MultiFunction {

	if m == nil {
		return nil
	}

	if len(param) < 1 {
		mf := NewMF(m)
		mf.IsPagin = false
		return mf
	}
	iDisplayStart := int(param["iDisplayStart"].(float64))
	iDisplayLength := int(param["iDisplayLength"].(float64))

	iSortCol_0 := int(param["iSortCol_0"].(float64))
	sortFieldName := param["mDataProp_"+strconv.Itoa(iSortCol_0)].(string)
	sortTypeString := param["sSortDir_0"].(string)

	mf := dao.NewMultiFunction(m)
	mf.DbAddr = getDbAddr()
	mf.DbType = getDbType()
	mf.SetPageSize(iDisplayLength)
	mf.SetBeginRow(iDisplayStart)

	if sortFieldName != "" && sortTypeString != "" {
		mf.OrderBy = sortFieldName
		mf.OrderByType = sortTypeString
	}

	return mf
}

func GetUserByNameAndPassword(name, password string) map[int]map[string]string {
	if util.IsEmpty(name) || util.IsEmpty(password) {
		return make(map[int]map[string]string)
	}

	m := model.NewSysUser()
	m.Name = name
	h := md5.New()
	h.Write([]byte(password))
	m.Password = hex.EncodeToString(h.Sum(nil))

	mf := NewMF(m)
	mf.Rule["Name"] = dao.EQ
	mf.Rule["Password"] = dao.EQ
	mf.IsPagin = false
	mf.IsSort = false

	ex := dao.NewExecute()
	ex.QueryAllOrByCondition(mf)

	return mf.ResultSet
}

func GetDbInfo(dbConfig string) (dbType, dbAddr string) {
	dbType, adrr, err := dao.GetDBConnectAddr(dbConfig)
	if err != nil {
		util.HandleErr(err)
	}

	return dbType, adrr
}

func GetCurrentUserId() string {

	userId = web.GetSession().Get("userId").(string)
	fmt.Println("current userId is ", userId)

	return userId

}

func GetCurrentUserGroupId() string {
	userId := GetCurrentUserId()

	user := GetUserById(userId)

	return user[0]["GroupId"]
}
