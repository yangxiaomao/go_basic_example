package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 我们填写一个示例来比较下互斥锁和原子操作的性能。
type Counter interface {
	Inc()
	Load() int64
}

// 普通版
type CommonCounter struct {
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

// 互斥锁版
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	c1 := CommonCounter{} // 非并发安全
	test(c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	test(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	test(&c3)
}

// // sync.Map{}
// var m = sync.Map{}

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			m.Store(key, n)
// 			value, _ := m.Load(key)
// 			fmt.Printf("k=:%v,v:=%v\n", key, value)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

// var m = make(map[string]int)

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int) {
// 	m[key] = value
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			set(key, n)
// 			fmt.Printf("k=:%v, v:=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

// // 我们利用sync.WaitGroup记录并发声明周期
// var wg sync.WaitGroup

// func hello() {
// 	defer wg.Done()
// 	fmt.Println("Hello Goroutine!")
// }

// func main() {
// 	wg.Add(1)
// 	go hello() // 启动另外一个goroutine去执行hello函数
// 	fmt.Println("main goroutine done!")
// 	wg.Wait()
// }

// // 读写锁示例
// var (
// 	x      int64
// 	wg     sync.WaitGroup
// 	lock   sync.Mutex
// 	rwlock sync.RWMutex
// )

// // 写
// func write() {
// 	// lock.Lock   // 加互斥锁
// 	rwlock.Lock() // 加写锁
// 	x = x + 1
// 	time.Sleep(10 * time.Millisecond) //假设写操作耗时10毫秒
// 	rwlock.Unlock()                   // 解写锁
// 	// lock.Unlock() // 解互斥锁
// 	wg.Done()
// }

// // 读
// func read() {
// 	// lock.Lock()   // 加互斥锁
// 	rwlock.RLock()               // 加读锁
// 	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
// 	rwlock.RUnlock()             //解读锁
// 	// lock.Unlock()  // 解互斥锁
// 	wg.Done()
// }

// func main() {
// 	start := time.Now()
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go write()
// 	}

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go read()
// 	}

// 	wg.Wait()
// 	end := time.Now()
// 	fmt.Println(end.Sub(start))
// }

// // 互斥锁
// var x int64
// var wg sync.WaitGroup
// var lock sync.Mutex

// func add() {
// 	for i := 0; i < 5000; i++ {
// 		lock.Lock() //加锁
// 		x = x + 1
// 		lock.Unlock() // 解锁

// 	}
// 	wg.Done()
// }

// func main() {
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Println(x)
// }

// var x int64
// var wg sync.WaitGroup

// func add() {
// 	for i := 0; i < 5000; i++ {
// 		x = x + 1
// 	}
// 	wg.Done()
// }

// func main() {
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Println(x)
// }

// func main() {
// 	ch := make(chan int, 1)
// 	for i := 0; i < 10; i++ {
// 		select {
// 		case x := <-ch:
// 			fmt.Println(x)
// 		case ch <- i:
// 			fmt.Println(123456)
// 		}
// 	}
// }

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Printf("worker:%d start job:%d\n", id, j)
// 		time.Sleep(time.Second)
// 		fmt.Printf("worker:%d end job:%d\n", id, j)
// 		results <- j * 2
// 	}
// }

// func main() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)
// 	// 开启3个goroutine
// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}
// 	// 5个任务
// 	for j := 1; j <= 5; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
// 	//输出结果
// 	for a := 1; a <= 5; a++ {
// 		<-results
// 	}
// }

// func counter(out chan<- int) {
// 	for i := 0; i < 100; i++ {
// 		out <- i
// 	}
// 	close(out)
// }

// func squarer(out chan<- int, in <-chan int) {
// 	for i := range in {
// 		out <- i * i
// 	}
// 	close(out)
// }

// func printer(in <-chan int) {
// 	for i := range in {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	go counter(ch1)
// 	go squarer(ch2, ch1)
// 	printer(ch2)
// }

// // channel 练习
// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	// 开启goroutine将0~100的数发送到ch1中
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			ch1 <- i
// 		}
// 		close(ch1)
// 	}()
// 	// 开启goroutine从ch1中接收值， 并将该值的平方发送到ch2中
// 	go func() {
// 		for {
// 			i, ok := <-ch1 // 通道关闭后再取值ok=false
// 			if !ok {
// 				break
// 			}
// 			ch2 <- i * i
// 		}
// 		close(ch2)
// 	}()
// 	// 在主 goroutine中从ch2中接收值打印
// 	for i := range ch2 {
// 		// 通道关闭后会退出for range 循环
// 		fmt.Println(i)
// 	}
// }

// func recv(c chan int) {
// 	ret := <- c
// 	fmt.Println("接收成功", ret)

// }

// func main() {
// 	ch := make(chan int)
// 	go recv(ch) // 启动goroutine从通道接收值
// 	ch <- 10
// 	fmt.Println("发送成功")
// }

// var ch chan int

// func main() {
// 	ch = make(chan int, 1)
// 	ch <- 10 //把10发送到ch中
// 	close(ch)
// 	x := <-ch // 从ch中接收值并赋值给变量X
// 	//<-ch      // 从ch中接收值，忽略结果
// 	y := <-ch // 当通道关闭后，还可以取值，当值被取完后，则为0

// 	fmt.Println(ch)
// 	fmt.Println(x)
// 	fmt.Println(y)
// }

// func a() {
// 	for i := 1; i < 20; i++ {
// 		fmt.Println("A:", i)
// 	}
// }

// func b() {
// 	for i := 1; i < 20; i++ {
// 		fmt.Println("B:", i)
// 	}
// }

// func main() {
// 	runtime.GOMAXPROCS(2)
// 	fmt.Println(runtime.NumCPU()) // 查看CPU核数
// 	go a()
// 	go b()
// 	time.Sleep(time.Second)
// }

// var wg sync.WaitGroup

// func hello(i int) {
// 	defer wg.Done() // goroutine结束就登记-1
// 	fmt.Println("Hello Goroutine!", i)
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go hello(i) // 启动另外一个goroutine去执行hello函数
// 	}
// 	wg.Wait() // 等待所有登记的goroutine都结束

// 	fmt.Println("main goroutine done!")

// }
