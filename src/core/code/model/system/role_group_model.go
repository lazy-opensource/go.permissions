package model

/**
用户组关联角色中间表
**/
type SysRoleGroup struct {
	Id      string
	GroupId string
	RoleId  string
	tableN  string
}

func NewSysRoleGroup() *SysRoleGroup {

	return &SysRoleGroup{tableN: "sys_role_group"}
}
