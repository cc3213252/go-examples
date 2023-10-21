package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// step 1 单独创建
	//db.AutoMigrate(&User{})
	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//
	//result := db.Create(&user) // 通过数据的指针来创建
	//
	//fmt.Println(user.ID)                     // 返回插入数据的主键
	//fmt.Println(result.Error)                // 返回 error
	//fmt.Println("num:", result.RowsAffected) // 返回插入记录的条数

	// step 2  批量创建
	//users := []User{
	//	{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
	//	{Name: "Jackson", Age: 19, Birthday: time.Now()},
	//}
	//result := db.Create(users)
	//fmt.Println("insert num:", result.RowsAffected)

	// step3  Map创建
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "yudan", "Age": 20,
	})
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "xiaohong", "Age": 17},
		{"Name": "xiaolan", "Age": 34},
	})
}
