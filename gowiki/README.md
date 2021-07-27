## 测试

```bash
go build wiki.go
./wiki
```

## net/http简单实用

ListenAndServe这个函数正常会阻塞，不阻塞了返回错误  
![net例子](https://github.com/cc3213252/go-examples/raw/master/img/net-http-example.png)

## 升级后wiki

http://localhost:8080/view/TestPage