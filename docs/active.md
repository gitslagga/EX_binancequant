# 子账户激活

## 获取子账户激活状态 
* curl http://47.57.93.231:8000/api/account/activeFutures --header "token:56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf"
```
{"error_code":0,"error_message":"ok","data":false}
```


## 子账户创建与合约激活
* curl -X POST http://47.57.93.231:8000/api/account/activeFutures --header "token:56f7ac9b-b70d-4d4c-a8bc-7e60a9b2bcdf"
```
{"error_code":0,"error_message":"ok"}
```
