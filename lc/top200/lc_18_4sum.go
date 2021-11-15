package main

import (
	"fmt"
	"sort"
)
// -2 -1 0 0 1 2
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i+1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] { // j的起点是i+1，而不是0，所以 此处j的判断逻辑是j>i+1而不是j>0
				continue
			}

			sum := nums[i] + nums[j]
			if sum > target { // 此处不行，当target为负数，nums存在正负数的情况
				//continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				sub := target - nums[l] - nums[r] - sum
				if sub == 0 {
						res = append(res, []int{nums[l], nums[r], nums[i], nums[j]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sub > 0 {
					l++
				} else {
					r--
				}
			}
		}
	}

	return res
}
func main() {
	nums := []int{1,-2,-5,-4,-3,3,3,5}

	target := -11
	fmt.Println(fourSum(nums, target))
}
