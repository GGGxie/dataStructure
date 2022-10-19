package main_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		nums := generateWithCap(10000)
		b.StartTimer()
		bubbleSort(nums)
	}
}
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

var cnt = 10000

var sIdx int
var rIdx int

func BenchmarkProducerConsumerModelNRoutine(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan string)
	c := make(chan string, 100000000)

	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- "hello"
		}
	}

	receive := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2 * cnt)
	for i := 0; i < cnt; i++ {
		go sender()
		go receive()
	}
	b.StartTimer()
	close(begin)
	wg.Wait()
}

func BenchmarkProducerConsumerModelLock(b *testing.B) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	begin := make(chan string)
	pool := make([]string, 100000000)

	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			lock.Lock()
			if sIdx >= 100000000 {
				lock.Unlock()
				return
			}
			pool[sIdx] = "hello"
			sIdx++
			lock.Unlock()
		}
	}

	receive := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			lock.Lock()
			if len(pool) > 0 {
				if rIdx >= 100000000 {
					lock.Unlock()
					return
				}
				// rc := pool[rIdx]
				// handle(rc)
				rIdx++
			}
			lock.Unlock()
		}
	}
	wg.Add(2 * cnt)
	for i := 0; i < cnt; i++ {
		go sender()
		go receive()
	}
	b.StartTimer()
	close(begin)
	wg.Wait()
}
