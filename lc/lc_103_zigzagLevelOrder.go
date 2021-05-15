package main

import (
	"fmt"
	"myGoProject/leetcode/tree"
)

func zigzagLevelOrder(root *tree.TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}
	direct := 0
	queue := []*tree.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var tmp = make([]int, size)
		for i := 0; i < size; i++ {
			first := queue[i]

			if direct%2 == 0 {
				tmp[i] = first.Val
			} else {
				tmp[size-i-1] = first.Val
			}
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}
		res = append(res, tmp)
		queue = queue[size:]
		direct++
	}
	return res
}

func main() {
	root := tree.CreateTreeNode([]int{1, 2, 3, 4, tree.NilNode, tree.NilNode, 5})
	//fmt.Println(levelOrder(root))
	fmt.Println(zigzagLevelOrder(root))
}
