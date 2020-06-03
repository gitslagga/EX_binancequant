package tasks

import (
	"EX_binancequant/db"
	"fmt"
	"time"
)

func InitFutures() {
	fmt.Println("[Tasks] swap init ...")

	StartPingService()

	fmt.Println("[Tasks] swap init success.")
}

func StartPingService() {
	db.NewPingService()

	go func() {
		timer := time.NewTicker(24 * time.Hour)
		for {
			select {
			case <-timer.C:
				db.NewPingService()
			}
		}
	}()

	fmt.Println("[Tasks] StartPingService succeed.")
}
