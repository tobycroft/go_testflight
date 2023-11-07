package main

import (
	"github.com/tobycroft/Calc"
	"testing"
)

func Benchmark_any2String(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Calc.Any2String(1234567890)
	}
	b.StopTimer()
}
