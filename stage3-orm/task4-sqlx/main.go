package task4sqlx

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 如何初始化表结构体
func Run(db *gorm.DB) {
	// 这里可以添加初始化代码，例如创建表结构等
	db.AutoMigrate(&Book{})

	db.Create(&Book{Title: "Book A", Author: "Author A", Price: 45.0})
	db.Create(&Book{Title: "Book B", Author: "Author B", Price: 55.0})
	db.Create(&Book{Title: "Book C", Author: "Author C", Price: 65.0})
}

func QueryExpensiveBooks(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price > ?"
	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, err
	}
	fmt.Println("size of books:", len(books))
	return books, nil
}
