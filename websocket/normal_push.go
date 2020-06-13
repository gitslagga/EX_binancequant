package websocket

import (
	"EX_binancequant/mylog"
	"EX_binancequant/trade/futures"
	"encoding/json"
)

func InitNormalPush(wsConn *wsConnection, symbol, levels, listenKey string) {
	go func() {
		PushAllMarkPrice(wsConn)
	}()
	go func() {
		PushKline(wsConn, symbol, levels)
	}()
	go func() {
		PushAggTrade(wsConn, symbol)
	}()
	go func() {
		PushAllMarketsStat(wsConn)
	}()
	go func() {
		PushDepth(wsConn, symbol)
	}()
	go func() {
		PushUserData(wsConn, listenKey)
	}()
}

func PushDepth(wsConn *wsConnection, symbol string) {
	wsDepthHandler := func(event *futures.WsDepthEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[PushDepth] json Marshal fail err: %v", err)
			return
		}

		if !wsConn.isClosed {
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

	if wsConn.isClosed {
		stopC <- struct{}{}
	}
}

func PushKlineCustom(wsConn *wsConnection, symbol, levels string) {
	wsKlineHandler := func(event *futures.WsKlineEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[PushKlineCustom] json Marshal fail err: %v", err)
			return
		}

		if !wsConn.isClosed {
			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushKlineCustom] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushKlineCustom] WsKlineServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsKlineServe(symbol, levels, wsKlineHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushKlineCustom] WsKlineServe dial fail err: %v", err)
		return
	}

	if wsConn.isClosed {
		stopC <- struct{}{}
	}
}

func PushKline(wsConn *wsConnection, symbol, levels string) {
	wsKlineHandler := func(event *futures.WsKlineEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[PushKline] json Marshal fail err: %v", err)
			return
		}

		if !wsConn.isClosed {
			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[PushKline] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[PushKline] WsKlineServe handler fail err: %v", err)
	}
	_, stopC, err := futures.WsKlineServe(symbol, levels, wsKlineHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[PushKline] WsKlineServe dial fail err: %v", err)
		return
	}

	if wsConn.isClosed {
		stopC <- struct{}{}
	}
}

func PushAggTrade(wsConn *wsConnection, symbol string) {
	wsAggTradeHandler := func(event *futures.WsAggTradeEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] json Marshal fail err: %v", err)
			return
		}

		if !wsConn.isClosed {
			err = wsConn.wsWrite(MessageType, response)
			if err != nil {
				mylog.DataLogger.Error().Msgf("[BA_CLIENT] write message fail err: %v", err)
			}
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushAggTrade handler fail err: %v", err)
	}
	_, stopC, err := futures.WsAggTradeServe(symbol, wsAggTradeHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushAggTrade dial fail err: %v", err)
		return
	}

	if wsConn.isClosed {
		stopC <- struct{}{}
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

	if wsConn.isClosed {
		stopC <- struct{}{}
	}
}

func PushAllMarkPrice(wsConn *wsConnection) {
	wsAllMarkPrice := func(event futures.WsAllMarkPriceEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[PushAllMarkPrice] json Marshal fail err: %v", err)
			return
		}

		if !wsConn.isClosed {
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

	if wsConn.isClosed {
		stopC <- struct{}{}
	}
}

func PushAllMarketsStat(wsConn *wsConnection) {
	wsAllMarketsStatHandler := func(event futures.WsAllMarketsStatEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[PushAllMarketsStat] json Marshal fail err: %v", err)
			return
		}

		if !wsConn.isClosed {
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

	if wsConn.isClosed {
		stopC <- struct{}{}
	}
}
