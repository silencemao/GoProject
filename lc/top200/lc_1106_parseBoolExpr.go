package main

import "fmt"

func parseBoolExpr(expression string) bool {
	stk := []rune{}
	for _, c := range expression {
		if c != '(' && c != ')' && c != ',' {
			stk = append(stk, c)
		} else if c == ')' {
			var f, t int
			for len(stk) > 0 && (stk[len(stk)-1] == 'f' || stk[len(stk)-1] == 't') {
				if stk[len(stk)-1] == 'f' {
					f += 1
				} else {
					t += 1
				}
				stk = stk[:len(stk)-1]
			}
			op := stk[len(stk)-1]
			stk = stk[:len(stk)-1]

			c = 'f'
			if (op == '!' && f > 0) || (op == '|' && t > 0) || (op == '&' && f == 0) {
				c = 't'
			}
			stk = append(stk, c)
		}
	}
	return stk[0] == 't'
}

func main() {
	expression := "!(&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))))"
	fmt.Println(parseBoolExpr(expression))
}
