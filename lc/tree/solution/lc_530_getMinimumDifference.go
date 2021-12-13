package main

import (
	"GoProject/leetcode/tree"
	"fmt"
	"math"
)

func helpInorder(root *tree.TreeNode, vec *[]int) {
	if root == nil {
		return
	}
	helpInorder(root.Left, vec)
	*vec = append(*vec, root.Val)
	helpInorder(root.Right, vec)
}

func getMinimumDifference(root *tree.TreeNode) int {
	res := 1<<63 - 1
	var vec []int
	helpInorder(root, &vec)
	for i := 1; i < len(vec); i++ {
		if vec[i]-vec[i-1] < res {
			res = vec[i] - vec[i-1]
		}
	}
	return res
}

func help530(root, pre *tree.TreeNode, res *int) {
	if root == nil {
		return
	}
	help530(root.Left, pre, res)
	if pre != nil && root.Val-pre.Val < *res {
		*res = root.Val - pre.Val
	}
	//fmt.Println(root.Val)
	pre = root // todo bug
	help530(root.Right, pre, res)
}

func getMinimumDifference1(root *tree.TreeNode) int {
	res := 1<<63 - 1
	var pre *tree.TreeNode
	help530(root, pre, &res)
	return res
}

// 中序遍历的同时计算最小值
func getMinimumDifference2(root *tree.TreeNode) int {
	// 保留前一个节点的指针
	var prev *tree.TreeNode
	// 定义一个比较大的值
	min := math.MaxInt64
	var travel func(node *tree.TreeNode)
	travel = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return min
}

func main() {
	tRoot := tree.CreateTreeNode([]int{5, 4, 7, tree.NilNode, tree.NilNode, tree.NilNode, tree.NilNode})

	fmt.Println(getMinimumDifference1(tRoot))
	fmt.Println(getMinimumDifference2(tRoot))

}
