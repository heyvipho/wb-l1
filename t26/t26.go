package main

import (
	"fmt"
	"strings"
)

func checkUniqueSymbols(str string) bool {
	set := make(map[string]struct{})

	arr := strings.Split(str, "")
	for _, v := range arr {
		if _, ok := set[v]; ok {
			return false
		}
		set[v] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println("abcd: ", checkUniqueSymbols("abcd"))
	fmt.Println("abCdefAaf: ", checkUniqueSymbols("abCdefAaf"))
	fmt.Println("aabcd: ", checkUniqueSymbols("aabcd"))
}
