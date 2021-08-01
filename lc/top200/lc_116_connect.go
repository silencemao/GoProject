package main

import "GoProject/leetcode/tree"

/*
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有next 指针都被设置为 NULL。

填充二叉树的next节点，即将左节点的next 赋值为 右节点
bfs搞定
*/
func connect(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	var queue []*tree.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {

		sz := len(queue)
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[1].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		for i := 1; i < sz; i++ {
			cur := queue[i]
			queue[i-1].Next = cur

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[sz:]
	}
	return root
}
func main() {

}
