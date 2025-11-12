package task6gormadvance

import (
	"fmt"

	"gorm.io/gorm"
)

// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

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
	var userID = 3
	// user := User{ID: 2}
	// db.Debug().Table("User").Where("id = ?", userID).Find(&user)
	// gorm.G[User](db).First(2)

	var user User
	db.Preload("Post.Comment").First(&user, userID)
	fmt.Println("user 表", user.ID)

	// range 使用，用一个参数的时候需要关注原来的数据结构，
	// for k := range user.Post {
	// 	fmt.Println("用户发表的文章ID:", user.Post[k])
	// }

	for _, post := range user.Post {
		fmt.Println("用户发表的文章:", post.ID)
		for _, comment := range post.Comment {
			fmt.Println("文章对应的评论:", comment.Cotent)

		}
	}

	// 需求2
	// 评论数量最多的文章信息
	var mostCommentedPostV1 Post
	db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Find(&mostCommentedPostV1)

	fmt.Println("评论数量最多的文章信息 mostCommentedPostV1:", mostCommentedPostV1)

	var mostCommentedPostV2 Post
	db.Raw(`
		SELECT posts.* 
		FROM posts 
		JOIN (
			SELECT post_id, COUNT(*) as comment_count 
			FROM comments 
			GROUP BY post_id 
			ORDER BY comment_count DESC 
			LIMIT 1
		) as top_post ON posts.id = top_post.post_id
	`).Scan(&mostCommentedPostV2)
	// SELECT *  FROM comments GROUP BY post_id ORDER BY COUNT(*) desc limit 1

	fmt.Println("评论数量最多的文章信息 mostCommentedPostV2:", mostCommentedPostV2)

	// var result Post
	// gorm.G[Post](db).Raw("SELECT *  FROM comments GROUP BY post_id ORDER BY COUNT(*) desc limit 1").Scan(&result)
	// result, err := gorm.G[Post](db).Raw("SELECT *  FROM comments GROUP BY post_id ORDER BY COUNT(*) desc limit 1").Find(&result)
	// if err != nil {
	// 	fmt.Println("error", err)
	// }

	// db.Model(&Post{}).Select()

	// fmt.Println("query result ", result)

	// 查询语句，抄一个能用就行；
	// rang遍历

	// fmt.Println("用户 发表的所有文章： ", posts)

	// for _, v := range posts {

	// }

	// db.Model("Comment").Where()

	// 关联删除

}
