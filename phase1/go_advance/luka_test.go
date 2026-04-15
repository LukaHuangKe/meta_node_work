package go_advance

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutDemo(t *testing.T) {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case msg := <-ch:
		fmt.Println("收到:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("超时了") // 1秒后输出这个
	}
	// 因为发送需要2秒，但超时是1秒，所以会输出"超时了"
}

func TestLoopSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- fmt.Sprintf("msg-%d", i)
			time.Sleep(150 * time.Millisecond)
		}
		close(ch2)
	}()

	// 持续监听，直到所有channel都关闭
	for {
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil // 关闭的channel设为nil，select会忽略它
				continue
			}
			fmt.Println("ch1:", val)
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("ch2:", msg)
		default:
			// 如果没有数据，可以做其他事情
			//fmt.Println("default")
			if ch1 == nil && ch2 == nil {
				fmt.Println("所有channel已关闭")
				return
			}
		}
	}
}

func TestConsumerProducer(t *testing.T) {
	ch := make(chan int, 5)

	go producer(ch)

	// 多个消费者
	for i := 0; i < 3; i++ {
		go consumer(ch, i)
	}

	time.Sleep(2 * time.Second)
}

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, id int) {
	for value := range ch {
		fmt.Printf("Consumer %d received: %d\n", id, value)
		time.Sleep(50 * time.Millisecond)
	}
}
