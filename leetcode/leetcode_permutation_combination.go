package main

import "fmt"

/*
不存在重复数字的全排列
*/

// 1 dfs
func permutation(nums []int) [][]int {
	var res [][]int
	dfsnn(&res, nums, 0)
	return res
}

func dfsnn(res *[][]int, nums []int, index int) {
	if index >= len(nums) {
		*res = append(*res, append([]int{}, nums...))
	}

	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		dfsnn(res, nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

func testPerDfs() {
	nums := []int{1, 2, 3}
	fmt.Println(permutation(nums))
}

// 2

func permutation1(nums []int) [][]int {
	var res [][]int
	var out []int
	visited := make([]int, len(nums))
	permuteDfs1(&res, nums, out, visited, 0)
	return res
}

func permuteDfs1(res *[][]int, nums, out, visited []int, index int) {
	if index >= len(nums) {
		tmp := make([]int, len(out))
		copy(tmp, out)
		*res = append(*res, out)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		visited[i] = 1
		out = append(out, nums[i])
		permuteDfs1(res, nums, out, visited, index+1)
		out = out[:len(out)-1]
		visited[i] = 0
	}
}

func testPerDfs1() {
	nums := []int{3, 2, 1, 5}
	fmt.Println(permutation1(nums))
}

func Permute(nums []int) [][]int {
	//全排列
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var res [][]int
	for i, num := range nums {
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		sub := Permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}

func testPerDfs2() {
	nums := []int{3, 2, 1, 5}
	fmt.Println(Permute(nums))
}

// 3 非递归
func permutation3(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{}) //先生成一个空的数组

	for _, num := range nums {
		size1 := len(res)
		for i := 0; i < size1; i++ {
			front := res[0] // 取第一个
			res = res[1:]
			size := len(front)
			for j := 0; j <= size; j++ { // 在第一个数组的每一个位置尝试加入新的数字
				tmp := make([]int, len(front))
				copy(tmp, front)
				tmp = append(tmp[:j], append([]int{num}, tmp[j:]...)...)
				res = append(res, tmp)
			}
		}
	}
	return res
}

func testPerDfs3() {
	nums := []int{3, 2, 1, 5}
	fmt.Println(permutation3(nums))
}

/*
存在重复数字的全排列，产生无重复得到全排列
*/

func permuteUnique1(nums []int) [][]int {
	var res [][]int
	permuteUniqueDfs1(&res, nums, 0)
	return res
}

func permuteUniqueDfs1(res *[][]int, nums []int, index int) {
	if index >= len(nums) {
		*res = append(*res, append([]int{}, nums...))
	}
	set := make(map[int]int)
	for i := index; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			continue
		}
		set[nums[i]] = 1
		nums[i], nums[index] = nums[index], nums[i]
		permuteUniqueDfs1(res, nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

func testUnique1() {
	nums := []int{1, 2, 3}
	fmt.Println(permuteUnique1(nums))
}

func permuteUnique2(nums []int) [][]int {
	var res [][]int
	permuteUniqueDfs2(&res, nums, 0)
	return res
}

func permuteUniqueDfs2(res *[][]int, nums []int, index int) {
	if index >= len(nums) {
		*res = append(*res, append([]int{}, nums...))
	}
	for i := index; i < len(nums); i++ {
		if canSwap(nums, index, i) {
			nums[i], nums[index] = nums[index], nums[i]
			permuteUniqueDfs2(res, nums, index+1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
}

func canSwap(nums []int, start, end int) bool {
	for i := start; i < end; i++ {
		if nums[i] == nums[end] {
			return false
		}
	}
	return true
}

func testUnique2() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique2(nums))
}

func main() {
	testPerDfs()
	//testPerDfs1()
	//testPerDfs2()
	//testPerDfs3()

	testUnique1()
	testUnique2()
}
