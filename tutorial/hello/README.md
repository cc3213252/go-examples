正常应把greetings发布后才能使用，这里需要修改本地路径才能引用到自己的仓库
```bash
go mod edit -replace example.com/greetings=../greetings
```
再同步一下即可访问到
```bash
go mod tidy
```

require example.com/greetings v0.0.0-00010101000000-000000000000  
后面一串是伪版本号  

## 打包

```bash
go build
```
可以用以下命令查看发布目录
```bash
go list -f '{{.Target}}'
```

## 安装

GOBIN用来设置可执行文件的路径  
如设置：  
```bash
go env -w GOBIN=/Users/cyd/Desktop/tools/goinstall
```
那么当执行go install，会安装到这个目录下。  
如果再把安装目录加到PATH里面，那所有安装的程序可直接运行了  