package main

import (
	"fmt"
	"sort"
)

func threeSumHelper(res *[][]int, tmp *[]int, used *[]bool, nums []int, index int) {
	if len(*tmp)==3 {
		cur := 0
		for _, num := range *tmp {
			cur += num
		}
		if cur==0 {
			*res = append(*res, []int{(*tmp)[0], (*tmp)[1], (*tmp)[2]})
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			*tmp = append(*tmp, nums[i])
			threeSumHelper(res, tmp, used, nums, index+1)
			(*used)[i] = false
			*tmp = (*tmp)[:len(*tmp)-1]
		}
	}
}
func threeSum(nums []int) [][]int {
	var res [][]int
	var tmp []int
	used := make([]bool, len(nums))
	threeSumHelper(&res, &tmp, &used, nums, 0)
	return res
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums) - 1
		for left < right {
			if nums[i] + nums[left] + nums[right]==0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left]==nums[left+1] {
					left++
				}
				for left < right && nums[right]==nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i] + nums[left] + nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func main() {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum1(nums))
}
