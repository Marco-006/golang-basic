# 每个阶段的README文件都需要写清楚这个阶段的收获

# 1. goroutine 与 channel
## goroutine 中 defer 关键字是什么意思？  --- defer 跟 goroutine 无关，只跟‘当前函数’有关；函数一结束，defer 就执行，哪怕你在 goroutine 里
## goroutine 不是通常理解的线程，线程是操作系统调度的. 那么什么是真正的线程？  ---协成
## goroutine 是不是协成？   ---是


# 2. 使用select进行多路复用
## 


# 3. golang中
## 对应题目 6-Goroutine.go
## 在 Go 里没有语言级/标准库级的“线程池”类型；

```
在 Go 里没有语言级/标准库级的“线程池”类型，因为：
M:N 调度模型
用户代码只看见 goroutine（轻量级协程，几 KB 栈）
运行时把 N 个 goroutine 映射到 M 个操作系统线程，并自动复用、抢占
这套映射与管理逻辑官方已经帮你实现好，对开发者透明，所以不需要也不应该自己去维护“线程池”
Goroutine 就是“池”
想并发就 go f()，开几十万个 goroutine 也毫无压力
真正消耗 OS 线程的只有阻塞在系统调用或 runtime.LockOSThread 的场合，数量远小于 goroutine 数
如果想限制并发度，用 “goroutine 池”/“worker 池” 即可，而不是线程池。标准套路是：
把任务发到带缓冲的 chan Job
固定启动 N 个 worker goroutine 循环读取 channel
通过 channel 容量或 semaphore 控制并发级别