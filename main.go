package main

import (
	"EX_binancequant/trade"
	"EX_binancequant/config"
	"EX_binancequant/data"
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/proxy"
	"EX_binancequant/tasks"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configAddr = flag.String("config", "./config/config.toml", "base configuration files for server")

func main() {
	var err error

	flag.Parse()

	if *configAddr == "" {
		panic("Configuration file path is not set, server exit")
	}
	config.LoadConfig(*configAddr)

	data.Location, err = time.LoadLocation(config.Config.Server.Location)
	if err != nil {
		panic(fmt.Sprintf("load location failed, err=%v", err))
	}
	mylog.ConfigLoggers()

	proxy.InitProxy()
	defer proxy.CloseProxy()
	db.InitRedisCli()
	defer db.CloseRedisCli()
	//db.InitMysqlCli()
	//defer db.CloseMysqlCli()

	trade.InitTrade()
	db.InitMongoCli()
	defer db.CloseMongoCli()
	tasks.InitFutures()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	tasks.InitRouter(r)

	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen err:%v\n", err)
		}
	}()
	fmt.Println("the server start succeed!!!")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)
	sg := <-quit
	fmt.Printf("receive the signal:%v\n", sg)

	close(data.ShutdownChan)

	data.Wg.Wait()
	fmt.Println("wg return...")

	fmt.Println("main shutdown")
}
