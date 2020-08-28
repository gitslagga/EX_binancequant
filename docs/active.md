# 子账户激活

## 获取子账户激活状态 
* curl --location --request GET 'http://127.0.0.1:8000/api/account/activeFutures' --header "token:eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzM4NCJ9.eyJpc3MiOiJpc3N1ZXIiLCJleHAiOjE1OTg2MTk1NTMsImlhdCI6MTU5ODYwNTE1MywidG9rZW4iOiI1OTk2Mjg2ODEyMUQ0RjE0OTVBODk0OUUxOTk2MkJCRSJ9.zHxWhV3-xnsv0YA6zC_x05ZFqLRaYEbKvWu3teHhIb1hC_1lQ9cG58N4p-1YJd-p"
```
{"respCode":1,"respDesc":"SUCCESS","respData":true}
```


## 子账户创建与合约激活
* curl --location --request POST http://127.0.0.1:8000/api/account/activeFutures --header "token:eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzM4NCJ9.eyJpc3MiOiJpc3N1ZXIiLCJleHAiOjE1OTg2MTk1NTMsImlhdCI6MTU5ODYwNTE1MywidG9rZW4iOiI1OTk2Mjg2ODEyMUQ0RjE0OTVBODk0OUUxOTk2MkJCRSJ9.zHxWhV3-xnsv0YA6zC_x05ZFqLRaYEbKvWu3teHhIb1hC_1lQ9cG58N4p-1YJd-p"
```
{"respCode":1,"respDesc":"SUCCESS"}
```
