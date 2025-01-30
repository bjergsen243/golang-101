package main

import (
	"fmt"
	"testing"
)

func IntMint(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMintBasic(t *testing.T) {
	ans := IntMint(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMintTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMint(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarIntMint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMint(1, 2)
	}
}
