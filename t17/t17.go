package main

import "fmt"

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	left := make([]int, 0)
	right := make([]int, 0)
	pivot := a[0]

	for _, v := range a[1:] {
		if v < pivot {
			left = append(left, v)
			continue
		}
		right = append(right, v)
	}

	return append(quicksort(left), append([]int{pivot}, quicksort(right)...)...)
}

func binSearch(a []int, v int) int {
	index := len(a) / 2

	if a[index] == v {
		return index
	}

	if index == 0 {
		return -1
	}

	if v > a[index] {
		return binSearch(a[index:], v) + index
	} else {
		return binSearch(a[:index], v)
	}
}

func main() {
	slice := []int{3, 2, -1, 50, -100, 45, 22}

	sortedSlice := quicksort(slice)

	fmt.Println(sortedSlice)

	fmt.Printf("value: %v, key: %v\n", -100, binSearch(sortedSlice, -100))
	fmt.Printf("value: %v, key: %v\n", -1, binSearch(sortedSlice, -1))
	fmt.Printf("value: %v, key: %v\n", 2, binSearch(sortedSlice, 2))
	fmt.Printf("value: %v, key: %v\n", 3, binSearch(sortedSlice, 3))
	fmt.Printf("value: %v, key: %v\n", 22, binSearch(sortedSlice, 22))
	fmt.Printf("value: %v, key: %v\n", 45, binSearch(sortedSlice, 45))
	fmt.Printf("value: %v, key: %v\n", 50, binSearch(sortedSlice, 50))
}
