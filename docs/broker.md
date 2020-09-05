# 经纪商测试

## 创建子账户
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/subAccount/create' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"subaccountId":"485396905497952257"}}
```


## 开启合约权限
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/subAccount/enable' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "futures":true}'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"subaccountId":"485396905497952257","enableFutures":true,"updateTime":1593690239552}}
```


## 查询子账户
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccount' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[{"subaccountId":"485396905497952257","makerCommission":"0.0010","takerCommission":"0.0010","createTime":1593686040000}]}
```
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccount?subAccountId=485396905497952257' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[{"subaccountId":"485396905497952257","makerCommission":"0.0010","takerCommission":"0.0010","createTime":1593686040000}]}
```


## 创建子账户ApiKey
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/subAccountApi/create' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "canTrade":true, "futuresTrade":true}'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"subaccountId":"485396905497952257","apikey":"RoruLyLqUI4rNaWT0hJ64Wba0q00nGTm7Y0NCLfKaKdZ8SzsFrrUXVIvcEEOk3MI","secretkey":"Kb0m2V60w4RYjGioQsBuqSNYyjN3mwB3DBdoCcjZl4OHDIgvXOLCE03NIqnaNzQp","canTrade":true,"futuresTrade":true}}
```


## 删除子账户ApiKey
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/subAccountApi/close' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "subAccountApiKey":"RoruLyLqUI4rNaWT0hJ64Wba0q00nGTm7Y0NCLfKaKdZ8SzsFrrUXVIvcEEOk3MI"}'
```
{"respCode":1,"respDesc":"SUCCESS"}
```


## 查询子账户ApiKey
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccountApi?subAccountId=485396905497952257' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[{"subaccountId":"485396905497952257","apikey":"BAYb7zqjwDXcuU52RIT65DlB0SLzGSk02zxE3YPbJlDFCukYXQWZWDPeHfnfm9MW","canTrade":true,"futuresTrade":true},{"subaccountId":"485396905497952257","apikey":"iFgqdhh2n68lrwwTHldpMGUut2g4hdRi4Phffl7hybbW6KMn3mEp87nnP4S4XuIy","canTrade":true,"futuresTrade":true}]}
```


## 更改子账户ApiKey 交易权限，合约权限
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/subAccountApi/permission' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "subAccountApiKey":"BAYb7zqjwDXcuU52RIT65DlB0SLzGSk02zxE3YPbJlDFCukYXQWZWDPeHfnfm9MW", "canTrade":true, "futuresTrade":true}'
```
{"respCode":1,"respDesc":"SUCCESS"}
```


## 更改子账户合约手续费
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/subAccountApi/commission/futures' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "symbol":"BTCUSDT", "makerAdjustment":100, "takerAdjustment":100}'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"subAccountId":485396905497952257,"symbol":"BTCUSDT","makerAdjustment":10,"takerAdjustment":10,"makerCommission":210,"takerCommission":410}}
```


## 获取子账户合约手续费
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccountApi/commission/futures?subAccountId=485396905497952257&symbol=BTCUSDT' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[{"subaccountId":485396905497952257,"symbol":"BTCUSDT","makerAdjustment":10,"takerAdjustment":10,"makerCommission":210,"takerCommission":410}]}
```


## 获取经纪商账户信息
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/info' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"maxMakerCommission":"0.00200000","minMakerCommission":"0.00100000","maxTakerCommission":"0.00200000","minTakerCommission":"0.00100000","subAccountQty":7,"maxSubAccountQty":1000}}
```


## 经纪商和子账户划转
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/transfer' --header 'Content-Type: application/json' --data-raw '{"toId":"485396905497952257", "futuresType":1, "asset":"USDT", "amount":100}'
```
{"respCode":30110101,"respDesc":"\u003cAPIError\u003e code=-9000, msg=user have no avaliable amount"}
```
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/transfer' --header 'Content-Type: application/json' --data-raw '{"fromId":"485396905497952257", "futuresType":1, "asset":"USDT", "amount":100}'
```
{"respCode":30110101,"respDesc":"\u003cAPIError\u003e code=-9000, msg=user have no avaliable amount"}
```


## 获取经纪商和子账户划转记录
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/transfer?subAccountId=485396905497952257' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[]}
```
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/transfer?subAccountId=485396905497952257&startTime=0&endTime=1' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[]}
```


## 获取子账户充币记录
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccount/depositHist' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[]}
```
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccount/depositHist?subAccountId=485396905497952257&startTime=0&endTime=1' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[]}
```


## 获取子账户现货资产
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccount/spotSummary' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"respData":[{"subAccountId":"485396905497952257","totalBalanceOfBtc":"0.00000000"}],"sourceAddress":""}}
```
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccount/spotSummary?subAccountId=485396905497952257' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[]}
```


## 获取子账户合约资产
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccount/futuresSummary' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"respData":[{"futuresEnable":true,"subAccountId":"485396905497952257","totalInitialMarginOfUsdt":"0","totalMaintenanceMarginOfUsdt":"0","totalWalletBalanceOfUsdt":"0","totalUnrealizedProfitOfUsdt":"0","totalMarginBalanceOfUsdt":"0","totalPositionInitialMarginOfUsdt":"0","totalOpenOrderInitialMarginOfUsdt":"0"}],"timestamp":1593745369080}}
```
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/subAccount/futuresSummary?subAccountId=485396905497952257' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"respData":[{"futuresEnable":true,"subAccountId":"485396905497952257","totalInitialMarginOfUsdt":"0","totalMaintenanceMarginOfUsdt":"0","totalWalletBalanceOfUsdt":"0","totalUnrealizedProfitOfUsdt":"0","totalMarginBalanceOfUsdt":"0","totalPositionInitialMarginOfUsdt":"0","totalOpenOrderInitialMarginOfUsdt":"0"}],"timestamp":1593745369080}}
```


## 获取子账户七天返佣记录
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/rebate/recentRecord?subAccountId=485396905497952257&startTime=0&endTime=1&limit=10' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":[]}
```


## 开启子账户三十天返佣记录
* curl --location --request POST 'http://127.0.0.1:8000/binance/broker/rebate/historicalRecord' --header 'Content-Type: application/json' --data-raw '{"subAccountId":"485396905497952257", "startTime":1, "endTime":100}'
```
{"respCode":1,"respDesc":"SUCCESS","respData":{"code":200,"msg":"Historical data is collecting"}
```


## 获取子账户三十天返佣记录
* curl --location --request GET 'http://127.0.0.1:8000/binance/broker/rebate/historicalRecord?subAccountId=485396905497952257&startTime=0&endTime=1&limit=10' --header 'Content-Type: application/json'
```
{"respCode":1,"respDesc":"SUCCESS","respData":"https://bin-prod-user-rebate-bucket.s3.amazonaws.com/user-rebate/b4b6ca80-bcdc-11ea-8a61-0ad86c4d89f6/part-00000-d67a3f95-97ed-428f-89f0-44929fbd3405-c000.csv?AWSAccessKeyId=ASIAVL364M5ZHDRTC3ES\u0026Expires=1593833086\u0026x-amz-security-token=IQoJb3JpZ2luX2VjENP%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaDmFwLW5vcnRoZWFzdC0xIkcwRQIgf3MuxZjiULuxAGXOziJj5%2FjYSCXNyr9ZYRXyT6ZSbasCIQDz45ZUTIMN%2BYpnG%2B%2FVkGFsjfd2ODO0EGfitFmkQcGgyyq%2BAwhsEAAaDDM2OTA5NjQxOTE4NiIMNwJsOnb2HzG85xU4KpsDhgfDP6h5NMC8zNXNJBKHGv%2BtLSCjijh5hrADVnUtWJGylhLrb1YQfweynzwHe4tLm7LIHHvojaT1l62lKy2kUWNaBXTjW0KwQUltuO5EuvrCLHFsuPJbu9493NLI9Bdc9Tg%2BlGgdSIDXxHwt4SGtiPRXhELNvUlKK0HHAL6zDRMMuSsFHivn9NEdm3OoGW3m0XFitK4XRhDhjgxehm8xTJsznjj1UlXq7d%2BcUqrK2rO13%2BVNhOPYQdNE%2FAIy6CA8mVyioNVGDfMmVX9%2BeuGBWFeogIv%2BlkQC%2FiGlsfLT%2FejfYHVREX3NfH5C2MzB8VffKmeUfQbxjGh7GEFeQyraprx5iH6ukVJOCoWHQAnrMlLyitCkxuT7Bc19hRDzKXrw3NnewOuz2CWGD%2Fc8ALSV6xdAhaadq4mGekjHt%2Fyph7fT3Ctjx%2BSR1EXmtzvBe2X%2BR%2FFoTw5ismACvdQdWHoigs9ef66lojfLUWU3CNwITX7nvP%2Fx27ISTduQT2RgynM1XYWMJdC4ZcxLuC71cYxnM%2B7swrbSx8Jyw3jQMLLE%2BvcFOusBQ3q3UmnqDQZ5NzISfzFQRcT1%2Fb5YuWZBttswaf2bWwYy82P%2FeV%2BoFJBjXh3zkc6oTpA0w1FfE2LW3Pz3Rh1E8jOQyGb2IQGNvdByfQVdJztV%2F%2BMtDFd5w0ZkBZwFEPleikx%2Fn5P04VVPX6%2FjY%2BYtM1RDoa8%2FaDFeAxeFQnNnYNNvXJdKF8dax9cbT2tNy1AzmOQDLAywlUwvzbMJdAJMoS04%2FAn%2FjP%2BPpbwgIjvuieHzXOY8gAuqfpod85JlbLPwgpllulGj8wd1J8Ma4wx06UxcYALC%2Bp%2FIAlsW9mVp0jE5kRuwkMe62cCVwQ%3D%3D\u0026Signature=NXEdVtbPvjY5iv0I6ZTlQGp8baw%3D"}
```
