package dao

import (
	"core/code/model"
	"fmt"
	"lzy/framework/util"
	"testing"
	"time"
)

func ATestInsert(t *testing.T) {
	sysUser := model.NewSysUser()
	sysUser.Id = "4"
	sysUser.Name = "zhangsan"
	sysUser.Passworld = "123"
	sysUser.Email = "zhangsan@.com"
	sysUser.Status = 1
	sysUser.Remark = "test"
	sysUser.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	sysUser.UpdateTime = time.Now().Format("2006-01-02 15:04:05")

	//--------------------
	fmt.Println(util.GetCurrentPath())
	mf := util.NewMultiFunction(sysUser)
	mf.TableN = sysUser.TableN
	//	mf.DbFilePath = "F:/Go-Workspace/permissions/src/core/code/conf/db.properties"
	mf.DbFilePath = "../../../core/code/conf/db.properties"
	//--------------------

	ex := NewExecute()
	msg := ex.Insert(mf)
	if msg.Code == 200 {
		fmt.Println(mf.Tag)
		msg.Message = "success insert into " + mf.TableN + string(mf.Tag) + " data"
	}
	if msg.Code == 500 {
		fmt.Println("------------------")
		//		fmt.Println(msg.Message)
	}
}

func ATestDelete(t *testing.T) {
	sysUser := model.NewSysUser()

	//---------------------
	mf := util.NewMultiFunction(sysUser)
	mf.TableN = sysUser.TableN
	mf.DbFilePath = "F:/Go-Workspace/permissions/src/core/code/conf/db.properties"
	//----------------------

	ex := NewExecute()
	msg := ex.Delete(mf)
	if msg.Code == 200 {
		fmt.Println(mf.Tag)
		msg.Message = "success delete from " + mf.TableN + string(mf.Tag) + " data"
	}
	if msg.Code == 500 {
		fmt.Println("------------------")
		fmt.Println(msg.Message)
	}
}

func ATestQueryAllOrByCondition(t *testing.T) {

}

func ATestUpdata(t *testing.T) {

}

func TestCount(t *testing.T) {
	sysUser := model.NewSysUser()

	mf := util.NewMultiFunction(sysUser)
	mf.TableN = sysUser.TableN
	mf.DbFilePath = "F:/Go-Workspace/permissions/src/core/code/conf/db.properties"

	ex := NewExecute()
	ex.Count(mf)
	fmt.Println(mf.Total)
}
