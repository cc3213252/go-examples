
go mod init blueegg.com/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

go get -u gorm.io/driver/mysql

CREATE DATABASE IF NOT EXISTS gorm_test DEFAULT CHARSET utf8 COLLATE utf8_general_ci;