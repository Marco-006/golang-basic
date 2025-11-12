package task5gormadvance

import (
	"fmt"

	"gorm.io/gorm"
)

// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
// Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	gorm.Model
	Post []Post
}

type Post struct {
	gorm.Model
	UserID  uint
	Comment []Comment
}

type Comment struct {
	gorm.Model
	PostID uint
	Cotent string
}

func Run(db *gorm.DB) {
	// init table structure
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Post{})
	// db.AutoMigrate(&Comment{})

	// 创建
	// 创建可以关联吗

	// 插入
	// 一对一的关联
	// var user1 User
	// db.Debug().Create(&user1)
	// fmt.Println("新创建user，userID： ", user1.ID)

	// post1 := Post{UserID: user1.ID}
	// db.Debug().Create(&post1)

	// comment1 := Comment{PostID: post1.ID, Cotent: "66666"}
	// db.Debug().Create(&comment1)

	var user1 User
	db.Debug().Create(&user1)
	fmt.Println("新创建user，userID： ", user1.ID)

	post1 := Post{UserID: user1.ID}
	db.Debug().Create(&post1)
	fmt.Println("新创建post，postID： ", post1.ID)

	post12 := Post{UserID: user1.ID}
	db.Debug().Create(&post12)
	fmt.Println("新创建post，postID： ", post12.ID)

	comment1 := Comment{PostID: post1.ID, Cotent: "66666"}
	db.Debug().Create(&comment1)
	fmt.Println("新创建comment，commentID： ", comment1.ID)

	comment12 := Comment{PostID: post1.ID, Cotent: "7777"}
	db.Debug().Create(&comment12)
	fmt.Println("新创建comment，commentID： ", comment12.ID)
	// 一对一的修改

	// 关联删除

}
