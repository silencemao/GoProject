package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
先返回最底层的叶子节点，最后返回根结点
*/
func levelOrderBottom(root *tree.TreeNode) [][]int {
	var res [][]int

	var queue []*tree.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		sz := len(queue)
		var tmp []int
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]

			tmp = append(tmp, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append([][]int{tmp}, res...)
	}

	return res
}

func main() {
	root := tree.CreateTreeNode([]int{3, 9, 20, tree.NilNode, tree.NilNode, 15, 7})
	fmt.Println(levelOrderBottom(root))
}
