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