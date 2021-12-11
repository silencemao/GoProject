package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。
切分新的中序 后序遍历数组时区间全部维持左闭右开
*/
func buildTree(inorder []int, postorder []int) *tree.TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootNum := postorder[len(postorder)-1]
	root := &tree.TreeNode{Val: rootNum}
	leftEnd := 0
	for i := range inorder {
		if inorder[i] == rootNum {
			leftEnd = i
			break
		}
	}
	l1, r1 := 0, leftEnd                // 闭 开
	l2, r2 := leftEnd+1, len(postorder) // 闭 开
	leftInorder := inorder[l1:r1]
	rightInOrder := inorder[l2:r2]

	leftPostorder := postorder[0:r1]
	rightPostorder := postorder[r1 : len(postorder)-1]

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInOrder, rightPostorder)
	return root
}
func main() {
	inorder := []int{4, 6, 3, 5, 1, 7, 2}
	postorder := []int{4, 3, 6, 1, 2, 7, 5}
	root := buildTree(inorder, postorder)
	fmt.Println(root.Val)
}
