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

func main() {
	slice := []int{3, 2, -1, 50, -100, 45, 22}

	fmt.Println(quicksort(slice))
}
