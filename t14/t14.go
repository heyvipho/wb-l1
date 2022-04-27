package main

import "fmt"

func printType(entity interface{}) {
	switch entity.(type) {
	case int:
		fmt.Println("This is a int")
	case string:
		fmt.Println("This is a string")
	case bool:
		fmt.Println("This is a bool")
	case chan int:
		fmt.Println("This is a chan int")
	default:
		fmt.Println("type is not recognized")
	}
}

func main() {
	n := 1
	b := true
	str := "hello"
	ch := make(chan int)

	printType(n)
	printType(b)
	printType(str)
	printType(ch)
	printType(struct{}{})
}
