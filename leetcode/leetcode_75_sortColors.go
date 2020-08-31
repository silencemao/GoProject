package main

import "fmt"

func sortColors(nums []int) {
	second := len(nums) - 1
	zero := 0
	for i := 0; i <= second; i++ {
		for nums[i]==2 && i < second {
			nums[i], nums[second] = nums[second], nums[i]
			second--
		}
		for nums[i]==0 && i > zero {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
}

func main() {
	nums := []int{2, 0, 1, 2, 0, 1}
	fmt.Println(nums)
	sortColors(nums)
	fmt.Println(nums)
}
