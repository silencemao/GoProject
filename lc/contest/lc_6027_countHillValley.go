package main

import "fmt"

/*
给你一个下标从 0 开始的整数数组 nums 。如果两侧距 i 最近的不相等邻居的值均小于 nums[i] ，则下标 i 是 nums 中，某个峰的一部分。类似地，如果两侧距 i 最近的不相等邻居的值均大于 nums[i] ，则下标 i 是 nums 中某个谷的一部分。对于相邻下标 i 和 j ，如果 nums[i] == nums[j] ， 则认为这两下标属于 同一个 峰或谷。

注意，要使某个下标所做峰或谷的一部分，那么它左右两侧必须 都 存在不相等邻居。

返回 nums 中峰和谷的数量。
*/
func countHillValley(nums []int) int {
	res := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		s1, s2 := -1, -1
		for s := i - 1; s >= 0; s-- {
			if nums[s] != nums[i] {
				s1 = nums[s]
				break
			}
		}
		for s := i + 1; s < len(nums); s++ {
			if nums[s] != nums[i] {
				s2 = nums[s]
				break
			}
		}

		if s1 != -1 && s2 != -1 {
			if nums[i] < s1 && nums[i] < s2 {
				res += 1
			}
			if nums[i] > s1 && nums[i] > s2 {
				res += 1
			}
		}
	}
	return res
}

func main() {
	nums := []int{57, 57, 90, 90, 85, 85, 85, 86, 86, 86}
	fmt.Println(countHillValley(nums))
}
