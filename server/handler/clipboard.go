package handler

import (
    "clipper_server/models/entity"
    "clipper_server/service"
    err2 "clipper_server/util/err"
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
            c.JSON(200, err2.ParamsError)
            return
        }

        now := time.Now()
        fmt.Println(now.Format("2006-01-02 15:04:05"), ": ", clipboardText)
        msg := entity.Message{
            CreateAt: now,
            Text:     fmt.Sprintf("%s", clipboardText),
        }
        result, err := service.SERVICE.CreateClipboardMessage(&msg)
        if err != nil {
            c.JSON(200, err2.MysqlError)
        }
        c.JSON(200, result)
        return
    }
    c.JSON(200, err2.ParamsError)
}

func GetClipboard(c *gin.Context) {
    clipboardTexts := service.SERVICE.GetClipboardMessages()
    c.JSON(200,  clipboardTexts)
}
