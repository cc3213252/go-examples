package main

import (
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

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{}, &CreditCard2{})

	user := User{
		Name:        "yudan",
		CreditCard2: CreditCard2{Number: "11222111"},
	}
	db.Create(&user)
	db.Save(&user)

}

/*
one to one关系
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

mysql> show columns from credit_card2;
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
