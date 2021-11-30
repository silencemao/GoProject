package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func inOrderHelper(pRoot *tree.TreeNode, pNums *[]int) {
	if pRoot == nil {
		return
	}
	inOrderHelper(pRoot.Left, pNums)
	*pNums = append(*pNums, pRoot.Val)
	inOrderHelper(pRoot.Right, pNums)
}

func inOrderTraversal(root *tree.TreeNode) []int {
	var res []int
	inOrderHelper(root, &res)
	fmt.Println(res)
	return res
}

func inOrderTraversal1(root *tree.TreeNode) []int {
	var res []int

	var tNodes []tree.TreeNode
	for root != nil || len(tNodes) > 0 {
		for root != nil {
			tNodes = append([]tree.TreeNode{*root}, tNodes...)
			root = root.Left
		}
		pTop := tNodes[0]
		res = append(res, pTop.Val)
		root = pTop.Right
		tNodes = tNodes[1:]
	}
	return res
}

func inOrderTraversal2(root *tree.TreeNode) []int {
	var res []int
	var tNodes []tree.TreeNode
	for root != nil || len(tNodes) > 0 {
		if root != nil {
			tNodes = append([]tree.TreeNode{*root}, tNodes...)
			root = root.Left
		} else {
			pTop := tNodes[0]
			res = append(res, pTop.Val)
			root = pTop.Right
			tNodes = tNodes[1:]
		}
	}
	return res
}

func inOrderTraversal3(root *tree.TreeNode) []int {
	var res []int
	var stack []*tree.TreeNode

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, top.Val)
			root = top.Right
		}
	}
	return res
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

func main1() {
	tRoot := tree.CreateTreeNode([]int{5, 4, 8, 11, tree.NilNode, 13, 4, 7, 2, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, 1})
	inOrderTraversal(tRoot)
	fmt.Println(inOrderTraversal1(tRoot))
	fmt.Println(inOrderTraversal2(tRoot))
}
