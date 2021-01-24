package main

import "fmt"

/*
格雷码
给定数字n，生成长度为n的二进制数，且相邻的两个数的二进制之间只有一位不同

此题需要先明确两个问题，第一个如何保证相邻两个数的二进制只有一位不同。（下一个数字由上一个数字转换而来，利用位运算或 |）
                    第二个从n=1开始，生成二进制长度为1的，再生成二进制长度为2的，逐步生成
当我们计算n=2的时候，我们需要将1向右移动1位，然后分别于n=1的数字进行或运算
*/
func grayCode(n int) []int {
	var res []int
	res = append(res, 0)
	for i := 0; i < n; i++ {
		size := len(res)
		for j := size - 1; j >= 0; j-- {
			res = append(res, res[j]|1<<i)
		}
	}
	return res
}

// g[i] = i ^ (i>>1)

func main() {
	n := 2
	fmt.Println(grayCode(n))
}
