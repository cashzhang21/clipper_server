package service

import (
    "clipper_server/models/entity"
    "clipper_server/models/resp"
)

func (s *Service) CreateClipboardMessage(messageEntity *entity.Message) (*resp.Response, error) {

    s.dao.InsertClipboardText(messageEntity)
    r := &resp.Response{
        Code:    0,
        Message: "success",
        Data:    nil,
    }

    return r, nil
}

func (s *Service) GetClipboardMessages() []entity.Message {

    return s.dao.GetClipboardTexts()
}
