package handler

import (
	"clipper_server/models/entity"
	"clipper_server/models/resp"
	"clipper_server/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

func CreateClipboardMessage(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var bodyData map[string]interface{}
	if err := json.Unmarshal(body, &bodyData); err == nil {
		clipboardText := bodyData["text"]
		if clipboardText == "" || clipboardText == nil {
			c.JSON(200, &resp.Response{
				Code:    100501,
				Message: "参数错误",
				Data:    nil,
			})
			return
		}

		currentTime := time.Now()
		fmt.Println(currentTime.Format("2006-01-02 15:04:05"), ": ", clipboardText)
		msg := entity.Message{
			Id:       0,
			CreateAt: currentTime,
			Text:     fmt.Sprintf("%s", clipboardText),
		}
		result, err := service.SERVICE.CreateClipboardMessage(&msg)
		if err != nil {
			c.JSON(200, &resp.Response{
				Code:    100601,
				Message: "MySQL错误",
				Data:    nil,
			})
		}
		c.JSON(200, result)
		return
	}
	c.JSON(200, &resp.Response{
		Code:    100501,
		Message: "参数错误",
		Data:    nil,
	})
}

func GetClipboard(c *gin.Context) {

}
