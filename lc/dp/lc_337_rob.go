package main

import (
	"GoProject/leetcode/tree"
	"GoProject/leetcode/util"
	"fmt"
)

/*
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。
除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

*/

func help337(root *tree.TreeNode) []int {
	dp := make([]int, 2) // dp[0]不取当前元素，dp[1]取当前元素
	if root == nil {
		return dp
	}
	left := help337(root.Left)
	right := help337(root.Right)
	cur := root.Val
	dp[0] = util.MaxInt(left[0], left[1]) + util.MaxInt(right[0], right[1])
	dp[1] = cur + left[0] + right[0]
	return dp
}

func rob337(root *tree.TreeNode) int {
	dp := help337(root)
	return util.MaxInt(dp[0], dp[1])
}

var set map[*tree.TreeNode]int

func rob3371(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	val1 := root.Val
	if root.Left != nil {
		val1 += rob3371(root.Left.Left) + rob3371(root.Left.Right)
	}
	if root.Right != nil {
		val1 += rob3371(root.Right.Left) + rob3371(root.Right.Right)
	}
	val2 := rob3371(root.Left) + rob3371(root.Right)
	return util.MaxInt(val1, val2)
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
	fmt.Println(rob337(tRoot))
	fmt.Println(rob3371(tRoot))
}
