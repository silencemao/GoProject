package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if s[i] == ')' {
				if len(stack) < 1 || stack[len(stack)-1] != '(' {
					return false
				}
			} else if s[i] == '}' {
				if len(stack) < 1 || stack[len(stack)-1] != '{' {
					return false
				}
			} else if s[i] == ']' {
				if len(stack) < 1 || stack[len(stack)-1] != '[' {
					return false
				}
			} else {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValid1(s string) bool {
	pair := make(map[byte]byte, 0)
	pair[')'] = '('
	pair['}'] = '{'
	pair[']'] = '['
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if _, ok := pair[s[i]]; !ok {
				return false
			}
			if len(stack) < 1 || stack[len(stack)-1] != pair[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	s := "()(){[]}}"
	fmt.Println(isValid(s))
}
