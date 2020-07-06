# 永续合约

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
