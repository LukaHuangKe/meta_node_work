package go_advance

import (
	"fmt"
	"math/rand"
	"time"
)

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
func Add10(num *int64) int64 {
	return *num + 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func SliceMul2(nums []int64) []int64 {
	for i, _ := range nums {
		nums[i] = nums[i] * 2
	}
	return nums
}

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func GoroutineOddEvenNum() {
	go func() {
		for i := 1; i <= 100; i += 2 {
			fmt.Println(fmt.Sprintf("奇数协程打印：%d", i))
		}
	}()

	go func() {
		for i := 2; i <= 100; i += 2 {
			fmt.Println(fmt.Sprintf("偶数协程打印：%d", i))
		}
	}()
	time.Sleep(1 * time.Second)
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 貌似改成waitGroup的方式更合理
func JobTicker() {
	funcList := make([]func(num int64), 0)
	for i := 0; i < 10; i++ {
		funcList = append(funcList, func(num int64) {
			time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
			fmt.Println(fmt.Sprintf("this is NO:%d func", num))
		})
	}

	go func() {
		for i := 0; i < len(funcList); i += 2 {
			start := time.Now().UnixMilli()
			funcList[i](int64(i))
			end := time.Now().UnixMilli()
			fmt.Println(fmt.Sprintf("NO %d job exec %d ms", i, end-start))
		}
	}()

	go func() {
		for i := 1; i < len(funcList); i += 2 {
			start := time.Now().UnixMilli()
			funcList[i](int64(i))
			end := time.Now().UnixMilli()
			fmt.Println(fmt.Sprintf("NO %d job exec %d ms", i, end-start))
		}
	}()

	time.Sleep(3 * time.Second)
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
