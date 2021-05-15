package main

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

/*
员工的重要性
*/
func getImportance(employees []*Employee, id int) int {
	res := 0

	var queue []int
	for _, e := range employees {
		if e.Id == id {
			res += e.Importance
			queue = append(queue, e.Subordinates...)
			break
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, e := range employees {
			if e.Id == front {
				res += e.Importance
				queue = append(queue, e.Subordinates...)
				break
			}
		}
	}
	return res
}

// [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]],
func main() {
	e1 := &Employee{Id: 1, Importance: 5, Subordinates: []int{2, 3}}
	e2 := &Employee{Id: 2, Importance: 3, Subordinates: []int{}}
	e3 := &Employee{Id: 3, Importance: 3, Subordinates: []int{}}
	employees := []*Employee{e1, e2, e3}
	fmt.Println(getImportance(employees, 1))
}
