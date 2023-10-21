package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type UserSoft struct {
	gorm.Model
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.AutoMigrate(&UserSoft{})
	//db.Create(&UserSoft{Name: "kk", Age: 12, Birthday: time.Now()})
	//db.Delete(&UserSoft{}, 1)

	var user UserSoft
	db.First(&user)
	fmt.Println("result:", db.RowsAffected)

	tx := db.Unscoped().Where("age = 12").Find(&user)
	fmt.Println("after:", tx.RowsAffected)
}
