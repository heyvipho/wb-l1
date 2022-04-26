package main

import (
	"fmt"
	"strconv"
)

func transform(value int64, bitNum int, bitValue int) int64 {
	maxInt := int64(^uint64(0) >> 1) // 1111111...

	if bitValue == 0 {
		mask := int64(maxInt - 1<<bitNum) // ...1111011...
		return value & mask
	} else if bitValue == 1 {
		mask := int64(1 << bitNum) // 1000000...
		return value | mask
	}

	return 0
}

func main() {
	value := int64(123)
	fmt.Println(strconv.FormatInt(value, 2))
	fmt.Println(strconv.FormatInt(transform(value, 3, 0), 2))
}
