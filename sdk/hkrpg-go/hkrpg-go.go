package hkrpg_go

import (
	"net/http"

	"mhy-sdk/constant"
	"mhy-sdk/db"

	"github.com/gin-gonic/gin"
)

func GetComboToken(c *gin.Context) {
	accountId := c.Query("account_id")
	rsp := &constant.GateGetPlayerComboToken{
		Retcode:    0,
		AccountId:  accountId,
		ComboToken: "",
	}
	token := db.GetComboTokenByAccountId(accountId)
	if token == "" {
		rsp.Retcode = -1
	} else {
		rsp.ComboToken = token
	}
	c.JSON(http.StatusOK, rsp)
}
