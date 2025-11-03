package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// TaskStat 保存单个任务的统计信息
type TaskStat struct {
	ID   int
	Name string
	Cost time.Duration
}

// Scheduler 任务调度器
type Scheduler struct {
	stats []TaskStat     // 统计结果
	mu    sync.Mutex     // 保护 stats
	wg    sync.WaitGroup // 等待所有任务完成
}

// NewScheduler 构造器
func NewScheduler() *Scheduler {
	return &Scheduler{stats: make([]TaskStat, 0)}
}

// Submit 提交一个任务
// id 与 name 仅用于标识，方便查看结果
func (s *Scheduler) Submit(id int, name string, task func()) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		// start := time.Now()
		// s.mu.Lock()
		task() // 真正执行任务
		// s.mu.Unlock()
		// cost := time.Since(start)

		// 线程安全地记录结果
		// s.mu.Lock()
		// s.stats = append(s.stats, TaskStat{ID: id, Name: name, Cost: cost})
		// s.mu.Unlock()
	}()
}

// Wait 等待所有任务完成，并返回统计结果
func (s *Scheduler) Wait() []TaskStat {
	s.wg.Wait()
	return s.stats
}

func main10() {
	//共享计数器
	var index int64 = 0

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

	sched := NewScheduler()

	// 提交任务
	for i, t := range tasks {
		sched.Submit(i+1, fmt.Sprintf("task-%d", i+1), t)
	}

	// 阻塞直到全部完成
	sched.Wait()

	// 打印共享变量
	fmt.Println("共享变量： ", index)
}

func increment(index *int64) {
	for i := 0; i < 1000; i++ {
		// (*index)++
		atomic.AddInt64(index, 1)
	}
}

// 难点
// Q: 默认int  是int32还是 int64？ --平台相关（platform-specific）

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全
