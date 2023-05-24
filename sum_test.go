package main

import (
	"testing"
)


func BenchmarkSingleThreadThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSingleThread(1000)
	}
}

func BenchmarkMultithreadedThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumMultithreaded(1000)
	}
}

func BenchmarkSingleThreadMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSingleThread(1000000)
	}
}

func BenchmarkMultithreadedMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumMultithreaded(1000000)
	}
}

func BenchmarkSingleThreadBillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSingleThread(1000000000)
	}
}

func BenchmarkMultithreadedBillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumMultithreaded(1000000000)
	}
}
