package main

import "fmt"

func main7() {
	// version1  // 创建一个带缓冲的channel
	ch := make(chan int, 10)
	sendOnly7(ch)
	readOnly7(ch)
}

func readOnly7(ch <-chan int) {
	// 通道关闭后自动退出
	for v := range ch {
		fmt.Println("read， ", v)
	}
}

func sendOnly7(ch chan<- int) {
	for i := 1; i < 11; i++ {
		ch <- i
		fmt.Println("send， ", i)
	}
	close(ch) // 发送完关闭，通知消费者退出
}

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 	  一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
