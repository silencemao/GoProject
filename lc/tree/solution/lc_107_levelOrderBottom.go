package main

import "myGoProject/leetcode/tree"

func levelOrderBottom(root *tree.TreeNode) [][]int {
    return tree.LevelOrderBottom1(root)
}

func main() {
	tRoot := tree.CreateTreeNode([]int{3, 9, 20, tree.NilNode, tree.NilNode, 15, 7})
	//tree.LevelOrder(tRoot)
	levelOrderBottom(tRoot)

}
