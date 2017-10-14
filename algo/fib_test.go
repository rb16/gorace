package main

import (
	"testing"
)

// benchmarkfib a helper function for fibnocci number
func benchmarkfib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib(b *testing.B) {
	//run the fib function b.N times
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

func Benchmark1(b *testing.B)  { benchmarkfib(1, b) }
func Benchmark2(b *testing.B)  { benchmarkfib(2, b) }
func Benchmark3(b *testing.B)  { benchmarkfib(3, b) }
func Benchmark4(b *testing.B)  { benchmarkfib(4, b) }
func Benchmark5(b *testing.B)  { benchmarkfib(5, b) }
func Benchmark6(b *testing.B)  { benchmarkfib(6, b) }
func Benchmark40(b *testing.B) { benchmarkfib(50, b) }
