package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade/futures"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
更改持仓模式（TRADE）
*/
func ChangePositionModeService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	var positionModeRequest data.PositionModeRequest
	if err := c.ShouldBindJSON(&positionModeRequest); err != nil {
		mylog.Logger.Error().Msgf("[Task Account] ChangePositionModeService request param err: %v, %v",
			userID, err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] FuturesAccountService request param: %v, %v",
		userID, positionModeRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	positionMode := client.NewChangePositionModeService()
	positionMode.DualSide(positionModeRequest.DualSidePosition)

	err = positionMode.Do(data.NewContext())
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
查询持仓模式（USER_DATA）
*/
func GetPositionModeService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Futures] FuturesAccountService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewGetPositionModeService().Do(data.NewContext())
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
下单 (TRADE)
*/
func CreateOrderService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	var orderRequest data.OrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		mylog.Logger.Error().Msgf("[Task Account] CreateOrderService request param err: %v, %v",
			userID, err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] CreateOrderService request param: %v, %v",
		userID, orderRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	createOrder := client.NewCreateOrderService()
	createOrder.Symbol(orderRequest.Symbol)
	createOrder.Side(futures.SideType(orderRequest.Side))
	createOrder.Type(futures.OrderType(orderRequest.Type))
	if orderRequest.PositionSide != "" {
		createOrder.PositionSide(futures.PositionSideType(orderRequest.PositionSide))
	}
	if orderRequest.ReduceOnly != false {
		createOrder.ReduceOnly(orderRequest.ReduceOnly)
	}
	if orderRequest.Quantity != 0 {
		createOrder.Quantity(orderRequest.Quantity)
	}
	if orderRequest.Price != 0 {
		createOrder.Price(orderRequest.Price)
	}
	if orderRequest.NewClientOrderId != "" {
		createOrder.NewClientOrderID(orderRequest.NewClientOrderId)
	}
	if orderRequest.StopPrice != 0 {
		createOrder.StopPrice(orderRequest.StopPrice)
	}
	if orderRequest.ClosePosition != false {
		createOrder.ClosePosition(orderRequest.ClosePosition)
	}
	if orderRequest.ActivationPrice != 0 {
		createOrder.ActivationPrice(orderRequest.ActivationPrice)
	}
	if orderRequest.CallbackRate != 0 {
		createOrder.CallbackRate(orderRequest.CallbackRate)
	}
	if orderRequest.TimeInForce != "" {
		createOrder.TimeInForce(futures.TimeInForceType(orderRequest.TimeInForce))
	}
	if orderRequest.WorkingType != "" {
		createOrder.WorkingType(futures.WorkingType(orderRequest.WorkingType))
	}
	if orderRequest.NewOrderRespType != "" {
		createOrder.NewOrderRespType(futures.NewOrderRespType(orderRequest.NewOrderRespType))
	}

	list, err := createOrder.Do(data.NewContext())
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
查询订单 (USER_DATA)
*/
func GetOrderService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	symbol := c.Query("symbol")
	orderId := c.Query("orderId")
	origClientOrderId := c.Query("origClientOrderId")

	mylog.Logger.Info().Msgf(
		"[Task Futures] GetOrderService request param: %v, %v, %v, %v",
		userID, symbol, orderId, origClientOrderId)

	if symbol == "" || (orderId == "" && origClientOrderId == "") {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	getOrder := client.NewGetOrderService()
	getOrder.Symbol(symbol)
	if orderId != "" {
		iOrderId, err := strconv.ParseInt(orderId, 10, 64)
		if err == nil {
			getOrder.OrderID(iOrderId)
		}
	}
	if origClientOrderId != "" {
		getOrder.OrigClientOrderID(origClientOrderId)
	}

	list, err := getOrder.Do(data.NewContext())
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
撤销订单 (TRADE)
*/
func CancelOrderService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	symbol := c.Query("symbol")
	orderId := c.Query("orderId")
	origClientOrderId := c.Query("origClientOrderId")

	mylog.Logger.Info().Msgf(
		"[Task Futures] CancelOrderService request param: %v, %v, %v, %v",
		userID, symbol, orderId, origClientOrderId)

	if symbol == "" || (orderId == "" && origClientOrderId == "") {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	cancelOrder := client.NewCancelOrderService()
	cancelOrder.Symbol(symbol)
	if orderId != "" {
		iOrderId, err := strconv.ParseInt(orderId, 10, 64)
		if err == nil {
			cancelOrder.OrderID(iOrderId)
		}
	}
	if origClientOrderId != "" {
		cancelOrder.OrigClientOrderID(origClientOrderId)
	}

	list, err := cancelOrder.Do(data.NewContext())
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
撤销全部订单 (TRADE)
*/
func CancelAllOpenOrdersService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf(
		"[Task Futures] CancelAllOpenOrdersService request param: %v, %v",
		userID, symbol)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	cancelAllOpenOrders := client.NewCancelAllOpenOrdersService()
	cancelAllOpenOrders.Symbol(symbol)

	err = cancelAllOpenOrders.Do(data.NewContext())
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
查看当前全部挂单 (USER_DATA)
*/
func ListOpenOrdersService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf(
		"[Task Futures] ListOpenOrdersService request param: %v, %v",
		userID, symbol)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	listOpenOrders := client.NewListOpenOrdersService()
	listOpenOrders.Symbol(symbol)

	list, err := listOpenOrders.Do(data.NewContext())
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
查询所有订单（包括历史订单） (USER_DATA)
*/
func ListOrdersService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	symbol := c.Query("symbol")
	orderId := c.Query("orderId")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Futures] ListOrdersService request param: %v, %v, %v, %v, %v, %v",
		userID, symbol, orderId, startTime, endTime, limit)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	listOrders := client.NewListOrdersService()
	listOrders.Symbol(symbol)
	if orderId != "" {
		iOrderId, err := strconv.ParseInt(orderId, 10, 64)
		if err == nil {
			listOrders.OrderID(iOrderId)
		}
	}
	if startTime != "" {
		iStartTime, err := strconv.ParseInt(startTime, 10, 64)
		if err == nil {
			listOrders.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		iEndTime, err := strconv.ParseInt(endTime, 10, 64)
		if err == nil {
			listOrders.EndTime(iEndTime)
		}
	}
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			listOrders.Limit(iLimit)
		}
	}

	list, err := listOrders.Do(data.NewContext())
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
账户余额 (USER_DATA)
*/
func GetBalanceService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Futures] GetBalanceService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewGetBalanceService().Do(data.NewContext())
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
调整开仓杠杆 (TRADE)
*/
func ChangeLeverageService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	var leverageRequest data.LeverageRequest
	if err := c.ShouldBindJSON(&leverageRequest); err != nil {
		mylog.Logger.Error().Msgf("[Task Futures] ChangeLeverageService request param err: %v, %v",
			userID, err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] ChangeLeverageService request param: %v, %v",
		userID, leverageRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	changeLeverage := client.NewChangeLeverageService()
	changeLeverage.Symbol(leverageRequest.Symbol)
	changeLeverage.Leverage(leverageRequest.Leverage)

	list, err := changeLeverage.Do(data.NewContext())
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
变换逐全仓模式 (TRADE)
*/
func ChangeMarginTypeService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	var marginTypeRequest data.MarginTypeRequest
	if err := c.ShouldBindJSON(&marginTypeRequest); err != nil {
		mylog.Logger.Error().Msgf("[Task Futures] ChangeMarginTypeService request param err: %v, %v",
			userID, err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] ChangeMarginTypeService request param: %v, %v",
		userID, marginTypeRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	changeMarginType := client.NewChangeMarginTypeService()
	changeMarginType.Symbol(marginTypeRequest.Symbol)
	changeMarginType.MarginType(futures.MarginType(marginTypeRequest.MarginType))

	err = changeMarginType.Do(data.NewContext())
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
调整逐仓保证金 (TRADE)
*/
func UpdatePositionMarginService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	var positionMarginRequest data.PositionMarginRequest
	if err := c.ShouldBindJSON(&positionMarginRequest); err != nil {
		mylog.Logger.Info().Msgf("[Task Futures] UpdatePositionMarginService request param err: %v, %v",
			userID, err)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] UpdatePositionMarginService request param: %v, %v",
		userID, positionMarginRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	updatePositionMargin := client.NewUpdatePositionMarginService()
	updatePositionMargin.Symbol(positionMarginRequest.Symbol)
	updatePositionMargin.Amount(positionMarginRequest.Amount)
	updatePositionMargin.Type(positionMarginRequest.Type)
	if positionMarginRequest.PositionSide != "" {
		updatePositionMargin.PositionSide(futures.PositionSideType(positionMarginRequest.PositionSide))
	}

	err = updatePositionMargin.Do(data.NewContext())
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
逐仓保证金变动历史 (TRADE)
*/
func GetPositionMarginHistoryService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	symbol := c.Query("symbol")
	sType := c.Query("type")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Futures] GetPositionMarginHistoryService request param: %v, %v, %v, %v, %v, %v",
		userID, symbol, sType, startTime, endTime, limit)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	positionMarginHistory := client.NewGetPositionMarginHistoryService()
	positionMarginHistory.Symbol(symbol)
	if sType != "" {
		iType, err := strconv.Atoi(sType)
		if err == nil {
			positionMarginHistory.Type(iType)
		}
	}
	if startTime != "" {
		iStartTime, err := strconv.ParseInt(startTime, 10, 64)
		if err == nil {
			positionMarginHistory.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		iEndTime, err := strconv.ParseInt(endTime, 10, 64)
		if err == nil {
			positionMarginHistory.EndTime(iEndTime)
		}
	}
	if limit != "" {
		iLimit, err := strconv.ParseInt(limit, 10, 64)
		if err == nil {
			positionMarginHistory.Limit(iLimit)
		}
	}

	list, err := positionMarginHistory.Do(data.NewContext())
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
用户持仓风险 (USER_DATA)
*/
func GetPositionRiskService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Futures] GetPositionRiskService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewGetPositionRiskService().Do(data.NewContext())
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
账户成交历史 (USER_DATA)
*/
func GetTradeHistoryService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	symbol := c.Query("symbol")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	fromId := c.Query("fromId")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Futures] GetTradeHistoryService request param: %v, %v, %v, %v, %v, %v",
		userID, symbol, fromId, startTime, endTime, limit)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	positionMarginHistory := client.NewGetTradeHistoryService()
	positionMarginHistory.Symbol(symbol)
	if fromId != "" {
		iFromId, err := strconv.ParseUint(fromId, 10, 64)
		if err == nil {
			positionMarginHistory.FromId(iFromId)
		}
	}
	if startTime != "" {
		iStartTime, err := strconv.ParseInt(startTime, 10, 64)
		if err == nil {
			positionMarginHistory.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		iEndTime, err := strconv.ParseInt(endTime, 10, 64)
		if err == nil {
			positionMarginHistory.EndTime(iEndTime)
		}
	}
	if limit != "" {
		iLimit, err := strconv.ParseInt(limit, 10, 64)
		if err == nil {
			positionMarginHistory.Limit(iLimit)
		}
	}

	list, err := positionMarginHistory.Do(data.NewContext())
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
获取账户损益资金流水(USER_DATA)
*/
func GetIncomeHistoryService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	symbol := c.Query("symbol")
	incomeType := c.Query("incomeType")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Futures] GetIncomeHistoryService request param: %v, %v, %v, %v, %v, %v",
		userID, symbol, incomeType, startTime, endTime, limit)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	incomeHistory := client.NewGetIncomeHistoryService()
	incomeHistory.Symbol(symbol)
	if incomeType != "" {
		incomeHistory.IncomeType(incomeType)
	}
	if startTime != "" {
		iStartTime, err := strconv.ParseInt(startTime, 10, 64)
		if err == nil {
			incomeHistory.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		iEndTime, err := strconv.ParseInt(endTime, 10, 64)
		if err == nil {
			incomeHistory.EndTime(iEndTime)
		}
	}
	if limit != "" {
		iLimit, err := strconv.ParseInt(limit, 10, 64)
		if err == nil {
			incomeHistory.Limit(iLimit)
		}
	}

	list, err := incomeHistory.Do(data.NewContext())
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
杠杆分层标准 (USER_DATA)
*/
func GetLeverageBracketService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Futures] GetLeverageBracketService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewGetLeverageBracketService().Do(data.NewContext())
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
生成listenKey (USER_STREAM)
*/
func StartUserStreamService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Futures] StartUserStreamService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewStartUserStreamService().Do(data.NewContext())
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
延长listenKey有效期 (USER_STREAM)
*/
func KeepaliveUserStreamService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Futures] KeepaliveUserStreamService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	err = client.NewKeepaliveUserStreamService().Do(data.NewContext())
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
关闭listenKey (USER_STREAM)
*/
func CloseUserStreamService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Futures] CloseUserStreamService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NOT_ACTIVE
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_NOT_ACTIVE)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	err = client.NewCloseUserStreamService().Do(data.NewContext())
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
