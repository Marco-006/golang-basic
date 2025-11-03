package main

import "fmt"

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
func main() {
	p := Person{Name: "xxx", Age: 66}
	// Q： 实例化对象的时候， 要么全写名，要么全按位，混用不允许
	e := Employee{p, "123"}

	e.PrintInfo()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Println("employee name is: ", e.Name, "age is: ", e.Age, "id is :", e.EmployeeID)
}
