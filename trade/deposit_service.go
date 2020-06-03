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
func (s *ListDepositsService) Do(ctx context.Context, opts ...RequestOption) (deposits *ResponseDeposit, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/deposit/hisrec",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.coin != nil {
		m["coin"] = *s.coin
	}
	if s.status != nil {
		m["status"] = *s.status
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.endTime != nil {
		m["offset"] = *s.offset
	}
	if s.endTime != nil {
		m["limit"] = *s.limit
	}
	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	var res = new(ResponseDeposit)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

type ResponseDeposit []Deposit

// Deposit define deposit info
type Deposit struct {
	Address    string  `json:"address"`
	AddressTag string  `json:"addressTag"`
	Network    string  `json:"network"`
	InsertTime int64   `json:"insertTime"`
	Amount     float64 `json:"amount"`
	Coin       string  `json:"coin"`
	Status     int     `json:"status"`
	TxID       string  `json:"txId"`
}
