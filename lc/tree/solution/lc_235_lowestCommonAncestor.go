package main

import "GoProject/leetcode/tree"

/*
二叉搜索树的最近公共祖先
*/
func lowestCommonAncestorBST(root, p, q *tree.TreeNode) *tree.TreeNode {
	for root != nil {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}

func help235(root, p, q *tree.TreeNode) *tree.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		left := help235(root.Left, p, q)
		if left != nil {
			return left
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		right := help235(root.Right, p, q)
		if right != nil {
			return right
		}
	}
	return root
}

/*
普通二叉树的搜索公共祖先
*/
func help235_1(root, p, q *tree.TreeNode) *tree.TreeNode {
	left := help235(root.Left, p, q)
	right := help235(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return left
	}
	return nil
}

func lowestCommonAncestorBST1(root, p, q *tree.TreeNode) *tree.TreeNode {
	return help235(root, p, q)
	//return help235_1(root, p, q)
}

func main() {

}
