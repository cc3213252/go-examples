# 生成一个package，并用它

```bash
go install example.com/user/hello
```
这个命令会先build，然后安装

如果设置了GOBIN，二进制文件会装到这个目录，否则会装到GOPATH指定目录下的bin子目录中  

## 设置GOBIN

```bash
go env -w GOBIN=/somewhere/else/bin
```

## 解除GOBIN

```bash
go env -u GOBIN
```

## 让程序加入系统目录

```bash
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
```

## go build

go build不会产生输出文件，会保存编译包在本地cache中

## 碰到问题

[ok]package包的引入IDE解析不出来
加入一个外部包，再执行一下go mod tidy即可

## 删除所有已经下载的

```bash
go clean -modcache
```

## 测试

必须以_test.go命名文件，函数是Test***开头，用func(t *testing.T)作为入参