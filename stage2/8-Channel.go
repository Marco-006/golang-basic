package main

import (
	"fmt"
	"sync"
)

func main8() {
	// version1  // 创建一个带缓冲的channel

	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 50)
	// sendOnly(ch, wg)
	// readOnly(ch, wg)

	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			ch <- i
			fmt.Println("send， ", i)
		}
		close(ch) // 发送完关闭，通知消费者退出
	}()

	go func() {
		defer wg.Done()

		// 通道关闭后自动退出
		for v := range ch {
			fmt.Println("read， ", v)
		}
	}()
	wg.Wait()

}

// func readOnly(ch <-chan int, wg sync.WaitGroup) {
// 	defer wg.Done()

// 	// 通道关闭后自动退出
// 	for v := range ch {
// 		fmt.Println("read， ", v)
// 	}

// }

// func sendOnly(ch chan<- int, wg sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i < 100; i++ {
// 		ch <- i
// 		fmt.Println("send， ", i)
// 	}
// 	close(ch) // 发送完关闭，通知消费者退出
// }

// Q： 缓冲机制需要和  sync.WaitGroup  搭配一起使用；否则当缓冲队列填满的时候会发生阻塞，进而发生死锁
