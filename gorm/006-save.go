package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//type User struct {
//	ID       uint32    `gorm:"primary_key" json:"id"`
//	Name     string    `json:"name"`
//	Age      int       `json:"age"`
//	Birthday time.Time `json:"birthday"`
//}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var user User
	//var users []User

	user.Name = "fangliu"
	user.Age = 22
	user.Birthday = time.Now()
	db.Save(&user)
}
