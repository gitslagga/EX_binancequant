package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
获取充值地址 (支持多网络) (USER_DATA)
*/
func DepositsAddressService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	coin := c.Query("coin")
	network := c.Query("network")

	mylog.Logger.Info().Msgf("[Task Account] DepositsAddressService request param: %s, %s, %s",
		userID, coin, network)

	if coin == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	depositsAddress := client.NewDepositsAddressService()
	depositsAddress.Coin(coin)
	depositsAddress.Network(network)

	list, err := depositsAddress.Do(data.NewContext())
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
获取充值历史（支持多网络） (USER_DATA)
*/
func ListDepositsService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	coin := c.Query("coin")
	status, _ := strconv.Atoi(c.Query("status"))
	startTime, _ := strconv.ParseInt(c.Query("startTime"), 10, 64)
	endTime, _ := strconv.ParseInt(c.Query("endTime"), 10, 64)
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	mylog.Logger.Info().Msgf("[Task Account] ListDepositsService request param: %s, %s, %s, %s, %s, %s, %s",
		userID, coin, status, startTime, endTime, offset, limit)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	listDeposits := client.NewListDepositsService()
	listDeposits.Coin(coin)
	listDeposits.Status(status)
	listDeposits.StartTime(startTime)
	listDeposits.EndTime(endTime)
	listDeposits.Offset(offset)
	listDeposits.Limit(limit)

	list, err := listDeposits.Do(data.NewContext())
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
现货账户信息 (USER_DATA)
*/
func SpotAccountService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Account] SpotAccountService request param: %s", userID)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewGetAccountService().Do(data.NewContext())
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
合约账户信息 (USER_DATA)
*/
func FuturesAccountService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Account] FuturesAccountService request param: %s", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewGetAccountService().Do(data.NewContext())
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
