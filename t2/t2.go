package main

import (
	"fmt"
	"sync"
)

func main() {
	// Для того, чтобы высчитать квадраты параллельно, а потом вывести, создаем WaitGroup
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}
	wg.Add(len(arr))

	// Создаем Mutex для того, чтобы не произошло конфликтов при записи массива
	var mutex sync.Mutex
	// Массив квадратов
	resultArr := make([][]int, 0, 5)

	for _, v := range arr {
		go func(v int) {
			// Избавляемся от конфликтов
			mutex.Lock()
			defer mutex.Unlock()

			// Добоавляем квадрат в массив
			resultArr = append(resultArr, []int{v, v * v})

			wg.Done()
		}(v)
	}

	wg.Wait()

	// Готово! Теперь выводим результат
	for _, v := range resultArr {
		fmt.Printf("%v в квадрате = %v\n", v[0], v[1])
	}
}
