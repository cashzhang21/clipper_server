package main

import (
	"clipper_server/model"
	"clipper_server/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.NoRoute(NoRoute)

	v1 := router.Group("/v1")
	{
		//v1.GET("/clipboard", handler.GetClipboard)
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
