package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。
排序+双指针
*/
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0

	for i := len(nums) - 1; i >= 0; i-- { // 要先固定最大值
		l, r := 0, i-1
		for l < r {
			tmp := nums[l] + nums[r] - nums[i]
			if tmp > 0 {
				res += r - l
				r -= 1
			} else {
				l += 1
			}
		}
	}
	return res
}

func help611(res *int, nums []int, tmp []int, ind int) {
	if len(tmp) == 3 {
		sort.Ints(tmp)
		if tmp[0]+tmp[1]-tmp[2] > 0 {
			*res += 1
		}
		return
	}
	for i := ind; i < len(nums); i++ {
		if len(tmp) == 2 && tmp[0]+tmp[1] <= nums[i] {
			continue
		}
		tmp = append(tmp, nums[i])
		help611(res, nums, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

/*
超时
*/
func triangleNumber1(nums []int) int {
	res := 0
	sort.Ints(nums)
	help611(&res, nums, []int{}, 0)
	//fmt.Println(res)
	return res
}

// 二分法
func triangleNumber2(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			l, r := 0, j-1
			for l < r {
				mid := l + (r-l)>>1
				if nums[mid]+nums[j] > nums[i] {
					r = mid
				} else {
					l = mid + 1
				}
			}
			if r >= 0 && nums[r]+nums[j] > nums[i] {
				res += j - r
			}
		}
	}
	return res
}

func main() {
	nums := []int{48, 66, 61, 46, 94, 75}
	fmt.Println(triangleNumber(nums))
	fmt.Println(triangleNumber1(nums))
	fmt.Println(triangleNumber2(nums))
}
