# 合约账户

## 提币
* curl -X POST http://127.0.0.1:8000/binance/account/withdraw --header 'Content-Type: application/json' --data-raw '{"coin":"USDT", "address":"1", "amount":100}' --header "token:eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzM4NCJ9.eyJpc3MiOiJpc3N1ZXIiLCJleHAiOjE1OTg2MTk1NTMsImlhdCI6MTU5ODYwNTE1MywidG9rZW4iOiI1OTk2Mjg2ODEyMUQ0RjE0OTVBODk0OUUxOTk2MkJCRSJ9.zHxWhV3-xnsv0YA6zC_x05ZFqLRaYEbKvWu3teHhIb1hC_1lQ9cG58N4p-1YJd-p"
```
{"respCode":1,"respDesc":"SUCCESS"}
```