package db

import (
	"mhy-sdk/constant"
)

// QueryClientConfig 拉取全部客户端配置
func QueryClientConfig() []*constant.ClientConfig {
	var clientConfigList []*constant.ClientConfig
	gormDb.Find(&clientConfigList)
	return clientConfigList
}

// QueryRegionConfig 拉取全部区服配置
func QueryRegionConfig() []*constant.RegionConfig {
	var regionConfigList []*constant.RegionConfig
	gormDb.Find(&regionConfigList)
	return regionConfigList
}

// QueryClientRegionConfigByName 根据客户端拉取区服信息
func QueryClientRegionConfigByName(name string) []*constant.ClientRegionConfig {
	var clientRegionConfigList []*constant.ClientRegionConfig
	gormDb.Where("client_version = ?", name).Find(&clientRegionConfigList)
	return clientRegionConfigList
}
