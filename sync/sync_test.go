package sync

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func example() {
	// 增加锁
	mutex := sync.Mutex{}
	group := sync.WaitGroup{}
	count := 0
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			// 加锁
			mutex.Lock()
			for i := 0; i < 100; i++ {
				count++
			}
			mutex.Unlock()
			group.Done()
		}()
	}

	group.Wait()
	fmt.Printf("100 goroutine count 10 times  is %v times\n", count)
}

func atomicExam() {
	group := sync.WaitGroup{}
	var count int32
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				atomic.AddInt32(&count, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Printf("100 goroutine count 1000 times  is %v times", atomic.LoadInt32(&count))
}

func TestMutexExample(t *testing.T) {
	example()
}

func TestAtomicExample(t *testing.T) {
	atomicExam()
}

func TestChannel(t *testing.T) {
	ints := make(chan int, 1)
	go func(ints chan int) {
		fmt.Println("goroutline启动了")
		ints <- -1
	}(ints)

	select {
	case <-ints:
		fmt.Print("gorouline exits")
	}
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	i := 1
	once.Do(func() {
		i++
	})

	once.Do(func() {
		i++
	})

	fmt.Println("current value is ", i)
}

func TestCtx(t *testing.T) {
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)

	go func() {
		time.Sleep(2 * time.Second)
		cancelFunc()
	}()
	// defer cancelFunc()

	slowOperation(timeoutCtx)

	fmt.Println("都停止了")

}

func slowOperation(ctx context.Context) {
	go func() {
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s 肆意打印: %d\n", time.Now(), i)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("deadline已到")
	}
}

func TestCondMutex(t *testing.T) {
	signal := 1
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	go func() {
		cond.L.Lock()
		for signal == 1 {
			fmt.Println("线程开始等待")
			cond.Wait()
			fmt.Println("线程等待结束")
		}

		cond.L.Unlock()
	}()

	time.Sleep(2 * time.Second)
	cond.Broadcast()
	signal = 0
	time.Sleep(2 * time.Second)
}
