package cmds

/* 一级命令分类汇总 */
import (
	"leiax00.cn/gocmder/cmds/practice"
	"leiax00.cn/gocmder/cmds/usual"
	"leiax00.cn/gocmder/domain"
)

func init() {
	cmdList = append(cmdList,
		usual.CmdUsual,
		practice.CmdPra,
	)
}

var cmdList []*domain.SelfCommand

func GetWorkCmdList() []*domain.SelfCommand {
	return cmdList
}
