# 市场行情

## 获取服务器时间 
* curl --location --request GET 'http://127.0.0.1:8000/binance/market/time'
```
{"respCode":1,"respDesc":"SUCCESS","respData":1598599523041}
```


## 获取深度信息
* curl --location --request GET 'http://127.0.0.1:8000/binance/market/depth?symbol=BTCUSDT&limit=1'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{
      "lastUpdateId": 55396372838,
      "bids": [
          {
              "Price": "11407.68",
              "Quantity": "0.316"
          },
          {
              "Price": "11407.67",
              "Quantity": "0.130"
          },
          {
              "Price": "11407.65",
              "Quantity": "0.059"
          },
          {
              "Price": "11407.55",
              "Quantity": "0.015"
          },
          {
              "Price": "11407.49",
              "Quantity": "0.802"
          }
      ],
      "asks": [
          {
              "Price": "11408.14",
              "Quantity": "0.001"
          },
          {
              "Price": "11408.57",
              "Quantity": "0.054"
          },
          {
              "Price": "11408.88",
              "Quantity": "0.002"
          },
          {
              "Price": "11408.95",
              "Quantity": "0.024"
          },
          {
              "Price": "11409.13",
              "Quantity": "0.043"
          }
      ]
    }
}
```


## 获取近期成交
* curl --location --request GET 'http://127.0.0.1:8000/binance/market/aggTrades?symbol=BTCUSDT&limit=1'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[
      {
          "a": 135329287,
          "p": "11403.76",
          "q": "0.051",
          "f": 194733187,
          "l": 194733187,
          "T": 1598604457606,
          "m": true
      }
    ]
}
```


## 获取K线数据
* curl --location --request GET 'http://127.0.0.1:8000/binance/market/klines?symbol=BTCUSDT&interval=1d&limit=1'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[
      {
          "openTime": 1598572800000,
          "open": "11335.83",
          "high": "11447.62",
          "low": "11283.19",
          "close": "11397.97",
          "volume": "46303.742",
          "closeTime": 1598659199999,
          "quoteAssetVolume": "526823278.79653",
          "tradeNum": 198621,
          "takerBuyBaseAssetVolume": "22976.690",
          "takerBuyQuoteAssetVolume": "261440628.66481"
      }
    ]
}
```


## 最新标记价格和资金费率
* curl --location --request GET 'http://127.0.0.1:8000/binance/market/premiumIndex?symbol=BTCUSDT'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[
      {
          "symbol": "BTCUSDT",
          "markPrice": "11392.87000000",
          "lastFundingRate": "0.00010000",
          "nextFundingTime": 1598630400000,
          "time": 1598603486008
      }
    ]
}
```


## 24hr价格变动情况
* curl --location --request GET 'http://127.0.0.1:8000/binance/market/ticker/24hr?symbol=BTCUSDT'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[
      {
          "symbol": "BTCUSDT",
          "priceChange": "1.33",
          "priceChangePercent": "0.012",
          "weightedAvgPrice": "11345.61",
          "prevClosePrice": "",
          "lastPrice": "11391.37",
          "lastQty": "0.051",
          "openPrice": "11390.04",
          "highPrice": "11593.00",
          "lowPrice": "11130.82",
          "volume": "258164.086",
          "quoteVolume": "2929029243.36",
          "openTime": 1598517120000,
          "closeTime": 1598603557454,
          "firstId": 193872452,
          "lastId": 194727099,
          "count": 854646
      }
    ]
}
```


## 最新价格
* curl --location --request GET 'http://127.0.0.1:8000/binance/market/ticker/price?symbol=BTCUSDT'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[
      {
          "symbol": "BTCUSDT",
          "price": "11391.06"
      }
    ]
}
```


## 获取交易规则和交易对
* curl --location --request GET 'http://127.0.0.1:8000/binance/market/exchangeInfo?symbol=BTCUSDT'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{
      "timezone": "UTC",
      "serverTime": 1598603642625,
      "rateLimits": [
          {
              "rateLimitType": "REQUEST_WEIGHT",
              "interval": "MINUTE",
              "intervalNum": 1,
              "limit": 2400
          },
          {
              "rateLimitType": "ORDERS",
              "interval": "MINUTE",
              "intervalNum": 1,
              "limit": 1200
          }
      ],
      "exchangeFilters": [],
      "symbols": [
          {
              "symbol": "BTCUSDT",
              "status": "TRADING",
              "maintMarginPercent": "2.5000",
              "pricePrecision": 2,
              "quantityPrecision": 3,
              "requiredMarginPercent": "5.0000",
              "OrderType": null,
              "timeInForce": [
                  "GTC",
                  "IOC",
                  "FOK",
                  "GTX"
              ],
              "filters": [
                  {
                      "filterType": "PRICE_FILTER",
                      "maxPrice": "100000",
                      "minPrice": "0.01",
                      "tickSize": "0.01"
                  },
                  {
                      "filterType": "LOT_SIZE",
                      "maxQty": "1000",
                      "minQty": "0.001",
                      "stepSize": "0.001"
                  },
                  {
                      "filterType": "MARKET_LOT_SIZE",
                      "maxQty": "2419.848",
                      "minQty": "0.001",
                      "stepSize": "0.001"
                  },
                  {
                      "filterType": "MAX_NUM_ORDERS",
                      "limit": 0
                  },
                  {
                      "filterType": "PERCENT_PRICE",
                      "multiplierDecimal": "4",
                      "multiplierDown": "0.8500",
                      "multiplierUp": "1.1500"
                  }
              ]
          }, ...
      ]
    }
}
```
