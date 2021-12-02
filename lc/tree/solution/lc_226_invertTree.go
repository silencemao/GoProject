package main

import "GoProject/leetcode/tree"

func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

func help(root *tree.TreeNode) {
	root.Left, root.Right = root.Right, root.Left
}

func invertTree1(root *tree.TreeNode) *tree.TreeNode {
	var queue []*tree.TreeNode
	if root == nil {
		return root
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			help(front)
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
	}
	return root
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
	inOrderTraversal(tRoot)

	tRoot = invertTree(tRoot)
	inOrderTraversal(tRoot)
}
