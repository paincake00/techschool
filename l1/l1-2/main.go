package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	arrSize := len(arr)
	workersNum := 3
	arrCh := make(chan int, arrSize)
	resCh := make(chan int, arrSize)

	for i := 1; i <= workersNum; i++ {
		go worker(arrCh, resCh)
	}

	for _, value := range arr {
		arrCh <- value
	}
	close(arrCh)

	for range arrSize {
		fmt.Println(<-resCh)
	}
}

func worker(arrCh <-chan int, resCh chan<- int) {
	for el := range arrCh {
		resCh <- el * el
	}
}
