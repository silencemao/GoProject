package main

import (
	"fmt"
	"myGoProject/leetcode/tree"
)

/*
深度优先遍历
层序遍历二叉树
*/
func levelOrder(root *tree.TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}
	queue := []*tree.TreeNode{root}
	for len(queue) > 0 {
		var tmp []int
		size := len(queue)
		for i := 0; i < size; i++ {
			first := queue[i]
			tmp = append(tmp, first.Val)
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}
		queue = queue[size:]
		res = append(res, tmp)
	}

	return res
}

func levelOrder1(root *tree.TreeNode) []int {
	var res []int

	queue := []*tree.TreeNode{root}
	for len(queue) > 0 {
		first := queue[0]
		res = append(res, first.Val)
		if first.Left != nil {
			queue = append(queue, first.Left)
		}
		if first.Right != nil {
			queue = append(queue, first.Right)
		}
		queue = queue[1:]
	}

	return res
}

func main() {
	root := tree.CreateTreeNode([]int{3, 9, 20, tree.NilNode, tree.NilNode, 15, 7})
	fmt.Println(levelOrder1(root))
	fmt.Println(levelOrder(root))
}
