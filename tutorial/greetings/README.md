* 对外发布的变量首字符大写
* go中%v表示相应值的默认格式

```bash
go mod edit -replace example.com/hello=../hello
```
```bash
go mod tidy
```

## 测试用例

```bash
go test -v
```