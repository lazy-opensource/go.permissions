package model

/**
系统用户角色中间表结构体
*/
type SysRoleUser struct {
	Id         string
	UserId     string
	RoleId     string
	UpdateTime string
	tableN     string
	CreateTime string
}

/**
初始化用户角色
*/
func NewSysRoleUser() *SysRoleUser {
	return &SysRoleUser{tableN: "sys_role_user"}
}
