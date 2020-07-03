# My Version

## listenKey, 生成，延长，关闭
* curl --location --request POST 'http://47.57.93.231:8000/api/futures/listenKey' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok","data":"kOpLJN4gXrycDSKlewIlfV5mKnVYWTXlBnaTUCXNOS2QuObcRG3udSt62rnorgce"}
```
* curl --location --request PUT 'http://47.57.93.231:8000/api/futures/listenKey' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok"}
```
* curl --location --request DELETE 'http://47.57.93.231:8000/api/futures/listenKey' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok"}
```

## 更改持仓模式
* curl --location --request POST 'http://47.57.93.231:8000/api/futures/position/mode' --header 'Content-Type: application/json' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf' --data-raw '{"dualSidePosition":true}'
```
{"error_code":0,"error_message":"ok"}
```


## 获取持仓模式
* curl --location --request GET 'http://47.57.93.231:8000/api/futures/position/mode' --header 'Content-Type: application/json' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok","data":{"dualSidePosition":true}}
```


## 下单, 平仓
* curl --location --request POST 'http://47.57.93.231:8000/api/futures/order' --header 'Content-Type: application/json' --header 'Content-Type: application/json' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf' --data-raw '{"symbol": "BTCUSDT", "side": "BUY", "positionSide":"BOTH", "type": "LIMIT", "reduceOnly": false, "quantity": 10, "price": 5000, "stopPrice": 0,	"closePosition": false, "activationPrice": 1000, "callbackRate": 0.1, "timeInForce": "GTC", "workingType": "CONTRACT_PRICE",	"newOrderRespType":"ACK"}'
```
{"error_code":0,"error_message":"ok","data":{"symbol":"BTCUSDT","orderId":2494226119,"clientOrderId":"testOrder","price":"10000","origQty":"10","executedQty":"0","cumQuote":"0","reduceOnly":false,"status":"NEW","stopPrice":"0","timeInForce":"GTC","type":"LIMIT","side":"SELL","updateTime":1592192005711,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","positionSide":"BOTH"}}
```
* curl --location --request POST 'http://47.57.93.231:8000/api/futures/order' --header 'Content-Type: application/json' --header 'Content-Type: application/json' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf' --data-raw '{"symbol": "BTCUSDT", "side": "BUY", "positionSide":"BOTH", "type": "LIMIT", "reduceOnly": false, "quantity": 10, "price": 10000, "stopPrice": 0,	"closePosition": false, "activationPrice": 1000, "callbackRate": 0.1, "timeInForce": "GTC", "workingType": "CONTRACT_PRICE",	"newOrderRespType":"ACK"}'
```
{"error_code":0,"error_message":"ok","data":{"symbol":"BTCUSDT","orderId":2494263524,"clientOrderId":"buyOrder","price":"10000","origQty":"10","executedQty":"0","cumQuote":"0","reduceOnly":false,"status":"NEW","stopPrice":"0","timeInForce":"GTC","type":"LIMIT","side":"BUY","updateTime":1592192923649,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","positionSide":"BOTH"}}
```
* curl --location --request POST 'http://47.57.93.231:8000/api/futures/order' --header 'Content-Type: application/json' --header 'Content-Type: application/json' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf' --data-raw '{"symbol": "BTCUSDT", "side": "SELL", "positionSide":"BOTH", "type": "MARKET", "quantity": 10, "closePosition": false, "activationPrice": 1000, "callbackRate": 0.1, "workingType": "CONTRACT_PRICE",	"newOrderRespType":"ACK"}'
 ```
{"error_code":0,"error_message":"ok","data":{"symbol":"BTCUSDT","orderId":2494264483,"clientOrderId":"stopOrder","price":"0","origQty":"0","executedQty":"0","cumQuote":"0","reduceOnly":true,"status":"NEW","stopPrice":"9200","timeInForce":"GTC","type":"STOP_MARKET","side":"BUY","updateTime":1592194582808,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","positionSide":"BOTH"}}
```


## 查询订单
* curl --location --request GET 'http://47.57.93.231:8000/api/futures/order?symbol=BTCUSDT&orderId=2494226119' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok","data":{"symbol":"BTCUSDT","orderId":2494226119,"clientOrderId":"testOrder","price":"10000","reduceOnly":false,"origQty":"10","executedQty":"0","cumQty":"","cumQuote":"0","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0","time":1592192005711,"updateTime":1592192005711,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","origType":"LIMIT","positionSide":"BOTH","closePosition":false}}
```


## 撤销订单, 撤销所有订单
* curl --location --request DELETE 'http://47.57.93.231:8000/api/futures/order?symbol=BTCUSDT&orderId=2494454324' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok","data":{"clientOrderId":"testOrder","cumQty":"0","cumQuote":"0","executedQty":"0","orderId":2494226119,"origQty":"10","price":"10000","reduceOnly":false,"side":"SELL","status":"CANCELED","stopPrice":"0","symbol":"BTCUSDT","timeInForce":"GTC","type":"LIMIT","updateTime":1592193302168,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","origType":"LIMIT","positionSide":"BOTH"}}
```
* curl --location --request DELETE 'http://47.57.93.231:8000/api/futures/allOpenOrders?symbol=BTCUSDT' --header 'Content-Type: application/json' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok"}
```


## 查看当前所有挂单
* curl --location --request GET 'http://47.57.93.231:8000/api/futures/openOrders?symbol=BTCUSDT' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok","data":[{"symbol":"BTCUSDT","orderId":2494253712,"clientOrderId":"sellOrder","price":"10000","reduceOnly":false,"origQty":"10","executedQty":"0","cumQty":"","cumQuote":"0","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0","time":1592193816956,"updateTime":1592193854490,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","origType":"LIMIT","positionSide":"BOTH","closePosition":false}]}
```


## 查看所有订单
* curl --location --request GET 'http://47.57.93.231:8000/api/futures/allOrders?symbol=BTCUSDT&orderId=10000&limit=1000' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok","data":[{"symbol":"BTCUSDT","orderId":2486612018,"clientOrderId":"web_JnuPkxByhvuHNKeKt48z","price":"0","reduceOnly":false,"origQty":"1","executedQty":"1","cumQty":"","cumQuote":"9379.94346","status":"FILLED","timeInForce":"GTC","type":"MARKET","side":"BUY","stopPrice":"0","time":1591597901305,"updateTime":1591597901325,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"9379.94346","origType":"MARKET","positionSide":"BOTH","closePosition":false},{"symbol":"BTCUSDT","orderId":2486612063,"clientOrderId":"web_lQHOY9PUW8H8su7YgWEB","price":"8000","reduceOnly":false,"origQty":"1","executedQty":"0","cumQty":"","cumQuote":"0","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"BUY","stopPrice":"0","time":1591598029709,"updateTime":1591775938126,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","origType":"LIMIT","positionSide":"BOTH","closePosition":false},{"symbol":"BTCUSDT","orderId":2487922001,"clientOrderId":"XPdtJTgcyCFfEm9fIdRucl","price":"9740.79","reduceOnly":true,"origQty":"1","executedQty":"1","cumQty":"","cumQuote":"9740.79000","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0","time":1591775938128,"updateTime":1591775938154,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"9740.79000","origType":"LIMIT","positionSide":"BOTH","closePosition":false},{"symbol":"BTCUSDT","orderId":2494226119,"clientOrderId":"testOrder","price":"10000","reduceOnly":false,"origQty":"10","executedQty":"0","cumQty":"","cumQuote":"0","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0","time":1592192005711,"updateTime":1592193302169,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","origType":"LIMIT","positionSide":"BOTH","closePosition":false},{"symbol":"BTCUSDT","orderId":2494240266,"clientOrderId":"buyOrder","price":"10000","reduceOnly":false,"origQty":"10","executedQty":"10","cumQty":"","cumQuote":"91847.74052","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"BUY","stopPrice":"0","time":1592192923649,"updateTime":1592192923668,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"9184.77405","origType":"LIMIT","positionSide":"BOTH","closePosition":false},{"symbol":"BTCUSDT","orderId":2494242249,"clientOrderId":"web_lVcDzjYat84MgS4an4c2","price":"0","reduceOnly":true,"origQty":"10","executedQty":"10","cumQty":"","cumQuote":"91849.29682","status":"FILLED","timeInForce":"GTC","type":"MARKET","side":"SELL","stopPrice":"0","time":1592193065134,"updateTime":1592193065191,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"9184.92968","origType":"MARKET","positionSide":"BOTH","closePosition":false},{"symbol":"BTCUSDT","orderId":2494247226,"clientOrderId":"sellOrder","price":"10000","reduceOnly":false,"origQty":"10","executedQty":"0","cumQty":"","cumQuote":"0","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0","time":1592193425895,"updateTime":1592193514104,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","origType":"LIMIT","positionSide":"BOTH","closePosition":false},{"symbol":"BTCUSDT","orderId":2494253712,"clientOrderId":"sellOrder","price":"10000","reduceOnly":false,"origQty":"10","executedQty":"0","cumQty":"","cumQuote":"0","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0","time":1592193816956,"updateTime":1592193816956,"workingType":"CONTRACT_PRICE","activatePrice":"","priceRate":"","avgPrice":"0.00000","origType":"LIMIT","positionSide":"BOTH","closePosition":false}]}
```


## 账户余额
* curl --location --request GET 'http://47.57.93.231:8000/api/futures/balance' --header 'Content-Type: application/json' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf'
```
{"error_code":0,"error_message":"ok","data":[{"accountAlias":"mYSgXqSgAuoC","asset":"USDT","balance":"9928.07748480","withdrawAvailable":"8928.07748480"},{"accountAlias":"mYSgXqSgAuoC","asset":"BNB","balance":"0","withdrawAvailable":"0.00000000"}]}
```


## 调整开仓杠杆
* curl --location --request POST 'http://47.57.93.231:8000/api/futures/leverage' --header 'Content-Type: application/json' --header 'token: 56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf' --data-raw '{"symbol":"BTCUSDT", "leverage":100}'
```
{"error_code":0,"error_message":"ok","data":{"leverage":100,"maxNotionalValue":"250000","symbol":"BTCUSDT"}}
```


# 经纪商测试

## 创建子账户
* curl --location --request POST 'http://47.57.93.231:8000/api/broker/subAccount' --header 'Content-Type: application/json'
```
{"error_code":0,"error_message":"ok","data":{"subaccountId":"485396905497952257"}}
```


## 开启合约权限
* curl --location --request POST 'http://47.57.93.231:8000/api/broker/subAccount/futures' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "futures":"true"}'
```
{"error_code":0,"error_message":"ok","data":{"subaccountId":"485396905497952257","enableFutures":true,"updateTime":1593690239552}}
```


## 创建子账户ApiKey
* curl --location --request POST 'http://47.57.93.231:8000/api/broker/subAccountApi' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "canTrade":"true", "futuresTrade":"true"}'
```
{"error_code":0,"error_message":"ok","data":{"subaccountId":"485396905497952257","apikey":"RoruLyLqUI4rNaWT0hJ64Wba0q00nGTm7Y0NCLfKaKdZ8SzsFrrUXVIvcEEOk3MI","secretkey":"Kb0m2V60w4RYjGioQsBuqSNYyjN3mwB3DBdoCcjZl4OHDIgvXOLCE03NIqnaNzQp","canTrade":true,"futuresTrade":true}}
```


## 删除子账户ApiKey
* curl --location --request DELETE 'http://47.57.93.231:8000/api/broker/subAccountApi' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "subAccountApiKey":"RoruLyLqUI4rNaWT0hJ64Wba0q00nGTm7Y0NCLfKaKdZ8SzsFrrUXVIvcEEOk3MI"}'
```
{"error_code":0,"error_message":"ok"}
```


## 查询子账户ApiKey
* curl --location --request GET 'http://47.57.93.231:8000/api/broker/subAccountApi?subAccountId=485396905497952257' --header 'Content-Type: application/json'
```
{"error_code":0,"error_message":"ok","data":[{"subaccountId":"485396905497952257","apikey":"BAYb7zqjwDXcuU52RIT65DlB0SLzGSk02zxE3YPbJlDFCukYXQWZWDPeHfnfm9MW","canTrade":true,"futuresTrade":true},{"subaccountId":"485396905497952257","apikey":"iFgqdhh2n68lrwwTHldpMGUut2g4hdRi4Phffl7hybbW6KMn3mEp87nnP4S4XuIy","canTrade":true,"futuresTrade":true}]}
```


## 更改子账户ApiKey 交易权限，合约权限
* curl --location --request POST 'http://47.57.93.231:8000/api/broker/subAccountApi/permission' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "subAccountApiKey":"BAYb7zqjwDXcuU52RIT65DlB0SLzGSk02zxE3YPbJlDFCukYXQWZWDPeHfnfm9MW", "canTrade":"true", "futuresTrade":"true"}'
```
{"error_code":0,"error_message":"ok"}
```


## 查询子账户
* curl --location --request GET 'http://47.57.93.231:8000/api/broker/subAccount' --header 'Content-Type: application/json'
```
{"error_code":0,"error_message":"ok","data":[{"subaccountId":"485396905497952257","makerCommission":"0.0010","takerCommission":"0.0010","createTime":1593686040000}]}
```


## 更改子账户合约手续费
* curl --location --request POST 'http://47.57.93.231:8000/api/broker/subAccountApi/commission/futures' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "symbol":"BTCUSDT", "makerAdjustment":100, "takerAdjustment":100}'
```
{"error_code":0,"error_message":"ok","data":{"subAccountId":485396905497952257,"symbol":"BTCUSDT","makerAdjustment":10,"takerAdjustment":10,"makerCommission":210,"takerCommission":410}}
```


## 获取子账户合约手续费
* curl --location --request GET 'http://47.57.93.231:8000/api/broker/subAccountApi/commission/futures?subAccountId=485396905497952257&symbol=BTCUSDT' --header 'Content-Type: application/json'
```
{"error_code":0,"error_message":"ok","data":[{"subaccountId":485396905497952257,"symbol":"BTCUSDT","makerAdjustment":10,"takerAdjustment":10,"makerCommission":210,"takerCommission":410}]}
```


## 获取经纪商账户信息
* curl --location --request GET 'http://47.57.93.231:8000/api/broker/info' --header 'Content-Type: application/json'
```
{"error_code":0,"error_message":"ok","data":{"maxMakerCommission":"0.00200000","minMakerCommission":"0.00100000","maxTakerCommission":"0.00200000","minTakerCommission":"0.00100000","subAccountQty":7,"maxSubAccountQty":1000}}
```

