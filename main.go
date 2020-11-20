package main

import (
	"clipper_server/models/resp"
	"clipper_server/server/handler"
	"clipper_server/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	service.New()

	router := gin.Default()
	router.NoRoute(NoRoute)

	router.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello, friend")
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/clipboard", handler.GetClipboard)
		v1.POST("/clipboard", handler.CreateClipboardMessage)
	}

	router.Run(":8614")
}

func NoRoute(c *gin.Context) {
	c.JSON(200, &resp.Response{
		Code:    100404,
		Message: "没有找到路由",
		Data:    nil,
	})
}
