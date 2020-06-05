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

	mylog.Logger.Info().Msgf("[Task Account] FuturesAccountService request param: %v, %v",
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

	mylog.Logger.Info().Msgf("[Task Account] FuturesAccountService request param: %v", userID)

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
	reduceOnly, _ := strconv.ParseBool(c.Query("reduceOnly"))
	quantity := c.Query("quantity")
	price := c.Query("price")
	newClientOrderId := c.Query("newClientOrderId")
	stopPrice := c.Query("stopPrice")
	closePosition, _ := strconv.ParseBool(c.Query("closePosition"))
	activationPrice := c.Query("activationPrice")
	callbackRate := c.Query("callbackRate")
	timeInForce := c.Query("timeInForce")
	workingType := c.Query("workingType")
	newOrderRespType := c.Query("newOrderRespType")

	mylog.Logger.Info().Msgf(
		"[Task Account] CreateOrderService request param: %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v",
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
	createOrder.PositionSide(futures.PositionSideType(positionSide))
	createOrder.Type(futures.OrderType(oType))
	createOrder.ReduceOnly(reduceOnly)
	createOrder.Quantity(quantity)
	createOrder.Price(price)
	createOrder.NewClientOrderID(newClientOrderId)
	createOrder.StopPrice(stopPrice)
	createOrder.ClosePosition(closePosition)
	createOrder.ActivationPrice(activationPrice)
	createOrder.CallbackRate(callbackRate)
	createOrder.TimeInForce(futures.TimeInForceType(timeInForce))
	createOrder.WorkingType(futures.WorkingType(workingType))
	createOrder.NewOrderRespType(futures.NewOrderRespType(newOrderRespType))

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
	orderId, _ := strconv.ParseInt(c.Query("orderId"), 10, 64)
	origClientOrderId := c.Query("origClientOrderId")

	mylog.Logger.Info().Msgf(
		"[Task Account] GetOrderService request param: %v, %v, %v, %v",
		userID, symbol, orderId, origClientOrderId)

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

	getOrder := client.NewGetOrderService()
	getOrder.Symbol(symbol)
	getOrder.OrderID(orderId)
	getOrder.OrigClientOrderID(origClientOrderId)

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
