package main

import (
	"fmt"
	"strconv"
)

/*
6038. 向表达式添加括号后的最小结果
给你一个下标从 0 开始的字符串 expression ，格式为 "<num1>+<num2>" ，其中 <num1> 和 <num2> 表示正整数。

请你向 expression 中添加一对括号，使得在添加之后， expression 仍然是一个有效的数学表达式，并且计算后可以得到 最小 可能值。左括号 必须 添加在 '+' 的左侧，而右括号必须添加在 '+' 的右侧。

返回添加一对括号后形成的表达式 expression ，且满足 expression 计算得到 最小 可能值。如果存在多个答案都能产生相同结果，返回任意一个答案。

生成的输入满足：expression 的原始值和添加满足要求的任一对括号之后 expression 的值，都符合 32-bit 带符号整数范围。

输入：expression = "247+38"
输出："2(47+38)"
解释：表达式计算得到 2 * (47 + 38) = 2 * 85 = 170 。
注意 "2(4)7+38" 不是有效的结果，因为右括号必须添加在 '+' 的右侧。
可以证明 170 是最小可能值
*/
func help6038(num string, i int) (int, int) {
	l, _ := strconv.Atoi(num[:i])
	r, _ := strconv.Atoi(num[i:])
	return l, r
}

func minimizeResult(expression string) string {
	i := 0
	expre1 := expression
	for ; i < len(expression); i++ {
		if expression[i] == '+' {
			break
		}
	}
	num1, num2 := expression[:i], expression[i+1:]
	l1, l2 := len(num1), len(num2)
	res := 1<<31 - 1
	for i := 0; i <= l1-1; i++ {
		mul1, add1 := help6038(num1, i)
		if mul1 == 0 {
			mul1 = 1
		}
		add2, mul2 := 0, 1
		j := 0
		for ; j <= l2; j++ {
			add2, mul2 = help6038(num2, j)
			if mul2 == 0 {
				mul2 = 1
			}

			tmp := mul1 * mul2 * (add1 + add2)
			if tmp < res {
				fmt.Println(mul1, add1, add2, mul2)
				//fmt.Println(i, j)
				res = tmp
				if i == 0 {
					expre1 = "(" + num1[:i] + num1[i:] + "+"
				} else if i != 0 {
					expre1 = num1[:i] + "(" + num1[i:] + "+"
				}

				if j == 0 {
					expre1 = expre1 + num2[:j] + num2[j:] + ")"
				} else if j != 0 {
					expre1 = expre1 + num2[:j] + ")" + num2[j:]
				}

			}
		}

	}
	return expre1
}

func main() {
	expression := "1+1"
	fmt.Println(minimizeResult(expression))
}
