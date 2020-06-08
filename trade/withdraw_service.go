package trade

import (
	"context"
	"encoding/json"
)

// CreateWithdrawService create withdraw
type CreateWithdrawService struct {
	c                  *Client
	coin               string
	address            string
	amount             float64
	network            *string
	addressTag         *string
	withdrawOrderId    *string
	transactionFeeFlag *bool
	name               *string
}

// Coin set coin
func (s *CreateWithdrawService) Coin(coin string) *CreateWithdrawService {
	s.coin = coin
	return s
}

// Address set address
func (s *CreateWithdrawService) Address(address string) *CreateWithdrawService {
	s.address = address
	return s
}

// Amount set amount
func (s *CreateWithdrawService) Amount(amount float64) *CreateWithdrawService {
	s.amount = amount
	return s
}

// Network set network
func (s *CreateWithdrawService) Network(network string) *CreateWithdrawService {
	s.network = &network
	return s
}

// AddressTag set addressTag
func (s *CreateWithdrawService) AddressTag(addressTag string) *CreateWithdrawService {
	s.addressTag = &addressTag
	return s
}

// WithdrawOrderId set withdrawOrderId
func (s *CreateWithdrawService) WithdrawOrderId(withdrawOrderId string) *CreateWithdrawService {
	s.withdrawOrderId = &withdrawOrderId
	return s
}

// TransactionFeeFlag set transactionFeeFlag
func (s *CreateWithdrawService) TransactionFeeFlag(transactionFeeFlag bool) *CreateWithdrawService {
	s.transactionFeeFlag = &transactionFeeFlag
	return s
}

// Name set name
func (s *CreateWithdrawService) Name(name string) *CreateWithdrawService {
	s.name = &name
	return s
}

// Do send request
func (s *CreateWithdrawService) Do(ctx context.Context) (err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/capital/withdraw/apply",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)
	if s.withdrawOrderId != nil {
		r.setParam("withdrawOrderId", *s.withdrawOrderId)
	}
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	if s.addressTag != nil {
		r.setParam("addressTag", *s.addressTag)
	}
	if s.transactionFeeFlag != nil {
		r.setParam("transactionFeeFlag", *s.transactionFeeFlag)
	}
	if s.name != nil {
		r.setParam("name", *s.name)
	}
	_, err = s.c.callAPI(ctx, r)
	return err
}

// ListWithdrawsService list withdraws
type ListWithdrawsService struct {
	c         *Client
	coin      *string
	status    *int
	offset    *int
	limit     *int
	startTime *int64
	endTime   *int64
}

// Coin set coin
func (s *ListWithdrawsService) Coin(coin string) *ListWithdrawsService {
	s.coin = &coin
	return s
}

// Status set status
func (s *ListWithdrawsService) Status(status int) *ListWithdrawsService {
	s.status = &status
	return s
}

// Offset set offset
func (s *ListWithdrawsService) Offset(offset int) *ListWithdrawsService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *ListWithdrawsService) Limit(limit int) *ListWithdrawsService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *ListWithdrawsService) StartTime(startTime int64) *ListWithdrawsService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListWithdrawsService) EndTime(endTime int64) *ListWithdrawsService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ListWithdrawsService) Do(ctx context.Context) (withdraws *WithdrawHistoryResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/withdraw/history",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(WithdrawHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

// WithdrawHistoryResponse define withdraw history response
type WithdrawHistoryResponse []Withdraw

// Withdraw define withdraw info
type Withdraw struct {
	Amount          float64 `json:"amount"`
	Address         string  `json:"address"`
	Coin            string  `json:"coin"`
	TxID            string  `json:"txId"`
	ApplyTime       int64   `json:"applyTime"`
	Status          int     `json:"status"`
	Id              string  `json:"id"`
	WithdrawOrderId string  `json:"withdrawOrderId"`
	Network         string  `json:"network"`
	TransferType    string  `json:"transferType"`
}

// Discard
// GetWithdrawFeeService get withdraw fee
type GetWithdrawFeeService struct {
	c    *Client
	coin string
}

// Coin set coin
func (s *GetWithdrawFeeService) Coin(coin string) *GetWithdrawFeeService {
	s.coin = coin
	return s
}

// Do send request
func (s *GetWithdrawFeeService) Do(ctx context.Context, opts ...RequestOption) (res *WithdrawFee, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/withdrawFee.html",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(WithdrawFee)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// WithdrawFee withdraw fee
type WithdrawFee struct {
	Fee float64 `json:"withdrawFee"` // docs specify string value but api returns decimal
}
