package dao

import (
    "clipper_server/models/entity"
    "fmt"
    "github.com/gomodule/redigo/redis"
)

func (dao *Dao) InsertClipboardText(clipboardEntity *entity.Message) error {
    err := dao.MySQL.
        Create(clipboardEntity).
        Error
    return err
}

func (dao *Dao) GetClipboardText(index int) (string, error) {
    clipboard := new(entity.Message)

    err := WrapDo(dao.Redis, func(con redis.Conn) error {
        err := con.Do("GET", fmt.Sprintf("clipboard:message:%d", index)))
        return
    })

    result, err := dao.Redis.Get(fmt.Sprintf("clipboard:message:%d", index)).Result()

    if err == nil {
        return result, nil
    }

    err = dao.MySQL.
        First(&clipboard, "id = ?", index).
        Error

    if err != nil {
        return "", err
    }

    return clipboard.Text, err
}
