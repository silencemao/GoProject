package main

import "fmt"

/*
给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。

输入：n = 13
输出：[1,10,11,12,13,2,3,4,5,6,7,8,9]

*/
func help386(res *[]int, n, cur int) {
	if cur > n {
		return
	}
	*res = append(*res, cur)

	for i := 0; i <= 9; i++ {
		cur = cur*10 + i
		help386(res, n, cur)
		cur = cur / 10
	}
}

func lexicalOrder(n int) []int {
	var res []int
	for i := 1; i <= 9 && i <= n; i++ {
		help386(&res, n, i)
	}

	return res
}

func lexicalOrder2(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n { //直接在后面加0
			num *= 10
		} else { //加0会超出n的范围，
			// n=21，num=19时，字典序下一个数字应该是2，相当于去掉个位，十位先加1
			// n=21, num=21时，字典序下一位是3，去掉个位，十位先加1
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return ans
}

func main() {
	fmt.Println(lexicalOrder(21))
	fmt.Println(lexicalOrder2(21))
}
