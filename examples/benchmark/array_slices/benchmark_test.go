package main

import (
	"fmt"
	"runtime/pprof"
	"time"

	"os"
	"testing"
)

func BenchmarkArray(b *testing.B) {
	var arr [1000000]int
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr); j++ {
			arr[j] = j
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	slice := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(slice); j++ {
			slice[j] = j
		}
	}
}

func BenchmarkSliceAppend(b *testing.B) {
	slice := []int{}
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
}

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
	fmt.Println("Run benchmarks using: `go test -bench=.`")

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

/* Result:
goos: darwin
goarch: arm64
pkg: github.com/bjergsen243/golang-101/examples/benchmark/array_slices
cpu: Apple M2 Pro
BenchmarkArray-10                   4012            294549 ns/op
BenchmarkSlice-10                   3968            302529 ns/op
BenchmarkSliceAppend-10         153837738                7.708 ns/op
PASS
ok      github.com/bjergsen243/golang-101/examples/benchmark/array_slices       5.615s
*/

/*
-> append and iteration through slice is very slow due to reallocation on heap
-> iteration on array is fast cause using stack
*/
