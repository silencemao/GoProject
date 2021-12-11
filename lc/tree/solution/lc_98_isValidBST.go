package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
验证二叉搜索树

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func help98(root *tree.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val < min || root.Val > max {
		return false
	}
	return help98(root.Left, min, root.Val) && help98(root.Right, root.Val, max)
}

func isValidBST(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	min, max := -1<<63, 1<<63-1
	return help98(root, min, max)
}

/*

      5
     / \
    4   6
   /   / \
      3   7


*/

func main() {
	tRoot := tree.CreateTreeNode([]int{5, 4, 6, tree.NilNode, tree.NilNode, 3, 7, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode})
	tree.LevelOrder(tRoot)
	fmt.Println(isValidBST(tRoot))
}
