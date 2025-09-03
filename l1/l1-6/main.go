package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// остановка по условию
func stopByCondition(wg *sync.WaitGroup) {
	defer wg.Done()
	running := true
	go func(isRun *bool) {
		for {
			// исполнение обычного условия
			if !*isRun {
				fmt.Println("stop by condition: goroutine is stopped")
				return
			}
			fmt.Println("stop by condition: goroutine is running")
			time.Sleep(time.Second)
		}
	}(&running)
	time.Sleep(2 * time.Second)
	running = false
	time.Sleep(time.Second)
}

// остановка через канал
func stopByChan(wg *sync.WaitGroup) {
	defer wg.Done()
	ch := make(chan bool)
	go func(chIn <-chan bool) {
		for {
			select {
			// ожидание получения значения из канала
			case <-chIn:
				fmt.Println("stop by chan: goroutine is stopped")
				return
			default:
				fmt.Println("stop by chan: goroutine is running")
				time.Sleep(time.Second)
			}
		}
	}(ch)
	time.Sleep(2 * time.Second)
	ch <- true
	time.Sleep(time.Second)
}

// остановка через контекст
func stopByContext(wg *sync.WaitGroup) {
	defer wg.Done()
	cnt, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			// ожидание сигнала отмены контекста
			case <-ctx.Done():
				fmt.Println("stop by context: goroutine is stopped")
				return
			default:
				fmt.Println("stop by context: goroutine is running")
				time.Sleep(time.Second)
			}
		}
	}(cnt)
	time.Sleep(2 * time.Second)
	// отмена контекста
	cancel()
	time.Sleep(time.Second)
}

// остановка через runtime.GoExit
func stopByGoExit(wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		// вызывается перед GoExit
		defer fmt.Println("stop by go exit: goroutine is stopped")
		fmt.Println("stop by go exit: goroutine is running")
		time.Sleep(500 * time.Microsecond)
		// немедленное завершение горутины
		runtime.Goexit()
	}()
	time.Sleep(time.Second)
}

// остановка через закрытие канала
func stopByClosingChan(wg *sync.WaitGroup) {
	defer wg.Done()
	ch := make(chan int)
	go func(chIn <-chan int) {
		// получение значений до закрытия канала
		for i := range chIn {
			fmt.Printf("stop by closing chan: goroutine is running with value: %d\n", i)
		}
		fmt.Println("stop by closing chan: goroutine is stopped")
	}(ch)
	ch <- 4
	ch <- 5
	close(ch)
	time.Sleep(time.Second)
}

// остановка через таймаутом
func stopByTimeout(wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		timeout := time.After(2 * time.Second)
		for {
			select {
			// завершение по истечению таймаута
			case <-timeout:
				fmt.Println("stop by timeout: goroutine is stopped")
				return
			default:
				fmt.Println("stop by timeout: goroutine is running")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
}

// остановка через recover
func stopByRecover(wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		defer func() {
			// перехватываем panic и завершаем горутину
			if err := recover(); err != nil {
				fmt.Println("stop by recover: goroutine is stopped")
			}
		}()
		fmt.Println("stop by recover: goroutine is running")
		time.Sleep(500 * time.Millisecond)
		panic("sudden error")
	}()
	time.Sleep(time.Second)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(7)

	stopByCondition(&wg)
	stopByChan(&wg)
	stopByContext(&wg)
	stopByGoExit(&wg)
	stopByClosingChan(&wg)
	stopByTimeout(&wg)
	stopByRecover(&wg)

	wg.Wait()
	fmt.Println("All goroutines stopped")
}
