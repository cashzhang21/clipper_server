package entity

import "time"

type Message struct {
    Id       int
    CreateAt time.Time
    Text     string
}

func (Message) TableName() string {
    return "messages"
}
