package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//实现自旋锁
var Count int32 = 1

func SpinLock(i int32, fn func()) {
	for {
		if n := atomic.LoadInt32(&Count); n == i {
			fn()
			atomic.AddInt32(&Count, 1)
			break
		}
		time.Sleep(time.Nanosecond)
	}
}
func main() {
	for i := int32(1); i < 4; i++ {
		go func(i int32) {
			fn := func() {
				fmt.Println(i)
			}
			SpinLock(i, fn)
		}(i)
	}
	SpinLock(4, func() {})
}
