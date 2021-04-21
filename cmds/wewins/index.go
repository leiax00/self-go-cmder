package wewins

/* 一级命令及其子命令汇总 */
import (
	"go-demo/domain"
)

func init() {
	cmdList = append(cmdList,
		N2iCmdObj,
		PriCmdObj,
	)
	CmdWewins.ParseCmdObj2Cmd(cmdList)
}

var cmdList []*domain.CmdObj

var CmdWewins = domain.NewCommand().WithBase("wewins", "wewins")
