
go mod init blueegg.com/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

go get -u gorm.io/driver/mysql

CREATE DATABASE IF NOT EXISTS gorm_test DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

## 问题

004 生成带外键的表格时报错