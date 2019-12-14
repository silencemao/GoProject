package main

/*
给定数字n，生成n对括号
*/
func generateParenthesis(n int) []string {
	var result []string     // 注意slice的传递方式
	helper("", 0, 0, n, &result)
	return result
}

func helper(str string, left, right, max int, result *[]string) {
	if len(str)==2 * max {
		*result = append(*result, str)
		return
	}

	if left < max {
		helper(str + "(", left+1, right, max, result)
	}

	if right < left {
		helper(str + ")", left, right+1, max, result)
	}
}

//func main() {
//	result := generateParenthesis(3)
//	for _, str := range result {
//		fmt.Println(str)
//	}
//
//}
