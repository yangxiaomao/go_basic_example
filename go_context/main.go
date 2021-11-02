package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	fmt.Println("key:", key)
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:

		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

// func worker(ctx context.Context) {
// LOOP:
// 	for {
// 		fmt.Println("db connecting ...")
// 		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
// 		select {
// 		case <-ctx.Done(): // 50毫秒后自动调用
// 			break LOOP
// 		default:
// 		}
// 	}
// 	fmt.Println("worker done!")
// 	wg.Done()
// }

// func main() {
// 	// 设置一个50毫秒的超时
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
// 	wg.Add(1)
// 	go worker(ctx)
// 	time.Sleep(time.Second * 5)
// 	cancel() // 通知子goroutine结束
// 	wg.Wait()
// 	fmt.Println("over")
// }

// func gen(ctx context.Context) <-chan int {
// 	dst := make(chan int)
// 	n := 1
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return // return结束该goroutine, 防止泄露
// 			case dst <- n:
// 				n++
// 			}
// 		}
// 	}()
// 	return dst
// }

// func main() {
// d := time.Now().Add(50 * time.Millisecond)
// ctx, cancel := context.WithDeadline(context.Background(), d)
// // 尽管ctx会过期，但再任何情况下调用它的cancel函数都是很好的实践。
// // 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
// defer cancel()

// select {
// case <-time.After(1 * time.Second):
// 	fmt.Println("overslept")
// case <-ctx.Done():
// 	fmt.Println("err:", ctx.Err())
// }
// ctx, cancel := context.WithCancel(context.Background())
// defer cancel() // 当我们取完需要的整数后调用cancel

// for n := range gen(ctx) {
// 	fmt.Println(n)
// 	if n == 5 {
// 		break
// 	}
// }
// }

// var wg sync.WaitGroup

// // var exit bool

// // 官方版的方案
// func worker(ctx context.Context) {
// 	go worker2(ctx)
// LOOP:
// 	for {
// 		fmt.Println("worker")
// 		time.Sleep(time.Second)
// 		select {
// 		case <-ctx.Done(): // 等待上级通知
// 			break LOOP
// 		default:
// 		}
// 	}
// 	wg.Done()
// }
// func worker2(ctx context.Context) {
// LOOP:
// 	for {
// 		fmt.Println("worker2")
// 		time.Sleep(time.Second)
// 		select {
// 		case <-ctx.Done(): // 等待上级通知
// 			break LOOP
// 		default:
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	wg.Add(1)
// 	go worker(ctx)
// 	time.Sleep(time.Second * 3)
// 	cancel() // 通知子goroutine结束
// 	wg.Wait()
// 	fmt.Println("over")
// }

// // 管道方式存在的问题：
// // 1.使用全局变量在跨包调用时不容易实现规范和统一， 需要维护一个共用的channel

// func worker(exitChan chan struct{}) {
// LOOP:
// 	for {
// 		fmt.Println("worker")
// 		time.Sleep(time.Second)
// 		select {
// 		case <-exitChan: // 等待接收上级通知
// 			break LOOP
// 		default:

// 		}
// 	}
// 	wg.Done()
// }

// // 全局变量方式存在的问题：
// // 1.使用全局变量在跨包调用时不容易统一
// // 2. 如果worker中再启动goroutine, 就不太好控制了
// func worker() {
// 	for {
// 		fmt.Println("worker")
// 		time.Sleep(time.Second)
// 		if exit {
// 			break
// 		}
// 	}
// 	wg.Done()
// }

// // 初始的例子
// func worker() {
// 	for {
// 		fmt.Println("worker")
// 		time.Sleep(time.Second)
// 	}
// 	// 如何接收外部命令实现退出
// 	wg.Done()
// }

// func main() {
// 	var exitChan = make(chan struct{})
// 	wg.Add(1)
// 	go worker(exitChan)
// 	// 如何优雅的实现结束子goroutine
// 	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
// 	//exit = true                 // 修改全局变量实现子goroutine的退出
// 	exitChan <- struct{}{} // 给子goroutine发送退出信号
// 	close(exitChan)
// 	wg.Wait()
// 	fmt.Println("over")
// }
