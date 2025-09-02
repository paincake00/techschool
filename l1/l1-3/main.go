package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Println("Enter num of workers:")
	_, err := fmt.Fscan(in, &n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if n < 1 || n > 10 {
		fmt.Println("Number of workers should be between 1 and 10")
		return
	}

	genCh := make(chan int)
	for i := 1; i <= n; i++ {
		go worker(i, genCh)
	}

	ticker := time.NewTicker(100 * time.Millisecond)
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

	time.Sleep(5 * time.Second)
	done <- true
	close(genCh)
	fmt.Println("general channel stopped")
}

func worker(id int, ch <-chan int) {
	for i := range ch {
		fmt.Printf("worker %d: %d\n", id, i)
	}
}
