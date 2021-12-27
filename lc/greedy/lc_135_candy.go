package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

/*
错误
*/
func candy(ratings []int) int {
	res := 0
	nums := make([]int, len(ratings))
	nums[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		} else if ratings[i] == ratings[i-1] {
			if nums[i-1] > 1 {
				nums[i] = 1
			} else {
				nums[i] = nums[i-1] + 1
			}
		} else {
			if nums[i-1] > 1 {
				nums[i] = 1
			} else {
				nums[i] = 1
				for j := i - 1; j >= 0 && ratings[j] > ratings[j+1] && nums[j] <= nums[j+1]; j-- {
					nums[j] += 1
				}
			}
		}
	}
	for _, v := range nums {
		res += v
	}
	fmt.Println(ratings)
	fmt.Println(nums)
	return res
}

/*
1、从前向后遍历一次
2、从后向前遍历一次
3、求和
*/
func candy1(ratings []int) int {
	nums := make([]int, len(ratings))
	nums[0] = 1
	for i := 1; i < len(ratings); i++ {
		nums[i] = 1
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			nums[i] = util.MaxInt(nums[i], nums[i+1]+1)
		}
	}
	fmt.Println(ratings)
	fmt.Println(nums)
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

func main() {
	ratings := []int{1, 2, 3, 5, 4, 3, 2, 1, 4, 3, 2, 1, 3, 2, 1, 1, 2, 3, 4}
	fmt.Println(candy(ratings))
	fmt.Println(candy1(ratings))
}
