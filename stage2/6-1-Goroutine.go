package main

import (
	"fmt"
	"sync"
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

		start := time.Now()
		task() // 真正执行任务
		cost := time.Since(start)

		// 线程安全地记录结果
		s.mu.Lock()
		s.stats = append(s.stats, TaskStat{ID: id, Name: name, Cost: cost})
		s.mu.Unlock()
	}()
}

// Wait 等待所有任务完成，并返回统计结果
func (s *Scheduler) Wait() []TaskStat {
	s.wg.Wait()
	return s.stats
}

func main6() {
	// 构造 5 个模拟任务
	tasks := []func(){
		func() { time.Sleep(100 * time.Millisecond) },
		func() { time.Sleep(200 * time.Millisecond) },
		func() { time.Sleep(150 * time.Millisecond) },
		func() { time.Sleep(50 * time.Millisecond) },
		func() { time.Sleep(300 * time.Millisecond) },
	}

	sched := NewScheduler()

	// 提交任务
	for i, t := range tasks {
		sched.Submit(i+1, fmt.Sprintf("task-%d", i+1), t)
	}

	// 阻塞直到全部完成
	results := sched.Wait()

	// 打印统计
	fmt.Println("----------- 执行完毕，耗时统计 -----------")
	for _, r := range results {
		fmt.Printf("[%s] cost = %v\n", r.Name, r.Cost)
	}
}

// 难点
// 1. 任务，指针传递，并在具体的调度器方法体重执行（类似于在）
// 2. 任务调度器 组合任务
