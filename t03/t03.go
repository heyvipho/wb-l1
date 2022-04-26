package main

import (
	"fmt"
	"sync"
)

func main() {
	// Для того, чтобы высчитать квадраты параллельно, а потом вывести их сумму, создаем WaitGroup
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}
	wg.Add(len(arr))

	// Создаем Mutex для того, чтобы не произошло конфликтов при записи
	var mutex sync.Mutex
	// Сумма квадратов
	result := 0

	for _, v := range arr {
		go func(v int) {
			// Избавляемся от конфликтов
			mutex.Lock()
			defer mutex.Unlock()

			// Прибавляем квадрат
			result += v * v

			wg.Done()
		}(v)
	}

	wg.Wait()

	// Готово! Теперь выводим результат
	fmt.Printf("Сумма квадратов = %v\n", result)
}
