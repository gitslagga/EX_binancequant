package trade

import (
	"EX_binancequant/config"
	"EX_binancequant/mylog"
	"EX_binancequant/trade/futures"
	"context"
	"fmt"
)

var BAExClient *Client
var BAExFuturesClient *futures.Client

func InitTrade() {
	var (
		apiKey          = config.Config.Trade.ApiKey
		secretKey       = config.Config.Trade.SecretKey
		endpoint        = config.Config.Trade.Endpoint
		futuresEndpoint = config.Config.Trade.FuturesEndpoint
	)

	BAExClient = NewClient(apiKey, secretKey, endpoint)
	BAExFuturesClient = NewFuturesClient(apiKey, secretKey, futuresEndpoint)

	fmt.Println("[InitTrade] binance success.")
}

func NewClientByParam(apiKey, secretKey string) *Client {
	var endpoint = config.Config.Trade.Endpoint

	return NewClient(apiKey, secretKey, endpoint)
}

func NewFuturesClientByParam(apiKey, secretKey string) *futures.Client {
	var futuresEndpoint = config.Config.Trade.FuturesEndpoint

	return NewFuturesClient(apiKey, secretKey, futuresEndpoint)
}

/**
Ping服务器服务
*/
func NewPingService() {
	err := BAExClient.NewPingService().Do(context.Background())
	if err != nil {
		mylog.Logger.Error().Msgf("[NewPingService] Do failed, err:%v", err)
		return
	}
}
