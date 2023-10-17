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

## 验证器

```bash
go get -u github.com/go-playground/validator/v10
```

## 测试响应结果

curl -v http://127.0.0.1:8000/api/v1/articles/1

## 参数校验

cyd:~ cyd$ curl -X GET http://127.0.0.1:8000/api/v1/tags\?state\=6
{"code":10000001,"details":["State必须是[0 1]中的一个"],"msg":"入参错误"}cyd:~ cyd$ curl -X GET http://127.0.0.1:8000/api/v1/tags\?state\=1
{}cyd:~ cyd$

## swagger

```bash
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0 
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

## 问题

【fixed】model.go 中需引入 _ "github.com/go-sql-driver/mysql" ，否则会报 unknow driver mysql  

【fixed】swag报错，not spec，需要在router.go中增加_ "blueegg/blog-service/docs"

通过swagger提交POST请求报校验错误，实际postman调是好的

