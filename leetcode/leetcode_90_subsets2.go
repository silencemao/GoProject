package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}

	size, j := len(res), 0
	for i := range nums {
		j = 0
		if i > 0 && nums[i] == nums[i-1] {
			j = size
		}
		size = len(res)
		for ; j < size; j++ {
			// 下面两种写法会得到不同的结果
			res = append(res, res[j])
			res[len(res)-1] = append(res[len(res)-1], nums[i])

			//res = append(res, append([]int{nums[i]}, res[j]...))
		}
	}
	return res
}

func subsetsWithDup1(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	help90(&res, nums, 0, []int{})
	return res
}

func help90(res *[][]int, nums []int, ind int, tmp []int) {
	tmpCur := make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		tmpCur[i] = tmp[i]
	}
	*res = append(*res, tmpCur)

	for i := ind; i < len(nums); i++ {
		if i > ind && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		help90(res, nums, i+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}

func main() {
	nums := []int{2, 1, 2, 1, 3}
	fmt.Println(subsetsWithDup(nums))
	fmt.Println(subsetsWithDup1(nums))
}
