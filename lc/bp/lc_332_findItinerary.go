package main

import "fmt"

func help332(res *[]string, targets map[string]map[string]int, ticketNum int) bool {
	if len(*res) == ticketNum+1 {
		return true
	}
	for end, cnt := range targets[(*res)[len(*res)-1]] {
		if cnt > 0 {
			//cnt -= 1
			targets[(*res)[len(*res)-1]][end] = cnt - 1
			*res = append(*res, end)

			if help332(res, targets, ticketNum) {
				return true
			}

			*res = (*res)[:len(*res)-1]
			//cnt += 1
			targets[(*res)[len(*res)-1]][end] = cnt + 1
		}
	}
	return false
}

func findItinerary(tickets [][]string) []string {
	targets := make(map[string]map[string]int, 0)
	for _, ticket := range tickets {
		if tmp, ok := targets[ticket[0]]; ok {
			tmp[ticket[1]] += 1
			targets[ticket[0]] = tmp
		} else {
			targets[ticket[0]] = make(map[string]int, 0)
			targets[ticket[0]][ticket[1]] += 1
		}
	}

	var res []string
	res = append(res, "JFK")
	help332(&res, targets, len(tickets))
	return res
}

func main() {
	tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	fmt.Println(findItinerary(tickets))
}
