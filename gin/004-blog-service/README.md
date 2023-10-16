## 创建项目

```bash
go mod init blueegg/blog-service
go get -u github.com/gin-gonic/gin@v1.6.3
```

## 配置文件

```bash
go get -u github.com/spf13/viper@v1.4.0
```

## ORM

```bash
go get -u github.com/jinzhu/gorm@v1.9.12
```

## 日志组件

```bash
go get -u gopkg.in/natefinch/lumberjack.v2
```

## 测试响应结果

curl -v http://127.0.0.1:8000/api/v1/articles/1

## 问题

【fixed】model.go 中需引入 _ "github.com/go-sql-driver/mysql" ，否则会报 unknow driver mysql