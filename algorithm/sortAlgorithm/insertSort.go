package main

import "fmt"

func insertSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		arrI := arr[i]

		j := i - 1
		for j >= 0 && arr[j] > arrI {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = arrI
	}

	return arr
}

func main() {
	arr := []int{4, 3, 5, 1, 6, 2, 0, 10}
	fmt.Println(arr)
	fmt.Println(insertSort(arr))
}
