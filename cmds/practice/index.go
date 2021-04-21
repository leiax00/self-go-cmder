package practice

/* 一级命令及其子命令汇总 */
import (
	"go-demo/domain"
)

func init() {
	cmdList = append(cmdList,
		Pra01CmdObj,
		IterTreeCmdObj,
	)
	CmdPra.ParseCmdObj2Cmd(cmdList)
}

var cmdList []*domain.CmdObj

var CmdPra = domain.NewCommand().WithBase("pra", "practice")
