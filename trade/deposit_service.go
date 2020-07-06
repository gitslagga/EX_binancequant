package trade

import (
	"context"
	"encoding/json"
)

// ListDepositsService list deposits
type ListDepositsService struct {
	c         *Client
	coin      *string
	status    *int
	startTime *int64
	endTime   *int64
	offset    *int
	limit     *int
}

// Coin set coin
func (s *ListDepositsService) Coin(coin string) *ListDepositsService {
	s.coin = &coin
	return s
}

// Status set status
func (s *ListDepositsService) Status(status int) *ListDepositsService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *ListDepositsService) StartTime(startTime int64) *ListDepositsService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListDepositsService) EndTime(endTime int64) *ListDepositsService {
	s.endTime = &endTime
	return s
}

// EndTime set endTime
func (s *ListDepositsService) Offset(offset int) *ListDepositsService {
	s.offset = &offset
	return s
}

// EndTime set endTime
func (s *ListDepositsService) Limit(limit int) *ListDepositsService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListDepositsService) Do(ctx context.Context, opts ...RequestOption) (deposits []*Deposit, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/deposit/hisrec",
		secType:  secTypeSigned,
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
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	var res = make([]*Deposit, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// Deposit define deposit info
type Deposit struct {
	Address    string `json:"address"`
	AddressTag string `json:"addressTag"`
	Network    string `json:"network"`
	InsertTime int64  `json:"insertTime"`
	Amount     string `json:"amount"`
	Coin       string `json:"coin"`
	Status     int    `json:"status"`
	TxID       string `json:"txId"`
}

// DepositsAddressService deposits address
type DepositsAddressService struct {
	c       *Client
	coin    *string
	network *string
}

// Coin set coin
func (s *DepositsAddressService) Coin(coin string) *DepositsAddressService {
	s.coin = &coin
	return s
}

// Network set network
func (s *DepositsAddressService) Network(network string) *DepositsAddressService {
	s.network = &network
	return s
}

// Do send request
func (s *DepositsAddressService) Do(ctx context.Context, opts ...RequestOption) (address *DepositAddressResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/deposit/address",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.network != nil {
		r.setParam("network", *s.network)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	var res = new(DepositAddressResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

// DepositAddressResponse define deposit history
type DepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}
