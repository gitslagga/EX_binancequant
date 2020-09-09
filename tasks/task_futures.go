package tasks

import (
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade/futures"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

/**
更改持仓模式（TRADE）
*/
func ChangePositionModeService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var positionModeRequest PositionModeRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &positionModeRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] FuturesAccountService request param: %v, %v",
		userID, positionModeRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	positionMode := client.NewChangePositionModeService()
	positionMode.DualSide(positionModeRequest.DualSidePosition)

	err = positionMode.Do(context.Background())
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
查询持仓模式（USER_DATA）
*/
func GetPositionModeService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Futures] FuturesAccountService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	list, err := client.NewGetPositionModeService().Do(context.Background())
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

/**
下单 (TRADE)
*/
func CreateOrderService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var orderRequest OrderRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &orderRequest)
	if err != nil || orderRequest.Symbol == "" || orderRequest.Side == "" || orderRequest.Type == "" {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] CreateOrderService request param: %v, %v",
		userID, orderRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
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

	list, err := createOrder.Do(context.Background())
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

/**
查询订单 (USER_DATA)
*/
func GetOrderService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var getOrderRequest GetOrderRequest
	err := c.ShouldBindQuery(&getOrderRequest)
	if err != nil || (getOrderRequest.OrderId == 0 && getOrderRequest.OrigClientOrderId == "") {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf(
		"[Task Futures] GetOrderService request param: %v, %v",
		userID, getOrderRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	getOrder := client.NewGetOrderService()
	getOrder.Symbol(getOrderRequest.Symbol)
	if getOrderRequest.OrderId != 0 {
		getOrder.OrderID(getOrderRequest.OrderId)
	}
	if getOrderRequest.OrigClientOrderId != "" {
		getOrder.OrigClientOrderID(getOrderRequest.OrigClientOrderId)
	}

	list, err := getOrder.Do(context.Background())
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

/**
撤销订单 (TRADE)
*/
func CancelOrderService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var cancelOrderRequest CancelOrderRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &cancelOrderRequest)
	if err != nil || cancelOrderRequest.Symbol == "" || (cancelOrderRequest.OrderId == 0 && cancelOrderRequest.OrigClientOrderId == "") {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf(
		"[Task Futures] CancelOrderService request param: %v, %v",
		userID, cancelOrderRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	cancelOrder := client.NewCancelOrderService()
	cancelOrder.Symbol(cancelOrderRequest.Symbol)
	if cancelOrderRequest.OrderId != 0 {
		cancelOrder.OrderID(cancelOrderRequest.OrderId)
	}
	if cancelOrderRequest.OrigClientOrderId != "" {
		cancelOrder.OrigClientOrderID(cancelOrderRequest.OrigClientOrderId)
	}

	list, err := cancelOrder.Do(context.Background())
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

/**
撤销全部订单 (TRADE)
*/
func CancelAllOpenOrdersService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var cancelAllOpenOrdersRequest CancelAllOpenOrdersRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &cancelAllOpenOrdersRequest)
	if err != nil || cancelAllOpenOrdersRequest.Symbol == "" {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf(
		"[Task Futures] CancelAllOpenOrdersService request param: %v, %v",
		userID, cancelAllOpenOrdersRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	cancelAllOpenOrders := client.NewCancelAllOpenOrdersService()
	cancelAllOpenOrders.Symbol(cancelAllOpenOrdersRequest.Symbol)

	err = cancelAllOpenOrders.Do(context.Background())
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
查看当前全部挂单 (USER_DATA)
*/
func ListOpenOrdersService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var listOpenOrdersRequest ListOpenOrdersRequest
	err := c.ShouldBindQuery(&listOpenOrdersRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf(
		"[Task Futures] ListOpenOrdersService request param: %v, %v",
		userID, listOpenOrdersRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	listOpenOrders := client.NewListOpenOrdersService()
	if listOpenOrdersRequest.Symbol != "" {
		listOpenOrders.Symbol(listOpenOrdersRequest.Symbol)
	}

	list, err := listOpenOrders.Do(context.Background())
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

/**
查询所有订单（包括历史订单） (USER_DATA)
*/
func ListOrdersService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var listOrdersRequest ListOrdersRequest
	err := c.ShouldBindQuery(&listOrdersRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] ListOrdersService request param: %v, %v",
		userID, listOrdersRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	listOrders := client.NewListOrdersService()
	if listOrdersRequest.Symbol != "" {
		listOrders.Symbol(listOrdersRequest.Symbol)
	}
	listOrders.Symbol(listOrdersRequest.Symbol)
	if listOrdersRequest.OrderId != 0 {
		listOrders.OrderID(listOrdersRequest.OrderId)
	}
	if listOrdersRequest.StartTime != 0 {
		listOrders.StartTime(listOrdersRequest.StartTime)
	}
	if listOrdersRequest.EndTime != 0 {
		listOrders.EndTime(listOrdersRequest.EndTime)
	}
	if listOrdersRequest.Limit != 0 {
		listOrders.Limit(listOrdersRequest.Limit)
	}

	list, err := listOrders.Do(context.Background())
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

/**
账户余额 (USER_DATA)
*/
func GetBalanceService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Futures] GetBalanceService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	list, err := client.NewGetBalanceService().Do(context.Background())
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

/**
调整开仓杠杆 (TRADE)
*/
func ChangeLeverageService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var leverageRequest LeverageRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &leverageRequest)
	if err != nil || leverageRequest.Symbol == "" || leverageRequest.Leverage == 0 {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] ChangeLeverageService request param: %v, %v",
		userID, leverageRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	changeLeverage := client.NewChangeLeverageService()
	changeLeverage.Symbol(leverageRequest.Symbol)
	changeLeverage.Leverage(leverageRequest.Leverage)

	list, err := changeLeverage.Do(context.Background())
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

/**
变换逐全仓模式 (TRADE)
*/
func ChangeMarginTypeService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var marginTypeRequest MarginTypeRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &marginTypeRequest)
	if err != nil || marginTypeRequest.Symbol == "" || marginTypeRequest.MarginType == "" {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] ChangeMarginTypeService request param: %v, %v",
		userID, marginTypeRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	changeMarginType := client.NewChangeMarginTypeService()
	changeMarginType.Symbol(marginTypeRequest.Symbol)
	changeMarginType.MarginType(futures.MarginType(marginTypeRequest.MarginType))

	err = changeMarginType.Do(context.Background())
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
调整逐仓保证金 (TRADE)
*/
func UpdatePositionMarginService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var positionMarginRequest PositionMarginRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &positionMarginRequest)
	if err != nil || positionMarginRequest.Symbol == "" || positionMarginRequest.Amount == 0 || positionMarginRequest.Type == 0 {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] UpdatePositionMarginService request param: %v, %v",
		userID, positionMarginRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	updatePositionMargin := client.NewUpdatePositionMarginService()
	updatePositionMargin.Symbol(positionMarginRequest.Symbol)
	updatePositionMargin.Amount(positionMarginRequest.Amount)
	updatePositionMargin.Type(positionMarginRequest.Type)
	if positionMarginRequest.PositionSide != "" {
		updatePositionMargin.PositionSide(futures.PositionSideType(positionMarginRequest.PositionSide))
	}

	err = updatePositionMargin.Do(context.Background())
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
逐仓保证金变动历史 (TRADE)
*/
func GetPositionMarginHistoryService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var getPositionMarginHistoryRequest GetPositionMarginHistoryRequest
	err := c.ShouldBindQuery(&getPositionMarginHistoryRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] GetPositionMarginHistoryService request param: %v, %v",
		userID, getPositionMarginHistoryRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	positionMarginHistory := client.NewGetPositionMarginHistoryService()
	positionMarginHistory.Symbol(getPositionMarginHistoryRequest.Symbol)
	if getPositionMarginHistoryRequest.Type != 0 {
		positionMarginHistory.Type(getPositionMarginHistoryRequest.Type)
	}
	if getPositionMarginHistoryRequest.StartTime != 0 {
		positionMarginHistory.StartTime(getPositionMarginHistoryRequest.StartTime)
	}
	if getPositionMarginHistoryRequest.EndTime != 0 {
		positionMarginHistory.EndTime(getPositionMarginHistoryRequest.EndTime)
	}
	if getPositionMarginHistoryRequest.Limit != 0 {
		positionMarginHistory.Limit(getPositionMarginHistoryRequest.Limit)
	}

	list, err := positionMarginHistory.Do(context.Background())
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

/**
用户持仓风险 (USER_DATA)
*/
func GetPositionRiskService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var getPositionRiskRequest GetPositionRiskRequest
	err := c.ShouldBindQuery(&getPositionRiskRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] GetPositionRiskService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	positionRisk := client.NewGetPositionRiskService()
	if getPositionRiskRequest.Symbol != "" {
		positionRisk.Symbol(getPositionRiskRequest.Symbol)
	}

	list, err := positionRisk.Do(context.Background())
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

/**
账户成交历史 (USER_DATA)
*/
func GetTradeHistoryService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var getTradeHistoryRequest GetTradeHistoryRequest
	err := c.ShouldBindQuery(&getTradeHistoryRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] GetTradeHistoryService request param: %v, %v",
		userID, getTradeHistoryRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	positionMarginHistory := client.NewGetTradeHistoryService()
	if getTradeHistoryRequest.Symbol != "" {
		positionMarginHistory.Symbol(getTradeHistoryRequest.Symbol)
	}
	positionMarginHistory.Symbol(getTradeHistoryRequest.Symbol)
	if getTradeHistoryRequest.FromId != 0 {
		positionMarginHistory.FromId(getTradeHistoryRequest.FromId)
	}
	if getTradeHistoryRequest.StartTime != 0 {
		positionMarginHistory.StartTime(getTradeHistoryRequest.StartTime)
	}
	if getTradeHistoryRequest.EndTime != 0 {
		positionMarginHistory.EndTime(getTradeHistoryRequest.EndTime)
	}
	if getTradeHistoryRequest.Limit != 0 {
		positionMarginHistory.Limit(getTradeHistoryRequest.Limit)
	}

	list, err := positionMarginHistory.Do(context.Background())
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

/**
获取账户损益资金流水(USER_DATA)
*/
func GetIncomeHistoryService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var getIncomeHistoryRequest GetIncomeHistoryRequest
	err := c.ShouldBindQuery(&getIncomeHistoryRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] GetIncomeHistoryService request param: %v, %v",
		userID, getIncomeHistoryRequest)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	incomeHistory := client.NewGetIncomeHistoryService()
	if getIncomeHistoryRequest.Symbol != "" {
		incomeHistory.Symbol(getIncomeHistoryRequest.Symbol)
	}
	if getIncomeHistoryRequest.IncomeType != "" {
		incomeHistory.IncomeType(getIncomeHistoryRequest.IncomeType)
	}
	if getIncomeHistoryRequest.StartTime != 0 {
		incomeHistory.StartTime(getIncomeHistoryRequest.StartTime)
	}
	if getIncomeHistoryRequest.EndTime != 0 {
		incomeHistory.EndTime(getIncomeHistoryRequest.EndTime)
	}
	if getIncomeHistoryRequest.Limit != 0 {
		incomeHistory.Limit(getIncomeHistoryRequest.Limit)
	}

	list, err := incomeHistory.Do(context.Background())
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

/**
杠杆分层标准 (USER_DATA)
*/
func GetLeverageBracketService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var getLeverageBracketRequest GetLeverageBracketRequest
	err := c.ShouldBindQuery(&getLeverageBracketRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Futures] GetLeverageBracketService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	leverageBracket := client.NewGetLeverageBracketService()
	if getLeverageBracketRequest.Symbol != "" {
		leverageBracket.Symbol(getLeverageBracketRequest.Symbol)
	}

	list, err := leverageBracket.Do(context.Background())
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

/**
生成listenKey (USER_STREAM)
*/
func StartUserStreamService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Futures] StartUserStreamService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	list, err := client.NewStartUserStreamService().Do(context.Background())
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

/**
延长listenKey有效期 (USER_STREAM)
*/
func KeepaliveUserStreamService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Futures] KeepaliveUserStreamService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	err = client.NewKeepaliveUserStreamService().Do(context.Background())
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
关闭listenKey (USER_STREAM)
*/
func CloseUserStreamService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Futures] CloseUserStreamService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	err = client.NewCloseUserStreamService().Do(context.Background())
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
