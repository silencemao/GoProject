package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
[10, 2] -> 210
给定一个数组，返回该数组可以组成的最大数字

其实就是排序，两个数字a、b如何排序，才能是最后组成的数字最大，a+b > b+a
*/
func largestNumber(nums []int) string {
	res := ""
	sort.SliceStable(nums, func(i, j int) bool {
		a, b := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return a+b > b+a
	})
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	k := 0
	for ; k < len(res)-1; k++ { // 防止全为0
		if res[k] != '0' {
			break
		}
	}
	return res[k:]
}

func main() {
	nums := []int{0, 0}
	fmt.Println(largestNumber(nums))
}
