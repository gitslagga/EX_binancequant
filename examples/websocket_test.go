package examples

import (
	"EX_binancequant/trade/futures"
	"fmt"
	"testing"
	"time"
)

func Test_Depth(t *testing.T) {
	wsDepthHandler := func(event *futures.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := futures.WsDepthServe("BTCUSDT", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(9527 * time.Hour)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC
}

func Test_PartialDepth(t *testing.T) {
	wsDepthHandler := func(event *futures.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsPartialDepthServe("BTCUSDT", "5", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// remove this if you do not want to be blocked here
	<-doneC
}

func Test_CombinedPartial(t *testing.T) {
	wsDepthHandler := func(event *futures.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}

	symbolLevels := map[string]string{
		"BTCUSDT": "5",
		"ETHUSDT": "5",
	}
	doneC, _, err := futures.WsCombinedPartialDepthServe(symbolLevels, wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// remove this if you do not want to be blocked here
	<-doneC
}

func Test_Kline(t *testing.T) {
	wsKlineHandler := func(event *futures.WsKlineEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsKlineServe("BTCUSDT", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func Test_AggTrade(t *testing.T) {
	wsAggTradeHandler := func(event *futures.WsAggTradeEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsAggTradeServe("BTCUSDT", wsAggTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func Test_MarkPrice(t *testing.T) {
	wsMarkPriceHandler := func(event *futures.WsMarkPriceEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsMarkPriceServe("BTCUSDT", wsMarkPriceHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func Test_UserData(t *testing.T) {
	listenKey := "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"
	wsHandler := func(message []byte) {
		fmt.Println(string(message))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsUserDataServe(listenKey, wsHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
