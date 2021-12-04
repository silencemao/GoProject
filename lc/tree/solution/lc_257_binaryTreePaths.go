package main

import (
	"GoProject/leetcode/tree"
	"fmt"
	"strconv"
	"strings"
)

/*
二叉树的所有路径
*/
func help257(root *tree.TreeNode, res *[]string, tmp []string) {
	if root == nil {
		return
	}
	tmp = append(tmp, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(tmp, "->"))
		return
	}
	help257(root.Left, res, tmp)  // 状态为5-4-11时，递归下一步拿到7，加入到结果中，return 状态仍然为5-4-11
	help257(root.Right, res, tmp) // 从上一步返回，进入到这里时仍然是5-4-11 下一步拿到2 加入res，return 状态仍然是5-4-11
	tmp = tmp[:len(tmp)-1]        // 此处手动去除最后一个值，也就是去除倒数第二层节点的值，回退到倒数第三层
}

func binaryTreePaths(root *tree.TreeNode) []string {
	var res []string
	help257(root, &res, []string{})
	return res
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
	fmt.Println(binaryTreePaths(tRoot))
}
