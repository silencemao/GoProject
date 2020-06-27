package main

import (
	"fmt"
	"myGoProject/leetcode/tree"
)

func preOrderHelper(pRoot *tree.TreeNode, pNums *[]int) {
	if pRoot == nil {
		return
	}
	*pNums = append(*pNums, pRoot.Val)
	preOrderHelper(pRoot.Left, pNums)
	preOrderHelper(pRoot.Right, pNums)
}

func preOrderTraversal(root *tree.TreeNode) []int {
	var pNums []int
	preOrderHelper(root, &pNums)
	return pNums
}

func preOrderTraversal1(root *tree.TreeNode) []int {
	var pNums []int
	if root == nil {
		return pNums
	}
	var pNodes []tree.TreeNode
	pNodes = append(pNodes, *root)
	for len(pNodes) > 0 {
		pTop := pNodes[0]
		pNums = append(pNums, pTop.Val)

		pNodes = pNodes[1:]

		if pTop.Right != nil {
			pNodes = append([]tree.TreeNode{*pTop.Right}, pNodes...)
		}

		if pTop.Left != nil {
			pNodes = append([]tree.TreeNode{*pTop.Left}, pNodes...)
		}
	}
	return pNums
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
	fmt.Println(preOrderTraversal(tRoot))

	fmt.Println(preOrderTraversal1(tRoot))
}