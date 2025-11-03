package main

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

// func main() {

// }

type Result struct {
	cost int
}

type TaskScheduler struct {
}

// 任务  任务调度器: 1、统一开始、结束等待；2、 使用多少个协程 -- 在开始的时候，就这么定义好了

// Q：一组任务??
// 同时统计每个任务的执行时间  使用select获取返回值

// Q：go 有没有类似线程池的概念？ 有没有类似ForkJoin的机制？
