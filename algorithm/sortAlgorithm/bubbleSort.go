package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func main() {
	arr := []int{4, 3, 5, 1, 6, 2, 0, 9}
	fmt.Println(arr)
	fmt.Println(bubbleSort(arr))
}
