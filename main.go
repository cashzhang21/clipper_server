package main

import (
	"clipper_server/model"
	"clipper_server/server/handler"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.NoRoute(NoRoute)

	router.Use(static.Serve("/", static.LocalFile("./static", false)))
	router.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello, friend")
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/clipboard", handler.GetClipboard)
		v1.POST("/clipboard", handler.PostClipboard)
	}

	router.Run(":8614")
}

func NoRoute(c *gin.Context) {
	c.JSON(200, &model.Response{
		Code:    1040404,
		Message: "没有找到路由",
		Data:    nil,
	})
}
