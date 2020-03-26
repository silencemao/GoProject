package main

import (
	"fmt"
	"sort"
)
/*

给定一个字符串数组，按照每一个字符串中的字母相同的生成一组。
*/
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	strMapList := make(map[string][]string)
	for _, str := range strs {
		strSort := sortByteInStr(str)
		if list, ok := strMapList[strSort]; ok {
			list = append(list, str)
			strMapList[strSort] = list
		} else {
			strMapList[strSort] = []string{str}
		}
	}
	for k := range strMapList {
		sort.Strings(strMapList[k])
		res = append(res, strMapList[k])
	}

	// 这部分排序纯属多余
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res) - 1 - i; j++ {
			if res[j][0] > res[j+1][0] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}

	return res
}

func sortByteInStr(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int)bool{
		return b[i] < b[j]
	})
	str := string(b)
	return str
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
