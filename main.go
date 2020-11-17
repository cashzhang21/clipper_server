package main

import (
	"clipper_server/model"
	"clipper_server/server/handler"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
)

const (
	dbUser	string = "root"
	dbPassword	string = "h9fjj8fKEL6w6PMY@"
	dbHost		string = "127.0.0.1"
	dbPort		int	   = 33060
	dbName		string = "clipperdb"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

func main() {

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	d.AutoMigrate()

	router := gin.Default()
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
