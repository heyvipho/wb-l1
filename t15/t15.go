package main

import "fmt"

func createHugeString(amount int) string {
	str := ""
	for i := 0; i < amount; i++ {
		str += "a"
	}
	return str
}

func someFunc() string {
	// Тут непонятно, зачем мы создаем строку из 1024 символов, а потом ее обрезаем, но ладно
	v := createHugeString(1 << 10)
	// Вместо записи в глоальную переменную, вернем значение через функцию. Использование глобальной переменной нежелательно, так как в многопоточном режиме могут произойти конфликты
	return v[:100]
}

func main() {
	// Сделаем локальную переменную вместо глобальной
	justString := someFunc()
	fmt.Println(justString)
}
