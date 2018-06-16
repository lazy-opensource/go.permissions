package model

/**
系统菜单结构体
*/
type SysMenu struct {
	Id         string
	Name       string
	Remark     string
	Status     int
	UpdateTime string
	CreateTime string
	tableN     string
	ParentId   string
	Url        string
	Oper       string
	Icon       string
}

/**
初始化菜单
*/
func NewSysMenu() *SysMenu {
	return &SysMenu{tableN: "sys_menu"}
}
