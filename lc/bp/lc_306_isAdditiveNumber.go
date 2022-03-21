package main

import "fmt"

/*
累加数
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

func addStr(str1, str2 string) string {
	if len(str1) < len(str2) {
		return addStr(str2, str1)
	}

	l1, l2 := len(str1), len(str2)
	for l1 > l2 {
		str2 = "0" + str2
		l2++
	}
	res := ""
	carry := 0
	for i := l2 - 1; i >= 0; i-- {
		m := int(str1[i]-'0'+str2[i]-'0') + carry
		res = res + string(rune(m%10))
		carry = m / 10
	}
	return res
}

func isValid306(num string, i, j, k int) bool {
	if num[i] == '0' && j != i+1 {
		return false
	}
	if num[j] == '0' && k != j+1 {
		return false
	}
	sum := addStr(num[i:j], num[j:k])

	if len(sum) > len(num)-k {
		return false
	}

	tmp := num[k : k+len(sum)]
	for m := 0; m < len(tmp); m++ {
		if sum[m] != tmp[m] {
			return false
		}
	}
	if len(sum)+k == len(num) {
		return true
	}

	return isValid306(num, j, k, k+len(sum))
}

func isAdditiveNumber2(num string) bool {
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if isValid306(num, i, j, j+1) {
				return true
			}
		}
	}
	return false
}

func main() {
	num := "199100199"
	fmt.Println(isAdditiveNumber(num))
	fmt.Println(isAdditiveNumber2(num))
}
