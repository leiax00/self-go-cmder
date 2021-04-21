package domain

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"reflect"
	"strings"
)

type SelfCommand struct {
	command *cobra.Command
	mode    string
	modeMap map[string]*CmdObj
}

func NewCommand() *SelfCommand {
	cmd := &SelfCommand{
		command: &cobra.Command{},
		modeMap: make(map[string]*CmdObj),
	}
	cmd.SetRunFn()
	return cmd
}

func (cmd *SelfCommand) WithBase(use string, short string) *SelfCommand {
	cmd.command.Use = use
	cmd.command.Short = short
	return cmd
}

func (cmd *SelfCommand) ParseCmdObj2Cmd(cmdList []*CmdObj) {
	cmd.command.Flags().StringVarP(&cmd.mode, "mode", "m", "", "请输入操作的模式")
	modeDesc := []string{fmt.Sprintf("%s 命令模式:", cmd.command.Use)}
	for _, item := range cmdList {
		modeDesc = append(modeDesc, item.GetDesc())
		cmd.modeMap[item.Mode] = item
		if item.ArgsDef == nil {
			continue
		}
		item.Args = make([]string, len(item.ArgsDef))
		for i := 0; i < len(item.Args); i++ {
			cmd.command.Flags().StringVarP(&item.Args[i], item.ArgsDef[i].Name, item.ArgsDef[i].ShortName, item.ArgsDef[i].Value, item.ArgsDef[i].Usage)
		}
	}
	cmd.command.Long = strings.Join(modeDesc, "\n")
}

func (cmd *SelfCommand) SetRunFn() {
	cmd.command.Run = func(c *cobra.Command, args []string) {
		cmdObj, ok := cmd.modeMap[cmd.mode]
		if !ok || reflect.ValueOf(cmdObj.Fn).Kind() != reflect.Func {
			log.Fatalf("暂不支持该命令，请执行 help " + cmd.command.Use + " 查看帮助文档")
		}
		vargs := make([]reflect.Value, len(cmdObj.Args))
		for i, arg := range cmdObj.Args {
			vargs[i] = reflect.ValueOf(arg)
		}
		reflect.ValueOf(cmdObj.Fn).Call(vargs)
	}
}

func (cmd *SelfCommand) Execute() error {
	return cmd.command.Execute()
}

func (cmd *SelfCommand) SetArgs(args ...string) {
	cmd.command.SetArgs(args)
}

func (cmd *SelfCommand) AddCommand(cmds ...*SelfCommand) {
	for _, item := range cmds {
		cmd.command.AddCommand(item.command)
	}
}
