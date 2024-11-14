package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"mhy-sdk/conf"
	"mhy-sdk/db"
	"mhy-sdk/logger"
	"mhy-sdk/sdk"
)

func main() {
	confName := "config.json"
	err := conf.LoadConfig(confName)
	if err != nil {
		if err == conf.FileNotExist {
			fmt.Printf("找不到配置文件准备生成默认配置文件 %s \n", confName)
			p, _ := json.MarshalIndent(conf.DefaultConfig, "", "  ")
			cf, _ := os.Create(confName)
			_, err := cf.Write(p)
			cf.Close()
			if err != nil {
				fmt.Printf("生成默认配置文件失败 %s \n使用默认配置\n", err.Error())
				conf.SetDefaultConfig()
			} else {
				fmt.Printf("生成默认配置文件成功 \n")
				main()
			}
		} else {
			panic(err)
		}
	}
	cfg := conf.GetConfig()
	// 初始化日志
	logger.InitLogger("mhy-sdk", strings.ToUpper(cfg.LogLevel))
	logger.Info("mhy-sdk")
	// 初始化数据库
	newMysql(cfg.MysqlDsn)
	newRedis(cfg.RedisConf)
	// 初始化sdk
	s := sdk.NewSdk(cfg)
	s.Run()
}

func newMysql(dsn string) {
	err := db.NewMysql(dsn)
	if err != nil {
		panic(err)
	}
}

func newRedis(list map[string]conf.RedisConf) {
	for name, conf := range list {
		err := db.NewRedis(name, conf.Addr, conf.Password, conf.DB)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
	}
}
