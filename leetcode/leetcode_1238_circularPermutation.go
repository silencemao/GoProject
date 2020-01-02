package main

import (
	"fmt"
)
/*
给定一个数n和一个数start，n代表二进制数的位数，start代表第一个数。生成一个数组，
以start开始，每个数的二进制表示与前一个数字的二进制表示只有一位不同。每个数字的二进制长度均为n
使用格雷码，格雷码与二进制数的关系是G=B^(B>>1)

*/
func circularPermutation(n, start int) []int {
	var res []int
	if n < 0 || start > 1<<n {
		return res
	}
	for i := 0; i < 1<<n; i++ {
		res = append(res, i^(i>>1))
	}
	for res[0] != start {
		res = append(res, res[0])
		res = res[1:]
	}
	return res
}

func circularDfs(res *[]int, flag *[]bool, index, n, start int) bool {
	if index==0 {
		return true
	}
	for i := 0; i < n; i++ {
		num := start ^ (1<<i)
		if !(*flag)[num] {
			*res = append(*res, num)
			(*flag)[num] = true
			if circularDfs(res, flag, index-1, n, num) {
				return true
			}
			(*flag)[num] = false
		}
	}
	return false
}

func circularPermutation1(n, start int) []int {
	flag := make([]bool, 1<<n)
	var res []int
	res = append(res, start)
	flag[start] = true
	circularDfs(&res, &flag, (1<<n)-1, n, start)
	return res
}

func main() {
	fmt.Println(circularPermutation(3, 2))
	fmt.Println(circularPermutation1(3, 2))
}