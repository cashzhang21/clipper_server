package handler

import (
	"clipper_server/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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
