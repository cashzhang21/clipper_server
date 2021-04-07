package service

import (
    "clipper_server/models/entity"
    "clipper_server/models/resp"
)

func (s *Service) CreateClipboardMessage(messageEntity *entity.Message) (*resp.Response, error) {

    err := s.dao.InsertClipboardText(messageEntity)
    if err != nil {
        return nil, err
    }

    r := &resp.Response{
        Code:    0,
        Message: "success",
        Data:    nil,
    }

    return r, nil
}
