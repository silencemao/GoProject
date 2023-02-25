package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
找到树的最后一层的最左边的叶子节点的值
层序遍历
*/
func findBottomLeftValue(root *tree.TreeNode) int {
	var queue []*tree.TreeNode
	res := root.Val
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			if i == 0 {
				res = front.Val
			}
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
	}
	return res
}

var maxDep int
var leftValue int

func help513(root *tree.TreeNode, deep int) {
	if root.Left == nil && root.Right == nil {
		if deep > maxDep {
			maxDep = deep
			leftValue = root.Val
		}
		return
	}
	if root.Left != nil {
		help513(root.Left, deep+1)
	}
	if root.Right != nil {
		help513(root.Right, deep+1)
	}
}
func findBottomLeftValue1(root *tree.TreeNode) int {
	help513(root, 0)
	return leftValue
}

func help5131(root *tree.TreeNode, deep int) {
	if root == nil {
		return
	}

	if root.Right != nil {
		help5131(root.Right, deep+1)
	}
	if root.Left != nil {
		help5131(root.Left, deep+1)
	}
	if root.Left == nil && root.Right == nil {
		if deep > maxDep {
			maxDep = deep
			leftValue = root.Val
		}
	}

}

func findBottomRightValue(root *tree.TreeNode) int {
	help5131(root, 0)
	return leftValue
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
	//fmt.Println(findBottomLeftValue1(tRoot))
	//fmt.Println(findBottomLeftValue(tRoot))
	fmt.Println(findBottomRightValue(tRoot))
}
