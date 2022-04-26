package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	var seconds time.Duration
	fmt.Print("Введите количество секунд: ")
	_, err := fmt.Fscan(os.Stdin, &seconds)
	if err != nil {
		panic(err)
	}

	// Запускаем постоянную запись данных в канал
	go writer(ch, bye)
	// Запускаем постоянное чтение данных из канала
	go worker(ch, bye)

	// Сделаем так, чтобы процесс не завершался
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	timer := time.NewTimer(time.Second * seconds)

	select {
	case <-stop:
		close(bye)
		fmt.Println("До свидания.")
	case <-timer.C:
		close(bye)
		fmt.Println("До свидания.")
	}
}
