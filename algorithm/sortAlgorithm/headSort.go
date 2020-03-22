package main

import "fmt"

func heapAdjust(arr []int, beg int, length int) {
	left := 2 * beg + 1
	for left < length {
		if left+1 < length && arr[left] < arr[left+1] {
			left++
		}
		if arr[beg] < arr[left] {
			arr[beg], arr[left] = arr[left], arr[beg]
			beg = left
			left = 2 * left + 1
		} else {
			break
		}
	}
}

func heapSort(arr []int) {
	for i := len(arr) /2 - 1; i >=0 ; i-- {
		heapAdjust(arr[:], i, len(arr) - 1)
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapAdjust(arr, 0, i)
	}
}

func main() {
	arr := []int{4, 3, 5, 1, 6, 2, 0, 10, 1}
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println(arr)

}
