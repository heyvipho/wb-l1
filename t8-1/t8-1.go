package main

import (
	"fmt"
	"strconv"
)

func transform(value int64, bitNum, bitValue int) int64 {
	maxInt := int64(^uint64(0) >> 1)

	a := value & ((1 << bitNum) - 1)

	if bitValue == 0 {
		b := value & (maxInt << (bitNum + 1))
		return b | a
	}

	b := value | (1 << bitNum)
	return b | a
}

func main() {
	value := int64(123)
	fmt.Println(strconv.FormatInt(value, 2))
	fmt.Println(strconv.FormatInt(transform(value, 3, 1), 2))
}
