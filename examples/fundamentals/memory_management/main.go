package main

import (
	"fmt"
	"runtime"
	"time"
)

var x int

type Object struct {
	id   int
	data [1024]byte
}

func createObjects() {
	for i := 0; i < 5; i++ {
		obj := &Object{id: i} // allocated on the heap
		fmt.Printf("Created object %d at %p\n", obj.id, obj)
	}
}

func main() {
	fmt.Println("Starting program...")

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Memory usage before: HeapAlloc = %v KB\n", mem.HeapAlloc/1024)

	createObjects()

	fmt.Println("Forcing GC...")
	runtime.GC()

	runtime.ReadMemStats(&mem)
	fmt.Printf("Memoery Usage After: HeapAlloc = %v KB\n", mem.HeapAlloc/1024)

	fmt.Println("Waiting before exiting...")
	time.Sleep(2 * time.Second)
}
