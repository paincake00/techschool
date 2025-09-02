package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// кол-во воркеров
	var n int
	fmt.Println("Enter num of workers:")
	// считывание из стандартного потока вводе
	_, err := fmt.Fscan(os.Stdin, &n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// ограничение на кол-во воркеров
	if n < 1 || n > 10 {
		fmt.Println("Number of workers should be between 1 and 10")
		return
	}

	// общий канал, в который будет поступать постоянно данные
	genCh := make(chan int)
	// запуск N воркеров
	for i := 1; i <= n; i++ {
		go worker(i, genCh)
	}

	// тикер для интервала между добавлениями в канал
	ticker := time.NewTicker(100 * time.Millisecond)
	// канал для завершения главной горутины
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				genCh <- rand.Int()
			}
		}
	}()

	// добавление в течение 5 секунд
	time.Sleep(5 * time.Second)
	// остановка главной горутины
	done <- true
	// закрытие общего канала
	close(genCh)
	fmt.Println("general channel stopped")
}

// worker получает данные, преобразует и выводит в Stdout
func worker(id int, ch <-chan int) {
	// полчение данных из канала вплоть до закрытия
	for i := range ch {
		fmt.Printf("worker %d: %d\n", id, i)
	}
}
