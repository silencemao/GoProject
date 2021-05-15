package main

import (
	"fmt"
)

// 接雨水 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 暴力搜索，从当前位置向左右两侧扩散，找左边的最大值，找有边的最大值，真正能存储多少水 有左右两边最大值的较小值来决定
func trap(height []int) int {
	res := 0
	for i := 1; i < len(height)-1; i++ {
		l, r := 0, 0
		for j := i; j >= 0; j-- {
			if height[j] > l {
				l = height[j]
			}
		}
		for j := i; j < len(height); j++ {
			if height[j] > r {
				r = height[j]
			}
		}
		res += Min(l, r) - height[i] // 减去当前位置的高度
	}
	return res
}

// 双指针法
func trap1(height []int) int {
	res := 0

	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] >= lMax {
				lMax = height[l]
			} else {
				res += lMax - height[l] // 记录l位置可以存储的水
			}
			l++
		} else {
			if height[r] >= rMax {
				rMax = height[r]
			} else {
				res += rMax - height[r]
			}
			r--
		}
	}
	return res
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	fmt.Println(trap1(height))
}
