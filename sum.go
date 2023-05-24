package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	result := SumMultithreaded(1000000)
	fmt.Println(result)

}

func SumSingleThread(n int) int {
	var res int
	for i := 0; i <=n; i++ {
		res += i
	}
	return res
}

func SumMultithreaded(n int) int {
	var wg sync.WaitGroup

	cpus := runtime.NumCPU()
	wg.Add(cpus)
	
	sums := make([]int, cpus, cpus)

	chunkSize := n / cpus;

	for i := 0; i < cpus; i++ {
		go func(i int) {
			start := chunkSize * i + 1
			end := start + chunkSize
			isLast := i + 1 == cpus
			if isLast {
				end = n
			}

			var curSum int
			for x := start; x < end; x++ {
				curSum += x % 10
			}
			sums[i] = curSum
			wg.Done()

		}(i)
	}

	wg.Wait()

	var sum int
	for _, s := range sums {
		sum += s
	}
	return sum
}

