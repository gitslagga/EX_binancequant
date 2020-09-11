package trade

import (
	"context"
	"encoding/json"
)

// GetInfoService query broker account information
type GetInfoService struct {
	c *Client
}

// Do send request
func (s *GetInfoService) Do(ctx context.Context, opts ...RequestOption) (res *GetInfo, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/info",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetInfo define query broker info
type GetInfo struct {
	MaxMakerCommission string `json:"maxMakerCommission"`
	MinMakerCommission string `json:"minMakerCommission"`
	MaxTakerCommission string `json:"maxTakerCommission"`
	MinTakerCommission string `json:"minTakerCommission"`
	SubAccountQty      int    `json:"subAccountQty"`
	MaxSubAccountQty   int    `json:"maxSubAccountQty"`
}

// CreateTransferService create sub account transfer
type CreateTransferService struct {
	c           *Client
	fromId      *string
	toId        *string
	futuresType *int
	asset       *string
	amount      *float64
}

// FromId set fromId
func (s *CreateTransferService) FromId(fromId string) *CreateTransferService {
	s.fromId = &fromId
	return s
}

// ToId set toId
func (s *CreateTransferService) ToId(toId string) *CreateTransferService {
	s.toId = &toId
	return s
}

// FuturesType set futuresType 1:USDT Futures,2: COIN Futures
func (s *CreateTransferService) FuturesType(futuresType int) *CreateTransferService {
	s.futuresType = &futuresType
	return s
}

// Asset set asset
func (s *CreateTransferService) Asset(asset string) *CreateTransferService {
	s.asset = &asset
	return s
}

// Amount set amount
func (s *CreateTransferService) Amount(amount float64) *CreateTransferService {
	s.amount = &amount
	return s
}

// Do send request
func (s *CreateTransferService) Do(ctx context.Context, opts ...RequestOption) (res *CreateTransfer, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/transfer/futures",
		secType:  secTypeSigned,
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}
	r.setParam("futuresType", *s.futuresType)
	r.setParam("asset", *s.asset)
	r.setParam("amount", *s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateTransfer)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateTransfer define create sub account transfer
type CreateTransfer struct {
	TxnId        uint64 `json:"txnId"`
	ClientTranId string `json:"clientTranId,omitempty"`
}

// GetTransferService create sub account transfer history
type GetTransferService struct {
	c            *Client
	subAccountId *string
	clientTranId *string
	startTime    *uint64
	endTime      *uint64
	page         *int
	limit        *int
}

// SubAccountId set subAccountId
func (s *GetTransferService) SubAccountId(subAccountId string) *GetTransferService {
	s.subAccountId = &subAccountId
	return s
}

// ClientTranId set clientTranId
func (s *GetTransferService) ClientTranId(clientTranId string) *GetTransferService {
	s.clientTranId = &clientTranId
	return s
}

// StartTime set startTime
func (s *GetTransferService) StartTime(startTime uint64) *GetTransferService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetTransferService) EndTime(endTime uint64) *GetTransferService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *GetTransferService) Page(page int) *GetTransferService {
	s.page = &page
	return s
}

// Limit set limit
func (s *GetTransferService) Limit(limit int) *GetTransferService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetTransferService) Do(ctx context.Context, opts ...RequestOption) (res []*GetTransfer, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*GetTransfer, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTransfer define get sub account transfer history
type GetTransfer struct {
	FromId       string `json:"fromId"`
	ToId         string `json:"toId"`
	Asset        string `json:"asset"`
	Qty          string `json:"qty"`
	Time         uint64 `json:"time"`
	TxnId        string `json:"txnId"`
	ClientTranId string `json:"clientTranId,omitempty"`
}

// GetSubAccountDepositHistService get sub account depository history
type GetSubAccountDepositHistService struct {
	c            *Client
	subAccountId *string
	coin         *string
	status       *int
	startTime    *uint64
	endTime      *uint64
	limit        *int
	offset       *int
}

// SubAccountId set subAccountId
func (s *GetSubAccountDepositHistService) SubAccountId(subAccountId string) *GetSubAccountDepositHistService {
	s.subAccountId = &subAccountId
	return s
}

// Coin set coin
func (s *GetSubAccountDepositHistService) Coin(coin string) *GetSubAccountDepositHistService {
	s.coin = &coin
	return s
}

// Status set status
func (s *GetSubAccountDepositHistService) Status(status int) *GetSubAccountDepositHistService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *GetSubAccountDepositHistService) StartTime(startTime uint64) *GetSubAccountDepositHistService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetSubAccountDepositHistService) EndTime(endTime uint64) *GetSubAccountDepositHistService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetSubAccountDepositHistService) Limit(limit int) *GetSubAccountDepositHistService {
	s.limit = &limit
	return s
}

// Offset set offset
func (s *GetSubAccountDepositHistService) Offset(offset int) *GetSubAccountDepositHistService {
	s.offset = &offset
	return s
}

// Do send request
func (s *GetSubAccountDepositHistService) Do(ctx context.Context, opts ...RequestOption) (res []*GetSubAccountDepositHist, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount/depositHist",
		secType:  secTypeSigned,
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*GetSubAccountDepositHist, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSubAccountDepositHist define get sub account deposit history
type GetSubAccountDepositHist struct {
	SubAccountId  string `json:"subaccountId"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	InsertTime    uint64 `json:"insertTime"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	TxId          string `json:"txId"`
	SourceAddress string `json:"sourceAddress"`
}

// GetSubAccountSpotSummaryService get sub account spot summary
type GetSubAccountSpotSummaryService struct {
	c            *Client
	subAccountId *string
	page         *uint64
	size         *uint64
}

// SubAccountId set subAccountId
func (s *GetSubAccountSpotSummaryService) SubAccountId(subAccountId string) *GetSubAccountSpotSummaryService {
	s.subAccountId = &subAccountId
	return s
}

// Page set page
func (s *GetSubAccountSpotSummaryService) Page(page uint64) *GetSubAccountSpotSummaryService {
	s.page = &page
	return s
}

// Size set size
func (s *GetSubAccountSpotSummaryService) Size(size uint64) *GetSubAccountSpotSummaryService {
	s.size = &size
	return s
}

// Do send request
func (s *GetSubAccountSpotSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *GetSubAccountSpotSummary, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount/spotSummary",
		secType:  secTypeSigned,
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetSubAccountSpotSummary)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSubAccountSpotSummary define get sub account spot summary
type GetSubAccountSpotSummary struct {
	Data      []*subAccountSpotSummary
	Timestamp uint64 `json:"timestamp"`
}

type subAccountSpotSummary struct {
	SubAccountId      string `json:"subAccountId"`
	TotalBalanceOfBtc string `json:"totalBalanceOfBtc"`
}

// GetSubAccountFuturesSummaryService get sub account futures summary
type GetSubAccountFuturesSummaryService struct {
	c            *Client
	subAccountId *string
	page         *uint64
	size         *uint64
}

// SubAccountId set subAccountId
func (s *GetSubAccountFuturesSummaryService) SubAccountId(subAccountId string) *GetSubAccountFuturesSummaryService {
	s.subAccountId = &subAccountId
	return s
}

// Page set page
func (s *GetSubAccountFuturesSummaryService) Page(page uint64) *GetSubAccountFuturesSummaryService {
	s.page = &page
	return s
}

// Size set size
func (s *GetSubAccountFuturesSummaryService) Size(size uint64) *GetSubAccountFuturesSummaryService {
	s.size = &size
	return s
}

// Do send request
func (s *GetSubAccountFuturesSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *GetSubAccountFuturesSummary, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount/futuresSummary",
		secType:  secTypeSigned,
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetSubAccountFuturesSummary)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSubAccountFuturesSummary define get sub account futures summary
type GetSubAccountFuturesSummary struct {
	Data      []*subAccountFuturesSummary
	Timestamp uint64 `json:"timestamp"`
}

type subAccountFuturesSummary struct {
	FuturesEnable                     bool   `json:"futuresEnable"`
	SubAccountId                      string `json:"subAccountId"`
	TotalInitialMarginOfUsdt          string `json:"totalInitialMarginOfUsdt"`
	TotalMaintenanceMarginOfUsdt      string `json:"totalMaintenanceMarginOfUsdt"`
	TotalWalletBalanceOfUsdt          string `json:"totalWalletBalanceOfUsdt"`
	TotalUnrealizedProfitOfUsdt       string `json:"totalUnrealizedProfitOfUsdt"`
	TotalMarginBalanceOfUsdt          string `json:"totalMarginBalanceOfUsdt"`
	TotalPositionInitialMarginOfUsdt  string `json:"totalPositionInitialMarginOfUsdt"`
	TotalOpenOrderInitialMarginOfUsdt string `json:"totalOpenOrderInitialMarginOfUsdt"`
}

// GetRebateRecentRecordService get rebate recent record history
type GetRebateRecentRecordService struct {
	c            *Client
	subAccountId *string
	startTime    *uint64
	endTime      *uint64
	limit        *int
}

// SubAccountId set subAccountId
func (s *GetRebateRecentRecordService) SubAccountId(subAccountId string) *GetRebateRecentRecordService {
	s.subAccountId = &subAccountId
	return s
}

// StartTime set startTime
func (s *GetRebateRecentRecordService) StartTime(startTime uint64) *GetRebateRecentRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetRebateRecentRecordService) EndTime(endTime uint64) *GetRebateRecentRecordService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetRebateRecentRecordService) Limit(limit int) *GetRebateRecentRecordService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetRebateRecentRecordService) Do(ctx context.Context, opts ...RequestOption) (res []*GetRebateRecentRecord, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/rebate/recentRecord",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	r.setParam("startTime", *s.startTime)
	r.setParam("endTime", *s.endTime)
	r.setParam("limit", *s.limit)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*GetRebateRecentRecord, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetRebateRecentRecord define get rebate recent record
type GetRebateRecentRecord struct {
	SubAccountId string `json:"subaccountId"`
	Income       string `json:"income"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	TradeId      uint64 `json:"tradeId"`
	Time         uint64 `json:"time"`
}

// GenerateRebateHistoryService generate rebate history
type GenerateRebateHistoryService struct {
	c            *Client
	subAccountId *string
	startTime    *uint64
	endTime      *uint64
}

// SubAccountId set subAccountId
func (s *GenerateRebateHistoryService) SubAccountId(subAccountId string) *GenerateRebateHistoryService {
	s.subAccountId = &subAccountId
	return s
}

// StartTime set startTime
func (s *GenerateRebateHistoryService) StartTime(startTime uint64) *GenerateRebateHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GenerateRebateHistoryService) EndTime(endTime uint64) *GenerateRebateHistoryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *GenerateRebateHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *GenerateRebateHistory, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/rebate/historicalRecord",
		secType:  secTypeSigned,
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GenerateRebateHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GenerateRebateHistory define generate rebate history
type GenerateRebateHistory struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// GetRebateHistoryService get rebate history
type GetRebateHistoryService struct {
	c            *Client
	subAccountId *string
	startTime    *uint64
	endTime      *uint64
	limit        *int
}

// SubAccountId set subAccountId
func (s *GetRebateHistoryService) SubAccountId(subAccountId string) *GetRebateHistoryService {
	s.subAccountId = &subAccountId
	return s
}

// StartTime set startTime
func (s *GetRebateHistoryService) StartTime(startTime uint64) *GetRebateHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetRebateHistoryService) EndTime(endTime uint64) *GetRebateHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetRebateHistoryService) Limit(limit int) *GetRebateHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetRebateHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/rebate/historicalRecord",
		secType:  secTypeSigned,
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return data, nil
}
