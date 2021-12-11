package main

import "GoProject/leetcode/tree"

/*
给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

给定一个二叉树 和 一个值，返回二叉树中以该值为根节点的树，若不存在 返回nil
*/
func searchBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	queue := []*tree.TreeNode{root}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front.Val == val {
			return front
		}
		if front.Left != nil {
			queue = append(queue, front.Left)
		}
		if front.Right != nil {
			queue = append(queue, front.Right)
		}
	}
	return nil
}

func searchBST1(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST1(root.Right, val)
	}
	if root.Val > val {
		return searchBST1(root.Left, val)
	}
	return nil
}

func searchBST2(root *tree.TreeNode, val int) *tree.TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}

func main() {

}
