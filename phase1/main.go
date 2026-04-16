package main

import (
	"fmt"
	"phase1/go_basic"
)

func main() {
	//goBasic()
	goAdvance()
}

func goAdvance() {
	// 1.
	//fmt.Println(go_advance.Add10(proto.Int64(10)))

	// 2.
	//fmt.Println(go_advance.SliceMul2([]int64{1, 2, 3}))

	// 3.
	//go_advance.GoroutineOddEvenNum()

	// 4.
	//go_advance.JobTicker()

	// 5.
	//rec := &go_advance.Rectangle{
	//	Width: 10,
	//	Long:  20,
	//}
	//rec.Area()
	//rec.Perimeter()
	//
	//circle := &go_advance.Circle{
	//	Radius: 10,
	//}
	//circle.Perimeter()
	//circle.Area()

	// 6.
	//employee := &go_advance.Employee{
	//	Person: &go_advance.Person{
	//		Name: "hk test",
	//		Age:  18,
	//	},
	//	EmployeeID: 123,
	//}
	//employee.PrintInfo()

	// 7.
	//go_advance.ChannelWork1()

	// 8.
	//go_advance.ChannelWork2()

	// 9.
	//go_advance.MutexWork1()

	// 10.
	//go_advance.MutexWork2()
}

func goBasic() {
	// 只出现一次的数字
	fmt.Println(go_basic.SingleNumber([]int{2, 2, 1}))
	// 回文数
	fmt.Println(go_basic.IsPalindrome(123321))
}
