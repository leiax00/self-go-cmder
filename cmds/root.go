package cmds

/* 一级命令分类汇总 */
import (
	"go-demo/cmds/wewins"
	"go-demo/domain"
)

func init() {
	cmdList = append(cmdList,
		wewins.CmdWewins,
	)
}

var cmdList []*domain.SelfCommand

func GetWorkCmdList() []*domain.SelfCommand {
	return cmdList
}
