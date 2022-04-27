package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	arr := strings.Split(str, " ")

	reversedArr := make([]string, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		reversedArr = append(reversedArr, arr[i])
	}

	return strings.Join(reversedArr, " ")
}

func main() {
	fmt.Println("snow dog sun")
	fmt.Println(reverseString("snow dog sun"))
}
