## go.mod file not found in current directory or any parent directory

```bash
go env -w GO111MODULE=on
go mod init code
go mod tidy
go run main.go
```

## 无法安装swag

查看到之前GOBIN环境变量配错，环境变量正确配置如下：
```bash
vim ~/.bash_profile
export GOROOT=/usr/local/go
export GOPATH=/Users/cyd/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
```

需要设置权限
```bash
cd /usr/local/
sudo chown -R cyd:staff go
```

再次安装即可
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```