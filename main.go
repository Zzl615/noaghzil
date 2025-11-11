package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // 静态文件映射
    router.Static("/static", "./static")

    // 首页路由
    router.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")

    })

    // 监听 80 端口
    err := router.Run(":80")
    if err != nil {
        panic(err)
    }
}

