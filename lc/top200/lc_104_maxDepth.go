package main

import (
	"GoProject/leetcode/tree"
	"GoProject/leetcode/util"
	"fmt"
)

/*
二叉树的最大深度，深度优先搜索
*/
func help104(head *tree.TreeNode) int {
	if head == nil {
		return 0
	}
	return util.MaxInt(1+help104(head.Left), 1+help104(head.Right))
}

func maxDepth(head *tree.TreeNode) int {
	res := 0
	res = help104(head)
	return res
}

/*
依照层序遍历的思路来求解
*/

func maxDepth2(head *tree.TreeNode) int {
	if head == nil {
		return 0
	}
	var queue []*tree.TreeNode
	queue = append(queue, head)
	res := 0
	for len(queue) > 0 {
		sz := len(queue)
		res++
		for sz > 0 {
			top := queue[0]
			queue = queue[1:]
			sz--
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return res
}

func createTree104(nums []int) *tree.TreeNode {
	if len(nums) < 1 {
		return nil
	}

	var nodes []*tree.TreeNode
	for i := range nums {
		if nums[i] != -100 {
			nodes = append(nodes, &tree.TreeNode{Val: nums[i]})
		} else {
			nodes = append(nodes, nil)
		}
	}
	if len(nodes) == 0 {
		return nil
	}

	for i := 0; i < len(nums)/2; i++ {
		if nodes[i] != nil {
			l, r := i*2+1, i*2+2
			nodes[i].Left, nodes[i].Right = nodes[l], nodes[r]
		}
	}
	return nodes[0]
}

func main() {
	nums := []int{3, 9, 20, -100, -100, 15, 7}
	head := createTree104(nums)
	fmt.Println(maxDepth2(head))
}
