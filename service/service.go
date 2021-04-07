package service

import (
    "clipper_server/config"
    "clipper_server/dao"
)

var (
    SERVICE *Service
)

type Service struct {
    config *config.Config
    dao    *dao.Dao
}

func New() *Service {

    SERVICE = &Service{
        config: config.Read("./config/config.yaml"),
        dao:    dao.New(),
    }
    return SERVICE
}
