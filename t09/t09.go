package main

import (
	"fmt"
	"time"
)

type Conveyor struct {
	Ch       chan int
	squareCh chan int
	bye      chan struct{}
}

func CreateConveyor() Conveyor {
	conveyor := Conveyor{
		Ch:       make(chan int),
		squareCh: make(chan int),
		bye:      make(chan struct{}),
	}

	return conveyor
}

func (v *Conveyor) Run() {
	// Принимает значение и отправляет в первый канал
	go func() {
		for {
			select {
			case data := <-v.Ch:
				v.squareCh <- data
			case <-v.bye:
				return
			}
		}
	}()

	// Принимает значение из первого канала и отправляет его квадрат во второй канал
	go func() {
		for {
			select {
			case data := <-v.Ch:
				v.squareCh <- data * data
			case <-v.bye:
				return
			}
		}
	}()

	// Печатает значения квадратов из второго канала
	go func() {
		for {
			select {
			case data := <-v.squareCh:
				fmt.Println(data)
			case <-v.bye:
				return
			}
		}
	}()
}

func (v *Conveyor) Stop() {
	close(v.bye)
}

func main() {
	conveyor := CreateConveyor()
	conveyor.Run()
	defer conveyor.Stop()

	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		go func(v int) {
			conveyor.Ch <- v
		}(v)
	}

	time.Sleep(time.Second * 5)
}
