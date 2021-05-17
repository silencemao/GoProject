package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func levelOrderTraversal(root *tree.TreeNode) []int {
	var res []int
	var queue []*tree.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		res = append(res, top.Val)
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return res
}

func levelOrderTraversal1(root *tree.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue [][]*tree.TreeNode
	queue = append(queue, []*tree.TreeNode{root})
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		var num []int
		var tmp []*tree.TreeNode
		for _, n := range top {
			num = append(num, n.Val)
			if n.Left != nil {
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
			}
		}
		if len(num) > 0 {
			res = append(res, num)
		}
		if len(tmp) > 0 {
			queue = append(queue, tmp)
		}
	}
	return res
}

func preOrderTraversal(root *tree.TreeNode) []int {
	var res []int

	var stack []*tree.TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

func inOrderTraversal(root *tree.TreeNode) []int {
	var res []int
	var stack []*tree.TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		root = top.Right
	}
	return res
}

func postOrderTraversal(root *tree.TreeNode) []int {
	var res []int
	var stack1, stack2 []*tree.TreeNode
	stack1 = append(stack1, root)
	for len(stack1) > 0 {
		top := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, top)
		if top.Right != nil {
			stack1 = append(stack1, top.Right)
		}
		if top.Left != nil {
			stack1 = append(stack1, top.Left)
		}
	}
	for len(stack2) > 0 {
		res = append(res, stack2[len(stack2)-1].Val)
		stack2 = stack2[:len(stack2)-1]
	}
	return res
}

func main1() {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 5,
			},
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 6,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(levelOrderTraversal(root))
	fmt.Println(levelOrderTraversal1(root))
	fmt.Println(preOrderTraversal(root))
	fmt.Println(inOrderTraversal(root))
	fmt.Println(postOrderTraversal(root))

}
