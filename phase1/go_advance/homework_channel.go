package go_advance

import (
	"fmt"
	"sync"
	"time"
)

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func ChannelWork1() {
	ch := make(chan int)

	go func() {
		for i := range ch {
			fmt.Println(fmt.Sprintf("打印数字 %d", i))
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	time.Sleep(5 * time.Second)
	close(ch)
}

func ChannelWork2() {
	length := 10

	ch := make(chan int, length)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println(fmt.Sprintf("打印数字 %d", i))
		}
	}()

	go func() {
		defer wg.Done()
		defer close(ch) //关闭channel不影响消费
		for i := 0; i < length; i++ {
			ch <- i
		}

	}()

	wg.Wait()
}
