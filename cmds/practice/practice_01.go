package practice

import (
	"fmt"
	"go-demo/domain"
	"golang.org/x/tour/pic"
)

func calc1(x int, y int) uint8 {
	return uint8((x + y) / 2)
}

func calc2(x int, y int) uint8 {
	return uint8(x * y)
}

func calc3(x int, y int) uint8 {
	return uint8(x ^ y)
}

func calc4(x int, y int) uint8 {
	return uint8(x % (y + 1))
}

func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx)
	fmt.Println(dy)
	rst := make([][]uint8, dy)
	for i := range rst {
		item := make([]uint8, dx)
		for j := range item {
			item[j] = calc1(i, j)
			item[j] = calc2(i, j)
			item[j] = calc3(i, j)
			item[j] = calc4(i, j)
		}
		rst[i] = item
	}
	return rst
}

func PicShow() {
	pic.Show(Pic)
}

var Pra01CmdObj = &domain.CmdObj{
	Mode:     "pra_01",
	Fn:       PicShow,
	ModeDesc: "练习: 绘图测试",
}
