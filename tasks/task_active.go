package tasks

import (
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"context"
	"github.com/gin-gonic/gin"
)

/**
获取合约激活信息
*/
func GetActiveFuturesService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Account] GetActiveFuturesService request param: %v",
		userID)

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = db.GetActiveFuturesByUserID(userID)

	c.Set("responseData", out)
}

/**
为用户申请子账户，创建API
*/
func CreateActiveFuturesService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Account] CreateActiveFuturesService request param: %v",
		userID)

	if active := db.GetActiveFuturesByUserID(userID); active == true {
		out.RespCode = EC_ALREADY_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_ALREADY_ACTIVE)
		c.Set("responseData", out)
		return
	}

	//创建子账户
	createRes, err := trade.BAExClient.NewCreateSubAccountService().Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	//为子账户开启合约权限
	_, err = trade.BAExClient.NewEnableSubAccountFutures().
		SubAccountId(createRes.SubAccountId).Futures(true).Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	//创建子账户api key
	createApiRes, err := trade.BAExClient.NewCreateSubAccountApiService().
		SubAccountId(createRes.SubAccountId).CanTrade(true).FuturesTrade(true).Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	err = db.CreateFuturesSubAccount(userID, createApiRes.SubAccountId, createApiRes.ApiKey, createApiRes.SecretKey)
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()

	c.Set("responseData", out)
}
