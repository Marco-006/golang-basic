package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	task4sqlx "github.com/test/init_project/stage3-orm/task4-sqlx"
)

type Parent struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Child struct {
	Parent
	Age int
}

// func InitDB(dst ...interface{}) *gorm.DB {
// 	db, err := gorm.Open(mysql.Open("root:st123456@tcp(127.0.0.1:3306)/web3_learnning?charset=utf8mb4&parseTime=True&loc=Local"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.AutoMigrate(dst...)

// 	return db
// }

func main() {
	// gorm connection
	// gormDb, err := gorm.Open(mysql.Open("root:st123456@tcp(127.0.0.1:3306)/web3_learnning?charset=utf8mb4&parseTime=True&loc=Local"))
	// if err != nil {
	// 	panic(err)
	// }

	// sqlx connection :  need import _ "github.com/go-sql-driver/mysql";
	// 在 Go 语言中，_ "github.com/go-sql-driver/mysql" 表示导入一个包，但不直接使用它
	db, err := sqlx.Connect("mysql", "root:st123456@tcp(127.0.0.1:3306)/web3_learnning?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// lesson01.Run(db)
	// lesson02.Run(db)
	// lesson03.Run(db)
	// lesson03_02.Run(db)
	// lesson03_03.Run(db)
	// lesson03_04.Run(db)
	// lesson04.Run(db)

	// InitDB(&Parent{}, &Child{})

	// task1.Run(db)
	// task2.Run(db)
	// task3sqlx.Run(db)
	// task4sqlx.Run(gormDb)
	task4sqlx.QueryExpensiveBooks(db, 50)

	// task5gormadvance.Run(db)
	// task6gormadvance.Run(db)
	// task6gormadvance.Run(db)

	// task7hook.Run(db)
}
