package main

import (
	"GoProject/leetcode/tree"
)

/*
	在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
	如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
	我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
	只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。

给一个二叉树，给两个节点，判断两个节点是否属于同一层，但是父节点不同
*/

type NodeInfo struct {
	fa    int
	cur   *tree.TreeNode
	depth int
}

func bfs993(root *tree.TreeNode, val int) (int, int) {
	var queue []*NodeInfo
	queue = append(queue, &NodeInfo{fa: 0, cur: root, depth: 0})

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]

			if front.cur.Val == val {
				return front.fa, front.depth
			}

			if front.cur.Left != nil {
				queue = append(queue, &NodeInfo{fa: front.cur.Val, cur: front.cur.Left, depth: front.depth + 1})
			}
			if front.cur.Right != nil {
				queue = append(queue, &NodeInfo{fa: front.cur.Val, cur: front.cur.Right, depth: front.depth + 1})
			}
		}
	}
	return -1, -1
}

func isCousins(root *tree.TreeNode, x int, y int) bool {
	fa1, depth1 := bfs993(root, x)
	fa2, depth2 := bfs993(root, y)
	if fa1 == -1 || fa2 == -1 {
		return false
	}
	if fa1 != fa2 && depth1 == depth2 {
		return true
	}
	return false
}

func main() {

}
