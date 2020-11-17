package handler

import (
	"clipper_server/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"time"
)

func PostClipboard(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var bodyData map[string]interface{}
	if err := json.Unmarshal(body, &bodyData); err == nil {
		clipboardText := bodyData["text"]
		if clipboardText == "" || clipboardText == nil {
			c.JSON(200, &model.Response{
				Code:    1060003,
				Message: "参数错误",
				Data:    nil,
			})
			return
		}
		r := &model.Response{
			Code:    0,
			Message: "success",
			Data:    nil,
		}
		c.JSON(200, r)
		currentTime := time.Now()
		fmt.Println(currentTime.Format("2006-01-02 15:04:05"), ": ", clipboardText)

		db, err := sql.Open("mysql", "root:/clipper")
		if err != nil {
			panic(err)
		}
		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		return
	}
	c.JSON(200, &model.Response{
		Code:    1060003,
		Message: "参数错误",
		Data:    nil,
	})
}

func GetClipboard(c *gin.Context) {

}
