package main

import (
	"fmt"
	"time"
)

//轮流打印字母和数字

<<<<<<< HEAD
<<<<<<< HEAD
// 两个channel  一个控制字符 一个控制数字
=======
>>>>>>> 9c616fa... 测试
=======
>>>>>>> 41880d2a61d67bbf8a98c7c0711e733b45aa8b6c
func main() {
	strChan := make(chan int, 1)
	numChan := make(chan int, 1)
	strChan <- 0
	go func() {
		for i := 65; i <= 90; i++ {
			<-strChan
			fmt.Printf("%v ", string(rune(i)))
			numChan <- i
		}
		return
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-numChan
			fmt.Printf("%v ", i)
			strChan <- i
		}
		return
	}()
	time.Sleep(1 * time.Second)
	fmt.Println()
}
