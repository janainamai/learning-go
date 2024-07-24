package main

import (
	"testing"
)

// Benchmark para a função SumLoop
func BenchmarkSumLoop(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		SumLoop(nums)
	}
}

// Benchmark para a função SumRecursion
func BenchmarkSumRecursion(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		SumRecursion(nums)
	}
}
