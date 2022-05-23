package main_test

import (
	"sync"
	"testing"
)

// func setup() {
// 	fmt.Println("Before all tests")
// }

// func teardown() {
// 	fmt.Println("After all tests")
// }

// func Test1(t *testing.T) {
// 	fmt.Println("I'm test1")
// }

// func Test2(t *testing.T) {
// 	fmt.Println("I'm test2")
// }

// // func TestMain(m *testing.M) {
// // 	setup()
// // 	code := m.Run()
// // 	teardown()
// // 	os.Exit(code)
// // }

// // func Test_AA(t *testing.T) {
// // 	a := []struct {
// // 		A int
// // 	}{
// // 		{1},
// // 		{2},
// // 	}
// // 	// score := getScore(100, 100, 100, 2.1)
// // 	// fmt.Println(111)
// // 	// Output:
// // 	// 31.1
// // 	t.Run("ps", sort.F2)
// // 	t.Helper()
// // }

// // func TestMain(m *testing.M) {
// // 	m.Run()
// // }

// type Duck interface {
// 	Quack()
// }

// type Cat struct {
// 	Name string
// }

// func (c Cat) Quack() {}

// func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
// func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
// func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
// func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
// func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }

// // func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

// func benchmarkFib(i int, b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		Fib(i)
// 	}
// }
// func Fib(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	return Fib(n-1) + Fib(n-1)
// }

// // func BenchmarkDirectCall(b *testing.B) {
// // 	c := Cat{Name: "draven"}
// // 	for n := 0; n < b.N; n++ {
// // 		// MOVQ	AX, (SP)
// // 		// MOVQ	$6, 8(SP)
// // 		// CALL	"".Cat.Quack(SB)
// // 		c.Quack()
// // 	}
// // }

// // func BenchmarkDynamicDispatch(b *testing.B) {
// // 	c := Duck(Cat{Name: "draven"})
// // 	for n := 0; n < b.N; n++ {
// // 		// MOVQ	16(SP), AX
// // 		// MOVQ	24(SP), CX
// // 		// MOVQ	AX, "".d+32(SP)
// // 		// MOVQ	CX, "".d+40(SP)
// // 		// MOVQ	"".d+32(SP), AX
// // 		// MOVQ	24(AX), AX
// // 		// MOVQ	"".d+40(SP), CX
// // 		// MOVQ	CX, (SP)
// // 		// CALL	AX
// // 		c.Quack() //值类型的变量实现接口对性能影响较大
// // 	}
// // }

// func BenchmarkDynamicDispatch2(b *testing.B) {
// 	c := Duck(&Cat{Name: "draven"})
// 	for n := 0; n < b.N; n++ {
// 		// MOVQ	"".d+56(SP), AX
// 		// MOVQ	24(AX), AX
// 		// MOVQ	"".d+64(SP), CX
// 		// MOVQ	CX, (SP)
// 		// CALL	AX
// 		c.Quack() //指针类型的变量实现接口对性能影响小一点,但还是有影响
// 	}
// }

// func BenchmarkDynamicDispatch3(b *testing.B) {
// 	c := Duck(&Cat{Name: "draven"})
// 	for n := 0; n < b.N; n++ {
// 		c.(*Cat).Quack() //断言后再调用，几乎没有影响
// 	}
// }

//o test -benchmem -bench . -benchtime=10000x
//10000个消费者,10000个生产者,每个生产者生产10000条数据
//用channel和用lock比较
//结论:channel最后性能更高
//分析,用lock不好做,如果每次都加判断那么性能就会降低
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
