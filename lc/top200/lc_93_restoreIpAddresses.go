package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isIp(s string, i, j int) bool {
	if j-i > 1 && s[i] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s[i:j])
	if num > 255 {
		return false
	}
	return true
}

func help93(res *[]string, tmp []string, s string, ind int) {
	if ind < len(s) && len(tmp) > 4 {
		return
	}
	if ind >= len(s) {
		if len(tmp) == 4 {
			*res = append(*res, strings.Join(tmp, "."))
		}
		return
	}
	for i := ind; i < len(s); i++ {
		if isIp(s, ind, i+1) {
			tmp = append(tmp, s[ind:i+1])
			help93(res, tmp, s, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func restoreIpAddresses(s string) []string {
	var res []string
	if len(s) > 12 {
		return res
	}
	help93(&res, []string{}, s, 0)
	for _, r := range res {
		fmt.Println(r)
	}
	return res
}
func main() {
	s := "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	fmt.Println(restoreIpAddresses(s))
}
