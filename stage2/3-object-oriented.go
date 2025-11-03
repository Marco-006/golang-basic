package main

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。


package main

import (
	"fmt"
	"math"
)

// 1. Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Rectangle 结构体
type Rectangle struct {
	Width, Height float64
}

// Rectangle 实现 Shape 接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 3. Circle 结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main3() {
	// 4. 创建实例并通过 Shape 接口调用
	var s Shape

	s = Rectangle{Width: 3, Height: 4}
	fmt.Println("Rectangle")
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	s = Circle{Radius: 5}
	fmt.Println("Circle")
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}



// func main() {
// 	c = &Circle{l: 2, w: 3}
// 	c = &Circle{l: 2, w: 3}
// }

// type Shape interface {
// 	Area(l int, w int) int 
// 	Perimeter(l int, w int) int 
// }


// // 实现Shape interface
// type Rectangle struct{
// 	Area(l int, w int) int 
// 	Perimeter(l int, w int) int 
// }

// func (r *Rectangle) Area (l int, w int) int {
// 	return l * w 
// }

// func (r *Rectangle) Perimeter (l int, w int) int {
// 	return (l + w) * 2  
// }

// // 实现Shape interface
// type Circle struct{
// 	Area(l int, w int) int 
// 	Perimeter(l int, w int) int 
// }

// func (c *Circle) Area (l int, w int) int {
// 	return l * w
// }

// func (r *Circle) Perimeter (l int, w int) int {
// 	return (l + w) * 2  
// }

// Q：分别写两个方法 绑定给两个对象吗？
// A：方法参数的绑定方式至少有两种，可以是方法的入参，也可以是绑定对象自己的参数


// Q：方法怎么绑定给对象？
// A：

// 如何把不同对象的，相同的逻辑抽象出来？

