package main

import "fmt"

func merge88(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}

	k := m + n - 1
	m, n = m-1, n-1
	for m >= 0 && n >= 0 {
		if nums1[m] >= nums2[n] {
			nums1[k] = nums1[m]
			m--
		} else {
			nums1[k] = nums2[n]
			n--
		}
		k--
	}
	if n > m {
		for n >= 0 {
			nums1[k] = nums2[n]
			k, n = k-1, n-1
		}
	}
	if m > n {
		for m >= 0 {
			nums1[k] = nums1[m]
			k, m = k-1, m-1
		}
	}
	fmt.Println(nums1)
}

func merge881(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	m, n = m-1, n-1
	cur := 0
	for m >= 0 || n >= 0 {
		if m == -1 {
			cur = nums2[n]
			n--
		} else if n == -1 {
			cur = nums1[m]
			m--
		} else if nums1[m] >= nums2[n] {
			cur = nums1[m]
			m--
		} else {
			cur = nums2[n]
			n--
		}
		nums1[k] = cur
		k--
	}
	fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, len(nums2)
	//merge88(nums1, m, nums2, n)
	merge881(nums1, m, nums2, n)
}
