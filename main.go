package main

import (
	"EX_binancequant/config"
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/proxy"
	"EX_binancequant/tasks"
	"EX_binancequant/trade"
	"EX_binancequant/websocket"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var configAddr = flag.String("config", "./config/config.toml", "base configuration files for server")

func main() {
	flag.Parse()

	if *configAddr == "" {
		panic("Configuration file path is not set, server exit")
	}
	config.LoadConfig(*configAddr)

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

	go func() {
		http.HandleFunc("/ws", websocket.WSHandler)

		// service connections
		err := http.ListenAndServe(config.Config.Server.WSAddress, nil)
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("websocket listen err:%v\n", err)
		}
	}()

	go func() {
		srv := &http.Server{
			Addr:    config.Config.Server.Address,
			Handler: r,
		}

		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("server listen err:%v\n", err)
		}
	}()
	fmt.Println("the server start succeed!!!")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)
	sg := <-quit
	fmt.Printf("receive the signal:%v\n", sg)
}
