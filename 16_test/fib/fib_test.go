package fib

import "testing"

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func benchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func benchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func benchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func benchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func benchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func benchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
