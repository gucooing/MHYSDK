package conf

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel   string               `json:"LogLevel"`
	AutoCreate bool                 `json:"AutoCreate"`
	HttpNet    *HttpNet             `json:"HttpNet"`
	MysqlDsn   string               `json:"MysqlDsn"`
	RedisConf  map[string]RedisConf `json:"RedisConf"`
}

type RedisConf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type HttpNet struct {
	InnerAddr string `json:"InnerAddr"`
	InnerPort string `json:"InnerPort"`
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
}

func SetDefaultConfig() {
	CONF = DefaultConfig
}

var FileNotExist = errors.New("config file not found")

func LoadConfig(confName string) error {
	filePath := confName
	f, err := os.Open(filePath)
	if err != nil {
		return FileNotExist
	}
	defer func() {
		_ = f.Close()
	}()
	c := new(Config)
	d := json.NewDecoder(f)
	if err := d.Decode(c); err != nil {
		return err
	}
	CONF = c
	return nil
}

var DefaultConfig = &Config{
	LogLevel:   "Info",
	AutoCreate: true,
	HttpNet: &HttpNet{
		InnerAddr: "127.0.0.1",
		InnerPort: "8080",
	},
	MysqlDsn: "root:password@tcp(127.0.0.1:3306)/mhy_sdk?charset=utf8mb4&parseTime=True&loc=Local",
	RedisConf: map[string]RedisConf{
		"player_login": {
			Addr:     "127.0.0.1:6379",
			Password: "password",
			DB:       1,
		},
	},
}
