package main

import "fmt"

//删除有序数组的重复项，且返回数组的长度
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	l, r := 0, 0
	for r < len(nums) {
		if nums[l] == nums[r] {
			r++
		} else {
			l++
			nums[l] = nums[r]
			r++
		}
	}
	return l + 1 // 注意此处要加一
}

func removeDuplicates1(nums []int) int {
	pos := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}

func main() {
	nums := []int{0, 0, 2}
	l := removeDuplicates1(nums)

	for i := 0; i < l; i++ {
		fmt.Println(nums[i])
	}
}
