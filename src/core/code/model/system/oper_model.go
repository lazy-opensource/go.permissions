package model

/**
系统操作按钮结构体
*/
type SysOper struct {
	Id         string
	Name       string
	Remark     string
	Status     int
	UpdateTime string
	CreateTime string
	tableN     string
	Url        string
	MenuId     string
	Oper       string
	Expression string
}

/**
初始化操作
*/
func NewSysOper() *SysOper {
	return &SysOper{tableN: "sys_oper"}
}
