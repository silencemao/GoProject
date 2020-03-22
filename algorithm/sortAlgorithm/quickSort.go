package main

import "fmt"

func quickSort(arr []int, left int, right int) []int {
	if left < right {
		key := arr[left]
		low, high := left + 1, right

		for low <= high{
			for low < high && arr[low] > key && arr[high] < key {
				arr[low], arr[high] = arr[high], arr[low]
				low++
				high--
			}
			if arr[low] <= key {
				low++
			}
			if arr[high] >= key {
				high--
			}
		}
		arr[left], arr[high] = arr[high], arr[left]
		quickSort(arr, left, high-1)
		quickSort(arr, high+1, right)
	}

	return arr
}

func quickSort1(arr []int, left int, right int) []int {

	if left < right {
		key := arr[left]
		low, high := left, right
		for low < high {
			for low < high && arr[high] >= key {
				high--
			}
			if low < high {
				arr[low] = arr[high]
				low++
			}

			for low < high && arr[low] <= key {
				low++
			}
			if low < high {
				arr[high] = arr[low]
			}
		}

		arr[low] = key
		quickSort(arr, left, low-1)
		quickSort(arr, low+1, right)
	}

	return arr
}

func main() {
	arr := []int{4, 3, 5, 1, 6, 2, 0, 10, 1}
	fmt.Println(arr)
	fmt.Println(quickSort(arr, 0, len(arr)-1))
	fmt.Println(quickSort1(arr, 0, len(arr)-1))
}
