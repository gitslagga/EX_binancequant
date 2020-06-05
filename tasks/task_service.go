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
	route := r.Use(beforeHandler())

	/****************************** 永续合约 *********************************/
	route.GET("/api/account/deposits/list", ListDepositsService)
	route.GET("/api/account/deposits/address", DepositsAddressService)
	route.GET("/api/account/spot", SpotAccountService)
	route.GET("/api/account/futures", FuturesAccountService)
	route.POST("/api/account/transfer", FuturesTransferService)
	route.GET("/api/account/transfer", ListFuturesTransferService)

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
	route.GET("/api/futures/income", GetIncomeHistoryService)
	route.GET("/api/futures/leverageBracket", GetLeverageBracketService)

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
