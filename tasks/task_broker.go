package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
Create a Sub Account
*/
func CreateSubAccountService(c *gin.Context) {
	out := data.CommonResp{}

	list, err := trade.BAExClient.NewCreateSubAccountService().Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
Enable Futures for Sub Account
*/
func EnableSubAccountFuturesService(c *gin.Context) {
	out := data.CommonResp{}

	var enableFuturesRequest data.EnableFuturesRequest
	if err := c.ShouldBindJSON(&enableFuturesRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] EnableSubAccountFuturesService request param err: %v",
			err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] EnableSubAccountFuturesService request param: %v",
		enableFuturesRequest)

	list, err := trade.BAExClient.NewEnableSubAccountFutures().
		SubAccountId(enableFuturesRequest.SubAccountId).
		Futures(enableFuturesRequest.Futures).
		Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
Create Api Key for Sub Account
*/
func CreateSubAccountApiService(c *gin.Context) {
	out := data.CommonResp{}

	var createApiKeyRequest data.CreateApiKeyRequest
	if err := c.ShouldBindJSON(&createApiKeyRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] CreateSubAccountApiService request param err: %v",
			err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] CreateSubAccountApiService request param: %v",
		createApiKeyRequest)

	list, err := trade.BAExClient.NewCreateSubAccountApiService().
		SubAccountId(createApiKeyRequest.SubAccountId).
		CanTrade(createApiKeyRequest.CanTrade).
		FuturesTrade(createApiKeyRequest.FuturesTrade).
		Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
Delete Sub Account Api Key
*/
func DeleteSubAccountApiService(c *gin.Context) {
	out := data.CommonResp{}

	var subAccountApiKeyRequest data.SubAccountApiKeyRequest
	if err := c.ShouldBindJSON(&subAccountApiKeyRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] DeleteSubAccountApiService request param err: %v",
			err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] DeleteSubAccountApiService request param: %v",
		subAccountApiKeyRequest)

	err := trade.BAExClient.NewDeleteSubAccountApiService().
		SubAccountId(subAccountApiKeyRequest.SubAccountId).
		SubAccountApiKey(subAccountApiKeyRequest.SubAccountApiKey).
		Do(data.NewContext())
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

/**
Query Sub Account Api Key
*/
func GetSubAccountApiService(c *gin.Context) {
	out := data.CommonResp{}

	var subAccountApiKeyRequest data.SubAccountApiKeyRequest
	if err := c.ShouldBindJSON(&subAccountApiKeyRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountApiService request param err: %v",
			err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountApiService request param: %v",
		subAccountApiKeyRequest)

	list, err := trade.BAExClient.NewGetSubAccountApiService().
		SubAccountId(subAccountApiKeyRequest.SubAccountId).
		SubAccountApiKey(subAccountApiKeyRequest.SubAccountApiKey).
		Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

func ChangeSubAccountApiPermissionService(c *gin.Context) {

}

func GetSubAccountService(c *gin.Context) {

}

func ChangeCommissionFuturesService(c *gin.Context) {

}

func GetCommissionFuturesService(c *gin.Context) {

}

func GetInfoService(c *gin.Context) {

}

func CreateTransferService(c *gin.Context) {

}

func GetTransferService(c *gin.Context) {

}

func GetSubAccountDepositHistService(c *gin.Context) {

}

func GetSubAccountSpotSummaryService(c *gin.Context) {

}

func GetSubAccountFuturesSummaryService(c *gin.Context) {

}

func GetRebateRecentRecordService(c *gin.Context) {

}

func GenerateRebateHistoryService(c *gin.Context) {

}

func GetRebateHistoryService(c *gin.Context) {

}
