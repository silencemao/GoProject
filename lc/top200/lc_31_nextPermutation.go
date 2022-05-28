package main

import "fmt"

/*
整数数组的一个 排列 就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。
更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。
如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。
必须 原地 修改，只允许使用额外常数空间。
1 2 4 3 1
1 3 1 2 4
1、在nums中从后向前找，找到nums[i]<nums[i+1]
2、从i+1位置向后找，找出大于nums[i]的最小值的索引j
3、nums[i]，nums[j]互换
4、将nums[i]后面的数字升序排列

注意nums本身是降序的情况
*/
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i < 0 {
		for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
			nums[l], nums[r] = nums[r], nums[l]
		}
		return
	}

	replaceInd := i + 1
	for j := i + 1; j < len(nums); j++ {
		if nums[j] > nums[i] && nums[j] < nums[replaceInd] {
			replaceInd = j
		}
	}
	nums[i], nums[replaceInd] = nums[replaceInd], nums[i]

	for l := i + 1; l < len(nums); l++ {
		for r := l + 1; r < len(nums); r++ {
			if nums[l] > nums[r] {
				nums[l], nums[r] = nums[r], nums[l]

			}
		}
	}
}

func main() {
	nums := []int{1, 2, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
