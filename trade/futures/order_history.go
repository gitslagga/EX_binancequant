package futures

import (
	"context"
	"encoding/json"
)

// GetTradeHistoryService get order history after fromId
type GetTradeHistoryService struct {
	c         *Client
	symbol    string
	fromId    *uint64
	startTime *int64
	endTime   *int64
	limit     *int64
}

// Symbol set symbol
func (s *GetTradeHistoryService) Symbol(symbol string) *GetTradeHistoryService {
	s.symbol = symbol
	return s
}

// Order history set after fromId
func (s *GetTradeHistoryService) FromId(fromId uint64) *GetTradeHistoryService {
	s.fromId = &fromId
	return s
}

// StartTime set startTime
func (s *GetTradeHistoryService) StartTime(startTime int64) *GetTradeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetTradeHistoryService) EndTime(endTime int64) *GetTradeHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetTradeHistoryService) Limit(limit int64) *GetTradeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetTradeHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *[]TradeHistory, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/userTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.fromId != nil {
		r.setParam("fromId", s.fromId)
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
	res = new([]TradeHistory)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TradeHistory define order history after fromId
type TradeHistory struct {
	Buyer           bool   `json:"buyer"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Id              uint64 `json:"id"`
	Maker           bool   `json:"maker"`
	OrderId         uint64 `json:"orderId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	RealizedPnl     string `json:"realizedPnl"`
	Side            string `json:"side"`
	PositionSide    string `json:"positionSide"`
	Symbol          string `json:"symbol"`
	Time            uint64 `json:"time"`
}
