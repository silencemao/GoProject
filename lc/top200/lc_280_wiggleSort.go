package main

import (
	"fmt"
	"sort"
)

/*
摆动排序 nums[0] <= nums[1] >= nums[2] <= nums[4]

允许相邻的数字相等，所以我们可以在排序之后，两两进行交换，不需要从最后一个数字开始逐渐插入到数组中
*/
func wiggleSortI(nums []int) {
	sort.Ints(nums)

	for i := 2; i < len(nums); i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
	fmt.Println(nums)
}

/*
根据奇数、偶数位置不同的值来判断
奇数位置的值满足 nums[i] >= nums[i-1]
偶数位置的值满足 nums[i] <= nums[i-1]
若不满足 可以直接交换i与i-1位置处的值
*/
func wiggleSortI2(nums []int) {
	for i := 1; i < len(nums); i++ {
		if (i%2 == 1 && nums[i] < nums[i-1]) || (i%2 == 0 && nums[i] > nums[i-1]) {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{6, 5, 4, 3, 2, 1}
	//wiggleSortI(nums)
	wiggleSortI2(nums)
}
