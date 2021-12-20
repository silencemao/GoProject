package main

import (
	"GoProject/leetcode/tree"
)

/*
删除二叉搜索树中的节点
*/
func deleteNode(root *tree.TreeNode, key int) *tree.TreeNode {
	if root == nil {
		return root
	} else if root.Val == key { // 找到节点和key的值相等
		if root.Left == nil && root.Right == nil { // 被删除节点的左右子树都为空
			return nil
		} else if root.Left == nil { // 被删除节点的左子树为空
			return root.Right
		} else if root.Right == nil { // 被删除节点的右子树为空
			return root.Left
		} else { // 被删除节点的左右子树都不为空
			cur := root.Right     // 找到被删除节点的右子树
			for cur.Left != nil { // 找到被删除节点右子树最左边的节点
				cur = cur.Left
			}
			cur.Left = root.Left // 将被删除节点的左子树 作为 被删除节点的右子树 的 最左边节点的 子节点
			root = root.Right    // 删除当前节点
			return root
		}
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

func deleteNode1(root *tree.TreeNode, key int) *tree.TreeNode {
	if root == nil {
		return root
	} else if root.Val == key {
		if root.Right == nil {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		root.Val, cur.Val = cur.Val, root.Val
	}
	root.Left = deleteNode1(root.Left, key)
	root.Right = deleteNode1(root.Right, key)
	return root
}

func main() {
	tRoot := tree.CreateTreeNode([]int{5, 3, 6, 2, 4, tree.NilNode, 7, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode})
	deleteNode(tRoot, 3)
}
