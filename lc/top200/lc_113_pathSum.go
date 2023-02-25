package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

func help113(root *tree.TreeNode, res *[][]int, tmp []int, sum int) {
	if root == nil {
		return
	}
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		*res = append(*res, append([]int{}, tmp...))
		return
	}
	help113(root.Left, res, tmp, sum-root.Val)
	help113(root.Right, res, tmp, sum-root.Val)
	tmp = tmp[:len(tmp)-1]
}

func pathSumII(root *tree.TreeNode, sum int) [][]int {
	var res [][]int
	help113(root, &res, []int{}, sum)
	return res
}

/*
         5
        / \
       4   8
      /   / \
     11  13  4
    /  \    / \
   7    2  5   1
*/
func main() {
	arr := []int{5, 4, 8, 11, tree.NilNode, 13, 4, 7, 2, tree.NilNode, tree.NilNode, 5, 1}
	root := tree.Ints2Node(arr)
	tree.LevelOrder1(root)
	sum := 22
	fmt.Println(pathSumII(root, sum))
}
