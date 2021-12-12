package main

import (
	"fmt"
)

/*
https://leetcode-cn.com/problems/sum-of-subarray-ranges/

给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。
返回 nums 中 所有 子数组范围的 和 。
子数组是数组中一个连续 非空 的元素序列。

输入：nums = [1,2,3]
输出：4
解释：nums 的 6 个子数组如下所示：
[1]，范围 = 最大 - 最小 = 1 - 1 = 0
[2]，范围 = 2 - 2 = 0
[3]，范围 = 3 - 3 = 0
[1,2]，范围 = 2 - 1 = 1
[2,3]，范围 = 3 - 2 = 1
[1,2,3]，范围 = 3 - 1 = 2
所有范围的和是 0 + 0 + 0 + 1 + 1 + 2 = 4

给定一个数组，找出所有的子数组，将子数组最大最小值的的差值加起来，返回
1、暴力解法，dfs找出所有子数组 超时
2、双重循环，外层循环遍历数组，内层循环 寻找以i为起点，长度为2、3、4...的子数组，计算子数组的最大值最小值，将差值累加
判断最大最小值时不必再遍历子数组，以min，max记录最小最大值，每当子数组长度加一之后，只需要对比新加入的数字与最大最小值的大小即可

*/
func combine5953(nums []int, tmp []int, res *[][]int, ind int) {
	if len(tmp) > 1 {
		*res = append(*res, append([]int{}, tmp...))
	}
	for i := ind; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		combine5953(nums, tmp, res, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
func help1(nums []int) int64 {
	var res2 int64
	for s := 0; s < len(nums); s += 1 {
		min, max := nums[s], nums[s]           // 以s为起点的最小最大值
		for k := 2; k <= len(nums)-s; k += 1 { // 子数组的长度
			e := s + k
			if min <= nums[e-1] && nums[e-1] <= max { // 若没有这个判断就会超时

			} else { // 更新最小最大值
				if nums[e-1] > max {
					max = nums[e-1]
				}
				if nums[e-1] < min {
					min = nums[e-1]
				}
			}

			res2 += int64(max - min)
		}
	}
	return res2
}

func help2(nums []int) int64 { // 暴力解法
	var res int64
	for i, num := range nums {
		min, max := num, num
		for _, v := range nums[i+1:] {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
			res += int64(max - min)
		}
	}
	return res
}

func subArrayRanges(nums []int) int64 {
	return help1(nums)
}

func main() {
	nums := []int{4, -2, -3, 4, 1}
	fmt.Println(subArrayRanges(nums))
}
