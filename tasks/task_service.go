package tasks

import (
	"EX_binancequant/trade"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter(r *gin.Engine) {
	/****************************** 永续合约 *********************************/
	r.POST("/api/account/deposits/list", ListDepositsService)

}

func InitFutures() {
	fmt.Println("[Tasks] futures init ...")

	StartPingService()

	fmt.Println("[Tasks] futures init success.")
}

func StartPingService() {
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
