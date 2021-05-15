package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for _, num := range nums {
		for _, arr := range res {
			//arr = append(arr, num) //这种写法会改变res中已有arr的值
			tmp := append([]int{num}, arr...)
			res = append(res, tmp)
		}
	}
	return res
}

func subsets1(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(subsets1(nums))
}
