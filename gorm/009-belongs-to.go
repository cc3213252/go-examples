package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User3 struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User3{}, &Company{})

	user := User3{
		Name:    "yudan",
		Company: Company{Name: "greatwall"},
	}
	db.Create(&user)
	db.Save(&user)

}

/*
mysql> show tables;
+---------------------+
| Tables_in_gorm_test |
+---------------------+
| companies           |
| user3               |
| user_softs          |
| users               |
+---------------------+
4 rows in set (0.00 sec)

mysql> show columns from user3;
+------------+-----------------+------+-----+---------+----------------+
| Field      | Type            | Null | Key | Default | Extra          |
+------------+-----------------+------+-----+---------+----------------+
| id         | bigint unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | datetime(3)     | YES  |     | NULL    |                |
| updated_at | datetime(3)     | YES  |     | NULL    |                |
| deleted_at | datetime(3)     | YES  | MUL | NULL    |                |
| name       | longtext        | YES  |     | NULL    |                |
| company_id | bigint          | YES  | MUL | NULL    |                |
+------------+-----------------+------+-----+---------+----------------+
6 rows in set (0.00 sec)

mysql> show columns from companies;
+-------+----------+------+-----+---------+----------------+
| Field | Type     | Null | Key | Default | Extra          |
+-------+----------+------+-----+---------+----------------+
| id    | bigint   | NO   | PRI | NULL    | auto_increment |
| name  | longtext | YES  |     | NULL    |                |
+-------+----------+------+-----+---------+----------------+
2 rows in set (0.00 sec)
*/
