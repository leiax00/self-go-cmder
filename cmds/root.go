package cmds

/* 一级命令分类汇总 */
import (
	"go-demo/cmds/usual"
	"go-demo/domain"
)

func init() {
	cmdList = append(cmdList,
		usual.CmdWewins,
	)
}

var cmdList []*domain.SelfCommand

func GetWorkCmdList() []*domain.SelfCommand {
	return cmdList
}
