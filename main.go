package main

import (
	"leiax00.cn/gocmder/conf"
	"log"
)

var args []string

func init() {
	// 通过code 设置参数
	//cmd := "usual -m=n2i"
	//cmd := "usual -m=pri -a=Hello -b=World"
	//cmd := "help"
	//args = append(args, strings.FieldsFunc(cmd, func(r rune) bool {
	//	return strings.ContainsRune(" =", r)
	//})...)
}

func main() {
	//conf.SetArgs(args...)
	err := conf.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
