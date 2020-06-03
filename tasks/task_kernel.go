package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
下单
*/
func PostSwapOrder(c *gin.Context) {
	out := data.CommonResp{}
	orderParam := data.OrderParam{}

	if err := c.ShouldBindJSON(&orderParam); err != nil {
		mylog.Logger.Info().Msgf("[Task Service] PostSwapOrder request orderParam: %s", orderParam)
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	token := c.GetHeader("token")
	userID, err := db.ConvertTokenToUserID(token)
	if err != nil {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Service] PostSwapOrder request param: %s, %s", userID, orderParam)

	client, err := db.GetClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	order := &trade.BasePlaceOrderInfo{}
	order.Size = orderParam.Size
	order.Type = orderParam.Type
	order.MatchPrice = orderParam.MatchPrice
	order.Price = orderParam.Price

	list, err := client.PostSwapOrder(orderParam.InstrumentID, order)
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
撤单
*/
func CancelSwapInstrumentOrder(c *gin.Context) {
	out := data.CommonResp{}

	token := c.GetHeader("token")
	userID, err := db.ConvertTokenToUserID(token)
	instrumentID := c.Param("instrument_id")
	orderID := c.Param("order_id")

	if err != nil || instrumentID == "" || orderID == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Service] CancelSwapInstrumentOrder request param: %s, %s, %s", userID, instrumentID, orderID)

	client, err := db.GetClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.PostSwapCancelOrder(instrumentID, orderID)
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
