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

## 框架不足

错误是返回不是抛异常方式，后果就是未能捕获的异常就不知道是啥问题后台一律显示失败

## 问题

【fixed】model.go 中需引入 _ "github.com/go-sql-driver/mysql" ，否则会报 unknow driver mysql  

【fixed】swag报错，not spec，需要在router.go中增加_ "blueegg/blog-service/docs"

【fixed】通过swagger提交POST请求报校验错误，实际postman调是好的，swagger参数类型不对导致

【fixed】更新tag老失败，最后发现是名字有约束，最少3个字符，而框架没有抛异常原因出来

更新tag时，ModifiedBy必须这样写，否则出错，不知为啥