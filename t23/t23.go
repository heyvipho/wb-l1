package main

import "fmt"

func removeElementByIndex(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

func main() {
	fmt.Println(removeElementByIndex([]int{1, 2, 3, 4}, 1))
}
