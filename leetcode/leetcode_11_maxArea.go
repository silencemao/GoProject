package main

import (
	"math"
)
/*
最大蓄水面积
给定一个数组a，计算以(i, ai), (j, aj)为坐标组成的最大面积
*/
func maxArea(height []int) int {
	res := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			tmp := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

func maxArea1(height []int) int {
	res := 0
	left, right := 0, len(height) - 1
	for left < right {
		tmp := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if tmp > res {
			res = tmp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

//func main() {
//	var nums = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
//	fmt.Println(maxArea(nums), maxArea1(nums))
//}