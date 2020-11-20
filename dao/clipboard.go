package dao

import (
	"clipper_server/models/entity"
)

func (dao *Dao) InsertClipboardText(clipboardEntity *entity.Message) error {
	err := dao.MySQL.Table(clipboardEntity.TableName()).Create(clipboardEntity).Error
	return err
}