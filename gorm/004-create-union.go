package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User2 struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// step 1 单独创建
	err := db.AutoMigrate(&User2{})
	if err != nil {
		return
	}
	db.Create(&User2{
		Name:       "jinzhu",
		CreditCard: CreditCard{Number: "411111111111"},
	})
	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

}
