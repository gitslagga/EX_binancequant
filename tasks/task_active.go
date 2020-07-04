package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
获取合约激活信息
*/
func GetActiveFuturesService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Account] GetActiveFuturesService request param: %v",
		userID)

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = db.GetActiveFuturesByUserID(userID)

	c.JSON(http.StatusOK, out)
	return
}

/**
为用户申请子账户，创建API
*/
func CreateActiveFuturesService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Account] CreateActiveFuturesService request param: %v",
		userID)

	if active := db.GetActiveFuturesByUserID(userID); active == true {
		out.ErrorCode = data.EC_ALREADY_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_ALREADY_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	createRes, err := trade.BAExClient.NewCreateSubAccountService().Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	_, err = trade.BAExClient.NewEnableSubAccountFutures().
		SubAccountId(createRes.SubAccountId).Futures(true).Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	createApiRes, err := trade.BAExClient.NewCreateSubAccountApiService().
		SubAccountId(createRes.SubAccountId).CanTrade(true).FuturesTrade(true).Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	err = db.CreateFuturesSubAccount(userID, createApiRes.SubAccountId, createApiRes.ApiKey, createApiRes.SecretKey)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()

	c.JSON(http.StatusOK, out)
	return
}
