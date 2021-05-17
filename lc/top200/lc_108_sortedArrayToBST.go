package main

import (
	"GoProject/leetcode/tree"
	"fmt"
)

/*
给定一个排序数组，转换成二叉搜索树，且左右子树的高度差不超过1
方案一：
要求高度差不超过1，那我们就从中间分，中间的值作为根结点，将数组左边的数据作为左子树的节点，右边的数据作为右子树的节点
有了一个根结点root 和两个数组nums1，nums2，然后nums1去构建左子树，构建的过程和上面的过程一致(一个根结点，两段数组)
                                        nums2构建右子树
依次类推，整个树构建完毕
*/

func helper108(root *tree.TreeNode, nums1, nums2 []int) {
	if len(nums1) > 0 {
		root.Left = &tree.TreeNode{Val: nums1[len(nums1)/2]}
		helper108(root.Left, nums1[:len(nums1)/2], nums1[len(nums1)/2+1:])
	}
	if len(nums2) > 0 {
		root.Right = &tree.TreeNode{Val: nums2[len(nums2)/2]}
		helper108(root.Right, nums2[:len(nums2)/2], nums2[len(nums2)/2+1:])
	}
}

func sortedArrayToBST(nums []int) *tree.TreeNode {
	if len(nums) < 1 {
		return nil
	}
	root := &tree.TreeNode{Val: nums[len(nums)/2]}
	helper108(root, nums[:len(nums)/2], nums[len(nums)/2+1:])
	return root
}

func dfs(nums []int, l, r int) *tree.TreeNode {
	if l > r {
		return nil
	}
	mid := l + (r-l)>>1
	root := &tree.TreeNode{Val: nums[mid]}
	root.Left = dfs(nums, l, mid-1)
	root.Right = dfs(nums, mid+1, r)
	return root
}

func sortedArrayToBST1(nums []int) *tree.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return dfs(nums, 0, len(nums)-1)
}

func main() {
	nums := []int{-10, 3, 0, 5, 9}
	//root := sortedArrayToBST(nums)
	root := sortedArrayToBST1(nums)

	fmt.Println(root)
}
