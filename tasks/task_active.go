package tasks

import (
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
获取合约激活信息
*/
func GetActiveFuturesService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Account] GetActiveFuturesService request param: %v",
		userID)

	responseData, err := json.Marshal(db.GetActiveFuturesByUserID(userID))
	if err != nil {
		out.RespCode = EC_JSON_MARSHAL_ERR
		out.RespDesc = ErrorCodeMessage(EC_JSON_MARSHAL_ERR)
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = responseData

	c.Set("responseData", out)
}

func activeFutures(userID uint64) CommonResp {
	out := CommonResp{}

	//创建子账户
	createRes, err := trade.BAExClient.NewCreateSubAccountService().Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		return out
	}

	//为子账户开启合约权限
	_, err = trade.BAExClient.NewEnableSubAccountFutures().
		SubAccountId(createRes.SubAccountId).Futures(true).Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		return out
	}

	//创建子账户api key
	createApiRes, err := trade.BAExClient.NewCreateSubAccountApiService().
		SubAccountId(createRes.SubAccountId).CanTrade(true).FuturesTrade(true).Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		return out
	}

	err = db.CreateFuturesSubAccount(userID, createApiRes.SubAccountId, createApiRes.ApiKey, createApiRes.SecretKey)
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		return out
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	return out
}

/**
为用户申请子账户，创建API
*/
func CreateActiveFuturesService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Account] CreateActiveFuturesService request param: %v",
		userID)

	if active := db.GetActiveFuturesByUserID(userID); active == true {
		out.RespCode = EC_ALREADY_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_ALREADY_ACTIVE)
		c.Set("responseData", out)
		return
	}

	out = activeFutures(userID)

	c.Set("responseData", out)
}

/**
合约账户余额 (USER_DATA)
*/
func GetBalanceNoTokenService(c *gin.Context) {
	out := CommonResp{}

	var getBalanceNoTokenRequest GetBalanceNoTokenRequest
	if err := c.ShouldBindJSON(&getBalanceNoTokenRequest); err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.JSON(http.StatusOK, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] FuturesAccountNoTokenService request param: %v", getBalanceNoTokenRequest)

	if active := db.GetActiveFuturesByUserID(getBalanceNoTokenRequest.UserId); active == false {
		out = activeFutures(getBalanceNoTokenRequest.UserId)
		if out.RespCode != EC_NONE.Code() {
			c.JSON(http.StatusOK, out)
			return
		}
	}

	client, err := db.GetFuturesClientByUserID(getBalanceNoTokenRequest.UserId)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.JSON(http.StatusOK, out)
		return
	}

	list, err := client.NewGetBalanceService().Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusOK, out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
子母账户合约资产划转
*/
func CreateTransferNoTokenService(c *gin.Context) {
	out := CommonResp{}

	var createTransferNoTokenRequest CreateTransferNoTokenRequest
	if err := c.ShouldBindJSON(&createTransferNoTokenRequest); err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.JSON(http.StatusOK, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] CreateTransferNoTokenService request param: %v",
		createTransferNoTokenRequest)

	if active := db.GetActiveFuturesByUserID(createTransferNoTokenRequest.UserId); active == false {
		out = activeFutures(createTransferNoTokenRequest.UserId)
		if out.RespCode != EC_NONE.Code() {
			c.JSON(http.StatusOK, out)
			return
		}
	}

	client, err := db.GetSpotClientByUserID(createTransferNoTokenRequest.UserId)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.JSON(http.StatusOK, out)
		return
	}

	createTransferService := client.NewCreateTransferService()
	if createTransferNoTokenRequest.FromId != "" {
		createTransferService.FromId(createTransferNoTokenRequest.FromId)
	}
	if createTransferNoTokenRequest.ToId != "" {
		createTransferService.ToId(createTransferNoTokenRequest.ToId)
	}

	createTransferService.FuturesType(createTransferNoTokenRequest.FuturesType)
	createTransferService.Asset(createTransferNoTokenRequest.Asset)
	createTransferService.Amount(createTransferNoTokenRequest.Amount)

	list, err := createTransferService.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	responseData, err := json.Marshal(list)
	if err != nil {
		out.RespCode = EC_JSON_MARSHAL_ERR
		out.RespDesc = ErrorCodeMessage(EC_JSON_MARSHAL_ERR)
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = responseData

	c.Set("responseData", out)
}
