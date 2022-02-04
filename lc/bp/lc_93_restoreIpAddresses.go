package main

import (
	"fmt"
	"strconv"
	"strings"
)
/*
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入'.' 来形成。
你不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

1、判断ip是否有效
2、有效则加入当前缓存ip中(tmp) 无效则continue

*/
func isOk(s string, i, j int) bool {
	tmp := s[i:j]
	if len(tmp) < 1 || len(tmp) > 3 {
		return false
	}
	if len(tmp) > 1 && tmp[0]-'0' == 0 {
		return false
	}
	num, err := strconv.Atoi(tmp)
	if err != nil || num < 0 || num > 255 {
		return false
	}
	return true
}

func help93(res *[]string, s string, tmp []string, ind int) {
	if ind >= len(s) {
		if len(tmp) == 4 {
			*res = append(*res, strings.Join(tmp, "."))
		}
		return
	}
	for i := ind; i < len(s); i++ {
		if !isOk(s, ind, i+1) {
			continue
		}
		tmp = append(tmp, s[ind:i+1])
		help93(res, s, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func restoreIpAddresses(s string) []string {
	var res []string
	help93(&res, s, []string{}, 0)
	return res
}

func main() {
	s := "101023"
	fmt.Println(restoreIpAddresses(s))
}
