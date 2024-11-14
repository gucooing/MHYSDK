package db

import (
	"fmt"
	"time"

	"mhy-sdk/constant"
)

// QueryAccountByFieldUsername 查询账号
func QueryAccountByFieldUsername(Username string) *constant.Account {
	var account constant.Account
	gormDb.Model(&constant.Account{}).Where("username = ?", Username).First(&account)
	return &account
}

// GetAccountByFieldAccountId 查询账号
func GetAccountByFieldAccountId(AccountId uint32) *constant.Account {
	var account constant.Account
	gormDb.Model(&constant.Account{}).Where("account_id = ?", AccountId).First(&account)
	return &account
}

// AddAccountFieldByFieldName 添加新账号
func AddAccountFieldByFieldName(account *constant.Account) (uint32, error) {
	if err := gormDb.Create(account).Error; err == nil {
		return account.AccountId, nil
	} else {
		return 0, err
	}
}

// SetComboTokenByAccountId 更新ComboToken
func SetComboTokenByAccountId(accountId uint32, comboToken string) {
	key := fmt.Sprintf("player_comboToken:%v", accountId)
	err := GetRedisClient("player_login").Set(ctx, key, comboToken, 168*time.Hour).Err()
	if err != nil {
		return
	}
	return
}

// GetComboTokenByAccountId 获取更新ComboToken
func GetComboTokenByAccountId(accountId string) string {
	key := fmt.Sprintf("player_comboToken:%s", accountId)
	comboToken, _ := GetRedisClient("player_login").Get(ctx, key).Result()
	return comboToken
}
