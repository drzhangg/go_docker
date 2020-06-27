package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MysqlConfig struct {
	Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func DBConfig() MysqlConfig {
	file, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		fmt.Errorf("read config file failed, err:%v", err.Error())
	}

	var mysqlConf MysqlConfig
	err = yaml.Unmarshal(file, &mysqlConf)
	if err != nil {
		fmt.Errorf("read config file failed, err:%v", err.Error())
	}
	fmt.Println(mysqlConf)
	return mysqlConf
}
