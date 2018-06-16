package model

/**
系统角色操作按钮中间表结构体
*/
type SysRoleOper struct {
	Id         string
	RoleId     string
	OperId     string
	UpdateTime string
	tableN     string
}

/**
初始化角色操作
*/
func NewSysRoleOper() *SysRoleOper {
	return &SysRoleOper{tableN: "sys_role_oper"}
}
