package main

import (
	"fmt"
	"runtime"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct") // running synchronously

	go f("goroutine") // execute concurrently with the calling one

	// start a goroutine for an annonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// wait for goroutines finish
	fmt.Println("waiting goroutines finish") // try WaitGroup
	time.Sleep(time.Second)
	fmt.Println("done")

	go func() {
		for i := 1; i <= 50; i++ {
			fmt.Println("I am Goroutine 1")
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 1; i <= 50; i++ {
			fmt.Println("I am Goroutine 2")
			runtime.Gosched()
		}
	}()

	time.Sleep(time.Second)
}
