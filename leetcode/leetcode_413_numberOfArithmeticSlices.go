package main

import "fmt"

/*
算术序列：至少包含三个数字，且任意两个相邻数字之间的差值相等
给定一个数组，返回这一数组包含多少个算术序列，
存在多少个，可以考虑采用动态规划

若数组长度小于3，则直接返回0
若大于3，生成一个与数组长度相等的dp数组，表示以i位置结尾的数组包含多少个算术数组。
如何判断是否为算术序列，只要满足A[i]-A[i-1] == A[i-1]-A[i-2]，则为算术序列。
从i=2开始，逐个遍历i，只要满足上面的表达式，则说明A[i]与i位置前面的数组是可以生成一个算术序列的，
即在dp[i-1]的值的基础上 + 1表示以A[i]结尾的算术序列的个数，最后对dp数组求和，即表示所有的算术序列的个数

*/
func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}
	dp := make([]int, len(A))
	if A[2] - A[1] == A[1] - A[0] {
		dp[2] = 1
	}
	res := dp[2]
	for i := 3; i < len(A); i++ {
		if A[i] - A[i-1] == A[i-1] - A[i-2] {
			dp[i] = dp[i-1] + 1
		}
		res += dp[i]
	}
	return res
}

func numberOfArithmeticSlices1(A []int) int {
	if len(A) < 3 {
		return 0
	}
	pre := 0
	if A[2] - A[1] == A[1] - A[0] {
		pre = 1
	}
	res := pre
	for i := 3; i < len(A); i++ {
		if A[i] - A[i-1] == A[i-1] - A[i-2] {
			pre = pre + 1
		} else {
			pre = 0
		}
		res += pre
	}
	return res
}

func main() {
	A := []int{1, 3, 5, 7}
	fmt.Println(numberOfArithmeticSlices(A))
	fmt.Println(numberOfArithmeticSlices1(A))
}
