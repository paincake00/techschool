package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// кол-во секунд
	var n int
	fmt.Println("Enter seconds:")
	// считывание из стандартного потока вводе
	_, err := fmt.Fscan(os.Stdin, &n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// ограничение на кол-во секунд
	if n < 0 || n > 999 {
		fmt.Println("Seconds value should be between 1 and 999")
		return
	}

	// тикер для интервала между добавлениями в канал
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	ch := make(chan int)
	go func() {
		i := 0
		for {
			<-ticker.C
			ch <- i
			i++
		}
	}()

	// таймаут вне цикла, чтобы он не пересоздавался с каждой итерацией
	timeout := time.After(time.Duration(n) * time.Second)

	for {
		select {
		case i := <-ch:
			fmt.Println("Added:", i)
		case <-timeout:
			fmt.Println("Timed out")
			return
		}
	}
}
