package main

import (
	"GoProject/leetcode/tree"
	"GoProject/leetcode/util"
	"fmt"
)

/*
二叉树的最小深度，从根节点到叶子节点的最小深度
   1
  2
最小深度为2，不是1
*/
func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	} else if root.Right == nil {
		return 1 + minDepth(root.Left)
	} else {
		return 1 + util.MinInt(minDepth(root.Left), minDepth(root.Right))
	}
}

func minDepth1(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	var queue []*tree.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
	}

	return depth
}

func main() {
	tRoot := tree.CreateTreeNode([]int{1, 2, 3, 4, 5})
	tree.LevelOrder(tRoot)
	fmt.Println(minDepth(tRoot))
	fmt.Println(minDepth1(tRoot))
}
