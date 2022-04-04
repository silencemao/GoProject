package main

import (
	"fmt"
)

func isValid526(arr []int) bool {
	for i, val := range arr {
		if val%(i+1) != 0 && (i+1)%val != 0 {
			return false
		}
	}

	return true
}

func help526(res *int, n int, used, tmp []int) {
	if len(tmp) == n {
		if isValid526(tmp) {
			*res += 1
			//*res = append(*res, append([]int{}, tmp...))
		}
		return
	}
	for i := 1; i <= n; i++ {
		if used[i] == 1 {
			continue
		}
		used[i] = 1
		tmp = append(tmp, i)
		help526(res, n, used, tmp)
		tmp = tmp[:len(tmp)-1]
		used[i] = 0
	}
}

/*
优美的排列
假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件 之一 ，该数组就是一个 优美的排列 ：
perm[i] 能够被 i 整除
i 能够被 perm[i] 整除
给你一个整数 n ，返回可以构造的 优美排列 的 数量 。
*/
// 纯bp会超时
func countArrangement(n int) int {
	var res int
	used := make([]int, n+1)
	help526(&res, n, used, []int{})
	fmt.Println(res)
	return res
}

func backtrack526(res *int, ind, n int, tmp []int, match map[int][]int, used []int) {
	if len(tmp) == n {
		*res += 1
		return
	}
	for _, x := range match[ind] {
		if used[x] == 1 {
			continue
		}
		used[x] = 1
		tmp = append(tmp, x)
		backtrack526(res, ind+1, n, tmp, match, used)
		used[x] = 0
		tmp = tmp[:len(tmp)-1]
	}
}

func backtrack526II(res *int, ind, n int, tmp []int, match map[int][]int, used []int) {
	if ind > n { // 此处必须是ind>n
		*res += 1
		return
	}
	for _, x := range match[ind] {
		if used[x] == 1 {
			continue
		}
		used[x] = 1
		//tmp = append(tmp, x)
		backtrack526II(res, ind+1, n, tmp, match, used)
		used[x] = 0
		//tmp = tmp[:len(tmp)-1]
	}
}

//提前用match记录i和j之间的关系
func countArrangement2(n int) int {
	match := make(map[int][]int, 0)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}

	var res int
	used := make([]int, n+1)
	//backtrack526(&res, 1, n, []int{}, match, used)
	backtrack526II(&res, 1, n, []int{}, match, used)
	return res
}

func isValid526III(ind, i int) bool {
	tmp := int(i / ind)
	return tmp*ind == i
}

func help526III(res *int, ind, n int, used []int) {
	if ind > n {
		*res += 1
		return
	}
	for i := 1; i <= n; i++ {
		if used[i] == 1 {
			continue
		}
		if isValid526III(ind, i) || isValid526III(i, ind) {
			used[i] = 1
			help526III(res, ind+1, n, used)
			used[i] = 0
		}
	}
}

func countArrangement3(n int) int {
	var res int
	used := make([]int, n+1)
	help526III(&res, 1, n, used)
	return res
}

func main() {
	n := 10
	fmt.Println(countArrangement(n))
	fmt.Println(countArrangement2(n))
	fmt.Println(countArrangement3(n))
}
