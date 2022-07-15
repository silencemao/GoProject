package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

二叉树转换为全部向右的树

1、前序遍历
2、找前驱节点
*/

func help114(root *tree.TreeNode) []*tree.TreeNode {
	var stack []*tree.TreeNode
	var res []*tree.TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top)

		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		if top.Left != nil {
			stack = append(stack, top.Left)
		}

	}
	return res
}
func flatten1(root *tree.TreeNode) {
	res := help114(root)
	for i, _ := range res {
		fmt.Println(res[i].Val)
	}

	for i := 1; i < len(res); i++ {
		pre, cur := res[i-1], res[i]
		pre.Left, pre.Right = nil, cur
	}
}

func flatten(root *tree.TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Right = next
			cur.Left = nil
		}
		cur = cur.Right
	}
}

/*
      1
     /  \
    2     5
   / \    / \
  3   4  7   8

1-2-3-4-5-7-8
*/

func main() {
	tRoot := tree.CreateTreeNode([]int{1, 2, 5, 3, 4, 7, 8, tree.NilNode, tree.NilNode, tree.NilNode})
	//tRoot := tree.CreateTreeNode([]int{1, tree.NilNode, tree.NilNode})
	tree.LevelOrderBottom1(tRoot)
	flatten1(tRoot)
	tree.LevelOrderBottom1(tRoot)

}
