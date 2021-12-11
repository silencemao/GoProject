package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func sumOfLeftLeaves(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func sumOfLeftLeaves1(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves1(root.Left)
	rightValue := sumOfLeftLeaves1(root.Right)

	leaveNum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leaveNum = root.Left.Val
	}
	sum := leaveNum + leftValue + rightValue
	return sum
}

func sumOfLeftLeaves2(root *tree.TreeNode) int {
	var stack []*tree.TreeNode
	if root == nil {
		return 0
	}
	res := 0
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Left != nil && top.Left.Left == nil && top.Left.Right == nil {
			res += top.Left.Val
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

func sumOfLeaves(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return sumOfLeaves(root.Left) + sumOfLeaves(root.Right)
}

/*
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1

*/

func main() {
	tRoot := tree.CreateTreeNode([]int{5, 4, 8, 11, tree.NilNode, 13, 4, 7, 2, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, 1})
	fmt.Println(sumOfLeftLeaves(tRoot))
	fmt.Println(sumOfLeftLeaves1(tRoot))
	fmt.Println(sumOfLeftLeaves2(tRoot))
	//fmt.Println(sumOfLeaves(tRoot))
}
