package sdk

import (
	"mhy-sdk/sdk/combo"
	hkrphgo "mhy-sdk/sdk/hkrpg-go"
	"mhy-sdk/sdk/mdk"

	"github.com/gin-gonic/gin"
)

func (s *Sdk) newRouter() {
	s.router.POST("/account/risky/api/check", riskyApiCheckHandler)
	s.router.POST("/apm/dataUpload", apmdataUpload)
	s.router.POST("/data_abtest_api/config/experiment/list", getExperimentListHandler)
	hkrpg := s.router.Group("/hkrpg_:type")
	{
		hkrpg.GET("/combo/granter/api/getConfig", combo.ComboGranterApiGetConfigHandler) // 获取服务器配置
		hkrpg.GET("/combo/box/api/config/sdk/combo", combo.Combo)
		hkrpg.POST("/combo/granter/api/compareProtocolVersion", combo.CompareProtocolVersion)
		hkrpg.POST("/combo/granter/login/beforeVerify", combo.BeforeVerify)
		hkrpg.POST("/combo/red_dot/list", combo.RedDotList)
		hkrpg.POST("/combo/granter/login/v2/login", combo.V2LoginRequestHandler) // 获取combo token

		hkrpg.POST("/mdk/shield/api/login", mdk.LoginRequestHandler)   // 账号登录
		hkrpg.POST("/mdk/shield/api/verify", mdk.VerifyRequestHandler) // token登录
		hkrpg.GET("/mdk/shield/api/loadConfig", mdk.LoadConfig)
		hkrpg.GET("/mdk/agreement/api/getAgreementInfos", mdk.GetAgreementInfos)
		hkrpg.POST("/mdk/shopwindow/shopwindow/listPriceTier", mdk.ListPriceTier)
		hkrpg.POST("/mdk/shopwindow/shopwindow/listPriceTierV2", mdk.ListPriceTier)
		hkrpg.GET("/mdk/shopwindow/shopwindow/listPriceTierV2", mdk.ListPriceTier)
		hkrpg.POST("/mdk/luckycat/luckycat/createOrder", mdk.CreateOrder)
		hkrpg.POST("/mdk/shopwindow/shopwindow/getCurrencyAndCountryByIp", mdk.GetCurrencyAndCountryByIp)
	}
	server := s.router.Group("/hkrpg-go")
	{
		server.GET("/getComboToken", hkrphgo.GetComboToken)
	}
}

func riskyApiCheckHandler(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":\"none\",\"action\":\"ACTION_NONE\",\"geetest\":null}}")
}

func getExperimentListHandler(c *gin.Context) {
	c.Header("Content-type", "application/json")
	_, _ = c.Writer.WriteString("{\"retcode\":0,\"success\":true,\"message\":\"\",\"data\":[{\"code\":1000,\"type\":2,\"config_id\":\"14\",\"period_id\":\"6125_197\",\"version\":\"1\",\"configs\":{\"cardType\":\"direct\"}}]}")
}

func apmdataUpload(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}
