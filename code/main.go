/**
@author: yudan.chen
@date: 2023/9/22
**/
package main
// 导入gin包
import "github.com/gin-gonic/gin"

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := gin.Default()

	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	r.GET("/ping", func(c *gin.Context) {
		// 通过请求上下文对象Context, 直接往客户端返回一个json
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}