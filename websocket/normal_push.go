package websocket

import (
	"EX_binancequant/mylog"
	"EX_binancequant/trade/futures"
	"encoding/json"
	"time"
)

func PushDepth(wsConn *wsConnection, symbol string) {
	wsDepthHandler := func(event *futures.WsDepthEvent) {
		if !wsConn.isClosed {
			response, err := json.Marshal(event)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushDepth] json Marshal fail err: %v", err)
				return
			}

			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushDepth] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushDepth] WsDepthServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsDepthServe(symbol, wsDepthHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushDepth] WsDepthServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}

func PushDepthLevels(wsConn *wsConnection, symbol, levels string) {
	wsDepthHandler := func(event *futures.WsDepthEvent) {
		if !wsConn.isClosed {
			response, err := json.Marshal(event)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushDepthLevel] json Marshal fail err: %v", err)
				return
			}

			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushDepthLevel] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushDepthLevel] WsPartialDepthServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsPartialDepthServe(symbol, levels, wsDepthHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushDepthLevel] WsPartialDepthServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}

func PushKline(wsConn *wsConnection, symbol, interval string) {
	wsKlineHandler := func(event *futures.WsKlineEvent) {
		if !wsConn.isClosed {
			response, err := json.Marshal(event)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushKline] json Marshal fail err: %v", err)
				return
			}

			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushKline] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushKline] WsKlineServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsKlineServe(symbol, interval, wsKlineHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushKline] WsKlineServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}

func PushKlineInterval(wsConn *wsConnection, symbol, interval string) {
	wsKlineHandler := func(event *futures.WsKlineEvent) {
		if !wsConn.isClosed {
			response, err := json.Marshal(event)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushKlineInterval] json Marshal fail err: %v", err)
				return
			}

			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushKlineInterval] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushKlineInterval] WsKlineServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsKlineServe(symbol, interval, wsKlineHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushKlineInterval] WsKlineServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}

func PushAggTrade(wsConn *wsConnection, symbol string) {
	wsAggTradeHandler := func(event *futures.WsAggTradeEvent) {
		if !wsConn.isClosed {
			response, err := json.Marshal(event)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushAggTrade] json Marshal fail err: %v", err)
				return
			}

			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushAggTrade] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushAggTrade] WsAggTradeServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsAggTradeServe(symbol, wsAggTradeHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushAggTrade] WsAggTradeServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}

func PushUserData(wsConn *wsConnection, listenKey string) {
	//listenKey := "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"
	wsHandler := func(message []byte) {
		if !wsConn.isClosed {
			err := wsConn.wsWrite(MessageType, message)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushUserData] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushUserData] WsUserDataServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsUserDataServe(listenKey, wsHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushUserData] WsUserDataServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}

func PushAllMarkPrice(wsConn *wsConnection) {
	wsAllMarkPrice := func(event futures.WsAllMarkPriceEvent) {
		if !wsConn.isClosed {
			response, err := json.Marshal(event)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushAllMarkPrice] json Marshal fail err: %v", err)
				return
			}

			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushAllMarkPrice] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushAllMarkPrice] WsAllMarkPriceServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsAllMarkPriceServe(wsAllMarkPrice, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushAllMarkPrice] WsAllMarkPriceServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}

func PushAllMarketsStat(wsConn *wsConnection) {
	wsAllMarketsStatHandler := func(event futures.WsAllMarketsStatEvent) {
		if !wsConn.isClosed {
			response, err := json.Marshal(event)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushAllMarketsStat] json Marshal fail err: %v", err)
				return
			}

			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushAllMarketsStat] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushAllMarketsStat] WsAllMarketsStatServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsAllMarketsStatServe(wsAllMarketsStatHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushAllMarketsStat] WsAllMarketsStatServe dial fail err: %v", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)
		if wsConn.isClosed {
			stopC <- struct{}{}
			break
		}
	}
}
