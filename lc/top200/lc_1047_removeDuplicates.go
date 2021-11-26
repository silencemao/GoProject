package main

import (
	"fmt"
)

/*
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
在 S 上反复执行重复项删除操作，直到无法继续删除。
在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，
其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
*/
func removeDuplicates1047(s string) string { // 结果对，方法不好
	var stack []byte
	bs := []byte(s)
	for i := range bs {
		if len(stack) == 0 || stack[len(stack)-1] != bs[i] {
			stack = append(stack, bs[i])
		} else {
			isOK := false // 针对字母全部相同的情况，全部相同就不会进入下面的else部分，不会删除栈顶的字母
			for j := i; j < len(bs); {
				if bs[j] == bs[i] {
					j += 1
				} else {
					i = j
					stack = stack[:len(stack)-1]
					isOK = true
					break
				}
			}
			if !isOK {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(stack)
}

func removeDuplicates2(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		// 栈不空 且 与栈顶元素不等
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			// 弹出栈顶元素 并 忽略当前元素(s[i])
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func main() {
	s := "aaa"
	fmt.Println(removeDuplicates2(s))

	fmt.Println(removeDuplicates1047(s))
}
