package main

import (
	"GoProject/leetcode/util"
	"fmt"
)

func help440(res *[]int, cur, n int) {
	if cur > n {
		return
	}
	*res = append(*res, cur)

	for i := 0; i <= 9; i++ {
		help440(res, cur*10+i, n)
	}
}

func findKthNumber(n int, k int) int {
	var res []int
	for i := 1; i <= 9 && i <= n; i++ {
		help440(&res, i, n)
		if len(res) > k {
			break
		}
	}
	fmt.Println(res)
	return res[k-1]
}

/*
计算以prefix为前缀的数字，其中小于等于n的数字的个数
算prefix和prefix+1分别为前缀，二者之间的数字
prefix=1 prefix+1=2，n=21
[1 10 11 12 13 14 15 16 17 18 19 2 20 21 3 4 5 6 7 8 9]
1 2之间的数字数字有1个
10， 20之间的数字有10个
二者之间的数字个数有 1 + 10 个
*/
func getCount(prefix, n int) int {
	cur, next, cnt := prefix, prefix+1, 0
	for cur <= n {
		cnt += util.MinInt(next, n+1) - cur
		cur, next = cur*10, next*10
	}
	return cnt
}

func findKthNumber1(n int, k int) int {
	p := 1
	prefix := 1
	for p < k { // 此处不能取等号，p=k时立马退出
		cnt := getCount(prefix, n)
		// 此处不能包含等于，因为p+cnt=k时，那结果的长度肯定和prefix一致，不能再*10了
		if p+cnt > k { // 第k个数字在以prefix为前缀的子树中，所以prefix向前移动一位(11 -> 110)，p计数加一
			prefix *= 10
			p += 1
		} else { // 不在以prefix为前缀的子树中，prefix+1，p计数加cnt
			prefix += 1
			p += cnt
		}
	}
	return prefix
}

func main() {
	n, k := 13, 4
	fmt.Println(findKthNumber(n, k))
	fmt.Println(findKthNumber1(n, k))
}
