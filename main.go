package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var counter int32

func main() {
	counter = 0
	for i := 0; i < 10000; i++ {
		go func(i int) {
			if atomic.CompareAndSwapInt32(&counter, counter, counter+1) {
				time.Sleep(time.Duration(rand.Intn(2)) * time.Microsecond) //加一个睡眠
				if counter > 3 {
					return
				}
				fmt.Println(counter, "处理服务: ", i)
				time.Sleep(time.Hour)
			} else {
			}
		}(i)
	}
	time.Sleep(time.Hour)
}
