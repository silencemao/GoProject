package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
返回所有二叉树的根节点
*/
func helper(start, end int) []*tree.TreeNode {
	if start > end {
		return []*tree.TreeNode{nil}
	}
	var allTree []*tree.TreeNode
	for i := start; i <= end; i++ { // 以i为根节点
		leftNode := helper(start, i-1) //构造左子树
		rightNode := helper(i+1, end)  //构造右子树

		for _, l := range leftNode {
			for _, r := range rightNode {
				cur := &tree.TreeNode{Val: i}
				cur.Left = l
				cur.Right = r
				allTree = append(allTree, cur)
			}
		}
	}
	return allTree
}

func generateTrees(n int) []*tree.TreeNode {
	if n < 1 {
		return nil
	}
	return helper(1, n)
}

func main() {
	n := 2
	res := generateTrees(n)
	fmt.Println(res)

}
