package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	arrSize := len(arr)
	workersNum := 3
	arrCh := make(chan int, arrSize)
	resCh := make(chan int, arrSize)

	// создаем воркеров
	for i := 1; i <= workersNum; i++ {
		go worker(arrCh, resCh)
	}

	// отправляем входные данные в arrCh
	for _, value := range arr {
		arrCh <- value
	}
	// закрываем канал после всей отправки
	close(arrCh)

	// ждем и выводим все результаты
	for range arrSize {
		fmt.Println(<-resCh)
	}
}

// worker получает из канала arrCh, преобразует и записывает в канал resCh
func worker(arrCh <-chan int, resCh chan<- int) {
	for el := range arrCh {
		resCh <- el * el
	}
}
