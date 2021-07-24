## 设置GOPROXY

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## mod

初始化module（需要和别人分享自己写的module时，需要init）：
```bash
go mod init example.com/hello
```

拉取缺少的模块，移出不用的模块
```bash
go mod tidy
```

## 搜索外部库

pkg.go.dev

## 设置GOPROXY

系统的GOPROXY在Goland中会覆盖，所以每次创建新项目，都需要在Goland中设置一遍，如下：
![效果图](https://github.com/cc3213252/go-examples/raw/master/img/set-proxy.png)