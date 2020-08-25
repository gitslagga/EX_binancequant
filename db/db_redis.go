package db

import (
	"EX_binancequant/config"
	"EX_binancequant/mylog"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	redisPool *redis.Pool
)

func InitRedisCli() {
	address := config.Config.Redis.Address
	password := config.Config.Redis.Password
	maxActive := config.Config.Redis.MaxActive
	maxIdle := config.Config.Redis.MaxIdle
	idleMills := config.Config.Redis.IdleMills

	redisPool = newPool(address, password, maxIdle, maxActive, idleMills)
	_, err := redisPool.Dial()
	if err != nil {
		mylog.Logger.Fatal().Msgf("[InitRedis] dial redis failed, address=%v, password=%v", address, password)
	}

	fmt.Println("[Init] redis init succeed.")
}

func CloseRedisCli() {
	redisPool.Close()
}

func newPool(server, password string, maxidle int, maxactive int, idleMills int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxidle,
		IdleTimeout: time.Duration(idleMills) * time.Millisecond,
		MaxActive:   maxactive,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				mylog.Logger.Error().Msgf("[Dial] Dial redis pool failed, err=%v", err)
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					mylog.Logger.Error().Msgf("[Dial] Auth redis cluster failed, err=%v", err)
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				mylog.Logger.Error().Msgf("[TestOnBorrow] Ping to redis cluster failed, err=%v", err)
			}
			return err
		},
	}
}

func ConvertTokenToUserID(identity string) (string, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	redisKey := fmt.Sprintf("loginUserSession:%s", identity)
	result, err := redis.String(redisConn.Do("GET", redisKey))
	if err != nil {
		mylog.Logger.Error().Msgf("redis GET %v error, err:%v", redisKey, err)
		return "", err
	}

	return result, err
}

func ConvertUserTokenToUserInfo(userToken string) ([]byte, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	redisKey := fmt.Sprintf("loginUser:%s", userToken)
	result, err := redis.Bytes(redisConn.Do("GET", redisKey))
	if err != nil {
		mylog.Logger.Error().Msgf("redis GET %v error, err:%v", redisKey, err)
		return nil, err
	}

	return result, err
}
