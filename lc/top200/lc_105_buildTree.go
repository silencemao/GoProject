package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func build106(preOrder []int, p_s, p_e int, inOrder []int, i_s, i_e int) *tree.TreeNode {
	if p_s == p_e {
		return nil
	}
	root := &tree.TreeNode{Val: preOrder[p_s]}

	i := i_s
	for ; i < i_e; i++ {
		if inOrder[i] == preOrder[p_s] {
			break
		}
	}

	leftNum := i - i_s                                                    // 左子树的长度
	root.Left = build106(preOrder, p_s+1, p_s+leftNum+1, inOrder, i_s, i) // 右边界一律不取
	root.Right = build106(preOrder, p_s+leftNum+1, p_e, inOrder, i+1, i_e)

	return root
}

func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}
	head := build106(preorder, 0, len(preorder), inorder, 0, len(inorder))
	return head
}

func levelOrder105(head *tree.TreeNode) {
	if head == nil {
		return
	}
	var queue []*tree.TreeNode
	queue = append(queue, head)
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			top := queue[i]
			fmt.Print(top.Val, " ")
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		fmt.Println()
		queue = queue[sz:]
	}
}

func main() {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	head := buildTree(preOrder, inOrder)
	levelOrder105(head)
}
