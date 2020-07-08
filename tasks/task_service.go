package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/db"
	"EX_binancequant/trade"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter(r *gin.Engine) {
	/****************************** 通用 - 永续合约行情接口 *********************************/
	r.GET("/api/market/time", ServerTimeService)
	r.GET("/api/market/depth", DepthService)
	r.GET("/api/market/aggTrades", AggTradesService)
	r.GET("/api/market/klines", KlinesService)
	r.GET("/api/market/premiumIndex", PremiumIndexService)
	r.GET("/api/market/ticker/24hr", ListPriceChangeStatsService)
	r.GET("/api/market/ticker/price", ListPricesService)
	r.GET("/api/market/exchangeInfo", ExchangeInfoService)

	/****************************** 后台 - 经纪人接口 *********************************/
	//管理子账户权限
	r.POST("/api/broker/subAccount", CreateSubAccountService)
	r.POST("/api/broker/subAccount/futures", EnableSubAccountFuturesService)
	r.POST("/api/broker/subAccountApi", CreateSubAccountApiService)
	r.DELETE("/api/broker/subAccountApi", DeleteSubAccountApiService)
	r.GET("/api/broker/subAccountApi", GetSubAccountApiService)
	r.POST("/api/broker/subAccountApi/permission", ChangeSubAccountApiPermissionService)
	r.GET("/api/broker/subAccount", GetSubAccountService)

	//调整子账户手续费
	r.POST("/api/broker/subAccountApi/commission/futures", ChangeCommissionFuturesService)
	r.GET("/api/broker/subAccountApi/commission/futures", GetCommissionFuturesService)
	r.GET("/api/broker/info", GetInfoService)

	//经纪商账户与子账户划转
	r.POST("/api/broker/transfer", CreateTransferService)
	r.GET("/api/broker/transfer", GetTransferService)

	//子账户充币与资产
	r.GET("/api/broker/subAccount/depositHist", GetSubAccountDepositHistService)
	r.GET("/api/broker/subAccount/spotSummary", GetSubAccountSpotSummaryService)
	r.GET("/api/broker/subAccount/futuresSummary", GetSubAccountFuturesSummaryService)

	//查询返佣记录
	r.GET("/api/broker/rebate/recentRecord", GetRebateRecentRecordService)
	r.POST("/api/broker/rebate/historicalRecord", GenerateRebateHistoryService)
	r.GET("/api/broker/rebate/historicalRecord", GetRebateHistoryService)

	/****************************** 前台 - 永续合约接口 *********************************/
	//开启子账户认证
	route := r.Use(beforeHandler())
	route.GET("/api/account/activeFutures", GetActiveFuturesService)
	route.POST("/api/account/activeFutures", CreateActiveFuturesService)

	//子账户资产，充币，提币，划转
	route.GET("/api/account/deposits/list", ListDepositsService)
	route.GET("/api/account/deposits/address", DepositsAddressService)
	route.GET("/api/account/spot", SpotAccountService)
	route.GET("/api/account/futures", FuturesAccountService)
	route.POST("/api/account/transfer", FuturesTransferService)
	route.GET("/api/account/transfer", ListFuturesTransferService)
	route.POST("/api/account/withdraw", CreateWithdrawService)
	route.GET("/api/account/withdraw", ListWithdrawsService)

	//永续合约交易
	route.POST("/api/futures/position/mode", ChangePositionModeService)
	route.GET("/api/futures/position/mode", GetPositionModeService)
	route.POST("/api/futures/order", CreateOrderService)
	route.GET("/api/futures/order", GetOrderService)
	route.DELETE("/api/futures/order", CancelOrderService)
	route.DELETE("/api/futures/allOpenOrders", CancelAllOpenOrdersService)
	route.GET("/api/futures/openOrders", ListOpenOrdersService)
	route.GET("/api/futures/allOrders", ListOrdersService)
	route.GET("/api/futures/balance", GetBalanceService)
	route.POST("/api/futures/leverage", ChangeLeverageService)
	route.POST("/api/futures/marginType", ChangeMarginTypeService)
	route.POST("/api/futures/positionMargin", UpdatePositionMarginService)
	route.GET("/api/futures/positionMargin", GetPositionMarginHistoryService)
	route.GET("/api/futures/positionRisk", GetPositionRiskService)
	route.GET("/api/futures/userTrades", GetTradeHistoryService)
	route.GET("/api/futures/income", GetIncomeHistoryService)
	route.GET("/api/futures/leverageBracket", GetLeverageBracketService)

	//WS许可证权限（每小时更新）
	route.POST("/api/futures/listenKey", StartUserStreamService)
	route.PUT("/api/futures/listenKey", KeepaliveUserStreamService)
	route.DELETE("/api/futures/listenKey", CloseUserStreamService)
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

func beforeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userID, err := db.ConvertTokenToUserID(token)

		if err != nil {
			out := data.CommonResp{}
			out.ErrorCode = data.EC_USER_NOT_EXIST
			out.ErrorMessage = data.ErrorCodeMessage(data.EC_USER_NOT_EXIST)
			c.JSON(http.StatusBadRequest, out)
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
