package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Реализует постоянную запись данных в канал
func writer(ch chan int, bye chan struct{}) {
	for i := 0; ; i++ {
		select {
		default:
			ch <- i
		case <-bye:
			return
		}
	}
}

// Выводит произвольные данные
func worker(ch chan int, bye chan struct{}) {
	for {
		select {
		default:
			fmt.Println(<-ch)
		case <-bye:
			return
		}
	}
}

func main() {
	bye := make(chan struct{})
	ch := make(chan int)

	// Запускаем постоянную запись данных в канал
	go writer(ch, bye)

	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	_, err := fmt.Fscan(os.Stdin, &workerCount)
	if err != nil {
		panic(err)
	}

	// Создаем воркеры количеством workerCount
	for i := 0; i < workerCount; i++ {
		go worker(ch, bye)
	}

	// Сделаем так, чтобы процесс не завершался
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		close(bye)
		fmt.Println("До свидания.")
	}
}
