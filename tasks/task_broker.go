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

	var deleteApiKeyRequest data.DeleteApiKeyRequest
	if err := c.ShouldBindJSON(&deleteApiKeyRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] DeleteSubAccountApiService request param err: %v",
			err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] DeleteSubAccountApiService request param: %v",
		deleteApiKeyRequest)

	err := trade.BAExClient.NewDeleteSubAccountApiService().
		SubAccountId(deleteApiKeyRequest.SubAccountId).
		SubAccountApiKey(deleteApiKeyRequest.SubAccountApiKey).
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

	subAccountId := c.Query("subAccountId")
	subAccountApiKey := c.Query("subAccountApiKey")

	mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountApiService request param: %v, %v",
		subAccountId, subAccountApiKey)

	if subAccountId == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	getSubAccountApiService := trade.BAExClient.NewGetSubAccountApiService()
	getSubAccountApiService.SubAccountId(subAccountId)
	if subAccountApiKey != "" {
		getSubAccountApiService.SubAccountApiKey(subAccountApiKey)
	}

	list, err := getSubAccountApiService.Do(data.NewContext())
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
Change Sub Account Api Permission
*/
func ChangeSubAccountApiPermissionService(c *gin.Context) {
	out := data.CommonResp{}

	var changeApiPermissionRequest data.ChangeApiPermissionRequest
	if err := c.ShouldBindJSON(&changeApiPermissionRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] ChangeSubAccountApiPermissionService request param err: %v",
			err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] ChangeSubAccountApiPermissionService request param: %v",
		changeApiPermissionRequest)

	list, err := trade.BAExClient.NewChangeSubAccountApiPermissionService().
		SubAccountId(changeApiPermissionRequest.SubAccountId).
		SubAccountApiKey(changeApiPermissionRequest.SubAccountApiKey).
		CanTrade(changeApiPermissionRequest.CanTrade).
		FuturesTrade(changeApiPermissionRequest.FuturesTrade).
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
Query Sub Account
*/
func GetSubAccountService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")

	mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountService request param: %v",
		subAccountId)

	if subAccountId == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := trade.BAExClient.NewGetSubAccountService().
		SubAccountId(subAccountId).
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
Change Sub Account Futures Commission Adjustment
*/
func ChangeCommissionFuturesService(c *gin.Context) {
	out := data.CommonResp{}

	var changeCommissionFuturesRequest data.ChangeCommissionFuturesRequest
	if err := c.ShouldBindJSON(&changeCommissionFuturesRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] ChangeCommissionFuturesService request param err: %v",
			err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] ChangeCommissionFuturesService request param: %v",
		changeCommissionFuturesRequest)

	list, err := trade.BAExClient.NewChangeCommissionFuturesService().
		SubAccountId(changeCommissionFuturesRequest.SubAccountId).
		Symbol(changeCommissionFuturesRequest.Symbol).
		MakerAdjustment(changeCommissionFuturesRequest.MakerAdjustment).
		TakerAdjustment(changeCommissionFuturesRequest.TakerAdjustment).
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
Query Sub Account Futures Commission Adjustment
*/
func GetCommissionFuturesService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf("[Task Broker] GetCommissionFuturesService request param: %v, %v",
		subAccountId, symbol)

	if subAccountId == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	getCommissionFuturesService := trade.BAExClient.NewGetCommissionFuturesService()
	getCommissionFuturesService.SubAccountId(subAccountId)
	if symbol != "" {
		getCommissionFuturesService.Symbol(symbol)
	}

	list, err := getCommissionFuturesService.Do(data.NewContext())
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
Broker Account Information
*/
func GetInfoService(c *gin.Context) {
	out := data.CommonResp{}

	list, err := trade.BAExClient.NewGetInfoService().Do(data.NewContext())
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
