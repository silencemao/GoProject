package main

import "fmt"

/*
全排列
*/
func help46(res *[][]int, nums []int, ind int) {
	fmt.Println(nums)
	if ind >= len(nums) {
		*res = append(*res, append([]int{}, nums...))
	}
	for i := ind; i < len(nums); i++ {
		nums[i], nums[ind] = nums[ind], nums[i]
		help46(res, nums, ind+1) //
		nums[i], nums[ind] = nums[ind], nums[i]
	}
}

//排列写法，每次遍历都从0开始遍历，used数组记录元素是否使用，每个元素在一个排列里面只能使用一次
func help46II(res *[][]int, nums, tmp, used []int) {
	if len(nums) == len(tmp) {
		*res = append(*res, append([]int{}, tmp...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] == 1 {
			continue
		}
		used[i] = 1
		tmp = append(tmp, nums[i])
		help46II(res, nums, tmp, used)
		tmp = tmp[:len(tmp)-1]
		used[i] = 0
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	//help46(&res, nums, 0)

	used := make([]int, len(nums))
	help46II(&res, nums, []int{}, used)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
