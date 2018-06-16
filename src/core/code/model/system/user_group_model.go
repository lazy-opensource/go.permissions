package model

/**
系统用户组/部门
*/
type SysUserGroup struct {
	Id         string
	Name       string
	CreateTime string
	UpdateTime string
	Remark     string
	tableN     string
}

func NewSysUserGroup() *SysUserGroup {

	return &SysUserGroup{tableN: "sys_user_group"}
}
