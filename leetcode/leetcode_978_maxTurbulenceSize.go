package main

import (
	"fmt"
)

/*
最长湍流子数组，湍流 是指数组中相邻的数交替递增与递减
即 [2, 4, 1, 5, 3, 10] 为湍流的数组，长度为6
给定一个数组arr，求出其湍流数组的长度
*/

/*
动态规划，生成两个长度与arr一致的数组 inc，dec
inc记录以i结尾&&arr[i] > arr[i-1]的长度
dec记录以i结尾&&arr[i] < arr[i-1]的长度

若arr[i] > arr[i-1] 表明当前是递增，所以我们应该在 i-1 位置是递减的前提的加一 inc[i] = dec[i-1]+1 
既然当前i位置是递增，那么对于dec数组来讲是不满足湍流情况的，所以相当于递减的数组重新开始。即dec[i]=1
每次计算完inc dec，重新计算res的结果
*/
func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	res := 0
	inc, dec := make([]int, len(arr)), make([]int, len(arr))
	inc[0], dec[0] = 1, 1
	for i := 1; i < len(arr); i++ {
		inc[i], dec[i] = 1, 1
		if arr[i] > arr[i-1] {
			inc[i] = dec[i-1] + 1
		} else if arr[i] < arr[i-1] {
			dec[i] = inc[i-1] + 1
		}
	}
	for i := 0; i < len(arr); i++ {
		if inc[i] > dec[i] {
			if inc[i] > res {
				res = inc[i]
			}
		} else {
			if dec[i] > res {
				res = dec[i]
			}
		}
	}
	return res
}

/*
可以不使用数组来记录结果，改用两个变量inc dec来记录状态变化，当其中一个变化时，记得对另一个赋值为1
*/
func maxTurbulenceSize1(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	inc, dec, res := 1, 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			inc = dec + 1
			dec = 1
		} else if arr[i] < arr[i-1] {
			dec = inc + 1
			inc = 1
		} else {
			inc, dec = 1, 1
		}
		if inc > dec && inc > res {
			res = inc
		} else if dec > inc && dec > res {
			res = dec
		}
	}
	return res
}

func main() {
	arr := []int{4, 8, 12, 16}
	fmt.Println(maxTurbulenceSize(arr))
	fmt.Println(maxTurbulenceSize1(arr))
}
