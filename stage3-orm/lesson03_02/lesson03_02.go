package lesson03_02

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// 这个一对多的关系中，User是“一”，CreditCard是“多”；  并没有表现在User结构体中的字段上，而是通过CreditCard结构体中的UserID字段来体现的
	CreditCards []CreditCard
}

// type User struct {
// 	gorm.Model
// 	Name       string     `gorm:"unique;index;size:50"`
// 	CreditCard CreditCard `gorm:"foreignKey:UserName;references:Name"`
// }

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func Run(db *gorm.DB) {
	// db.Debug().AutoMigrate(&User{})
	// db.Debug().AutoMigrate(&CreditCard{})

	// user := User{}
	// db.Debug().Create(&user)

	// card := CreditCard{Number: "1111", UserID: 1}
	// db.Debug().Create(&card)

	// card2 := CreditCard{Number: "2222", UserID: 1}
	// db.Debug().Create(&card2)

	user := User{}
	db.Debug().Preload("CreditCards").First(&user, 1)
	fmt.Println(user)
}
