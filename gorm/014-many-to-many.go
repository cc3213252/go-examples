package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

// 检索 User 列表并预加载 Language
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Languages").Find(&users).Error
	return users, err
}

// 检索 Language 列表并预加载 User
func GetAllLanguages(db *gorm.DB) ([]Language, error) {
	var languages []Language
	err := db.Model(&Language{}).Preload("Users").Find(&languages).Error
	return languages, err
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{}, &Language{})

	lang1 := Language{Name: "english"}
	lang2 := Language{Name: "japan"}

	user := User{
		Name:      "yudan",
		Languages: []*Language{&lang1, &lang2},
	}
	db.Create(&user)
	db.Save(&user)

	langs, _ := GetAllLanguages(db)
	for i, lang := range langs {
		fmt.Printf("result: 序号 %d 值 %s\n", i, lang.Name)
	}
}
