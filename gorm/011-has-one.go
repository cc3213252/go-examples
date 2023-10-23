package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	CreditCard2 CreditCard2
}

type CreditCard2 struct {
	gorm.Model
	Number string
	UserID uint
}

// 检索用户列表并预加载信用卡
func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("CreditCard2").Find(&users).Error
	return users, err
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	users, _ := GetAll(db)
	for i, user := range users {
		fmt.Printf("result: 序号 %d 值 %s\n", i, user.Name)
	}
}
