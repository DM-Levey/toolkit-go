package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printNumbers(start, end int, wg *sync.WaitGroup, ch1, ch2 chan struct{}) {
	defer wg.Done()

	for i := start; i <= end; i += 3 {
		<-ch1 // 等待通道可用
		fmt.Println(i)
		ch2 <- struct{}{} // 释放通道，让下一个协程打印
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	ch1 := make(chan struct{}, 1) // 创建一个容量为 1 的通道
	ch2 := make(chan struct{}, 1) // 创建一个容量为 1 的通道
	ch3 := make(chan struct{}, 1) // 创建一个容量为 1 的通道
	ch1 <- struct{}{}             // 初始化通道，使第一个协程可以开始打印

	var wg sync.WaitGroup
	wg.Add(3)

	go printNumbers(1, 100, &wg, ch1, ch2)

	go printNumbers(3, 100, &wg, ch3, ch1)

	go printNumbers(2, 100, &wg, ch2, ch3)

	wg.Wait()
}
