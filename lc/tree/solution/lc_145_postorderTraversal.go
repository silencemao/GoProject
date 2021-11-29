package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func postOrderHelper(pRoot *tree.TreeNode, pNums *[]int) {
	if pRoot == nil {
		return
	}

	postOrderHelper(pRoot.Left, pNums)
	postOrderHelper(pRoot.Right, pNums)
	*pNums = append(*pNums, pRoot.Val)
}

func postOrderTraversal(root *tree.TreeNode) []int {
	var pNums []int
	postOrderHelper(root, &pNums)
	return pNums
}

func postOrderTraversal1(root *tree.TreeNode) []int {
	var pNums []int
	if root == nil {
		return pNums
	}

	var pLast tree.TreeNode
	var pNodes []tree.TreeNode
	for root != nil || len(pNodes) > 0 {
		if root != nil {
			pNodes = append([]tree.TreeNode{*root}, pNodes...)
			root = root.Left
		} else {
			pTop := pNodes[0]
			if pTop.Right != nil && pLast != *pTop.Right {
				root = pTop.Right
			} else {
				pNums = append(pNums, pTop.Val)
				pLast = pTop
				pNodes = pNodes[1:]
			}
		}
	}
	return pNums
}

func postorderTraversal(root *tree.TreeNode) []int {
	var stack []*tree.TreeNode
	var res []int
	var pLast *tree.TreeNode

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			top := stack[len(stack)-1]
			if top.Right != nil && top.Right != pLast {
				// stack = append(stack, top.Right) 此处不能把top.right加进来，否则会重复
				root = top.Right
			} else {
				res = append(res, top.Val)
				stack = stack[:len(stack)-1]
				pLast = top
			}
		}
	}
	return res
}

func main() {
	tRoot := tree.CreateTreeNode([]int{5, 4, 8, 11, tree.NilNode, 13, 4, 7, 2, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode, 1})
	fmt.Println(postOrderTraversal(tRoot))
	fmt.Println(postOrderTraversal1(tRoot))
	fmt.Println(postorderTraversal(tRoot))
}
