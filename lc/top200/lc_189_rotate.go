package main

import "fmt"

/*
旋转数组

*/
func rotate(nums []int, k int) {
	k = k % len(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := k, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate1(nums []int, k int) {
	k = k % len(nums)
	tmp := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate1(nums, k)
	fmt.Println(nums)
}
