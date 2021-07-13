package entity

import "time"

type Message struct {
    Id       int        `json:"id"`
    CreateAt time.Time  `json:"create_at"`
    Text     string     `json:"text"`
}

/*func (Message) TableName() string {
    return "messages"
}*/
