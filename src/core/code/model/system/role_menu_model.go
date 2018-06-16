package model

/**
系统角色菜单中间表结构体
*/
type SysRoleMenu struct {
	Id         string
	RoleId     string
	MenuId     string
	UpdateTime string
	tableN     string
}

/**
初始化角色菜单
*/
func NewSysRoleMenu() *SysRoleMenu {
	return &SysRoleMenu{tableN: "sys_role_menu"}
}
