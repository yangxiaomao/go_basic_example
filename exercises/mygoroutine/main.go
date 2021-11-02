package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3.主goroutine从resultChan取出结果并打印到终端输出

*/
type Job struct {
	X int64
}

// Result ...
type Result struct {
	Job    *Job
	Result int64
}

var jobChan = make(chan *Job, 10)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func maozong(mz chan<- *Job) {
	defer wg.Done()
	// 循环生成int64类型的随机数，发送到jobChan
	for {
		value := rand.Int63()
		newJob := &Job{
			X: value,
		}
		mz <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func xiaomao(mz <-chan *Job, resultChan chan<- *Result) {
	defer wg.Done()
	// 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan中
	for {
		job := <-mz
		sum := int64(0)
		n := job.X
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &Result{
			Job:    job,
			Result: sum,
		}
		resultChan <- newResult
	}
}
func main() {
	wg.Add(1)
	go maozong(jobChan)
	// 开启24个goroutine执行小毛
	wg.Add(24)
	for i := 1; i < 25; i++ {
		go xiaomao(jobChan, resultChan)
	}
	// 主goroutine从resultChan取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.Job.X, result.Result)
	}
	wg.Wait()
}
