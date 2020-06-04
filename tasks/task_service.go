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
	route.POST("/api/account/deposits/list", ListDepositsService)
	route.POST("/api/account/deposits/address", DepositsAddressService)
	route.POST("/api/account/spot", SpotAccountService)
	route.POST("/api/account/futures", FuturesAccountService)
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
