package constant

import (
	"database/sql"
)

// Account 账号信息
type Account struct {
	AccountId  uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	Username   string
	Token      string
	CreateTime int64
}

// ClientConfig 客户端配置
type ClientConfig struct {
	ClientVersion string `gorm:"primarykey"`
	Name          string
	DispatchSeed  string
	StopBeginTime sql.NullTime
	StopEndTime   sql.NullTime
	StopTips      string
	StopUrl       string
}

// ClientRegionConfig 客户端分区配置
type ClientRegionConfig struct {
	ClientVersion string
	RegionName    string
}

// RegionConfig 分区配置
type RegionConfig struct {
	RegionName  string `gorm:"primarykey"`
	Title       string
	DispatchUrl string
	SdkEnv      string
	// ClientSecret []byte
	// CdkUrl       string
}
