## 增加外部库

代码写完后，执行  
```bash
go get .
```
表示把依赖加到当前目录

## 增加条目

```bash
curl http://localhost:8080/albums \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": 49.99}'
```
## 查看条目列表

```bash
curl http://localhost:8080/albums \
  --header "Content-Type: application/json" \
  --request "GET"
```

## 查看单条

```bash
curl http://localhost:8080/albums/2
```