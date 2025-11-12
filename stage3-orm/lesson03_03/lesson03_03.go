package lesson03_03

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Refer     uint       `gorm:"unique"`
	Languages []Language `gorm:"many2many:user_lang;foreignkey:Refer;joinForeignKey:UserReferID;References:Refer;joinReferences:LangReferID;"`
}

// foreignKey:ID           → 本模型用哪一列作为主键（用于中间表外键）。
// joinForeignKey:StudentID → 中间表里指向本模型的外键列名。
// references:ID           → 对方模型用哪一列被引用。
// joinReferences:CourseID → 中间表里指向对方模型的外键列名。

// type User struct {
// 	gorm.Model
// 	Languages []Language
// }

type Language struct {
	gorm.Model
	Name  string
	Refer uint `gorm:"unique"`
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Language{})

	// user := User{}
	// db.Create(&user)

	// language := Language{Name: "english"}
	// db.Create(&language)
	//
	// language2 := Language{Name: "chinese"}
	// db.Create(&language2)

	// user := User{Languages: []Language{{Name: "en1"}, {Name: "en2"}}}
	// db.Create(&user)

	// user := User{}
	// err := db.Preload("Languages").Find(&user, 3).Error
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user)
}
