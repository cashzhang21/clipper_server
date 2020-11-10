package handler

import (
	"clipper_server/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
		fmt.Printf("clipboard: %s", clipboardText)
		fmt.Println()
		return
	}
	c.JSON(200, &model.Response{
		Code:    1060003,
		Message: "参数错误",
		Data:    nil,
	})
}
