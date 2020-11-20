package dao

import (
	"clipper_server/config"
	"clipper_server/models/entity"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strings"
)

type Dao struct {
	MySQL *gorm.DB
	//Redis	*Redis
}

func New() (dao *Dao) {
	dao = &Dao{
		MySQL: mysqlClient(),
	}

	return dao
}

func mysqlClient() (mysql *gorm.DB) {
	c := config.Conf.Mysql
	dsnConfig := c.DSN
	d, err := gorm.Open(
		"mysql",
		constructDsn(dsnConfig),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	d.DB().SetMaxOpenConns(c.Active)
	d.DB().SetMaxIdleConns(c.Idle)
	d.DB().SetConnMaxLifetime(c.IdleTimeout)
	d.AutoMigrate(&entity.Message{})
	return d
}

func constructDsn(dsnConfig *config.DSNConfig) string {
	uri := fmt.Sprintf("%s:%s@(%s:%d)/%s",
		dsnConfig.UserName,
		dsnConfig.Password,
		dsnConfig.Endpoint.Address,
		dsnConfig.Endpoint.Port,
		dsnConfig.DBName,
	)
	if len(dsnConfig.Options) != 0 {
		uri = fmt.Sprintf("%s?%s", uri, strings.Join(dsnConfig.Options, "&"))
	}

	return uri
}
