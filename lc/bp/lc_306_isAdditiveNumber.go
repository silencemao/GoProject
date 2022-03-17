package main

import "fmt"

/*
雷加数
https://leetcode-cn.com/problems/additive-number/solution/tong-ge-lai-shua-ti-la-dfs-jian-zhi-by-t-jxsf/
*/
func help306(num string, ind int, cnt *int, preprev, prev int64) bool {
	//fmt.Println(preprev, prev, cnt)
	if ind >= len(num) {
		return *cnt >= 3
	}
	var cur int64 = 0
	for i := ind; i < len(num); i++ {
		if num[ind] == '0' && i > ind {
			return false
		}
		cur = cur*10 + int64(byte(num[i])-'0')
		if *cnt >= 2 {
			if cur > preprev+prev {
				return false
			}
			if cur < preprev+prev {
				continue
			}
		}

		*cnt += 1
		if help306(num, i+1, cnt, prev, cur) {
			return true
		}
		*cnt -= 1
		//if help306(num, i+1, cnt+1, prev, cur) {
		//	return true
		//}
	}
	return false
}

func isAdditiveNumber(num string) bool {
	cnt := 0
	res := help306(num, 0, &cnt, 0, 0)
	return res
}

func main() {
	num := "199100199"
	fmt.Println(isAdditiveNumber(num))
}
