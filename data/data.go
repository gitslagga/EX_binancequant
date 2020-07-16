package data

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	Location     *time.Location
	Wg           sync.WaitGroup
	ShutdownChan = make(chan int)
)

type ErrorCode int

const (
	EC_NONE               ErrorCode = 1
	EC_PARAMS_ERR                   = 30110100
	EC_NETWORK_ERR                  = 30110101
	EC_INTERNAL_ERR                 = 30110102
	EC_INTERNAL_ERR_DB              = 30110103
	EC_INTERNAL_ERR_REDIS           = 30110104

	EC_USER_NOT_EXIST = 30200000 + 10
	EC_NOT_ACTIVE     = 30200000 + 11
	EC_FORMAT_ERR     = 30200000 + 12
	EC_ALREADY_ACTIVE = 30100000 + 13
	EC_NO_BALANCE     = 30200000 + 14
)

func (c ErrorCode) Code() (r int) {
	r = int(c)
	return
}

func (c ErrorCode) Error() (r error) {
	r = errors.New(c.String())
	return
}

func (c ErrorCode) String() (r string) {
	switch c {
	case EC_NONE:
		r = "SUCCESS"
	case EC_NETWORK_ERR:
		r = "Network error"
	case EC_PARAMS_ERR:
		r = "Parameter error"
	case EC_INTERNAL_ERR:
		r = "Server error"
	case EC_INTERNAL_ERR_DB:
		r = "Server error"
	case EC_INTERNAL_ERR_REDIS:
		r = "Server error"

	case EC_USER_NOT_EXIST:
		r = "User does not exist"
	case EC_NOT_ACTIVE:
		r = "Not activated"
	case EC_ALREADY_ACTIVE:
		r = "Already activated"
	case EC_NO_BALANCE:
		r = "Insufficient balance"
	case EC_FORMAT_ERR:
		r = "Format error"
	default:
	}
	return
}

func ErrorCodeMessage(c ErrorCode) (r string) {
	return c.String()
}

type CommonResp struct {
	RespCode int         `json:"respCode" form:"respCode"`
	RespDesc string      `json:"respDesc" form:"respDesc"`
	RespData interface{} `json:"respData,omitempty" form:"respData"`
}

func NewContext() context.Context {
	return context.Background()
}

/*********************************** future trading *************************************/
type TransferRequest struct {
	Asset  string  `json:"asset" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Type   int     `json:"type" binding:"required"`
}

type WithdrawRequest struct {
	Coin               string  `json:"coin" binding:"required"`
	Address            string  `json:"address" binding:"required"`
	Amount             float64 `json:"amount" binding:"required"`
	WithdrawOrderId    string  `json:"withdrawOrderId"`
	Network            string  `json:"network"`
	AddressTag         string  `json:"addressTag"`
	TransactionFeeFlag bool    `json:"transactionFeeFlag"`
	Name               string  `json:"name"`
}

type PositionModeRequest struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

type OrderRequest struct {
	Symbol           string  `json:"symbol" binding:"required"`
	Side             string  `json:"side" binding:"required"`
	Type             string  `json:"type" binding:"required"`
	PositionSide     string  `json:"positionSide"`
	ReduceOnly       bool    `json:"reduceOnly"`
	Quantity         float64 `json:"quantity"`
	Price            float64 `json:"price"`
	NewClientOrderId string  `json:"newClientOrderId"`
	StopPrice        float64 `json:"stopPrice"`
	ClosePosition    bool    `json:"closePosition"`
	ActivationPrice  float64 `json:"activationPrice"`
	CallbackRate     float64 `json:"callbackRate"`
	TimeInForce      string  `json:"timeInForce"`
	WorkingType      string  `json:"workingType"`
	NewOrderRespType string  `json:"newOrderRespType"`
}

type LeverageRequest struct {
	Symbol   string `json:"symbol" binding:"required"`
	Leverage int    `json:"leverage" binding:"required"`
}

type MarginTypeRequest struct {
	Symbol     string `json:"symbol" binding:"required"`
	MarginType string `json:"marginType" binding:"required"`
}

type PositionMarginRequest struct {
	Symbol       string  `json:"symbol" binding:"required"`
	Amount       float64 `json:"amount"  binding:"required"`
	Type         int     `json:"type" binding:"required"`
	PositionSide string  `json:"positionSide"`
}

/*********************************** broker sub account *************************************/
type EnableFuturesRequest struct {
	SubAccountId string `json:"subAccountId" binding:"required"`
	Futures      bool   `json:"futures"`
}

type CreateApiKeyRequest struct {
	SubAccountId string `json:"subAccountId" binding:"required"`
	CanTrade     bool   `json:"canTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

type DeleteApiKeyRequest struct {
	SubAccountId     string `json:"subAccountId" binding:"required"`
	SubAccountApiKey string `json:"subAccountApiKey" binding:"required"`
}

type ChangeApiPermissionRequest struct {
	SubAccountId     string `json:"subAccountId" binding:"required"`
	SubAccountApiKey string `json:"subAccountApiKey" binding:"required"`
	CanTrade         bool   `json:"canTrade"`
	FuturesTrade     bool   `json:"futuresTrade"`
}

type ChangeCommissionFuturesRequest struct {
	SubAccountId    string `json:"subAccountId" binding:"required"`
	Symbol          string `json:"symbol" binding:"required"`
	MakerAdjustment int    `json:"makerAdjustment" binding:"required"`
	TakerAdjustment int    `json:"takerAdjustment" binding:"required"`
}

type CreateTransferRequest struct {
	FromId       string  `json:"fromId"`
	ToId         string  `json:"toId"`
	ClientTranId string  `json:"clientTranId"`
	Asset        string  `json:"asset" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
}

type GenerateRebateHistoryRequest struct {
	SubAccountId string `json:"subAccountId"`
	StartTime    uint64 `json:"startTime"`
	EndTime      uint64 `json:"endTime"`
}
