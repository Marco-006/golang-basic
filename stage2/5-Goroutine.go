package main

import (
	"fmt"
	// "sync"
	"time"
)

func main5() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	go func() {
		// defer wg.Done()
		fmt.Println("test1--- start")
		for i := 1; i < 10; i++ {
			if i&1 != 0 {
				fmt.Println("test1, ", i)
			}
		}
		fmt.Println("test1--- end")
	}()
	go func() {
		// defer wg.Done()
		fmt.Println("test2--- start")
		for i := 2; i < 10; i++ {
			if i&1 == 0 {
				fmt.Println("test2, ", i)
			}
		}
		fmt.Println("test2--- end")
	}()
	// wg.Wait()
	time.Sleep(500 * time.Millisecond)
}

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。

// func test1() {
// 	defer wg.Done()
// 	fmt.Println("test1--- start")
// 	for i := 1; i < 10; i++ {
// 		if i&1 != 0 {
// 			fmt.Println("test1, ", i)
// 		}
// 	}
// 	fmt.Println("test1--- end")
// }

// func test2() {
// 	defer wg.Done()
// 	fmt.Println("test2--- start")
// 	for i := 2; i < 10; i++ {
// 		if i&1 == 0 {
// 			fmt.Println("test2, ", i)
// 		}
// 	}
// 	fmt.Println("test2--- end")
// }

// Q： go 启动协成后，为什么没有启动起来呢？
// A:  主线程需要等待一下，否则多线程只启动还没开始执行的时候，主线程就结束了

// Q：go 中线程暂停的办法？time.Sleep(500 * time.Millisecond)
