package main

import "fmt"

func isPalindrome131(s string, i, j int) bool {
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func help131(s string, res *[][]string, tmp []string, ind int) {
	if ind == len(s) {
		*res = append(*res, append([]string{}, tmp...))
		return
	}
	for i := ind; i < len(s); i++ {
		if !isPalindrome131(s, ind, i) {
			continue
		}
		tmp = append(tmp, s[ind:i+1])
		help131(s, res, tmp, i+1) //此处是 i+1 不是 ind+1
		tmp = tmp[:len(tmp)-1]
	}
}

func partition131(s string) [][]string {
	var res [][]string
	help131(s, &res, []string{}, 0)
	return res
}

/*
先用二维数组记录任意子串是否为回文串，这样就不用每次都判断了
*/

func help131II(s string, res *[][]string, tmp []string, ind int, dp [][]bool) {
	if ind == len(s) {
		*res = append(*res, append([]string{}, tmp...))
		return
	}
	for i := ind; i < len(s); i++ {
		if !dp[ind][i] {
			continue
		}
		tmp = append(tmp, s[ind:i+1])
		help131II(s, res, tmp, i+1, dp)
		tmp = tmp[:len(tmp)-1]
	}
}
func partition131II(s string) [][]string {
	var res [][]string
	size := len(s)

	dp := make([][]bool, size)
	for r := 0; r < size; r++ {
		dp[r] = make([]bool, size)
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}
		}
	}

	help131II(s, &res, []string{}, 0, dp)
	return res
}

func main() {
	s := "abaamnm"
	fmt.Println(partition131(s))
	fmt.Println(partition131II(s))
}
