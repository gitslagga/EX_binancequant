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
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] json Marshal fail err: %v", err)
			return
		}

		err = wsConn.wsWrite(MessageType, response)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] write message fail err: %v", err)
			return
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushDepth handler fail err: %v", err)
	}
	_, _, err := futures.WsDepthServe(symbol, wsDepthHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushDepth dial fail err: %v", err)
		return
	}
}

func PushDepthLevel(wsConn *wsConnection, symbol, levels string) {
	wsDepthHandler := func(event *futures.WsDepthEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] json Marshal fail err: %v", err)
			return
		}

		err = wsConn.wsWrite(MessageType, response)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] write message fail err: %v", err)
			return
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushDepthLevel handler fail err: %v", err)
	}
	_, _, err := futures.WsPartialDepthServe(symbol, levels, wsDepthHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushDepthLevel dial fail err: %v", err)
		return
	}
}

func PushKline(wsConn *wsConnection, symbol, levels string) {
	wsKlineHandler := func(event *futures.WsKlineEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] json Marshal fail err: %v", err)
			return
		}

		err = wsConn.wsWrite(MessageType, response)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] write message fail err: %v", err)
			return
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushKline handler fail err: %v", err)
	}
	_, _, err := futures.WsKlineServe(symbol, levels, wsKlineHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushKline dial fail err: %v", err)
		return
	}
}

func PushAggTrade(wsConn *wsConnection, symbol string) {
	wsAggTradeHandler := func(event *futures.WsAggTradeEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] json Marshal fail err: %v", err)
			return
		}

		err = wsConn.wsWrite(MessageType, response)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] write message fail err: %v", err)
			return
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushAggTrade handler fail err: %v", err)
	}
	_, _, err := futures.WsAggTradeServe(symbol, wsAggTradeHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushAggTrade dial fail err: %v", err)
		return
	}
}

func PushUserData(wsConn *wsConnection, listenKey string) {
	//listenKey := "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"
	wsHandler := func(message []byte) {
		err := wsConn.wsWrite(MessageType, message)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] write message fail err: %v", err)
			return
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushUserData handler fail err: %v", err)
	}
	doneC, _, err := futures.WsUserDataServe(listenKey, wsHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushUserData dial fail err: %v", err)
		return
	}
	<-doneC
}

func PushAllMarkPrice(wsConn *wsConnection) {
	wsAllMarkPrice := func(event futures.WsAllMarkPriceEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] json Marshal fail err: %v", err)
			return
		}

		err = wsConn.wsWrite(MessageType, response)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] write message fail err: %v", err)
			return
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushAllMarkPrice handler fail err: %v", err)
	}
	doneC, _, err := futures.WsAllMarkPriceServe(wsAllMarkPrice, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushAllMarkPrice dial fail err: %v", err)
		return
	}
	<-doneC
}

func PushAllMarketsStat(wsConn *wsConnection) {
	wsAllMarketsStatHandler := func(event futures.WsAllMarketsStatEvent) {
		response, err := json.Marshal(event)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] json Marshal fail err: %v", err)
			return
		}

		err = wsConn.wsWrite(MessageType, response)
		if err != nil {
			mylog.DataLogger.Error().Msgf("[BA_CLIENT] write message fail err: %v", err)
			return
		}
	}
	errHandler := func(err error) {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushAllMarketsStat handler fail err: %v", err)
	}
	doneC, _, err := futures.WsAllMarketsStatServe(wsAllMarketsStatHandler, errHandler)
	if err != nil {
		mylog.DataLogger.Error().Msgf("[BA_CLIENT] PushAllMarketsStat dial fail err: %v", err)
		return
	}
	<-doneC
}
