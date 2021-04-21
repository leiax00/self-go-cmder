package conf

/* 命令集合汇总 */
import (
	"go-demo/cmds"
	"go-demo/domain"
)

func init() {
	rootCmd.AddCommand(cmds.GetWorkCmdList()...)
}

var rootCmd = domain.NewCommand()

func Execute() error {
	return rootCmd.Execute()
}

func SetArgs(args ...string) {
	rootCmd.SetArgs(args...)
}
