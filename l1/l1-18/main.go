package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	value atomic.Uint64
}

func main() {
	var wg sync.WaitGroup

	var c counter

	for range 10 {
		wg.Add(1)
		go func(c *counter) {
			defer wg.Done()
			for range 1000 {
				c.value.Add(1)
			}
		}(&c)
	}

	wg.Wait()

	fmt.Println(c.value.Load())
}
