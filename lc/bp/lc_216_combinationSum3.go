package main

import "fmt"

/*
找出所有相加之和为n 的k个数的组合。组合中只允许含有 1-9 的正整数，并且每种组合中不存在重复的数字。
说明：
所有数字都是正整数。
解集不能包含重复的组合。
*/
func help216(res *[][]int, tmp []int, k, n, ind int, sum int) {
	if len(tmp) == k {
		if sum == n {
			*res = append(*res, append([]int{}, tmp...))
			return
		}
	}
	for i := ind; i <= 9 && sum+i <= n; i++ {
		sum += i
		tmp = append(tmp, i)
		help216(res, tmp, k, n, i+1, sum)
		sum -= i
		tmp = tmp[:len(tmp)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	help216(&res, []int{}, k, n, 1, 0)
	return res
}

func main() {
	k, n := 3, 9
	fmt.Println(combinationSum3(k, n))
}
