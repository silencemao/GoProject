package main

import "fmt"

/*
给你一个整数数组 nums ，除某个元素仅出现 一次外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
问题很清晰，找出只出现一次的数字。
1、将所有数字按位相加
2、因为除了结果之外其它数字都出现了三次，则每一位相加之后，是3n+1 或 3n+0
3、3n表示出现了三次的数字在某一位的加和，1 或 0是结果在某一位的值
4、我们需要的就是拿到结果各个位的值，然后组成一个值返回。此处取余即可拿到该位的值
5、如何取每一个位置的值呢？我们可以将nums[i]右移m位，然后与1 进行 与运算，即可拿到nums[i]第m位的值
6、如何保存每一个位置的值呢，我们定义一个变量res，或运算即可将指定位置的值加到res中
*/
func singleNumber137(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		sum := int32(0)
		for j := 0; j < len(nums); j++ {
			sum += (int32(nums[j]) >> i) & 1
		}
		sum %= 3
		res |= sum << i
	}
	return int(res)
}
func main() {
	nums := []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}
	fmt.Println(singleNumber137(nums))
}
