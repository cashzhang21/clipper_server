package dao

import (
	"clipper_server/models/entity"
)

func (dao *Dao) InsertClipboardText(clipboardEntity *entity.Message) error {
	err := dao.MySQL.Table(clipboardEntity.TableName()).
		Create(clipboardEntity).
		Error
	return err
}

func (dao *Dao) GetClipboardText() error {
	clipboard := new(entity.Message)
	err := dao.MySQL.Table(clipboard.TableName()).
		Find(&clipboard).
		Error
	return err
}
