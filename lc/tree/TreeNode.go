package tree

import (
	"fmt"
	"math"
)

const(
	NilNode  int = -10000
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func CreateTreeNode(arr []int) *TreeNode {
	tNodeList := make([]*TreeNode, len(arr))
	for i := 0; i < len(arr); i++ {
		tmpNode := new(TreeNode)
		if arr[i]==NilNode {
			tmpNode = nil
		} else {
			tmpNode.Val = arr[i]
		}

		tNodeList[i] = tmpNode
	}

	for i := 0; i < len(arr) / 2; i++ {
		if tNodeList[i] == nil {
			continue
		}
		if 2*i+1 < len(arr) {
			tNodeList[i].Left = tNodeList[2 * i + 1]
		}
		if 2*i+2 < len(arr) {
			tNodeList[i].Right = tNodeList[2 * i + 2]
		}
	}
	return tNodeList[0]
}

func LevelOrder(tRoot *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, tRoot)
	for len(queue) > 0 {
		first := queue[0]
		fmt.Print(first.Val, " ")
		if first.Left != nil {
			queue = append(queue, first.Left)
		}
		if first.Right != nil {
			queue = append(queue, first.Right)
		}
		queue = queue[1:]
	}
	fmt.Println()
}

func LevelOrderBottom(tRoot *TreeNode){
	queue := make([]*TreeNode, 0)
	queue = append(queue, tRoot)

	stack := make([]int, 0)
	for len(queue) > 0 {

		first := queue[0]
		stack = append([]int{first.Val}, stack...)

		if first.Right != nil {
			queue = append(queue, first.Right)
		}
		if first.Left != nil {
			queue = append(queue, first.Left)
		}

		queue = queue[1:]
	}

	for _, t := range stack {
		fmt.Print(t, " ")
	}
}

func LevelOrderBottom1(tRoot *TreeNode) [][]int {
	res := make([][]int, 0)
	if tRoot==nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, tRoot)


	for len(queue) > 0 {
		n := len(queue)
		stack := make([]int, 0)
		for n > 0 {
			first := queue[0]
			n--
			queue = queue[1:]
			stack = append([]int{first.Val}, stack...)

			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
		}
		res = append([][]int{stack}, res...)
	}
	fmt.Println(res)


	return res
}

func MaxDepth(tRoot *TreeNode) int {
	if tRoot==nil {
		return 0
	}
	return int(1 + math.Max(float64(MaxDepth(tRoot.Left)), float64(MaxDepth(tRoot.Right))))
}

