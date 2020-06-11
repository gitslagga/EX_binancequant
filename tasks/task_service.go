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
	/****************************** 永续合约行情接口 *********************************/
	r.GET("/api/market/time", ServerTimeService)
	r.GET("/api/market/depth", DepthService)
	r.GET("/api/market/aggTrades", AggTradesService)
	r.GET("/api/market/klines", KlinesService)
	r.GET("/api/market/premiumIndex", PremiumIndexService)
	r.GET("/api/market/ticker/24hr", ListPriceChangeStatsService)
	r.GET("/api/market/ticker/price", ListPricesService)
	r.GET("/api/market/exchangeInfo", ExchangeInfoService)

	route := r.Use(beforeHandler())

	/****************************** 永续合约认证接口 *********************************/
	route.GET("/api/account/deposits/list", ListDepositsService)
	route.GET("/api/account/deposits/address", DepositsAddressService)
	route.GET("/api/account/spot", SpotAccountService)
	route.GET("/api/account/futures", FuturesAccountService)
	route.POST("/api/account/transfer", FuturesTransferService)
	route.GET("/api/account/transfer", ListFuturesTransferService)
	route.POST("/api/account/withdraw", CreateWithdrawService)
	route.GET("/api/account/withdraw", ListWithdrawsService)

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

	route.POST("/api/futures/listenKey ", StartUserStreamService)
	route.PUT("/api/futures/listenKey ", KeepaliveUserStreamService)
	route.DELETE("/api/futures/listenKey ", CloseUserStreamService)
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
			out.ErrorCode = data.EC_PARAMS_ERR
			out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
			c.JSON(http.StatusBadRequest, out)
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
