package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
创建子账户
*/
func CreateSubAccountService(c *gin.Context) {
	out := data.CommonResp{}

	list, err := trade.BAExClient.NewCreateSubAccountService().Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
为子帐户启用合约
*/
func EnableSubAccountFuturesService(c *gin.Context) {
	out := data.CommonResp{}

	var enableFuturesRequest data.EnableFuturesRequest
	if err := c.ShouldBindJSON(&enableFuturesRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] EnableSubAccountFuturesService request param err: %v",
			err)
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] EnableSubAccountFuturesService request param: %v",
		enableFuturesRequest)

	list, err := trade.BAExClient.NewEnableSubAccountFutures().
		SubAccountId(enableFuturesRequest.SubAccountId).
		Futures(enableFuturesRequest.Futures).
		Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
为子帐户创建Api密钥
*/
func CreateSubAccountApiService(c *gin.Context) {
	out := data.CommonResp{}

	var createApiKeyRequest data.CreateApiKeyRequest
	if err := c.ShouldBindJSON(&createApiKeyRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] CreateSubAccountApiService request param err: %v",
			err)
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] CreateSubAccountApiService request param: %v",
		createApiKeyRequest)

	list, err := trade.BAExClient.NewCreateSubAccountApiService().
		SubAccountId(createApiKeyRequest.SubAccountId).
		CanTrade(createApiKeyRequest.CanTrade).
		FuturesTrade(createApiKeyRequest.FuturesTrade).
		Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
删除子帐户Api密钥
*/
func DeleteSubAccountApiService(c *gin.Context) {
	out := data.CommonResp{}

	var deleteApiKeyRequest data.DeleteApiKeyRequest
	if err := c.ShouldBindJSON(&deleteApiKeyRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] DeleteSubAccountApiService request param err: %v",
			err)
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] DeleteSubAccountApiService request param: %v",
		deleteApiKeyRequest)

	err := trade.BAExClient.NewDeleteSubAccountApiService().
		SubAccountId(deleteApiKeyRequest.SubAccountId).
		SubAccountApiKey(deleteApiKeyRequest.SubAccountApiKey).
		Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()

	c.JSON(http.StatusOK, out)
	return
}

/**
查询子帐户api密钥
*/
func GetSubAccountApiService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	subAccountApiKey := c.Query("subAccountApiKey")

	mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountApiService request param: %v, %v",
		subAccountId, subAccountApiKey)

	if subAccountId == "" {
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	getSubAccountApiService := trade.BAExClient.NewGetSubAccountApiService()
	getSubAccountApiService.SubAccountId(subAccountId)
	if subAccountApiKey != "" {
		getSubAccountApiService.SubAccountApiKey(subAccountApiKey)
	}

	list, err := getSubAccountApiService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
更改子帐户Api权限
*/
func ChangeSubAccountApiPermissionService(c *gin.Context) {
	out := data.CommonResp{}

	var changeApiPermissionRequest data.ChangeApiPermissionRequest
	if err := c.ShouldBindJSON(&changeApiPermissionRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] ChangeSubAccountApiPermissionService request param err: %v",
			err)
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
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
		Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
查询子账户
*/
func GetSubAccountService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")

	mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountService request param: %v",
		subAccountId)

	getSubAccountService := trade.BAExClient.NewGetSubAccountService()
	if subAccountId != "" {
		getSubAccountService.SubAccountId(subAccountId)
	}

	list, err := getSubAccountService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
更改子账户合约佣金调整
*/
func ChangeCommissionFuturesService(c *gin.Context) {
	out := data.CommonResp{}

	var changeCommissionFuturesRequest data.ChangeCommissionFuturesRequest
	if err := c.ShouldBindJSON(&changeCommissionFuturesRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] ChangeCommissionFuturesService request param err: %v",
			err)
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
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
		Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
查询子账户合约佣金调整
*/
func GetCommissionFuturesService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf("[Task Broker] GetCommissionFuturesService request param: %v, %v",
		subAccountId, symbol)

	if subAccountId == "" {
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	getCommissionFuturesService := trade.BAExClient.NewGetCommissionFuturesService()
	getCommissionFuturesService.SubAccountId(subAccountId)
	if symbol != "" {
		getCommissionFuturesService.Symbol(symbol)
	}

	list, err := getCommissionFuturesService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
经纪人账户信息
*/
func GetInfoService(c *gin.Context) {
	out := data.CommonResp{}

	list, err := trade.BAExClient.NewGetInfoService().Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
子账户划转
*/
func CreateTransferService(c *gin.Context) {
	out := data.CommonResp{}

	var createTransferRequest data.CreateTransferRequest
	if err := c.ShouldBindJSON(&createTransferRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] CreateTransferService request param err: %v",
			err)
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] CreateTransferService request param: %v",
		createTransferRequest)

	createTransferService := trade.BAExClient.NewCreateTransferService()
	if createTransferRequest.FromId != "" {
		createTransferService.FromId(createTransferRequest.FromId)
	}
	if createTransferRequest.ToId != "" {
		createTransferService.ToId(createTransferRequest.ToId)
	}
	if createTransferRequest.ClientTranId != "" {
		createTransferService.ClientTranId(createTransferRequest.ClientTranId)
	}
	createTransferService.Asset(createTransferRequest.Asset)
	createTransferService.Amount(createTransferRequest.Amount)

	list, err := createTransferService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
查询子账户划转历史
*/
func GetTransferService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	clientTranId := c.Query("clientTranId")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	page := c.Query("page")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Broker] GetTransferService request param: %v, %v, %v, %v, %v, %v",
		subAccountId, clientTranId, startTime, endTime, page, limit)

	if subAccountId == "" {
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	getTransferService := trade.BAExClient.NewGetTransferService()
	getTransferService.SubAccountId(subAccountId)
	if clientTranId != "" {
		getTransferService.ClientTranId(clientTranId)
	}
	if startTime != "" {
		if iStartTime, err := strconv.ParseUint(startTime, 10, 64); err == nil {
			getTransferService.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		if iEndTime, err := strconv.ParseUint(endTime, 10, 64); err == nil {
			getTransferService.EndTime(iEndTime)
		}
	}
	if page != "" {
		if iPage, err := strconv.Atoi(page); err == nil {
			getTransferService.Page(iPage)
		}
	}
	if limit != "" {
		if iLimit, err := strconv.Atoi(limit); err == nil {
			getTransferService.Limit(iLimit)
		}
	}

	list, err := getTransferService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
获取子账户充币历史
*/
func GetSubAccountDepositHistService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	coin := c.Query("coin")
	status := c.Query("status")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	limit := c.Query("limit")
	offset := c.Query("offset")

	mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountDepositHistService request param: %v, %v, %v, %v, %v, %v, %v",
		subAccountId, coin, status, startTime, endTime, limit, offset)

	getSubAccountDepositHistService := trade.BAExClient.NewGetSubAccountDepositHistService()
	if subAccountId != "" {
		getSubAccountDepositHistService.SubAccountId(subAccountId)
	}
	if coin != "" {
		getSubAccountDepositHistService.Coin(coin)
	}
	if status != "" {
		iStatus, err := strconv.Atoi(status)
		if err == nil {
			getSubAccountDepositHistService.Status(iStatus)
		}
	}
	if startTime != "" {
		if iStartTime, err := strconv.ParseUint(startTime, 10, 64); err == nil {
			getSubAccountDepositHistService.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		if iEndTime, err := strconv.ParseUint(endTime, 10, 64); err == nil {
			getSubAccountDepositHistService.EndTime(iEndTime)
		}
	}
	if limit != "" {
		if iLimit, err := strconv.Atoi(limit); err == nil {
			getSubAccountDepositHistService.Limit(iLimit)
		}
	}
	if offset != "" {
		if iOffset, err := strconv.Atoi(offset); err == nil {
			getSubAccountDepositHistService.Offset(iOffset)
		}
	}

	list, err := getSubAccountDepositHistService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
查询子账户现货资产信息
*/
func GetSubAccountSpotSummaryService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	page := c.Query("page")
	size := c.Query("size")

	mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountSpotSummaryService request param: %v, %v, %v",
		subAccountId, page, size)

	getSubAccountSpotSummaryService := trade.BAExClient.NewGetSubAccountSpotSummaryService()
	if subAccountId != "" {
		getSubAccountSpotSummaryService.SubAccountId(subAccountId)
	}
	if page != "" {
		if iPage, err := strconv.ParseUint(page, 10, 64); err == nil {
			getSubAccountSpotSummaryService.Page(iPage)
		}
	}
	if size != "" {
		if iSize, err := strconv.ParseUint(size, 10, 64); err == nil {
			getSubAccountSpotSummaryService.Size(iSize)
		}
	}

	list, err := getSubAccountSpotSummaryService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
查询子账户合约资产信息
*/
func GetSubAccountFuturesSummaryService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	page := c.Query("page")
	size := c.Query("size")

	mylog.Logger.Info().Msgf("[Task Broker] GetSubAccountFuturesSummaryService request param: %v, %v, %v",
		subAccountId, page, size)

	getSubAccountFuturesSummaryService := trade.BAExClient.NewGetSubAccountFuturesSummaryService()
	if subAccountId != "" {
		getSubAccountFuturesSummaryService.SubAccountId(subAccountId)
	}
	if page != "" {
		if iPage, err := strconv.ParseUint(page, 10, 64); err == nil {
			getSubAccountFuturesSummaryService.Page(iPage)
		}
	}
	if size != "" {
		if iSize, err := strconv.ParseUint(size, 10, 64); err == nil {
			getSubAccountFuturesSummaryService.Size(iSize)
		}
	}

	list, err := getSubAccountFuturesSummaryService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
查询经纪人佣金回扣最近记录
*/
func GetRebateRecentRecordService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Broker] GetRebateRecentRecordService request param: %v, %v, %v, %v",
		subAccountId, startTime, endTime, limit)

	if subAccountId == "" || startTime == "" || endTime == "" || limit == "" {
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	getRebateRecentRecordService := trade.BAExClient.NewGetRebateRecentRecordService()
	getRebateRecentRecordService.SubAccountId(subAccountId)
	if iStartTime, err := strconv.ParseUint(startTime, 10, 64); err == nil {
		getRebateRecentRecordService.StartTime(iStartTime)
	}
	if iEndTime, err := strconv.ParseUint(endTime, 10, 64); err == nil {
		getRebateRecentRecordService.EndTime(iEndTime)
	}
	if iLimit, err := strconv.Atoi(limit); err == nil {
		getRebateRecentRecordService.Limit(iLimit)
	}

	list, err := getRebateRecentRecordService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
生成经纪人佣金回扣历史记录
*/
func GenerateRebateHistoryService(c *gin.Context) {
	out := data.CommonResp{}

	var generateRebateHistory data.GenerateRebateHistoryRequest
	if err := c.ShouldBindJSON(&generateRebateHistory); err != nil {
		mylog.Logger.Info().Msgf("[Task Broker] GenerateRebateHistoryService request param err: %v",
			err)
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Broker] GenerateRebateHistoryService request param: %v",
		generateRebateHistory)

	generateRebateHistoryService := trade.BAExClient.NewGenerateRebateHistoryService()
	if generateRebateHistory.SubAccountId != "" {
		generateRebateHistoryService.SubAccountId(generateRebateHistory.SubAccountId)
	}
	if generateRebateHistory.StartTime != 0 {
		generateRebateHistoryService.StartTime(generateRebateHistory.StartTime)
	}
	if generateRebateHistory.EndTime != 0 {
		generateRebateHistoryService.EndTime(generateRebateHistory.EndTime)
	}

	list, err := generateRebateHistoryService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = list

	c.JSON(http.StatusOK, out)
	return
}

/**
查询经纪人佣金回扣记录
*/
func GetRebateHistoryService(c *gin.Context) {
	out := data.CommonResp{}

	subAccountId := c.Query("subAccountId")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Broker] GetRebateHistoryService request param: %v, %v, %v, %v",
		subAccountId, startTime, endTime, limit)

	if subAccountId == "" || startTime == "" || endTime == "" || limit == "" {
		out.RespCode = data.EC_PARAMS_ERR
		out.RespDesc = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	getRebateHistoryService := trade.BAExClient.NewGetRebateHistoryService()
	if subAccountId != "" {
		getRebateHistoryService.SubAccountId(subAccountId)
	}
	if startTime != "" {
		if iStartTime, err := strconv.ParseUint(startTime, 10, 64); err == nil {
			getRebateHistoryService.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		if iEndTime, err := strconv.ParseUint(endTime, 10, 64); err == nil {
			getRebateHistoryService.EndTime(iEndTime)
		}
	}
	if limit != "" {
		if iLimit, err := strconv.Atoi(limit); err == nil {
			getRebateHistoryService.Limit(iLimit)
		}
	}

	list, err := getRebateHistoryService.Do(context.Background())
	if err != nil {
		out.RespCode = data.EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.RespCode = data.EC_NONE.Code()
	out.RespDesc = data.EC_NONE.String()
	out.RespData = string(list)

	c.JSON(http.StatusOK, out)
	return
}
