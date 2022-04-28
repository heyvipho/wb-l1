package main

import (
	"fmt"
	"time"
)

func sleep(dur time.Duration) {
	timer := time.NewTimer(dur)
	<-timer.C
}

func main() {
	fmt.Println("Начало отсчета")
	sleep(time.Second * 3)
	fmt.Println("Отсчет завершен")
}
