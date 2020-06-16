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
		debug           = config.Config.Trade.Debug
		futuresDebug    = config.Config.Trade.FuturesDebug
	)

	BAExClient = NewClient(apiKey, secretKey, endpoint, debug)
	BAExFuturesClient = NewFuturesClient(apiKey, secretKey, futuresEndpoint, futuresDebug)

	fmt.Println("[InitTrade] binance success.")
}

func NewClientByParam(apiKey, secretKey string) *Client {
	var (
		endpoint = config.Config.Trade.Endpoint
		debug    = config.Config.Trade.Debug
	)

	return NewClient(apiKey, secretKey, endpoint, debug)
}

func NewFuturesClientByParam(apiKey, secretKey string) *futures.Client {
	var (
		futuresEndpoint = config.Config.Trade.FuturesEndpoint
		futuresDebu     = config.Config.Trade.FuturesDebug
	)

	return NewFuturesClient(apiKey, secretKey, futuresEndpoint, futuresDebu)
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
