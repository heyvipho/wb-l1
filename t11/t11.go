package main

import "fmt"

func main() {
	arr1 := []string{"yellow", "green", "red", "blue"}
	arr2 := []string{"blue", "red", "white", "black"}

	set := make(map[string]struct{})
	result := make([]string, 0)

	// Делаем Set из первого массива
	for _, v := range arr1 {
		set[v] = struct{}{}
	}

	// Добавляем в массив result перескающиеся элементы
	for _, v := range arr2 {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}

	fmt.Println(result)
}
