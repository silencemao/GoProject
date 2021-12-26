package main

import "fmt"

func help5864(n int, direct map[byte][]int, s byte, curPos []int) ([]int, bool) {
	x, y := direct[s][0], direct[s][1]
	nexX, nexY := curPos[0]+x, curPos[1]+y
	if nexX < 0 || nexX >= n || nexY < 0 || nexY >= n {
		return []int{-1, -1}, false
	}
	return []int{nexX, nexY}, true
}

func executeInstructions(n int, startPos []int, s string) []int {
	direct := make(map[byte][]int)
	direct['L'] = []int{0, -1}
	direct['R'] = []int{0, 1}
	direct['U'] = []int{-1, 0}
	direct['D'] = []int{1, 0}

	answer := make([]int, len(s))
	if n == 1 {
		return answer
	}

	curPos := make([]int, 2)
	for i := 0; i < len(s); i++ {
		curPos[0], curPos[1] = startPos[0], startPos[1]
		for j := i; j < len(s); j++ {
			isOK := false
			if curPos, isOK = help5864(n, direct, s[j], curPos); isOK {
				answer[i] += 1
			} else {
				break
			}
		}
	}
	return answer
}
func main() {
	n := 1
	startPos := []int{0, 0}
	s := "LURD"
	fmt.Println(executeInstructions(n, startPos, s))
}
