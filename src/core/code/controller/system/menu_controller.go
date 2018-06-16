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
	SYSTEM_MENU_TOLIST                       = "/system/menu/toList"
	SYSTEM_MENU_LIST                         = "/system/menu/list"
	SYSTEM_MENU_ADD                          = "/system/menu/add"
	SYSTEM_MENU_EDIT                         = "/system/menu/edit"
	SYSTEM_MENU_DEL                          = "/system/menu/del"
	SYSTEM_MENU_LISTBYCONDITION              = "/system/menu/listByCondition"
	SYSTEM_MENU_MENUBYID                     = "/system/menu/getMenuById"
	SYSTEM_MENU_MENUCONFIRMCONTACTPARENTMENU = "/system/menu/menuConfirmContactParentMenu"

	//------------ html name -------------------
	MENU_LIST = "system/menu.html"
)

func menuConfirmContactParentMenu(w http.ResponseWriter, r *http.Request) {
	pids := r.FormValue("pids")
	mainIds := r.FormValue("mainIds")

	num := service.MenuConfirmContactParentMenu(pids, mainIds)

	io.WriteString(w, strconv.FormatInt(num, 10))
}

func menuAdd(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	url := r.FormValue("url")
	remark := r.FormValue("remark")
	icon := r.FormValue("icon")

	m := model.NewSysMenu()
	m.Name = name
	m.Remark = remark
	m.ParentId = "0"
	m.Url = url
	m.Icon = icon

	dispatcher = "menuAdd"
	adapterAdd(w, m)
}

func menuEdit(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	name := r.FormValue("name")
	remark := r.FormValue("remark")
	icon := r.FormValue("icon")

	m := model.NewSysMenu()
	m.Id = id
	m.Name = name
	m.Remark = remark
	m.Icon = icon

	dispatcher = "menuEdit"
	adapterEdit(w, m)
}

func getMenuById(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	rows := service.GetMenuById(id)

	data, _ := json.Marshal(rows)
	io.WriteString(w, string(data))
}

func menuDel(w http.ResponseWriter, r *http.Request) {

	dispatcher = "menuDel"
	adapterDel(w, r)
}

//to menu list
func menuToList(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("menuId")

	web.DevEnvExecuteHtml(w, MENU_LIST, id)
}

//menu list
func menuList(w http.ResponseWriter, r *http.Request) {

	menus := service.GetCurrentUserMenus()

	data, _ := json.Marshal(menus)
	io.WriteString(w, string(data))
}

func menuListByCondition(w http.ResponseWriter, r *http.Request) {

	dispatcher = "menu"
	adapterList(w, r)
}
