package main

func twoSum(nums []int, target int) []int {
	var res []int
	var numMap = make(map[int]int, 0)
	numMap[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		sub := target - nums[i]
		if index, ok := numMap[sub]; ok {
			res = append(res, index)
			res = append(res, i)
			return res
		} else {
			numMap[nums[i]] = i
		}
	}
	return res
}

//func main() {
//	var nums = []int{2, 7, 11, 15}
//	target := 17
//	fmt.Println(twoSum(nums, target))
//}
