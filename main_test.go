package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"unicode"
)

// func generateWithCap(n int) []int {
// 	rand.Seed(time.Now().UnixNano())
// 	nums := make([]int, 0, n)
// 	for i := 0; i < n; i++ {
// 		nums = append(nums, rand.Int())
// 	}
// 	return nums
// }

// func bubbleSort(nums []int) {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 1; j < len(nums)-i; j++ {
// 			if nums[j] < nums[j-1] {
// 				nums[j], nums[j-1] = nums[j-1], nums[j]
// 			}
// 		}
// 	}
// }

// func BenchmarkBubbleSort(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		b.StopTimer()
// 		nums := generateWithCap(10000)
// 		b.StartTimer()
// 		bubbleSort(nums)
// 	}
// }
// func BenchmarkFib(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		fib(30) // run fib(30) b.N times
// 	}
// }

// func fib(n int) int {
// 	if n == 0 || n == 1 {
// 		return n
// 	}
// 	return fib(n-2) + fib(n-1)
// }

// var cnt = 10000

// var sIdx int
// var rIdx int

// func BenchmarkProducerConsumerModelNRoutine(b *testing.B) {
// 	var wg sync.WaitGroup
// 	begin := make(chan string)
// 	c := make(chan string, 100000000)

// 	sender := func() {
// 		defer wg.Done()
// 		<-begin
// 		for i := 0; i < b.N; i++ {
// 			c <- "hello"
// 		}
// 	}

// 	receive := func() {
// 		defer wg.Done()
// 		<-begin
// 		for i := 0; i < b.N; i++ {
// 			<-c
// 		}
// 	}
// 	wg.Add(2 * cnt)
// 	for i := 0; i < cnt; i++ {
// 		go sender()
// 		go receive()
// 	}
// 	b.StartTimer()
// 	close(begin)
// 	wg.Wait()
// }

// func BenchmarkProducerConsumerModelLock(b *testing.B) {
// 	var wg sync.WaitGroup
// 	var lock sync.Mutex
// 	begin := make(chan string)
// 	pool := make([]string, 100000000)

// 	sender := func() {
// 		defer wg.Done()
// 		<-begin
// 		for i := 0; i < b.N; i++ {
// 			lock.Lock()
// 			if sIdx >= 100000000 {
// 				lock.Unlock()
// 				return
// 			}
// 			pool[sIdx] = "hello"
// 			sIdx++
// 			lock.Unlock()
// 		}
// 	}

// 	receive := func() {
// 		defer wg.Done()
// 		<-begin
// 		for i := 0; i < b.N; i++ {
// 			lock.Lock()
// 			if len(pool) > 0 {
// 				if rIdx >= 100000000 {
// 					lock.Unlock()
// 					return
// 				}
// 				// rc := pool[rIdx]
// 				// handle(rc)
// 				rIdx++
// 			}
// 			lock.Unlock()
// 		}
// 	}
// 	wg.Add(2 * cnt)
// 	for i := 0; i < cnt; i++ {
// 		go sender()
// 		go receive()
// 	}
// 	b.StartTimer()
// 	close(begin)
// 	wg.Wait()
// }

// func Test_0(t *testing.T) {
// 	ch := make(chan int, 1)
// 	ch2 := make(chan int, 1)
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10; i += 2 {
// 			<-ch
// 			t.Log(i)
// 			ch2 <- 1
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 1; i < 10; i += 2 {
// 			<-ch2
// 			t.Log(i)
// 			ch <- 1
// 		}
// 	}()
// 	ch <- 1
// 	wg.Wait()
// 	close(ch)
// 	close(ch2)
// }

// 测试函数,Test*
func TestA(t *testing.T) {

}
func F(s string) bool {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	return true
}

// 基准测试,Benchmark*
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		F("A man, a plan, a canal: Panama")
	}
}

func BenchmarkP(b *testing.B) {
	slice := []string{"http://www.baidu.com", "http://www.baidu.com"}
	for i := range slice {
		fetch2(slice[i])
	}
}
func BenchmarkP2(b *testing.B) {
	ch := make(chan string)
	slice := []string{"http://www.baidu.com", "http://www.baidu.com"}
	length := len(slice)
	for i := range slice {
		go fetch(slice[i], ch)
	}
	for i := 0; i < length; i++ {
		<-ch
	}
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func fetch2(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// // 实例函数,作为文档使用
// func ExampleA() {
// }
