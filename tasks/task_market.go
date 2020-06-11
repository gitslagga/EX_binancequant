package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
获取服务器时间
*/
func ServerTimeService(c *gin.Context) {
	out := data.CommonResp{}

	list, err := trade.BAExFuturesClient.NewServerTimeService().Do(data.NewContext())
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
深度信息
*/
func DepthService(c *gin.Context) {
	out := data.CommonResp{}

	symbol := c.Query("symbol")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Market] DepthService request param: %v, %v",
		symbol, limit)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	depth := trade.BAExFuturesClient.NewDepthService()
	depth.Symbol(symbol)
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			depth.Limit(iLimit)
		}
	}

	list, err := depth.Do(data.NewContext())
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
近期成交(归集)
*/
func AggTradesService(c *gin.Context) {
	out := data.CommonResp{}

	symbol := c.Query("symbol")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Market] AggTradesService request param: %v, %v",
		symbol, limit)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	aggTrade := trade.BAExFuturesClient.NewAggTradesService()
	aggTrade.Symbol(symbol)
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			aggTrade.Limit(iLimit)
		}
	}

	list, err := aggTrade.Do(data.NewContext())
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
K线数据
*/
func KlinesService(c *gin.Context) {
	out := data.CommonResp{}

	symbol := c.Query("symbol")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	interval := c.Query("interval")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Market] KlinesService request param: %v, %v, %v, %v, %v",
		symbol, startTime, endTime, interval, limit)

	if symbol == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	klines := trade.BAExFuturesClient.NewKlinesService()
	klines.Symbol(symbol)
	if startTime != "" {
		iStartTime, err := strconv.ParseInt(startTime, 10, 64)
		if err == nil {
			klines.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		iEndTime, err := strconv.ParseInt(endTime, 10, 64)
		if err == nil {
			klines.EndTime(iEndTime)
		}
	}
	if interval != "" {
		klines.Interval(interval)
	}
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			klines.Limit(iLimit)
		}
	}

	list, err := klines.Do(data.NewContext())
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
最新标记价格和资金费率
*/
func PremiumIndexService(c *gin.Context) {
	out := data.CommonResp{}

	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf("[Task Market] PremiumIndexService request param: %v", symbol)

	premiumIndex := trade.BAExFuturesClient.NewPremiumIndexService()

	if symbol != "" {
		premiumIndex.Symbol(symbol)
	}

	list, err := premiumIndex.Do(data.NewContext())
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
24hr价格变动情况
*/
func ListPriceChangeStatsService(c *gin.Context) {
	out := data.CommonResp{}

	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf("[Task Market] ListPriceChangeStatsService request param: %v", symbol)

	priceChangeStats := trade.BAExFuturesClient.NewListPriceChangeStatsService()

	if symbol != "" {
		priceChangeStats.Symbol(symbol)
	}

	list, err := priceChangeStats.Do(data.NewContext())
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
最新价格
*/
func ListPricesService(c *gin.Context) {
	out := data.CommonResp{}

	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf("[Task Market] ListPricesService request param: %v", symbol)

	listPrices := trade.BAExFuturesClient.NewListPricesService()

	if symbol != "" {
		listPrices.Symbol(symbol)
	}

	list, err := listPrices.Do(data.NewContext())
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
获取交易规则和交易对
*/
func ExchangeInfoService(c *gin.Context) {
	out := data.CommonResp{}

	list, err := trade.BAExFuturesClient.NewExchangeInfoService().Do(data.NewContext())

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
