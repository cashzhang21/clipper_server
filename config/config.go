package config

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "time"
)

type Config struct {
    Mysql *MysqlConfig `yaml:"mysql"`
    Redis *RedisConfig `yaml:"redis"`
}

var (
    Conf = &Config{}
)

func Read(path string) *Config {
    Conf = new(Config)
    configFile, err := os.Open(path)
    if err != nil {
        fmt.Println("open config profile failed")
        panic(err)
    }
    defer configFile.Close()
    configData, err := ioutil.ReadAll(configFile)
    if err != nil {
        fmt.Println("read config profile failed")
        panic(err)
    }
    err = yaml.Unmarshal(configData, Conf)
    if err != nil {
        fmt.Println("parse config profile failed")
        panic(err)
    }
    return Conf
}

type EndpointConfig struct {
    Address string `yaml:"address"`
    Port    int    `yaml:"port"`
}

type DSNConfig struct {
    UserName string          `yaml:"userName"`
    Password string          `yaml:"password"`
    Endpoint *EndpointConfig `yaml:"endpoint"`
    DBName   string          `yaml:"dbName"`
    Options  []string        `yaml:"options"`
}

type MysqlConfig struct {
    DSN         *DSNConfig    `yaml:"dsn"`
    Active      int           `yaml:"active"`
    Idle        int           `yaml:"idle"`
    IdleTimeout time.Duration `yaml:"idleTimeout"`
}

type RedisConfig struct {
    *PoolConfig `yaml:",inline"`
    Proto string `yaml:"proto"`
    DB int `yaml:"db"`
    Endpoint *EndpointConfig `yaml:"endpoint"`
    Auth string `yaml:"auth"`
    MaxConnLifetime time.Duration `yaml:"maxConnLifetime"`
    *RenderConfig `yaml:",inline"`
}

type PoolConfig struct {
    Active int `yaml:"active"`
    Idle int `yaml:"idle"`
    IdleTimeout time.Duration `yaml:"idleTimeout"`
    CheckTime time.Duration `yaml:"checkTime"`
    Wait bool `yaml:"wait"`
    ReadTimeout time.Duration `yaml:"readTimeout"`
    WriteTimeout time.Duration `yaml:"writeTimeout"`
    ConnectTimeout time.Duration `yaml:"connectTimeout"`
}

type RenderConfig struct {
    Stdout bool `yaml:"stdout"`
    StdoutPattern string `yaml:"stdoutPattern"`
    OutDir string `yaml:"outDir"`
    OutFile string `yaml:"outFile"`
    OutPattern string `yaml:"outPattern"`
    FileBufferSize int64 `yaml:"fileBufferSize"`
    MaxLogFile int `yaml:"maxLogFile"`
    RotateSize int64 `yaml:"rotateSize"`
}
