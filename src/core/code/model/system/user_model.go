package model

//用户结构体
type SysUser struct {
	Id         string
	Name       string
	Remark     string
	Status     string
	UpdateTime string
	CreateTime string
	tableN     string
	Email      string
	Password   string
	Oper       string
	GroupId    string
}

func NewSysUser() *SysUser {
	o := new(SysUser)
	o.tableN = "sys_user"
	return o
}
