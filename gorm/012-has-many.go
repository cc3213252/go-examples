package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

// 检索用户列表并预加载信用卡
func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
	return users, err
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.AutoMigrate(&User{}, &CreditCard{})
	//
	//user := User{
	//	Name: "yudan",
	//	CreditCards: []CreditCard{
	//		{Number: "222222"},
	//		{Number: "333333"},
	//	},
	//}
	//db.Create(&user)
	//db.Save(&user)

	users, _ := GetAll(db)
	for i, user := range users {
		fmt.Printf("result: 序号 %d 值 %s\n", i, user.CreditCards)
	}
}

/*
mysql> show columns from users;
+------------+-----------------+------+-----+---------+----------------+
| Field      | Type            | Null | Key | Default | Extra          |
+------------+-----------------+------+-----+---------+----------------+
| id         | bigint unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | datetime(3)     | YES  |     | NULL    |                |
| updated_at | datetime(3)     | YES  |     | NULL    |                |
| deleted_at | datetime(3)     | YES  | MUL | NULL    |                |
| name       | longtext        | YES  |     | NULL    |                |
+------------+-----------------+------+-----+---------+----------------+
5 rows in set (0.00 sec)

mysql> show columns from credit_cards;
+------------+-----------------+------+-----+---------+----------------+
| Field      | Type            | Null | Key | Default | Extra          |
+------------+-----------------+------+-----+---------+----------------+
| id         | bigint unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | datetime(3)     | YES  |     | NULL    |                |
| updated_at | datetime(3)     | YES  |     | NULL    |                |
| deleted_at | datetime(3)     | YES  | MUL | NULL    |                |
| number     | longtext        | YES  |     | NULL    |                |
| user_id    | bigint unsigned | YES  | MUL | NULL    |                |
+------------+-----------------+------+-----+---------+----------------+
6 rows in set (0.00 sec)

*/
