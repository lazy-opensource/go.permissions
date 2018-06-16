package model

type SysGroupOper struct {
	Id      string
	GroupId string
	OperId  string
	tableN  string
}

func NewSysGroupOper() *SysGroupOper {

	return &SysGroupOper{tableN: "sys_group_oper"}
}
