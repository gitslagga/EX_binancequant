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
	dualSidePosition, err := strconv.ParseBool(c.Query("dualSidePosition"))

	mylog.Logger.Info().Msgf("[Task Futures] FuturesAccountService request param: %v, %v",
		userID, dualSidePosition)

	if err != nil {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	positionMode := client.NewChangePositionModeService()
	positionMode.DualSide(dualSidePosition)

	err = positionMode.Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = ""

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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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

	symbol := c.Query("symbol")
	side := c.Query("side")
	positionSide := c.Query("positionSide")
	oType := c.Query("type")
	reduceOnly := c.Query("reduceOnly")
	quantity := c.Query("quantity")
	price := c.Query("price")
	newClientOrderId := c.Query("newClientOrderId")
	stopPrice := c.Query("stopPrice")
	closePosition := c.Query("closePosition")
	activationPrice := c.Query("activationPrice")
	callbackRate := c.Query("callbackRate")
	timeInForce := c.Query("timeInForce")
	workingType := c.Query("workingType")
	newOrderRespType := c.Query("newOrderRespType")

	mylog.Logger.Info().Msgf(
		"[Task Futures] CreateOrderService request param: %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v",
		userID, symbol, side, positionSide, oType, reduceOnly, quantity, price, newClientOrderId, stopPrice, closePosition,
		activationPrice, callbackRate, timeInForce, workingType, newOrderRespType)

	if symbol == "" || side == "" || oType == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	createOrder := client.NewCreateOrderService()
	createOrder.Symbol(symbol)
	createOrder.Side(futures.SideType(side))
	createOrder.Type(futures.OrderType(oType))
	if positionSide != "" {
		createOrder.PositionSide(futures.PositionSideType(positionSide))
	}
	if reduceOnly != "" {
		bReduceOnly, err := strconv.ParseBool(reduceOnly)
		if err == nil {
			createOrder.ReduceOnly(bReduceOnly)
		}
	}
	if quantity != "" {
		createOrder.Quantity(quantity)
	}
	if price != "" {
		createOrder.Price(price)
	}
	if newClientOrderId != "" {
		createOrder.NewClientOrderID(newClientOrderId)
	}
	if stopPrice != "" {
		createOrder.StopPrice(stopPrice)
	}
	if closePosition != "" {
		bClosePosition, err := strconv.ParseBool(closePosition)
		if err == nil {
			createOrder.ClosePosition(bClosePosition)
		}
	}
	if activationPrice != "" {
		createOrder.ActivationPrice(activationPrice)
	}
	if callbackRate != "" {
		createOrder.CallbackRate(callbackRate)
	}
	if timeInForce != "" {
		createOrder.TimeInForce(futures.TimeInForceType(timeInForce))
	}
	if workingType != "" {
		createOrder.WorkingType(futures.WorkingType(workingType))
	}
	if newOrderRespType != "" {
		createOrder.NewOrderRespType(futures.NewOrderRespType(newOrderRespType))
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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
	out.Data = ""

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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
func ChangeLeveragevService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	symbol := c.Query("symbol")
	leverage, err := strconv.Atoi(c.Query("leverage"))

	mylog.Logger.Info().Msgf("[Task Futures] ChangeLeverageService request param: %v, %v, %v",
		userID, symbol, leverage)

	if symbol == "" || err != nil {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	changeLeverage := client.NewChangeLeverageService()
	changeLeverage.Symbol(symbol)
	changeLeverage.Leverage(leverage)

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
	symbol := c.Query("symbol")
	marginType := c.Query("marginType")

	mylog.Logger.Info().Msgf("[Task Futures] ChangeMarginTypeService request param: %v, %v, %v",
		userID, symbol, marginType)

	if symbol == "" || marginType == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	changeMarginType := client.NewChangeMarginTypeService()
	changeMarginType.Symbol(symbol)
	changeMarginType.MarginType(futures.MarginType(marginType))

	err = changeMarginType.Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = ""

	c.JSON(http.StatusOK, out)
	return
}

/**
调整逐仓保证金 (TRADE)
*/
func UpdatePositionMarginService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	symbol := c.Query("symbol")
	positionSide := c.Query("positionSide")
	amount := c.Query("amount")
	iType, err := strconv.Atoi(c.Query("type"))

	mylog.Logger.Info().Msgf("[Task Futures] UpdatePositionMarginService request param: %v, %v, %v, %v, %v",
		userID, symbol, positionSide, amount, iType)

	if symbol == "" || amount == "" || err != nil {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	updatePositionMargin := client.NewUpdatePositionMarginService()
	updatePositionMargin.Symbol(symbol)
	updatePositionMargin.Amount(amount)
	updatePositionMargin.Type(iType)
	if positionSide != "" {
		updatePositionMargin.PositionSide(futures.PositionSideType(positionSide))
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
	out.Data = ""

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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
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
