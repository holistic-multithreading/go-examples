package main

import (
	"testing"
)


func BenchmarkSimpleSumThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSimple(1000)
	}
}

func BenchmarkSumMultithreadedThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumMultithreaded(1000)
	}
}

func BenchmarkSimpleSumMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSimple(1000000)
	}
}

func BenchmarkSumMultithreadedMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumMultithreaded(1000000)
	}
}

func BenchmarkSimpleSumBillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSimple(1000000000)
	}
}

func BenchmarkSumMultithreadedBillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumMultithreaded(1000000000)
	}
}
