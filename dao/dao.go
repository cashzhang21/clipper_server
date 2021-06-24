package dao

import (
    "clipper_server/config"
    "clipper_server/models/entity"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
    "strings"
    "time"
)

type Dao struct {
    MySQL *gorm.DB
    Redis *redis.Pool
}

func New() (dao *Dao) {
    dao = &Dao{
        MySQL: mysqlClient(),
        Redis: redisClient(),
    }

    return dao
}

func redisClient() (r *redis.Pool) {
    redisConfig := config.Conf.Redis
    r = &redis.Pool{
        MaxIdle:         redisConfig.PoolConfig.Idle,
        IdleTimeout:     redisConfig.PoolConfig.IdleTimeout,
        MaxActive:       redisConfig.PoolConfig.Active,
        Wait:            redisConfig.PoolConfig.Wait,
        MaxConnLifetime: redisConfig.MaxConnLifetime,
        Dial: func() (redis.Conn, error) {
            endpoint := fmt.Sprintf("%s:%d", redisConfig.Endpoint.Address, redisConfig.Endpoint.Port)
            c, err := redis.Dial(
                redisConfig.Proto,
                endpoint,
                redis.DialConnectTimeout(redisConfig.ConnectTimeout),
                redis.DialReadTimeout(redisConfig.ReadTimeout),
                redis.DialWriteTimeout(redisConfig.WriteTimeout),
            )
            if err != nil {
                return nil, err
            }

            if redisConfig.Auth != "" {
                if _, err := c.Do("AUTH", redisConfig.Auth); err != nil {
                    c.Close()
                    return nil, err
                }
            }

            if redisConfig.DB != 0 {
                if _, err := c.Do("SELECT", redisConfig.DB); err != nil {
                    c.Close()
                    return nil, err
                }
                return c, nil
            }

            return c, nil
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if time.Since(t) < redisConfig.PoolConfig.CheckTime {
                return nil
            }
            _, err := c.Do("PING")
            return err
        },
    }

    return r
}

func mysqlClient() (mysqlClient *gorm.DB) {
    c := config.Conf.Mysql
    dsnConfig := c.DSN
    d, err := gorm.Open(
        mysql.New(
            mysql.Config{
                DriverName:                "mysql",
                DSN:                       constructDsn(dsnConfig),
                SkipInitializeWithVersion: false,
                DefaultStringSize:         256,
                DisableDatetimePrecision:  true,
                DontSupportRenameIndex:    true,
                DontSupportRenameColumn:   true,
            }), &gorm.Config{})

    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }

    sqlDb, err := d.DB()
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    sqlDb.SetMaxOpenConns(c.Active)
    sqlDb.SetMaxIdleConns(c.Idle)
    sqlDb.SetConnMaxLifetime(c.IdleTimeout)

    err = d.AutoMigrate(&entity.Message{})
    if err != nil {
        return nil
    }
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

func WrapDo(p *redis.Pool, doFunction func(con redis.Conn) error) error {
    con := p.Get()
    defer con.Close()

    return doFunction(con)
}
