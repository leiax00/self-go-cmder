package domain

import (
	"fmt"
)

type ParamObj struct {
	Name      string
	ShortName string
	Value     string
	Usage     string
}

type CmdObj struct {
	Mode     string
	ArgsDef  []*ParamObj
	Args     []string
	Fn       interface{}
	ModeDesc string
}

func (cmdObj *CmdObj) GetDesc() string {
	return fmt.Sprintf("%s: %s", cmdObj.Mode, cmdObj.ModeDesc)
}
