package go_advance

import (
	"encoding/json"
	"fmt"
)

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Width int64
	Long  int64
}

func (r *Rectangle) Area() {
	fmt.Println(fmt.Sprintf("Rectangle area is %v", r.Long*r.Width))
}

func (r *Rectangle) Perimeter() {
	fmt.Println(fmt.Sprintf("Rectangle Perimeter is %v", 2*(r.Long+r.Width)))
}

type Circle struct {
	Radius int64
}

func (r *Circle) Area() {
	fmt.Println(fmt.Sprintf("Circle area is %v", 3.14*float64(r.Radius)*float64(r.Radius)))
}

func (r *Circle) Perimeter() {
	fmt.Println(fmt.Sprintf("Circle Perimeter is %v", 3.14*2.0*float64(r.Radius)))
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	Age  int64
}

type Employee struct {
	*Person
	EmployeeID int64
}

func (e *Employee) PrintInfo() {
	marshal, _ := json.Marshal(e)
	fmt.Println(string(marshal))
}
