# go 语言编程之旅

https://golang2.eddycjy.com/posts/ch1/01-simple-flag/

## flag 命令行输入

接受命令行入参
go run 001-tool-flag.go -n=煎鱼

## cobra 单词转换小项目

生成go.mod
```bash
go mod init blueegg.com/gin
```
增加cobra库
```bash
go get -u github.com/spf13/cobra@v1.0.0
```