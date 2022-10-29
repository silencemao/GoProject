package main

import "fmt"

/*
字符串解码
*/
func decodeString(s string) string {
	var nums []int
	num := 0
	res := ""
	var stack []string

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] >= 'a' && s[i] <= 'z' {
			res += string(s[i])
		} else if s[i] == '[' {
			nums = append(nums, num)
			num = 0
			stack = append(stack, res)
			res = ""
		} else {
			times := nums[len(nums)-1]
			for k := 0; k < times; k++ {
				stack[len(stack)-1] += res
			}
			res = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nums = nums[:len(nums)-1]
		}
	}

	return res
}

func main() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}
