package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Mysql *MysqlConfig `yaml:"mysql"`
}

var (
	Conf = &Config{}
)

func Read(path string) *Config {
	Conf = new(Config)
	configFile, err := os.Open(path)
	if err != nil {
		fmt.Println("open config profile failed")
		return nil
	}
	defer configFile.Close()
	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Println("read config profile failed")
		return nil
	}
	err = yaml.Unmarshal(configData, Conf)
	if err != nil {
		fmt.Println("parse config profile failed")
		return nil
	}
	return Conf
}
