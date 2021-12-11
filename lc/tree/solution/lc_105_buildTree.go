package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
前序和中序构造二叉树
1、中序最后一位为根节点，
2、只从前序中找到根结点的索引
3、生成对中序数组切分的索引，左子树的l1 r1 右子树的l2 r2
*/
func buildTree2(preorder []int, inorder []int) *tree.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootNum := preorder[0]
	root := &tree.TreeNode{Val: rootNum}
	leftEnd := 0
	for i := range inorder {
		if inorder[i] == rootNum {
			leftEnd = i
			break
		}
	}

	l1, r1 := 0, leftEnd
	l2, r2 := leftEnd+1, len(inorder)

	leftInorder := inorder[l1:r1]
	rightInorder := inorder[l2:r2]

	leftPreorder := preorder[1 : r1+1]
	rightPreorder := preorder[r1+1:]

	root.Left = buildTree2(leftPreorder, leftInorder)
	root.Right = buildTree2(rightPreorder, rightInorder)
	return root
}

func main() {
	preorder := []int{5, 6, 4, 3, 7, 1, 2}
	inorder := []int{4, 6, 3, 5, 1, 7, 2}

	root := buildTree2(preorder, inorder)
	fmt.Println(root.Val)
}
