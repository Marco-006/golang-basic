package task7hook

import (
	"fmt"

	"gorm.io/gorm"
)

// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

type User struct {
	gorm.Model
	Post      []Post
	PostCount int64
}

type Post struct {
	gorm.Model
	UserID        uint
	Comment       []Comment
	CommentCount  int64
	CommentStatus string
}

type Comment struct {
	gorm.Model
	PostID uint
	Cotent string
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var currentCount int64
	tx.Model(&Post{}).Where("user_id = ?", p.UserID).Select("count(*)").Count(&currentCount)
	fmt.Println("currentCount is :", currentCount)
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", currentCount)
	return
}

func (p *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var currentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", p.PostID).Select("count(*)").Count(&currentCount)
	tx.Model(&Post{}).Where("id = ?", p.PostID).Update("comment_count", currentCount)

	if currentCount == 0 {
		tx.Model(&Post{}).Where("id = ?", p.PostID).Update("comment_status", "无评论")
	}

	fmt.Println("after delete, comment_count is: ", currentCount)
	return
}

// comment_count
// comment_count

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

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

	comment12 := Comment{PostID: post12.ID, Cotent: "7777"}
	db.Debug().Create(&comment12)
	fmt.Println("新创建comment，commentID： ", comment12.ID)

	comment13 := Comment{PostID: post12.ID, Cotent: "888"}
	db.Debug().Create(&comment13)
	fmt.Println("新创建comment，commentID： ", comment13.ID)

	db.Debug().Delete(&comment13)
}
