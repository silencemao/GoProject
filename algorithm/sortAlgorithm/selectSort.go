package main

import "fmt"

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minInd := i

		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[minInd] {
				minInd = j
			}
		}
		arr[i], arr[minInd] = arr[minInd], arr[i]
	}

	return arr
}

func main() {
	arr := []int{4, 3, 5, 1, 6, 2, 0, 10, 1}
	fmt.Println(arr)
	fmt.Println(selectSort(arr))
}
