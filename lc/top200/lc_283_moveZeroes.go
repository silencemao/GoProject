package main

import "fmt"

/*
移动0，给定数组，将0全部移动至尾部，非0数字保持原来的顺序

而重遍历，以i为结尾时，遇到第一个0 就将其和i位置进行交换，
将0逐步移动到最后
*/
func moveZeroes(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		for j := 0; j < i; j++ {
			if nums[j] == 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
}

/*
pos计数 遍历数组，遇到非0的数字就放到pos位置处，最后pos以后的位置都填充为0
*/
func moveZeroes1(nums []int) {
	n := len(nums)
	pos := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}
	for ; pos < n; pos++ {
		nums[pos] = 0
	}
}

func main() {
	nums := []int{0, 1, 2, 0, 3, 5}
	moveZeroes1(nums)
	fmt.Println(nums)
}
