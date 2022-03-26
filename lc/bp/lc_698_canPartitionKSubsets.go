package main

import (
	"fmt"
	"sort"
)

/*
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
示例 1：

输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
*/
func canPartitionSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	//sort.Ints(nums)
	used := make([]bool, len(nums))
	target := sum / k
	return backtrack(k, 0, nums, []int{}, 0, used, target)
}

func backtrack(k int, bucket int, nums, tmp []int, start int, used []bool, target int) bool {
	if k == 0 {
		return true
	}
	if bucket == target {
		//fmt.Println(tmp)
		tmp = tmp[:0]
		return backtrack(k-1, 0, nums, tmp, 0, used, target)
	}

	for i := start; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if bucket+nums[i] > target {
			continue
		}
		used[i] = true
		bucket += nums[i]
		tmp = append(tmp, nums[i])
		if backtrack(k, bucket, nums, tmp, i+1, used, target) {
			return true
		}
		used[i] = false
		bucket -= nums[i]
		tmp = tmp[:len(tmp)-1]
	}

	return false
}

func help6982(res *[][]int, ind, target, k, curSum int, used []bool, nums []int) bool {
	if k == 1 {
		return true
	}
	if curSum == target {
		return help6982(res, 0, target, k-1, 0, used, nums)
	}
	for i := ind; i < len(nums); i++ {
		if used[i] || curSum+nums[i] > target {
			continue
		}
		used[i] = true
		if help6982(res, i+1, target, k, curSum+nums[i], used, nums) {
			return true
		}
		used[i] = false
	}
	return false
}

func canPartitionKSubsets2(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	var res [][]int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	return help6982(&res, 0, sum/k, k, 0, used, nums)
}

func main() {
	nums := []int{3, 2, 1, 3, 6, 1, 4, 8, 10, 8, 9, 1, 7, 9, 8, 1}
	k := 9
	fmt.Println(canPartitionKSubsets2(nums, k))
	fmt.Println(canPartitionSubsets(nums, k))
}
