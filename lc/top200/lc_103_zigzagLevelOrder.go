package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
z字形遍历二叉树

按照层序遍历的思路往队列中添加节点，将节点的值取出来num，
然后利用一个flag控制num放入res中的顺序，正序or逆序
*/
func zigzagLevelOrder(root *tree.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue [][]*tree.TreeNode
	queue = append(queue, []*tree.TreeNode{root})
	flag := true
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		var num []int
		var tmp []*tree.TreeNode
		for _, n := range top {
			num = append(num, n.Val)
			if n.Left != nil {
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
			}
		}
		if len(tmp) > 0 {
			queue = append(queue, tmp)
		}
		if !flag {
			for i := 0; i < len(num)/2; i++ {
				num[i], num[len(num)-1-i] = num[len(num)-1-i], num[i]
			}
		}
		res = append(res, num)
		flag = !flag
	}
	return res
}

func main() {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 5,
			},
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 6,
			},
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}
