package trade

import (
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

// SubAccountApiKey set subAccountApiKey (true, false)
func (s *GetSubAccountApiService) SubAccountApiKey(subAccountApiKey string) *GetSubAccountApiService {
	s.subAccountApiKey = &subAccountApiKey
	return s
}

// Do send request
func (s *GetSubAccountApiService) Do(ctx context.Context, opts ...RequestOption) (res *CreateSubAccountApi, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", *s.subAccountId)
	r.setParam("subAccountApiKey", *s.subAccountApiKey)
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
