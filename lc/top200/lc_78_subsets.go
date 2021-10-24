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

func help78(res *[][]int, tmp, nums []int, ind int) {
	*res = append(*res, append([]int{}, tmp...)) // 每个组合都保留
	if ind >= len(nums) {
		return
	}
	for i := ind; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		help78(res, tmp, nums, i+1) // 从下一个位置取数，不重复
		tmp = tmp[:len(tmp)-1]
	}
}

func subsets2(nums []int) [][]int {
	var res [][]int
	help78(&res, []int{}, nums, 0)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets2(nums))
}
