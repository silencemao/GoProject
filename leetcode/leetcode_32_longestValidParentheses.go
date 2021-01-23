package main

import "fmt"

// https://leetcode.com/problems/longest-valid-parentheses/discuss/14355/My-solution-using-one-stack-in-one-pass

/*
最长的有效括号对，给定字符串，返回最长有效括号对的长度 有效 即括号是闭合的 ()(()())()
利用堆栈的思路，有点像1003的思路
从左至右遍历字符串，遇到'('则将其索引压入栈中，
遇到')'分情况，
此时若栈顶是'('则将栈顶出栈，相当于i与栈顶位置的元素可以组成有效的()，那应该如何计算长度呢，
		我下面写了两种方式：
        第一种方式是错误，我们直接找得到当前栈顶的索引front i-front+1，这样计算是小于最终的最大长度的，因为若j-2与j-1也配对组成了括号，而j-2
        已经弹出栈了，所以我们计算的长度小于真正的最大长度的
        第二种方式是正确的，把栈顶索引弹出之后，此时若栈不为空，则直接i-新的栈顶索引即可
                                             若为空，直接i+1即可
若栈为空 或栈顶不是'('，相当于栈内剩余的索引没有可以组成()的，直接将当前i入栈即可
*/
func longestValidParentheses(s string) int {
	res := 0
	stack := make([]int, 0)
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 && ss[stack[len(stack)-1]] == '(' {
				//front := stack[len(stack)-1]
				//stack = stack[:len(stack)-1]
				//if i-front+1 > res {
				//	res = i - front + 1
				//}
				front := -1
				if len(stack) > 0 {
					front = stack[len(stack)-1]
				}
				if i-front+1 > res {
					res = i - front + 1
				}
			} else {
				stack = append(stack, i)
			}
		}
	}
	return res
}

func main() {
	s := "()))()(())"
	fmt.Println(longestValidParentheses(s))
}
