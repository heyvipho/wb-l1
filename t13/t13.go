package main

import "fmt"

func main() {
	num1 := 1
	num2 := 2

	fmt.Println(num1, num2)

	num1, num2 = num2, num1

	fmt.Println(num1, num2)
}
