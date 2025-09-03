package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]int{"k1": 0, "k2": 0}

	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(3)
	// вызов трез горутин, которые конкурентно обращются к map
	go addNForKey(&wg, &mu, m, "k1", 100)
	go addNForKey(&wg, &mu, m, "k1", 80)
	go addNForKey(&wg, &mu, m, "k2", 100)
	wg.Wait()

	fmt.Println(m)
}

// выполенение функции инкерментирования N раз
func addNForKey(wg *sync.WaitGroup, mu *sync.Mutex, m map[string]int, k string, n int) {
	defer wg.Done()
	for range n {
		incrementByKey(mu, m, k)
	}
}

// инкерементирование значения по ключу в map
func incrementByKey(mu *sync.Mutex, m map[string]int, k string) {
	// блокируем mutex (mutual exclusion)
	mu.Lock()
	// освобождаем mutex
	defer mu.Unlock()

	m[k]++
}
