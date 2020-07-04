package db

import (
	"EX_binancequant/config"
	"EX_binancequant/data"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"EX_binancequant/trade/futures"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	client *mongo.Client
)

type UserKeys struct {
	//struct里面获取ObjectID
	UserKeysID   primitive.ObjectID `bson:"_id"`
	UserID       string             `bson:"user_id"`
	SubAccountId string             `bson:"sub_account_id"`
	ApiKey       string             `bson:"api_key"`
	SecretKey    string             `bson:"secret_key"`
	CreateTime   int64              `bson:"create_time"`
}

func InitMongoCli() {
	var err error
	uri := config.Config.Mongo.ApplyURI
	localThreshold := config.Config.Mongo.LocalThreshold
	maxConnIdleTime := config.Config.Mongo.MaxConnIdleTime
	maxPoolSize := config.Config.Mongo.MaxPoolSize

	opt := options.Client().ApplyURI(uri)
	opt.SetLocalThreshold(time.Duration(localThreshold) * time.Second)   //只使用与mongo操作耗时小于3秒的
	opt.SetMaxConnIdleTime(time.Duration(maxConnIdleTime) * time.Second) //指定连接可以保持空闲的最大毫秒数
	opt.SetMaxPoolSize(maxPoolSize)                                      //使用最大的连接数

	client, err = mongo.Connect(getContext(), opt)
	if err != nil {
		mylog.Logger.Fatal().Msgf("[InitMongoCli] mongo connection failed, err=%v, client=%v", err, client)
	}

	fmt.Println("[InitMongo] mongo succeed.")
}

func CloseMongoCli() {
	client.Disconnect(getContext())
}

func getContext() context.Context {
	timeout := config.Config.Mongo.Timeout
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	return ctx
}

func getUserKeys(userID string) (*UserKeys, error) {
	ctx := data.NewContext()
	var userKeys UserKeys

	collection := client.Database("main_quantify").Collection("user_keys")
	err := collection.FindOne(ctx, bson.D{{"user_id", userID}}).Decode(&userKeys)
	if err != nil {
		mylog.Logger.Error().Msgf("[GetSpotClientByUserID] collection FindOne failed, err=%v", err)
		return nil, err
	}

	return &userKeys, nil
}

/**
根据用户ID获取客户端
*/
func GetSpotClientByUserID(userID string) (*trade.Client, error) {
	userKeys, err := getUserKeys(userID)
	if err != nil {
		return nil, err
	}

	client := trade.NewClientByParam(userKeys.ApiKey, userKeys.SecretKey)
	return client, nil
}

/**
根据用户ID获取合约客户端
*/
func GetFuturesClientByUserID(userID string) (*futures.Client, error) {
	userKeys, err := getUserKeys(userID)
	if err != nil {
		return nil, err
	}

	client := trade.NewFuturesClientByParam(userKeys.ApiKey, userKeys.SecretKey)
	return client, nil
}

/**
获取用户开户状态
*/
func GetActiveFuturesByUserID(userID string) bool {
	userKeys, err := getUserKeys(userID)
	if userKeys == nil || err != nil {
		return false
	}

	return true
}

/**
创建合约子账户
*/
func CreateFuturesSubAccount(userID, subAccountId, apiKey, secretKey string) error {
	ctx := data.NewContext()

	collection := client.Database("main_quantify").Collection("user_keys")
	userKeys := UserKeys{
		UserKeysID:   primitive.NewObjectID(),
		UserID:       userID,
		SubAccountId: subAccountId,
		ApiKey:       apiKey,
		SecretKey:    secretKey,
		CreateTime:   time.Now().Unix(),
	}

	_, err := collection.InsertOne(ctx, userKeys)
	if err != nil {
		mylog.Logger.Error().Msgf("[CreateFuturesSubAccount] collection InsertOne failed, err=%v", err)
		return err
	}

	return nil
}
