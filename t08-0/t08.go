package main

import (
	"fmt"
	"strconv"
)

func transform(value int64, bitNum, bitValue int) int64 {
	switch bitValue {
	case 0:
		if value&(1<<bitNum) == 0 {
			return value
		}
		return value ^ (1 << bitNum)
	case 1:
		if value&(1<<bitNum) == 0 {
			return value ^ (1 << bitNum)
		}
		return value
	}

	return value
}

func main() {
	value := int64(123)
	fmt.Println(strconv.FormatInt(value, 2))
	fmt.Println(strconv.FormatInt(transform(value, 3, 1), 2))
}
