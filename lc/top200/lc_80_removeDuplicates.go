package main

import "fmt"

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。

*/
func removeDuplicatesII(nums []int) int {
	pos, cnt := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[pos] {
			if cnt < 2 {
				pos, cnt = pos+1, cnt+1
				nums[pos] = nums[i]
			} else {
				continue
			}
		} else {
			pos, cnt = pos+1, 1
			nums[pos] = nums[i]
		}
	}
	return pos + 1
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 4, 4}
	pos := removeDuplicatesII(nums)
	for i := 0; i < pos; i++ {
		fmt.Println(nums[i])
	}
}
