package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Increment struct {
	i int
	m sync.RWMutex
}

func CreateIncrement() *Increment {
	return &Increment{
		i: 0,
	}
}

func (v *Increment) Add() {
	v.m.Lock()
	defer v.m.Unlock()

	v.i++
}

func (v *Increment) Get() int {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.i
}

func worker(inc *Increment, bye chan struct{}) {
	for {
		select {
		default:
			inc.Add()
		case <-bye:
			return
		}
	}
}

func main() {
	bye := make(chan struct{})

	inc := CreateIncrement()

	for i := 0; i < 8; i++ {
		go func() {
			worker(inc, bye)
		}()
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		close(bye)
		fmt.Printf("Счетчик досчитал до: %v\n", inc.Get())
		fmt.Println("До свидания.")
	}
}
