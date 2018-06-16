package model

type SysGroupMenu struct {
	Id      string
	GroupId string
	MenuId  string
	tableN  string
}

func NewSysGroupMenu() *SysGroupMenu {

	return &SysGroupMenu{tableN: "sys_group_menu"}
}
