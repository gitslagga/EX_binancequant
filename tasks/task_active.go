package tasks

import (
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
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

/**
子母账户合约资产划转
*/
func CreateTransferFuturesService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var createTransferRequest CreateTransferRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &createTransferRequest)
	if err != nil || createTransferRequest.FuturesType == 0 || createTransferRequest.Asset == "" || createTransferRequest.Amount == 0 {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] CreateTransferFuturesService request param: %v",
		createTransferRequest)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	createTransferService := client.NewCreateTransferService()
	if createTransferRequest.FromId != "" {
		createTransferService.FromId(createTransferRequest.FromId)
	}
	if createTransferRequest.ToId != "" {
		createTransferService.ToId(createTransferRequest.ToId)
	}

	createTransferService.FuturesType(createTransferRequest.FuturesType)
	createTransferService.Asset(createTransferRequest.Asset)
	createTransferService.Amount(createTransferRequest.Amount)

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
