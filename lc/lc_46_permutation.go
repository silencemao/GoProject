package main

func dfs(res *[][]int, nums *[]int, index int) {
	if index >= len(*nums) {
		var tmp []int
		for _, num := range *nums {
			tmp = append(tmp, num)
		}
		*res = append(*res, tmp)
		return
	}
	for i := index; i < len(*nums); i++ {
		(*nums)[index], (*nums)[i] = (*nums)[i], (*nums)[index]
		dfs(res, nums, index+1)
		(*nums)[index], (*nums)[i] = (*nums)[i], (*nums)[index]
	}
}
func permute(nums []int) [][]int {
	var res [][]int
	dfs(&res, &nums, 0)
	return res
}

//func main() {
//	var nums = []int{1, 2, 3}
//	fmt.Println(permute(nums))
//}