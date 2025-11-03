package main

import (
	"fmt"
	"sync"
	"time"
)

// TaskStat9 保存单个任务的统计信息
type TaskStat9 struct {
	ID   int
	Name string
	Cost time.Duration
}

// Scheduler9 任务调度器
type Scheduler9 struct {
	stats []TaskStat9    // 统计结果
	mu    sync.Mutex     // 保护 stats
	wg    sync.WaitGroup // 等待所有任务完成
}

// NewScheduler9 构造器
func NewScheduler9() *Scheduler9 {
	return &Scheduler9{stats: make([]TaskStat9, 0)}
}

// Submit 提交一个任务
// id 与 name 仅用于标识，方便查看结果
func (s *Scheduler9) Submit(id int, name string, task func()) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		// start := time.Now()
		s.mu.Lock()
		task() // 真正执行任务
		s.mu.Unlock()
		// cost := time.Since(start)

		// 线程安全地记录结果
		// s.mu.Lock()
		// s.stats = append(s.stats, TaskStat9{ID: id, Name: name, Cost: cost})
		// s.mu.Unlock()
	}()
}

// Wait 等待所有任务完成，并返回统计结果
func (s *Scheduler9) Wait() []TaskStat9 {
	s.wg.Wait()
	return s.stats
}

func main9() {
	//共享计数器
	index := 0

	// 构造 5 个模拟任务
	tasks := []func(){
		func() { increment(&index) },
		func() { increment(&index) },
		func() { increment(&index) },
		func() { increment(&index) },
		func() { increment(&index) },
		func() { increment(&index) },
		func() { increment(&index) },
		func() { increment(&index) },
		func() { increment(&index) },
		func() { increment(&index) },
	}

	sched := NewScheduler9()

	// 提交任务
	for i, t := range tasks {
		sched.Submit(i+1, fmt.Sprintf("task-%d", i+1), t)
	}

	// 阻塞直到全部完成
	sched.Wait()

	// 打印共享变量
	fmt.Println("共享变量： ", index)
}

func increment(index *int) {
	for i := 0; i < 1000; i++ {
		(*index)++
	}
}

// 难点

// 需求：
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
