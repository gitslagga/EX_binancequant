package tasks

import (
	"errors"
)

type ErrorCode int

const (
	EC_NONE              ErrorCode = 1
	EC_PARAMS_ERR                  = 10000
	EC_NETWORK_ERR                 = 9999
	EC_JSON_MARSHAL_ERR            = 9998
	EC_TOKEN_INVALID               = 9997
	EC_RESPONSE_DATA_ERR           = 9996
	EC_REQUEST_DATA_ERR            = 9995

	EC_NOT_ACTIVE      = 8999
	EC_FORMAT_ERR      = 8998
	EC_ALREADY_ACTIVE  = 8997
	EC_NO_BALANCE      = 8996
	EC_INVALID_OPERATE = 8995
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
		r = "请求错误|Request error"
	case EC_PARAMS_ERR:
		r = "参数错误|Params error"
	case EC_JSON_MARSHAL_ERR:
		r = "json格式异常|Json format exception"
	case EC_TOKEN_INVALID:
		r = "请登录后操作|Please log in to operate"
	case EC_RESPONSE_DATA_ERR:
		r = "请重新登录|Please login again"
	case EC_REQUEST_DATA_ERR:
		r = "非法请求|Illegal request"

	case EC_NOT_ACTIVE:
		r = "暂未激活|Not activated"
	case EC_ALREADY_ACTIVE:
		r = "已经激活|Already activated"
	case EC_NO_BALANCE:
		r = "余额不足|Insufficient balance"
	case EC_FORMAT_ERR:
		r = "格式化错误|Format error"
	case EC_INVALID_OPERATE:
		r = "无效的操作|Invalid operate"
	default:
	}
	return
}

func ErrorCodeMessage(c ErrorCode) (r string) {
	return c.String()
}

type CommonResp struct {
	RespCode int         `form:"respCode" json:"respCode"`
	RespDesc string      `form:"respDesc" json:"respDesc"`
	RespData interface{} `form:"respData,omitempty" json:"respData,omitempty"`
}

/*********************************** future trading *************************************/
type UserInfo struct {
	ID            uint64 `json:"id"`
	Email         string `json:"email"`
	EmailStatus   int    `json:"emailStatus"`
	Preliminary   string `json:"preliminary"`
	Phone         string `json:"phone"`
	PhoneStatus   int    `json:"phoneStatus"`
	UserName      string `json:"userName"`
	AvatarUrl     string `json:"avatarUrl"`
	UserTrueName  string `json:"userTrueName"`
	UserStatus    int    `json:"userStatus"`
	OtcSellStatus int    `json:"otcSellStatus"`
	LoginIp       string `json:"loginIp"`
}

type FutureRequest struct {
	Data string `json:"d" binding:"required"`
	Key  string `json:"k" binding:"required"`
}

type FutureResponse struct {
	RespCode int    `json:"respCode" binding:"required"`
	RespDesc string `json:"respDesc" binding:"required"`
	RespData string `json:"respData,omitempty"`
}

type DepositsAddressRequest struct {
	Coin    string `form:"coin" json:"coin" binding:"required"`
	Network string `form:"network" json:"network"`
}

type ListDepositsRequest struct {
	Coin      string `form:"coin" json:"coin"`
	Status    int    `form:"status" json:"status"`
	StartTime int64  `form:"startTime" json:"startTime"`
	EndTime   int64  `form:"endTime" json:"endTime"`
	Offset    int    `form:"offset" json:"offset"`
	Limit     int    `form:"limit" json:"limit"`
}

type TransferRequest struct {
	Asset  string  `json:"asset" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Type   int     `json:"type" binding:"required"`
}

type ListFuturesTransferRequest struct {
	Asset     string `form:"asset" json:"asset" binding:"required"`
	StartTime int64  `form:"startTime" json:"startTime" binding:"required"`
	EndTime   int64  `form:"endTime" json:"endTime"`
	Current   int64  `form:"current" json:"current"`
	Size      int64  `form:"size" json:"size"`
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

type ListWithdrawsRequest struct {
	Coin      string `form:"coin" json:"coin"`
	Status    int    `form:"status" json:"status"`
	StartTime int64  `form:"startTime" json:"startTime"`
	EndTime   int64  `form:"endTime" json:"endTime"`
	Offset    int    `form:"offset" json:"offset"`
	Limit     int    `form:"limit" json:"limit"`
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

type GetOrderRequest struct {
	Symbol            string `form:"symbol" json:"symbol" binding:"required"`
	OrderId           int64  `form:"orderId" json:"orderId"`
	OrigClientOrderId string `form:"origClientOrderId" json:"origClientOrderId"`
}

type CancelOrderRequest struct {
	Symbol            string `json:"symbol" binding:"required"`
	OrderId           int64  `json:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId"`
}

type CancelAllOpenOrdersRequest struct {
	Symbol string `json:"symbol" binding:"required"`
}

type ListOpenOrdersRequest struct {
	Symbol string `form:"symbol" json:"symbol"`
}

type ListOrdersRequest struct {
	Symbol    string `form:"symbol" json:"symbol"`
	OrderId   int64  `form:"orderId" json:"orderId"`
	StartTime int64  `form:"startTime" json:"startTime"`
	EndTime   int64  `form:"endTime" json:"endTime"`
	Limit     int    `form:"limit" json:"limit"`
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

type GetPositionMarginHistoryRequest struct {
	Symbol    string `form:"symbol" json:"symbol" binding:"required"`
	Type      int    `form:"type" json:"type"`
	StartTime int64  `form:"startTime" json:"startTime"`
	EndTime   int64  `form:"endTime" json:"endTime"`
	Limit     int64  `form:"limit" json:"limit"`
}

type GetPositionRiskRequest struct {
	Symbol string `form:"symbol" json:"symbol"`
}

type GetTradeHistoryRequest struct {
	Symbol    string `form:"symbol" json:"symbol"`
	FromId    uint64 `form:"fromId" json:"fromId"`
	StartTime int64  `form:"startTime" json:"startTime"`
	EndTime   int64  `form:"endTime" json:"endTime"`
	Limit     int64  `form:"limit" json:"limit"`
}

type GetIncomeHistoryRequest struct {
	Symbol     string `form:"symbol" json:"symbol"`
	IncomeType string `form:"incomeType" json:"incomeType"`
	StartTime  int64  `form:"startTime" json:"startTime"`
	EndTime    int64  `form:"endTime" json:"endTime"`
	Limit      int64  `form:"limit" json:"limit"`
}

type GetLeverageBracketRequest struct {
	Symbol string `form:"symbol" json:"symbol"`
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
	FromId      string  `json:"fromId"`
	ToId        string  `json:"toId"`
	FuturesType int     `json:"futuresType" binding:"required"`
	Asset       string  `json:"asset" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
}

type GenerateRebateHistoryRequest struct {
	SubAccountId string `json:"subAccountId"`
	StartTime    uint64 `json:"startTime"`
	EndTime      uint64 `json:"endTime"`
}

/*********************************** internal no token *************************************/
type GetBalanceNoTokenRequest struct {
	UserId uint64 `form:"userId" json:"userId" binding:"required"`
}

type CreateTransferNoTokenRequest struct {
	UserId      uint64  `form:"userId" json:"userId" binding:"required"`
	Type        int     `form:"type" json:"type" binding:"gte=1,lte=2"`
	FuturesType int     `form:"futuresType" json:"futuresType" binding:"required"`
	Asset       string  `form:"asset" json:"asset" binding:"required"`
	Amount      float64 `form:"amount" json:"amount" binding:"required"`
}
