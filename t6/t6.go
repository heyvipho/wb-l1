package main

import (
	"context"
	"time"
)

func main() {
	// Обычное завершение горутины
	go func() {
	}()

	// Завершение горутины через return
	go func() {
		return
	}()

	// Завершение горутиы через select, закрытие канала и return
	go func() {
		ch := make(chan struct{})
		close(ch)
		select {
		case <-ch:
			return
		}
	}()

	// Завершение горутиы через select, таймер и return
	go func() {
		timer := time.NewTimer(time.Second)
		select {
		case <-timer.C:
			return
		}
	}()

	// Завершение горутиы через select, context и return
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		select {
		case <-ctx.Done():
			return
		}
	}()

	////
	//go func() {
	//}()
}
