package main

import "fmt"

/*
模糊坐标
我们有一些二维坐标，如"(1, 3)"或"(2, 0.5)"，然后我们移除所有逗号，小数点和空格，得到一个字符串S。返回所有可能的原始字符串到一个列表中。
原始的坐标表示法不会存在多余的零，所以不会出现类似于"00", "0.0", "0.00", "1.0", "001", "00.01"或一些其他更小的数来表示坐标。
此外，一个小数点前至少存在一个数，所以也不会出现“.1”形式的数字。
最后返回的列表可以是任意顺序的。而且注意返回的两个数字中间（逗号之后）都有一个空格。
示例 1:
输入: "(123)"
输出: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]

*/
func getPos(s string) []string {
	var res []string
	if s[0] != '0' || s == "0" {
		res = append(res, s)
	}
	for p := 1; p < len(s); p++ {
		if p != 1 && s[0] == '0' || s[len(s)-1] == '0' {
			continue
		}
		res = append(res, s[:p]+"."+s[p:])
	}
	return res
}

func ambiguousCoordinates(s string) []string {
	var res []string
	n := len(s) - 2
	s = s[1 : len(s)-1]
	for i := 1; i < n; i++ {
		lt := getPos(s[:i])
		if len(lt) == 0 {
			continue
		}
		rt := getPos(s[i:])
		if len(rt) == 0 {
			continue
		}

		for _, t1 := range lt {
			for _, t2 := range rt {
				res = append(res, "("+t1+", "+t2+")")
			}
		}
	}
	return res
}

func main() {
	s := "(123)"
	fmt.Println(ambiguousCoordinates(s))
}
