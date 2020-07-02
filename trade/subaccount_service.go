package trade

import (
	"EX_binancequant/trade/common"
	"context"
	"encoding/json"
)

// CreateSubAccountService create broker subAccount
type CreateSubAccountService struct {
	c *Client
}

// Do send request
func (s *CreateSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccount, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubAccount define subAccount info
type SubAccount struct {
	SubAccountId string `json:"subaccountId"`
}

// EnableSubAccountFuturesService enable futures for sub account
type EnableSubAccountFuturesService struct {
	c            *Client
	subAccountId *string
	futures      *bool
}

// SubAccountId set subAccountId
func (s *EnableSubAccountFuturesService) SubAccountId(subAccount string) *EnableSubAccountFuturesService {
	s.subAccountId = &subAccount
	return s
}

// Futures set futures (true, false)
func (s *EnableSubAccountFuturesService) Futures(futures bool) *EnableSubAccountFuturesService {
	s.futures = &futures
	return s
}

// Do send request
func (s *EnableSubAccountFuturesService) Do(ctx context.Context, opts ...RequestOption) (res *EnableSubAccountFutures, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	r.setParam("futures", *s.futures)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableSubAccountFutures)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EnableSubAccountFutures define enableFutures info
type EnableSubAccountFutures struct {
	SubAccountId  string `json:"subaccountId"`
	EnableFutures bool   `json:"enableFutures"`
	UpdateTime    uint64 `json:"updateTime"`
}

// CreateSubAccountApiService create api key for sub account
type CreateSubAccountApiService struct {
	c            *Client
	subAccountId *string
	canTrade     *bool
	futuresTrade *bool
}

// SubAccountId set subAccountId
func (s *CreateSubAccountApiService) SubAccountId(subAccountId string) *CreateSubAccountApiService {
	s.subAccountId = &subAccountId
	return s
}

// CanTrade set canTrade (true, false)
func (s *CreateSubAccountApiService) CanTrade(canTrade bool) *CreateSubAccountApiService {
	s.canTrade = &canTrade
	return s
}

// FuturesTrade set futuresTrade (true, false)
func (s *CreateSubAccountApiService) FuturesTrade(futuresTrade bool) *CreateSubAccountApiService {
	s.futuresTrade = &futuresTrade
	return s
}

// Do send request
func (s *CreateSubAccountApiService) Do(ctx context.Context, opts ...RequestOption) (res *CreateSubAccountApi, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	r.setParam("canTrade", *s.canTrade)
	r.setParam("futuresTrade", *s.futuresTrade)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateSubAccountApi)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateSubAccountApi define create subAccountApi info
type CreateSubAccountApi struct {
	SubAccountId string `json:"subaccountId"`
	ApiKey       string `json:"apikey"`
	SecretKey    string `json:"secretkey,omitempty"`
	CanTrade     bool   `json:"canTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

// DeleteSubAccountApiService delete api key for sub account
type DeleteSubAccountApiService struct {
	c                *Client
	subAccountId     *string
	subAccountApiKey *string
}

// SubAccountId set subAccountId
func (s *DeleteSubAccountApiService) SubAccountId(subAccountId string) *DeleteSubAccountApiService {
	s.subAccountId = &subAccountId
	return s
}

// SubAccountApiKey set subAccountApiKey
func (s *DeleteSubAccountApiService) SubAccountApiKey(subAccountApiKey string) *DeleteSubAccountApiService {
	s.subAccountApiKey = &subAccountApiKey
	return s
}

// Do send request
func (s *DeleteSubAccountApiService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	r.setParam("subAccountApiKey", *s.subAccountApiKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}

// GetSubAccountApiService query sub account api key
type GetSubAccountApiService struct {
	c                *Client
	subAccountId     *string
	subAccountApiKey *string
}

// SubAccountId set subAccountId
func (s *GetSubAccountApiService) SubAccountId(subAccountId string) *GetSubAccountApiService {
	s.subAccountId = &subAccountId
	return s
}

// SubAccountApiKey set subAccountApiKey
func (s *GetSubAccountApiService) SubAccountApiKey(subAccountApiKey string) *GetSubAccountApiService {
	s.subAccountApiKey = &subAccountApiKey
	return s
}

// Do send request
func (s *GetSubAccountApiService) Do(ctx context.Context, opts ...RequestOption) (res []*CreateSubAccountApi, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	if s.subAccountApiKey != nil {
		r.setParam("subAccountApiKey", *s.subAccountApiKey)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*CreateSubAccountApi{}, err
	}
	data = common.ToJSONList(data)
	res = make([]*CreateSubAccountApi, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CreateSubAccountApi{}, err
	}
	return res, nil
}

// ChangeSubAccountApiPermissionService change api permission
type ChangeSubAccountApiPermissionService struct {
	c                *Client
	subAccountId     *string
	subAccountApiKey *string
	canTrade         *bool
	futuresTrade     *bool
}

// SubAccountId set subAccountId
func (s *ChangeSubAccountApiPermissionService) SubAccountId(subAccountId string) *ChangeSubAccountApiPermissionService {
	s.subAccountId = &subAccountId
	return s
}

// SubAccountApiKey set subAccountApiKey
func (s *ChangeSubAccountApiPermissionService) SubAccountApiKey(subAccountApiKey string) *ChangeSubAccountApiPermissionService {
	s.subAccountApiKey = &subAccountApiKey
	return s
}

// CanTrade set canTrade (true, false)
func (s *ChangeSubAccountApiPermissionService) CanTrade(canTrade bool) *ChangeSubAccountApiPermissionService {
	s.canTrade = &canTrade
	return s
}

// FuturesTrade set futuresTrade (true, false)
func (s *ChangeSubAccountApiPermissionService) FuturesTrade(futuresTrade bool) *ChangeSubAccountApiPermissionService {
	s.futuresTrade = &futuresTrade
	return s
}

// Do send request
func (s *ChangeSubAccountApiPermissionService) Do(ctx context.Context, opts ...RequestOption) (res *CreateSubAccountApi, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccountApi/permission",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	r.setParam("subAccountApiKey", *s.subAccountApiKey)
	r.setParam("canTrade", *s.canTrade)
	r.setParam("futuresTrade", *s.futuresTrade)
	r.setParam("marginTrade", false)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateSubAccountApi)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSubAccountService query sub account
type GetSubAccountService struct {
	c            *Client
	subAccountId *string
}

// SubAccountId set subAccountId
func (s *GetSubAccountService) SubAccountId(subAccountId string) *GetSubAccountService {
	s.subAccountId = &subAccountId
	return s
}

// Do send request
func (s *GetSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res []*GetSubAccount, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*GetSubAccount, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSubAccount define query subAccountApi info
type GetSubAccount struct {
	SubAccountId    string `json:"subaccountId"`
	MakerCommission int    `json:"makerCommission"`
	TakerCommission int    `json:"takerCommission"`
	CreateTime      uint64 `json:"createTime"`
}

// ChangeCommissionFuturesService change futures commission
type ChangeCommissionFuturesService struct {
	c               *Client
	subAccountId    *string
	symbol          *string
	makerAdjustment *int
	takerAdjustment *int
}

// SubAccountId set subAccountId
func (s *ChangeCommissionFuturesService) SubAccountId(subAccountId string) *ChangeCommissionFuturesService {
	s.subAccountId = &subAccountId
	return s
}

// Symbol set symbol
func (s *ChangeCommissionFuturesService) Symbol(symbol string) *ChangeCommissionFuturesService {
	s.symbol = &symbol
	return s
}

// MakerAdjustment set makerAdjustment
func (s *ChangeCommissionFuturesService) MakerAdjustment(makerAdjustment int) *ChangeCommissionFuturesService {
	s.makerAdjustment = &makerAdjustment
	return s
}

// TakerAdjustment set takerAdjustment
func (s *ChangeCommissionFuturesService) TakerAdjustment(takerAdjustment int) *ChangeCommissionFuturesService {
	s.takerAdjustment = &takerAdjustment
	return s
}

// Do send request
func (s *ChangeCommissionFuturesService) Do(ctx context.Context, opts ...RequestOption) (res *ChangeCommissionFutures, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccountApi/commission/futures",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	r.setParam("symbol", *s.symbol)
	r.setParam("makerAdjustment", *s.makerAdjustment)
	r.setParam("takerAdjustment", *s.takerAdjustment)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ChangeCommissionFutures)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ChangeCommissionFutures define change futures commission
type ChangeCommissionFutures struct {
	SubAccountId    string `json:"subaccountId"`
	Symbol          string `json:"symbol"`
	MakerAdjustment int    `json:"makerAdjustment"`
	TakerAdjustment int    `json:"takerAdjustment"`
	MakerCommission int    `json:"makerCommission"`
	TakerCommission int    `json:"takerCommission"`
}

// GetCommissionFuturesService query futures commission
type GetCommissionFuturesService struct {
	c            *Client
	subAccountId *string
	symbol       *string
}

// SubAccountId set subAccountId
func (s *GetCommissionFuturesService) SubAccountId(subAccountId string) *GetCommissionFuturesService {
	s.subAccountId = &subAccountId
	return s
}

// Symbol set symbol
func (s *GetCommissionFuturesService) Symbol(symbol string) *GetCommissionFuturesService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *GetCommissionFuturesService) Do(ctx context.Context, opts ...RequestOption) (res []*ChangeCommissionFutures, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccountApi/commission/futures",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*ChangeCommissionFutures{}, err
	}
	data = common.ToJSONList(data)
	res = make([]*ChangeCommissionFutures, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*ChangeCommissionFutures{}, err
	}
	return res, nil
}
