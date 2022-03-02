package main

import (
	"fmt"
	"sort"
)

/*
给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。

所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。

例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。

*/
// 这种方法不是字典序最小
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

type TicketCnt struct {
	ticket  string
	visited bool
}

type TicketCnts []*TicketCnt

func help3321(res *[]string, targets map[string]TicketCnts, ticketNum int) bool {
	if len(*res) == ticketNum+1 {
		return true
	}
	for _, ticket := range targets[(*res)[len(*res)-1]] {
		if ticket.visited {
			continue
		}
		*res = append(*res, ticket.ticket)
		ticket.visited = true

		if help3321(res, targets, ticketNum) {
			return true
		}
		ticket.visited = false
		*res = (*res)[:len(*res)-1]
	}
	return false
}

func findItinerary1(tickets [][]string) []string {
	var res []string
	targets := make(map[string]TicketCnts, 0)

	for _, ticket := range tickets {
		_, ok := targets[ticket[0]]
		if !ok {
			targets[ticket[0]] = make(TicketCnts, 0)
		}
		targets[ticket[0]] = append(targets[ticket[0]], &TicketCnt{visited: false, ticket: ticket[1]})
	}

	for ticket, tc := range targets {
		sort.Slice(tc, func(i, j int) bool {
			return tc[i].ticket < tc[j].ticket
		})
		targets[ticket] = tc
	}

	res = append(res, "JFK")
	help3321(&res, targets, len(tickets))
	return res
}

func main() {
	tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	fmt.Println(findItinerary(tickets))
	fmt.Println(findItinerary1(tickets))
}
