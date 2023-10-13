
```bash
cyd:003-tool-cobra cyd$ go run main.go word -s=blueegg -m=1
2023/10/13 09:41:54 输出结果: BLUEEGG
```


## 时间结果跟预计不一致

```bash
cyd:003-tool-cobra cyd$ go run main.go time calc -c="2029-09-04 12:02:33" -d=5m
2023/10/13 13:50:21 输出结果: 1970-01-01 08:05, 300
cyd:003-tool-cobra cyd$ go run main.go time calc -c="2029-09-04 12:02:33" -d=-2h
2023/10/13 13:50:54 输出结果: 1970-01-01 06:00, -7200
```