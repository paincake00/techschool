package main

import (
	"fmt"
)

func main() {
	// Исходный массив чисел
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// канал содержит x
	ch1 := make(chan int)
	// канал содержит x*2
	ch2 := make(chan int)

	// запуск генератора (берет из массивы)
	go generator(nums[:], ch1)
	// запуск обработчика (x -> x*2)
	go handler(ch1, ch2)

	// чтение рещультатов x*2 до закрытия канала
	for i := range ch2 {
		fmt.Printf("new value: %d\n", i)
	}
	fmt.Println("Printer stopped")
	fmt.Println("All goroutines finished")
}

func generator(nums []int, out chan<- int) {
	for _, x := range nums {
		out <- x
	}
	fmt.Println("Generator stopped")
	close(out)
}

func handler(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * 2
	}
	fmt.Println("Handler stopped")
	close(out)
}
