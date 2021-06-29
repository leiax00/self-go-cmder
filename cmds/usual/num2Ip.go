package usual

import (
	"fmt"
	"leiax00.cn/gocmder/domain"
	"strconv"
	"strings"
)

func Num2Ip() {
	//export function num2ipList(num) {
	//	// 数字转ip
	//	return [(num >> 24) & 0xff, (num >> 16) & 0xff, (num >> 8) & 0xff, num & 0xff]
	//}
	numList := []int{-1062730240, -1062730177}
	for _, num := range numList {
		tmp := []string{
			strconv.Itoa((num >> 24) & 0xff),
			strconv.Itoa((num >> 16) & 0xff),
			strconv.Itoa((num >> 8) & 0xff),
			strconv.Itoa(num & 0xff),
		}
		fmt.Println(strings.Join(tmp, "."))
	}
}

var N2iCmdObj = &domain.CmdObj{
	Mode:     "n2i",
	Fn:       Num2Ip,
	ModeDesc: "数字通过二进制计算出IP",
}

// PriCmdObj 测试对象
var PriCmdObj = &domain.CmdObj{
	Mode: "pri",
	ArgsDef: []*domain.ParamObj{
		{
			Name:      "aaa",
			ShortName: "a",
			Value:     "",
			Usage:     "值A",
		},
		{
			Name:      "bbb",
			ShortName: "b",
			Value:     "",
			Usage:     "值B",
		},
	},
	Fn: func(a string, b string) {
		fmt.Printf("%s --> %s", a, b)
	},
	ModeDesc: "测试对象:a --> b",
}
