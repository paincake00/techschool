package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// канал для принятия сигнала
	stop := make(chan os.Signal, 1)
	// "прослушка" на сигнал SIGINT
	signal.Notify(stop, syscall.SIGINT)

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

	// создаем группу для ожидания завершения горутин-воркеров
	var wg sync.WaitGroup

	// контекст с отменой, который будет передаваться в каждый воркер
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// общий канал, в который будет поступать постоянно данные
	genCh := make(chan int)
	// запуск N воркеров
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i, genCh)
	}

	// тикер для интервала между добавлениями в канал
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				// отмена контекста произошла
				// закрытие общего канала
				close(genCh)
				return
			case <-ticker.C:
				genCh <- rand.Int()
			}
		}
	}()

	select {
	case <-time.After(5 * time.Second):
		// завершаем по истечении таймера, вызвав cancel()
		fmt.Println("Timer is done, shutting down...")
		cancel()
	case <-stop:
		// досрочное завершение пользователем
		fmt.Println("Received stop signal, shutting down...")
		cancel()
	}

	// ждём завершения всех воркеров
	wg.Wait()
	fmt.Println("All workers finished")
}

// worker получает данные, преобразует и выводит в Stdout
func worker(ctx context.Context, wg *sync.WaitGroup, id int, ch <-chan int) {
	// уменьшение счетчика при завершении воркера
	defer wg.Done()

	// полчение данных из канала вплоть до закрытия
	for {
		select {
		case <-ctx.Done():
			// контекст отменен, воркер завершает работу
			fmt.Printf("Worker %d is done\n", id)
			return
		case i, ok := <-ch:
			if !ok {
				// ok==false канал закрыт, новых значений не будет
				return
			}
			fmt.Printf("worker %d: %d\n", id, i)
		}
	}
}
