# 合约账户

## 提币
* curl -X POST http://47.57.93.231:8000/api/account/withdraw --header "token:56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf" --header 'Content-Type: application/json' --data-raw '{"coin":"USDT", "address":"1", "amount":100}'
```
{"error_code":0,"error_message":"ok"}
```