package db

import (
	"EX_binancequant/config"
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
	UserID       uint64             `bson:"user_id"`
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

	client, err = mongo.Connect(context.Background(), opt)
	if err != nil {
		mylog.Logger.Fatal().Msgf("[InitMongoCli] mongo connection failed, err=%v, client=%v", err, client)
	}

	fmt.Println("[InitMongo] mongo succeed.")
}

func CloseMongoCli() {
	client.Disconnect(context.Background())
}

func getUserKeys(userID uint64) (*UserKeys, error) {
	timeout := time.Duration(config.Config.Mongo.Timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

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
func GetSpotClientByUserID(userID uint64) (*trade.Client, error) {
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
func GetFuturesClientByUserID(userID uint64) (*futures.Client, error) {
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
func GetActiveFuturesByUserID(userID uint64) bool {
	_, err := getUserKeys(userID)
	if err != nil {
		return false
	}

	return true
}

/**
获取子账户ID
*/
func GetSubAccountIdByUserID(userID uint64) (string, error) {
	userKeys, err := getUserKeys(userID)
	if err != nil {
		return "", err
	}

	return userKeys.SubAccountId, nil
}

/**
创建合约子账户
*/
func CreateFuturesSubAccount(userID uint64, subAccountId, apiKey, secretKey string) error {
	timeout := time.Duration(config.Config.Mongo.Timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

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
