package db

import (
	"context"
	"time"

	"mhy-sdk/constant"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gromlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var gormDb *gorm.DB
var redisClient map[string]*redis.Client
var err error

func NewMysql(dsn string) error {
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gromlogger.Default.LogMode(gromlogger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqlDB, err := gormDb.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(1000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(1 * time.Second) // 10 秒钟

	gormDb.AutoMigrate(&constant.Account{},
		&constant.ClientConfig{},
		&constant.ClientRegionConfig{},
		&constant.RegionConfig{},
	)

	return nil
}

var ctx = context.Background()

func NewRedis(name, addr, password string, db int) error {
	rdb := redis.NewClient(&redis.Options{
		Network:               "",
		Addr:                  addr,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              "",
		Password:              password,
		CredentialsProvider:   nil,
		DB:                    db,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		MaxActiveConns:        0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
		DisableIndentity:      false,
		IdentitySuffix:        "",
	})
	if _, err = rdb.Ping(ctx).Result(); err != nil {
		return err
	}
	SetRedisClient(name, rdb)
	return nil
}

func SetRedisClient(name string, rc *redis.Client) {
	if redisClient == nil {
		redisClient = make(map[string]*redis.Client)
	}
	redisClient[name] = rc
}

func GetRedisClient(name string) *redis.Client {
	if redisClient == nil {
		return nil
	}
	return redisClient[name]
}
