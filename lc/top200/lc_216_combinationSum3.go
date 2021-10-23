package main

import "fmt"

/*
从1-9中选出k个数，其和为n，返回这些组合，每个组合中不能使用重复的数字
就是在[1,2,3,4,5,6,7,8,9]这个集合中找到和为n的k个数的组合。
*/

func help216(res *[][]int, tmp []int, k, n, ind, sum int) {
	if len(tmp) == k { //控制个数
		if sum == n { // 控制求和
			*res = append(*res, append([]int{}, tmp...))
		}
		return
	}

	for i := ind; i <= 9 && sum+i <= n; i++ {
		sum += i
		tmp = append(tmp, i)
		help216(res, tmp, k, n, i+1, sum) // 每个组合内不能使用重复数字
		sum -= i
		tmp = tmp[:len(tmp)-1]
	}
}
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	help216(&res, []int{}, k, n, 1, 0)
	for _, r := range res {
		fmt.Println(r)
	}
	return res
}

func main() {
	k, n := 3, 9
	fmt.Println(combinationSum3(k, n))
}
