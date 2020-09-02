package examples

import (
	"EX_binancequant/trade/futures"
	"fmt"
	"github.com/gorilla/websocket"
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

func Test_MarksPrice(t *testing.T) {
	wsAllMarkPriceHandler := func(event futures.WsAllMarkPriceEvent) {
		fmt.Println(event[0])
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsAllMarkPriceServe(wsAllMarkPriceHandler, errHandler)
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

func Test_MarketStat(t *testing.T) {
	wsMarketStatHandler := func(event *futures.WsMarketStatEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsMarketStatServe("BTCUSDT", wsMarketStatHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func Test_MarketsStat(t *testing.T) {
	wsAllMarketsStatHandler := func(event futures.WsAllMarketsStatEvent) {
		fmt.Println(event[0])
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsAllMarketsStatServe(wsAllMarketsStatHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func Test_BinanceStream(t *testing.T) {
	ws := &wsConnection{Data: "TestBinanceStream"}

	go func() {
		ws.binanceRequest([]byte(`{
		"method": "SUBSCRIBE",
		"params":
		[
			"btcusdt@aggTrade",
			"btcusdt@depth"
		],
		"id": 1
	}`))
	}()

	time.Sleep(60 * time.Second)

	go func() {
		ws.binanceRequest([]byte(`{
			"method": "UNSUBSCRIBE",
			"params":
			[
			"btcusdt@aggTrade",
			"btcusdt@depth"
			],
			"id": 312
		}`))
	}()

	time.Sleep(60 * time.Second)
}

type wsConnection struct {
	Data string
}

var entityChannel = make(map[*wsConnection]*websocket.Conn, 10000)

func (w *wsConnection) binanceRequest(message []byte) (stopC chan struct{}) {
	if _, ok := entityChannel[w]; ok {
		fmt.Printf("len entityChannel: %v \n", len(entityChannel))
		entityChannel[w].WriteMessage(websocket.TextMessage, message)
		return
	}

	c, _, err := websocket.DefaultDialer.Dial("wss://fstream.binancezh.com/stream", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	entityChannel[w] = c
	c.WriteMessage(websocket.TextMessage, message)

	stopC = make(chan struct{})

	go func() {
		defer c.Close()

		for {
			select {
			case <-stopC:
				return
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(message))
			}
		}
	}()

	for {
		time.Sleep(5 * time.Second)

		fmt.Printf("[Websocket] conn entityChannel data: %v \n", w)
		fmt.Printf("[Websocket] conn entityChannel data: %v \n", entityChannel[w])
	}

	fmt.Println("binanceRequest finished")

	return
}
