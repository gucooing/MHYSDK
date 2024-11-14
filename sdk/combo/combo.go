package combo

import (
	"encoding/json"

	"mhy-sdk/alg"
	"mhy-sdk/constant"
	"mhy-sdk/db"
	"mhy-sdk/logger"

	"github.com/gin-gonic/gin"
)

func ComboGranterApiGetConfigHandler(c *gin.Context) {
	getConfigrsq := new(constant.GranterApiGetConfig)

	data := &constant.GranterApiGetConfigData{
		Protocol:                true,
		QrEnabled:               true,
		LogLevel:                "INFO",
		AnnounceURL:             "https://hkrpg.alsl.xyz",
		PushAliasType:           0,
		DisableYsdkGuard:        true,
		EnableAnnouncePicPopup:  true,
		AppName:                 "崩坏RPG",
		FunctionalSwitchConfigs: make([]string, 0),
	}
	getConfigrsq.Retcode = 0
	getConfigrsq.Message = "OK"
	getConfigrsq.Data = data

	c.JSON(200, getConfigrsq)
}

func CompareProtocolVersion(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"modified\":false,\"protocol\":null}}")
}

func BeforeVerify(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"is_heartbeat_required\":false,\"is_realname_required\":false,\"is_guardian_required\":false}}")
}

func RedDotList(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"infos\":[]}}")
}

/*
流程:
1.检查token是否正确
2.若正确则生成token返回
3.若错误或不存在则返回错误
*/
func V2LoginRequestHandler(c *gin.Context) {
	requestData := new(constant.ComboTokenReq)
	err := c.ShouldBindJSON(requestData)
	if err != nil {
		logger.Error("parse ComboTokenReq error: %v", err)
		return
	}
	data := requestData.Data
	if len(data) == 0 {
		logger.Error("requestData.Data len == 0")
		return
	}
	loginData := new(constant.ComboTokenReqLoginTokenData)
	err = json.Unmarshal([]byte(data), loginData)
	if err != nil {
		logger.Error("Unmarshal ComboTokenReqLoginTokenData error: %v", err)
		return
	}
	accountId := alg.S2U32(loginData.Uid)
	responseData := new(constant.ComboTokenRsp)
	var account *constant.Account
	account = db.GetAccountByFieldAccountId(accountId)
	if account.AccountId != accountId {
		logger.Warn("查询不到此账户,uid: %s", loginData.Uid)
		c.Header("Content-type", "application/json")
		_, _ = c.Writer.WriteString("{\"data\":null,\"message\":\"游戏信息账号缓存错误\",\"retcode\":-103}")
		return
	} else {
		if account.Token == loginData.Token {
			comboToken := alg.GetRandomByteHexStr(20)
			db.SetComboTokenByAccountId(account.AccountId, comboToken)
			responseData.Retcode = 0
			responseData.Message = "OK"
			responseData.Data = &constant.ComboTokenRspLoginData{
				ComboID:       "0",
				OpenID:        loginData.Uid,
				ComboToken:    comboToken,
				Data:          "{\"guest\":false}",
				Heartbeat:     false,
				AccountType:   1,
				FatigueRemind: nil,
			}
			c.JSON(200, responseData)
			return
		} else {
			logger.Error("token验证失败,uid: %s", loginData.Uid)
			c.Header("Content-type", "application/json")
			_, _ = c.Writer.WriteString("{\"data\":null,\"message\":\"token验证失败\",\"retcode\":-103}")
			return
		}
	}
}
