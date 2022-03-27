package main

import "fmt"

/*
 拆分字符串使唯一子字符串的数目最大
给你一个字符串 s ，请你拆分该字符串，并返回拆分后唯一子字符串的最大数目。
字符串 s 拆分后可以得到若干 非空子字符串 ，这些子字符串连接后应当能够还原为原字符串。但是拆分出来的每个子字符串都必须是 唯一的 。
注意：子字符串 是字符串中的一个连续字符序列。

给定字符串s，拆分为多个子字符串，每个字符串唯一，s最多可以拆成多少个子字符串
*/
func check1593(tmp []string) bool {
	for i := 0; i < len(tmp); i++ {
		for j := i + 1; j < len(tmp); j++ {
			if tmp[i] == tmp[j] {
				return false
			}
		}
	}
	return true
}

func help1593(s string, ind int, cnt *int, res *[][]string, set map[string]int, tmp []string) {
	if ind >= len(s) {
		//fmt.Println(tmp)
		*res = append(*res, append([]string{}, tmp...))
		if len(tmp) > *cnt {
			*cnt = len(tmp)
		}
	}

	for i := ind; i < len(s); i++ {
		if _, ok := set[s[ind:i+1]]; ok {
			continue
		}

		if len(s[ind:i+1]) > 0 {
			tmp = append(tmp, s[ind:i+1])
			set[s[ind:i+1]] = 1
			help1593(s, i+1, cnt, res, set, tmp)
			delete(set, tmp[len(tmp)-1])
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func maxUniqueSplit(s string) int {
	var cnt int
	var res [][]string
	set := make(map[string]int, 0)
	help1593(s, 0, &cnt, &res, set, []string{})
	return cnt
}

func main() {
	s := "ababccc"
	fmt.Println(maxUniqueSplit(s))
}
