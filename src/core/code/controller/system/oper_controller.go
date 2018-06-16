package controller

import (
	"core/code/model/system"
	"core/code/service/system"
	"encoding/json"
	"io"
	"lzy/framework/web"
	"net/http"
	"strconv"
)

const (
	//------------ route name  ---------------
	SYSTEM_OPER_TOLIST      = "/system/oper/toList"
	SYSTEM_OPER_LIST        = "/system/oper/list"
	SYSTEM_OPER_ADD         = "/system/oper/add"
	SYSTEM_OPER_EDIT        = "/system/oper/edit"
	SYSTEM_OPER_DEL         = "/system/oper/del"
	SYSTEM_OPER_GETOPERBYID = "/system/oper/getOperById"
	SYSTEM_OPER_CONTACTMENU = "/system/oper/contactMenu"

	//------------ html name -------------------
	OPER_LIST = "system/oper.html"
)

func operConfirmContactMenu(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	pid := r.FormValue("pid")

	num := service.OperContactMenu(id, pid)

	if num > 0 {
		io.WriteString(w, strconv.FormatInt(num, 10))
	} else {
		io.WriteString(w, "")
	}
}

func getOperById(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	rows := service.GetOperById(id)

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))

}

func operAdd(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	url := r.FormValue("url")
	remark := r.FormValue("remark")

	m := model.NewSysOper()
	m.Name = name
	m.Url = url
	m.Remark = remark

	dispatcher = "operAdd"
	adapterAdd(w, m)
}

func operEdit(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	url := r.FormValue("url")
	remark := r.FormValue("remark")

	m := model.NewSysOper()
	m.Name = name
	m.Url = url
	m.Remark = remark
	m.Id = id

	dispatcher = "operEdit"
	adapterEdit(w, m)
}

func operDel(w http.ResponseWriter, r *http.Request) {

	dispatcher = "operDel"
	adapterDel(w, r)
}

//to oper list
func operToList(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("menuId")
	web.DevEnvExecuteHtml(w, OPER_LIST, id)
}

//oper list
func operListByPagin(w http.ResponseWriter, r *http.Request) {

	dispatcher = "oper"
	adapterList(w, r)
}
