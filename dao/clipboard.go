package dao

import (
    "clipper_server/models/entity"
    "github.com/go-redis/redis"
    "math"
    "time"
)

func (dao *Dao) InsertClipboardText(clipboardEntity *entity.Message) {
    dao.Redis.Incr("messages:id")
    dao.Redis.ZAdd("messages:", redis.Z{
        Score:  float64(clipboardEntity.CreateAt.Unix()),
        Member: clipboardEntity.Text,
    })
}

func (dao *Dao) GetClipboardTexts() (messageSlice []entity.Message) {
    messages := dao.Redis.ZRangeWithScores("messages:", 0, -1).Val()
    for _, message := range messages {
        sec, dec := math.Modf(message.Score)
        messageSlice = append(messageSlice, entity.Message{
            CreateAt: time.Unix(int64(sec), int64(dec*(1e9))),
            Text:     message.Member.(string),
        })
    }
    return messageSlice
}
