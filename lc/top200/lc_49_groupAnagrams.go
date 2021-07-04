package main

import (
	"fmt"
	"sort"
)

/*
给定一个字符串数组 将 组成字母一样的字符串 放在一起 返回
*/
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	if len(strs) < 1 {
		return res
	}
	f := func(str string) string {
		strByte := []byte(str)
		sort.Slice(strByte, func(i, j int) bool {
			return strByte[i] < strByte[j]
		})
		return string(strByte)
	}

	strMapList := make(map[string][]string, 0)
	strMapList[f(strs[0])] = []string{strs[0]}
	for i := 1; i < len(strs); i++ {
		strByte := f(strs[i])
		if list, ok := strMapList[strByte]; ok {
			list = append(list, strs[i])
			strMapList[strByte] = list
		} else {
			strMapList[strByte] = []string{strs[i]}
		}
	}
	for _, v := range strMapList {
		res = append(res, v)
	}
	return res
}
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
