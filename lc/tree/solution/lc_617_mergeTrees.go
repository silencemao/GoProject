package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func mergeTrees(root1 *tree.TreeNode, root2 *tree.TreeNode) *tree.TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
		//return &tree.TreeNode{Val: root2.Val}
	}
	if root2 == nil {
		return root1
		//return &tree.TreeNode{Val: root1.Val}
	}

	root := &tree.TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees(root1.Left, root2.Left) // 上面注释处的写法就结束了递归，则进入求解right(下一行)，同样下一行结束后就返回了，这样结果就不对
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

func main() {
	tRoot1 := tree.CreateTreeNode([]int{1, 2, tree.NilNode, 3, tree.NilNode, tree.NilNode, tree.NilNode})
	tRoot2 := tree.CreateTreeNode([]int{1, tree.NilNode, 2, tree.NilNode, tree.NilNode, tree.NilNode, 3})

	res := mergeTrees(tRoot1, tRoot2)
	fmt.Println(res.Val)
	//levelOrderBottom(res)
}
