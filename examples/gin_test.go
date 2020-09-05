package examples

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义ip白名单
		whiteList := []string{
			"127.0.0.1",
		}

		ip := c.ClientIP()

		fmt.Println("ip: ", ip)

		flag := false

		for _, host := range whiteList {
			if ip == host {
				flag = true
				break
			}
		}

		if !flag {
			c.String(http.StatusNetworkAuthenticationRequired, "your ip is not trusted: %s", ip)
			c.Abort()
			return
		}

		c.Next()
	}
}

func TestGin(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(Auth())

	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	router.Run(":3000")
}
