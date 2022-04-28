package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var aStr string
	var bStr string

	fmt.Printf("Введите первое число: ")
	_, err := fmt.Fscan(os.Stdin, &aStr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Введите второе число: ")
	_, err = fmt.Fscan(os.Stdin, &bStr)
	if err != nil {
		panic(err)
	}

	a := new(big.Int)
	a.SetString(aStr, 10)

	b := new(big.Int)
	b.SetString(bStr, 10)

	fmt.Println("Add: ", new(big.Int).Add(a, b))

	fmt.Println("Sub: ", new(big.Int).Sub(a, b))

	fmt.Println("Mul: ", new(big.Int).Mul(a, b))

	fmt.Println("Div: ", new(big.Int).Div(a, b))
}
