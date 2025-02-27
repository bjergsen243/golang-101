package main

import (
	"fmt"
	"time"
)

// use tickers when you want to do something repeatedly at regular intervals

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()

	// this will sleep 1600 milliseconds
	// then the for loop will print 3 times ( 500 * 3 = 1500 < 1600)
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
