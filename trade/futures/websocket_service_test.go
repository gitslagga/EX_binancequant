package futures

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type websocketServiceTestSuite struct {
	baseTestSuite
	origWsServe func(*WsConfig, WsHandler, ErrHandler) (chan struct{}, chan struct{}, error)
	serveCount  int
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(websocketServiceTestSuite))
}

func (s *websocketServiceTestSuite) SetupTest() {
	s.origWsServe = wsServe
}

func (s *websocketServiceTestSuite) TearDownTest() {
	wsServe = s.origWsServe
	s.serveCount = 0
}

func (s *websocketServiceTestSuite) mockWsServe(data []byte, err error) {
	wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, innerErr error) {
		s.serveCount++
		doneC = make(chan struct{})
		stopC = make(chan struct{})
		go func() {
			<-stopC
			close(doneC)
		}()
		handler(data)
		if err != nil {
			errHandler(err)
		}
		return doneC, stopC, nil
	}
}

func (s *websocketServiceTestSuite) assertWsServe(count ...int) {
	e := 1
	if len(count) > 0 {
		e = count[0]
	}
	s.r().Equal(e, s.serveCount)
}

func (s *websocketServiceTestSuite) TestDepthServe() {
	data := []byte(`{
        "e": "depthUpdate",
        "E": 1499404630606,
        "s": "ETHBTC",
        "u": 7913455,
        "U": 7913452,
        "pu": 7913450,
		"T": 1499404630606,
        "b": [
            [
                "0.10376590",
                "59.15767010",
                []
            ]
        ],
        "a": [
            [
                "0.10376586",
                "159.15767010",
                []
            ],
            [
                "0.10383109",
                "345.86845230",
                []
            ],
            [
                "0.10490700",
                "0.00000000",
                []
            ]
        ]
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsDepthServe("ETHBTC", func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:         "depthUpdate",
			Time:          1499404630606,
			Symbol:        "ETHBTC",
			FirstUpdateID: 7913452,
			LastUpdateID:  7913455,
			PreviousID:    7913450,
			TradeTime:     1499404630606,
			Bids: []Bid{
				{
					Price:    "0.10376590",
					Quantity: "59.15767010",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.10376586",
					Quantity: "159.15767010",
				},
				{
					Price:    "0.10383109",
					Quantity: "345.86845230",
				},
				{
					Price:    "0.10490700",
					Quantity: "0.00000000",
				},
			},
		}
		s.assertWsDepthEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsDepthEventEqual(e, a *WsDepthEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.FirstUpdateID, a.FirstUpdateID, "FirstUpdateID")
	r.Equal(e.LastUpdateID, a.LastUpdateID, "LastUpdateID")
	r.Equal(e.PreviousID, a.PreviousID, "PreviousID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	for i := 0; i < len(e.Bids); i++ {
		r.Equal(e.Bids[i].Price, a.Bids[i].Price, "Price")
		r.Equal(e.Bids[i].Quantity, a.Bids[i].Quantity, "Quantity")
	}
	for i := 0; i < len(e.Asks); i++ {
		r.Equal(e.Asks[i].Price, a.Asks[i].Price, "Price")
		r.Equal(e.Asks[i].Quantity, a.Asks[i].Quantity, "Quantity")
	}
}

func (s *websocketServiceTestSuite) TestPartialDepthServe() {
	data := []byte(`{
        "e": "depthUpdate",
        "E": 1499404630606,
        "s": "ETHBTC",
        "u": 7913455,
        "U": 7913452,
        "pu": 7913450,
		"T": 1499404630606,
        "b": [
            [
                "0.10376590",
                "59.15767010",
                []
            ]
        ],
        "a": [
            [
                "0.10376586",
                "159.15767010",
                []
            ],
            [
                "0.10383109",
                "345.86845230",
                []
            ],
            [
                "0.10490700",
                "0.00000000",
                []
            ]
        ]
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsPartialDepthServe("ETHBTC", "5", func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:         "depthUpdate",
			Time:          1499404630606,
			Symbol:        "ETHBTC",
			FirstUpdateID: 7913452,
			LastUpdateID:  7913455,
			PreviousID:    7913450,
			TradeTime:     1499404630606,
			Bids: []Bid{
				{
					Price:    "0.10376590",
					Quantity: "59.15767010",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.10376586",
					Quantity: "159.15767010",
				},
				{
					Price:    "0.10383109",
					Quantity: "345.86845230",
				},
				{
					Price:    "0.10490700",
					Quantity: "0.00000000",
				},
			},
		}
		s.assertWsPartialDepthEventEqual(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestCombinedPartialDepthServe() {
	data := []byte(`{
      "stream":"ethusdt@depth5",
      "data": {
        "e": "depthUpdate",
        "E": 1499404630606,
        "s": "ETHBTC",
        "u": 7913455,
        "U": 7913452,
        "pu": 7913450,
		"T": 1499404630606,
        "b": [
            [
                "0.10376590",
                "59.15767010",
                []
            ]
        ],
        "a": [
            [
                "0.10376586",
                "159.15767010",
                []
            ],
            [
                "0.10383109",
                "345.86845230",
                []
            ],
            [
                "0.10490700",
                "0.00000000",
                []
            ]
        ]
      }
	}`)
	symbolLevels := map[string]string{
		"BTCUSDT": "5",
		"ETHUSDT": "5",
	}
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()
	doneC, stopC, err := WsCombinedPartialDepthServe(symbolLevels, func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:         "depthUpdate",
			Time:          1499404630606,
			Symbol:        "ETHBTC",
			FirstUpdateID: 7913452,
			LastUpdateID:  7913455,
			PreviousID:    7913450,
			TradeTime:     1499404630606,
			Bids: []Bid{
				{
					Price:    "0.10376590",
					Quantity: "59.15767010",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.10376586",
					Quantity: "159.15767010",
				},
				{
					Price:    "0.10383109",
					Quantity: "345.86845230",
				},
				{
					Price:    "0.10490700",
					Quantity: "0.00000000",
				},
			},
		}
		s.assertWsPartialDepthEventEqual(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsPartialDepthEventEqual(e, a *WsDepthEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.FirstUpdateID, a.FirstUpdateID, "FirstUpdateID")
	r.Equal(e.LastUpdateID, a.LastUpdateID, "LastUpdateID")
	r.Equal(e.PreviousID, a.PreviousID, "PreviousID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	for i := 0; i < len(e.Bids); i++ {
		r.Equal(e.Bids[i].Price, a.Bids[i].Price, "Price")
		r.Equal(e.Bids[i].Quantity, a.Bids[i].Quantity, "Quantity")
	}
	for i := 0; i < len(e.Asks); i++ {
		r.Equal(e.Asks[i].Price, a.Asks[i].Price, "Price")
		r.Equal(e.Asks[i].Quantity, a.Asks[i].Quantity, "Quantity")
	}
}

func (s *websocketServiceTestSuite) TestKlineServe() {
	data := []byte(`{
        "e": "kline",
        "E": 1499404907056,
        "s": "ETHBTC",
        "k": {
            "t": 1499404860000,
            "T": 1499404919999,
            "s": "ETHBTC",
            "i": "1m",
            "f": 77462,
            "L": 77465,
            "o": "0.10278577",
            "c": "0.10278645",
            "h": "0.10278712",
            "l": "0.10278518",
            "v": "17.47929838",
            "n": 4,
            "x": false,
            "q": "1.79662878",
            "V": "2.34879839",
            "Q": "0.24142166",
            "B": "13279784.01349473"
        }
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsKlineServe("ETHBTC", "1m", func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   1499404907056,
			Symbol: "ETHBTC",
			Kline: WsKline{
				StartTime:            1499404860000,
				EndTime:              1499404919999,
				Symbol:               "ETHBTC",
				Interval:             "1m",
				FirstTradeID:         77462,
				LastTradeID:          77465,
				Open:                 "0.10278577",
				Close:                "0.10278645",
				High:                 "0.10278712",
				Low:                  "0.10278518",
				Volume:               "17.47929838",
				TradeNum:             4,
				IsFinal:              false,
				QuoteVolume:          "1.79662878",
				ActiveBuyVolume:      "2.34879839",
				ActiveBuyQuoteVolume: "0.24142166",
			},
		}
		s.assertWsKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsKlineEventEqual(e, a *WsKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.EndTime, ak.EndTime, "EndTime")
	r.Equal(ek.Symbol, ak.Symbol, "Symbol")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.FirstTradeID, ak.FirstTradeID, "FirstTradeID")
	r.Equal(ek.LastTradeID, ak.LastTradeID, "LastTradeID")
	r.Equal(ek.Open, ak.Open, "Open")
	r.Equal(ek.Close, ak.Close, "Close")
	r.Equal(ek.High, ak.High, "High")
	r.Equal(ek.Low, ak.Low, "Low")
	r.Equal(ek.Volume, ak.Volume, "Volume")
	r.Equal(ek.TradeNum, ak.TradeNum, "TradeNum")
	r.Equal(ek.IsFinal, ak.IsFinal, "IsFinal")
	r.Equal(ek.QuoteVolume, ak.QuoteVolume, "QuoteVolume")
	r.Equal(ek.ActiveBuyVolume, ak.ActiveBuyVolume, "ActiveBuyVolume")
	r.Equal(ek.ActiveBuyQuoteVolume, ak.ActiveBuyQuoteVolume, "ActiveBuyQuoteVolume")
}

func (s *websocketServiceTestSuite) TestWsAggTradeServe() {
	data := []byte(`{
        "e": "aggTrade",
        "E": 1499405254326,
        "s": "ETHBTC",
        "a": 70232,
        "p": "0.10281118",
        "q": "8.15632997",
        "f": 77489,
        "l": 77489,
        "T": 1499405254324,
        "m": false,
        "M": true
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAggTradeServe("ETHBTC", func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:                 "aggTrade",
			Time:                  1499405254326,
			Symbol:                "ETHBTC",
			AggTradeID:            70232,
			Price:                 "0.10281118",
			Quantity:              "8.15632997",
			FirstBreakdownTradeID: 77489,
			LastBreakdownTradeID:  77489,
			TradeTime:             1499405254324,
			IsBuyerMaker:          false,
		}
		s.assertWsAggTradeEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAggTradeEventEqual(e, a *WsAggTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.AggTradeID, a.AggTradeID, "AggTradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.FirstBreakdownTradeID, a.FirstBreakdownTradeID, "FirstBreakdownTradeID")
	r.Equal(e.LastBreakdownTradeID, a.LastBreakdownTradeID, "LastBreakdownTradeID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
}

func (s *websocketServiceTestSuite) testWsUserDataServe(data []byte) {
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsUserDataServe("listenKey", func(event []byte) {
		s.r().Equal(data, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

//listenKey 过期推送
func (s *websocketServiceTestSuite) TestWsUserDataServeListen() {
	s.testWsUserDataServe([]byte(`{
		'e': 'listenKeyExpired',      // 事件类型
		'E': 1576653824250            // 事件时间
	}`))
}

//追加保证金通知
func (s *websocketServiceTestSuite) TestWsUserDataServeMargin() {
	s.testWsUserDataServe([]byte(`{
		"e":"MARGIN_CALL",      // 事件类型
		"E":1587727187525,      // 事件时间
		"cw":"3.16812045",      // 除去逐仓仓位保证金的钱包余额, 仅在全仓 margin call 情况下推送此字段
		"p":[                   // 涉及持仓
		  {
			"s":"ETHUSDT",      // symbol
			"ps":"LONG",        // 持仓方向
			"pa":"1.327",       // 仓位
			"mt":"CROSSED",     // 保证金模式
			"iw":"0",           // 若为逐仓，仓位保证金
			"mp":"187.17127",   // 标记价格
			"up":"-1.166074",   // 未实现盈亏
			"mm":"1.614445"     // 持仓需要的维持保证金
		  }
		]
	}`))
}

//Balance和Position更新推送
func (s *websocketServiceTestSuite) TestWsUserDataServeAccount() {
	s.testWsUserDataServe([]byte(`{
	  "e": "ACCOUNT_UPDATE",                // 事件类型
	  "E": 1564745798939,                   // 事件时间
	  "T": 1564745798938 ,                  // 撮合时间
	  "a":                                  // 账户更新事件
		{
		  "m":"ORDER",                      // 事件推出原因 
		  "B":[                             // 余额信息
			{
			  "a":"USDT",                   // 资产名称
			  "wb":"122624.12345678",       // 钱包余额
			  "cw":"100.12345678"           // 除去逐仓仓位保证金的钱包余额
			},
		  ],
		  "P":[
		   {
			  "s":"BTCUSDT",            // 交易对
			  "pa":"0",                 // 仓位
			  "ep":"0.00000",            // 入仓价格
			  "cr":"200",               // (费前)累计实现损益
			  "up":"0",                     // 持仓未实现盈亏
			  "mt":"isolated",              // 保证金模式
			  "iw":"0.00000000",            // 若为逐仓，仓位保证金
			  "ps":"BOTH"                   // 持仓方向
		   },
		  ]
		}
	}`))
}

//Balance和Position更新推送
func (s *websocketServiceTestSuite) TestWsUserDataServeOrderTrade() {
	s.testWsUserDataServe([]byte(`{
	  "e":"ORDER_TRADE_UPDATE",         // 事件类型
	  "E":1568879465651,                // 事件时间
	  "T":1568879465650,                // 撮合时间
	  "o":{                             
		"s":"BTCUSDT",                  // 交易对
		"c":"TEST",                     // 客户端自定订单ID
		  // 特殊的自定义订单ID:
		  // "autoclose-"开头的字符串: 系统强平订单
		  // "adl_autoclose": ADL自动减仓订单
		"S":"SELL",                     // 订单方向
		"o":"TRAILING_STOP_MARKET", // 订单类型
		"f":"GTC",                      // 有效方式
		"q":"0.001",                    // 订单原始数量
		"p":"0",                        // 订单原始价格
		"ap":"0",                       // 订单平均价格
		"sp":"7103.04",                 // 条件订单触发价格，对追踪止损单无效
		"x":"NEW",                      // 本次事件的具体执行类型
		"X":"NEW",                      // 订单的当前状态
		"i":8886774,                    // 订单ID
		"l":"0",                        // 订单末次成交数量
		"z":"0",                        // 订单累计已成交数量
		"L":"0",                        // 订单末次成交价格
		"N": "USDT",                    // 手续费资产类型
		"n": "0",                       // 手续费数量
		"T":1568879465651,              // 成交时间
		"t":0,                          // 成交ID
		"b":"0",                        // 买单净值
		"a":"9.91",                     // 卖单净值
		"m": false,                     // 该成交是作为挂单成交吗？
		"R":false   ,                   // 是否是只减仓单
		"wt": "CONTRACT_PRICE",         // 触发价类型
		"ot": "TRAILING_STOP_MARKET",   // 原始订单类型
		"ps":"LONG"                     // 持仓方向
		"cp":false,                     // 是否为触发平仓单; 仅在条件订单情况下会推送此字段
		"AP":"7476.89",                 // 追踪止损激活价格, 仅在追踪止损单时会推送此字段
		"cr":"5.0",                     // 追踪止损回调比例, 仅在追踪止损单时会推送此字段
		"rp":"0"                        // 该交易实现盈亏
	  }
	}`))
}

func (s *websocketServiceTestSuite) TestWsMarkPriceServe() {
	data := []byte(`{
		"e": "markPriceUpdate",
		"E": 1562305380000,
		"s": "BTCUSDT",
		"p": "11185.87786614",
		"r": "0.00030000",
		"T": 1562306400000
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarkPriceServe("BTCUSDT", func(event *WsMarkPriceEvent) {
		e := &WsMarkPriceEvent{
			Event:       "markPriceUpdate", // 事件类型
			Time:        1562305380000,     // 事件时间
			Symbol:      "BTCUSDT",         // 交易对
			Price:       "11185.87786614",  // 标记价格
			FundingRate: "0.00030000",      // 资金费率
			TradeTime:   1562306400000,     // 下个资金时间
		}

		s.assertWsMarkPriceEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarkPriceEventEqual(e, a *WsMarkPriceEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.FundingRate, a.FundingRate, "FundingRate")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
}

func (s *websocketServiceTestSuite) TestWsAllMarkPriceServe() {
	data := []byte(`[{
		"e": "markPriceUpdate",
		"E": 1562305380000,
		"s": "BTCUSDT",
		"p": "11185.87786614",
		"r": "0.00030000",
		"T": 1562306400000
	},{
		"e": "markPriceUpdate",
		"E": 1562305380000,
		"s": "BTCUSDT",
		"p": "11185.87786614",
		"r": "0.00030000",
		"T": 1562306400000
	}]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMarkPriceServe(func(event WsAllMarkPriceEvent) {
		e := WsAllMarkPriceEvent{
			&WsMarkPriceEvent{
				Event:       "markPriceUpdate",
				Time:        1562305380000,
				Symbol:      "BTCUSDT",
				Price:       "11185.87786614",
				FundingRate: "0.00030000",
				TradeTime:   1562306400000,
			},
			&WsMarkPriceEvent{
				Event:       "markPriceUpdate",
				Time:        1562305380000,
				Symbol:      "BTCUSDT",
				Price:       "11185.87786614",
				FundingRate: "0.00030000",
				TradeTime:   1562306400000,
			},
		}
		s.assertWsAllMarkPriceEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAllMarkPriceEventEqual(e, a WsAllMarkPriceEvent) {
	for i := range e {
		s.assertWsMarkPriceEventEqual(e[i], a[i])
	}
}

func (s *websocketServiceTestSuite) TestWsMarketStatServe() {
	data := []byte(`{
  		"e": "24hrTicker",
  		"E": 123456789,
  		"s": "BNBBTC",
  		"p": "0.0015",
  		"P": "250.00",
  		"w": "0.0018",
  		"c": "0.0025",
  		"Q": "10",
  		"o": "0.0010",
  		"h": "0.0026",
  		"l": "0.0010",
  		"v": "10000",
  		"q": "18",
 		"O": 0,
  		"C": 86400000,
  		"F": 0,
  		"L": 18150,
  		"n": 18151
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarketStatServe("BNBBTC", func(event *WsMarketStatEvent) {
		e := &WsMarketStatEvent{
			Event:              "24hrTicker",
			Time:               123456789,
			Symbol:             "BNBBTC",
			PriceChange:        "0.0015",
			PriceChangePercent: "250.00",
			WeightedAvgPrice:   "0.0018",
			LastPrice:          "0.0025",
			CloseQty:           "10",
			OpenPrice:          "0.0010",
			HighPrice:          "0.0026",
			LowPrice:           "0.0010",
			BaseVolume:         "10000",
			QuoteVolume:        "18",
			OpenTime:           0,
			CloseTime:          86400000,
			FirstID:            0,
			LastID:             18150,
			Count:              18151,
		}
		s.assertWsMarketStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarketStatEventEqual(e, a *WsMarketStatEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.CloseQty, a.CloseQty, "CloseQty")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.BaseVolume, a.BaseVolume, "BaseVolume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstID, a.FirstID, "FirstID")
	r.Equal(e.LastID, a.LastID, "LastID")
	r.Equal(e.Count, a.Count, "Count")
}

func (s *websocketServiceTestSuite) TestWsAllMarketsStatServe() {
	data := []byte(`[{
  		"e": "24hrTicker",
  		"E": 123456789,
  		"s": "BNBBTC",
  		"p": "0.0015",
  		"P": "250.00",
  		"w": "0.0018",
  		"c": "0.0025",
  		"Q": "10",
  		"o": "0.0010",
  		"h": "0.0026",
  		"l": "0.0010",
  		"v": "10000",
  		"q": "18",
 		"O": 0,
  		"C": 86400000,
  		"F": 0,
  		"L": 18150,
  		"n": 18151
	},{
  		"e": "24hrTicker",
  		"E": 123456789,
  		"s": "ETHBTC",
  		"p": "0.0015",
  		"P": "250.00",
  		"w": "0.0018",
  		"c": "0.0025",
  		"Q": "10",
  		"o": "0.0010",
  		"h": "0.0026",
  		"l": "0.0010",
  		"v": "10000",
  		"q": "18",
 		"O": 0,
  		"C": 86400000,
  		"F": 0,
  		"L": 18150,
  		"n": 18151
	}]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMarketsStatServe(func(event WsAllMarketsStatEvent) {
		e := WsAllMarketsStatEvent{
			&WsMarketStatEvent{
				Event:              "24hrTicker",
				Time:               123456789,
				Symbol:             "BNBBTC",
				PriceChange:        "0.0015",
				PriceChangePercent: "250.00",
				WeightedAvgPrice:   "0.0018",
				LastPrice:          "0.0025",
				CloseQty:           "10",
				OpenPrice:          "0.0010",
				HighPrice:          "0.0026",
				LowPrice:           "0.0010",
				BaseVolume:         "10000",
				QuoteVolume:        "18",
				OpenTime:           0,
				CloseTime:          86400000,
				FirstID:            0,
				LastID:             18150,
				Count:              18151,
			},
			&WsMarketStatEvent{
				Event:              "24hrTicker",
				Time:               123456789,
				Symbol:             "ETHBTC",
				PriceChange:        "0.0015",
				PriceChangePercent: "250.00",
				WeightedAvgPrice:   "0.0018",
				LastPrice:          "0.0025",
				CloseQty:           "10",
				OpenPrice:          "0.0010",
				HighPrice:          "0.0026",
				LowPrice:           "0.0010",
				BaseVolume:         "10000",
				QuoteVolume:        "18",
				OpenTime:           0,
				CloseTime:          86400000,
				FirstID:            0,
				LastID:             18150,
				Count:              18151,
			},
		}
		s.assertWsAllMarketsStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAllMarketsStatEventEqual(e, a WsAllMarketsStatEvent) {
	for i := range e {
		s.assertWsMarketStatEventEqual(e[i], a[i])
	}
}
