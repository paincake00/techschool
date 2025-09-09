package main

import (
	"fmt"
	"time"
)

func main() {
	time.NewTimer(time.Second)

	fmt.Println(time.Now())
	sleep(3 * time.Second)
	fmt.Println(time.Now())
}

func sleep(d time.Duration) {
	done := make(chan struct{})

	go func() {
		timer := time.NewTimer(d)
		<-timer.C
		close(done)
	}()

	<-done
}
