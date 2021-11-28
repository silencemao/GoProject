package main

import (
	"fmt"
	"strconv"
)

/*
 逆波兰表达式， 是一种后缀表达式，所谓后缀就是指算符写在后面。
给定一个字符串数组，包含数字和+ - * / 符号，
{"2", "1", "+", "3", "*"} 表示 （2+1）*3

是栈的一种应用，遇到数字就放到栈中，遇到符号就把栈顶的两个数字拿出来，进行计算，然后再将结果压入栈顶
*/
func evalRPN(tokens []string) int {
	res := 0

	stack := make([]int, 0)
	set := make(map[string]int)
	set["+"] = 1
	set["-"] = 2
	set["*"] = 3
	set["/"] = 4
	for _, str := range tokens {
		if method, ok := set[str]; ok {
			if len(stack) < 2 {
				panic("param num is error")
			}
			front2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			front1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			tmp := 0
			if method == 1 {
				tmp = front1 + front2
			} else if method == 2 {
				tmp = front1 - front2
			} else if method == 3 {
				tmp = front1 * front2
			} else {
				tmp = front1 / front2
			}
			stack = append(stack, tmp)
		} else {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic("str is not num or signal")
			}
			stack = append(stack, num)
		}
	}
	res = stack[len(stack)-1]

	return res
}

func evalRPN1(tokens []string) int {
	stack := make([]int, 0)
	for _, str := range tokens {
		num, err := strconv.Atoi(str)
		if err == nil {
			stack = append(stack, num)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch str {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[len(stack)-1]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN1(tokens))
}
