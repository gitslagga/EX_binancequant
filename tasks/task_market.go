package tasks

import (
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
获取服务器时间
*/
func ServerTimeService(c *gin.Context) {
	out := CommonResp{}

	list, err := trade.BAExFuturesClient.NewServerTimeService().Do(context.Background())
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
深度信息
*/
func DepthService(c *gin.Context) {
	out := CommonResp{}

	symbol := c.Query("symbol")
	limit := c.Query("limit")

	if symbol == "" {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Market] DepthService request param: %v, %v",
		symbol, limit)

	depth := trade.BAExFuturesClient.NewDepthService()
	depth.Symbol(symbol)
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			depth.Limit(iLimit)
		}
	}

	list, err := depth.Do(context.Background())
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
近期成交(归集)
*/
func AggTradesService(c *gin.Context) {
	out := CommonResp{}

	symbol := c.Query("symbol")
	limit := c.Query("limit")

	if symbol == "" {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Market] AggTradesService request param: %v, %v",
		symbol, limit)

	aggTrade := trade.BAExFuturesClient.NewAggTradesService()
	aggTrade.Symbol(symbol)
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			aggTrade.Limit(iLimit)
		}
	}

	list, err := aggTrade.Do(context.Background())
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
K线数据
*/
func KlinesService(c *gin.Context) {
	out := CommonResp{}

	symbol := c.Query("symbol")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	interval := c.Query("interval")
	limit := c.Query("limit")

	if symbol == "" || interval == "" {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Market] KlinesService request param: %v, %v, %v, %v, %v",
		symbol, startTime, endTime, interval, limit)

	klines := trade.BAExFuturesClient.NewKlinesService()
	klines.Symbol(symbol)
	klines.Interval(interval)
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
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			klines.Limit(iLimit)
		}
	}

	list, err := klines.Do(context.Background())
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
最新标记价格和资金费率
*/
func PremiumIndexService(c *gin.Context) {
	out := CommonResp{}

	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf("[Task Market] PremiumIndexService request param: %v", symbol)

	premiumIndex := trade.BAExFuturesClient.NewPremiumIndexService()

	if symbol != "" {
		premiumIndex.Symbol(symbol)
	}

	list, err := premiumIndex.Do(context.Background())
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
24hr价格变动情况
*/
func ListPriceChangeStatsService(c *gin.Context) {
	out := CommonResp{}

	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf("[Task Market] ListPriceChangeStatsService request param: %v", symbol)

	priceChangeStats := trade.BAExFuturesClient.NewListPriceChangeStatsService()

	if symbol != "" {
		priceChangeStats.Symbol(symbol)
	}

	list, err := priceChangeStats.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	for i := 0; i < len(list); i++ {
		if list[i].Symbol == "LENDUSDT" {
			list = append(list[:i], list[i+1:]...)
			break
		}
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
最新价格
*/
func ListPricesService(c *gin.Context) {
	out := CommonResp{}

	symbol := c.Query("symbol")

	mylog.Logger.Info().Msgf("[Task Market] ListPricesService request param: %v", symbol)

	listPrices := trade.BAExFuturesClient.NewListPricesService()

	if symbol != "" {
		listPrices.Symbol(symbol)
	}

	list, err := listPrices.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	for i := 0; i < len(list); i++ {
		if list[i].Symbol == "LENDUSDT" {
			list = append(list[:i], list[i+1:]...)
			break
		}
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
获取交易规则和交易对
*/
func ExchangeInfoService(c *gin.Context) {
	out := CommonResp{}

	list, err := trade.BAExFuturesClient.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	for i := 0; i < len(list.Symbols); i++ {
		if list.Symbols[i].Symbol == "LENDUSDT" {
			list.Symbols = append(list.Symbols[:i], list.Symbols[i+1:]...)
			break
		}
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
