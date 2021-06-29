package usual

/* 一级命令及其子命令汇总 */
import (
	"leiax00.cn/gocmder/domain"
)

func init() {
	cmdList = append(cmdList,
		N2iCmdObj,
		PriCmdObj,
	)
	CmdUsual.ParseCmdObj2Cmd(cmdList)
}

var cmdList []*domain.CmdObj

var CmdUsual = domain.NewCommand().WithBase("usual", "usual")
