package go_advance

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func MutexWork1() {
	var mu sync.Mutex
	num := 0

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				mu.Lock()
				num++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println(fmt.Sprintf("num: %d", num))
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
// MutexWork2 使用原子操作实现无锁计数器
func MutexWork2() {
	var num int32

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				// 使用原子操作递增计数器
				atomic.AddInt32(&num, 1)
			}
		}()
	}

	wg.Wait()

	// 使用原子操作读取最终值
	fmt.Println(fmt.Sprintf("num: %d", atomic.LoadInt32(&num)))
}
