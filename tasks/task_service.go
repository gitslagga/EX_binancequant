package tasks

import (
	"EX_binancequant/db"
	"EX_binancequant/trade"
	"EX_binancequant/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter(router *gin.Engine) {
	router.Use(cors.Default())

	//no authorized
	noAuthorized := router.Group("/api", requestHandler(), responseHandler())

	/****************************** 通用 - 永续合约行情接口 *********************************/
	noAuthorized.GET("/market/time", ServerTimeService)
	noAuthorized.GET("/market/depth", DepthService)
	noAuthorized.GET("/market/aggTrades", AggTradesService)
	noAuthorized.GET("/market/klines", KlinesService)
	noAuthorized.GET("/market/premiumIndex", PremiumIndexService)
	noAuthorized.GET("/market/ticker/24hr", ListPriceChangeStatsService)
	noAuthorized.GET("/market/ticker/price", ListPricesService)
	noAuthorized.GET("/market/exchangeInfo", ExchangeInfoService)

	/****************************** 后台 - 经纪人接口 *********************************/
	//TODO Backend Authorized
	/*//管理子账户权限
	noAuthorized.POST("/broker/subAccount", CreateSubAccountService)
	noAuthorized.POST("/broker/subAccount/futures", EnableSubAccountFuturesService)
	noAuthorized.POST("/broker/subAccountApi", CreateSubAccountApiService)
	noAuthorized.DELETE("/broker/subAccountApi", DeleteSubAccountApiService)
	noAuthorized.GET("/broker/subAccountApi", GetSubAccountApiService)
	noAuthorized.POST("/broker/subAccountApi/permission", ChangeSubAccountApiPermissionService)
	noAuthorized.GET("/broker/subAccount", GetSubAccountService)

	//调整子账户手续费
	noAuthorized.POST("/broker/subAccountApi/commission/futures", ChangeCommissionFuturesService)
	noAuthorized.GET("/broker/subAccountApi/commission/futures", GetCommissionFuturesService)
	noAuthorized.GET("/broker/info", GetInfoService)

	//经纪商账户与子账户划转
	noAuthorized.POST("/broker/transfer", CreateTransferService)
	noAuthorized.GET("/broker/transfer", GetTransferService)

	//子账户充币与资产
	noAuthorized.GET("/broker/subAccount/depositHist", GetSubAccountDepositHistService)
	noAuthorized.GET("/broker/subAccount/spotSummary", GetSubAccountSpotSummaryService)
	noAuthorized.GET("/broker/subAccount/futuresSummary", GetSubAccountFuturesSummaryService)

	//查询返佣记录
	noAuthorized.GET("/broker/rebate/recentRecord", GetRebateRecentRecordService)
	noAuthorized.POST("/broker/rebate/historicalRecord", GenerateRebateHistoryService)
	noAuthorized.GET("/broker/rebate/historicalRecord", GetRebateHistoryService)*/

	/****************************** 前台 - 永续合约接口 *********************************/
	//开启子账户认证
	authorized := router.Group("/api", tokenHandler(), requestHandler(), responseHandler())
	authorized.GET("/account/activeFutures", GetActiveFuturesService)
	authorized.POST("/account/activeFutures", CreateActiveFuturesService)

	//子账户资产，充币，提币，划转
	authorized.GET("/account/deposits/list", ListDepositsService)
	authorized.GET("/account/deposits/address", DepositsAddressService)
	authorized.GET("/account/spot", SpotAccountService)
	authorized.GET("/account/futures", FuturesAccountService)
	authorized.POST("/account/transfer", FuturesTransferService)
	authorized.GET("/account/transfer", ListFuturesTransferService)
	authorized.POST("/account/withdraw", CreateWithdrawService)
	authorized.GET("/account/withdraw", ListWithdrawsService)

	//永续合约交易
	authorized.POST("/futures/position/mode", ChangePositionModeService)
	authorized.GET("/futures/position/mode", GetPositionModeService)
	authorized.POST("/futures/order", CreateOrderService)
	authorized.GET("/futures/order", GetOrderService)
	authorized.DELETE("/futures/order", CancelOrderService)
	authorized.DELETE("/futures/allOpenOrders", CancelAllOpenOrdersService)
	authorized.GET("/futures/openOrders", ListOpenOrdersService)
	authorized.GET("/futures/allOrders", ListOrdersService)
	authorized.GET("/futures/balance", GetBalanceService)
	authorized.POST("/futures/leverage", ChangeLeverageService)
	authorized.POST("/futures/marginType", ChangeMarginTypeService)
	authorized.POST("/futures/positionMargin", UpdatePositionMarginService)
	authorized.GET("/futures/positionMargin", GetPositionMarginHistoryService)
	authorized.GET("/futures/positionRisk", GetPositionRiskService)
	authorized.GET("/futures/userTrades", GetTradeHistoryService)
	authorized.GET("/futures/income", GetIncomeHistoryService)
	authorized.GET("/futures/leverageBracket", GetLeverageBracketService)

	//WS许可证权限（每小时更新）
	authorized.POST("/futures/listenKey", StartUserStreamService)
	authorized.PUT("/futures/listenKey", KeepaliveUserStreamService)
	authorized.DELETE("/futures/listenKey", CloseUserStreamService)
}

func InitFutures() {
	fmt.Println("[Tasks] futures init ...")

	startPingService()

	fmt.Println("[Tasks] futures init success.")
}

func startPingService() {
	trade.NewPingService()

	go func() {
		timer := time.NewTicker(24 * time.Hour)
		for {
			select {
			case <-timer.C:
				trade.NewPingService()
			}
		}
	}()

	fmt.Println("[Tasks] StartPingService succeed.")
}

func tokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("token")
		token, valid := utils.GetVerifyToken(jwtToken)
		if token == "" || valid == false {
			out := CommonResp{}
			out.RespCode = EC_TOKEN_INVALID
			out.RespDesc = ErrorCodeMessage(EC_TOKEN_INVALID)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		var userInfo UserInfo
		byteUserInfo, err := db.ConvertUserTokenToUserInfo(token)
		if err != nil {
			out := CommonResp{}
			out.RespCode = EC_TOKEN_INVALID
			out.RespDesc = ErrorCodeMessage(EC_TOKEN_INVALID)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		err = json.Unmarshal(byteUserInfo, &userInfo)
		if err != nil || userInfo.ID == 0 {
			out := CommonResp{}
			out.RespCode = EC_TOKEN_INVALID
			out.RespDesc = ErrorCodeMessage(EC_TOKEN_INVALID)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		c.Set("user_id", userInfo.ID)
		c.Next()
	}
}

func requestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		out := CommonResp{}
		var futureRequest FutureRequest

		if err := c.ShouldBindJSON(&futureRequest); err != nil {
			out.RespCode = EC_PARAMS_ERR
			out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		encryptKey, err := base64.StdEncoding.DecodeString(futureRequest.Key)
		if err != nil {
			out.RespCode = EC_REQUEST_DATA_ERR
			out.RespDesc = ErrorCodeMessage(EC_REQUEST_DATA_ERR)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		encryptData, err := base64.StdEncoding.DecodeString(futureRequest.Data)
		if err != nil {
			out.RespCode = EC_REQUEST_DATA_ERR
			out.RespDesc = ErrorCodeMessage(EC_REQUEST_DATA_ERR)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		privateKey := []byte(`-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAJzIo0OkvKH5a2NyPYuAOorElm8OX3QVrArMsX2X/8B9pBEtcp5QSe6CBx5p4TcZf73scHDgubcp0DBM6gdWEcyUFkl+Z+XKwdSgkVyHQK01klJcjJvMErjIKGeUggsFRMNduFyJePWA6pT3MJ+OKhmIKH6YhF1uiYXvnRbNvWrXAgMBAAECgYEAm5DpjsDy+rUFvVMphjXh4LdXnTJhvEmUv9KDet9LQbBpDzJNPJDmCuayMZdVhNqkSctFdntFS10N2h83R7g7R3gTOE2xLNj51XMMtPmL86HGdVsw9JyBkI4vBwJDkQ65Y7p85mtOPQ+8FP742acn1W8XdS+77x7zp6O8gWmR5aECQQDYfjnQcLe/OUiC78o6KKe/cos7zdqY/Vq/oNEQOl4fL8Fu2RojAQnpP6swGOP/5JmB/hxua5H0XOqTqC4RcvLLAkEAuWUB/eRODOSQSQWEm6sD+XJHtkEgGfFNZ8kHLo7/Guto72OoOE+rWjfUrixwBF5cxUi+IWZqpycVz44nryEKpQJAERXFEkIS/jBTHKI332cd9engOxP/0FsOMllKpnE0xFlMdqcDfQez9IhlxiHwvF0aEDwxmjU7C4HZsVVwbUgZCQJAZJ6WiyaK2eJ/ELKm+xnBCXRlyVvlQU8+lJJ9jF5dxE154Vs0JIPQ2yEsE+/YR/ay4PwO/O+p+Nh0tPZRQXJsZQJAMkYsybGSLwl76GOFQfT5xyqNHaQj3rAiBD7W5XWkLApRxbLuBzWpHk9IaV3GNEQuLIEjGcFO6tCt56w97D+QYA==
-----END PRIVATE KEY-----`)
		key, err := utils.RsaDecrypt(privateKey, encryptKey)
		if err != nil {
			out.RespCode = EC_REQUEST_DATA_ERR
			out.RespDesc = ErrorCodeMessage(EC_REQUEST_DATA_ERR)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		requestData, err := utils.AesDecryptECB(encryptData, key)
		if err != nil {
			out.RespCode = EC_REQUEST_DATA_ERR
			out.RespDesc = ErrorCodeMessage(EC_REQUEST_DATA_ERR)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		fmt.Println(string(requestData))
		c.Set("requestData", requestData)
	}
}

func responseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		out := CommonResp{}
		responseData := c.MustGet("responseData").(CommonResp)
		if responseData.RespCode != 1 {
			out.RespCode = responseData.RespCode
			out.RespDesc = responseData.RespDesc
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		err, privateKey, publicKey := utils.GenRsaKey(1024)
		if err != nil {
			out.RespCode = EC_RESPONSE_DATA_ERR
			out.RespDesc = ErrorCodeMessage(EC_RESPONSE_DATA_ERR)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		randomString := utils.GetRandomString(16)
		rsaKey, err := utils.RsaEncrypt(publicKey, randomString)
		if err != nil {
			out.RespCode = EC_RESPONSE_DATA_ERR
			out.RespDesc = ErrorCodeMessage(EC_RESPONSE_DATA_ERR)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}
		stringKey := base64.StdEncoding.EncodeToString(rsaKey)

		if responseData.RespData == nil {
			responseData.RespData = []byte{}
		}
		aesData, err := utils.AesEncryptECB(responseData.RespData.([]byte), randomString)
		if err != nil {
			out.RespCode = EC_RESPONSE_DATA_ERR
			out.RespDesc = ErrorCodeMessage(EC_RESPONSE_DATA_ERR)
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}
		stringData := base64.StdEncoding.EncodeToString(aesData)

		out.RespData = privateKey + "," + stringKey + "," + stringData
		out.RespCode = EC_NONE.Code()
		out.RespDesc = EC_NONE.String()

		c.JSON(http.StatusOK, out)
	}
}
