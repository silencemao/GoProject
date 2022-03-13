package main

import "fmt"

/*
n 位格雷码序列 是一个由 2n 个整数组成的序列，其中：
每个整数都在范围 [0, 2n - 1] 内（含 0 和 2n - 1）
第一个整数是 0
一个整数在序列中出现 不超过一次
每对 相邻 整数的二进制表示 恰好一位不同 ，且
第一个 和 最后一个 整数的二进制表示 恰好一位不同
给你一个整数 n ，返回任一有效的 n 位格雷码序列 。

*/
func grayCode(n int) []int {
	//ans := make([]int, 1<<n)
	ans := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		ans[i] = gray(i)
	}
	return ans
}

// 数字转换为格雷码 公式 n xor (n>>1)
func gray(n int) int {
	return n ^ (n >> 1)
}

/*
格雷码 G(n)
G(n+1) = G(n)' U R(n)'
G(n)'为在G(n)序列中每一个数首位添加0
R(n)'为在G(n)逆序之后每一个数首位添加1 逆序是指，先遍历G(n)中最后一个数，再遍历倒数地个数
*/
func grayCode1(n int) []int {
	ans := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, head+ans[j])
		}
		head = head << 1
	}
	return ans
}

func main() {
	n := 2
	fmt.Println(grayCode(n))
	fmt.Println(grayCode1(n))
}
