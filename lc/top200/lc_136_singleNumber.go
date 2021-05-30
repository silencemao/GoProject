package main

import (
	"fmt"
	"sort"
)

// 数组中只有一个数字只出现一次，其余数字均出现两次，找出只出现一次的数字
func singleNumber(nums []int) int {
	sort.Ints(nums)
	res := -1
	for i := 0; i < len(nums); {
		if i+1 == len(nums) {
			res = nums[i]
			break
		}
		if nums[i] != nums[i+1] {
			res = nums[i]
			break
		}
		i += 2
	}
	return res
}

// 异或 a^b^a = b
func singleNumber2(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	nums := []int{2, 2, 1, 1, 3, 4, 4}
	fmt.Println(singleNumber(nums))
	fmt.Println(singleNumber2(nums))
}
