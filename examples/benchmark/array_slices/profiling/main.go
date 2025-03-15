package main

import (
	"fmt"
	"runtime/pprof"
	"time"

	"os"
)

func arrayTest() {
	var arr [1000000]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
}

// Function to test memory usage of slices
func sliceTest() {
	slice := make([]int, 1000000)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
}

// Function to test slice appending
func sliceAppendTest() {
	slice := []int{}
	for i := 0; i < 1000000; i++ {
		slice = append(slice, i)
	}
}

func main() {
	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Println("running array test...")
	arrayTest()
	time.Sleep(time.Second)

	fmt.Println("running slice test...")
	sliceTest()
	time.Sleep(time.Second)

	fmt.Println("running slice append test...")
	sliceAppendTest()
	time.Sleep(time.Second)

	memFile, err := os.Create("mem_profile.prof")
	if err != nil {
		fmt.Println("Could not create memory profile:", err)
		return
	}
	defer memFile.Close()
	pprof.WriteHeapProfile(memFile)
	fmt.Println("Profiling complete. Use `go tool pprof` to analyze")
}
