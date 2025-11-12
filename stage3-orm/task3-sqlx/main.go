package task3sqlx

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

// type Employee struct {
// 	ID         int     `db:"id"`
// 	Name       string  `db:"name"`
// 	Department string  `db:"department"`
// 	Salary     float64 `db:"salary"`
// }

// func QueryEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
// 	var employees []Employee
// 	query := "SELECT id, name, department, salary FROM employees WHERE department = ?"
// 	err := db.Select(&employees, query, department)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return employees, nil
// }

// func QueryHighestSalaryEmployee(db *sqlx.DB) (*Employee, error) {
// 	var employee Employee
// 	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
// 	err := db.Get(&employee, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &employee, nil
// }

// func Run(db *sqlx.DB) {

// 	employees, err := QueryEmployeesByDepartment(db, "B")
// 	if err != nil {
// 		fmt.Println("Error querying employees by department:", err)
// 		return
// 	}
// 	fmt.Println("Employees in 技术部:", employees)

// 	highestSalaryEmployee, err := QueryHighestSalaryEmployee(db)
// 	if err != nil {
// 		fmt.Println("Error querying highest salary employee:", err)
// 		return
// 	}
// 	fmt.Println("Highest salary employee:", highestSalaryEmployee)

// }

type Employee struct {
	// 测试GORM是不是可用
	ID         int `db:"id"`
	Name       string
	Department string
	Salary     int64
}

func Run(db *sqlx.DB) {
	// dbGorm.AutoMigrate(&Employee{})

	// dbGorm.Create(&Employee{Name: "a1", Department: "A", Salary: 23})
	// dbGorm.Create(&Employee{Name: "a2", Department: "A", Salary: 34})
	// dbGorm.Create(&Employee{Name: "b1", Department: "B", Salary: 78})
	// dbGorm.Create(&Employee{Name: "b2", Department: "B", Salary: 89})
	// dbGorm.Create(&Employee{Name: "c1", Department: "C", Salary: 66})
	// dbGorm.Create(&Employee{Name: "c2", Department: "C", Salary: 55})

	var employees []Employee
	// 为什么使用 * 不能查询出来结果？
	db.Select(&employees, "select id, name, department, salary  from employees where department = ?", "B")
	fmt.Println("所有部门为技术部的员工信息: ", employees)

	var e Employee
	db.Get(&e, "SELECT id, name, department, salary  from employees order by salary desc limit 1")
	fmt.Println("工资最高的员工ID为: ", e.ID)
}
