package model

/**
系统角色结构体
*/
type SysRole struct {
	Id         string
	Name       string
	Remark     string
	Status     int
	UpdateTime string
	CreateTime string
	tableN     string
	Oper       string
}

/**
初始化角色
*/
func NewSysRole() *SysRole {
	return &SysRole{tableN: "sys_role"}
}
