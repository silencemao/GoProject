package main

import "fmt"

/*
和为k的连续子数组的个数
*/
func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums)+1) // 前缀和 preSum[i]表示0-i-1之间的数值之和
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	res := 0
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				res += 1
			}
		}
	}
	return res
}

func subarraySum1(nums []int, k int) int {
	set := make(map[int]int, 0) // 记录前缀和的个数
	set[0] = 1
	res, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		subNum := sum - k
		if cnt, ok := set[subNum]; ok { // 先计算是否存在subNum
			res += cnt
		}

		if cnt, ok := set[sum]; ok { // 再将前缀和放入set中
			set[sum] = cnt + 1
		} else {
			set[sum] = 1
		}

	}
	return res
}

func main() {
	nums := []int{1}
	k := 0
	fmt.Println(subarraySum(nums, k))
	fmt.Println(subarraySum1(nums, k))
}
