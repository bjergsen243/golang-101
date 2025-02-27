package main

import (
	"fmt"
	"sync"
	"time"
)

// to wait for multiple goroutines to finish
// we can use wait group

func worker(id int) {
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// if pass wg to function, should use pointer
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}
