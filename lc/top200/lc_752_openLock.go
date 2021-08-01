package main

import "fmt"

func plusOne1(s string, j int) string {
	bs := []byte(s)
	if bs[j] == '9' {
		bs[j] = '0'
	} else {
		bs[j] += 1
	}
	return string(bs)
}

func minusOne1(s string, j int) string {
	bs := []byte(s)
	if bs[j] == '0' {
		bs[j] = '9'
	} else {
		bs[j] -= 1
	}
	return string(bs)
}

func openLock(deadends []string, target string) int {
	var queue []string
	queue = append(queue, "0000")
	visited := make(map[string]bool, 0)
	visited["0000"] = true
	step := 0

	deadMap := make(map[string]bool)
	for _, d := range deadends {
		deadMap[d] = true
	}

	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]

			if deadMap[cur] {
				continue
			}
			if cur == target {
				return step
			}

			for i := 0; i < len(target); i++ {
				up := plusOne1(cur, i)
				down := minusOne1(cur, i)
				if !visited[up] {
					queue = append(queue, up)
					visited[up] = true
				}
				if !visited[down] {
					queue = append(queue, down)
					visited[down] = true
				}

			}
		}
		step++
	}

	return -1
}
func main() {
	deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}
