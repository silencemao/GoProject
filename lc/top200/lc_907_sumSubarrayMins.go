package main

import "fmt"

/*
给定一个整数数组 arr，找到 min(b)的总和，其中 b 的范围为 arr 的每个（连续）子数组。由于答案可能很大，因此 返回答案模 10^9 + 7 。

输入：arr = [3,1,2,4]
输出：17
解释：
子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。

避免重复 左边界开区间，右边界闭区间
arr=[1,2,4,2,3,1]
第一个2的左右边界是(0, 3) -> [0, 3]
第二个2的左右边界是(0, 5) -> [0, 5]
左边找到第一个比当前元素小的索引
右边找到第一个比当前元素小于等于的索引

单调栈
以i为中心，向左向右找到第一个比自己小的元素索引
*/
func sumSubarrayMins(arr []int) int {
	res, n := 0, len(arr)
	const mod int = 1e9 + 7
	l, r := make([]int, n), make([]int, n)
	for i := range l {
		l[i], r[i] = -1, n
	}
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			r[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			l[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		res = (res + (i-l[i])*(r[i]-i)*arr[i]) % mod
	}

	return res

}

//两次遍历

func sumSubarrayMins1(arr []int) int {
	res, n := 0, len(arr)
	const mod int = 1e9 + 7
	l, r := make([]int, n), make([]int, n)
	stack := []int{-1}
	for i := range r {
		l[i], r[i] = -1, n
	}

	for i := 0; i < n; i++ {
		for len(stack) > 1 && arr[stack[len(stack)-1]] >= arr[i] { // 栈顶元素遇到第一个比其小的元素，则当前元素相当于栈顶元素的右边界
			r[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		l[i] = stack[len(stack)-1]
		stack = append(stack, i)
	}
	fmt.Println(arr)
	fmt.Println(l)
	fmt.Println(r)
	for i := 0; i < n; i++ {
		res = (res + (i-l[i])*(r[i]-i)*arr[i]) % mod
	}
	return res
}

func main() {
	arr := []int{1, 2, 4, 2, 3, 1}
	fmt.Println(sumSubarrayMins(arr))
	fmt.Println(sumSubarrayMins1(arr))
}
