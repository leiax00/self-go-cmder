package practice

import (
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	//lastLeftNode := findLastLeftNode(t)

}

func IsRoot(t *tree.Tree) bool {
	return false
}

func findLastLeftNode(t *tree.Tree) *tree.Tree  {
	tmp := t
	for t.Left != nil {
		tmp = t.Left
	}
	return tmp
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func IterTree()  {
	ch := make(chan int)
	Walk(tree.New(1), ch)
}