package main

import "fmt"

/*
给定一个字符串s="PAYPALISHIRING"，s中字符顺序按照Z字形排列，如下图，返回按行遍历下图之后的结果"PAHNAPLSIIGYIR"

P   A   H   N
A P L S I I G
Y   I   R

1、每一行存储到一个字符串中，最后合并多个字符串
2、从0开始遍历s，确定每一个字符属于的行，加入到对应行中的字符串中
3、row、step控制行的变化，step=1从上往下遍历，step=-1从下网上遍历
*/
func convert(s string, numRows int) string {
	if numRows < 2 || numRows > len(s) {
		return s
	}
	res := ""
	var strs []string
	for i := 0; i < numRows; i++ {
		strs = append(strs, "")
	}
	row, step := 0, 1
	for i := 0; i < len(s); i++ {
		if row == numRows-1 { // 当row从上遍历到numRows-1行时，下一步row要从下往上遍历
			step = -1
		}
		if row == 0 { // row从下网上遍历到第0行时，下一步要从上往下遍历
			step = 1
		}
		strs[row] += string(s[i])
		row += step
	}
	for _, str := range strs {
		res += str
	}
	return res
}

func convert1(s string, numRows int) string {
	if numRows < 2 || numRows > len(s) {
		return s
	}
	res := ""
	var strs []string
	for i := 0; i < numRows; i++ {
		strs = append(strs, "")
	}

	row, goDown := 0, false
	for _, c := range s {
		strs[row] += string(c)
		if row == 0 || row == numRows-1 {
			goDown = !goDown
		}
		if goDown {
			row += 1
		} else {
			row -= 1
		}
	}

	for _, str := range strs {
		res += str
	}
	return res
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert1(s, numRows))
}
