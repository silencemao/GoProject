package main

/*
给定字符串S,找出所有子串中只出现过一次的字符个数，子串哥重复
S="ABC"
子串有 A AB B BC ABC
      1  2 1  2  3   =  10
采用动态规划的方式，
*/
func uniqueLetterString(S string) int {
	res, cur := 0, 0
	var M int64 = 1e9 + 7
	var first [26]int
	var second [26]int
	for i := 0; i < len(S); i++ {
		c := S[i] - 'A'
		cur = cur + 1 + i - first[c] * 2 + second[c]
		res = int(int64(res + cur) % M)
		second[c] = first[c]
		first[c] = i + 1
	}
	return res
}

//func main() {
//	S := "ABA"
//	fmt.Println(uniqueLetterString(S))
//}
